package choose

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Settings struct {
	Title    string
	ErrorMsg string
}

func (cs *Settings) initialModel(items []list.Item) model {
	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = cs.Title
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	//l.Styles.HelpStyle = helpStyle
	l.SetShowHelp(false)

	return model{
		list: l,
		err:  errors.New(cs.ErrorMsg),
	}
}

// Run is used to prompt a list of available options to the user and retrieve the selection.
func Run(cs *Settings, items []list.Item) (string, error) {
	ti := cs.initialModel(items)
	p := tea.NewProgram(ti, tea.WithOutput(os.Stderr))

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
