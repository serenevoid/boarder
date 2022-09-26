package tui

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			return m, tea.Quit
        case "h", "left":
            return m, nil
        case "l", "right":
            return m, nil
		}
	case tea.WindowSizeMsg:
		windowSize = tea.WindowSizeMsg{Width: msg.Width, Height: msg.Height}
	}

	return m, nil
}

func (m model) View() string {
	doc := strings.Builder{}

	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		windowStyle.
			Width(windowSize.Width/20).
			Height((windowSize.Height)-2).
			Render(m.WindowContent[0]),
		windowStyle.
			Width(windowSize.Width*2/14).
			Height((windowSize.Height)-2).
			Render(m.WindowContent[1]),
		windowStyle.
			Width((windowSize.Width*3/4)).
			Height(windowSize.Height-2).
			Render(m.WindowContent[2]),
	)

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
