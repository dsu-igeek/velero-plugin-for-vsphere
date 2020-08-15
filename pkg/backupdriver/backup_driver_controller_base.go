/*
Copyright 2020 the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package backupdriver

import (
	"context"
	"time"

	"k8s.io/client-go/rest"

	"github.com/sirupsen/logrus"
	backupdriverapi "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/apis/backupdriver/v1"
	backupdriverclientset "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/generated/clientset/versioned/typed/backupdriver/v1"
	backupdriverinformers "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/generated/informers/externalversions"
	backupdriverlisters "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/generated/listers/backupdriver/v1"
	"github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/snapshotmgr"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

// BackupDriverController is the interface for backup driver controller
type BackupDriverController interface {
	// Run starts the controller.
	Run(ctx context.Context, workers int)
}

type backupDriverController struct {
	name   string
	logger logrus.FieldLogger

	// Supervisor Cluster KubeClient for Guest Cluster
	svcKubeConfig *rest.Config

	backupdriverClient *backupdriverclientset.BackupdriverV1Client

	// Supervisor Cluster namespace
	supervisorNamespace string

	rateLimiter workqueue.RateLimiter

	// PV Lister
	pvLister corelisters.PersistentVolumeLister
	// PV Synced
	pvSynced cache.InformerSynced

	// PVC Lister
	pvcLister corelisters.PersistentVolumeClaimLister
	// PVC Synced
	pvcSynced cache.InformerSynced
	// claim queue
	claimQueue workqueue.RateLimitingInterface

	// Supervisor Cluster PVC Lister
	svcPVCLister corelisters.PersistentVolumeClaimLister
	// Supervisor Cluster PVC Synced
	svcPVCSynced cache.InformerSynced

	// BackupRepositoryClaim Lister
	backupRepositoryClaimLister backupdriverlisters.BackupRepositoryClaimLister
	// BackupRepositoryClaim Synced
	backupRepositoryClaimSynced cache.InformerSynced
	// backupRepositoryClaim queue
	backupRepositoryClaimQueue workqueue.RateLimitingInterface

	// Supervisor Cluster BackupRepositoryClaim Lister
	svcBackupRepositoryClaimLister backupdriverlisters.BackupRepositoryClaimLister
	// Supervisor Cluster BackupRepositoryClaim Synced
	svcBackupRepositoryClaimSynced cache.InformerSynced

	// BackupRepository Lister
	backupRepositoryLister backupdriverlisters.BackupRepositoryLister
	// BackupRepository Synced
	backupRepositorySynced cache.InformerSynced

	// Snapshot queue
	snapshotQueue workqueue.RateLimitingInterface
	// Snapshot Lister
	snapshotLister backupdriverlisters.SnapshotLister
	// Snapshot Synced
	snapshotSynced cache.InformerSynced

	// Supervisor Cluster Snapshot Lister
	svcSnapshotLister backupdriverlisters.SnapshotLister
	// Supervisor Cluster Snapshot Synced
	svcSnapshotSynced cache.InformerSynced

	// CloneFromSnapshot queue
	cloneFromSnapshotQueue workqueue.RateLimitingInterface
	// CloneFromSnapshot Lister
	cloneFromSnapshotLister backupdriverlisters.CloneFromSnapshotLister
	// CloneFromSnapshot Synced
	cloneFromSnapshotSynced cache.InformerSynced

	// Supervisor Cluster CloneFromSnapshot Lister
	svcCloneFromSnapshotLister backupdriverlisters.CloneFromSnapshotLister
	// Supervisor Cluster CloneFromSnapshot Synced
	svcCloneFromSnapshotSynced cache.InformerSynced

	// Snapshot queue
	deleteSnapshotQueue workqueue.RateLimitingInterface
	// Snapshot Lister
	deleteSnapshotLister backupdriverlisters.DeleteSnapshotLister
	// Snapshot Synced
	deleteSnapshotSynced cache.InformerSynced

	// Snapshot manager
	snapManager *snapshotmgr.SnapshotManager
}

// NewBackupDriverController returns a BackupDriverController.
func NewBackupDriverController(
	name string,
	logger logrus.FieldLogger,
	backupdriverClient *backupdriverclientset.BackupdriverV1Client,
	svcKubeConfig *rest.Config,
	supervisorNamespace string,
	resyncPeriod time.Duration,
	informerFactory informers.SharedInformerFactory,
	backupdriverInformerFactory backupdriverinformers.SharedInformerFactory,
	snapManager *snapshotmgr.SnapshotManager,
	rateLimiter workqueue.RateLimiter) BackupDriverController {
	// TODO: Fix svcPVCInformer to use svcInformerFactory
	svcPVCInformer := informerFactory.Core().V1().PersistentVolumeClaims()
	//svcPVCInformer := svcInformerFactory.Core().V1().PersistentVolumeClaims()
	pvcInformer := informerFactory.Core().V1().PersistentVolumeClaims()
	pvInformer := informerFactory.Core().V1().PersistentVolumes()

	backupRepositoryInformer := backupdriverInformerFactory.Backupdriver().V1().BackupRepositories()

	backupRepositoryClaimInformer := backupdriverInformerFactory.Backupdriver().V1().BackupRepositoryClaims()
	// TODO: Use svcBackupdriverInformerFactory for svcBackupRepositoryClaimInformer
	svcBackupRepositoryClaimInformer := backupdriverInformerFactory.Backupdriver().V1().BackupRepositoryClaims()

	snapshotInformer := backupdriverInformerFactory.Backupdriver().V1().Snapshots()
	// TODO: Use svcBackupdriverInformerFactory for svcSnapshotInformer
	svcSnapshotInformer := backupdriverInformerFactory.Backupdriver().V1().Snapshots()

	cloneFromSnapshotInformer := backupdriverInformerFactory.Backupdriver().V1().CloneFromSnapshots()
	// TODO: Use svcBackupdriverInformerFactor for svcCloneFromSnapshotInformer
	svcCloneFromSnapshotInformer := backupdriverInformerFactory.Backupdriver().V1().CloneFromSnapshots()

	deleteSnapshotInformer := backupdriverInformerFactory.Backupdriver().V1().DeleteSnapshots()

	claimQueue := workqueue.NewNamedRateLimitingQueue(
		rateLimiter, "backup-driver-claim-queue")
	snapshotQueue := workqueue.NewNamedRateLimitingQueue(
		rateLimiter, "backup-driver-snapshot-queue")
	cloneFromSnapshotQueue := workqueue.NewNamedRateLimitingQueue(
		rateLimiter, "backup-driver-clone-queue")
	backupRepositoryClaimQueue := workqueue.NewNamedRateLimitingQueue(
		rateLimiter, "backup-driver-brc-queue")
	deleteSnapshotQueue := workqueue.NewNamedRateLimitingQueue(
		rateLimiter, "backup-driver-delete-snapshot-queue")

	ctrl := &backupDriverController{
		name:                           name,
		logger:                         logger.WithField("controller", name),
		svcKubeConfig:                  svcKubeConfig,
		backupdriverClient:             backupdriverClient,
		supervisorNamespace:            supervisorNamespace,
		pvLister:                       pvInformer.Lister(),
		pvSynced:                       pvInformer.Informer().HasSynced,
		pvcLister:                      pvcInformer.Lister(),
		pvcSynced:                      pvcInformer.Informer().HasSynced,
		claimQueue:                     claimQueue,
		svcPVCLister:                   svcPVCInformer.Lister(),
		svcPVCSynced:                   svcPVCInformer.Informer().HasSynced,
		backupRepositoryClaimLister:    backupRepositoryClaimInformer.Lister(),
		backupRepositoryClaimSynced:    backupRepositoryClaimInformer.Informer().HasSynced,
		backupRepositoryClaimQueue:     backupRepositoryClaimQueue,
		svcBackupRepositoryClaimLister: svcBackupRepositoryClaimInformer.Lister(),
		svcBackupRepositoryClaimSynced: svcBackupRepositoryClaimInformer.Informer().HasSynced,
		backupRepositoryLister:         backupRepositoryInformer.Lister(),
		backupRepositorySynced:         backupRepositoryInformer.Informer().HasSynced,
		snapshotQueue:                  snapshotQueue,
		snapshotLister:                 snapshotInformer.Lister(),
		snapshotSynced:                 snapshotInformer.Informer().HasSynced,
		svcSnapshotLister:              svcSnapshotInformer.Lister(),
		svcSnapshotSynced:              svcSnapshotInformer.Informer().HasSynced,
		cloneFromSnapshotQueue:         cloneFromSnapshotQueue,
		cloneFromSnapshotLister:        cloneFromSnapshotInformer.Lister(),
		cloneFromSnapshotSynced:        cloneFromSnapshotInformer.Informer().HasSynced,
		svcCloneFromSnapshotLister:     svcCloneFromSnapshotInformer.Lister(),
		svcCloneFromSnapshotSynced:     svcCloneFromSnapshotInformer.Informer().HasSynced,
		deleteSnapshotQueue:            deleteSnapshotQueue,
		deleteSnapshotLister:           deleteSnapshotInformer.Lister(),
		deleteSnapshotSynced:           deleteSnapshotInformer.Informer().HasSynced,
		snapManager:                    snapManager,
	}

	pvcInformer.Informer().AddEventHandlerWithResyncPeriod(
		cache.ResourceEventHandlerFuncs{
			//AddFunc:    func(obj interface{}) { ctrl.enqueueClaim(obj) },
			//UpdateFunc: func(oldObj, newObj interface{}) { ctrl.enqueueClaim(newObj) },
			//DeleteFunc: func(obj interface{}) { ctrl.delClaim(obj) },
		},
		resyncPeriod,
	)

	svcPVCInformer.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{
		//AddFunc:    rc.svcAddPVC,
		//UpdateFunc: rc.svcUpdatePVC,
		//DeleteFunc: rc.svcDeletePVC,
	}, resyncPeriod)

	pvInformer.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{
		//AddFunc:    rc.addPV,
		//UpdateFunc: rc.updatePV,
		//DeleteFunc: rc.deletePV,
	}, resyncPeriod)

	snapshotInformer.Informer().AddEventHandlerWithResyncPeriod(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    func(obj interface{}) { ctrl.enqueueSnapshot(obj) },
			UpdateFunc: func(_, obj interface{}) { ctrl.enqueueSnapshot(obj) },
			//DeleteFunc: func(obj interface{}) { ctrl.delSnapshot(obj) },
		},
		resyncPeriod,
	)

	svcSnapshotInformer.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{
		//AddFunc:    rc.svcAddSnapshot,
		//UpdateFunc: rc.svcUpdateSnapshot,
		//DeleteFunc: rc.svcDeleteSnapshot,
	}, resyncPeriod)

	cloneFromSnapshotInformer.Informer().AddEventHandlerWithResyncPeriod(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) { ctrl.enqueueCloneFromSnapshot(obj) },
			// UpdateFunc is not needed now. When adding it backup, make sure no multiple CloneFromSnapshot CR's created for the same request
			//UpdateFunc: func(oldObj, newObj interface{}) { ctrl.enqueueCloneFromSnapshot(newObj) },
			//DeleteFunc: func(obj interface{}) { ctrl.delCloneFromSnapshot(obj) },
		},
		resyncPeriod,
	)

	svcCloneFromSnapshotInformer.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{
		//AddFunc:    rc.svcAddCloneFromSnapshot,
		//UpdateFunc: rc.svcUpdateCloneFromSnapshot,
		//DeleteFunc: rc.svcDeleteCloneFromSnapshot,
	}, resyncPeriod)

	backupRepositoryInformer.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{
		//AddFunc:    rc.addBackupRepository,
		//UpdateFunc: rc.updateBackupRepository,
		//DeleteFunc: rc.deleteBackupRepository,
	}, resyncPeriod)

	backupRepositoryClaimInformer.Informer().AddEventHandlerWithResyncPeriod(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) { ctrl.enqueueBackupRepositoryClaim(obj) },
			//UpdateFunc: func(oldObj, newObj interface{}) { ctrl.enqueueBackupRepositoryClaim(newObj) },
			DeleteFunc: func(obj interface{}) { ctrl.dequeBackupRepositoryClaim(obj) },
		},
		resyncPeriod,
	)

	svcBackupRepositoryClaimInformer.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{
		//AddFunc:    rc.svcAddBackupRepositoryClaim,
		//UpdateFunc: rc.svcUpdateBackupRepositoryClaim,
		//DeleteFunc: rc.svcDeleteBackupRepositoryClaim,
	}, resyncPeriod)

	deleteSnapshotInformer.Informer().AddEventHandlerWithResyncPeriod(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    func(obj interface{}) { ctrl.enqueueDeleteSnapshot(obj) },
			UpdateFunc: func(oldObj, newObj interface{}) { ctrl.enqueueDeleteSnapshot(newObj) },
			DeleteFunc: func(obj interface{}) { ctrl.dequeDeleteSnapshot(obj) },
		},
		resyncPeriod,
	)

	return ctrl
}

// getKey helps to get the resource name from resource object
func (ctrl *backupDriverController) getKey(obj interface{}) (string, error) {
	if unknown, ok := obj.(cache.DeletedFinalStateUnknown); ok && unknown.Obj != nil {
		obj = unknown.Obj
	}
	objKey, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		ctrl.logger.Errorf("Failed to get key from object: %v", err)
		return "", err
	}
	ctrl.logger.Debugf("getKey: key %s", objKey)
	return objKey, nil
}

// Run starts the controller.
func (ctrl *backupDriverController) Run(
	ctx context.Context, workers int) {
	defer ctrl.claimQueue.ShutDown()
	defer ctrl.snapshotQueue.ShutDown()
	defer ctrl.cloneFromSnapshotQueue.ShutDown()
	defer ctrl.deleteSnapshotQueue.ShutDown()

	ctrl.logger.Infof("Starting backup driver controller")
	defer ctrl.logger.Infof("Shutting down backup driver controller")

	stopCh := ctx.Done()

	ctrl.logger.Infof("Waiting for caches to sync")
	if !cache.WaitForCacheSync(stopCh, ctrl.pvSynced, ctrl.pvcSynced, ctrl.svcPVCSynced, ctrl.backupRepositorySynced, ctrl.backupRepositoryClaimSynced, ctrl.svcBackupRepositoryClaimSynced, ctrl.snapshotSynced, ctrl.svcSnapshotSynced, ctrl.cloneFromSnapshotSynced, ctrl.svcCloneFromSnapshotSynced) {
		ctrl.logger.Errorf("Cannot sync caches")
		return
	}
	ctrl.logger.Infof("Caches are synced")

	for i := 0; i < workers; i++ {
		//go wait.Until(ctrl.pvcWorker, 0, stopCh)
		//go wait.Until(ctrl.pvWorker, 0, stopCh)
		//go wait.Until(ctrl.svcPvcWorker, 0, stopCh)
		go wait.Until(ctrl.snapshotWorker, 0, stopCh)
		//go wait.Until(ctrl.svcSnapshotWorker, 0, stopCh)
		go wait.Until(ctrl.cloneFromSnapshotWorker, 0, stopCh)
		//go wait.Until(ctrl.svcCloneFromSnapshotWorker, 0, stopCh)
		go wait.Until(ctrl.backupRepositoryClaimWorker, 0, stopCh)
		go wait.Until(ctrl.deleteSnapshotWorker, 0, stopCh)
	}

	<-stopCh
}

// snapshotWorker is the main worker for snapshot request.
func (ctrl *backupDriverController) snapshotWorker() {
	ctrl.logger.Infof("snapshotWorker: Enter snapshotWorker")

	key, quit := ctrl.snapshotQueue.Get()
	if quit {
		return
	}
	defer ctrl.snapshotQueue.Done(key)

	if err := ctrl.syncSnapshotByKey(key.(string)); err != nil {
		// Put snapshot back to the queue so that we can retry later.
		ctrl.snapshotQueue.AddRateLimited(key)
	} else {
		ctrl.snapshotQueue.Forget(key)
	}
}

// syncSnapshotByKey processes one Snapshot CRD
func (ctrl *backupDriverController) syncSnapshotByKey(key string) error {
	ctrl.logger.Infof("syncSnapshotByKey: Started Snapshot processing %s", key)

	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		ctrl.logger.Errorf("Split meta namespace key of snapshot %s failed: %v", key, err)
		return err
	}

	// Always retrieve up-to-date snapshot CR from API server
	snapshot, err := ctrl.backupdriverClient.Snapshots(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		if k8serrors.IsNotFound(err) {
			ctrl.logger.Infof("Snapshot %s/%s is deleted, no need to process it", namespace, name)
			return nil
		}
		ctrl.logger.Errorf("Get Snapshot %s/%s failed: %v", namespace, name, err)
		return err
	}

	if snapshot.Status.Phase != backupdriverapi.SnapshotPhaseNew {
		ctrl.logger.Debugf("Skipping snapshot, %v, which is not in New phase. Current phase: %v", key, snapshot.Status.Phase)
		return nil
	}

	if snapshot.Spec.SnapshotCancel == false {
		ctrl.logger.Infof("syncSnapshotByKey: calling CreateSnapshot %s/%s", snapshot.Namespace, snapshot.Name)
		return ctrl.createSnapshot(snapshot)
	}

	return nil
}

// enqueueSnapshot adds Snapshotto given work queue.
func (ctrl *backupDriverController) enqueueSnapshot(obj interface{}) {
	ctrl.logger.Infof("enqueueSnapshot: %+v", obj)

	// Beware of "xxx deleted" events
	if unknown, ok := obj.(cache.DeletedFinalStateUnknown); ok && unknown.Obj != nil {
		obj = unknown.Obj
	}
	if snapshot, ok := obj.(*backupdriverapi.Snapshot); ok {
		objName, err := cache.DeletionHandlingMetaNamespaceKeyFunc(snapshot)
		if err != nil {
			ctrl.logger.Errorf("failed to get key from object: %v, %v", err, snapshot)
			return
		}
		ctrl.logger.Infof("enqueueSnapshot: enqueued %q for sync", objName)
		ctrl.snapshotQueue.Add(objName)
	}
}

func (ctrl *backupDriverController) pvcWorker() {
}

func (ctrl *backupDriverController) svcPvcWorker() {
}

func (ctrl *backupDriverController) pvWorker() {
}

func (ctrl *backupDriverController) svcSnapshotWorker() {
}

// cloneFromSnapshotWorker is the main worker for restore request.
func (ctrl *backupDriverController) cloneFromSnapshotWorker() {
	ctrl.logger.Infof("cloneFromSnapshotWorker: Enter cloneFromSnapshotWorker")

	key, quit := ctrl.cloneFromSnapshotQueue.Get()
	if quit {
		return
	}
	defer ctrl.cloneFromSnapshotQueue.Done(key)

	if err := ctrl.syncCloneFromSnapshotByKey(key.(string)); err != nil {
		// Put cloneFromSnapshot back to the queue so that we can retry later.
		ctrl.cloneFromSnapshotQueue.AddRateLimited(key)
	} else {
		ctrl.cloneFromSnapshotQueue.Forget(key)
	}
}

// syncCloneFromSnapshotByKey processes one CloneFromSnapshot CRD
func (ctrl *backupDriverController) syncCloneFromSnapshotByKey(key string) error {
	ctrl.logger.Infof("syncCloneFromSnapshotByKey: Started CloneFromSnapshot processing %s", key)

	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		ctrl.logger.Errorf("Split meta namespace key of CloneFromSnapshot %s failed: %v", key, err)
		return err
	}

	cloneFromSnapshot, err := ctrl.backupdriverClient.CloneFromSnapshots(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		if k8serrors.IsNotFound(err) {
			ctrl.logger.Infof("CloneFromSnapshot %s/%s is deleted, no need to process it", namespace, name)
			return nil
		}
		ctrl.logger.Errorf("Get CloneFromSnapshot %s/%s failed: %v", namespace, name, err)
		return err
	}

	switch phase := cloneFromSnapshot.Status.Phase; phase {
	case backupdriverapi.ClonePhaseInProgress:
	case backupdriverapi.ClonePhaseCompleted:
	case backupdriverapi.ClonePhaseFailed:
	case backupdriverapi.ClonePhaseCanceling:
	case backupdriverapi.ClonePhaseCanceled:
		ctrl.logger.Infof("Skipping cloneFromSnapshot %s which is not in New or Retry phase. Current phase: %v", key, cloneFromSnapshot.Status.Phase)
		return nil
	default:
		// If phase is ClonePhaseNew or ClonePhaseRetry,
		// or if phase is not initialized, proceed
		// to clone from snapshot
	}

	if cloneFromSnapshot.Spec.CloneCancel == false {
		ctrl.logger.Infof("syncCloneFromSnapshotByKey: calling CloneFromSnapshot %s/%s", cloneFromSnapshot.Namespace, cloneFromSnapshot.Name)
		err := ctrl.cloneFromSnapshot(cloneFromSnapshot)
		if err != nil {
			ctrl.logger.Errorf("cloneFromSnapshot %s/%s failed: %v", namespace, name, err)
			return err
		}
	}

	return nil
}

// enqueueCloneFromSnapshot adds CloneFromSnapshotto given work queue.
func (ctrl *backupDriverController) enqueueCloneFromSnapshot(obj interface{}) {
	ctrl.logger.Infof("enqueueCloneFromSnapshot: %+v", obj)

	// Beware of "xxx deleted" events
	if unknown, ok := obj.(cache.DeletedFinalStateUnknown); ok && unknown.Obj != nil {
		obj = unknown.Obj
	}
	if cloneFromSnapshot, ok := obj.(*backupdriverapi.CloneFromSnapshot); ok {
		objName, err := cache.DeletionHandlingMetaNamespaceKeyFunc(cloneFromSnapshot)
		if err != nil {
			ctrl.logger.Errorf("failed to get key from object: %v, %v", err, cloneFromSnapshot)
			return
		}
		ctrl.logger.Infof("enqueueCloneFromSnapshot: enqueued %q for sync", objName)
		ctrl.cloneFromSnapshotQueue.Add(objName)
	}
}

func (ctrl *backupDriverController) svcCloneFromSnapshotWorker() {
}

func (ctrl *backupDriverController) backupRepositoryClaimWorker() {
	ctrl.logger.Debugf("backupRepositoryClaim: Enter backupRepositoryClaimWorker")

	key, quit := ctrl.backupRepositoryClaimQueue.Get()
	if quit {
		return
	}
	defer ctrl.backupRepositoryClaimQueue.Done(key)

	if err := ctrl.syncBackupRepositoryClaimByKey(key.(string)); err != nil {
		// Put backuprepositoryclaim back to the queue so that we can retry later.
		ctrl.backupRepositoryClaimQueue.AddRateLimited(key)
	} else {
		ctrl.backupRepositoryClaimQueue.Forget(key)
	}

}

// syncBackupRepositoryClaim processes one BackupRepositoryClaim CRD
func (ctrl *backupDriverController) syncBackupRepositoryClaimByKey(key string) error {
	ctrl.logger.Infof("syncBackupRepositoryClaimByKey: Started BackupRepositoryClaim processing %s", key)

	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		ctrl.logger.Errorf("Split meta namespace key of backupRepositoryClaim %s failed: %v", key, err)
		return err
	}

	brc, err := ctrl.backupRepositoryClaimLister.BackupRepositoryClaims(namespace).Get(name)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			ctrl.logger.Infof("syncBackupRepositoryClaimByKey: BackupRepositoryClaim %s/%s is deleted, no need to process it", namespace, name)
			return nil
		}
		ctrl.logger.Errorf("Get BackupRepositoryClaim %s/%s failed: %v", namespace, name, err)
		return err
	}

	if brc.ObjectMeta.DeletionTimestamp == nil {
		ctx := context.Background()
		var svcBackupRepositoryName string
		// In case of guest clusters, create BackupRepositoryClaim in the supervisor namespace
		if ctrl.svcKubeConfig != nil {
			svcBackupRepositoryName, err = ClaimSvcBackupRepository(ctx, brc, ctrl.svcKubeConfig, ctrl.supervisorNamespace, ctrl.logger)
			if err != nil {
				ctrl.logger.Errorf("Failed to create Supervisor BackupRepositoryClaim")
				return err
			}
			ctrl.logger.Infof("Created Supervisor BackupRepositoryClaim with BackupRepository %s", svcBackupRepositoryName)
		}

		// Create BackupRepository when a new BackupRepositoryClaim is added
		// Save the supervisor backup repository name to be passed to snapshot manager
		ctrl.logger.Infof("syncBackupRepositoryClaimByKey: Create BackupRepository for BackupRepositoryClaim %s/%s", brc.Namespace, brc.Name)
		// Create BackupRepository when a new BackupRepositoryClaim is added and if the BackupRepository is not already created
		br, err := CreateBackupRepository(ctx, brc, svcBackupRepositoryName, ctrl.backupdriverClient, ctrl.logger)
		if err != nil {
			ctrl.logger.Errorf("Failed to create BackupRepository")
			return err
		}

		err = PatchBackupRepositoryClaim(brc, br.Name, brc.Namespace, ctrl.backupdriverClient)
		if err != nil {
			return err
		}
	}

	return nil
}

// enqueueBackupRepositoryClaimWork adds BackupRepositoryClaim to given work queue.
func (ctrl *backupDriverController) enqueueBackupRepositoryClaim(obj interface{}) {
	ctrl.logger.Debugf("enqueueBackupRepositoryClaim: %+v", obj)

	// Beware of "xxx deleted" events
	if unknown, ok := obj.(cache.DeletedFinalStateUnknown); ok && unknown.Obj != nil {
		obj = unknown.Obj
	}
	if brc, ok := obj.(*backupdriverapi.BackupRepositoryClaim); ok {
		objName, err := cache.DeletionHandlingMetaNamespaceKeyFunc(brc)
		if err != nil {
			ctrl.logger.Errorf("failed to get key from object: %v, %v", err, brc)
			return
		}
		ctrl.logger.Infof("enqueueBackupRepositoryClaim: enqueued %q for sync", objName)
		ctrl.backupRepositoryClaimQueue.Add(objName)
	}
}

func (ctrl *backupDriverController) dequeBackupRepositoryClaim(obj interface{}) {
	ctrl.logger.Debugf("dequeBackupRepositoryClaim: Remove BackupRepositoryClaim %v from queue", obj)

	var key string
	var err error
	if key, err = cache.DeletionHandlingMetaNamespaceKeyFunc(obj); err != nil {
		return
	}

	brc, ok := obj.(*backupdriverapi.BackupRepositoryClaim)
	if !ok || brc == nil {
		return
	}
	// Delete BackupRepository from API server
	if brc.BackupRepository != "" {
		ctrl.logger.Infof("dequeBackupRepositoryClaim: Delete BackupRepository %s for BackupRepositoryClaim %s/%s", brc.BackupRepository, brc.Namespace, brc.Name)
		err = ctrl.backupdriverClient.BackupRepositories().Delete(brc.BackupRepository, &metav1.DeleteOptions{})
		if err != nil && !k8serrors.IsNotFound(err) {
			ctrl.logger.Errorf("Delete BackupRepository %s failed: %v", brc.BackupRepository, err)
			return
		}
	}

	ctrl.backupRepositoryClaimQueue.Forget(key)
	ctrl.backupRepositoryClaimQueue.Done(key)
}

func (ctrl *backupDriverController) deleteSnapshotWorker() {
	ctrl.logger.Infof("deleteSnapshotWorker: Enter deleteSnapshotWorker")

	key, quit := ctrl.deleteSnapshotQueue.Get()
	if quit {
		return
	}
	defer ctrl.deleteSnapshotQueue.Done(key)

	if err := ctrl.syncDeleteSnapshotByKey(key.(string)); err != nil {
		// Put deleteSnapshot back to the queue so that we can retry later.
		ctrl.deleteSnapshotQueue.AddRateLimited(key)
	} else {
		ctrl.deleteSnapshotQueue.Forget(key)
	}
}

func (ctrl *backupDriverController) syncDeleteSnapshotByKey(key string) error {
	ctrl.logger.Infof("syncDeleteSnapshotByKey: Started DeleteSnapshot processing %s", key)
	//ctx := context.Background()
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		ctrl.logger.Errorf("Split meta namespace key of DeleteSnapshot %s failed: %v", key, err)
		return err
	}

	delSnapshot, err := ctrl.backupdriverClient.DeleteSnapshots(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		if k8serrors.IsNotFound(err) {
			ctrl.logger.Infof("DeleteSnapshot %s/%s is deleted, no need to process it", namespace, name)
			return nil
		}
		ctrl.logger.Errorf("Get DeleteSnapshot %s/%s failed: %v", namespace, name, err)
		return err
	}

	if delSnapshot.Status.Phase != backupdriverapi.DeleteSnapshotPhaseNew {
		ctrl.logger.Debugf("Skipping DeleteSnapshot, %v, which is not in New phase. Current phase: %v",
			key, delSnapshot.Status.Phase)
		return nil
	}

	if delSnapshot.ObjectMeta.DeletionTimestamp == nil {
		ctrl.logger.Infof("syncDeleteSnapshotByKey: calling deleteSnapshot %s/%s ", delSnapshot.Namespace, delSnapshot.Name)
		return ctrl.deleteSnapshot(delSnapshot)
	}

	return nil
}

// enqueueDeleteSnapshot adds DeleteSnap given work queue.
func (ctrl *backupDriverController) enqueueDeleteSnapshot(obj interface{}) {
	ctrl.logger.Infof("enqueueDeleteSnapshot: %+v", obj)
	if unknown, ok := obj.(cache.DeletedFinalStateUnknown); ok && unknown.Obj != nil {
		obj = unknown.Obj
	}
	if brc, ok := obj.(*backupdriverapi.DeleteSnapshot); ok {
		objName, err := cache.DeletionHandlingMetaNamespaceKeyFunc(brc)
		if err != nil {
			ctrl.logger.Errorf("failed to get key from object: %v, %v", err, brc)
			return
		}
		ctrl.logger.Infof("enqueueDeleteSnapshot: enqueued %q for sync", objName)
		ctrl.deleteSnapshotQueue.Add(objName)
	}
}

func (ctrl *backupDriverController) dequeDeleteSnapshot(obj interface{}) {
	ctrl.logger.Debugf("dequeDeleteSnapshot: Remove DeleteSnapshot %v from queue", obj)

	var key string
	var err error
	if key, err = cache.DeletionHandlingMetaNamespaceKeyFunc(obj); err != nil {
		return
	}

	delSnap, ok := obj.(*backupdriverapi.DeleteSnapshot)
	if !ok || delSnap == nil {
		return
	}
	ctrl.logger.Infof("dequeDeleteSnapshot: Delete DeleteSnapshot for SnapshotID: %s with Namespace: %s Name: %s",
		delSnap.Spec.SnapshotID, delSnap.Namespace, delSnap.Name)
	err = ctrl.backupdriverClient.DeleteSnapshots(delSnap.Namespace).Delete(delSnap.Name, &metav1.DeleteOptions{})
	if err != nil && !k8serrors.IsNotFound(err) {
		ctrl.logger.Errorf("Delete DeleteSnapshot SnapshotID: %s Name: %s failed: %v",
			delSnap.Spec.SnapshotID, delSnap.Name, err)
		return
	}

	ctrl.deleteSnapshotQueue.Forget(key)
	ctrl.deleteSnapshotQueue.Done(key)
}
