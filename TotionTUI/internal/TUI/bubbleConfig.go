package bubble

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

type model struct {
	CreateFileInputVisible bool
	textInputs             []textinput.Model
	focusIndex             int
}

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7c12e0"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7c12e0"))
)

func Initialization() model {
	m := model{
		textInputs:             make([]textinput.Model, 2),
		CreateFileInputVisible: false,
	}

	// Text Input Setup
	var t textinput.Model
	for i := range m.textInputs {
		t = textinput.New()
		t.CharLimit = 0
		t.SetWidth(156)

		s := t.Styles()
		s.Cursor.Color = lipgloss.Color("#7c12e0")
		s.Focused.Prompt = focusedStyle
		s.Focused.Text = focusedStyle
		s.Blurred.Prompt = blurredStyle
		s.Focused.Text = focusedStyle
		t.SetStyles(s)

		// Input Styles Setup
		switch i {
		case 0:
			t.Placeholder = "What would you like to name it?"
			t.Focus()
		case 1:
			t.Placeholder = "Write the Content Here........"
			t.CharLimit = 0
		}
		m.textInputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "esc":
			return m, tea.Quit

		// These keys should add a new note.
		case "ctrl+n":
			m.CreateFileInputVisible = true
			return m, nil

		case "ctrl+s", "enter":

			return m, nil

		case "ctrl+b":
			m.CreateFileInputVisible = false
			return m, nil

		case "up", "down":
			s := msg.String()
			// Cycle indexes
			if s == "up" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex >= len(m.textInputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.textInputs) - 1
			}

			cmds := make([]tea.Cmd, len(m.textInputs))
			for i := 0; i <= len(m.textInputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.textInputs[i].Focus()
					continue
				}
				// Remove focused state
				m.textInputs[i].Blur()
			}
			return m, tea.Batch(cmds...)
		}
	}

	cmds := make([]tea.Cmd, len(m.textInputs))
	if m.CreateFileInputVisible {

		// Only text inputs with Focus() set will respond, so it's safe to simply
		// update all of them here without any further logic.
		for i := range m.textInputs {
			m.textInputs[i], cmds[i] = m.textInputs[i].Update(msg)
		}
	}
	return m, tea.Batch(cmds...)
}

func (m model) View() tea.View {

	// Styling using the lipglose
	var style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#7c12e0")).
		PaddingLeft(2).PaddingBottom(0).PaddingTop(0).
		PaddingRight(2).Align(lipgloss.Center)

	// The header
	s := "Welcome to TUI based Notes App!"

	// rendering the style
	s = style.Render(s)

	// Notes Part
	notesView := ""
	var b strings.Builder
	if m.CreateFileInputVisible {
		b.WriteString("\nCreating New Note....\n\n")

		for i := range m.textInputs {
			b.WriteString(m.textInputs[i].View() + "\n")
		}
	}

	notesView = b.String()

	// The footer
	style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#1df916")).
		Padding(0, 1, 0).
		Border(lipgloss.NormalBorder())

	help := "Controls:  "
	help += "New-File:Ctrl + N, List:Ctrl+L, Back/Save:ESC, Save:Ctrl+S, Quit:Ctrl+Q."
	// rendering the style
	help = style.Render(help)

	// Combined String
	finalString := fmt.Sprintf("%s\n%s\n%s\n", s, notesView, help)

	// Send the UI for rendering
	return tea.NewView(finalString)
}
