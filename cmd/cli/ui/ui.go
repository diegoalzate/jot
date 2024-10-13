package ui

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/diegoalzate/jot/internal/database"
	"github.com/diegoalzate/jot/internal/query"
	"github.com/google/uuid"
)

type model struct {
	database  database.Service
	textInput textinput.Model
	err       error
	context   context.Context
}

func New(db database.Service) model {
	ti := textinput.New()
	ti.Placeholder = "write your thoughts down here"
	ti.Focus()

	return model{
		textInput: ti,
		err:       nil,
		database:  db,
		context:   context.Background(),
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

			q := query.New(m.database.Conn)

			_, err = q.CreateTask(m.context, query.CreateTaskParams{
				ID:          uuid.New().String(),
				Name:        m.textInput.Value(),
				Description: sql.NullString{},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			})

			if err != nil {
				log.Printf("failed to save task: %v", err)
			}

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
