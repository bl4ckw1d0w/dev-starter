package profiles

import (
    "fmt"
    "os/exec"
    "github.com/bl4ckw1d0w/dev-starter/config"
)

// promptUser recebe uma mensagem e retorna true se o usuário responder "s" (sim).
func promptUser(message string) bool {
    fmt.Print(message + " (s/n): ")
    var response string
    fmt.Scanln(&response)
    return response == "s" || response == "S"
}

// InstallWebProfile instala as ferramentas necessárias para o perfil de Desenvolvimento Web.
func InstallWebProfile() {
    // Pergunta se o usuário deseja instalar Node.js e NPM
    if promptUser("Você gostaria de instalar Node.js e NPM?") {
        // Obtém os comandos básicos de instalação para o perfil Web
        commands := configs.GetWebCommands()
        
        // Executa cada comando para instalar Node.js e NPM
        for _, cmd := range commands {
            fmt.Println(cmd.Message)
            output, err := exec.Command("bash", "-c", cmd.Command).CombinedOutput()
            if err != nil {
                fmt.Printf("😱 Erro ao executar '%s': %s\n", cmd.Command, err)
                continue
            }
            fmt.Printf("✅ Saída: %s\n", output)
        }
    } else {
        fmt.Println("🙅‍♀️ Node.js e NPM não serão instalados.")
    }

    // Pergunta ao usuário qual CLI de front-end ele deseja instalar
    fmt.Print("Qual CLI de front-end você gostaria de instalar? (1 = Angular, 2 = React, 0 = Nenhuma): ")
    var choice int
    fmt.Scan(&choice)

    switch choice {
    case 1:
        fmt.Println("📦 Instalando Angular CLI!")
        output, err := exec.Command("bash", "-c", "sudo npm install -g @angular/cli").CombinedOutput()
        if err != nil {
            fmt.Printf("😱 Erro ao instalar Angular CLI: %s\n", err)
        } else {
            fmt.Printf("✅ Angular CLI instalada com sucesso! Saída: %s\n", output)
        }
    case 2:
        fmt.Println("📦 Instalando Create React App!")
        output, err := exec.Command("bash", "-c", "sudo npm install -g create-react-app").CombinedOutput()
        if err != nil {
            fmt.Printf("😱 Erro ao instalar Create React App: %s\n", err)
        } else {
            fmt.Printf("✅ Create React App instalada com sucesso! Saída: %s\n", output)
        }
    default:
        fmt.Println("❌ Nenhuma CLI foi instalada.")
    }

    fmt.Println("🎉 Configuração do perfil Web concluída. Tudo pronto para o desenvolvimento front-end!")
}
