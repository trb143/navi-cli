/*
Copyright © 2019 Adron Hall <adron@thrashingcode.com>

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
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// viewCmd represents the view command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "The 'get' subcommand will provide a list of keys and a map of the values.",
	Long: `The 'get' subcommand will provide a list of keys and a map of the values. For example:

'<cmd> config get'`,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		settings := viper.AllSettings()
		for i, v := range settings {
			fmt.Println(key)
			if i == key {
				fmt.Printf("%v\n", v)
			}
		}
	},
}

func init() {
	configCmd.AddCommand(getCmd)
}
