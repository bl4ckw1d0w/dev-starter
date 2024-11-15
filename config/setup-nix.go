package config

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/sirupsen/logrus"
)

func SetupNix() {
	fmt.Println("Bem-vindo ao instalador do Nix no Debian WSL!")

	// Executa a instalação das dependências
	err := installDependencies()
	if err != nil {
		fmt.Printf("Erro ao instalar o Nix: %v\n", err)
	} else {
		fmt.Println("Nix instalado com sucesso!")
	}
}

func installDependencies() error {
	// Lista de comandos a serem executados no shell WSL
	commands := []string{
		"sudo apt update && sudo apt upgrade -y",        // Atualizar pacotes do sistema
		"sudo apt install -y curl sudo",                 // Instalar dependências
		"sudo apt install -y xz-utils",                 // Instalar dependências
		"curl -L https://nixos.org/nix/install | sh",    // Instalar o Nix
		`echo "source ~/.nix-profile/etc/profile.d/nix.sh" >> ~/.bashrc`, // Adicionar script de ambiente ao bashrc
		"source ~/.nix-profile/etc/profile.d/nix.sh", // Recarregar o ambiente no shell interativo
					}

	// Executa os comandos
	for _, command := range commands {
		err := executeCommand(command)
		if err != nil {
			logrus.Errorf("Erro ao executar '%s': %s", command, err)
			return err
		}
	}

	return nil
}

// Função que executa o comando no shell WSL
func executeCommand(cmdStr string) error {
	// Executa o comando no WSL (Debian) com shell bash
	cmd := exec.Command("wsl", "-d", "Debian", "bash", "-c", cmdStr)

	// Conecta a entrada, saída e erro do terminal
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // Permite a interação com o terminal para senha do sudo

	// Executa o comando e retorna o erro se houver
	return cmd.Run()
}
