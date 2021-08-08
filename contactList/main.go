package main

import "fmt"

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Position  string
}

func (c *Contact) Info() string {
	return fmt.Sprint("id: ", c.ID, "\n", "name: ", c.FirstName, " ", c.LastName, "\n", "Position: ", c.Position, "\n")
}

func (c *Contact) PrintInfo() {
	fmt.Println(c.Info())
}

func NewContact(firstName, lastName, phone, email, position string, id int) *Contact {
	return &Contact{id, firstName, lastName, phone, email, position}
}

type IdTraker struct {
	id int
}

func (it *IdTraker) Get() int {
	it.id++
	return it.id
}

type ContactList struct {
	contacts map[int]*Contact
	id       *IdTraker
}

func NewContactList() *ContactList {
	return &ContactList{
		contacts: make(map[int]*Contact),
		id:       &IdTraker{},
	}
}

func (c *ContactList) Create(contact *Contact) {
	contact.ID = c.id.Get()
	c.contacts[contact.ID] = contact
}

func (c *ContactList) Update(contact *Contact) {
	c.contacts[contact.ID] = contact
}

func (c *ContactList) Get(id int) *Contact {
	return c.contacts[id]
}

func (c *ContactList) GetAll() map[int]*Contact {
	return c.contacts
}

func (c *ContactList) Delete(id int) {
	delete(c.contacts, id)
}

func main() {
	cl := SomeWorkThatReturnsListWithThreeContacts()

	fmt.Println("Before....")
	for _, contact := range cl.GetAll() {
		contact.PrintInfo()
	}

	contact1 := cl.Get(2)
	//contact1.PrintInfo()

	newContact := &Contact{
		FirstName: "Joel",
		LastName:  "Kinnaman",
		Position:  "Rick Flag",
	}

	updateСontact := &Contact{
		FirstName: "Christopher",
		LastName:  "Smith",
		Position:  "John Cena / Peacemaker",
		ID:        3,
	}

	cl.Delete(2)
	cl.Create(newContact)
	cl.Create(contact1)
	cl.Update(updateСontact)

	fmt.Println("After....")
	for _, contact := range cl.GetAll() {
		contact.PrintInfo()
	}

}

func SomeWorkThatReturnsListWithThreeContacts() *ContactList {
	contact1 := &Contact{
		FirstName: "Margot",
		LastName:  "Robbie",
		Position:  "Harley Quinn",
	}

	contact2 := &Contact{
		FirstName: "Idris",
		LastName:  "Elba",
		Position:  "Robert DuBois / Bloodsport",
	}

	contact3 := &Contact{
		FirstName: "John",
		LastName:  "Cena",
		Position:  "Christopher Smith / Peacemaker",
	}

	contactList := NewContactList()
	contactList.Create(contact1)
	contactList.Create(contact2)
	contactList.Create(contact3)

	return contactList
}
