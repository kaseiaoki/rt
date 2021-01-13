/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"errors"
	"fmt"
	"github.com/kaseiaoki/rt/requestCheck"
	"github.com/spf13/cobra"
	"net/url"
)

var RedirectAttemptedError = errors.New("redirect")

// hCmd represents the h command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "test with host name",
	Long:  `test with host name`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Default argument is url(https://example.io)")
		}
		u, err := url.Parse(args[0])
		if err != nil || u.Scheme == "" || u.Host == "" {
			return errors.New("Default argument is url(https://example.io)")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := requestCheck.AllRedirectHeader(args[0])
		if err != nil {
			fmt.Println(err)
			return nil
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(hostCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
