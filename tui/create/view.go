package create

import (
	"strings"

	"github.com/onioncall/dndgo/tui/shared"
)

func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}

	var content strings.Builder

	textLines := 3 // page header + empty line + page text
	topPadding := (m.height - textLines) / 2

	for range topPadding {
		content.WriteString("\n")
	}

	dndgoLeftPadding := max((m.width-len(shared.TuiHeader))/2, 0)
	content.WriteString(strings.Repeat(" ", dndgoLeftPadding))
	content.WriteString(shared.TuiHeader)
	content.WriteString("\n\n")

	leftPadding := max((m.width-len(m.pageText))/2, 0)
	content.WriteString(strings.Repeat(" ", leftPadding))
	content.WriteString(m.pageText)

	return content.String()
}
