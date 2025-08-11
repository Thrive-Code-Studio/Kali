package internal

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
	"fmt"
)

type model struct {
	txField textinput.Model
}

func KaliInitializer() model {
	tx := textinput.New()
	tx.Placeholder = "TCP"
	tx.Focus()
	tx.CharLimit = 256
	tx.Width = 20

	return model{ 
		txField: tx,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var command tea.Cmd

	switch message := message.(type) {
	case tea.KeyMsg:
		switch message.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	m.txField, command = m.txField.Update(message)
	return m, command
}


func (m model) View() string {
	return fmt.Sprintf("What Load Testing Type you want to utilize? \n %s \n %s",
			  m.txField.View(),
		          "(esc to quit)") + "\n"
}

