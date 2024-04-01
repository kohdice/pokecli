package main

// TODO: When constants are used, remove the nolint comments.
// nolint:unused // The part where this constant is used has not been implemented
const pokemonInfoFragment = `
	fragment PokemonInfo on Pokemon {
		id
		nationalPokedexNumber
		name
		hp
		attack
		defense
		specialAttack
		specialDefense
		speed
		baseTotal
		types {
			pokemonType {
				id
				typeName
			}
			slot
		}
		abilities {
			pokemonAbility {
				id
				abilityName
			}
			slot
			isHidden
		}
	}
`

// TODO: When constants are used, remove the nolint comments.
// nolint:unused // The part where this constant is used has not been implemented.
const pokemonByPokedexNumberQuery = `
	query getPokemonByPokedexNumber($number:Int!) {
		pokemonByPokedexNumber(pokedexNumber: $number) {
			...PokemonInfo
		}
	}
` + pokemonInfoFragment

// TODO: When constants are used, remove the nolint comments.
// nolint:unused // The part where this constant is used has not been implemented.
const pokemonByNameQuery = `
	query getPokemonByName($name:String!) {
		pokemonByName(name: $name) {
			...PokemonInfo
		}
	}
` + pokemonInfoFragment

type graphqlResponse struct {
	Data   *data          `json:"data"`
	Errors *[]errorDetail `json:"errors"`
}

type data struct {
	PokemonByPokedexNumber *pokemon `json:"pokemonByPokedexNumber"`
	PokemonByName          *pokemon `json:"pokemonByName"`
}

type errorDetail struct {
	Message   string         `json:"message"`
	Locations []errorLocaton `json:"locations"`
}

type errorLocaton struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type pokemon struct {
	ID                    string            `json:"id"`
	NationalPokedexNumber int               `json:"nationalPokedexNumber"`
	Name                  string            `json:"name"`
	HP                    int               `json:"hp"`
	Attack                int               `json:"attack"`
	Defense               int               `json:"defense"`
	SpecialAttack         int               `json:"specialAttack"`
	SpecialDefense        int               `json:"specialDefense"`
	Speed                 int               `json:"speed"`
	BaseTotal             int               `json:"baseTotal"`
	Types                 []pokemonsType    `json:"types"`
	Abilities             []pokemonsAbility `json:"abilities"`
}

type pokemonType struct {
	ID       string `json:"id"`
	TypeName string `json:"typeName"`
}

type pokemonsType struct {
	PokemonType pokemonType `json:"pokemonType"`
	Slot        int         `json:"slot"`
}

type pokemonAbility struct {
	ID          string `json:"id"`
	AbilityName string `json:"abilityName"`
}

type pokemonsAbility struct {
	PokemonAbility pokemonAbility `json:"pokemonAbility"`
	Slot           int            `json:"slot"`
	IsHidden       bool           `json:"isHidden"`
}
