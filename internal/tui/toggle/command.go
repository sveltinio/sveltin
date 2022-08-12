package toggle

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sveltinio/sveltin/internal/tui"
)

const (
	yesLabel    = "Yes"
	noLabel     = "No"
	cursorLabel = ">"
)

// Config represents the struct to configure the tui command.
type Config struct {
	Question            string
	OkButtonLabel       string
	CancelButtonLabel   string
	Cursor              string
	QuestionStyle       lipgloss.Style
	ButtonStyle         lipgloss.Style
	SelectedButtonStyle lipgloss.Style
}

// setDefaults sets default values for Config if not present.
func (cfg *Config) setDefaults() *Config {
	if cfg.OkButtonLabel == "" {
		cfg.OkButtonLabel = yesLabel
	}
	if cfg.CancelButtonLabel == "" {
		cfg.CancelButtonLabel = noLabel
	}
	if cfg.Cursor == "" {
		cfg.Cursor = cursorLabel
	}
	if tui.IsEmpty(cfg.QuestionStyle) {
		cfg.QuestionStyle = questionStyle
	}
	if tui.IsEmpty(cfg.ButtonStyle) {
		cfg.ButtonStyle = defaultButtonStyle
	}
	if tui.IsEmpty(cfg.SelectedButtonStyle) {
		cfg.SelectedButtonStyle = activeButtonStyle
	}
	return cfg
}

func (cfg *Config) initialModel() model {
	return model{
		question:          cfg.Question,
		okButtonLabel:     cfg.OkButtonLabel,
		cancelButtonLabel: cfg.CancelButtonLabel,
		confirmation:      true,

		cursor:        cfg.Cursor,
		questionStyle: cfg.QuestionStyle,
		buttonStyle:   cfg.ButtonStyle,
		selectedStyle: cfg.SelectedButtonStyle,
	}
}

// Run provides a shell script interface for prompting a user to confirm an
// action with an affirmative or negative answer.
func Run(cfg *Config) (bool, error) {
	defaultConfig := cfg.setDefaults()
	m, err := tea.NewProgram(defaultConfig.initialModel(), tea.WithOutput(os.Stderr)).StartReturningModel()

	if err != nil {
		return false, fmt.Errorf("unable to run confirm: %w", err)
	}

	if m.(model).confirmation {
		return true, nil
	}

	return false, nil
}
