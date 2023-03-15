package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/Aeswolf/cdx/cd"
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
			log.Printf("Root command is working\nSrc dir = %s\nDest dir = %s\n", cdx.GetSrc(), cdx.GetDest())
		},
	}

	// executing the root command
	if err := root.Execute(); err != nil {
		log.Fatalf("Error executing root command ==> %s", err.Error())
	}
}
