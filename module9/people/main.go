package main

import (
	"fmt"
	man "people/type"
)

func main() {
	var people map[string]man.Man
	people = make(map[string]man.Man)

	addPeople(people)

	suspects := []string{"Tony Stark", "John Doe", "Eren Yeager"}

	for _, suspect := range suspects {
		if _, ok := people[suspect]; ok {
			fmt.Println(people[suspect].Name, people[suspect].LastName, people[suspect].Age, people[suspect].Gender, people[suspect].Crimes)
		} else {
			fmt.Println("Not found")
		}
	}

}

func addPeople(people map[string]man.Man) {

	people["John Doe"] = man.Man{
		Name:     "John",
		LastName: "Doe",
		Age:      25,
		Gender:   "male",
		Crimes:   0,
	}

	people["Jane Doe"] = man.Man{
		Name:     "Jane",
		LastName: "Doe",
		Age:      25,
		Gender:   "female",
		Crimes:   0,
	}

	people["John Smith"] = man.Man{
		Name:     "John",
		LastName: "Smith",
		Age:      25,
		Gender:   "male",
		Crimes:   45,
	}

	people["Jane Smith"] = man.Man{
		Name:     "Jane",
		LastName: "Smith",
		Age:      25,
		Gender:   "female",
		Crimes:   0,
	}

	people["Alice Smith"] = man.Man{
		Name:     "Alice",
		LastName: "Smith",
		Age:      25,
		Gender:   "female",
		Crimes:   0,
	}

	people["Bob Smith"] = man.Man{
		Name:     "Bob",
		LastName: "Smith",
		Age:      34,
		Gender:   "male",
		Crimes:   5,
	}
}
