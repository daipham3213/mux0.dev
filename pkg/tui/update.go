package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type splashDoneMsg struct{}
type splashTickMsg struct{}

const splashDuration = 1 * time.Second
const splashFrameInterval = 350 * time.Millisecond

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.Tick(splashDuration, func(time.Time) tea.Msg {
			return splashDoneMsg{}
		}),
		tea.Tick(splashFrameInterval, func(time.Time) tea.Msg {
			return splashTickMsg{}
		}),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case splashDoneMsg:
		m.showSplash = false
		return m, nil
	case splashTickMsg:
		if !m.showSplash {
			return m, nil
		}
		m.splashFrame = (m.splashFrame + 1) & 1
		return m, tea.Tick(splashFrameInterval, func(time.Time) tea.Msg {
			return splashTickMsg{}
		})
	case tea.WindowSizeMsg:
		m.viewportWidth = msg.Width
		m.viewportHeight = msg.Height
		m.applyLayout()
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			m.focusNav = !m.focusNav
			return m, nil
		}
	}

	if m.focusNav {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		m.syncSelection()
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "g":
			m.viewport.GotoTop()
			return m, nil
		case "G":
			m.viewport.GotoBottom()
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}
