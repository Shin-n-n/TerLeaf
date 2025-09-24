package view

import "github.com/charmbracelet/lipgloss"

func HelpPageView(width, height int) string {
	content := "❓ Help Page\n\nHier findest du Hilfe und Anleitungen zur Nutzung der Anwendung."
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Padding(1, 2).
		Render(content)
}
