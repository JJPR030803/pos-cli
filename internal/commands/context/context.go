package context

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ContextCmd = &cobra.Command{
	Use:   "context",
	Short: "Collect project context and statistics/info for AI assistance",
	Long:  "Collect project context and statistics/info for AI assistance",
	RunE:  runContext,
}

func init() {
	ContextCmd.Flags().StringP("output", "o", "context.md", "Output file path")
	ContextCmd.Flags().StringSliceP("exclude", "e", []string{"node_modules", "dist"}, "Directories to exclude")
	ContextCmd.Flags().StringP("format", "f", "markdown", "Output format (markdown,json or txt)")

}

func runContext(cmd *cobra.Command, args []string) error {
	verbose := viper.GetBool("verbose")
	output, _ := cmd.Flags().GetString("output")
	exclude, _ := cmd.Flags().GetStringSlice("exclude")
	format, _ := cmd.Flags().GetString("format")

	if verbose {
		fmt.Println("Verbose:")
		fmt.Printf("%+v\n", verbose)
		fmt.Printf("Collecting context to: %s\n", output)
		fmt.Printf("Excluding directories: %v\n", exclude)
		fmt.Printf("Format: %s\n", format)
	}
	// Context collection here

	return nil
}
