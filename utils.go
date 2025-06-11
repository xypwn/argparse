package argparse

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func decideTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 80
	}
	return width
}

func formatHelpRow(head, content string, bareHeadLength, maxHeadLength, terminalWidth int, withBreak bool) string {
	content = strings.Replace(content, "\n", "", -1)
	result := fmt.Sprintf("  %s ", head)
	headLeftPadding := maxHeadLength - bareHeadLength - 3
	if headLeftPadding > 0 { // fill left padding
		result += strings.Repeat(" ", headLeftPadding)
	}
	contentPadding := strings.Repeat(" ", maxHeadLength)
	var rows []string
	if withBreak && headLeftPadding < 0 {
		rows = append(rows, result, contentPadding+content)
	} else {
		rows = append(rows, result+content)
	}
	for len(rows[len(rows)-1]) > terminalWidth { // break into lines
		lastIndex := len(rows) - 1
		lastOne := rows[lastIndex]
		rows[lastIndex] = rows[lastIndex][0:terminalWidth]
		rows = append(rows, contentPadding+lastOne[terminalWidth:])
	}
	return strings.Join(rows, "\n")
}
