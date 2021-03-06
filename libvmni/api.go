// Copyright 2015 CNI authors
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

package libvmni

import (
	"strings"

	"github.com/heartlock/vmni/pkg/invoke"
	"github.com/heartlock/vmni/pkg/types"
	"github.com/heartlock/vmni/pkg/version"
)

type RuntimeConf struct {
	PodID  string
	NetNS  string
	IfName string
	Args   [][2]string
}

type NetworkConfig struct {
	Network *types.NetConf
	Bytes   []byte
}

type CNI interface {
	AddNetwork(net *NetworkConfig, rt *RuntimeConf) (*types.Result, error)
	DelNetwork(net *NetworkConfig, rt *RuntimeConf) error
}

type CNIConfig struct {
	Path []string
}

// AddNetwork executes the plugin with the ADD command
func (c *CNIConfig) AddNetwork(net *NetworkConfig, rt *RuntimeConf) (*types.Result, error) {
	pluginPath, err := invoke.FindInPath(net.Network.Type, c.Path)
	if err != nil {
		return nil, err
	}

	return invoke.ExecPluginWithResult(pluginPath, net.Bytes, c.args("ADD", rt))
}

// DelNetwork executes the plugin with the DEL command
func (c *CNIConfig) DelNetwork(net *NetworkConfig, rt *RuntimeConf) error {
	pluginPath, err := invoke.FindInPath(net.Network.Type, c.Path)
	if err != nil {
		return err
	}

	return invoke.ExecPluginWithoutResult(pluginPath, net.Bytes, c.args("DEL", rt))
}

// GetVersionInfo reports which versions of the CNI spec are supported by
// the given plugin.
func (c *CNIConfig) GetVersionInfo(pluginType string) (version.PluginInfo, error) {
	pluginPath, err := invoke.FindInPath(pluginType, c.Path)
	if err != nil {
		return nil, err
	}

	return invoke.GetVersionInfo(pluginPath)
}

// =====
func (c *CNIConfig) args(action string, rt *RuntimeConf) *invoke.Args {
	return &invoke.Args{
		Command:     action,
		ContainerID: rt.ContainerID,
		NetNS:       rt.NetNS,
		PluginArgs:  rt.Args,
		IfName:      rt.IfName,
		Path:        strings.Join(c.Path, ":"),
	}
}
