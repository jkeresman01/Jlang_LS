package types

import (
	"fmt"
)

type CodeAction struct {
	Title          string         `json:"title"`
	CodeActionKind CodeActionKind `json:"kind"`
	IsPreferred    bool           `json:"isPreferred"`
	WorkspaceEdit  WorkspaceEdit  `json:"edit"`
}

func (c CodeAction) ToString() string {
	return fmt.Sprintf(
		"Title: %s\r\nKind: %d\r\nIs preferred: %s\r\nEdit: %s\r\n",
		c.Title,
		c.CodeActionKind,
		convertBoolToString(c.IsPreferred),
		c.WorkspaceEdit.ToString(),
	)
}

func convertBoolToString(isTrue bool) string {
	if isTrue {
		return "YES"
	}

	return "NO"
}
