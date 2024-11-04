package main

import (
    "fmt"
    "github.com/bl4ckw1d0w/dev-starter/profiles"
	"github.com/bl4ckw1d0w/dev-starter/config"
)

func main() {
    fmt.Println("Bem-vindo ao Dev Starter!")
	    
	// Verifica o ambiente e instala o WSL, se necessário.
	config.WSLInstallation()	

    fmt.Println("Escolha um perfil para configurar:")
    fmt.Println("1 - DevOps")
    fmt.Println("2 - Desenvolvimento Web")

    // Captura a escolha do usuário
    var choice int
    fmt.Print("Digite o número do perfil desejado: ")
    _, err := fmt.Scan(&choice)
    if err != nil {
        fmt.Println("❌ Entrada inválida. Por favor, digite um número.")
        
    }

    // Executa a configuração do perfil com base na escolha do usuário
    switch choice {
    case 1:
        fmt.Println("🔧 Configurando o perfil DevOps...")
        profiles.InstallDevOpsProfile() // Função que configura o perfil DevOps
    case 2:
        fmt.Println("🌐 Configurando o perfil de Desenvolvimento Web...")
        profiles.InstallWebProfile() // Função que configura o perfil Web
    default:
        fmt.Println("❌ Opção inválida. Por favor, selecione 1 para DevOps ou 2 para Desenvolvimento Web.")
    }

    fmt.Println("Obrigado por usar o Dev Starter!")
}
