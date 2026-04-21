package bubble

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	"fmt"
)

type model struct {
	CreateFileInputVisible bool
	textInput              textinput.Model
}

func Initialization() model {
	ti := textinput.New()
	ti.Placeholder = "What would you like to name Note?"
	// ti.SetVirtualCursor(false)
	ti.Focus()
	ti.CharLimit = 156
	ti.SetWidth(ti.CharLimit)

	return model{
		textInput:              ti,
		CreateFileInputVisible: false,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd 
	
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// These keys should add a new note.
		case "ctrl+n", "n":
			m.CreateFileInputVisible = true
			return m, nil

		}
	}

	if m.CreateFileInputVisible {
		m.textInput, cmd = m.textInput.Update(msg)
	}
	return m, cmd
}

func (m model) View() tea.View {

	// Styling using the lipglose
	var style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color("#7c12e0")).
		PaddingLeft(2).PaddingBottom(0).PaddingTop(0).
		PaddingRight(2).Align(lipgloss.Center)

	// The header
	s := "Welcome to TUI based Notes App!\n\n"

	// rendering the style
	s = style.Render(s)

	// Notes Part
	notesView := ""
	if m.CreateFileInputVisible {
		notesView = m.textInput.View()
	}

	// The footer
	style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#f91616")).Padding(0, 1, 0).Border(lipgloss.ASCIIBorder())
	help := "\nControls:"
	help += "\nNew-File:Ctrl + N, List:Ctrl+L, Back/Save:ESC, Save:Ctrl+S, Quit:Ctrl+Q.\n"
	// rendering the style
	help = style.Render(help)

	// Combined String
	finalString := fmt.Sprintf("%s\n%s\n%s", s, notesView, help)

	// Send the UI for rendering
	return tea.NewView(finalString)
}
