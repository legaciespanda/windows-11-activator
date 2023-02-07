/*
Copyright © 2023 ERNEST OBOT <youandinews@gmail.com>

*/
package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
)

var appVersion = "0.0.1"


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "windows-11-activator activate [key]",
	Example: `

	Activate Windows 11 Pro : windows-11-activator activate --key=WNMTR-4C88C-JK8YV-HQ7T2-76DF9

	`,
	Version: appVersion,
	Short:   "Activate Windows 11 Activator for Pro, Workstation, Home, Education, Enterprise, Core ",
	Long:    `Windows 11 Activator activates Windows 11 Pro, Workstation, Home, Education, Enterprise with a single command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Windows 11 Activato CLI")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}


