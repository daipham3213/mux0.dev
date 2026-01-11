package tui

import (
	"context"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/daipham3213/mux0.dev/pkg/portfolio"
	"github.com/daipham3213/mux0.dev/pkg/tui/theme"
)

type size int

const (
	undersized size = iota
	small
	medium
	large
)

type sectionItem struct {
	id    string
	title string
	desc  string
}

func (i sectionItem) Title() string       { return i.title }
func (i sectionItem) Description() string { return i.desc }
func (i sectionItem) FilterValue() string { return i.title }

type model struct {
	renderer        *lipgloss.Renderer
	context         context.Context
	command         []string
	fingerprint     string
	anonymous       bool
	list            list.Model
	viewport        viewport.Model
	sections        map[string]string
	sectionIDs      []string
	activeID        string
	focusNav        bool
	viewportWidth   int
	viewportHeight  int
	widthContainer  int
	heightContainer int
	widthContent    int
	heightContent   int
	size            size
	theme           theme.Theme
	showSplash      bool
	splashFrame     int
	navStyle        lipgloss.Style
	headStyle       lipgloss.Style
	bodyStyle       lipgloss.Style
	footStyle       lipgloss.Style
}

func NewModel(
	renderer *lipgloss.Renderer,
	fingerprint string,
	anonymous bool,
	clientIP *string,
	command []string,
) (tea.Model, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "client_ip", clientIP)

	p := portfolio.Default()
	th := theme.BasicTheme(renderer, nil)

	items := make([]list.Item, 0, len(p.Sections))
	sections := make(map[string]string, len(p.Sections))
	ids := make([]string, 0, len(p.Sections))
	for _, s := range p.Sections {
		items = append(items, sectionItem{id: s.ID, title: s.Title})
		sections[s.ID] = s.Body
		ids = append(ids, s.ID)
	}

	delegate := list.NewDefaultDelegate()
	delegate.Styles.NormalTitle = th.TextBody().Padding(0, 0, 0, 2)
	delegate.Styles.NormalDesc = th.TextBody().Faint(true).Padding(0, 0, 0, 2)
	delegate.Styles.SelectedTitle = th.TextHighlight().
		Bold(true).
		Padding(0, 0, 0, 1).
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(th.Highlight())
	delegate.Styles.SelectedDesc = th.TextHighlight().Padding(0, 0, 0, 1)
	delegate.Styles.DimmedTitle = th.TextBody().Faint(true).Padding(0, 0, 0, 2)
	delegate.Styles.DimmedDesc = th.TextBody().Faint(true).Padding(0, 0, 0, 2)

	l := list.New(items, delegate, 0, 0)
	l.Title = p.Name
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.Styles.Title = th.TextBrand().Bold(true).Padding(0, 1)
	l.Styles.TitleBar = th.TextBrand().Bold(true).Padding(0, 1, 1, 1)

	activeID := ids[0]
	if len(command) > 0 {
		activeID = selectSection(ids, p.Sections, command[0], activeID)
	}
	l.Select(indexForID(ids, activeID))

	vp := viewport.New(0, 0)
	vp.SetContent(sections[activeID])
	vp.Style = th.TextBody()

	return model{
		renderer:    renderer,
		context:     ctx,
		command:     command,
		fingerprint: fingerprint,
		anonymous:   anonymous,
		list:        l,
		viewport:    vp,
		sections:    sections,
		sectionIDs:  ids,
		activeID:    activeID,
		focusNav:    true,
		theme:       th,
		showSplash:  true,
		navStyle:    th.Base().Padding(0, 1),
		headStyle:   th.TextBrand().Bold(true),
		bodyStyle:   th.TextBody(),
		footStyle:   th.TextBody().Faint(true),
	}, nil
}

func selectSection(ids []string, sections []portfolio.Section, command string, fallback string) string {
	needle := strings.ToLower(command)
	for _, s := range sections {
		if strings.HasPrefix(strings.ToLower(s.ID), needle) ||
			strings.HasPrefix(strings.ToLower(s.Title), needle) {
			return s.ID
		}
	}
	return fallback
}

func indexForID(ids []string, id string) int {
	for i, value := range ids {
		if value == id {
			return i
		}
	}
	return 0
}
