package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/mpm1900/rpg.sh/pkg/state"
	// "github.com/mpm1900/rpg.sh/pkg/styles"
)

func main() {
	{
		game := tea.NewProgram(state.InitialGameState(), tea.WithAltScreen())
		if _, err := game.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
