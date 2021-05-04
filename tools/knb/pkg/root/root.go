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

package root

import (
	"knative.dev/client/tools/knb/pkg/plugin"

	"github.com/spf13/cobra"
)

// NewKnBuilderCmd represent root command level
func NewKnBuilderCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:           "knb",
		Short:         "Manage and build kn inline plugins.",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	rootCmd.AddCommand(plugin.NewPluginCmd())
	return rootCmd
}
