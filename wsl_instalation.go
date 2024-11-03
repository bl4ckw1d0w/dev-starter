package main

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
func wslInstalled() bool {
    cmd := exec.Command("wsl", "--list")
    err := cmd.Run()
    return err == nil
}

// installWSL tenta instalar o WSL.
func installWSL() error {
    // Implementar lógica de instalação do WSL
    return nil
}

func main() {
    if isWindows() {
        fmt.Println("Detectando WSL...")
        if !wslInstalled() {
            fmt.Println("WSL não está instalado. Instalando...")
            if err := installWSL(); err != nil {
                fmt.Println("Erro ao tentar instalar o WSL:", err)
            } else {
                fmt.Println("WSL instalado com sucesso!")
            }
        } else {
            fmt.Println("WSL já está instalado.")
        }
    } else {
        fmt.Println("Este script só funciona no Windows!")
    }
}