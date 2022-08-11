package confirm

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Settings struct {
	Question          string
	OkButtonLabel     string
	CancelButtonLabel string
}

func (cs *Settings) initialModel() model {
	_okButtonLabel := ""
	if cs.OkButtonLabel != "" {
		_okButtonLabel = cs.OkButtonLabel
	} else {
		_okButtonLabel = "Yes"
	}

	_cancelButtonLabel := ""
	if cs.CancelButtonLabel != "" {
		_cancelButtonLabel = cs.CancelButtonLabel
	} else {
		_cancelButtonLabel = "No"
	}

	return model{
		question:          cs.Question,
		okButtonLabel:     _okButtonLabel,
		cancelButtonLabel: _cancelButtonLabel,
		confirmation:      true,

		questionStyle: questionStyle,
		buttonStyle:   defaultButtonStyle,
		selectedStyle: activeButtonStyle,
	}
}

// Run provides a shell script interface for prompting a user to confirm an
// action with an affirmative or negative answer.
func Run(cs *Settings) (bool, error) {
	m, err := tea.NewProgram(cs.initialModel(), tea.WithOutput(os.Stderr)).StartReturningModel()

	if err != nil {
		return false, fmt.Errorf("unable to run confirm: %w", err)
	}

	if m.(model).confirmation {
		return true, nil
	}

	return false, nil
}
