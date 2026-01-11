package tui

import "github.com/charmbracelet/lipgloss"

func (m model) SlashView() string {
	logo := m.LogoView()
	splash := lipgloss.JoinVertical(lipgloss.Center, logo)
	if m.viewportWidth == 0 || m.viewportHeight == 0 {
		return splash
	}

	return m.renderer.Place(
		m.viewportWidth,
		m.viewportHeight,
		lipgloss.Center,
		lipgloss.Center,
		splash,
	)
}
