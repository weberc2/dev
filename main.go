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