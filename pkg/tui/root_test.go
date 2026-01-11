package tui

import (
	"testing"

	"github.com/daipham3213/mux0.dev/pkg/portfolio"
)

func TestSelectSection(t *testing.T) {
	ids := []string{"about", "experience", "social"}
	sections := []portfolio.Section{
		{ID: "about", Title: "About", Body: "body"},
		{ID: "experience", Title: "Experience", Body: "body"},
		{ID: "social", Title: "Socials", Body: "body"},
	}
	fallback := "about"

	selected := selectSection(ids, sections, "exp", fallback)
	if selected != "experience" {
		t.Fatalf("expected experience, got %q", selected)
	}

	selected = selectSection(ids, sections, "soc", fallback)
	if selected != "social" {
		t.Fatalf("expected social, got %q", selected)
	}

	selected = selectSection(ids, sections, "unknown", fallback)
	if selected != fallback {
		t.Fatalf("expected fallback %q, got %q", fallback, selected)
	}
}

func TestIndexForID(t *testing.T) {
	ids := []string{"about", "experience", "social"}
	if index := indexForID(ids, "social"); index != 2 {
		t.Fatalf("expected index 2, got %d", index)
	}
	if index := indexForID(ids, "missing"); index != 0 {
		t.Fatalf("expected index 0 for missing id, got %d", index)
	}
}
