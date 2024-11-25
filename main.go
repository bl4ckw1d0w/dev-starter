package main

import (
	"log"
	"os"
	"github.com/bl4ckw1d0w/dev-starter/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "wsl-tool",
		Short: "WSL Tool - Ferramenta para automatizar configurações do WSL",
		Long:  "WSL Tool permite configurar e gerenciar ambientes no WSL com perfis personalizados, backups e restauração.",
	}

	// Subcomandos
	rootCmd.AddCommand(cmd.SetupCmd)

	// Executa a CLI
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Erro ao executar: %v", err)
		os.Exit(1)
	}
}
