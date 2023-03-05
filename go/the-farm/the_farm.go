package thefarm

import (
	"fmt"
	"errors"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct{cows int}

func (err *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", err.cows)
}

var ErrNegativeFodder = errors.New("negative fodder")
var ErrNoCows = errors.New("division by zero")

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()

	errFodder := 0.0

	switch {
	case err == ErrScaleMalfunction && fodder > 0:
		fodder = fodder * 2
		err = nil
	case (err == ErrScaleMalfunction || err == nil) && fodder < 0:
		fodder = errFodder
		err = ErrNegativeFodder
	case cows == 0:
		err = ErrNoCows
	case cows < 0:
		err = &SillyNephewError{cows}
	}

	if err != nil {
		return errFodder, err
	}

	return fodder / float64(cows), err
}
