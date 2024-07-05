package cmd

import (
	"fmt"
	"fr/diginamic/go-cap/internals/services"
	"os"

	"github.com/spf13/cobra"
)

var dtoFields string
var dtoName string

var dtoCmd = &cobra.Command{
	Use:   "dto",
	Short: "Permet de générer une classe DTO, et éventuellement un mapper",
	Long: `
    Crée un dto en précisant le nom et les champs avec des flags flags
  `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("You should provide the name of the dto using the 'name' flag")
			os.Exit(1)
		}
		services.WriteDTO(dtoName, dtoFields)
	},
}

func init() {
	dtoCmd.Flags().StringVarP(&dtoFields, "fields", "f", "", "Liste des champs de la classe. Chaque se construit de la façon suivante : 'nom:type'")
	dtoCmd.Flags().StringVarP(&dtoFields, "name", "n", "", "Le nom du DTO")
    dtoCmd.MarkFlagRequired("name")
}
