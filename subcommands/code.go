package subcommands

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/Aeswolf/cdx/cd"
	"github.com/Aeswolf/cdx/funcs"
)

var (
	cdx  = cd.NewCdx(".")
	Code = &cobra.Command{
		Use:   "code [src] dest",
		Short: "Subcommand to open specified path in vscode",
		Long:  ``,
		Args:  cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			cdx.SetDest(args[0])
			if len(args) == 2 {
				cdx.SetSrc(args[0])
				cdx.SetDest(args[1])
			}

			path := funcs.GetPath(cdx)

			if path == "" {
				log.Fatalf("%s is not located in %s", cdx.GetDest(), cdx.GetSrc())
			}

			command := exec.Command("code", path)
			err := command.Run()
			if err != nil {
				log.Fatalf("Error running code command on %s ===> %s", path, err.Error())
			}
		},
	}
)
