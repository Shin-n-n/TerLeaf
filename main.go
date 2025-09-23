package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	viewport  viewport.Model
	cmdWidth  int
	cmdHeight int
}

func initialModel() model {
	vp := viewport.New(10, 10) // Breite x Höhe des Viewports
	vp.SetContent(`
    Hello, World!
    This is a simple Bubble Tea application using the viewport component.
    You can scroll if the content exceeds the viewport size.
    Enjoy!

    Keys:
    - down: Scroll down
    - up: Scroll up
    - space: Go to bottom
    - b: Go to top
    - q/ctrl+c: Quit`)

	vp.Style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true).Padding(1, 2)
	return model{
		viewport: vp,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.cmdWidth = msg.Width
		m.cmdHeight = msg.Height
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q":
			return m, tea.Quit
		case "down":
			m.viewport.ScrollDown(1)
		case "up":
			m.viewport.ScrollUp(1)
		case "left":
			m.viewport.ScrollLeft(1)
		case "right":
			m.viewport.ScrollRight(1)
		case "ctrl+o":
			m.viewport.SetContent(
				"Option page\n\n" +
					"This is a placeholder for the options page.\n" +
					"Press 'esc' to go back. (not working rgt now)",
			)
		case "ctrl+h":
			m.viewport.SetContent(
				"Help page\n\n" +
					"This is a placeholder for the help page.\n" +
					"Keys:\n" +
					"- down: Scroll down\n" +
					"- up: Scroll up\n" +
					"- space: Go to bottom\n" +
					"- b: Go to top\n" +
					"- q/ctrl+c: Quit\n\n" +
					"- ctrl+o: Options page\n" +
					"- ctrl+h: Help page\n\n" +
					"Press 'esc' to go back. (not working rgt now)",
			)
		}
	}
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.viewport.View()
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
