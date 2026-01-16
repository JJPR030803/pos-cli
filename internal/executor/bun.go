package executor

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

type BunCommand struct {
	Workspace string // e.g., "frontend", "backend"
	Script    string // dev, build, test
	Args      []string
}

func (e *Executor) RunBunScript(ctx context.Context, workspace, script string, args []string) error {
	workspacePath := filepath.Join(e.workDir, workspace)

	if _, err := os.Stat(workspacePath); os.IsNotExist(err) {
		return fmt.Errorf("workspace %s does not exist at %s", workspace, workspacePath)
	}

	cmdArgs := []string{"run", script}
	cmdArgs = append(cmdArgs, args...)

	// TODO implement reporter e.reporter

	return e.runCommandStreaming(ctx, "bun", cmdArgs, workspacePath)
}

func (e *Executor) RunBunInstall(ctx context.Context) error {

}
