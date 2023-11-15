package echologger

var (
	TagLatency = "${latency}"
	TagTime    = "${time}"
	TagPid     = "${pid}"
	TagStatus  = "${status}"
	TagMethod  = "${method}"
	TagHost    = "${host}"
	TagPath    = "${path}"
	TagError   = "${error}"

	spaceLatency          = 10 // 2.202771ms = 10 chars (the longest duration in Go on the Linux operating system)
	spaceTime             = 8  // 15:04:05 = 8 chars
	spacePid              = 5  // 65535 = 5 chars
	spaceStatus           = 3  // 200, 404, 500 = 3 chars
	spaceMethod           = 6  // DELETE = 6 chars
	spaceHost             = 15 // 100.100.100.100 = 15 chars
	spacePath             = 12
	absoluteErrorPosition = 74 // The error will be placed at the 50th character, or at the end of the line if the format is longer than 50 characters
)
