package config

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

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
		return
	}
	err = yaml.Unmarshal(yamlFile, &GlobalConfig)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}
	GlobalConfig.Setup()
	defaultProfile, ok := GlobalConfig.Profiles["default"]
	if ok {
		fmt.Println(defaultProfile)
	}
}
