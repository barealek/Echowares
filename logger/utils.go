package echologger

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func formatLog(s string, replace map[string]string, errorString string, shouldPad, shouldColor bool) string {

	if shouldColor {
		if replace[TagMethod] != "" {
			code := replace[TagMethod]
			s = strings.Replace(s, TagMethod, methodColor(code), -1)
		}
		if replace[TagStatus] != "" {
			code, _ := strconv.Atoi(replace[TagStatus])
			s = strings.Replace(s, TagStatus, statusColor(code), -1)
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
	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return color.GreenString(strconv.Itoa(code))
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return color.BlueString(strconv.Itoa(code))
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return color.YellowString(strconv.Itoa(code))
	default:
		return color.RedString(strconv.Itoa(code))
	}
}

func methodColor(method string) string {
	switch method {
	case http.MethodGet:
		return color.CyanString(method)
	case http.MethodPost:
		return color.GreenString(method)
	case http.MethodPut:
		return color.YellowString(method)
	case http.MethodDelete:
		return color.RedString(method)
	case http.MethodPatch:
		return color.MagentaString(method)
	case http.MethodHead:
		return color.WhiteString(method)
	case http.MethodOptions:
		return color.BlueString(method)
	default:
		return color.WhiteString(method)
	}
}

func pad(s string, tag string) string {
	switch tag {
	case TagLatency:
		return padLeft(s, spaceLatency)
	case TagTime:
		return padLeft(s, spaceTime)
	case TagPid:
		return padLeft(s, spacePid)
	case TagStatus:
		return padLeft(s, spaceStatus)
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
	if len(s) >= n {
		return s
	}
	return strings.Repeat(" ", n-len(s)) + s
}

func padRight(s string, n int) string {
	if len(s) > n {
		return s
	}
	return s + strings.Repeat(" ", n-len(s))
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
