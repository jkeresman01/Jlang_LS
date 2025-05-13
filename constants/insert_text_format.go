package constants

type InsertTextFormat int

const (
	INSERT_TEXT_FORMAT_PLAIN_TEXT InsertTextFormat = 1
	INSERT_TEXT_FORMAT_SNIPPET    InsertTextFormat = 2
)

func (format InsertTextFormat) ToString() string {
	switch format {
	case INSERT_TEXT_FORMAT_PLAIN_TEXT:
		return "PlainText"
	case INSERT_TEXT_FORMAT_SNIPPET:
		return "Snippet"
	default:
		return "Unknown"
	}
}
