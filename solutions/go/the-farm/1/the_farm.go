package thefarm

import "errors"
import "fmt"

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
    number_of_cows int
}

func (e SillyNephewError) Error() string {
    return fmt.Sprintf("silly nephew, there cannot be %d cows", e.number_of_cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
    var err error
    amount := 0.0
    
    if cows == 0 {
        err = errors.New("division by zero")
    }

    if cows < 0 {
        err = SillyNephewError{ number_of_cows: cows }
    }

    if err != nil {
        return amount, err
    }

    amount, err = weightFodder.FodderAmount()

    if amount < 0 {
        if err == nil || err == ErrScaleMalfunction {
            err = errors.New("negative fodder")
        }
    	amount = 0
    }

    if err == ErrScaleMalfunction {
        amount = amount * 2.0
        err = nil
    }

    if err != nil {
        amount = 0
    }

    return amount / float64(cows), err
}
