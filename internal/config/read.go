package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

// Load:
// Validates TOML syntax and loading
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

	if err := cfg.validate(); err != nil {
		return nil, err
	}
	//Validate missing

	return cfg, nil

}

// Validate TOML file fields
func (c *Config) validate() error {
	// TODO better validation
	ve := &ValidationErrors{}

	// Required fields (no defaults)
	if c.Project.Name == "" {
		ve.Add("project.name", "required field cannot be empty")
	}

	if c.Project.Root == "" {
		ve.Add("project.root", "required field cannot be empty")
	}

	// Validate database configuration
	if c.Database.Type == "" {
		ve.Add("database.type", "database type must be specified (e.g., 'postgres', 'mysql')")
	}

	if c.Database.Name == "" {
		ve.Add("database.name", "database name is required")
	}

	// Validate workspace paths exist (optional but recommended)
	if c.Workspace.Frontend == "" {
		ve.Add("workspace.frontend", "frontend path is required")
	}

	if c.Workspace.Backend == "" {
		ve.Add("workspace.backend", "backend path is required")
	}

	//Version validation
	if c.Version == "" {
		ve.Add("version", "required field cannot be empty")
	} else if !isValidSemanticVersion(c.Version) {
		ve.Add("version", "invalid semantic version")
	}

	//Port validation
	if len(c.Database.Ports) == 0 {
		ve.Add("database.ports", "required field cannot be empty")
	} else {
		for i, port := range c.Database.Ports {
			if port < 1024 || port > 65535 {
				ve.Add(fmt.Sprintf("database.ports[%d]", i), fmt.Sprintf("port %d must be between 1024-65533", port))
			}
		}
	}

	if ve.HasErrors() {
		return ve
	}
	return nil

}

func isValidSemanticVersion(version string) bool {
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return false
	}
	for _, part := range parts {
		if part == "" {
			return false
		}
		for _, ch := range part {
			if ch < '0' || ch > '9' {
				return false
			}
		}
	}
	return true
}

func (cfg *Config) ApplyDefaults() {
	if cfg.Title == "" {
		cfg.Title = "CLI"
	}
}
