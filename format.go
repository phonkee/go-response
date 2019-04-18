package response

var (
	// currentKeyFormat holds current key format
	currentKeyFormat = SnakeCaseFormat
)

// Format sets format to given format
func Format(f format) {
	currentKeyFormat = f
}

type format struct {
	ErrorKey      string
	MessageKey    string
	ResultKey     string
	ResultSizeKey string
	StatusKey     string
}

var (
	// CamelCaseFormat sets common keys as camel case
	CamelCaseFormat = format{
		ErrorKey:      "Error",
		MessageKey:    "Message",
		ResultKey:     "Result",
		ResultSizeKey: "ResultSize",
		StatusKey:     "Status",
	}

	// SnakeCaseFormat sets common keys as snake case
	SnakeCaseFormat = format{
		ErrorKey:      "error",
		MessageKey:    "message",
		ResultKey:     "result",
		ResultSizeKey: "result_size",
		StatusKey:     "status",
	}
)
