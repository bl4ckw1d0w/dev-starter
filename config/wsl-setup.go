package config

import (
	"os"
	"os/exec"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"
)

func executeCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func SetupWSL() {
	if runtime.GOOS == "windows" {
		// Habilitar WSL se necessário
		err := executeCommand("powershell.exe", "-Command", "Start-Process PowerShell -Verb RunAs -ArgumentList 'Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Windows-Subsystem-Linux -All'")
		if err != nil {
			log.Error("Opa! Não conseguimos habilitar o WSL:", err)
			return
		}
		log.Info("WSL habilitado com sucesso! Preparando o ambiente Linux no Windows...")

		// Instala o Debian
		err = executeCommand("powershell.exe", "-Command", "Start-Process PowerShell -Verb RunAs -ArgumentList 'wsl --install -d Debian'")
		if err != nil {
			log.Error("Algo deu errado ao instalar o Debian:", err)
			return
		}
		log.Info("Debian instalado! Linux está a caminho...")

		// Abre o Debian para configuração inicial
		cmdDebian := exec.Command("wsl", "-d", "Debian")
		err = cmdDebian.Start()
		if err != nil {
			log.Errorf("Erro ao abrir Debian para a configuração inicial: %s", err)
		} else {
			log.Info("Debian aberto para configuração inicial. Aproveite essa experiência Linux!")
		}

		// Aguarda o usuário terminar a configuração inicial
		time.Sleep(10 * time.Second)

		// Fecha o Debian para completar a configuração inicial
		err = executeCommand("wsl", "--shutdown")
		if err != nil {
			log.Error("Ops! Não conseguimos encerrar o WSL:", err)
		} else {
			log.Info("WSL encerrado. Vamos continuar a configuração!")
		}
	}

	// Reinicia o WSL
	err := executeCommand("wsl", "-d", "Debian")
	if err != nil {
		log.Errorf("Não conseguimos reiniciar o Debian: %s", err)
		return
	}
	}