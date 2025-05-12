package types

import (
	"fmt"
)

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (c ClientInfo) ToString() string {
	return fmt.Sprintf("%s, version: %s", c.Name, c.Version)
}
