package view

import "github.com/charmbracelet/lipgloss"

func HomePageView(width, height int) string {
	content := "🏠 Home Page\n\nWillkommen auf der Startseite!\nHier kannst du dein Projekt vorstellen oder Infos anzeigen."
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Padding(1, 2).
		Render(content)
}
