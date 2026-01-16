module github.com/JJPR030803/pos-cli

go 1.25.5

require (
	// Terminal color support - enables colors in Windows terminals
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect

	// Charm Bracelet TUI components - pre-built interactive UI elements (spinners, text inputs, lists)
	github.com/charmbracelet/bubbles v0.21.0 // indirect

	// Charm Bracelet TUI framework - The Elm Architecture for Go terminals (state management, rendering)
	github.com/charmbracelet/bubbletea v1.3.10 // indirect

	// Color profile detection - detects terminal color capabilities (true color, 256 colors, etc.)
	github.com/charmbracelet/colorprofile v0.2.3-0.20250311203215-f60798e515dc // indirect

	// Lipgloss - Style definitions for terminal output (like CSS for terminal)
	github.com/charmbracelet/lipgloss v1.1.0 // indirect

	// ANSI utilities - Low-level ANSI escape sequence handling
	github.com/charmbracelet/x/ansi v0.10.1 // indirect

	// Cell buffer - Terminal screen buffer management for rendering
	github.com/charmbracelet/x/cellbuf v0.0.13-0.20250311204145-2c3ea96c31dd // indirect

	// Terminal utilities - Terminal size detection, capabilities
	github.com/charmbracelet/x/term v0.2.1 // indirect

	// Windows console input - Better keyboard/mouse input on Windows
	github.com/erikgeiser/coninput v0.0.0-20211004153227-1c3628e74d0f // indirect

	// File system watcher - Watch files for changes (useful for hot reload)
	github.com/fsnotify/fsnotify v1.9.0 // indirect

	// Map decoder - Converts maps to structs (used by Viper for config parsing)
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect

	// Cobra helper - Detects mouse clicks on Windows (needed by Cobra)
	github.com/inconshreveable/mousetrap v1.1.0 // indirect

	// Color library - Color space conversions and manipulations
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect

	// TTY detection - Check if output is a terminal or pipe
	github.com/mattn/go-isatty v0.0.20 // indirect

	// Locale reader - Read system locale settings
	github.com/mattn/go-localereader v0.0.1 // indirect

	// Unicode width - Calculate display width of Unicode characters (for alignment)
	github.com/mattn/go-runewidth v0.0.16 // indirect

	// Color string parser - Parse color strings like "[red]text[reset]"
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect

	// ANSI parser - Parse ANSI escape sequences
	github.com/muesli/ansi v0.0.0-20230316100256-276c6243b2f6 // indirect

	// Cancel reader - Cancelable input reading for Bubbletea
	github.com/muesli/cancelreader v0.2.2 // indirect

	// Terminal environment - Query terminal capabilities and features
	github.com/muesli/termenv v0.16.0 // indirect

	// TOML parser - Parse TOML config files (your pos.toml)
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect

	// Unicode segmentation - Properly split Unicode strings (emojis, accents)
	github.com/rivo/uniseg v0.4.7 // indirect

	// Local file adapter - Afero adapter for local filesystem
	github.com/sagikazarmark/locafero v0.11.0 // indirect

	// Progress bar - Terminal progress indicators with ETA
	github.com/schollz/progressbar/v3 v3.19.0 // indirect

	// Concurrency utilities - Better error handling for goroutines
	github.com/sourcegraph/conc v0.3.1-0.20240121214520-5f936abd7ae8 // indirect

	// Abstract file system - Allows mocking filesystem for testing
	github.com/spf13/afero v1.15.0 // indirect

	// Type casting - Safe type conversions (used by Viper)
	github.com/spf13/cast v1.10.0 // indirect

	// Cobra CLI framework - THE standard for building CLI applications in Go
	github.com/spf13/cobra v1.10.2 // indirect

	// POSIX flags - Flag parsing library (used by Cobra)
	github.com/spf13/pflag v1.0.10 // indirect

	// Viper config management - Complete config solution (env vars, files, flags)
	github.com/spf13/viper v1.21.0 // indirect

	// .env file loader - Load environment variables from .env files
	github.com/subosito/gotenv v1.6.0 // indirect

	// Terminal info database - Terminal capability database
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect

	// YAML v3 parser - Parse YAML config files (Viper supports multiple formats)
	go.yaml.in/yaml/v3 v3.0.4 // indirect

	// System calls - Low-level OS interface (terminal control, signals)
	golang.org/x/sys v0.36.0 // indirect

	// Terminal package - Terminal control and raw mode
	golang.org/x/term v0.28.0 // indirect

	// Text processing - Unicode and encoding support
	golang.org/x/text v0.28.0 // indirect
)