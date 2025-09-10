package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	grossUnits := make(map[string]int)
	grossUnits["quarter_of_a_dozen"] = 3
	grossUnits["half_of_a_dozen"] = 6
	grossUnits["dozen"] = 12
	grossUnits["small_gross"] = 120
	grossUnits["gross"] = 144
	grossUnits["great_gross"] = 1728
	return grossUnits
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	emptyBill := make(map[string]int)
	return emptyBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
	_, itemExists := bill[item]
	if !exists {
		return false
	} else if itemExists {
		bill[item] += units[unit]
		return true
	} else {
		bill[item] = units[unit]
		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, itemsExists := bill[item]
	_, unitExists := units[unit]
	newQty := bill[item] - units[unit]
	if !itemsExists || !unitExists || newQty < 0 {
		return false
	} else if newQty == 0 {
		delete(bill, item)
		return true
	} else {
		bill[item] = newQty
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, itemExists := bill[item]
	return qty, itemExists
}
