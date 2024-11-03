package main

import (
    "os/exec"
    "testing"
)

// Testa a verificação do sistema operacional Windows.
func TestIsWindows(t *testing.T) {
    if !isWindows() {
        t.Error("Esperado true para sistema Windows, mas obteve false.")
    }
}

// Testa a detecção do WSL.
func TestWslInstalled(t *testing.T) {
    // Mock da função exec.Command para simular a presença do WSL.
    cmd := exec.Command("wsl", "--list")
    err := cmd.Run()
    if err != nil {
        t.Error("Esperado que o WSL esteja instalado, mas ocorreu um erro:", err)
    }
}

// Testa a instalação do WSL.
func TestInstallWSL(t *testing.T) {
    if isWindows() {
        err := installWSL()
        if err != nil {
            t.Error("Esperado sucesso na instalação do WSL, mas ocorreu um erro:", err)
        }
    } else {
        t.Skip("Este teste é apenas para sistemas Windows.")
    }
}