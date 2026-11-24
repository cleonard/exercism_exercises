package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(ingredients []string) (noodles int, sauce float64) {
	gramsNoodlesPerLayer := 50
	gramsSaucePerLayer := 0.2
	noodles, sauce = 0, 0
	for i := range ingredients {
		if ingredients[i] == "noodles" {
			noodles += gramsNoodlesPerLayer
		}
		if ingredients[i] == "sauce" {
			sauce += gramsSaucePerLayer
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) (scaled []float64) {
	for i := 0; i < len(quantities); i++ {
		newQuantity := quantities[i] * float64(portions) / 2.0
		scaled = append(scaled, newQuantity)
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
