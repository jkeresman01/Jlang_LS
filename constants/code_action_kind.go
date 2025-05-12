package constants

type CodeActionKind int32

const (
	CODE_ACTION_EMPTY CodeActionKind = iota
	CODE_ACTION_QUICK_FIX
	CODE_ACTION_REFACTOR
)

func (kind CodeActionKind) String() string {
	switch kind {
	case CODE_ACTION_EMPTY:
		return "Empty"
	case CODE_ACTION_QUICK_FIX:
		return "QuickFix"
	case CODE_ACTION_REFACTOR:
		return "Refactor"
	default:
		return "Unknown"
	}
}
