package echologger

import (
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
)

func formatLog(s string, replace map[string]string, shouldPad, shouldColor bool) string {

	if shouldColor {
		if replace[TagMethod] != "" {
			s = strings.Replace(s, TagMethod, methodColor(replace[TagMethod]), -1)
		}
		if replace[TagStatus] != "" {
			code, _ := strconv.Atoi(replace[TagStatus])
			s = strings.Replace(s, TagStatus, statusColor(code), -1)
		}
		if replace[TagError] != "" {
			s = strings.Replace(s, TagError, color.RedString("| "+replace[TagError]), -1)
		}
	}

	if shouldPad {
		for k, v := range replace {
			s = strings.Replace(s, k, pad(v, k), -1)
		}
	} else {

		for k, v := range replace {
			s = strings.Replace(s, k, v, -1)
		}
	}

	return s
}

func statusColor(code int) string {
	var col color.Attribute
	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		col = color.FgHiGreen
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		col = color.FgBlue
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		col = color.FgYellow
	default:
		col = color.FgRed
	}

	return color.New(col).Sprintf(TagStatus)
}

func methodColor(method string) string {
	var col color.Attribute
	switch method {
	case http.MethodGet:
		col = color.FgCyan
	case http.MethodPost:
		col = color.FgGreen
	case http.MethodPut:
		col = color.FgYellow
	case http.MethodDelete:
		col = color.FgRed
	case http.MethodPatch:
		col = color.FgMagenta
	case http.MethodHead:
		col = color.FgWhite
	case http.MethodOptions:
		col = color.FgBlue
	default:
		col = color.FgWhite
	}

	return color.New(col).Sprint(TagMethod)
}

func pad(s string, tag string) string {
	switch tag {
	case TagLatency:
		return padLeft(s, spaceLatency)
	case TagTime:
		return padLeft(s, spaceTime)
	case TagPid:
		return padLeft(s, spacePid)
	case TagMethod:
		return padRight(s, spaceMethod)
	case TagHost:
		return padLeft(s, spaceHost)
	case TagPath:
		return padRight(s, spacePath)
	case TagError:
		// The error will be placed at the 50th character, or at the end of the line if the format is longer than 50 characters
		if len(s) > absoluteErrorPosition {
			return s[:absoluteErrorPosition] + s[absoluteErrorPosition:len(s)-1] + " "
		}
		return padRight(s, absoluteErrorPosition)

	default:
		return s
	}
}

func padLeft(s string, n int) string {
	strlen := utf8.RuneCountInString(s)
	if strlen >= n {
		return s
	}
	return strings.Repeat(" ", n-strlen) + s
}

func padRight(s string, n int) string {
	strlen := utf8.RuneCountInString(s)
	if strlen > n {
		return s
	}
	return s + strings.Repeat(" ", n-strlen)
}

func findTags(format string) []string {
	var tags []string

	for _, tag := range []string{TagLatency, TagTime, TagPid, TagStatus, TagMethod, TagHost, TagPath, TagError} {
		if strings.Contains(format, tag) {
			tags = append(tags, tag)
		}
	}

	return tags
}
