package commands

import (
	"fmt"

	"github.com/devsimsek/plasm/core"
	"github.com/devsimsek/plasm/types"
)

func init() {
	core.RegisterCommand(types.Command{
		Name:  "hello",
		Alias: []string{"h"},
		Flags: []types.Flag{
			{Name: "name", Alias: []string{"n"}, Value: ""},
		},
		Description: "Say hello to someone",
		Handler: func(flags []types.Flag) {
			name := core.FindFlag(flags, "name")
			if name == "" {
				fmt.Println("Hello, World!")
			} else {
				fmt.Printf("Hello, %s!\n", name)
			}
		},
	})
}
