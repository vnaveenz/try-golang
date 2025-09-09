package lasagna

// define the 'PreparationTime()' function

func PreparationTime(layers []string, preptime int) (totalTime int) {
	if (preptime == 0) {
		preptime = 2
	}
	totalTime = len(layers) * preptime
	return
}

// func PreparationTime(layers []string, preptime int) int {
// 	totalTime := len(layers) * preptime
// 	return totalTime
// }

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for i :=0; i < len(layers); i++ {
		if (layers[i] == "noodles") {
			noodles += 50
		} else if (layers[i] == "sauce") {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(fl, myl []string) {
	myl[len(myl) -1] = fl[len(fl) -1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amtNeededfor2 []float64, numPortionsneeded int) (actualAmountsNeeded []float64) {
	actualAmountsNeeded = make([]float64, len(amtNeededfor2))
	for i := 0; i < len(amtNeededfor2); i++ {
		actualAmountsNeeded[i] = amtNeededfor2[i] * float64(numPortionsneeded) / 2
	}
	return
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
