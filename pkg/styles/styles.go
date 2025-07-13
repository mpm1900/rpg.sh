package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	BaseStyle = lipgloss.NewStyle().
			Foreground(Fg).
			Background(Bg).
			MarginBackground(Bg)

	ButtonStyle = BaseStyle.
			Foreground(SecondaryFg).
			Background(Secondary).
			Padding(0, 3).
			MarginTop(1)

	PrimaryButtonStyle = ButtonStyle.
				Background(Primary).
				Foreground(PrimaryFg)

	DialogBoxStyle = BaseStyle.
			Background(Card).
			MarginBackground(Card).
			Foreground(CardFg).
			Border(lipgloss.RoundedBorder()).
			BorderBackground(Card).
			BorderForeground(Border).
			Padding(1, 2).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	StatusBarStyle = BaseStyle.
			Foreground(MutedFg).
			Background(Muted)

	StatusStyle = BaseStyle.
			Inherit(StatusBarStyle).
			Foreground(Fg).
			Background(MutedFg).
			Padding(0, 1).
			MarginRight(1).
			MarginBackground(Muted)
)
