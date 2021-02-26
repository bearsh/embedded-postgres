package embeddedpostgres

import (
	"fmt"

	"github.com/mholt/archiver/v3"
)

type Unpacker func(source string, destination string) error

func defaultUnpacker() Unpacker {
	return func(source, destination string) error {
		if err := archiver.NewTarXz().Unarchive(source, destination); err != nil {
			return fmt.Errorf("unable to extract postgres archive %s to %s", source, destination)
		}
		return nil
	}
}
