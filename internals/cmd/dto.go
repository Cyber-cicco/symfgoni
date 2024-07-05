package cmd

import (
	"fmt"
	"fr/diginamic/go-cap/internals/services"
	"os"

	"github.com/spf13/cobra"
)

var fields string

var dtoCmd = &cobra.Command{
	Use:   "dto",
	Short: "Permet de générer une classe DTO, et éventuellement un mapper",
	Long: `
    Crée un dto en précisant le nom et les champs en argument
  `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a name for the new Dto as an argument")
			os.Exit(1)
		}
		services.WriteDTO(args, fields)
	},
}

func init() {
    dtoCmd.Flags().StringVarP(&fields, "fields", "f", "", "Liste des champs de la classe. Chaque se construit de la façon suivante : 'nom:type'")
}
