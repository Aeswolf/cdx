package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/Aeswolf/cdx/cd"
)


func main() {
	// cdx 
	cdx := cd.CDX{Src: "."}

	//root command
	root := &cobra.Command{
		Use: "cdx [src] dest",
		Short: "A tool for lazy transitioning between directories and/ files",
		Long: ``,
		Args: cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			cdx.Dest = args[0]

			if len(args) == 2 {
				cdx.Src, cdx.Dest = args[0], args[1]
			}

			log.Fatalf("\nSrc = %s\nDest = %s\n", cdx.Src, cdx.Dest)
		},
	}


	// executing root command
	if err := root.Execute(); err != nil {
		log.Fatalf("Error executing root command ==> %s", err.Error())
	}
}