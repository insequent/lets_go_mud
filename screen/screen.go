package main

import (
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

const ()

type keymap struct {
	quit key.Binding
}

type model struct {
	// Settings
	height int
	width  int
	keymap keymap

	// Models
	help      help.Model      // Help menu (TODO)
	textinput textinput.Model // Primary input screen
	viewport  viewport.Model  // Primary output screen
}

func NewModel() model {
	m := model{
		keymap: keymap{
			quit: key.NewBinding(
				key.WithKeys("#quit", "ctrl+c"),
				key.WithHelp("#quit", "quit lgm"),
			),
		},

		help:      help.New(),
		textinput: textinput.New(),
		viewport:  viewport.New(),
	}

	m.textinput.Focus()

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//var cmds []tea.Cmd

	switch t := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(t, m.keymap.quit):
			os.Exit(1)
		}
	}

	return nil, nil
}
