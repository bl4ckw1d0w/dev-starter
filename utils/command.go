package utils

import (
	"os"
	"os/exec"
	"github.com/sirupsen/logrus"
)

// Executa um comando no WSL (Debian) com shell bash
func ExecuteCommand(cmdStr string) error {
	// Log da execução do comando (útil para debugging)
	logrus.Infof("Executando comando: %s", cmdStr)

	// Configura o comando para rodar no WSL
	cmd := exec.Command("wsl", "-d", "Debian", "bash", "-c", cmdStr)

	// Conecta a entrada, saída e erro do terminal
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Executa o comando e retorna o erro, se houver
	err := cmd.Run()
	if err != nil {
		logrus.Errorf("Erro ao executar comando: %v", err)
	}
	return err
}
