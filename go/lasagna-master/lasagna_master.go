package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgLayerCookingTime int) int {
	if avgLayerCookingTime <= 0 {
		avgLayerCookingTime = 2
	}

	return len(layers) * avgLayerCookingTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	requiredNoodles := 0
	requiredSauce := 0.0

	for idx := 0; idx < len(layers); idx++ {
		switch {
		case layers[idx] == "noodles":
			requiredNoodles += 50
		case layers[idx] == "sauce":
			requiredSauce += 0.2
		}
	}

	return requiredNoodles, requiredSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(secretLayers, myLayers []string) {
	myLayers[len(myLayers) - 1] = secretLayers[len(secretLayers) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaledQuantities []float64

	floatPortions := float64(portions)

	for idx := 0; idx < len(quantities); idx++ {
		scaledQuantity := quantities[idx] * 0.5 * floatPortions

		scaledQuantities = append(scaledQuantities, scaledQuantity)
	}

	return scaledQuantities
}