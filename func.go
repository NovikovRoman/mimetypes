package mimetypes

import (
	"mime"
	"strings"
)

// ExtensionByType returns the extension of the file based on the content type
func ExtensionByType(contentType string) string {
	contentType = strings.ToLower(contentType)
	return mimes[contentType]
}

func te(contentType string) ([]string, error) {
	return mime.ExtensionsByType(contentType)
}
