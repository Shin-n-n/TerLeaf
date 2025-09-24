package style

import "github.com/charmbracelet/lipgloss"

var FrameStyle = lipgloss.NewStyle().
    Border(lipgloss.NormalBorder()).
    BorderForeground(lipgloss.Color("ffffff"))

var FrameActiveStyle = lipgloss.NewStyle().
    Border(lipgloss.ThickBorder()).
    BorderForeground(lipgloss.Color("205"))