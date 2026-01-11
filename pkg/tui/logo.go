package tui

func (m model) LogoView() string {
	return m.theme.TextAccent().Bold(true).Render("mux0.dev") + m.CursorView()
}
