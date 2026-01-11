package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	if m.showSplash {
		return m.SlashView()
	}
	if m.viewportWidth == 0 || m.viewportHeight == 0 {
		return "Loading..."
	}
	if m.size == undersized {
		return m.ResizeView()
	}

	nav := m.navStyle.Render(m.list.View())
	content := m.bodyStyle.Render(m.viewport.View())

	header := m.headStyle.Render(m.activeTitle())
	footer := m.footStyle.Render(m.footerText())
	contentBlock := lipgloss.JoinVertical(lipgloss.Left, header, content, footer)
	layout := lipgloss.JoinHorizontal(lipgloss.Top, nav, contentBlock)

	body := m.theme.Base().
		Width(m.widthContainer).
		Height(m.heightContainer).
		Render(layout)

	if m.widthContainer == m.viewportWidth && m.heightContainer == m.viewportHeight {
		return body
	}

	return m.renderer.Place(
		m.viewportWidth,
		m.viewportHeight,
		lipgloss.Center,
		lipgloss.Center,
		body,
	)
}

func (m *model) syncSelection() {
	item, ok := m.list.SelectedItem().(sectionItem)
	if !ok || item.id == "" {
		return
	}
	if item.id == m.activeID {
		return
	}
	m.activeID = item.id
	m.viewport.SetContent(m.sections[item.id])
}

func (m *model) applyLayout() {
	switch {
	case m.viewportWidth < 60 || m.viewportHeight < 10:
		m.size = undersized
		m.widthContainer = m.viewportWidth
		m.heightContainer = m.viewportHeight
	case m.viewportWidth < 90:
		m.size = small
		m.widthContainer = m.viewportWidth
		m.heightContainer = m.viewportHeight
	case m.viewportWidth < 120:
		m.size = medium
		m.widthContainer = 90
		m.heightContainer = min(m.viewportHeight, 32)
	default:
		m.size = large
		m.widthContainer = 100
		m.heightContainer = min(m.viewportHeight, 36)
	}

	navWidth := clamp(m.widthContainer/4, 22, 32)
	contentWidth := m.widthContainer - navWidth - 1
	if contentWidth < 20 {
		contentWidth = 20
	}

	headerHeight := 1
	footerHeight := 1
	contentHeight := m.heightContainer - headerHeight - footerHeight
	if contentHeight < 5 {
		contentHeight = 5
	}

	m.list.SetSize(navWidth, m.heightContainer)
	m.viewport.Width = contentWidth
	m.viewport.Height = contentHeight
}

func (m model) activeTitle() string {
	item, ok := m.list.SelectedItem().(sectionItem)
	if ok && item.title != "" {
		return item.title
	}
	return m.activeID
}

func (m model) footerText() string {
	focus := "content"
	if m.focusNav {
		focus = "nav"
	}
	return fmt.Sprintf("tab: switch focus (%s)  j/k: move  g/G: top/bottom  q: quit", focus)
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
