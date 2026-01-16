package config

type Config struct {
	Title     string        `toml:"title"`
	Version   string        `toml:"version"`
	Project   ProjectConfig `toml:"project"`
	Workspace Workspace     `toml:"workspace"`
	Commands  Commands      `toml:"commands"`
	Database  DatabaseCfg   `toml:"database"`
}

type ProjectConfig struct {
	Name string `toml:"name"`
	Root string `toml:"root"`
}

type Workspace struct {
	Frontend string `toml:"frontend"`
	Backend  string `toml:"backend"`
	Shared   string `toml:"shared"`
}

type Commands struct {
	Dev     CmdDev     `toml:"dev"`
	Test    CmdTest    `toml:"test"`
	Context CmdContext `toml:"context"`
}

type CmdDev struct {
	Frontend string     `toml:"frontend"`
	Backend  string     `toml:"backend"`
	Shared   string     `toml:"shared"`
	Context  CmdContext `toml:"context"`
}

type CmdTest struct {
	Coverage string `toml:"coverage"`
	Verbose  string `toml:"verbose"`
}

type CmdContext struct {
	Exclude   []string `toml:"exclude"`
	OutputFmt string   `toml:"output-format"`
}

type DatabaseCfg struct {
	Name          string `toml:"name"`
	Type          string `toml:"type"`
	Server        string `toml:"server"`
	Ports         []int  `toml:"port"`
	ConnectionMax int    `toml:"connection-max"`
}
