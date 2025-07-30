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
	"github.com/trb143/navi-cli/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var configFile string


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Navi CLI data Handler",
	Long: `
examples and usage of using your application. For example:

to quickly create a Cobra application.`,
} 

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVarP(&helper.Verbose , "debug", "d",  false, "verbose output")
}

func initConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("navi-cli.yml")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("NAVICLI")
	fmt.Println(viper.Get("global"))
	fmt.Println(os.Getenv("NAVICLI_GLOBAL"))

	if viper.Get("global") == "True" {
		fmt.Println("T")
		viper.AddConfigPath("$HOME/.config")
	} else {
		fmt.Println("L")
		viper.AddConfigPath(".")
		}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("A")
		helper.CheckPrint(fmt.Sprintf("Using configuration file: %s", viper.ConfigFileUsed()))
	}
	fmt.Println(viper.Get("ip"))
	fmt.Println(viper.ConfigFileUsed())
}
