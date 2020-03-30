package commands

import (
	kernel2 "github.com/firmeve/firmeve/bootstrap"
	"github.com/firmeve/firmeve/kernel"
	"github.com/firmeve/firmeve/kernel/contract"
	"github.com/spf13/cobra"
)

type MarkdownCommand struct {
	kernel.Command
	command *cobra.Command
}

func (c *MarkdownCommand) newCmd() *cobra.Command {
	c.command = new(cobra.Command)
	c.command.Use = "markdown"
	c.command.Short = "Parse all markdown file"
	c.command.Flags().StringP("path", "", "", "file path")

	c.command.Run = c.run

	return c.command
}

func (m *MarkdownCommand) Cmd() *cobra.Command {
	if m.command == nil {
		m.command = c.newCmd()
	}

	return c.command
}

func (m MarkdownCommand) SetApplication(app contract.Application) {
	panic("implement me")
}

func (m MarkdownCommand) SetProviders(providers []contract.Provider) {
	panic("implement me")
}

func (m MarkdownCommand) Application() contract.Application {
	panic("implement me")
}

func (m MarkdownCommand) Providers() []contract.Provider {
	panic("implement me")
}

func (c *MarkdownCommand) run(cmd *cobra.Command, args []string) {
	// bootstrap
	kernel2.BootFromCommand(c)
}