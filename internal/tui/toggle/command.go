package toggle

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Settings struct {
	Question          string
	OkButtonLabel     string
	CancelButtonLabel string
	Cursor            string
}

func (ts *Settings) initialModel() model {
	_okButtonLabel := ""
	if ts.OkButtonLabel != "" {
		_okButtonLabel = ts.OkButtonLabel
	} else {
		_okButtonLabel = "Yes"
	}

	_cancelButtonLabel := ""
	if ts.CancelButtonLabel != "" {
		_cancelButtonLabel = ts.CancelButtonLabel
	} else {
		_cancelButtonLabel = "No"
	}

	_cursor := ""
	if ts.Cursor != "" {
		_cursor = ts.Cursor
	} else {
		_cursor = ">"
	}

	return model{
		question:          ts.Question,
		okButtonLabel:     _okButtonLabel,
		cancelButtonLabel: _cancelButtonLabel,
		confirmation:      true,

		cursor:        _cursor,
		questionStyle: questionStyle,
		buttonStyle:   defaultButtonStyle,
		selectedStyle: activeButtonStyle,
	}
}

// Run provides a shell script interface for prompting a user to confirm an
// action with an affirmative or negative answer.
func Run(ts *Settings) (bool, error) {
	m, err := tea.NewProgram(ts.initialModel(), tea.WithOutput(os.Stderr)).StartReturningModel()

	if err != nil {
		return false, fmt.Errorf("unable to run confirm: %w", err)
	}

	if m.(model).confirmation {
		return true, nil
	}

	return false, nil
}
