package main

type Type int

const (
	TypeNormal Type = iota
	TypeFire
)

type Status int

const (
	StatusParalysis Status = iota
	StatusPoison
	StatusSleep
	StatusFreeze
)

type MoveCategory int

const(
	MoveCategoryPhysical MoveCategory = iota
	MoveCategorySpecial
)

type MoveID int

type Move struct {
	ID MoveID
	Name string
	Category MoveCategory
	Accuracy uint8
	Power uint8
	ParalysisRisk uint8
	SleepRisk uint8
	PoisonRisk uint8
	FreezeRisk uint8
}

type Experience uint64

type Level uint8

type ExperienceGroup int

const (
	ExperienceGroupFast ExperienceGroup = iota
	ExperienceGroupMediumFast
	ExperienceGroupMediumSlow
	ExperiemceGroupSlow
)

func (group ExperienceGroup) Level(exp Experience) Level {
	switch group {
	case ExperienceGroupFast:
		return Level(4 * exp * exp * exp / 5)
	case ExperienceGroupMediumFast:
		return Level(exp * exp * exp)
	case ExperienceGroupMediumSlow:
		return Level(6 * exp * exp * exp / 5 - 15 * exp * exp + 100 * exp - 140)
	case ExperiemceGroupSlow:
		return Level(5 * exp * exp * exp / 4)
	default:
		panic(fmt.Sprintf("invalid ExperienceGroup: %d", group))
	}
}

type SpeciesID int

const (
	SpeciesBulbasaur SpeciesID = iota
	SpeciesIvysaur
	SpeciesVenosaur
	SpeciesCharmander
	SpeciesCharmeleon
	SpeciesCharizard
)

type Species struct {
	ID SpeciesID
	Name string
	Type [2]Type
}

type PokemonID string

type Pokemon struct {
	ID PokemonID
	Name string
	Species SpeciesID
	Experience Experience
}