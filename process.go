package main

import (
	"errors"
	"fmt"
	"strconv"
)

func createMixArray(argumentString []string, c string) MixArray {
	var IDs []int

	for _, i := range argumentString {
		if i == "f" {
			IDs = append(IDs, -1)
		} else {
			convI, _ := strconv.Atoi(i)
			IDs = append(IDs, convI)
		}
	}

	return MixArray{IDs, c}
}

func (m *MixArray) execute(r *Recipe) error {
	if !r.Locked {
		return errors.New("RecipeNotLocked")
	}

	for _, i := range m.IDs {
		if i >= len(r.ExecIngredients) {
			return errors.New("OutOfBounds")
		}
	}

	for _, i := range m.IDs {

		if i < 0 {
			if r.ExecFlower.Amount == 0 {
				return errors.New("RedundantCommand")
			}

			r.ExecFlower.Amount = 0
		} else {
			if r.ExecIngredients[i].Amount == 0 {
				return errors.New("RedundantCommand")
			}
			r.ExecIngredients[i].Amount = 0
		}
	}

	return nil
}

func (m *MixArray) print(r *Recipe) error {
	if len(m.IDs) == 0 {
		return errors.New("EmptyMixArray")
	}

	argumentString := ""
	for _, i := range m.IDs {
		if i < 0 {
			argumentString = argumentString + " f"
		} else {
			argumentString = argumentString + " " + strconv.Itoa(i)
		}
	}

	CommentString := ""
	if m.Comment != "" {
		CommentString = " [" + m.Comment + "]"
	}

	fmt.Println("mix" + argumentString + CommentString)
	return nil
}

func createMixSingle(argumentString []string, r *Recipe, c string) MixSingle {
	ID, _ := strconv.Atoi(argumentString[0])
	Amount, _ := strconv.Atoi(argumentString[1])

	Factor := float64(Amount) / float64(r.Flower.Amount)
	return MixSingle{ID, Factor, c}
}

func (m *MixSingle) execute(r *Recipe) error {
	if !r.Locked {
		return errors.New("RecipeNotLocked")
	}

	if m.ID > len(r.ExecIngredients) {
		return errors.New("OutOfBounds")
	}

	Amount := m.Factor * float64(r.Flower.Amount)
	AmountInt := int(Amount)

	if r.ExecIngredients[m.ID].Amount-AmountInt < 0 {
		return errors.New("TooMuch")
	} else {
		r.ExecIngredients[m.ID].Amount = r.ExecIngredients[m.ID].Amount - AmountInt
	}

	return nil
}

func (m *MixSingle) print(r *Recipe) error {
	IDstr := strconv.Itoa(m.ID)

	Amount := m.Factor * float64(r.Flower.Amount)
	AmountInt := int(Amount)
	AmountStr := strconv.Itoa(AmountInt)

	fmt.Println("smix " + IDstr + " " + AmountStr)

	return nil
}
