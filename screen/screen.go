package screen

import (
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"golang.org/x/term"
)

var (
	borderStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder())
	textStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color(lipgloss.Color("#0099FF")))
)

type keymap struct {
	quit key.Binding
}

type model struct {
	// Settings
	keymap keymap

	// Terminal size (i.e. overall max
	height int
	width  int

	// Models
	help      help.Model      // Help menu (TODO)
	textinput textinput.Model // Primary input screen
	viewport  viewport.Model  // Primary output screen

	// TODO: End goal here is to have flexibility in layout. Allow the user
	// to define what goes where...
}

func NewModel() *model {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Print("Failed to find terminal dimensions! Error:", err)
		os.Exit(1)
	}

	m := model{
		keymap: keymap{
			quit: key.NewBinding(
				key.WithKeys("#quit", "ctrl+c"),
				key.WithHelp("#quit", "quit lgm"),
			),
		},

		height: height,
		width:  width,

		help:      help.New(),
		textinput: textinput.New(),
		viewport:  viewport.New(width, height),
	}

	// Viewport styling
	m.viewport.Style = borderStyle.Copy()
	text := "Waiting for connection status..."
	// Content Width: -2XBorder, Height: -(2XBorder + textinput height)
	content := textStyle.Copy().Width(width - 2).Height(height - 3).Render(text)
	m.viewport.SetContent(content)

	// Textinput styling
	m.textinput.Width = width
	m.textinput.Focus()

	return &m
}

func (m *model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			os.Exit(1)
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.viewport.Height = msg.Height - 3
		m.viewport.Width = msg.Width - 2
		m.textinput.Width = msg.Width // Height always 1
	}

	return m, nil
}

func (m *model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		m.viewport.View(),
		m.textinput.View(),
	)
}
