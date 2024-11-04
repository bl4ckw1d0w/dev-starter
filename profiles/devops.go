package profiles

import (
    "fmt"
    "os/exec"
    "github.com/bl4ckw1d0w/dev-starter/config"
)

// InstallDevOpsProfile instala as ferramentas necessárias para o perfil DevOps.
func InstallDevOpsProfile() {
    // Obtém os comandos e mensagens para o perfil DevOps
    commands := configs.GetDevOpsCommands()

    // Executando cada comando
    for _, cmd := range commands {
        fmt.Println(cmd.Message)
        fmt.Print("Você aceita instalar esta ferramenta? (s/n): ")
        
        var response string
        fmt.Scanln(&response)
        
        if response != "s" && response != "S" {
            fmt.Println("🙅‍♀️ A instalação da ferramenta foi pulada. Seguindo para a próxima...")
            continue
        }

        fmt.Println("🔄 Instalando...")
        output, err := exec.Command("bash", "-c", cmd.Command).CombinedOutput()
        if err != nil {
            fmt.Printf("😱 Erro ao executar '%s': %s\n", cmd.Command, err)
            continue
        }
        fmt.Printf("✅ Saída: %s\n", output)
    }

    fmt.Println("🎉 Configuração do perfil DevOps concluída. Você está pronto para arrasar!")
}