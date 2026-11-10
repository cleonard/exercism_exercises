package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	quantity, unitExists := units[unit]
	if unitExists == false {
		return false
	}

	_, itemExists := bill[item]
	if itemExists == false {
		bill[item] = 0
	}
	bill[item] += quantity
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// Return false if the given item is not in the bill
	currentQty, ok := bill[item]
	if ok == false {
		return false
	}

	// Return false if the given unit is not in the units map.
	unitQty, ok := units[unit]
	if ok == false {
		return false
	}

	// Return false if the new quantity would be less than 0.
	if currentQty-unitQty < 0 {
		return false
	}

	// If the new quantity is 0, completely remove the item from the bill then return true.
	if currentQty-unitQty == 0 {
		delete(bill, item)
		return true
	}

	// Otherwise, reduce the quantity of the item and return true.
	bill[item] -= unitQty
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	return quantity, exists
}
