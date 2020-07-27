/*
Copyright the Velero contributors.

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

// Code generated by crds_generate.go; DO NOT EDIT.

package crds

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"

	apiextinstall "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/install"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/client-go/kubernetes/scheme"
)

var rawCRDs = [][]byte{
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xb4V\xcbn+7\f\xdd\xcfW\x10\xb7\x8b\xb4@<F\xd0M1\xbb[\xa7\x05\x82\xbe\x82$Ȧ肖h\x9b\x8dF\x9a\x8a\x94oӯ/\xa4yy\x1c\x17^\xd5\x1bC$u\xf88$5\xd5j\xb5\xaa\xb0\xe3W\x8a\xc2\xc17\x80\x1d\xd3\xdfJ>\x9f\xa4~\xfbNj\x0e\xeb\xe3ݖ\x14\xef\xaa7\xf6\xb6\x81M\x12\r\xed\x13IH\xd1\xd0=\xedسr\xf0UK\x8a\x16\x15\x9b\n\x00\xbd\x0f\x8aY,\xf9\b`\x82\xd7\x18\x9c\xa3\xb8ړ\xaf\xdfҖ\xb6\x89\x9d\xa5X<\x8c\xfe\xbf\xb6t$\xf7M\x05`\"\x95\xfb/ܒ(\xb6]\x03>9W\x01xl\xa9\x81-\x9a\xb7\xd4Eꂰ\x86\xc8$u/\xb2\x91\x8f\x05\xb5\x92\x8eL\xf6\xbe\x8f!u\xe3\x8dY\xdd#\r\xf1\xf5\xb9}_L\x9eF\xd0\xf7\xa2r,\xfa\xd3E\xf5\xcf,ZL:\x97\"\xbaKA\x15\xb5\xb0\xdf'\x87\xf1\x83Av &t\xd4\xc0\xc6%Q\x8a\x15\xc0\x11\x1dے{\x1fZ\xe8\xc8\u007f~|x\xfd\xf6\xd9\x1c\xa8\xc5^\b`IL\xe4\xae\xd8\xc1͇\u0600\x05\x10L\x8f\xba*N,ā\xb5\x1a\xe0A\xb3\xc5D\x8b\x1dP\x01\xb6\xef\xa0\a\x1a\xf0ྔ\v\xd0\xe7\xcb;\x8a\xe4\r\xd9l\x03\xcf\x1e;9\x04\xbd\x85\x8d\v\x9e~\x8c\xa1\x1dE\xc5\xfc\x9e\x1ci\xf6\xf4r\xa0\t|\x8cR\xa60ǐJ(\xc8^\x8aw\x13ɒWF\a\xbb\x10\x01\x87\xc2\xc1\\\xb93\xe09\xf1!b\x9b\x1b\x93z\xb4\x9es\xd0\x03*|a\xe7`K\x90\x84,h\x00E\xf7V\xfe\x0ft\x82\x0e\xf0\x9bw\xef\x13\xfa\x94+\xa9\xa9a\xf3$\xb0\x8b\xa1\xed\xfb\xa7CSܠ\x02F*\xddB\x16\xd8\xc3g\xe7\xc2\x17\xb2\xbf\xceF\xd9\xf7\\g\x024\xd94\xf8[\xe0]\t`\x02\xcc܀\x0fz\x19'\x9b\x86\x8ebi\x92%\xea\x0e\xd9\xd57ñ\x8b\xd9Jyl\xf2\xfc\xc3s\xb4Y\x05\xc0J\xedB\x00\xa0\xef\xb9;E#\xfb}u.\xc6\x18q\xae\xd2\xc9\"9\xb1\\\xf6in\xe4\xdef\xc1а\x00Ȃ\x94&\x87\x90\v\u0092)\x89$\xe4\xfbe\xb2\x88,\xec\x00=\x84\xed\x9fd\xb4\x86g\x8a\x19\x04\xe4\x10\x92\xb3\xb9\x9b\x8e\x14\x15\"\x99\xb0\xf7\xfcτ,#\xd9\x0e\x95\x86\xf1\x9d\xd2\xf7Jѣ\xcb#\x98趴q\x8b\xb9G\xb3\x0fH\xfe\x04\xad\x98H\r\xbf\x84H\xc0~\x17\x1a8\xa8vҬ\xd7{\xd6qu\x9aжɳ\xbe\xafˤ\xf16i\x88\xb2.[n-\xbc_a4\aV2\x9a\"\xad\xb1\xe3U\tܗ\xcdY\xb7\xf6\xabq<\xe4\xe6C\xf1\xcf8ٞ-\x80\x8dCn\x9bk\xb7\xca\xde\xfbO\xb6\xf2\xda\xeb\x17I\u007f\xad\xcfz&%\x8br-\x9f~x~\x99'9\x13\xb7d\xaap4_\x93\x99\xae\\^\xf6\xbb2\x9d<\x8cUF$o\xbb\xc0^\xfbm\xe0\x98\xfc\x92*Iۖ5\xf7\xc7_\x89D3\xab5lʳS\xa6\xbb\xb3\xa8dkx\xf0\xb0\xc1\x96\xdc\x06\x85\xfew\xb2r\x85e\x95Kz\x9d\xae\xd3\xd7ri\xd8Wk\x12\xe7\xa5\xd4\xf5\x94>bĖ\x94\xe2bF\xd1\xda\xf2\x00\xa3{\xbc0\xefWg\xf8\x82\xbb\xd3Ez\xb5\x83\xe4h\xceߞ\xbc\\\xae\xdc˼q\xa4\xa9\xf7V\x1f\x17Ӥ\xb9\xd8ٓ\xf6b}\x16\xda\xd3t\xaa\x8b\xa9O\xdf\x1fǻ\xf9T\x8a\xb8\x1a>7\x8a\x02@r\xdb\xda\x064\xa6\xfe\xed\x11\r\x11\xf74HDQS\xb9\x87\xc6P\xa7C:\xa7\x9f\x18\x9f>-\xbe\x18\xca\xd1\x04\xdfS(\r\xfc\xfeGգ\x92}\x1d\xe3\xc8\xc2\u007f\x03\x00\x00\xff\xff\xd2\x1e\v\x14\xab\t\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xb4UKo+7\x0f\xddϯ \xee\xb7\xc8W\xe0z\x8c\xa0\x9bbv\xa9o\x17\x17} \xb8\t\xb2)\xba\xa0G\xb4\xcdF#\xa9$\xe54\xfd\xf5\x854\xe3\xf1#\x0e\xb2\xeaR\x14ux\xc4\xc3G\xb3X,\x1aL\xfcD\xa2\x1cC\a\x98\x98\xfe6\n\xe5\xa4\xed\xf3\x0f\xdar\\\xeeo\xd7dx\xdb<sp\x1d\xac\xb2Z\x1c\xbe\x91\xc6,=}\xa1\r\a6\x8e\xa1\x19\xc8Сa\xd7\x00`\bѰ\x98\xb5\x1c\x01\xfa\x18L\xa2\xf7$\x8b-\x85\xf69\xafi\x9d\xd9;\x92\x1a\xe1\x10\xff\xff\x8e\xf6\xe4\xbfk\x00z\xa1\xfa\xfe\x91\aR\xc3!u\x10\xb2\xf7\r@\xc0\x81:Xc\xff\x9c\x93P\x8a\xca\x16\xe5\xb5\xf7ȃ\xb6\xa3\xd9\t\xef+r\xa3\x89\xfa\xc2`+1\xa7ë\xe3\xf5\x886q\x1c\xff\xf7cu\xf96\x03\xaf\np\xbd\xf7\xac\xf6\xf3\xfb>\xbf\xb0Z\xf5K>\v\xfa\xf7(V\x17\xe5\xb0\xcd\x1e\xe5\x1d\xa7\x06@\xfb\x98\xa8\x83\xdf\n\xbd\x84=\xb9\x06`\x8f\x9e]\xcd\xcaH8&\nw\xf7_\x9f\xbe\u007f\xe8w4\xe0h\x04p\xa4\xbdp\xaa~ps\x9d,\xb0BVr`\x11\\ѐ\x96\xd8\xf7\xa4\n\xf8\xe6A\vp7A\x03\x04zy\xe3\x00/\xec=\xaci\x14\x8d\x1c\xc0\v\xdb\x0elGpt\xfaR\x93\xfe\x19VB\x8e\x821\xfa\x19\x13\x83\x83;\xef\xe3\v\xb9\xf9\xbf:\x82\x12ێ\xa4`\x17\xb4p\xb8\x05ۡU\xd3%\x97\x87D=\xc0\v\xea\x8c~ \xc5\x01\xa2\xd47oc\x952\xe1\r\x8f^\xef\xc1\xb6\x00\x8f;\x9aq\xdfda\xc3\xe4\xddH\xbb\x10\xce\xc9ոs.\n{\x88\x9b\xab\xf8\x97lۛɒ$&\x12\xe3C\x91\xd6|]\xf2?^\x01\xb0\xd1pf\x00\xb0\xd7RJj\xc2a\xdb\\\x9aQ\x04\x8f\xe1O\x86\xc1\x89\xe7yE\x95\x92\x1b}\xa6\xd2\xd1\xfa\xa5\xa9\x89Ɂ\xd6r\x1c\xbf\xca\nBIH)\x8c\x03\xe1\x8cY\xdc\x00\x06\x88\xeb?\xa9\xb7\x16\x1eH\n\b\xe8.f\xef\xca\xccؓ\x18\b\xf5q\x1b\xf8\x9f\x19YKՖ\x90\x1e\x8d\xa6\xa6\x9b\xbf\x1f\x8c$\xa0/͒\xe9s\xad\xad\x01_A\xa8Ā\x1cNЪ\x8b\xb6\xf0k\x14\x02\x0e\x9b\xd8\xc1\xce,i\xb7\\n\xd9\x0e㯏Ð\x03\xdb\xeb\xb2\x0e1^g\x8b\xa2\xcb:\xa9\x96\xca\xdb\x05J\xbfc\xa3\u07b2\xd0\x12\x13/*\xf1P\xa7_;\xb8\xff\xc94+\xf5\xe6M\xf2/4Y_TE\xf7у:\xb2\xde\x15\xaa\f\xab\xd2\xe88=\x1b?|ԣ\x98j\x93\xfe\xf4\xf0\b\a\x96U\xb3s\x91\xaa<\xc7gzT\xaad\x96ÆdTz#q\xa8\x88\x14\\\x8a\x1c\xc6\x16\xed=S8WI\xf3z`+\xa5\xf1W&\xb5\"h\v\xab\xba5NZ\xa7\x85\xaf\x01V8\x90_\xa1\xd2\u007f\xaeSɰ.JJ?V\xeatٝ;\x8eٚ\xcde\xb2\xa7Q\xcd{\x14\x1c\xc8H\xce\xda\x13\x9d\xab\xfb\x13\xfd\xfd\x95V\xff\xb0}\xaf\x84;\x9d\xb7\x1fTP\xc9?\v\xcd5\xb4\xb8\xce\xf7\xec\xf6\x14\xbe\xb9Je^\xe7\xfb\xdb\xe3\xa9~j1m\xefz\x01\xa0\xa5\x8c\\\a&y\x9c\xabjQpK\x93E\r-\xd7we7%\x9bF\xde\xe9\xb6\xfe\xf4\xe9l\xe5\xd6c\x1fØR\xed\xe0\xf7?\x9a\x11\x95\xdcӁG1\xfe\x1b\x00\x00\xff\xff\xc0J\x05\xd1\xfa\b\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xb4XM\x8f\x1b7\x12\xbd\xebW<x\x0fc\x03\xa3\x16\x8c\xbd,t3\xe4\xf5\xae\xb0\xebفǘ\xcb\"\a6YR3\xc3&;$[c%\xc8\u007f\x0f\x8a\xec\x0f\xb5\xd4\x1aO\x90D7\x16\xc9WU\xaf\xbe\xa8^,\x97˅h\xf4#\xf9\xa0\x9d]C4\x9a\xbeE\xb2\xbc\n\xc5\xd3?B\xa1\xdd\xea\xf0\xbe\xa4(\xde/\x9e\xb4Ukl\xda\x10]\xfd\x85\x82k\xbd\xa4\x8f\xb4\xd3VG\xed좦(\x94\x88b\xbd\x00\x84\xb5.\n\x16\a^\x02\xd2\xd9\xe8\x9d1\xe4\x97{\xb2\xc5S[R\xd9j\xa3\xc8'\r\xbd\xfe\xb7\x8a\x0ed\xde-\x00\xe9)\xdd\xff\xaak\nQ\xd4\xcd\x1a\xb65f\x01XQ\xd3\x1a\xd28K;\xef\xea`E\x13*\x17CQ\n\xf9\xd46\xca\xebCB]\x84\x86$k\xdf{\xd76k\x9cog\xa4ξ\xce7\x06\xfd\xe4]\xfdЁ\xa6=\xa3C\xfc\xcf\xfc\xfe\u007fu\xc8g\x1a\xd3za\xe6\xccJ\xdbA\xdb}k\x84\x9f9\xb0\x00\x82t\r\xadq\xc7\xe64B\x92bY[\xfa\x8e\xe3\xce\xc4\x10El\xc3\x1a\xbf\xfc\xba\x00\x0e\xc2h\x95\bʛ\xae!\xfb\xe1~\xfb\xf8\xf7\aYQ-\xb2\x10P\x14\xa4\xd7M:\x87\x9bK\xfb\xa1\x03\xda@\n\xd1e\xc6\t\x02\x96\x9e\xd1\xeb\xc6\xdbxl\xb4\x14\xc6\x1c;H@\xe0\xfeq\xf3\x0e\xec\x04\x04z?\n\xe0\u007fV\x12bE\xe8\xe1on\x02\xee+\x11\b\x95\b@\xed\x0eYU\xbf\x1f\x93\xab\xf9\xa7\x931ɯ\xeb\xd6$\x9d\xac\xa1\u05ca\xedǛ\x0e\xa2\xf1\xae!\x1fu\xcfW2u\xcc\xedAv\xce\nӖ\xcf@q6SH\x1a\xba\x9c$\x85\x90(\x85\xdb!V:\xc0S\xe3)\x90\xcd\xf9}\x02\v>\",\\\xf9#\xc9X\xe0\x81<\x83 T\xae5\x8aK\xe0@>\u0093t{\xab\u007f\x1e\x90\x03\xfb\xcb*\x8d\x88\xd4\xe5\xd3\xc0\x8a\x8d\xe4\xad0LLK\xb7\x10V\xa1\x16Gxb\x1dh\xed\tZ:\x12\n|v\x9e\xa0\xedέQ\xc5\u0604\xf5j\xb5ױ\xaff\xe9꺵:\x1eW\xa9&u\xd9F\xe7\xc3*\x15\xde*\xe8\xfdRxY\xe9H2\xb6\x9eV\xa2\xd1\xcbd\xb8M\xc5\\\xd4\xeaoCZޜX\x1a\x8f\x9c\xc1!zm\xf7\x838\x15\xd5U\u07b9\xa48䢻\x96\xed\x1f\xe9e\x11\xb3\xf2\xe5\x9f\x0f_\xc7\f\xe0\x10L9Ol\x8f\xd7\xc2H<\x13\xa5\xed\x8e|\x0eܐ<dU㴍i!\x8d&;%=\xb4e\xad#G\xfa\xa7\x96B\xe4\xf8\x14ؤ\x9e\x86\x92\xd06JDR\x05\xb6\x16\x1bQ\x93و@\u007f9\xed\xccpX2\xa5\xdf'\xfe\xb4\x15O\x0ff\xb6\x06q\xdf\"g#t\xd1-\x1e\x1a\x92\xe9\x8a\xdei\nc\xaas\xfe\x96\x94[\x9bJ4O\xe8\x14\xa7\xd5Z\x00_+\xc2\xe7\xce\xc2\x14\xa3\x92\xe0\x0e\xe4\xbdV\x8a\xecm\x8a\xca\xce\xf9Z\xc4\\s4\xf83-\x8dЫ\xefL\x92\x05\xf0\xe1~\xfb/n\xf7\xa9PR\x86\xe5\xcdc\xc2a\x0e\x18s4;\xb7\x99\xe2\x04x\xae\x8dt\xad$!O\xa5g\x94\r\xea;Ç\xb4-\x89\xd39kSg\b\xb3!\xe4\x1fO\xac\xe6\v5.\xe8\xe8\xfc\xf1E\xcdLj\x9ep\\\n\xdd\r\xf6\xd0S\xf4\x9a\x0e4m\x9b\x1c\xa4\x1c\x8a3\xd0n*\xf2\x18\x9a\xb4\xf2\xd5\xfd\xe3\x06F\x1f(@[\xd4m\x88\xa8ā \xa4\xa40t\xb0Q\xf5k}LI\xb3\x11V\x92yѿގ|\x14\xda*-\xb9]\xf6E\x9a\x82\x99\xf7\x9c\xdd;f\xfbd2Mo_\xb8,\x85\xe5$\f\x14!\"\x84=F]s\xccv\\\xd2\x13\xde<\tYq\xee#\x92\xaf5w\xe6\xa6J\xe5\x8f\xed\xee\x02wr\x95G`\xbe\xae.\xaeϲU:gHL'\xccyK\xbd\xe0\xa9盛\xe9\xfeǲp\xae\x99\xf0/\x97\xe8\x1a\xe51\xce\xdb?\x83Փ\xb1\xfdx\x8ev\xe5\nGW{\x9a\xf8\xbc\x1cjq\"<\xab\x96\xc9\xdeI\x96M\xe4L\xe7D0\x1a\xf8\xdd֙\x9fc\xaf\xe8\x1c5\x85 \xf6\xf4J\x8f\x913\xe2\xc5 \xdf 7\xe7\xfc\xb6\x1aG_~E\xe8\x1dɣ4\x94\x81ң$\x1f/\x00\xdc\xd1\xf3E\x9a.q\xe7\xf0\xec\xfc\x13\x8e\x14oa\xe9[\xec\xeeꀭ\xbd\xf7n\xef\xb9\xc8q\xba\x18\xb9\xcaiu\x99\xfc\xe2\x89,\x80\x8d\xab\x1bC\x91\x14\x96\xe91\xd75_.\x87\x92\xc8\xf6\xe9\b\xe0\x93\xd0&\x1d#\xee\xdbQD\xba\xbd,Uv\x04\xbbt\xf2\x16֝B>\x8bp\x82\x96\xc3\xcd\t\xbf\xc4sE6\x91\x93x\xb8\x00\xdd\x19\xb1g_\x03\xbb\xafw\xe3\xc9d%\x8f}a<\tu잰\xdaFwZ\xc3Wle\x98\xe9/'\r\x9e\xb51\t\x8a\xbb\xd6`g7\x1a\xfb\xe1T\x89x\x81\xf8\x9cH;)\xe2\fUr\x12d\xd3:\xb8D\xe3\xe8Fb\xa6۸\xb4\xf3*\x8b7\xafMپ\xc3\xfc[Xe^\xce]v\xb1J\xc7.\x86$\xbb\x9cՏ/\x89I\x0f=ýVpxa\\_\xd83\x8c\xec\xaei\xa6\u007f\x8b\xdc\xdc\xe6\x06\xb8\xa7\x1dy\xb2\x92T1\x83\xcb3`\x82Ǚӿ\x97T~\xd8\f\xcbܩ\xd3 -\xf9\xddȻ\xb3\x98\x92\x87Ї\xfbm\xb6\xac\xc0'\xe7yD\xc1\xc5*\xbfl\xbdZ6\xc2\xc7c\nO\xb8\x9dX\xd0\xf7\xcf9s\xafF\x13W\x06\xcd\x05w\xaf\x186#c\xbfׂ\xf4?\xff{\x16\xf0\u007f\xe6\xde\x02\xbe\xf0'Z07z075\xb2\x90\xb5ϖ\xcb\xd9ܸ6Ӻ)1\x91\x9d\xbf\rf\x00\xcf\xc1\x96)\xc3\xc6Ej8\x8b\xd9\xeb\xc3W\x97\xc3\xfbq\x95Ji\xd9}dI\x1b@\xe0\xffSj\x8d\xe8[\xea>E8\xcf\x13-K\xc6Iȏ\xc1&\x92\xba;\xff\xb0\xf2\xe6\xcd\xe4+IZJg\x95Ο\x88\xf0\xff\x1f\x16\x19\x95\xd4co\a\v\u007f\v\x00\x00\xff\xff\xcfw\x89\x10\xa1\x12\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xb4XMo\x1b7\x13\xbe\xebW\f\xfc\x1e\xfc\x16\xb0$\x04-\x8aB7W\ue1d0&0,ח\xa2\a.9Ҳ\xe6\x92[r\xa8D-\xfaߋ!\xf7C\xbbZ\xd9\t\x90\xea\xb6\xe4|>3\xf3\x90\xd4l>\x9f\xcfD\xad\x9f\xd0\a\xed\xec\nD\xad\xf1#\xa1寰x\xfe.,\xb4[\x1e\xde\x14H\xe2\xcd\xecY[\xb5\x82u\f\xe4\xaa\a\f.z\x89w\xb8\xd3V\x93vvV!\t%H\xacf\x00\xc2ZG\x82\x97\x03\u007f\x02Hg\xc9;c\xd0\xcf\xf7h\x17ϱ\xc0\"j\xa3\xd0'\x0f\xad\xff\xff+<\xa0\xf9j\x06 =&\xfdG]a Q\xd5+\xb0ј\x19\x80\x15\x15\xae XQ\x87\xd2QX\x14B>\xc7Zy}H\xc6f\xa1F\xc9N\xf7\xde\xc5z\x05\xe3\xedl\xa0\t+\xa7\xb4ml\xa5%\xa3\x03\xbd\x1d,\xff\xa2CުM\xf4\u009c\xf8N\xabA\xdb}4\xc2\xf7\xeb3\x80 ]\x8d+xϮj!Q\xf1Z,|\x03[\xe3>\x90\xa0\x18V\xf0\xf7?3\x80\x830Z\xa5\x9c\xf3\xa6\xab\xd1\xde\xdeo\x9e\xbe\xde\xca\x12+\x91\x17\x01\x14\x06\xe9u\x9d\xe4\xe0\xba\v\x12t\x80\x18P\x019\xf0\xf8g\xc4@@\xa5 \x10]X,B\xe2\x19\xed\x02`\xc3_\x8dE\x00\xeb\xa8S\xae\x84\x15{\x04*\x11\xb4=\xa0%\xe7\x8f\xe0v}\xd2 \xac\x02\xe50$5\xb0\x98\xf5\xf0c\v\x13\xff\xb4\x05\xe7\x15zޑ\xc6\xd9l\xb0M\x1fv\xdeU'\x91]7z\xb5w5z\xd2->\xfc;i\xcfnm\x8c\x02Ôe@qCbH\ue6b6B\x05!A\xc8iP\xa9\x03x\xac=\x06\xb4\xb9EO\xcc\x02\x8b\b\v\xae\xf8\x03%-`\x8b\x9e\x8d@(]4\x8a\xbb\xf8\x80\x9e\xc0\xa3t{\xab\xff\xea,\aΓ]\x1aAx\x02C\x86\x82\xd0[a\xb8\xc0\x11o\x12|\x958\x82G\xf6\x01ўXK\"a\x01\xef\x9cg\xf8wn\x05%Q\x1dV\xcb\xe5^S;\x90\xd2UU\xb4\x9a\x8e\xcb4V\xba\x88\xe4|X\xa6\xd9Y\x06\xbd\x9f\v/KM()z\\\x8aZ\xcfS\xe06\xcd\xe3\xa2R\xff\xeb\xda\xf0\xfa$R:r\xc7\x06\xf2\xda\xee\xbb\xe54 \x17q\xe79\xe1\xae\x12\x8dZ\x8e\xbf\x87\x97\x97\x18\x95\x87\x1f\xb6\x8f}\xf1\xb9\x04C\xcc\x13ڽZ\xe8\x81g\xa0\xb4\xddq#q\xe1R߰E\xb4\xaav\xdaR\xfa\x90F\xa3\x1d\x82\x1ebQi\n\xed(p}\x16\xb0N\xb4\x04\x05B\xac\x95 T\v\xd8XX\x8b\n\xcdZ\x04\xfc\xcfag\x84Ü!}\x1d\xf8S6\x1d\nf\xb4\xba\xe5\x96\xee&+\xb4\xadQ\xa6\xb1g\x94\x12q\xf7e`\xc5\x13\xbd\xa9\xd9\xe3_\xe6\xcf\a\xac]\xd0\xcc\x05\xc3ݑ\xbf\xc7\x12\x1b\x05\xaef\xa3\xc1\xb3\xd1s\x90\xe5J$Aے\xe3\xc8\"\xa4p[b[\xde?\xad\xc1\xe8\x03\x06&\x95*\x06\x82R\x1c\x10\x84\x94\x18\xba\xb9뽍\x8cM\x82˿\x16\x87\x9f\x85U\x06_\xcc\xeaa \n\x1ewܚ\xe4@\xc0\xdbX\xa0\xb7H\x18:\x837 \xa3\xf7hɌc\x01\x10\xc0\xd9\x14\x91r;sw\x17\b\xe9\xc8U\xa88ANf\x17\xb9\x83Fʗ\xea\x03\x99#\u007fJ\xa7\xdd\xd9\xce(\x93\xdb\xfbM\x12l{\"\x9d\x91\xb0s~H\xcf\x05\xf2\xe4\xa6<\xd1JT\x8b\t\xbb\x00\x9b\xdd\xc0\x1e\x8f\x16\xf7\x94\xdeiT7\xc9`\xf7\t\x89)R\xf1\nlҜ\xb4)y\x02o\xef79\xb2\x05\xfc\xe8<\b{\x04Ge\xe6\x00\xaf\xe6\xb5\xf0tL\x85\r7\x83\bxص\x9f\x0e\xf7b\x1f\xc0\x04\xcbMbג\x1d'\xc6\xd6\xf8\xa8\xb8\x88\xd8\xe7F\x90.5\xafE\xc0\xb7\x896\x02V\xf8\x82\x11\xb4Ѝc\x98'l\xce\x16\xd9\xfb䠍ȉ\u007f\xed诅\x95h^\x1c\xb4\xed@\x14\xb4UZ\x8a<[͍Ɓ\xcc{\xce\xee\x1d'\xdcZ_\xc0H\xfb,\u007f),7_@\x02\xbe\x16\xd9#\xe9\x8aa\xdbq˥nm9ʣ\x90%\xf2\xb1F\xe8+\xcdgw]\xa6\x03\x026\xbbI\xb2\xeaTK\x11\x1auu\xa6>\tX\xe1\x9cAqz\a\x99*\xc5\xfc\x8c\x86\a\x9bC.{\xf5\xc4ȷ\xceKg\xc6:\xb3W#\x96\xefL'\x192\v\x8d\xafM\x97\xb8\xa9\xc2\x10\xc4\xfeer}\x97e\xf25\xa2Q\x00Q\xb8H\x03\xbfס\th<\xdc\x17\x9bz\xea\x04\x9d\xf0\x9e\x85:\x0el\xfd\x11\xaa\xa9f\x06\x16\xac\x04\xad\xa08\xd2tE'BI\xe5\u007f1\x8e{\x96\xe8O\xea\xbe\x02\xd8\x16\xe0\xd7\xda8qFm\x97]z\xb7\xf7\x18\xce\x0e\x8b\xa1\xd7F\xa8\xcb>&'\x9fq\xf20\n\xe1\xce\xd9I\xf2j\xa1Җ\xbe\xfd\xe6\"\x1f\xf1\xfdx\x8f\xfe|\xae\x1c\t\xf3=\xdb\xffҶ_\xa5\xaa\xcd\xdd'\xd1\x14l\xee\xf2\x1b\x8b\xa7\xbe@\xb4\xdd\xf3\xea\x91O\xf7\x0f\xda\x18f\x9c\x9d6&\x1d\xeega~(1\x9f\xf8\xa9A`Ϗ*rp\xb5\xed\x9b\xf0\xea\xd3\n>\x91ҘG\xe6\xa77\xbe\x91|\xf7\x00?\xbc\xe9\xbf\x12\xee\xf3潝6\x00\x02\xdf\xcb\xd5\n\xc8Gl\x9e\xb0\xce\xf3\x88畞Z\xf8zV\x13\xaa\xf7\xe3\xc7\xf6\xd5\xd5\xe0-\x9d>\xa5\xb3J\xe7\u007f\v\xe0\xb7\xdfg\xd9*\xaa\xa76\x0e^\xfc7\x00\x00\xff\xff~\x94[,\xac\x10\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xbcX\xcdn#7\f\xbe\xfb)\x88\xed!-P;\xd8nQ\x14\xbeu\x93-\x10\xb4\x1b\x04q\x9aKуfD\xdbj4ҬH9q\x9f\xbe\xa04\u007f\x9eq\x9clۭ/\x899\x14)~\x1f\xff<\xb3\xf9|>S\xb5\xb9\xc7@ƻ%\xa8\xda\xe0\x13\xa3\x93o\xb4x\xf8\x91\x16Ɵ\xef\xde\x16\xc8\xea\xed\xec\xc18\xbd\x84\x8bH\xec\xab[$\x1fC\x89\x97\xb86ΰ\xf1nV!+\xadX-g\x00\xca9\xcfJ\xc4$_\x01J\xef8xk1\xcc7\xe8\x16\x0f\xb1\xc0\"\x1a\xab1$\x0f\xad\xff\xaf5\xee\xd0~3\x03(\x03\xa6\xf3w\xa6BbU\xd5Kp\xd1\xda\x19\x80S\x15.A\xfbGg\xbdҴء\xc5\xe0k\x1b7\xc6-\x8c\x9fQ\x8d\xa58\xdd\x04\x1f\xeb%\x8c\x1fg\x03͵rH\x97\x8d\xad$\xb2\x86\xf8\x97\x03\xf1\xaf\x868=\xaam\f\xca\x0e|')\x19\xb7\x89V\x85^>\x03\xa0\xd2\u05f8\x84kqU\xab\x12E\xb6S\xd6\xe8\x14Tv\xeekt?\xdd\\ݿ[\x95[\xacT\x16\x02h\xa42\x98:\xe9uwh\xa4\x05\x82j\"\x9a\xe7\x90  \xb1\x0f\xd8\x1c\xae\x83\xaf1\xb0i\x03\x94π\xe0N6rs&\xf7\xc8:\xa0\x85R$\xe0-BC\fj\xa0tG\xf0k\xe0\xad!\bX\a$t\x99\xe4\x81Y\x10\x15\xe5\xc0\x17\u007fb\xc9\vXa\x10#@[\x1f\xad\x96<\xd8a`\bX\xfa\x8d3\u007fu\x96\t\xd8'\x97V16x\xb7\x1f\xe3\x18\x83SV\x10\x8c\xf8-(\xa7\xa1R{\b(> \xba\x81\xb5\xa4B\v\xf8\xe8\x03\x82qk\xbf\x84-sM\xcb\xf3\xf3\x8d\xe16\xa5K_U\xd1\x19ޟ\xa7\xc44Ed\x1f\xe8<e\xdf9\x99\xcd\\\x85rk\x18K\x8e\x01\xcfUm\xe6\xe9\xe2.e\xf4\xa2\xd2_\x85&\xff\xe9lpS\xde\v\xe7\xc4\xc1\xb8M'N)\xf6,\xee\x92i`\bTs,߿\x87WD\x82\xca\xed\x87\xd5\x1d\xb4N\x13\x05\x87\x98'\xb4\xfbc\xd4\x03/@\x19\xb7Ɛ\x89[\a_%\x8b\xe8t\xed\x8d\xe3\xf4\xa5\xb4\x06\xdd!\xe8\x14\x8bʰ0\xfd)\"\xb1\U0003300bT\xd8P \xc4Z+F\xbd\x80+\a\x17\xaaB{\xa1\b\xbf8\xec\x820\xcd\x05җ\x81\x1f\xf6\xa3CŌV'n\x1b\xc6Q\x86V5\x96BPB)\xb5\xbe\x9e\x06988w\xac\xf6\xe4S\xa8\xf2!ַX{2\xec\xc3^\x9a¡\xc6\xc8\xe7\xfb\xd1\x01\xb1\xbc3\x1a\xa91%<\xb7\x8f\x04kX\xfb\xd05\x9f\xc5\xc82L̉\xff6$酹\xa8q\xa27\xb6t\x14\xe5&p\xe1\f\xf5\aǆ\xf7W\x97'\xa3\xbb\x19k\xb7w1Z\x98^\x1b\f) \x11u\x96\x01\x93\xf24\xb8;IjC\xe0\x105ji\"~\x87\xe11\x18FiC\xf8d(\x15\xd1\xce\xdbX\xe1\xabCj\xbaj?}NEt;RN\xdd-\xe8\x1c\x15\x9b\n\xd3?\x8dIxT\x04\xa5\xb2\x16\x8fPu\xb7E\xa0T\xb8g\x94O\x1a\x82H\xa8\x13\"+\xa7j\xdaz\xee\x1c\x8dί}\xa8\x14/A*s.\xa7_\x1b-5\x86_`nթ\x9d\xa0\xac\xb5\xd5f\xd5g\x01\u007f\xbc>Yq\xa4g+\xb4\x1d\x91\xab\xa4\xd6\xd7j\b\xe8\xb89\x9c\x86R\xa7\xb9xEі\xbe\xaa-\x1en \xa7\xb0\xb9\x98\xeaO\xd3@\xb9\xaeLs\x1a\xe4C\xc72\xa1\xb7\xd7\xe5A6'\xb5\xb0C\a\xde\xc1Z\x19\x8bz\xb0\n\x1d\xe4\xcf\xc4\xe4$\x9f\x8e\xdcy|\ua954\x92\x9dL\x15\x16\x97\xc0!\xbe:\xdf*$R\x9b\xd3M\xf0c\xd6\xc9ñ9\x00\xaa\xf01O\xac6\xe83j(~ui;|\xe2[\xe4\xb0\u007f\x1d\xb1\xd7\x13\xf5v\x97)\xb0c6\xcby\xab\x18\x8cӦT<\xc6\x02r\xbb\xc5'\x99\xd2\x1c\xf6\x89\x80\x11\x81pq\xbb\x80\xdf(\xb7\xb1\xb5\xb1\x8c\x01\xc6\xf1N\xcc6\xf3\x19\x1e\xb7\xa6\xdcJF!\x81qP\xe0Z\x9a\xcd\xc0\xa1\xdcs\fҗ\xa1\xb7\xde*:M\xee\x8dh\x1c+\xd5n\x18\x1d\xabU\xf9\xa0\x8b\xd5\xd8\xf4\x1c\xae\xf1q\"\xbbr7\xc1o\x02\xd28\xa7\xe7m\xe2\xe3\x18\xcd9$\xa6'ҟ\x13O\x9f1\x0eK$\xf9Yp\xed\xf5i\x1c\xa4^/\x15\xab\x8fʩ\r\x06p^cN\xa3\xad\"\xa8M\xf9\x80\x1ab}\x80\x88\x906ɂ\xde\xe7\"\x8f\xc4Gc\xed`O\x03E@\xde;\xf9{`̴n&&c\x9drth\xf9\xaa\xe1fp\xe3RJ\xc1\x9dq\xabw\xfa\x9a\xe4+\x19\x84\x8a\xbc\x03\xc3\xdd%\xfb@\x8b=(\xe7y\xdb`\xf19+H\xa2\xfa\xa5\xcd#)\xc1\xd6۶-{V\x16\\\xac\n\xa9\xb55\x14{F:\x1c]\x93 \xe4\xe7\xc70o\xfb\xd3\xed\x88O\x810R\xc3E\xa9\\\xa2\xa2\xa9lm\xa8\xb6j\x9ce}\fi\xa7\x93\xba\x94\xe6\xdfw\xbcָ̫\xf4l\f\xces\xa3L>\xe9:\x97\xdeM\xd2\x11\x06]\xc08\xfe\xe1\xfb#\xcf3\xea\xf2+l\x83SN\x13\x84\xef\xc5\xfe\u007fm\xfb\xe86\x00\xa9\xe7q\xd8_\xf8\xe8\xf8\x85\xbd\xacU;\x18\xc5C\xbe\xda\xc6H)\xd7\x03Ε\xd6\xe3ܒ\x8fj{\xf5e߫\x9b\x96\xdb6\xf1(\xd9\x04\x0e\xf9ч\a0D\x11Ӿ.\xd2O\x11\xe3$\x99!\xf7zq\x1cI~\x9f\x05U>Ȯ*\t\xa6\xb1\x88\x9b\x8dT\xdd3-\xdb8~\xf7\xddQ\xc0\x8e\x81I\xac\x02\xbfn\xe0\xad\x0eT_^b\x92\xe9\u007f\xb0\xcc\x1e\xb8\xf9\u007f\xf6\x8e\\\xd0/l\xb9\xf7\x8d҉\x1d\xb7)E\xfd\xaf\x97[\x19\xde&`\xf7\x82`>\xfcM9\xd2\xef^\x92\xed\xde\xf6\xdfR\xcd͛wb\xe9\x01d\xcc\xf5\x00\x19\xb9\xac\xac[YүӪ,\xb1f\xd4\xd7\xe3\x17bo\xde\x1c\xbc\xefJ_K\xef\xb4\xc9o\xf4\xe0\xf7?f\xd9*\xea\xfb\xf6\x1e\"\xfc;\x00\x00\xff\xff-K\xfc\xc5P\x14\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xbcX\xcdn\x1b7\x10\xbe\xeb)\x06\xe9\xc1-P\xc9HS\x14\x85n\x8d\x9c\x02F\x1b7\xb0\x92\\\x8a\x1e\xb8\xcbY-k.\xc9pH9\xea\xd3\x17C\ue7f4kY-\xd2\xeab\xefp8?\xdf|3\xe4\xeeb\xb9\\.\x84S\x1fѓ\xb2f\r\xc2)\xfc\x1c\xd0\xf0\x13\xad\x1e~\xa4\x95\xb2\xd7\xfb\x97\x05\x06\xf1r\xf1\xa0\x8c\\\xc3&R\xb0\xcd=\x92\x8d\xbe\xc4\x1b\xac\x94QAY\xb3h0\b)\x82X/\x00\x8416\b\x16\x13?\x02\x94\xd6\x04o\xb5F\xbfܡY=\xc4\x02\x8b\xa8\xb4D\x9f<t\xfe\xbf\x96\xb8G\xfd\xcd\x02\xa0\xf4\x98\xf6\xbfW\rR\x10\x8d[\x83\x89Z/\x00\x8chp\r\xd1i+$\xad\xf6\xa8\xd1[\xa7\xe3N\x99\x95\xb2\vrX\xb2˝\xb7ѭ\xe1t9oo\x83\xca\t}H\x96\x92@+\n\xbf\x8c\x84\xbf*\ni\xc1\xe9\xe8\x85\xee\xbd&\x19)\xb3\x8bZ\xf8N\xba\x00\xa0\xd2:\\\xc3\x1d\xbbp\xa2D\x96\xed\x85V2\xa5\x92\x9dZ\x87\xe6\xa7w\xb7\x1f_m\xcb\x1a\x1b\x91\x85\x00\x12\xa9\xf4\xca%\xbd\xd6{++\x10D\x9b\xc72'\x02\x85(\x1f\xa2kw:o\x1d\xfa\xa0\xba\xac\xf87\xaai/;\xf1q\xc5Ad\x1d\x90\\E$\b5B[\v\x94@)@\xb0\x15\x84Z\x11xt\x1e\tM\xae\xeb\xc8,\xb0\x8a0`\x8b?\xb1\f+آg#@\xb5\x8dZr\xe9\xf7\xe8\x03x,\xedΨ\xbfz\xcb\x04\xc1&\x97Z\x04l\x81\xee~\xca\x04\xf4Fh\x86/\xe2\xb7 \x8c\x84F\x1c\xc0#\xfb\x80hF֒\n\xad\xe0\xad\xf5\b\xcaTv\ru\b\x8e\xd6\xd7\xd7;\x15:\x16\x97\xb6i\xa2Q\xe1p\x9d\xb8\xa8\x8a\x18\xac\xa7\xebD\xb8kR\xbb\xa5\xf0e\xad\x02\x96!z\xbc\x16N-S\xe0&\x91x\xd5ȯ|Ky\xba\x1aE\x1a\x0e\\p\n^\x99]/N\xbcz\x12w&\x18(\x02\xd1n\xcb\xf1\x0f\xf0\xb2\x88Q\xb9\u007f\xb3}\x0f\x9d\xd3T\x82c\xcc\x13\xda\xc36\x1a\x80g\xa0\x94\xa9\xd0\xe7\xc2U\xde6\xc9\"\x1a\xe9\xac2!=\x94Z\xa19\x06\x9dbѨ\xc0\x95\xfe\x14\x91\x02\xd7g\x05\x9b\xd4\xcbP D'E@\xb9\x82[\x03\x1bѠ\xde\b\xc2\xff\x1cvF\x98\x96\f\xe9\xf3\xc0\x8fGбbF\xab\x17wSb\xb6B[\x87%\x17(\xa1\x94\xa6\xddP\x06\xde8\xda7\xd7{\xfc\xcb\rz\x8fΒ\n\xd6\x1f\x8eWO\xfc\xbd>Qf\xab{%\x91Z3\\\xe3n\x89q\x86\xca\xfav\xe6\xacN\xec\x02| \x94I\xa1\x89:(\xa7qj\xe4t\xd3,\x94C\x16\xc3\x04~>\x89^7u\xbb\x97\x19Ġ\x1aL\xff\xb4\xa1<\n\x82Rh\x8d3\xf1\xbf\xaf\x11(\xf1\xf8\x8a\xf2FE\x10\xbb\xa4\xb6F8\xaam\xe8\xfd\x9c쯬oDX\x03\x13uɻ/M\x95Z÷7g\xb3\xdc\xf6j\x1dA\x94d\xaeV\n}\n\x90E\x9d\xad<9\x11\xf6V\xc7\x06/F=Wv#L\x89\xfal0\x1fF\x8a\xa0\x8cT%\xcfѮ{y\xb8\x96y͚\x9d\xe5\xa92ϙ\x1cHa\xadFa\x9em\x9c B\xa4'['\x87\xb4MJC\vy\x8f&\xb4[\xd3Y\xd1\xea\xad.\xe8\xa4\xd26N\xe3\xf1M\xe0\x1c(\x9b\xa9\xfe\x94\x8b´Xd*\xe6-sl\x1c\xac\xf5\\\xcc\xc6P\x02\xeeр5P\t\xa5Q\xf6\x17\x92)\x83'Vǌ\x9e\x89\x97\xfe!\xa9\xf9^$\n\x8dk\b>^\xcc\xf8\xb6,ܷ\xbfU\xd5yL\x8fT\x8f\xe0䎶U\xc58x\f\xfe\x90rʈL\xd2n\xe9\a\xf7I\xd1\xf65h\xef\t\xb6\xc0\x03\xe0gg\rw\x94н\xed\x06\xcbZ\x18E\xcdi}:`\x94\t\xaf\xbe\x9b͛\xaf\x11;\xf4Gk\r\x12\x89\x1d\x9e\xcd\xf8m\xd6\xc9\xc7t\xbb\x01Dac>;s\xe0W\xd4r\xfa\xe2\xd66\xf89\xa4\xec/\xe3\xf2\xddD\xbdê\xc0\x9e\xccY\x1ej\x11\xfa\x190A\x9eu\xd9\xf7\xa4F]\t6\xf7\xab|r\x04\v\x95\xd2\x01=\x1c\xe7:1\xd9M\x99\xc7Z\x955\xb7\x10\x12\xf0\xcd\x14+\xbe\x0e\x8c\x9cq\x8cO\x15\xee\xcb2\xdaՂΗ\xf5\x1dk\xcc\xcd%\xecf\xf5t0\xf1\x0fMlN\r/\xe1\x0e\x1f'\xb2[\xf3\xce\u06ddG:m\xe2e\xd7\xe9\x93\xc6X\xb6N\xdfxo\xfdt\x17O\xe5\xe8~\x9ek\xa9%\xe4\xe1\xff\xe4\xc2)B\xe7\xc0\xf3\xb6D\xe2w\x9a;+ϣ\xc8\x03\xeeF\x04\xf1V\x18\xb1C\x0f\xc6J\xcc\x04\xac\x05\x81S\xe5C\"\xd6\bO.\xf8\x84A\x83G\x9e\x99\x8a\xe0Qi=\xbag\x82  k\r\xff\x1d\x99R#\x17\x13\x9b\xcc\xec\xb1\xdd\xdb\\\xd5q\xb4%7\x90\xb9\n\x9d\xde(D ;3\xad=\n\xb2\x06T\xe8\x03\x1cR,\x0e \x8c\ru\x8b\xc2Ń\xc0\xb5$9\xcf\xd6V\tj\xab\xbb\xd3\xcb\x06\xa1\xc1Ħ\xe0\x0e\xad\xa08\xf0\x89\u007ftϘ\xc4ϯNc\xbe\x8fv\x8b!\x8f\x80Ԗ\xa1\x14&U\xa1\x1d\aR\x91\xd3\xe20W\xbe\x1c\x1d_I\xb9\x9d\xf9\x90\x1cFd{\xd5\xe3C=-\x9dB\xf3\xd4yϿ\x14͍5\x13\x1a\xc2\xf1\xd0\xff\xe1\xfb\x99\xf5\xa7\a\u007fZe\x00_\xb3\xfd/m{\xf6\xba\x04\x89?\xc1\x1f66\x9ap\xb6\xda\xf7\xbd\xda\xd1\x01;Tk\x18\xa7\x94\xc8*\xa4\x9c\xf67\x80\xe8F{\xcb\xeb<\xa0[\x99\x8c\xcc!0\x18\x1e\xad\u007f\x00E\x14s\xa9X\xfa)b\xc4|\x0e\xcc6V$~\xa3\xf4\xa2|\xe0\xcb$\xd3Jb\x11w;\xee\xb4/r2S\x10>\\v4n\x8fT\x9f\xbb\xe1%\xc3\xff\xe2m\xe3\xc8\xc9\xffq-\x9b!\x11\x1f\xb4\xcac\xffAa9~\a=\xd1￣\xed_\x0eO\x89\xe9\xcb\xf6\xb3YZ\x80\x9c\xb3\x1cEF\xc1z\xbe\x14e\xc9p\xcb\x17e\x89.\xa0\xbc;\xfdj\xf6\xe2\xc5ч\xb1\xf4XZ#U\xfe\xe8\a\xbf\xff\xb1\xc8VQ~\xec\xe2`\xe1\xdf\x01\x00\x00\xff\xffh \x93[s\x14\x00\x00"),
}

var CRDs = crds()

func crds() []*apiextv1beta1.CustomResourceDefinition {
	apiextinstall.Install(scheme.Scheme)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	var objs []*apiextv1beta1.CustomResourceDefinition
	for _, crd := range rawCRDs {
		gzr, err := gzip.NewReader(bytes.NewReader(crd))
		if err != nil {
			panic(err)
		}
		bytes, err := ioutil.ReadAll(gzr)
		if err != nil {
			panic(err)
		}
		gzr.Close()

		obj, _, err := decode(bytes, nil, nil)
		if err != nil {
			panic(err)
		}
		objs = append(objs, obj.(*apiextv1beta1.CustomResourceDefinition))
	}
	return objs
}
