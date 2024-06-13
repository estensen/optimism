package cannon

import (
	"fmt"

	"github.com/ethereum-optimism/optimism/cannon/mipsevm"
	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/vm"
	"github.com/ethereum-optimism/optimism/op-service/ioutil"
)

func parseState(path string) (*mipsevm.State, error) {
	file, err := ioutil.OpenDecompressed(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open state file (%v): %w", path, err)
	}
	var state mipsevm.State
	if err := vm.ParseStateFromReader(file, &state); err != nil {
		return nil, fmt.Errorf("invalid mipsevm state: %w", err)
	}
	return &state, nil
}
