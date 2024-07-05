package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "go-cap",
  Short: "Génération de code pour le stack symfony de Digicap",
  Long: 
`
    Offre des utilitaires de génération de code supplémentaires pour symfony
    Permet de créer des services, des mappers, des interfaces, des controllers json et des controllers html
    Permet également de transformer les SVG classiques en SVG adaptés à Twig.
`,
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}

func init() {
    rootCmd.AddCommand(dtoCmd)
}
