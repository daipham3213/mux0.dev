package tui

func (m model) CursorView() string {
	if m.splashFrame&1 == 0 {
		return m.theme.Base().Background(m.theme.Brand()).Render(" ")
	}
	return m.theme.Base().Render(" ")
}
