package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	windowSize          = tea.WindowSizeMsg{Width: 0, Height: 0}
	docStyle            = lipgloss.NewStyle()
	windowStyle = lipgloss.NewStyle().
                Align(lipgloss.Left).
                Padding(0, 1, 0, 1).
                Border(lipgloss.RoundedBorder())
	boardWindowStyle = windowStyle.
				Width((windowSize.Width / 8) - 2).
				Height((windowSize.Height) - 3)
	threadWindowStyle = windowStyle.
				Width((windowSize.Width * 2 / 8) - 2).
				Height((windowSize.Height) - 3)
	contentWindowStyle = windowStyle.
				Width((windowSize.Width * 5 / 8) - 1).
				Height((windowSize.Height) - 3)
)

type model struct {
	Windows       []string
	WindowContent []string
	activeWindow  int
}
