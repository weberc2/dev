package main

type Type int

const (
	TypeNormal Type = iota
	TypeFire
	TypeGrass
	TypeWater
)

type Status int

const (
	StatusParalysis Status = iota
	StatusPoison
	StatusSleep
	StatusFreeze
	StatusBurn
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

	// StatusLikelihoods lists the likelihood that the move will
	// change the target's status *if the move hits its target*.
	// Pokemon with an existing status condition will be immune to
	// status changes irrespective of the probability listed here.
	// Similarly, fire-type Pokemon cannot be burned, nor can
	// poison types be poisoned, nor ice-types frozen irrespective
	// of the probability listed here.
	StatusLikelihoods[Status]uint8
}

type Experience uint64

type Level uint8

const (
	LevelMax Level = 100
)

type ExperienceGroup int

const (
	ExperienceGroupFast ExperienceGroup = iota
	ExperienceGroupMediumFast
	ExperienceGroupMediumSlow
	ExperienceGroupSlow
	ExperienceGroupMax = ExperienceGroupSlow
)

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

type EvolutionLevel struct {
	Target SpeciesID
	MinLevel Level
}

var (
	LevelsByExperience = [ExperienceGroupMax][LevelMax]Experience{}
	ExperienceGroups = map[SpeciesID]ExperienceGroup{}
	EvolutionLevels = map[SpeciesID]EvolutionLevel{}
	EvolutionStoneFire = map[SpeciesID]SpeciesID{}
	EvolutionStoneWater = map[SpeciesID]SpeciesID{}
	EvolutionStoneThunder = map[SpeciesID]SpeciesID{}
	EvolutionTrade = map[SpeciesID]SpeciesID{}
)