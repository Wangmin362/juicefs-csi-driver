/*
Copyright 2018 The Kubernetes Authors.

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

package driver

import (
	"github.com/juicedata/juicefs-csi-driver/pkg/juicefs"
	"k8s.io/kubernetes/pkg/util/mount"
)

// NewFakeDriver creates a new mock driver used for testing
func NewFakeDriver(endpoint string, fakeJuiceFS juicefs.JuiceFS, fakeMounter *mount.FakeMounter) *Driver {
	return &Driver{
		endpoint: endpoint,
		controllerService: controllerService{
			juicefs: fakeJuiceFS,
		},
		nodeService: nodeService{
			juicefs: fakeJuiceFS,
			mounter: &NodeMounter{mount.SafeFormatAndMount{Interface: fakeMounter, Exec: mount.NewFakeExec(nil)}},
			nodeID: "fake-node-id",
		},
	}
}
