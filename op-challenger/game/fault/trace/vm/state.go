package vm

import (
	"encoding/json"
	"fmt"
	"io"
)

func ParseStateFromReader(in io.ReadCloser, state any) error {
	defer in.Close()
	if err := json.NewDecoder(in).Decode(state); err != nil {
		return fmt.Errorf("cannot decode state: %w", err)
	}
	return nil
}
