package cmd

import (
	"fmt"
	"fr/diginamic/go-cap/internals/services"
	"os"

	"github.com/spf13/cobra"
)

var ctrlHName string

var ctrlHCmd = &cobra.Command{
	Use:   "html-ctrl",
	Short: "Crée un Controller HTML",
	Long: `
    Permet de générer une classe controller renvoyant du HTML, ainsi que le template twig à l'endroit correspondant. 
    Peut en créer un basique, ou un en lui injectant un service déjà existant.
  `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("You should provide the name of the dto using the 'name' flag")
			os.Exit(1)
		}
        if ctrlHName == "" {
			fmt.Println("Can't provide blank name for this command. Exiting")
			os.Exit(1)
        }
        services.WriteHtmlController(ctrlHName)
	},
}

func init() {
	ctrlHCmd.Flags().StringVarP(&ctrlHName, "name", "n", "", "Le nom du controller")
	ctrlHCmd.MarkFlagRequired("name")
}
