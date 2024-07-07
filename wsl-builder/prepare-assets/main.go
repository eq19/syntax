package main

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "prepare-assets CSV_FILE",
		Short: "Update all releases in WSL distribution info and assets from template",
		Long: `This tool update releases from an incoming CSV file and generate
		       every icons and xml files needed to build a WSL application`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("this command accepts exactly one CSV file")
			}

			return updateAssets(args[0])
		},
	}

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
