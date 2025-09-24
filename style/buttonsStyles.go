package style

import "github.com/charmbracelet/lipgloss"

var SidePanelButtonStyle = lipgloss.NewStyle().
    Border(lipgloss.RoundedBorder(), true).
    Foreground(lipgloss.Color("fffff")).
    Background(lipgloss.Color("0")).
    AlignHorizontal(lipgloss.Center).
    Padding(0, 2).
    Margin(0, 0)

var SidePanelButtonActiveStyle = lipgloss.NewStyle().
    Border(lipgloss.RoundedBorder(), true).
    Foreground(lipgloss.Color("#000000")).
    Background(lipgloss.Color("#fffff")).
    AlignHorizontal(lipgloss.Center).
    Padding(0, 2).
    Margin(0, 0)