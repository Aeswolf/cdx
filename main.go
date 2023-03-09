package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/Aeswolf/cdx/cd"
	"github.com/Aeswolf/cdx/funcs"
	"github.com/Aeswolf/cdx/subcommands"
)

func main() {
	// cdx
	cdx := cd.CDX{Src: "."}

	//root command
	root := &cobra.Command{
		Use:   "cdx [src] dest",
		Short: "A tool for lazy transitioning between directories and/ files",
		Long:  ``,
		Args:  cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			cdx.Dest = args[0]

			if len(args) == 2 {
				cdx.Src, cdx.Dest = args[0], args[1]
			}

			log.Println("Path =", funcs.GetPath(cdx))
		},
	}

	// adding subcommands
	root.AddCommand(subcommands.Code)

	// executing root command
	if err := root.Execute(); err != nil {
		log.Fatalf("Error executing root command ==> %s", err.Error())
	}
}
