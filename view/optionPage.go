package view

import "github.com/charmbracelet/lipgloss"

func OptionPageView(width, height int) string {
	content := "⚙️ Option Page\n\nHier kannst du verschiedene Einstellungen vornehmen."
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Padding(1, 2).
		Render(content)
}
