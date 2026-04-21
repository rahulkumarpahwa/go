package bubble

import (
	// "fmt"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func Initialization() model {
	return model{
		// Our to-do list is a grocery list
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}

}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the space bar toggle the selected state
		// for the item that the cursor is pointing at.
		case "enter", "space":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
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

	// Iterate over our choices
	// for i, choice := range m.choices {

	// 	// Is the cursor pointing at this choice?
	// 	cursor := " " // no cursor
	// 	if m.cursor == i {
	// 		cursor = ">" // cursor!
	// 	}

	// 	// Is this choice selected?
	// 	checked := " " // not selected
	// 	if _, ok := m.selected[i]; ok {
	// 		checked = "x" // selected!
	// 	}

	// 	// Render the row
	// 	s += fmt.Sprintf("%s [%s] %d %s\n", cursor, checked, i+1, choice)
	// }

	s += "\n\n"

	style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#f91616")).Padding(0, 1, 0).Border(lipgloss.ASCIIBorder())
	// The footer
	help := "\nControls:"
	help += "\nNew-File:Ctrl + N, List:Ctrl+L, Back/Save:ESC, Save:Ctrl+S, Quit:Ctrl+Q.\n"

	help = style.Render(help)

	// adding help to string s
	s += help

	// Send the UI for rendering
	return tea.NewView(s)
}
