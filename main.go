package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/Aeswolf/cdx/cd"
	"github.com/Aeswolf/cdx/funcs"
	"github.com/Aeswolf/cdx/subcommands"
)

// entry point
func main() {

	// instance of cdx
	cdx := cd.NewCdx(".")

	// root command
	root := &cobra.Command{
		Use:   "cdx [src] dest",
		Short: "A tool for lazy transitioning between directories",
		Long:  ``,
		Args:  cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			cdx.SetDest(args[0])

			if len(args) == 2 {
				cdx.SetSrc(args[0])
				cdx.SetDest(args[1])
			}
			// log.Printf("Root command is working\nSrc dir = %s\nDest dir = %s\n", cdx.GetSrc(), cdx.GetDest())
			log.Printf("Needed Path = %s", funcs.GetPath(cdx))
		},
	}

	// Adding subcommands
	root.AddCommand(subcommands.Code)

	// executing the root command
	if err := root.Execute(); err != nil {
		log.Fatalf("Error executing root command ==> %s", err.Error())
	}
}
