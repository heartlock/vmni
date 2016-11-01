// Copyright 2015 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vmni

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	vmnitypes "github.com/heartlock/vmni/pkg/types"
	"k8s.io/frakti/pkg/hyper/types"

	"github.com/heartlock/vmni/libvmni"
)

type Vmni struct {
	Vmninet   *libvmni.CNIConfig
	NetConfig *libvmni.NetworkConfig
}

//add pod to network
func (vm *Vmni) Add(runtimeconf *libvmni.RuntimeConf) (*vmnitypes.Result, error) {
	netconf := vm.NetConfig
	return vm.Vmninet.AddNetwork(netconf, runtimeconf)
}

//delete pod from network
func (vm *Vmni) Delete(runtimeconf *libvmni.RuntimeConf) error {
	netconf := vm.NetConf
	return vm.Vmninet.DelNetwork(netconf, vm.RuntimeConf)
}
