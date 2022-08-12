package tui

import (
	"reflect"

	"github.com/charmbracelet/lipgloss"
)

// IsEmpty returns true if lipgloss.Style struct is empty
func IsEmpty(s lipgloss.Style) bool {
	return reflect.DeepEqual(s, lipgloss.Style{})
}
