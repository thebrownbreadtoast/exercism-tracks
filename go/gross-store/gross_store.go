package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int)

	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	addQuantity, quantityExists := units[unit]

	if quantityExists {
		_, itemExists := bill[item]

		if itemExists {
			bill[item] += addQuantity
		} else {
			bill[item] = addQuantity
		}

		return true
	}

	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	removeQuantity, quantityExists := units[unit]

	if quantityExists {
		_, itemExists := bill[item]

		if itemExists {
			currentQuantity := bill[item]

			updatedQuantity := currentQuantity - removeQuantity

			if updatedQuantity < 0 {
				return false
			} else if updatedQuantity == 0 {
				delete(bill, item)
			} else {
				bill[item] = updatedQuantity
			}

			return true
		}
	}

	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemQuantity, itemExists := bill[item]

	return itemQuantity, itemExists
}
