package cmd

import "github.com/spf13/cobra"

var ConfigPath = []string{
	"./configs",
	"../configs",
	"../../configs",
	"../../../configs"}

var WelkomText = `
========================================================================================  
Messaging Services
========================================================================================
- port    : %d
- logrus     : %s
-----------------------------------------------------------------------------------------`

// RootCmd this function for root command
func RootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "serve",
		Short: "",
		Long:  "",
		Args:  cobra.MinimumNArgs(1),
	}
	return root
}
