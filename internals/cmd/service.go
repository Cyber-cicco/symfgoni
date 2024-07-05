package cmd

import (
	"fmt"
	"fr/diginamic/go-cap/internals/services"
	"os"

	"github.com/spf13/cobra"
)


var srvEntity bool
var srvName string

var srvCmd = &cobra.Command{
	Use:   "srv",
	Short: "Crée une classe de service",
	Long: `
    Permet de générer une classe de service. Peut en créer un vide, ou un avec un mapper et un repository d'injecté.
  `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("You should provide the name of the service using the 'name' flag.")
			os.Exit(1)
		}
        if srvEntity {
            services.WriteServiceLinkedToEntity(srvName)
        } else {
            services.WriteService(srvName)
        }
	},
}

func init() {
	srvCmd.Flags().BoolVarP(&srvEntity, "entity", "e", false, "Lie le service à une entité du même nom. Importe le mapper et le repository associé à l'entité.")
	srvCmd.Flags().StringVarP(&srvName, "name", "n", "", "Le nom du service")
    srvCmd.MarkFlagRequired("name")
}
