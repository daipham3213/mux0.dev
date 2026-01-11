package tui

import (
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestNewModel(t *testing.T) {
	teaModel, err := NewModel(lipgloss.DefaultRenderer(), "test", false, nil, []string{})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	m, ok := teaModel.(model)
	if !ok {
		t.Fatal("expected model type assertion to succeed")
	}
	if m.activeID == "" {
		t.Fatal("expected active section to be set")
	}
	if m.list.Title == "" {
		t.Fatal("expected list title to be set")
	}
	if !m.focusNav {
		t.Fatal("expected navigation to be focused by default")
	}
	if m.sections[m.activeID] == "" {
		t.Fatalf("expected content for section %q", m.activeID)
	}
}
