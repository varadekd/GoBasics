package methods

import "../structs"

// Purpose of this program is to access value of user defined type
func AccessPerson(p structs.Person) {
	p.Firstname = "Om"
	p.Lastname = "Narayan"
}
