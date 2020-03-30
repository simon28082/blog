package main

import (
    "github.com/crcms/blog/web/internal/app/commands"
    "github.com/crcms/blog/web/internal/app/providers"
    "github.com/firmeve/firmeve"
    "github.com/firmeve/firmeve/kernel/contract"
)

func main() {
    firmeve.RunDefault(
        firmeve.WithProviders([]contract.Provider{
           new(providers.AppProvider),
           new(providers.DocumentProvider),
        }),
        firmeve.WithCommands([]contract.Command{
           new(commands.MarkdownCommand),
        }),
    )
}
