package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	windowSize          = tea.WindowSizeMsg{Width: 0, Height: 0}
	docStyle            = lipgloss.NewStyle()
	inactiveWindowStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("62"))
	activeWindowStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("122"))
	windowStyle = lipgloss.NewStyle().
			Align(lipgloss.Left).
			Padding(0, 1, 0, 1).
			Border(lipgloss.RoundedBorder())
)

type model struct {
	Windows       []string
	WindowContent []string
	activeWindow  int
}
