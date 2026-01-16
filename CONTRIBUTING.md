# Contributing to POS CLI

Thank you for your interest in contributing to POS CLI! This document provides guidelines and information for contributors.

## Code of Conduct

Be respectful, constructive, and professional. This is a learning project and portfolio piece, so we welcome contributions of all experience levels.

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/pos-cli.git
   cd pos-cli
   ```
3. **Create a branch** for your changes:
   ```bash
   git checkout -b feature/my-new-feature
   ```

## Development Setup

### Prerequisites

- Go 1.25 or higher
- Familiarity with Go modules and packages
- Basic understanding of CLI development

### Install Dependencies

```bash
go mod download
```

### Run Tests

```bash
go test ./...
```

### Build and Run

```bash
# Build
go build -o pos ./cmd/pos

# Run without building
go run ./cmd/pos --help
```

## Project Structure

```
pos-cli/
├── cmd/pos/              # Entry point
├── internal/             # Private application code
│   ├── commands/        # CLI commands
│   ├── executor/        # Process execution
│   ├── config/          # Configuration
│   └── reporter/        # Output formatting
├── pkg/                 # Public, reusable packages
│   └── ui/             # UI components and styling
└── configs/            # Configuration examples
```

## Code Style

### Go Conventions

- Follow standard Go formatting: `gofmt` and `go vet`
- Use meaningful variable and function names
- Write comments for exported functions and types
- Keep functions focused and small

### Example

```go
// Good
func (o *Output) PrintSuccess(msg string) {
    fmt.Println(o.Styles.SuccessMsg(msg))
}

// Not ideal
func ps(m string) {
    fmt.Println(o.Styles.SuccessMsg(m))
}
```

### UI Styling

When adding UI elements, use the centralized `pkg/ui` package:

```go
// Good - uses centralized styling
out := ui.NewOutput(verbose)
out.PrintTitle("My Title")
out.PrintSuccess("Done!")

// Avoid - inline styling
titleStyle := lipgloss.NewStyle().Bold(true)...
fmt.Println(titleStyle.Render("My Title"))
```

## Adding New Commands

1. Create a new package in `internal/commands/`:
   ```bash
   mkdir internal/commands/mycommand
   ```

2. Create `mycommand.go`:
   ```go
   package mycommand

   import (
       "github.com/spf13/cobra"
       "github.com/JJPR030803/pos-cli/pkg/ui"
   )

   var MyCmd = &cobra.Command{
       Use:   "mycommand",
       Short: "Brief description",
       Long:  `Detailed description`,
       Run:   runMyCommand,
   }

   func init() {
       MyCmd.Flags().StringP("option", "o", "default", "Description")
   }

   func runMyCommand(cmd *cobra.Command, args []string) {
       out := ui.NewOutput(viper.GetBool("verbose"))
       out.PrintTitle("My Command")
       // Implementation
   }
   ```

3. Register in `internal/commands/root.go`:
   ```go
   import "github.com/JJPR030803/pos-cli/internal/commands/mycommand"
   
   func init() {
       rootCmd.AddCommand(mycommand.MyCmd)
   }
   ```

## Commit Messages

Use clear, descriptive commit messages:

```
Add context collection command

- Implement file system traversal
- Add exclusion patterns
- Support markdown and JSON output
```

Format:
- **First line**: Brief summary (50 chars or less)
- **Body**: Detailed explanation if needed
- **Footer**: Reference issues if applicable

## Pull Request Process

1. **Update documentation** if you've added features
2. **Add tests** for new functionality
3. **Ensure all tests pass**: `go test ./...`
4. **Update README.md** if needed
5. **Create a Pull Request** with a clear description

### PR Title Format

```
Add [feature]: Brief description
Fix [bug]: Brief description
Update [docs]: Brief description
```

## Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./pkg/ui

# Verbose output
go test -v ./...
```

## Documentation

- Comment exported functions and types
- Update README.md for new features
- Add examples in command descriptions
- Keep CONTRIBUTING.md current

## Questions?

Feel free to:
- Open an issue for discussion
- Reach out to the maintainer
- Ask in Pull Request comments

## Recognition

Contributors will be recognized in the README.md file.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to POS CLI! 🚀