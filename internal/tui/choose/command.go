package choose

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sveltinio/sveltin/internal/tui"
)

const (
	listHeight   = 14
	defaultWidth = 50
)

// Config represents the struct to configure the tui command.
type Config struct {
	Title        string
	ErrorMsg     string
	ListHeight   int
	DefaultWidth int
	TitleStyle   lipgloss.Style
}

// setDefaults sets default values for Config if not present.
func (cfg *Config) setDefaults() *Config {
	if cfg.ListHeight == 0 {
		cfg.ListHeight = listHeight
	}
	if cfg.DefaultWidth == 0 {
		cfg.DefaultWidth = defaultWidth
	}
	if tui.IsEmpty(cfg.TitleStyle) {
		cfg.TitleStyle = titleStyle
	}
	return cfg
}

func (cfg *Config) initialModel(items []list.Item) model {
	l := list.New(items, itemDelegate{}, cfg.DefaultWidth, cfg.ListHeight)
	l.Title = cfg.Title
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = cfg.TitleStyle
	l.SetShowHelp(false)

	return model{
		list: l,
		err:  errors.New(cfg.ErrorMsg),
	}
}

// Run is used to prompt a list of available options to the user and retrieve the selection.
func Run(cfg *Config, items []list.Item) (string, error) {
	defaultConfig := cfg.setDefaults()
	p := tea.NewProgram(defaultConfig.initialModel(items), tea.WithOutput(os.Stderr))

	tm, err := p.StartReturningModel()
	if err != nil {
		return "", fmt.Errorf("aborted")
	}
	m := tm.(model)

	if m.quitting {
		return "", fmt.Errorf(m.err.Error())
	}
	return m.list.SelectedItem().FilterValue(), nil
}
