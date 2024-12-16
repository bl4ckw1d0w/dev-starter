package config

import (
	"os"
	"os/exec"
	"runtime"
	"time"

	logrus "github.com/sirupsen/logrus"
)

// CommandExecutor define uma interface para executar comandos
type CommandExecutor interface {
	Execute(name string, arg ...string) error
}

// RealCommandExecutor implementa CommandExecutor e executa comandos reais
type RealCommandExecutor struct{}

func (r RealCommandExecutor) Execute(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

// SetupWSL Configura o WSL, instala o Debian e prepara o ambiente
func SetupWSL(executor CommandExecutor) {
	if runtime.GOOS == "windows" {
		// Habilitar WSL se necessário
		err := executor.Execute("powershell.exe", "-Command", "Start-Process PowerShell -Verb RunAs -ArgumentList 'Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Windows-Subsystem-Linux -All'")
		if err != nil {
			logrus.Error("Opa! Não conseguimos habilitar o WSL:", err)
		}
		logrus.Info("WSL habilitado com sucesso! Preparando o ambiente Linux no Windows...")

		// Instala o Debian
		err = executor.Execute("powershell.exe", "-Command", "Start-Process PowerShell -Verb RunAs -ArgumentList 'wsl --install -d Debian'")
		if err != nil {
			logrus.Error("Algo deu errado ao instalar o Debian:", err)
		}
		logrus.Info("Debian instalado! Linux está a caminho...")

		// Abre o Debian para configuração inicial
		err = executor.Execute("wsl", "-d", "Debian")
		if err != nil {
			logrus.Errorf("Erro ao abrir Debian para a configuração inicial: %s", err)
		} else {
			logrus.Info("Debian aberto para configuração inicial. Aproveite essa experiência Linux!")
		}

		// Aguarda o usuário terminar a configuração inicial
		time.Sleep(30 * time.Second)

		// Fecha o Debian para completar a configuração inicial
		err = executor.Execute("wsl", "--shutdown")
		if err != nil {
			logrus.Error("Ops! Não conseguimos encerrar o WSL:", err)
		} else {
			logrus.Info("WSL encerrado. Vamos continuar a configuração!")
		}
	}

	// Reinicia o WSL
	err := executor.Execute("wsl", "-d", "Debian")
	if err != nil {
		logrus.Errorf("Não conseguimos reiniciar o Debian: %s", err)
	}
}
