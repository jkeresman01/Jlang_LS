package constants

type DiagnosticSeverity int32

const (
	DIAGNOSTIC_SEVERITY_ERROR       DiagnosticSeverity = 1
	DIAGNOSTIC_SEVERITY_WARNING     DiagnosticSeverity = 2
	DIAGNOSTIC_SEVERITY_INFORMATION DiagnosticSeverity = 3
	DIAGNOSTIC_SEVERITY_HINT        DiagnosticSeverity = 4
)

func (severity DiagnosticSeverity) ToString() string {
	switch severity {
	case DIAGNOSTIC_SEVERITY_ERROR:
		return "Error"
	case DIAGNOSTIC_SEVERITY_WARNING:
		return "Warning"
	case DIAGNOSTIC_SEVERITY_INFORMATION:
		return "Information"
	case DIAGNOSTIC_SEVERITY_HINT:
		return "Hint"
	default:
		return "Unknown"
	}
}
