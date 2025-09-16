package mailwrap

import (
	"os"
	"path/filepath"
	"strings"
)

var resourcePath string = ""

func RegisterResourcePath(rp string) {
	resourcePath = rp
}

func GetEmailResource(filename string, fullPath bool, parse map[string]string) (string, error) {
	var currentPath string
	if fullPath {
		currentPath = filename
	} else {
		currentPath = filepath.Join(resourcePath, filename)
	}

	contentBytes, err := os.ReadFile(currentPath)
	if err != nil {
		return "", err
	}

	content := string(contentBytes)

	for key, value := range parse {
		placeholder := "{{" + key + "}}"
		content = strings.ReplaceAll(content, placeholder, value)
	}

	return content, nil
}
