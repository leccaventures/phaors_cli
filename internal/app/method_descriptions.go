package app

import (
	_ "embed"
	"strings"
	"sync"
)

//go:embed assets/ABI_REFERENCE.md
var abiReferenceMarkdown string

var (
	methodDescriptionsOnce sync.Once
	methodDescriptions     map[string]string
)

func loadMethodDescriptions() map[string]string {
	methodDescriptionsOnce.Do(func() {
		methodDescriptions = parseMethodDescriptions(abiReferenceMarkdown)
	})
	return methodDescriptions
}

func parseMethodDescriptions(markdown string) map[string]string {
	descriptions := make(map[string]string)
	lines := strings.Split(markdown, "\n")

	currentHeading := ""

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "### ") {
			currentHeading = strings.TrimSpace(strings.TrimPrefix(trimmed, "### "))
			continue
		}

		if currentHeading == "" {
			continue
		}

		if strings.HasPrefix(trimmed, "- **Description**:") {
			description := strings.TrimSpace(strings.TrimPrefix(trimmed, "- **Description**:"))
			if description != "" {
				descriptions[currentHeading] = description
			}
			currentHeading = ""
		}
	}

	return descriptions
}

func methodDescription(methodName string) string {
	return loadMethodDescriptions()[methodName]
}
