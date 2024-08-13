package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	textInput textinput.Model
	err       error
}

func New() model {
	ti := textinput.New()
	ti.Placeholder = "write your thoughts down here"
	ti.Focus()

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// check terminal messages
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			f, err := tea.LogToFile("log.txt", "debug")

			if err != nil {
				fmt.Println("fatal:", err)
				os.Exit(1)
			}

			defer f.Close()

			fmt.Fprint(f, m.textInput.Value())

			// save jot to sql

			return m, tea.Quit
		}

	case error:
		m.err = msg
		return m, nil
	}

	// assume user input
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprint(
		"write down just enough info \n",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
