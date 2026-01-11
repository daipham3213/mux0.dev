package portfolio

import "testing"

func TestDefaultPortfolio(t *testing.T) {
	p := Default()
	if p.Name == "" {
		t.Fatal("expected name to be set")
	}
	if len(p.Sections) == 0 {
		t.Fatal("expected at least one section")
	}

	seen := make(map[string]struct{}, len(p.Sections))
	for _, s := range p.Sections {
		if s.ID == "" {
			t.Fatal("expected section ID to be set")
		}
		if s.Title == "" {
			t.Fatal("expected section title to be set")
		}
		if s.Body == "" {
			t.Fatalf("expected section %q to have body content", s.ID)
		}
		if _, exists := seen[s.ID]; exists {
			t.Fatalf("duplicate section ID %q", s.ID)
		}
		seen[s.ID] = struct{}{}
	}
}
