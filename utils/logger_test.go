package utils

import (
	"testing"

	"github.com/matryer/is"
)

func TestLogger(t *testing.T) {
	is := is.New(t)

	tc1 := "\ntest case 1:\n------------\n"

	textLogger := NewTextLogger()
	textLogger.SetTitle("test case 1")
	textLogger.SetContent("")
	is.Equal(tc1, PrettyPrinter(textLogger).ToString())

	logger := NewListWriter()
	logger.Reset()
	logger.AppendItem("first")
	logger.AppendItem("second")

	tc2 := "\ntest case 2:\n------------\n* first\n* second\n"

	textLogger.Reset()
	textLogger.SetTitle("test case 2")
	textLogger.SetContent(logger.Render())
	is.Equal(tc2, PrettyPrinter(textLogger).ToString())

	logger.Reset()
	logger.AppendItem("first")
	logger.AppendItem("second")
	logger.Indent()
	logger.AppendItem("subItem")
	logger.UnIndent()

	textLogger.Reset()
	textLogger.SetTitle("test case 3")
	textLogger.SetContent(logger.Render())

	tc4 := "\ntest case 3:\n------------\n* first\n* second\n  * subItem\n"
	is.Equal(tc4, PrettyPrinter(textLogger).ToString())
}
