package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JJPR030803/pos-cli/internal/commands"
	"github.com/JJPR030803/pos-cli/internal/config"
	"github.com/JJPR030803/pos-cli/pkg/ui"
)

func main() {
	cfg, err := config.Load("configs/pos.toml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	fmt.Println(cfg.Database.Name, cfg.Title)
	if err := commands.Execute(); err != nil {
		os.Exit(1)
	}
}

func testUI() {
	// Create output handler (verbose = true for demo)
	out := ui.NewOutput(true)

	// Test all the different outputs
	out.PrintTitle("🎨 UI Package Demo")
	out.PrintSubtitle("Testing all available styles and functions")
	out.PrintSeparator()

	// Test section headers
	out.PrintSection("Message Types")
	out.PrintSuccess("This is a success message")
	out.PrintError("This is an error message")
	out.PrintWarning("This is a warning message")
	out.PrintInfo("This is an info message")
	out.PrintVerbose("This is a verbose message (only in verbose mode)")
	out.Newline()

	// Test status indicators
	out.PrintSection("Status Indicators")
	out.PrintStatus("Frontend Server", true)
	out.PrintStatus("Backend Server", true)
	out.PrintStatus("Database", false)
	out.Newline()

	// Test command display
	out.PrintSection("Commands")
	out.PrintCommand("bun run dev")
	out.PrintCommand("go run ./cmd/pos")
	out.Newline()

	// Test list
	out.PrintSection("Features")
	out.PrintList([]string{
		"Centralized styling",
		"Consistent icons",
		"Industrial aesthetic",
		"Easy to maintain",
	})
	out.Newline()

	// Test table
	out.PrintSection("Configuration")
	out.PrintTable([][2]string{
		{"Frontend Port", "3000"},
		{"Backend Port", "3001"},
		{"Hot Reload", "Enabled"},
		{"Verbose Mode", "true"},
	})
	out.Newline()

	// Test key-value pairs
	out.PrintSection("Environment")
	out.PrintKeyValue("Node Version", "v20.11.0")
	out.PrintKeyValue("Bun Version", "1.1.0")
	out.Newline()

	// Test box
	out.PrintSection("Boxed Content")
	out.PrintBox("This content is in a box!\nMultiple lines work too.\nPretty cool, right?")
	out.Newline()

	// Test badge
	out.PrintSection("Badges")
	out.PrintBadge("PRODUCTION")
	out.PrintBadge("DEVELOPMENT")
	out.Newline()

	// Test spinner (basic version)
	out.PrintSection("Spinner Demo")
	out.StartSpinner("Loading configuration")
	// Simulate some work...
	out.StopSpinner(true, "Configuration loaded successfully")
	out.Newline()

	out.PrintSeparator()
	out.PrintSuccess("UI Package is working perfectly!")
}
