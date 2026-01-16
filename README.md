# POS CLI

A professional command-line tool for managing Bun monorepo development workflows with style.

![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-green.svg)

## Overview

`pos` is a CLI tool designed to streamline development workflows for Bun-based monorepos. Built with Go, it provides a consistent, beautiful terminal interface for common development tasks like starting servers, running tests, and managing your project.

Originally developed for the [web-pos](https://github.com/JJPR030803/web-pos) project, it has been extracted as a standalone tool with broader applicability to any Bun workspace project.

## Features

- 🚀 **Concurrent Server Management** - Start multiple dev servers with a single command
- 🧪 **Unified Testing** - Run tests across all workspaces
- 📝 **Context Collection** - Generate project context files for AI assistance
- 🎨 **Beautiful Terminal UI** - Industrial aesthetic with Lipgloss styling
- ⚙️ **Flexible Configuration** - TOML-based configuration with sensible defaults
- 🔧 **Extensible Architecture** - Built with Cobra for easy command additions

## Installation

### From Source

```bash
go install github.com/JJPR030803/pos-cli/cmd/pos@latest
```

Make sure `$GOPATH/bin` is in your PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### From Release (Coming Soon)

Download pre-built binaries from the [releases page](https://github.com/JJPR030803/pos-cli/releases).

## Quick Start

```bash
# Start all development servers
pos dev

# Start only frontend
pos dev --frontend-only

# Run tests with coverage
pos test --coverage

# Collect project context for AI
pos context --output context.md

# Get help
pos --help
```

## Usage

### Development Servers

```bash
# Start all servers
pos dev

# Frontend only
pos dev --frontend-only
pos dev -f

# Backend only  
pos dev --backend-only
pos dev -b

# Custom port
pos dev --port 3000

# Verbose output
pos dev --verbose
pos dev -v
```

### Testing

```bash
# Run all tests
pos test

# With coverage report
pos test --coverage
pos test -c

# Watch mode
pos test --watch
pos test -w
```

### Context Collection

Generate a comprehensive context file for AI assistants:

```bash
# Default output (context.md)
pos context

# Custom output path
pos context --output my-context.md
pos context -o my-context.md

# Exclude directories
pos context --exclude node_modules,dist,.turbo

# JSON format
pos context --format json
```

### Configuration

```bash
# Use custom config file
pos --config ./custom-config.toml dev

# Global verbose mode
pos --verbose dev
pos -v test
```

## Configuration

Create a `pos.toml` in your project root or in a `configs/` directory:

```toml
[project]
name = "my-project"
root = "."

[workspace]
frontend = "apps/frontend"
backend = "apps/backend"
shared = "packages/shared"

[commands.dev]
frontend_cmd = "bun run dev"
backend_cmd = "bun run dev"

[commands.test]
coverage = true
verbose = false

[commands.context]
exclude = ["node_modules", "dist", "build", ".turbo"]
output_format = "markdown"
```

## Project Structure

This tool works best with Bun workspaces structured like:

```
your-project/
├── apps/
│   ├── frontend/
│   │   └── package.json
│   └── backend/
│       └── package.json
├── packages/
│   └── shared/
│       └── package.json
├── package.json          # Root workspace config
└── pos.toml             # Optional pos-cli config
```

## Tech Stack

- **[Go](https://golang.org/)** - Modern, efficient systems language
- **[Cobra](https://github.com/spf13/cobra)** - Industry-standard CLI framework
- **[Viper](https://github.com/spf13/viper)** - Complete configuration solution
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)** - Terminal styling framework
- **[Bubbletea](https://github.com/charmbracelet/bubbletea)** - TUI framework (planned)

## Development

### Prerequisites

- Go 1.25 or higher
- Bun (for testing with actual projects)

### Setup

```bash
# Clone the repository
git clone https://github.com/JJPR030803/pos-cli.git
cd pos-cli

# Install dependencies
go mod download

# Build
go build -o pos ./cmd/pos

# Run
./pos --help

# Or run directly without building
go run ./cmd/pos dev
```

### Project Structure

```
pos-cli/
├── cmd/
│   └── pos/
│       └── main.go           # Entry point
├── internal/
│   ├── commands/             # CLI commands
│   │   ├── root.go          # Root command setup
│   │   ├── dev/             # Development commands
│   │   ├── test/            # Testing commands
│   │   └── context/         # Context collection
│   ├── executor/            # Process execution
│   ├── config/              # Configuration handling
│   └── reporter/            # Output formatting
├── pkg/
│   └── ui/                  # Reusable UI components
│       ├── styles.go        # Styling definitions
│       └── output.go        # Output helpers
├── configs/
│   └── pos.toml.example     # Example configuration
├── go.mod
└── README.md
```

### Adding New Commands

```go
package mycommand

import (
    "github.com/spf13/cobra"
    "github.com/JJPR030803/pos-cli/pkg/ui"
)

var MyCmd = &cobra.Command{
    Use:   "mycommand",
    Short: "Description of my command",
    Run:   runMyCommand,
}

func init() {
    MyCmd.Flags().StringP("option", "o", "default", "Option description")
}

func runMyCommand(cmd *cobra.Command, args []string) {
    out := ui.NewOutput(viper.GetBool("verbose"))
    out.PrintTitle("My Command")
    // Your logic here
}
```

Then register it in `internal/commands/root.go`:

```go
import "github.com/JJPR030803/pos-cli/internal/commands/mycommand"

func init() {
    rootCmd.AddCommand(mycommand.MyCmd)
}
```

## Roadmap

- [x] Basic command structure
- [x] Centralized UI styling
- [x] Development server management
- [ ] Actual process execution
- [ ] Test command implementation
- [ ] Context collection implementation
- [ ] Database management commands
- [ ] Interactive TUI mode (optional)
- [ ] Configuration wizard
- [ ] Plugin system
- [ ] Release binaries for all platforms

## Contributing

Contributions are welcome! This project is primarily a portfolio piece and learning project, but feel free to:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Example Projects

- [web-pos](https://github.com/JJPR030803/web-pos) - A modern Point of Sale system built with React, TypeScript, and Elysia

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Cobra](https://cobra.dev/) CLI framework
- Styled with [Charm Bracelet](https://charm.sh/) tools
- Inspired by modern CLI tools like [gh](https://cli.github.com/), [k9s](https://k9scli.io/), and [lazygit](https://github.com/jesseduffield/lazygit)

## Author

**Juan Julian** - [JJPR030803](https://github.com/JJPR030803)

---

<p align="center">
  Made with ❤️ and Go
</p>