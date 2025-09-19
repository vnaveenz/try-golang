package thefarm

import (
	"errors"
	"fmt"
	"strconv"
)

// var errGeneric = errors.New("something went wrong")
var errCows = errors.New("invalid number of cows")
// var errFattening = errors.New("factor could not be determined")
// var errFddrAmt = errors.New("amount could not be determined")

type InvalidCowsError struct {
  cowsCount int
  customMessage string
}

func (ice *InvalidCowsError) Error() string {
	return fmt.Sprintf("%s cows are invalid: %s", strconv.Itoa(ice.cowsCount), ice.customMessage)
}
func DivideFood(weightFodder FodderCalculator, cows int) (float64, error) {
	if fodder, err := weightFodder.FodderAmount(cows); err == nil {
		if fatFactor, err := weightFodder.FatteningFactor(); err == nil {
			return fodder * fatFactor / float64(cows), nil
		} else {
			return 0, err
		}
	} else {
		return 0, err
	}
}

func ValidateInputAndDivideFood(fc FodderCalculator, ncows int) (float64, error) {

	if ncows > 0 {
		return DivideFood(fc, ncows)
	} else {
		return 0, errCows
	}
}

func ValidateNumberOfCows(ncows int) error {
	if ncows < 0 {
		return &InvalidCowsError{ncows, "there are no negative cows"}
	} else if ncows == 0 {
		return &InvalidCowsError{ncows, "no cows don't need food"}
	} else {
		return nil
	}
}
