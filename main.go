package main

import (
	"bufio"
	"os"
	"os/exec"
	"strings"

	"github.com/bl4ckw1d0w/dev-starter/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.SetupWSL() // Inicializa o WSL para Windows (ou outros SOs no futuro)

// Escolha do perfil
reader := bufio.NewReader(os.Stdin)

log.Info("Escolha um perfil para começar:")
log.Info("1 - Desenvolvimento Web")
log.Info("2 - DevOps")
log.Info("Digite o número do perfil e pressione Enter.")

profileChoice, err := reader.ReadString('\n')
if err != nil {
	log.Error("Houve um problema ao ler sua escolha:", err)
}
profileChoice = strings.TrimSpace(profileChoice)

var commands []string

switch profileChoice {
case "1":
	commands = []string{
		"sudo apt update",
		"sudo apt upgrade -y",
		"sudo apt install -y git",
		"sudo apt install -y software-properties-common apt-transport-https wget",
		// TODO: aicionar comando para instalar o VS Code
	}
	log.Info("Você escolheu Desenvolvimento Web! Preparando o ambiente...")
case "2":
	commands = []string{
		"sudo apt update",
		"sudo apt install docker.io -y",
		//TODO: adicionar comando para instalar ferramentas devops
	}
	log.Info("Perfil DevOps escolhido! Instalando as ferramentas para automação.")
default:
	log.Warn("Perfil inválido. Por favor, execute o programa novamente e escolha um perfil válido.")
	return
}

log.Info("Executando os comandos de configuração. Isso pode demorar um pouco, então relaxe!")

for _, cmd := range commands {
	log.Infof("Para executar o comando '%s', precisamos da sua senha sudo.", cmd)
	log.Info("Digite a senha sudo e pressione Enter:")

	sudoPassword, err := reader.ReadString('\n')
	if err != nil {
		log.Error("Erro ao ler a senha sudo:", err)
		return
	}
	sudoPassword = strings.TrimSpace(sudoPassword)

	execCmd := exec.Command("wsl", "-d", "Debian", "--", "sh", "-c", "echo "+sudoPassword+" | sudo -S sh -c \""+cmd+"\"")
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	err = execCmd.Run()
	if err != nil {
		log.Errorf("Algo deu errado ao executar o comando '%s': %s", cmd, err)
	} else {
		log.Infof("Comando '%s' executado com sucesso!", cmd)
	}
}
log.Info("Configuração concluída! Aproveite o seu ambiente de desenvolvimento.")
}

