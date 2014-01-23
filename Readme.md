# DM

This repo is a proof of concept. Things written in the Readme may not have been
implemented yet.

```go
package main

import "github.com/daneharrigan/dm"

type Gender struct {
	Id   string
	Name string
}

type Person struct {
	Id       string
	FullName string
	GenderId string
	OfficeId string

	Gender Gender
	Friends []Person
}

func init() {
	dm.Validate(Person,
		dm.Presence{ Field: "Id" },
		dm.Presence{ Field: "GenderId" },
		dm.Includes{ Field: "GenderId", Values:  })	

	dm.BeforeSave(Person, func(person *Person) error {
		
	})
}

func main() {
	person := new(Person)
	if err := dm.Find(person, "uuid-goes-here", dm.Condition{ With: "Gender" }); err != nil {
		// log problem here
	}

	if err := dm.HasOne(person, person.Gender);  err != nil {
		// log problem here
	}

	if err := dm.HasMany(person, person.Friends);  err != nil {
		// log problem here
	}

	err := dm.Save(person)
}
```
