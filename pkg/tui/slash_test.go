package tui

import (
	"io"
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"

	"github.com/daipham3213/mux0.dev/pkg/tui/theme"
)

func TestLogoViewIncludesText(t *testing.T) {
	m := newSplashTestModel()
	got := stripANSIForSplash(m.LogoView())
	if !strings.Contains(got, "mux0.dev") {
		t.Fatalf("expected logo to include mux0.dev")
	}
}

func TestSlashViewIncludesLogoAndLoading(t *testing.T) {
	m := newSplashTestModel()
	m.viewportWidth = 0
	m.viewportHeight = 0

	got := stripANSIForSplash(m.SlashView())
	if !strings.Contains(got, "mux0.dev") {
		t.Fatalf("expected slash view to include logo text")
	}
}

func newSplashTestModel() model {
	renderer := lipgloss.NewRenderer(io.Discard)
	renderer.SetColorProfile(termenv.TrueColor)
	th := theme.BasicTheme(renderer, nil)
	return model{
		renderer:    renderer,
		theme:       th,
		showSplash:  true,
		splashFrame: 0,
	}
}

func stripANSIForSplash(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	inEscape := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if inEscape {
			if c >= '@' && c <= '~' {
				inEscape = false
			}
			continue
		}
		if c == 0x1b && i+1 < len(s) && s[i+1] == '[' {
			inEscape = true
			i++
			continue
		}
		b.WriteByte(c)
	}
	return b.String()
}
