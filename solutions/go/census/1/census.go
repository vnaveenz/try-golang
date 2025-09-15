// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	newRes := Resident{Name: name, Age: age, Address: address}
	return &newRes
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	v := *r
	if v.Name != "" && v.Address["street"] != "" {
		return true
	} else {
		return false
	}
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	ls := r
	ls.Name = ""
	ls.Age = 0
	ls.Address = nil

}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var count int
	for _,resp := range residents {
		if resp.HasRequiredInfo() {
			count++
		}
	}
	return count
}
