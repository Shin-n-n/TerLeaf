package model

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shin_n_n/terleaf/style"
	"github.com/shin_n_n/terleaf/view"
)

type Model struct {
	Viewport     viewport.Model
	CmdWidth     int
	CmdHeight    int
	CurrentPage  string
	ActivePanel  string // "leftSidePanel" or "rightMainFrame"
	ActiveButton int
}

func InitialModel() Model {
	vp := viewport.New(44, 24)
	return Model{
		Viewport:     vp,
		CurrentPage:  "home",
		ActivePanel:  "leftSidePanel",
		ActiveButton: 0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		minWidth := 44
		minHeight := 24
		m.CmdWidth = msg.Width
		if m.CmdWidth < minWidth {
			m.CmdWidth = minWidth
		}
		m.CmdHeight = msg.Height
		if m.CmdHeight < minHeight {
			m.CmdHeight = minHeight
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			if m.ActivePanel == "leftSidePanel" {
				m.ActivePanel = "rightMainFrame"
			} else {
				m.ActivePanel = "leftSidePanel"
			}
		case "ctrl+o":
			m.CurrentPage = "option"
			m.ActiveButton = 1
		case "ctrl+h":
			m.CurrentPage = "help"
			m.ActiveButton = 2
		case "h":
			m.CurrentPage = "home"
			m.ActiveButton = 0
		case "up":
			if m.ActivePanel == "leftSidePanel" {
				m.ActiveButton--
				if m.ActiveButton < 0 {
					m.ActiveButton = 2
				}
				switch m.ActiveButton {
				case 0:
					m.CurrentPage = "home"
				case 1:
					m.CurrentPage = "option"
				case 2:
					m.CurrentPage = "help"
				}
			}
		case "down":
			if m.ActivePanel == "leftSidePanel" {
				m.ActiveButton++
				if m.ActiveButton > 2 {
					m.ActiveButton = 0
				}
				switch m.ActiveButton {
				case 0:
					m.CurrentPage = "home"
				case 1:
					m.CurrentPage = "option"
				case 2:
					m.CurrentPage = "help"
				}
			}
		case "right":
			if m.ActivePanel == "leftSidePanel" {
				m.ActivePanel = "rightMainFrame"
			}
		case "enter":
			if m.ActivePanel == "leftSidePanel" {

			}
		case "ctrl+q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	leftWidth := int(float64(m.CmdWidth) * 0.15)
	if leftWidth < 20 {
		leftWidth = 20
	}
	rightWidth := m.CmdWidth - leftWidth - 4
	if rightWidth < 20 {
		rightWidth = 20
	}

	// Sidepanel Buttons
	var buttonHome, buttonOption, buttonHelp string
	buttonHome = style.SidePanelButtonStyle.Width(leftWidth - 2).Render("Home (h)")
	buttonOption = style.SidePanelButtonStyle.Width(leftWidth - 2).Render("Option (ctrl+o)")
	buttonHelp = style.SidePanelButtonStyle.Width(leftWidth - 2).Render("Help (ctrl+h)")

	//if m.ActivePanel == "leftSidePanel" {
	switch m.ActiveButton {
	case 0:
		buttonHome = style.SidePanelButtonActiveStyle.Width(leftWidth - 2).Render("Home (h)")
	case 1:
		buttonOption = style.SidePanelButtonActiveStyle.Width(leftWidth - 2).Render("Option (ctrl+o)")
	case 2:
		buttonHelp = style.SidePanelButtonActiveStyle.Width(leftWidth - 2).Render("Help (ctrl+h)")
	}
	//}

	leftSidePanelContent := lipgloss.JoinVertical(
		lipgloss.Left,
		buttonHome,
		buttonOption,
		buttonHelp,
	)

	// Page Content
	var rightContent string
	switch m.CurrentPage {
	case "home":
		rightContent = view.HomePageView(rightWidth-2, m.CmdHeight-4)
	case "option":
		rightContent = view.OptionPageView(rightWidth-2, m.CmdHeight-4)
	case "help":
		rightContent = view.HelpPageView(rightWidth-2, m.CmdHeight-4)
	default:
		rightContent = view.HomePageView(rightWidth-2, m.CmdHeight-4)
	}

	// Aktiver Rahmen
	var leftFrameStyle, rightFrameStyle lipgloss.Style
	if m.ActivePanel == "leftSidePanel" {
		leftFrameStyle = style.FrameActiveStyle
		rightFrameStyle = style.FrameStyle
	} else {
		leftFrameStyle = style.FrameStyle
		rightFrameStyle = style.FrameActiveStyle
	}

	leftFrame := leftFrameStyle.Width(leftWidth).
		Height(m.CmdHeight - 3).
		Render(leftSidePanelContent)

	rightFrame := rightFrameStyle.Width(rightWidth).
		Height(m.CmdHeight - 3).
		Render(rightContent)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftFrame,
		rightFrame,
	)
}
