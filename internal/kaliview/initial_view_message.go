package kaliview


import (
	tea "github.com/charmbracelet/bubbletea"
)


type Model struct {
	err error
}

func NewKaliMainView() Model {
	return Model{}
}

func (m Model) View() string {
	return ""
}
