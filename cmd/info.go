package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// func init() {
// 	RootCmd.AddCommand(versionCmd)
// }
func versionCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "info",
		Short: "Print the workers office info",
		Long:  `All software has versions. This is gorkers`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("gorkers gorkers")
		},
	}
	return cmd
}
