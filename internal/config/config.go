package config

// Config represents the overall structure of the application's configuration.
// It includes settings for the project, workspace, commands, and database.
type Config struct {
	// Title of the project.
	Title string `toml:"title"`
	// Version of the application.
	Version string `toml:"version"`
	// Project-specific configuration.
	Project ProjectConfig `toml:"project"`
	// Workspace directory configurations.
	Workspace Workspace `toml:"workspace"`
	// Command-specific configurations.
	Commands Commands `toml:"commands"`
	// Database connection settings.
	Database DatabaseCfg `toml:"database"`
}

// ProjectConfig holds configuration specific to the project.
type ProjectConfig struct {
	// Name of the project.
	Name string `toml:"name"`
	// Root directory of the project.
	Root string `toml:"root"`
}

// Workspace defines the paths for different parts of the workspace.
type Workspace struct {
	// Path to the frontend directory.
	Frontend string `toml:"frontend"`
	// Path to the backend directory.
	Backend string `toml:"backend"`
	// Path to the shared directory.
	Shared string `toml:"shared"`
}

// Commands holds configurations for various command-line commands.
type Commands struct {
	// Development command configuration.
	Dev CmdDev `toml:"dev"`
	// Test command configuration.
	Test CmdTest `toml:"test"`
	// Context command configuration.
	Context CmdContext `toml:"context"`
}

// CmdDev configures the development command.
type CmdDev struct {
	// Frontend-specific settings for the dev command.
	Frontend string `toml:"frontend"`
	// Backend-specific settings for the dev command.
	Backend string `toml:"backend"`
	// Shared settings for the dev command.
	Shared string `toml:"shared"`
	// Context-specific settings for the dev command.
	Context CmdContext `toml:"context"`
}

// CmdTest configures the test command.
type CmdTest struct {
	// Coverage settings for the test command.
	Coverage string `toml:"coverage"`
	// Verbose flag for the test command.
	Verbose string `toml:"verbose"`
}

// CmdContext configures the context command.
type CmdContext struct {
	// List of patterns to exclude.
	Exclude []string `toml:"exclude"`
	// Output format for the context command.
	OutputFmt string `toml:"output-format"`
}

// DatabaseCfg holds configuration for database connection.
type DatabaseCfg struct {
	// Name of the database.
	Name string `toml:"name"`
	// Type of the database (e.g., "mysql", "postgres").
	Type string `toml:"type"`
	// Server address for the database.
	Server string `toml:"server"`
	// List of ports to connect to the database.
	Ports []int `toml:"port"`
	// Maximum number of connections allowed to the database.
	ConnectionMax int `toml:"connection-max"`
}
