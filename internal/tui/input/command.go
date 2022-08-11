package input

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Settings struct {
	Focus        bool
	Placeholder  string
	DefaultValue string
	ErrorMsg     string
	// styles
	TextStyle   lipgloss.Style
	PromptStyle lipgloss.Style
}

func (is *Settings) initialModel() model {
	ti := textinput.New()

	if is.Focus {
		ti.Focus()
	}
	ti.Placeholder = is.Placeholder
	ti.SetValue(is.DefaultValue)
	ti.TextStyle = textStyle
	ti.PromptStyle = promptStyle

	return model{
		textInput:    ti,
		placeholder:  is.Placeholder,
		defaultValue: is.DefaultValue,
		err:          errors.New(is.ErrorMsg),
		textStyle:    textStyle,
		promptStyle:  promptStyle,
	}
}

// Run is used to prompt an input the user and retrieve the value.
func Run(is *Settings) (string, error) {
	ti := is.initialModel()
	p := tea.NewProgram(ti, tea.WithOutput(os.Stderr))

	tm, err := p.StartReturningModel()
	m := tm.(model)

	if m.quitting {
		return "", fmt.Errorf(m.err.Error())
	}

	if m.textInput.Value() == "" {
		return m.defaultValue, err
	}

	return m.textInput.Value(), err
}
