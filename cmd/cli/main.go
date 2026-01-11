package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/daipham3213/mux0.dev/pkg/tui"
)

func main() {
	m, err := tui.NewModel(lipgloss.DefaultRenderer(), "local", false, nil, []string{})
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	}
}
