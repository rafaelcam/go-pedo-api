package validators

import (
	"errors"
)

func ValidateRange(start int, end int) error {
	if start < 0 || end < 0 {
		return errors.New("Only positive values ​​are accepted.")
	}

	if start > end {
		return errors.New("Start value can not be greater than the End value.")
	}

	return nil
}