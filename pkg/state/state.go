package state

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/mpm1900/rpg.sh/pkg/styles"
)

type View interface {
	tea.Model
	ID() string
	Status() string
	SetParent(parent GameState)
}

type GameState struct {
	Height int
	Width  int

	ViewID string
	Views  map[string]View

	StatusText string
}

func InitialGameState() GameState {
	startup := NewStartup()
	state := GameState{
		Height: 0,
		Width:  0,
		ViewID: startup.ID(),
		Views: map[string]View{
			startup.ID(): startup,
		},
	}

	state.Views[startup.ID()] = startup
	return state
}

func (s GameState) Init() tea.Cmd {
	return nil
}

func (s GameState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.Height = msg.Height
		s.Width = msg.Width
		return s, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return s, tea.Quit
		}
	}

	view := s.Views[s.ViewID]
	if view == nil {
		return s, nil
	}
	view.SetParent(s)
	return view.Update(msg)
}

func (s GameState) View() string {
	view := s.Views[s.ViewID]
	if view == nil {
		return ""
	}

	buf := strings.Builder{}
	view.SetParent(s)
	buf.WriteString(view.View() + "\n")

	status := styles.StatusStyle.Render(view.Status())
	bar := styles.StatusBarStyle.
		Width(s.Width).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, status))
	buf.WriteString(bar)

	return buf.String()
}
