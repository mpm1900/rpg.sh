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
	parent GameState
}

func NewStartup() *Startup {
	return &Startup{
		id:    uuid.New().String(),
		focus: "start",
	}
}

func (s *Startup) SetParent(parent GameState) {
	s.parent = parent
}

func (s Startup) ID() string {
	return s.id
}

func (s Startup) Status() string {
	return "STARTUP"
}

func (s Startup) Init() tea.Cmd {
	return nil
}

func (s Startup) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}

func (s Startup) View() string {
	buf := strings.Builder{}

	startButton := styles.PrimaryButtonStyle.MarginBackground(styles.Card).MarginRight(2).Render("Start")
	quitButton := styles.ButtonStyle.MarginBackground(styles.Card).Render("Quit")

	buttons := lipgloss.JoinHorizontal(lipgloss.Top, startButton, quitButton)
	ui := lipgloss.JoinVertical(lipgloss.Center, "rpg.ssh", buttons)

	dialog := lipgloss.Place(
		s.parent.Width,
		s.parent.Height-1,
		lipgloss.Center,
		lipgloss.Center,

		styles.DialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceBackground(styles.Bg),
		lipgloss.WithWhitespaceForeground(styles.Card),
		lipgloss.WithWhitespaceChars("#"),
	)
	buf.WriteString(dialog + "\n")

	return buf.String()
}
