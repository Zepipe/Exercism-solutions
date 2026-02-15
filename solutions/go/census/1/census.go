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
	return &Resident{name, age, address}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
    nameIsMissing, addressIsMissing := false, false
	if r.Name == "" {
        nameIsMissing = true
    }
    
    if r.Address == nil || len(r.Address) == 0 {
        addressIsMissing = true
    } else {
        for key, val := range r.Address {
            if val == "" || key != "street" {
                addressIsMissing = true
                break
            }
        }
    }

    return !(nameIsMissing || addressIsMissing)
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
    r.Age = 0
    r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
    countResidents := 0
    for _, resident := range residents {
        if resident.HasRequiredInfo() {
            countResidents++
        }
    }
    return countResidents
}
