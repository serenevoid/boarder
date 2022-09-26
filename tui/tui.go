package tui

import (
	"fmt"
	"os"
	"strings"

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

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		windowSize = tea.WindowSizeMsg{Width: msg.Width, Height: msg.Height}
	}

	return m, nil
}

func (m model) View() string {
	doc := strings.Builder{}
	sidePanels := lipgloss.JoinVertical(
		lipgloss.Top,
		windowStyle.
			Width(windowSize.Width/4).
			Height(1).
            Align(lipgloss.Center).
			Render("Boarder V0.0.1"),
		windowStyle.
			Width(windowSize.Width/4).
			Height((windowSize.Height/2)-4).
			Render(m.WindowContent[0]),
		windowStyle.
			Width(windowSize.Width/4).
			Height((windowSize.Height/2)-3).
			Render(m.WindowContent[1]),
	)
	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		sidePanels,
		windowStyle.
			Width((windowSize.Width*3/4)-3).
			Height(windowSize.Height-2).
			Render(m.WindowContent[2]),
	)
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row)
	doc.WriteString(row)
	return docStyle.Render(doc.String())
}

func Setup_UI() {
	windows := []string{"Board", "Thread", "Content"}
	windowContent := []string{"Board\n/w/\n/wg/\n/p/", "Thread\n283948\n289437\n289372", "Content\nHello there\nThis is the content of this thread"}
	m := model{Windows: windows, WindowContent: windowContent}
	if err := tea.NewProgram(m, tea.WithAltScreen()).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
