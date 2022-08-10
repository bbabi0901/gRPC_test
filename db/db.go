package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	abpb "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/proto/ability"
	pkpb "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/proto/pokemon"
	typepb "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/proto/type"
	"github.com/bbabi0901/learngo/gRPC/gRPCPokemon/utils"
	"github.com/go-gota/gota/dataframe"
)

const (
	abilityCSV = "db/ability.csv"
	MoveDB     = "db/move.csv"
	PokemonCSV = "db/pokemon.csv"
	typeAPI    = "https://pokeapi.co/api/v2/type/"
	typeCSV    = "type.csv"
	typeJSON   = "type.json"
	numType    = 18
)

type PokemonType struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	DamageRelations struct {
		NoDamageTo       []struct{ Name string } `json:"no_damage_to"`
		HalfDamageTo     []struct{ Name string } `json:"half_damage_to"`
		DoubleDamageTo   []struct{ Name string } `json:"double_damage_to"`
		NoDamageFrom     []struct{ Name string } `json:"no_damage_from"`
		HalfDamageFrom   []struct{ Name string } `json:"half_damage_from"`
		DoubleDamageFrom []struct{ Name string } `json:"double_damage_from"`
	} `json:"damage_relations"`
	Pokemons []Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Slot        int `json:"slot"`
	PokemonName struct {
		Name string
	} `json:"pokemon"`
}

var df *dataframe.DataFrame
var f *os.File

func (p *PokemonType) unmarshalResponse(resp *http.Response) {
	data, err := ioutil.ReadAll(resp.Body)
	utils.HandleErr(err)

	err = json.Unmarshal(data, p)
	utils.HandleErr(err)
}
func TypeInfos() []*typepb.TypeInfo {
	var typeInfos []*typepb.TypeInfo

	for _, r := range scrapeTypeInfo() {
		p := PokemonType{}
		p.unmarshalResponse(r)

		var typeInfo typepb.TypeInfo

		typeInfo.TypeId = fmt.Sprint(p.ID)
		typeInfo.TypeName = p.Name
		typeInfo.NoDamageTo = appendTypes(typeInfo.NoDamageTo, p.DamageRelations.NoDamageTo)
		typeInfo.DoubleDamageTo = appendTypes(typeInfo.DoubleDamageTo, p.DamageRelations.DoubleDamageTo)
		typeInfo.HalfDamageTo = appendTypes(typeInfo.HalfDamageTo, p.DamageRelations.HalfDamageTo)
		typeInfo.DoubleDamageTo = appendTypes(typeInfo.DoubleDamageTo, p.DamageRelations.DoubleDamageTo)
		typeInfo.NoDamageFrom = appendTypes(typeInfo.NoDamageFrom, p.DamageRelations.NoDamageFrom)
		typeInfo.HalfDamageFrom = appendTypes(typeInfo.HalfDamageFrom, p.DamageRelations.HalfDamageFrom)
		typeInfo.DoubleDamageFrom = appendTypes(typeInfo.DoubleDamageFrom, p.DamageRelations.DoubleDamageFrom)

		typeInfos = append(typeInfos, &typeInfo)
	}

	return typeInfos
}

func TypePokemonInfos() []*typepb.PokemonsOfType {
	var typePokemonInfos []*typepb.PokemonsOfType

	for _, r := range scrapeTypeInfo() {
		p := PokemonType{}
		p.unmarshalResponse(r)

		var typePokemonInfo typepb.PokemonsOfType

		typePokemonInfo.TypeId = fmt.Sprint(p.ID)
		typePokemonInfo.TypeName = p.Name
		for _, p := range p.Pokemons {
			typePokemonInfo.PokemonName = append(typePokemonInfo.PokemonName, p.PokemonName.Name)
		}

		typePokemonInfos = append(typePokemonInfos, &typePokemonInfo)
	}

	return typePokemonInfos
}

func scrapeTypeInfo() []*http.Response {
	var responses []*http.Response
	for i := 1; i <= numType; i++ {
		resp := utils.GetResponse(typeAPI, i)
		responses = append(responses, resp)
	}
	return responses
}

func initDB(fileName string) *dataframe.DataFrame {
	if df == nil {
		f, err := os.Open(fileName)
		utils.HandleErr(err)

		dfPointer := dataframe.ReadCSV(f)
		df = &dfPointer
	}
	return df
}

func AbilityInfos() []*abpb.AbilityInfo {
	df = initDB(abilityCSV)

	var abilityInfos []*abpb.AbilityInfo

	for i := 0; i <= df.Nrow()-1; i++ {
		var a abpb.AbilityInfo

		a.AbilityId = fmt.Sprint(df.Elem(i, 0))
		a.AbilityName = fmt.Sprint(df.Elem(i, 1))
		a.Effect = fmt.Sprint(df.Elem(i, 2))

		abilityInfos = append(abilityInfos, &a)
	}

	return abilityInfos
}

func PokemonInfos() []*pkpb.PokemonInfo {
	df = initDB(PokemonCSV)

	var pokemonInfos []*pkpb.PokemonInfo

	for i := 0; i <= df.Nrow()-1; i++ {
		var p pkpb.PokemonInfo

		p.PokemonId = fmt.Sprint(df.Elem(i, 0))
		p.PokemonName = fmt.Sprint(df.Elem(i, 1))
		p.Type = append(p.Type, checkNone(fmt.Sprint(df.Elem(i, 2)), fmt.Sprint(df.Elem(i, 3)))...)
		p.Ability = append(p.Ability, checkNone(fmt.Sprint(df.Elem(i, 4)), fmt.Sprint(df.Elem(i, 5)), fmt.Sprint(df.Elem(i, 6)))...)

		pokemonInfos = append(pokemonInfos, &p)
	}

	return pokemonInfos
}

func checkNone(elements ...string) []string {
	var notNone []string
	for _, e := range elements {
		if e != "None" {
			notNone = append(notNone, e)
		}
	}
	return notNone
}

func appendTypes(ti []string, e []struct{ Name string }) []string {
	for _, t := range e {
		ti = append(ti, t.Name)
	}
	return ti
}

func Close() {
	f.Close()
}
