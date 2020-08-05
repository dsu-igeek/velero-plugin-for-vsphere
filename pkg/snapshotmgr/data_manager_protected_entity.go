package snapshotmgr

import (
	"context"
	"fmt"
	"github.com/vmware-tanzu/astrolabe/pkg/astrolabe"
)

type DataManagerProtectedEntity struct{
	astrolabe.ProtectedEntity
	dm_petm *DataManagerProtectedEntityTypeManager
}

func newDataManagerProtectedEntity(pe astrolabe.ProtectedEntity, dm_petm *DataManagerProtectedEntityTypeManager) DataManagerProtectedEntity {
	return DataManagerProtectedEntity{
		ProtectedEntity: pe,
		dm_petm: dm_petm,
	}
}


func (DataManagerProtectedEntity) Overwrite(ctx context.Context, sourcePE astrolabe.ProtectedEntity, params map[string]map[string]interface{}, overwriteComponents bool) error {
	fmt.Println("Overwrite called")
	return nil
}
