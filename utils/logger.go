package utils

import (
	"os"
	"runtime"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type icon string

const (
	iconDebug icon = "\u25B8" // ▸
	iconFatal icon = "\u2718" // ✘
	iconError icon = "\u2718" // ✘
	iconWarn  icon = "\u0021" // esclamation mark
	iconInfo  icon = "\u25B8" // ▸
)

var levelIconMap = map[log.Level]icon{
	log.DebugLevel: iconDebug,
	log.FatalLevel: iconFatal,
	log.ErrorLevel: iconError,
	log.WarnLevel:  iconWarn,
	log.InfoLevel:  iconInfo,
}

var (
	// Light: black, Dark: white
	nocolor = lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FFFFFF"}
	// Light: gray-900, Dark: gray-300
	gray = lipgloss.AdaptiveColor{Light: "#0f172a", Dark: "#d1d5db"}
	// Light: red-500, Dark: red-500
	red = lipgloss.AdaptiveColor{Light: "#ef4444", Dark: "#ef4444"}
	// Light: orange-500, Dark: orange-400
	orange = lipgloss.AdaptiveColor{Light: "#f97316", Dark: "#fb923c"}
	// Light: blue-700, Dark: blue-500
	blue = lipgloss.AdaptiveColor{Light: "#1d4ed8", Dark: "#3b82f6"}
)

var levelColorMap = map[log.Level]lipgloss.AdaptiveColor{
	log.DebugLevel: gray,
	log.FatalLevel: red,
	log.ErrorLevel: red,
	log.WarnLevel:  orange,
	log.InfoLevel:  blue,
}

// SetupLogger creates a new logger, set styles for log levels and returns a pointer to the newly created logger.
func SetupLogger() *log.Logger {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: false,
		ReportCaller:    false,
	})

	log.DebugLevelStyle = makeLevelStyle(log.DebugLevel)
	log.WarnLevelStyle = makeLevelStyle(log.WarnLevel)
	log.ErrorLevelStyle = makeLevelStyle(log.ErrorLevel)
	log.FatalLevelStyle = makeLevelStyle(log.FatalLevel)
	log.InfoLevelStyle = makeLevelStyle(log.InfoLevel)

	return logger
}

func makeLevelStyle(level log.Level) lipgloss.Style {
	return lipgloss.NewStyle().SetString(getLevelIcon(level)).Foreground(getLevelColor(level))
}

func getLevelIcon(level log.Level) string {
	if isWindows() {
		return ">"
	}
	if _, ok := levelIconMap[level]; ok {
		return string(levelIconMap[level])
	}
	return "undefined"
}

// isWindows returns true if the OS is MS Windows.
func isWindows() bool {
	return runtime.GOOS == "windows"
}

func getLevelColor(level log.Level) lipgloss.AdaptiveColor {
	if _, ok := levelColorMap[level]; ok {
		return levelColorMap[level]
	}
	return nocolor
}
