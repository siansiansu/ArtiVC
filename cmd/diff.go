/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/infuseai/art/internal/core"
	"github.com/spf13/cobra"
)

var diffCommand = &cobra.Command{
	Use:   "diff",
	Short: "Diff workspace/commits/references",
	Long: `List files in the repository. For example:

# list the files for the latest version
art list

# list the files for the specific version
art list v1.0.0`,
	Args: cobra.RangeArgs(0, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var left, right string
		if len(args) == 0 {
			left = core.RefLatest
			right = core.RefLocal
		} else if len(args) == 1 {
			left = args[0]
			right = core.RefLocal
		} else if len(args) == 2 {
			left = args[0]
			right = args[1]
		} else {
			exitWithFormat("argument number cannot be more than 2\n")
		}

		config, err := core.LoadConfig("")
		if err != nil {
			exitWithError(err)
		}

		mngr, err := core.NewArtifactManager(config)
		if err != nil {
			exitWithError(err)
		}

		err = mngr.Diff(left, right)
		if err != nil {
			exitWithError(err)
		}
	},
}

func init() {
}
