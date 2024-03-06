package lectures

import (
	"fmt"
	"reflect"
)

// Maps - what are they?, Creating, Manipulation
// Structs - what are they?, Creating, Name Conventions,Embedding, Tags

func Agenda5() {
	embedding()
}

func maps() {
	states := map[string]string{
		"SP": "São Paulo",
		"RJ": "Rio de Janeiro",
	}

	fmt.Println(states["SP"])
	fmt.Println(states)
	fmt.Println("Length: ", len(states))

	states["BA"] = "Bahia"
	fmt.Println(states)

	delete(states, "RJ")
	fmt.Println(states)
	pop, ok := states["RJ"]
	fmt.Printf("Value: %v, founded? %v\n", pop, ok)

	states = make(map[string]string)
	fmt.Println(states)
}

func structs() {
	type Doctor struct {
		number     int
		actorName  string
		companions []string
	}

	aDoc := Doctor{
		number:    2,
		actorName: "Leonardo F.",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoc)

	anonymousDoctor := struct{ name string }{"Leonardo Anonymous"}
	anotherDoctor := anonymousDoctor // Deep Copy
	anotherDoctor.name = "Tom Baker"
	fmt.Println(anonymousDoctor)
	fmt.Println(anotherDoctor)

}

func embedding() {
	type animal struct {
		name   string `required max:"100"`
		origin string
	}

	type bird struct {
		animal   // embedding
		speedKPH float32
		canFly   bool
	}

	b := bird{
		canFly:   true,
		speedKPH: 12.3,
		animal: animal{
			origin: "naruto",
			name:   "blue bird"},
	}
	fmt.Println(b)

	t := reflect.TypeOf(animal{})
	field, _ := t.FieldByName("name")
	fmt.Println("tag: ", field.Tag)
}
