package view

import (
    "github.com/charmbracelet/lipgloss"
    "github.com/shin_n_n/terleaf/style"
)

func SidePanelContent(currentPage string, width int) string {
    var buttonHomePage, buttonOptionPage, buttonHelpPage string
    if currentPage == "home" {
        buttonHomePage = style.SidePanelButtonActiveStyle.Width(width - 2).Render("Home (h)")
    } else {
        buttonHomePage = style.SidePanelButtonStyle.Width(width - 2).Render("Home (h)")
    }
    if currentPage == "option" {
        buttonOptionPage = style.SidePanelButtonActiveStyle.Width(width - 2).Render("Option (ctrl+o)")
    } else {
        buttonOptionPage = style.SidePanelButtonStyle.Width(width - 2).Render("Option (ctrl+o)")
    }
    if currentPage == "help" {
        buttonHelpPage = style.SidePanelButtonActiveStyle.Width(width - 2).Render("Help (ctrl+h)")
    } else {
        buttonHelpPage = style.SidePanelButtonStyle.Width(width - 2).Render("Help (ctrl+h)")
    }

    return lipgloss.JoinVertical(
        lipgloss.Left,
        buttonHomePage,
        buttonOptionPage,
        buttonHelpPage,
    )
}