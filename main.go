package main

import (
	// "bufio"
	// "os"
	// "os/exec"
	// "strings"

	"github.com/bl4ckw1d0w/dev-starter/config"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("inicializando")
	config.SetupWSL() // Inicializa o WSL para Windows (ou outros SOs no futuro)
	logrus.Info("Instalando o nix")	
	config.SetupNix()

}
