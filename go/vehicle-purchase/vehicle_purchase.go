package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return (kind == "car" || kind == "truck")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	betterChoiceMsg := " is clearly the better choice."

	if option1 < option2 {
		return option1 + betterChoiceMsg
	} else {
		return option2 + betterChoiceMsg
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age >= 10 {
		return originalPrice * 0.50
	} else if age >= 3 {
		return originalPrice * 0.70
	} else {
		return originalPrice * 0.80
	}
}
