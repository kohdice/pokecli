package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGraphqlClient_CallApi_PokedexNumber(t *testing.T) {
	data := []struct {
		name     string
		number   int
		expected graphqlResponse
	}{
		{"Get ピカチュウ", 25, graphqlResponse{
			Data: &data{
				PokemonByPokedexNumber: &pokemon{
					ID:                    "UG9rZW1vbjoyNQ==",
					NationalPokedexNumber: 25,
					HP:                    35,
					Attack:                55,
					Defense:               40,
					SpecialAttack:         50,
					SpecialDefense:        50,
					Speed:                 90,
					BaseTotal:             320,
					Types: []pokemonsType{
						{PokemonType: pokemonType{ID: "UG9rZW1vblR5cGU6NA==", TypeName: "でんき"}, Slot: 1},
					},
					Abilities: []pokemonsAbility{
						{PokemonAbility: pokemonAbility{ID: "UG9rZW1vbkFiaWxpdHk6OQ==", AbilityName: "せいでんき"}, Slot: 1, IsHidden: false},
						{PokemonAbility: pokemonAbility{ID: "UG9rZW1vbkFiaWxpdHk6MzE=", AbilityName: "ひらいしん"}, Slot: 3, IsHidden: true},
					},
				},
				PokemonByName: nil,
			},
			Errors: &[]errorDetail{},
		}},
		{"Get けつばん", 152, graphqlResponse{
			Data: &data{
				PokemonByPokedexNumber: nil,
				PokemonByName:          nil,
			},
			Errors: &[]errorDetail{},
		}},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				responseBytes, _ := json.Marshal(d.expected)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, err := w.Write(responseBytes)
				if err != nil {
					t.Errorf("error writing response: %v", err)
				}
			}))
			defer mockServer.Close()

			client := GraphqlClient{url: mockServer.URL}
			_, err := client.callApi("pokemonByPokedexNumberQuery", map[string]interface{}{})
			if err != nil {
				t.Errorf("error calling API: %v", err)
			}
		})
	}
}

func TestGraphqlClient_CallApi_PokemonName(t *testing.T) {
	data := []struct {
		name        string
		pokemonName string
		expected    graphqlResponse
	}{
		{"Get Pikachu", "ピカチュウ", graphqlResponse{
			Data: &data{
				PokemonByPokedexNumber: nil,
				PokemonByName: &pokemon{
					ID:                    "UG9rZW1vbjoyNQ==",
					NationalPokedexNumber: 25,
					HP:                    35,
					Attack:                55,
					Defense:               40,
					SpecialAttack:         50,
					SpecialDefense:        50,
					Speed:                 90,
					BaseTotal:             320,
					Types: []pokemonsType{
						{PokemonType: pokemonType{ID: "UG9rZW1vblR5cGU6NA==", TypeName: "でんき"}, Slot: 1},
					},
					Abilities: []pokemonsAbility{
						{PokemonAbility: pokemonAbility{ID: "UG9rZW1vbkFiaWxpdHk6OQ==", AbilityName: "せいでんき"}, Slot: 1, IsHidden: false},
						{PokemonAbility: pokemonAbility{ID: "UG9rZW1vbkFiaWxpdHk6MzE=", AbilityName: "ひらいしん"}, Slot: 3, IsHidden: true},
					},
				},
			},
			Errors: &[]errorDetail{},
		}},
		{"Get けつばん", "けつばん", graphqlResponse{
			Data: &data{
				PokemonByPokedexNumber: nil,
				PokemonByName:          nil,
			},
			Errors: &[]errorDetail{},
		}},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				responseBytes, _ := json.Marshal(d.expected)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, err := w.Write(responseBytes)
				if err != nil {
					t.Errorf("error writing response: %v", err)
				}
			}))
			defer mockServer.Close()

			client := GraphqlClient{url: mockServer.URL}
			_, err := client.callApi("pokemonByPokedexNumberQuery", map[string]interface{}{})
			if err != nil {
				t.Errorf("error calling API: %v", err)
			}
		})
	}
}

func TestGraphqlClient_CallApi_ErrorHandling(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer mockServer.Close()

	client := GraphqlClient{url: mockServer.URL}
	_, err := client.callApi("testQuery", map[string]interface{}{})
	if err == nil {
		t.Error("expected an error, but got none")
	}

	expectedErrorMessage := "failed parsing JSON: unexpected end of JSON input"
	if !strings.Contains(err.Error(), expectedErrorMessage) {
		t.Errorf("expected error message to contain '%s', but got '%s'", expectedErrorMessage, err.Error())
	}
}
