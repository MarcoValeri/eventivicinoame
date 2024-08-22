package util

import (
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

func FormSanitizeStringInput(getStringInput string) string {
	/**
	* Avoid HTML injection
	* Remove space at the left and right position
	 */
	sanitizeHtml := bluemonday.StrictPolicy()
	outputInput := sanitizeHtml.Sanitize(getStringInput)

	outputInput = strings.TrimSpace(outputInput)
	return outputInput
}
