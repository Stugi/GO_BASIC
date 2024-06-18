package main

import (
	"fmt"
	man "people/type"
)

func main() {
	var people map[string]man.Man
	people = make(map[string]man.Man)

	addPeople(people)

	suspects := []string{"Tony Stark", "John Doe", "Eren Yeager", "Jane Doe"}

	var suspectWithMaxCrimes man.Man

	for _, suspect := range suspects {
		if _, ok := people[suspect]; ok {
			fmt.Println(people[suspect].Name, people[suspect].LastName, people[suspect].Age, people[suspect].Gender, people[suspect].Crimes)
			if suspectWithMaxCrimes == (man.Man{}) || people[suspect].Crimes > suspectWithMaxCrimes.Crimes {
				suspectWithMaxCrimes = people[suspect]
			}
		} else {
			fmt.Println("Not found")
		}
	}

	fmt.Println(suspectWithMaxCrimes)

}

func addPeople(people map[string]man.Man) {

	people["John Doe"] = man.Man{
		Name:     "John",
		LastName: "Doe",
		Age:      25,
		Gender:   "male",
		Crimes:   1,
	}

	people["Jane Doe"] = man.Man{
		Name:     "Jane",
		LastName: "Doe",
		Age:      25,
		Gender:   "female",
		Crimes:   12,
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
		Crimes:   2,
	}

	people["Alice Smith"] = man.Man{
		Name:     "Alice",
		LastName: "Smith",
		Age:      25,
		Gender:   "female",
		Crimes:   2,
	}

	people["Bob Smith"] = man.Man{
		Name:     "Bob",
		LastName: "Smith",
		Age:      34,
		Gender:   "male",
		Crimes:   5,
	}
}
