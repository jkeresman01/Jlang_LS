package constants

type TraceValue int32

const (
	TRACE_VALUE_OFF TraceValue = iota
	TRACE_VALUE_MESSAGES
	TRACE_VALUE_VERBOSE
)

func (traceValue TraceValue) ToString() string {
	switch traceValue {
	case TRACE_VALUE_OFF:
		return "off"
	case TRACE_VALUE_MESSAGES:
		return "messages"
	case TRACE_VALUE_VERBOSE:
		return "verbose"
	default:
		return "unknown"
	}
}
