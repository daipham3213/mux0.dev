package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) ResizeView() string {
	return lipgloss.Place(
		m.viewportWidth,
		m.viewportHeight,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.theme.TextAccent().Render("window"),
			m.theme.TextBrand().Bold(true).Render("is too small"),
			m.theme.TextAccent().Render("resize to continue"),
		),
	)
}
