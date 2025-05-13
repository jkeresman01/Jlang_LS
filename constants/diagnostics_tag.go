package constants

type DiagnosticTag int32

const (
	DIAGNOSTIC_TAG_UNNECESSARY DiagnosticTag = 1
	DIAGNOSTIC_TAG_DEPRECATED  DiagnosticTag = 2
)

func (tag DiagnosticTag) ToString() string {
	switch tag {
	case DIAGNOSTIC_TAG_UNNECESSARY:
		return "Unnecessary"
	case DIAGNOSTIC_TAG_DEPRECATED:
		return "Deprecated"
	default:
		return "Unknown"
	}
}
