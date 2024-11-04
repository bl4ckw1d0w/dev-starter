package config

import (
    "fmt"
    "os/exec"
    "runtime"
)

// isWindows verifica se o sistema operacional é Windows.
func isWindows() bool {
    return runtime.GOOS == "windows"
}

// wslInstalled verifica se o WSL está instalado.
func wslInstalled() (bool, error) {
    cmd := exec.Command("wsl", "--list")
    err := cmd.Run()
    return err == nil, err
}

// promptUser pergunta ao usuário se ele deseja instalar o WSL.
func promptUser() bool {
    fmt.Print("WSL não está instalado. Deseja instalar? (s/n): ")
    var response string
    fmt.Scanln(&response)
    return response == "s" || response == "S"
}

// installWSL tenta instalar o WSL.
func installWSL() error {
    fmt.Println("Iniciando a instalação do WSL...")
    cmd := exec.Command("wsl", "--install")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("falha ao instalar o WSL: %s", err)
    }
    fmt.Printf("Saída da instalação do WSL:\n%s\n", string(output))
    return nil
}

// WSLInstallation verifica e instala o WSL se necessário.
func WSLInstallation() {
    if !isWindows() {
        fmt.Println("⚠️ Este script só funciona no Windows!")
        return
    }

    fmt.Println("🔍 Detectando sistema operacional Windows...")

    installed, err := wslInstalled()
    if err != nil {
        fmt.Println("❌ Erro ao verificar instalação do WSL:", err)
        return
    }

    if installed {
        fmt.Println("✅ WSL já está instalado.")
        return
    }

    if promptUser() {
        if err := installWSL(); err != nil {
            fmt.Println("❌ Erro ao tentar instalar o WSL:", err)
        } else {
            fmt.Println("✅ WSL instalado com sucesso!")
        }
    } else {
        fmt.Println("🚫 Instalação do WSL cancelada pelo usuário.")
    }
}