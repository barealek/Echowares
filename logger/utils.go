package echologger

import (
	"strings"
)

func replaceAll(s string, replace map[string]string) string {
	for k, v := range replace {
		s = strings.Replace(s, k, v, -1)
	}
	return s
}

func findTags(format string) []string {
	var tags []string

	for _, tag := range []string{TagLatency, TagTime, TagStatus, TagMethod, TagPath, TagHost} {
		if strings.Contains(format, tag) {
			tags = append(tags, tag)
		}
	}

	return tags
}
