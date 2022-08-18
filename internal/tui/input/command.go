package input

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sveltinio/sveltin/internal/tui"
)

const (
	defaultPlaceholder = "Placeholder text..."
	defaultText        = "Default value..."
	defaultErrorMsg    = "Error Message..."
)

// Config represents the struct to configure the tui command.
type Config struct {
	Placeholder  string
	DefaultValue string
	ErrorMsg     string
	// styles
	Focus       *bool
	TextStyle   lipgloss.Style
	PromptStyle lipgloss.Style
}

// setDefaults sets default values for Config if not present.
func (cfg *Config) setDefaults() *Config {
	if cfg.Placeholder == "" {
		cfg.Placeholder = defaultPlaceholder
	}
	if cfg.DefaultValue == "" {
		cfg.DefaultValue = defaultText
	}
	if cfg.ErrorMsg == "" {
		cfg.ErrorMsg = defaultErrorMsg
	}
	if cfg.Focus == nil {
		_focus := true
		cfg.Focus = &_focus
	}
	if tui.IsEmpty(cfg.TextStyle) {
		cfg.TextStyle = textStyle
	}
	if tui.IsEmpty(cfg.PromptStyle) {
		cfg.PromptStyle = promptStyle
	}
	return cfg
}

func (cfg *Config) initialModel() model {
	ti := textinput.New()
	if *cfg.Focus {
		ti.Focus()
	}
	ti.Placeholder = cfg.Placeholder
	//ti.SetValue(cfg.DefaultValue)
	ti.TextStyle = cfg.TextStyle
	ti.PromptStyle = cfg.PromptStyle

	return model{
		textInput:    ti,
		placeholder:  cfg.Placeholder,
		defaultValue: cfg.DefaultValue,
		err:          errors.New(cfg.ErrorMsg),
		textStyle:    cfg.TextStyle,
		promptStyle:  cfg.PromptStyle,
	}
}

// Run is used to prompt an input the user and retrieve the value.
func Run(cfg *Config) (string, error) {
	defaultConfig := cfg.setDefaults()
	p := tea.NewProgram(defaultConfig.initialModel(), tea.WithOutput(os.Stderr))

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
