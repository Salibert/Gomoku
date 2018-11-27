package rules

import (
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

// Schema ...
type Schema struct {
	Schema [][]int32
	Report *ReportCheckRules
}

// ReportCheckRules create a report of this move
type ReportCheckRules struct {
	ListCapturedStone []*pb.Node
	PartyFinish       bool
	ListWin           [][]*pb.Node
	NextMovesOrWin    []*pb.Node
	NbFreeThree       int
}

type rules int

var rulesFuncArray = []func(p, o int32) []int32{
	createCapture,
	createFreeThreeSpace,
	createFreeThreeNoSpace,
	createWin,
	createProbableCapture,
}

const (
	// Capture ...
	Capture rules = iota
	// FreeThreeSpace ...
	FreeThreeSpace
	// FreeThreeNoSpace ...
	FreeThreeNoSpace
	// Win ...
	Win
	// ProbableCapture ...
	ProbableCapture
)

// New create a instance of Schema
func New(player int32) Schema {
	var opposent int32
	switch player {
	case 1:
		opposent = 2
	default:
		opposent = 1
	}
	checker := Schema{
		Schema: make([][]int32, 4, 4),
		Report: &ReportCheckRules{
			ListCapturedStone: make([]*pb.Node, 0, 16),
			ListWin:           make([][]*pb.Node, 0, 8),
		},
	}
	for i, f := range rulesFuncArray {
		checker.Schema[i] = f(player, opposent)
	}
	return checker
}

func compareNodesSchema(list []*pb.Node, schema []int32, index int, direction int) bool {
	len := len(list)
	for _, player := range schema {
		if 0 <= index && index < len && player == list[index].Player {
			index += direction
		} else {
			return false
		}
	}
	return true
}

func (schema Schema) checkFreeThreeNoSpace(list []*pb.Node, index int) bool {
	var isSuccess bool
	for i := 3; i > 0; i-- {
		if isSuccess = compareNodesSchema(list, schema.Schema[FreeThreeNoSpace], index-i, 1); isSuccess == true {
			schema.Report.NbFreeThree++
			return true
		}
	}
	return false
}

func (schema Schema) checkFreeThreeSpace(list []*pb.Node, index int) bool {
	var isFreeThree1, isFreeThree2 bool
	if isFreeThree1 = compareNodesSchema(list, schema.Schema[FreeThreeSpace], index-1, 1); isFreeThree1 == true {
		schema.Report.NbFreeThree++
	}
	if isFreeThree2 = compareNodesSchema(list, schema.Schema[FreeThreeSpace], index+1, -1); isFreeThree2 == true {
		schema.Report.NbFreeThree++
	}
	if isFreeThree1 == true || isFreeThree2 == true {
		return true
	}
	return false
}

func (schema Schema) checkCapture(list []*pb.Node, index int) bool {
	var isCapture bool
	if isCapture = compareNodesSchema(list, schema.Schema[Capture], index, 1); isCapture == true {
		schema.Report.ListCapturedStone = append(schema.Report.ListCapturedStone, list[index+1], list[index+2])
	}
	if isCapture = compareNodesSchema(list, schema.Schema[Capture], index, -1); isCapture == true {
		schema.Report.ListCapturedStone = append(schema.Report.ListCapturedStone, list[index-1], list[index-2])
	}
	return isCapture
}

func (schema Schema) checkWin(list []*pb.Node, index int) bool {
	schemaWin := schema.Schema[Win]
	lenSchema := len(schemaWin)
	lenList := len(list)
loop:
	for i := 0; i+lenSchema <= lenList; i++ {
		for index := 0; i < lenSchema; index++ {
			if list[i+index].Player != schemaWin[index] {
				continue loop
			}
		}
		schema.Report.ListWin = append(schema.Report.ListWin, list[i:lenSchema])
		return true
	}
	return false
}

// ProccessCheckRules ...
func (schema Schema) ProccessCheckRules(list []*pb.Node, index int) {
	var isSuccessChecked bool
	if isSuccessChecked = schema.checkCapture(list, index); isSuccessChecked == true {
		return
	} else if isSuccessChecked = schema.checkFreeThreeSpace(list, index); isSuccessChecked == true {
		return
	} else if isSuccessChecked = schema.checkFreeThreeNoSpace(list, index); isSuccessChecked == true {
		return
	} else if isSuccessChecked = schema.checkWin(list, index); isSuccessChecked == true {
		return
	}
	return
}

func (schema Schema) CheckIfPartyIsFinish(list []*pb.Node, index int) {
	var isSuccessChecked bool
	schemaProbableCapture := schema.Schema[ProbableCapture]
	len := len(list)
	if index != 0 && index != len-1 {
		switch list[index+1].Player {
		case list[index].Player:
			if isSuccessChecked = compareNodesSchema(list, schemaProbableCapture, index-1, 1); isSuccessChecked == true {
				schema.Report.NextMovesOrWin = append(schema.Report.NextMovesOrWin, list[index+2])
			} else if isSuccessChecked = compareNodesSchema(list, schemaProbableCapture, index+2, -1); isSuccessChecked == true {
				schema.Report.NextMovesOrWin = append(schema.Report.NextMovesOrWin, list[index-2])
			}
		case 0:
			if isSuccessChecked = compareNodesSchema(list, schemaProbableCapture, index-2, 1); isSuccessChecked == true {
				schema.Report.NextMovesOrWin = append(schema.Report.NextMovesOrWin, list[index+1])
			}
		default:
			if isSuccessChecked = compareNodesSchema(list, schemaProbableCapture, index+1, -1); isSuccessChecked == true {
				schema.Report.NextMovesOrWin = append(schema.Report.NextMovesOrWin, list[index-2])
			}
		}
	}
	return
}

func createCapture(player, opposent int32) []int32 {
	return []int32{player, opposent, opposent, player}
}

func createFreeThreeSpace(player, opposent int32) []int32 {
	return []int32{0, player, 0, player, player, 0}
}

func createFreeThreeNoSpace(player, opposent int32) []int32 {
	return []int32{0, player, player, player, 0}
}

func createWin(player, opposent int32) []int32 {
	return []int32{player, player, player, player, player}
}

func createProbableCapture(player, opposent int32) []int32 {
	return []int32{opposent, player, player, 0}
}
