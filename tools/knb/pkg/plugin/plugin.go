/*
Copyright 2021 The Knative Authors

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

package plugin

import "github.com/spf13/cobra"

// Plugin represents plugin configuration data struct
type Plugin struct {
	Name             string    `yaml:"name"`
	Description      string    `yaml:"description,omitempty"`
	Module           string    `yaml:"module"`
	Version          string    `yaml:"version"`
	PluginImportPath string    `yaml:"pluginImportPath,omitempty"`
	CmdParts         []string  `yaml:"cmdParts,omitempty"`
	Replace          []Replace `yaml:"replace,omitempty"`
}

// Plugin represents go module replacement declaration
type Replace struct {
	Module  string `yaml:"module"`
	Version string `yaml:"version"`
}

// NewPluginCmd represents plugin command group
func NewPluginCmd() *cobra.Command {
	var pluginCmd = &cobra.Command{
		Use:   "plugin",
		Short: "Manage kn plugins.",
	}
	pluginCmd.AddCommand(NewPluginInitCmd())
	pluginCmd.AddCommand(NewDistroGenerateCmd())
	return pluginCmd
}
