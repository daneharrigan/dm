package dm_test

import (
	"testing"
	. "github.com/daneharrigan/dm"
)

type Gender struct {
	Id string
	Label string
}

type Person struct {
	Id string
	FullName string
	GenderId string
	PersonId string

	Gender Gender
	Friends []*Person
}

func TestFind(t *testing.T) {
	person := new(Person)
	err := Find(person, "person-id")
	t.Logf("fn=Find error=%q", err)
}

func TestHasOne(t *testing.T) {
	person := new(Person)
	Find(person, "person-id")
	err := HasOne(person, person.Gender)
	t.Logf("fn=HasOne error=%q", err)
}

func TestHasMany(t *testing.T) {
	person := new(Person)
	Find(person, "person-id")
	err := HasMany(person, person.Friends)
	t.Logf("fn=HasOne error=%q", err)
}
