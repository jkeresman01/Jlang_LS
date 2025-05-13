package constants

type MarkupKind int32

const (
	MARKUP_KIND_PLAIN_TEXT MarkupKind = iota
	MARKUP_KIND_MARKDOWN
	MARKUP_KIND_OPTIONS_NO
)

func (kind MarkupKind) ToString() string {
	switch kind {
	case MARKUP_KIND_PLAIN_TEXT:
		return "PlainText"
	case MARKUP_KIND_MARKDOWN:
		return "Markdown"
	case MARKUP_KIND_OPTIONS_NO:
		return "OptionsNo"
	default:
		return "Unknown"
	}
}
