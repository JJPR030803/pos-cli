package dev

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var DevCmd = &cobra.Command{
	Use:     "dev",
	Short:   "Start development servers",
	Long:    "Starts both frontend and backend development servers with hot reload",
	RunE:    runDev,
	Example: "pos dev",
}

func init() {
	DevCmd.Flags().BoolP("frontend-only", "f", false, "Run in frontend only")
	DevCmd.Flags().BoolP("backend-only", "b", false, "Run in backend only")
	DevCmd.Flags().IntP("port", "p", 8080, "Port to listen on")
	//Bind them if i want them to show on config
	err := viper.BindPFlag("dev.frontend-only", DevCmd.Flags().Lookup("frontend-only"))
	if err != nil {
		return
	}
}

func runDev(cmd *cobra.Command, args []string) error {
	verbose := viper.GetBool("verbose")
	frontendOnly, _ := cmd.Flags().GetBool("frontend-only")
	backendOnly, _ := cmd.Flags().GetBool("backend-only")
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4")).
		MarginTop(1)
	fmt.Println(titleStyle.Render("Starting development servers...."))
	if verbose {
		fmt.Printf("Frontend only: %t\n", frontendOnly)
		fmt.Printf("Backend only: %t\n", backendOnly)
		fmt.Printf("Port %d\n", viper.GetInt("port"))
	}

	// TODO use executor package to run bun commands
	// concurrent process management
	return nil
}
