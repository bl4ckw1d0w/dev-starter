package cmd

import (
	"github.com/spf13/cobra"
	"github.com/bl4ckw1d0w/dev-starter/config"
)
// TODO: Mudar pra cada perfil quando finalizar esse perfil
// ProfileCmd comando para configurar o perfil de desenvolvedor web
var ProfileCmd = &cobra.Command{
	Use:   "Profiles",
	Short: "Configura o perfil de desenvolvedor",
	Long:  "Permite selecionar ferramentas específicas para desenvolvimento.",
	Run: func(cmd *cobra.Command, args []string) {
		// Chamando a função para configurar o perfil web
		config.ConfigureWebProfile()
	},
}
