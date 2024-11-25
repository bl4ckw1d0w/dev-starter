package cmd

import (
	"github.com/spf13/cobra"
	"github.com/bl4ckw1d0w/dev-starter/config"
)

// SetupCmd adiciona o comando "setup" à CLI
var SetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configura o ambiente WSL com Debian e Nix",
	Long: `Habilita o WSL no Windows, instala a distribuição Debian e prepara o ambiente com o gerenciador de pacotes Nix.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Executa o setup do WSL e do Nix
		config.SetupWSL()
		config.SetupNix()
	},
}
