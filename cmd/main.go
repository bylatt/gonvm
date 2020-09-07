package main

import (
	"fmt"

	"github.com/clozed2u/gonvm"
	"github.com/spf13/cobra"
)

const version = "1.0.0"

func main() {
	useCmd := &cobra.Command{
		Use:   "use [version]",
		Short: "Use specify node version",
		Long:  "Use automatically download, install, and symlink of specify node version.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			version := args[0]
			err := gonvm.Use(version)
			if err != nil {
				fmt.Printf("%s\n", err)
			} else {
				fmt.Printf("Switched node version to %s successfully :)\n", version)
			}
		},
	}
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "GoNVM version",
		Long:  "Show version of GoNVM",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("GoNVM %s\n", version)
		},
	}

	rootCmd := &cobra.Command{Use: "gonvm"}
	rootCmd.AddCommand(useCmd, versionCmd)
	rootCmd.Execute()
}
