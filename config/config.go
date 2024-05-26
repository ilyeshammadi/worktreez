package config

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Profile struct {
	RepositoriesPath string `yaml:"repositories_path"`
	WorktreesPath    string `yaml:"worktrees_path"`
}

type Config struct {
	WithIcons bool               `yaml:"icons"`
	Profiles  map[string]Profile `yaml:"profiles"`
}

var GlobalConfig Config

func SetGlobalConfig() {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	configPath := filepath.Join(usr.HomeDir, ".config", "worktreez", "config.yaml")
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, &GlobalConfig)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}
	fmt.Println(GlobalConfig)
}
