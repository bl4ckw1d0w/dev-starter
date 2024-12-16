package config

import (
	"testing"
)

// MockCommandExecutor é um mock para CommandExecutor
type MockCommandExecutor struct {
	Calls []string // Armazena os comandos executados
	Err   error    // Simula um erro, se necessário
}

func (m *MockCommandExecutor) Execute(name string, arg ...string) error {
	// Registra o comando chamado
	command := name + " " + concatArgs(arg)
	m.Calls = append(m.Calls, command)
	return m.Err
}

func concatArgs(args []string) string {
	result := ""
	for _, arg := range args {
		result += arg + " "
	}
	return result[:len(result)-1]
}

func TestSetupWSL(t *testing.T) {
	mockExecutor := &MockCommandExecutor{}

	// Executa a função com o mock
	SetupWSL(mockExecutor)

	// Verifica se os comandos esperados foram chamados
	expectedCommands := []string{
		"powershell.exe -Command Start-Process PowerShell -Verb RunAs -ArgumentList 'Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Windows-Subsystem-Linux -All'",
		"powershell.exe -Command Start-Process PowerShell -Verb RunAs -ArgumentList 'wsl --install -d Debian'",
		"wsl -d Debian",
		"wsl --shutdown",
		"wsl -d Debian",
	}

	for i, expected := range expectedCommands {
		if i >= len(mockExecutor.Calls) {
			t.Fatalf("Esperava mais comandos: %s", expected)
		}

		if mockExecutor.Calls[i] != expected {
			t.Errorf("Comando #%d incorreto: esperado '%s', obteve '%s'", i+1, expected, mockExecutor.Calls[i])
		}
	}

	if len(mockExecutor.Calls) != len(expectedCommands) {
		t.Errorf("Número de comandos incorreto: esperado %d, obteve %d", len(expectedCommands), len(mockExecutor.Calls))
	}
}
