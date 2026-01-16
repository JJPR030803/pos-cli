package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

func Load(path string) (*Config, error) {
	cfg := &Config{
		Title:   "pos-cli",
		Version: "1.0",
		Project: ProjectConfig{
			Name: "Pos CLI",
			Root: ".",
		},
		Workspace: Workspace{
			Frontend: "frontend/src",
			Backend:  "backend/src",
			Shared:   "shared/src",
		},
		Commands: Commands{
			Dev: CmdDev{
				Frontend: "bun run dev",
				Backend:  "bun run backend",
				Shared:   "",
				Context: CmdContext{
					Exclude:   []string{"node_modules", ".git", ".idea"},
					OutputFmt: "text",
				},
			},
			Test:    CmdTest{},
			Context: CmdContext{},
		},
		Database: DatabaseCfg{
			Name:          "pos_db",
			Type:          "postgres",
			Server:        "8000",
			Ports:         []int{8000},
			ConnectionMax: 5000,
		},
	}

	file, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return cfg, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := toml.Unmarshal(file, cfg); err != nil {
		return nil, fmt.Errorf("Invalid TOML syntax: %w", err)
	}

	//Validate missing

	return cfg, nil

}

func (c *Config) validate() error {
	// TODO better validation
	if c.Workspace.Frontend == "" {
		c.Workspace.Frontend = "frontend/src"
	}
	if c.Workspace.Backend == "" {
		c.Workspace.Backend = "backend/src"
	}
	if c.Workspace.Shared == "" {
		c.Workspace.Shared = "shared/src"
	}
	if c.Database.Name == "" {
		c.Database.Name = "pos_db"
	}

	return nil

}
