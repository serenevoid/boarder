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
            if m.activeWindow > 0 {
                m.activeWindow--
            } else {
                m.activeWindow = 0
            }
			return m, nil
		case "l", "right":
            if m.activeWindow < 2 {
                m.activeWindow++
            } else {
                m.activeWindow = 2
            }
			return m, nil
		}
	case tea.WindowSizeMsg:
		windowSize = tea.WindowSizeMsg{Width: msg.Width, Height: msg.Height}
	}

	return m, nil
}

func (m model) View() string {
	doc := strings.Builder{}

    colors := []string{"62", "62", "62"}
    colors[m.activeWindow] = "122"

	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		windowStyle.
			Width((windowSize.Width/8)-2).
			Height((windowSize.Height)-3).
            BorderForeground(lipgloss.Color(colors[0])).
			Render(m.WindowContent[0]),
		windowStyle.
			Width((windowSize.Width*2/8)-2).
			Height((windowSize.Height)-3).
            BorderForeground(lipgloss.Color(colors[1])).
			Render(m.WindowContent[1]),
		windowStyle.
			Width((windowSize.Width*5/8)-1).
			Height((windowSize.Height)-3).
            BorderForeground(lipgloss.Color(colors[2])).
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
