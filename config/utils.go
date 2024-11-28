package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

// Estrutura do Perfil
type Profile struct {
	Name     string                       `yaml:"name"`
	Commands map[string]map[string][]string `yaml:"commands"`
}

// Carrega o perfil do arquivo YAML
func LoadProfile(path string) (*Profile, error) {
	// Abre o arquivo
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decodifica o arquivo YAML diretamente
	var profile Profile
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}
