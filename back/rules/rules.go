package rules

import (
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

// Schema ...
type Schema struct {
	schema [][]int
	report *ReportCheckRules
}

// ReportCheckRules create a report of this move
type ReportCheckRules struct {
	listCapturedStone []pb.Node
	partyFinish       bool
	nbFreeThree       int
}

type rules int

var rulesFuncArray = []func(p, o int) []int{
	createCapture,
	createFreeThreeSpace,
	createFreeThreeNoSpace,
	createWin,
}

const (
	// Capture ...
	Capture rules = iota + 1
	// FreeThreeSpace ...
	FreeThreeSpace
	// FreeThreeNoSpace ...
	FreeThreeNoSpace
	// Win ...
	Win
)

// New create a instance of Schema
func New(player int) Schema {
	opposent := player * -1
	checker := Schema{
		schema: make([][]int, 4, 4),
		report: &ReportCheckRules{
			listCapturedStone: make([]pb.Node, 0, 16),
		},
	}
	for i, f := range rulesFuncArray {
		checker.schema[i] = f(player, opposent)
	}

	return checker
}

// ProccessCheckRules ...
func (schema Schema) ProccessCheckRules(list []pb.Node, index int) {

}

func createCapture(player, opposent int) []int {
	return []int{player, opposent, opposent, player}
}

func createFreeThreeSpace(player, opposent int) []int {
	return []int{0, player, 0, player, player, 0}
}

func createFreeThreeNoSpace(player, opposent int) []int {
	return []int{0, player, player, player, 0}
}

func createWin(player, opposent int) []int {
	return []int{player, player, player, player, player}
}
