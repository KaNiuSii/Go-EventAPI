package main

import (
	"fmt"
	"math/rand"
)

type Character struct {
	Event   string `json:"event"`
	Name    string `json:"name"`
	Race    string `json:"race"`
	Ability string `json:"ability"`
}

func generateRandomCharacter(event string) (Character, error) {
	basePath := fmt.Sprintf("./assets/%s/", event)

	names, err := readFromFile(basePath + "names.txt")
	if err != nil {
		return Character{}, err
	}
	races, err := readFromFile(basePath + "race.txt")
	if err != nil {
		return Character{}, err
	}
	abilities, err := readFromFile(basePath + "ability.txt")
	if err != nil {
		return Character{}, err
	}

	name := names[rand.Intn(len(names))] + " " + names[rand.Intn(len(names))]
	race := races[rand.Intn(len(races))]
	ability := abilities[rand.Intn(len(abilities))]

	character := Character{
		Event:   event,
		Name:    name,
		Race:    race,
		Ability: ability,
	}

	return character, nil
}
