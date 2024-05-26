package config

import "path/filepath"

type Profile struct {
	RepositoriesPath string `yaml:"repositories_path"`
	WorktreesPath    string `yaml:"worktrees_path"`
}

func (profile *Profile) Setup() {
	if profile.WorktreesPath == "" {
		profile.WorktreesPath = filepath.Join(profile.RepositoriesPath, profile.WorktreesPath)
	}
}

type Config struct {
	WithIcons bool                `yaml:"icons"`
	Profiles  map[string]*Profile `yaml:"profiles"`
}

func (config *Config) Setup() {
	for _, profile := range config.Profiles {
		profile.Setup()
	}
}
