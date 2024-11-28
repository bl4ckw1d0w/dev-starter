package main

import (
	"log"
	"os"
	"github.com/bl4ckw1d0w/dev-starter/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "dev-starter",
		Short: "Dev Starter - Ferramenta para automatizar configurações de ambiente",
		Long:  "Dev Starter permite configurar e gerenciar ambientes no WSL com perfis personalizados, backups e restauração.",
	}

	// Subcomandos
	rootCmd.AddCommand(cmd.SetupCmd, cmd.ProfileCmd)

	// Executa a CLI
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Erro ao executar: %v", err)
		os.Exit(1)
	}
}
