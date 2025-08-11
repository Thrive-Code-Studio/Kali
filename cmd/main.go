package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"kali/internal"
	"log"
)


func main() {
	init := internal.KaliInitializer()
	p := tea.NewProgram(init)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
