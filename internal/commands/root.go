package commands

import (
	"fmt"

	"github.com/JJPR030803/pos-cli/internal/commands/context"
	"github.com/JJPR030803/pos-cli/internal/commands/dev"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//TODO import subcommands example: internal/commands/dev etc
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "POS",
	Short: "POS Command Line Interface",
	Long:  "A comprehensive tooling for POS monorepo Command Line Interface",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	// Persistent flags (AVAILABLE TO ALL SUBCOMMANDS)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/configs/pos.toml)")

	// bind flags to viper so they can be checked whether they were activated
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolP("interactive", "i", false, "Activate interactive TUI") //TODO maybe
	err := viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	if err != nil {
		return
	}

	// TODO here add commands
	// Example: rootCmd.AddCommand(dev.DevCmd)
	rootCmd.AddCommand(dev.DevCmd)
	rootCmd.AddCommand(context.ContextCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./configs")
		viper.SetConfigName("pos")
		viper.SetConfigType("toml")
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if viper.GetBool("verbose") {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}
