package constants

type CompletionTriggerKind int32

const (
	COMPLETION_TRIGGER_INVOKED                    CompletionTriggerKind = 1
	COMPLETION_TRIGGER_CHARACTER                  CompletionTriggerKind = 2
	COMPLETION_TRIGGER_FOR_INCOMPLETE_COMPLETIONS CompletionTriggerKind = 3
)

func (kind CompletionTriggerKind) ToString() string {
	switch kind {
	case COMPLETION_TRIGGER_INVOKED:
		return "Invoked"
	case COMPLETION_TRIGGER_CHARACTER:
		return "TriggerCharacter"
	case COMPLETION_TRIGGER_FOR_INCOMPLETE_COMPLETIONS:
		return "TriggerForIncompleteCompletions"
	default:
		return "Unknown"
	}
}
