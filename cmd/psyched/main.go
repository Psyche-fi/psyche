package main

import (
	"os"

	"github.com/psyche-fi/psyche/cmd/psyched/cmd"
    svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/psyche-fi/psyche/app"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
    if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
