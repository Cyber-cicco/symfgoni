package cmd

import (
	"fmt"
	"fr/diginamic/go-cap/internals/services"
	"os"

	"github.com/spf13/cobra"
)

var ctrlJName string
var ctrlJEntity bool

var ctrlJCmd = &cobra.Command{
	Use:   "json-ctrl",
	Short: "Crée un Controller JSON",
	Long: `
    Permet de générer une classe controller renvoyant du JSON. 
    Peut en créer un basique, ou un en lui injectant un service déjà existant.
  `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("You should provide the name of the dto using the 'name' flag")
			os.Exit(1)
		}
		if ctrlJEntity {
			services.WriteJsonControllerWithServ(ctrlJName)
		} else {
			services.WriteJsonController(ctrlJName)
		}
	},
}

func init() {
	ctrlJCmd.Flags().BoolVarP(&ctrlJEntity, "entity", "e", false, "Lie le controller à une entité du même nom. Importe le service associé à l'entité.")
	ctrlJCmd.Flags().StringVarP(&ctrlJName, "name", "n", "", "Le nom du controller")
	ctrlJCmd.MarkFlagRequired("name")
}
