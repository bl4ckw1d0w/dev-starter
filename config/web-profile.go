package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"

	"github.com/bl4ckw1d0w/dev-starter/utils"
)

// Função principal para configurar o perfil web
func ConfigureWebProfile() {
	fmt.Println("Configurando perfil de desenvolvedor web...")

	// Carrega o perfil do YAML
	profilePath := filepath.Join("profiles", "web-dev.yaml")
	profile, err := LoadProfile(profilePath)
	if err != nil {
		fmt.Printf("Erro ao carregar perfil: %v\n", err)
		return
	}

	// Definir as categorias e suas funções de configuração
	profileConfig := map[string]func(*Profile){
		"Editores e IDEs": chooseEditor,
	}

	// Iterar sobre o map e executar cada função de categoria
	for category, function := range profileConfig {
		fmt.Printf("\nConfigurando categoria: %s\n", category)
		function(profile) // Passa o perfil carregado para a função
	}
}

// Função para escolher e configurar o editor ou IDE
func chooseEditor(profile *Profile) {
	var editor string
	fmt.Println("Escolha um editor ou IDE (vscode, intellij, sublime):")
	fmt.Scanln(&editor)

	// Verifica se o editor está no YAML
	editorCommands, exists := profile.Commands["editors"][editor]
	if !exists {
		logrus.Errorf("Editor %s não encontrado no perfil.", editor)
	}

	// Executa os comandos do editor escolhido
	logrus.Infof("Instalando editor: %s", editor)
	for _, cmd := range editorCommands {
		err := utils.ExecuteCommand(cmd)
		if err != nil {
			logrus.Errorf("Erro ao executar comando '%s': %v", cmd, err)
		}
	}
	logrus.Infof("Editor %s configurado com sucesso!", editor)
}
