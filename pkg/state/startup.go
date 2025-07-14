package state

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"

	"github.com/mpm1900/rpg.sh/pkg/styles"
)

type Startup struct {
	id     string
	focus  string
	parent *GameState
}

func NewStartup() *Startup {
	return &Startup{
		id:    uuid.New().String(),
		focus: "start",
	}
}

func (s *Startup) SetParent(parent *GameState) {
	s.parent = parent
}

func (s Startup) ID() string {
	return s.id
}

func (s Startup) Status() string {
	return "STARTUP"
}

func (s *Startup) Init() tea.Cmd {
	return nil
}

func (s *Startup) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "right", "l":
			s.focus = "quit"
		case "left", "h":
			s.focus = "start"
		case "enter":
			if s.focus == "quit" {
				return s, tea.Quit
			}
		}
	}
	return s, nil
}

func (s *Startup) View() string {
	buf := strings.Builder{}

	startButton := styles.Button(styles.ButtonProps{
		Active: s.focus == "start",
	}).
		MarginBackground(styles.Card).
		MarginRight(2)
	quitButton := styles.Button(styles.ButtonProps{
		Active: s.focus == "quit",
	}).
		MarginBackground(styles.Card)

	buttons := lipgloss.JoinHorizontal(lipgloss.Top, startButton.Render("Start"), quitButton.Render("Quit"))
	ui := lipgloss.JoinVertical(lipgloss.Center, "rpg.ssh", buttons)

	dialog := lipgloss.Place(
		s.parent.Width,
		s.parent.Height,
		lipgloss.Center,
		lipgloss.Center,

		styles.DialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceBackground(styles.Bg),
		lipgloss.WithWhitespaceForeground(styles.Card),
		// lipgloss.WithWhitespaceChars("#"),
	)
	buf.WriteString(dialog)

	return buf.String()
}
