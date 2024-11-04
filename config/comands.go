
package config

// CommandInfo armazena o comando e a mensagem associada.
type CommandInfo struct {
    Command string // O comando a ser executado
    Message string // A mensagem a ser exibida para o usuário
}

// GetDevOpsCommands retorna os comandos e mensagens para o perfil DevOps.
func GetDevOpsCommands() []CommandInfo {
    return []CommandInfo{
        {"sudo apt update", "🛠️ Hora de atualizar os pacotes! Preparando o terreno..."},
        {"sudo apt install -y docker.io", "🐳 Vamos colocar o Docker a bordo! Pronto para navegar?"},
        {"curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -", "🔑 Adicionando a chave do Kubernetes! Como se fosse um segredo..."},
        {"echo 'deb https://apt.kubernetes.io/ kubernetes-xenial main' | sudo tee /etc/apt/sources.list.d/kubernetes.list", "📦 Preparando o repositório do Kubernetes! Está quase lá!"},
        {"sudo apt update", "🔄 Atualizando novamente! Quase pronto para a festa do Kubernetes!"},
        {"sudo apt install -y kubectl", "🚀 Instalando o Kubernetes CLI (kubectl)! Vamos ao espaço!"},
        {"sudo apt install -y software-properties-common", "📦 Instalando dependências do Terraform! Preparando o caminho para a magia!"},
        {"sudo add-apt-repository ppa:hashicorp/ppa -y", "🧙‍♂️ Adicionando o repositório do Terraform! Feitiços a caminho!"},
        {"sudo apt update", "🔄 Atualizando de novo! Um último toque antes do grande final!"},
        {"sudo apt install -y terraform", "⚡ Instalando o Terraform! Preparando o terreno para a infraestrutura mágica!"},
    }
}
func GetWebCommands() []CommandInfo {
    return []CommandInfo{
        {"sudo apt update", "🔄 Atualizando pacotes do sistema..."},
        {"sudo apt install -y nodejs", "🟢 Instalando Node.js, o motor do JavaScript!"},
        {"sudo apt install -y npm", "📦 Instalando NPM, o gerenciador de pacotes do Node!"},
    }
}
