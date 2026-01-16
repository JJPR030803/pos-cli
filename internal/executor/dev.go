package executor

import (
	"context"
	"errors"
	"os"
	"os/exec"

	"github.com/JJPR030803/pos-cli/internal/commands/reporter"
	"github.com/JJPR030803/pos-cli/internal/config"
)

type Executor struct {
	config   *config.Config
	reporter *reporter.Reporter
	workDir  string
}

func New(config *config.Config, reporter *reporter.Reporter) *Executor {
	return &Executor{
		config:   config,
		reporter: reporter,
	}
}

type CommandResult struct {
	Output   string
	ExitCode int
	Error    error
}

func (e *Executor) runCommand(ctx context.Context, name string, args []string, dir string) *CommandResult {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Dir = dir

	cmd.Env = os.Environ()

	output, err := cmd.CombinedOutput()
	result := &CommandResult{
		Output: string(output),
	}

	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			result.ExitCode = exitError.ExitCode()
		}
		result.Error = err
	}
	return result
}

func (e *Executor) runCommandStreaming(ctx context.Context, name string, args []string, dir string) error {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Dir = dir
	cmd.Env = os.Environ()

	//Stream stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()

}
