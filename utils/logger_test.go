package utils

import (
	"testing"

	"github.com/matryer/is"
)

func TestLogger(t *testing.T) {
	is := is.New(t)

	tc1 := "\ntest case 1:\n------------\n"

	printer := PrinterContent{
		Title: "test case 1",
	}
	printer.SetContent("")
	is.Equal(tc1, PrettyPrinter(&printer).ToString())

	logger := NewLoggerWriter()
	logger.Reset()
	logger.AppendItem("first")
	logger.AppendItem("second")

	tc2 := "\ntest case 2:\n------------\n* first\n* second\n"

	printer = PrinterContent{
		Title: "test case 2",
	}
	printer.SetContent(logger.Render())
	is.Equal(tc2, PrettyPrinter(&printer).ToString())

	logger.Reset()
	logger.AppendItem("first")
	logger.AppendItem("second")
	logger.Indent()
	logger.AppendItem("subItem")
	logger.UnIndent()

	printer = PrinterContent{
		Title: "test case 3",
	}
	printer.SetContent(logger.Render())

	tc4 := "\ntest case 3:\n------------\n* first\n* second\n  * subItem\n"
	is.Equal(tc4, PrettyPrinter(&printer).ToString())
}
