package rules

import (
	"github.com/Salibert/Gomoku/back/axis"
	"github.com/Salibert/Gomoku/back/server/inter"
	pb "github.com/Salibert/Gomoku/back/server/pb"
)

// FuncCheckRules ...
type FuncCheckRules func(schema Schema, list *axis.Radius, index, lenRadius int) bool

// Schema ...
type Schema struct {
	Schema    [][]int
	FuncCheck []FuncCheckRules
	Report    *ReportCheckRules
}

type rules int

var rulesFuncArray = []func(p, o int) []int{
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

// parseRules use pb.ConfigRules for create a array of function checking
func parseRules(config pb.ConfigRules) []FuncCheckRules {
	arrayFuncRulse := make([]FuncCheckRules, 0, 6)
	if config.IsActiveRuleFreeThree == true {
		arrayFuncRulse = append(arrayFuncRulse, checkFreeThreeNoSpace, checkFreeThreeSpace)
	}
	if config.IsActiveRuleCapture == true {
		arrayFuncRulse = append(arrayFuncRulse, checkCapture)
	}
	if config.IsActiveRuleAlignment == true {
		arrayFuncRulse = append(arrayFuncRulse, checkAlignment)
	}
	if config.IsActiveRuleWin == true {
		arrayFuncRulse = append(arrayFuncRulse, checkWin)
	}
	if config.IsActiveRuleBlock == true {
		arrayFuncRulse = append(arrayFuncRulse, checkBlock)
	}
	if config.IsActiveRuleProbableCapture == true {
		arrayFuncRulse = append(arrayFuncRulse, probableCapture)
	}
	if config.IsActiveRuleAmbientSearch == true {
		arrayFuncRulse = append(arrayFuncRulse, ambientScore)
	}
	return arrayFuncRulse
}

// New create a instance of Schema
func New(playerIndex, opposent int, config pb.ConfigRules) Schema {
	checker := Schema{
		FuncCheck: parseRules(config),
		Schema:    make([][]int, 5, 5),
		Report: &ReportCheckRules{
			ListCapturedStone: make([]*inter.Node, 0, 16),
			WinOrLose:         make([][]*inter.Node, 0, 8),
			NextMovesOrLose:   make([]*inter.Node, 0, 10),
		},
	}
	for i, f := range rulesFuncArray {
		checker.Schema[i] = f(playerIndex, opposent)
	}
	return checker
}

// Clone copy Schema struct
func (schema Schema) Clone() *Schema {
	clone := &Schema{}
	clone.Schema = schema.Schema[:len(schema.Schema)]
	clone.FuncCheck = schema.FuncCheck[:len(schema.FuncCheck)]
	clone.Report = schema.Report.Clone()
	return clone
}

func compareNodesSchema(list *axis.Radius, schema []int, index, lenRadius, direction int) bool {
	for _, player := range schema {
		if 0 <= index && index < lenRadius && player == list[index].Player {
			index += direction
		} else {
			return false
		}
	}
	return true
}

func checkFreeThreeNoSpace(schema Schema, list *axis.Radius, index, lenRadius int) bool {
	var isSuccess bool
	for i := 3; i > 0; i-- {
		if isSuccess = compareNodesSchema(list, schema.Schema[FreeThreeNoSpace], index-i, lenRadius, 1); isSuccess == true {
			schema.Report.NbFreeThree++
			return true
		}
	}
	return false
}

func checkFreeThreeSpace(schema Schema, list *axis.Radius, index, lenRadius int) bool {
	var isFreeThree1, isFreeThree2 bool
	if isFreeThree1 = compareNodesSchema(list, schema.Schema[FreeThreeSpace], index-1, lenRadius, 1); isFreeThree1 == true {
		schema.Report.NbFreeThree++
	}
	if isFreeThree2 = compareNodesSchema(list, schema.Schema[FreeThreeSpace], index+1, lenRadius, -1); isFreeThree2 == true {
		schema.Report.NbFreeThree++
	}
	if isFreeThree1 == true || isFreeThree2 == true {
		return true
	}
	return false
}

func checkCapture(schema Schema, list *axis.Radius, index, lenRadius int) bool {
	var isCapture bool
	if isCapture = compareNodesSchema(list, schema.Schema[Capture], index, lenRadius, 1); isCapture == true {
		schema.Report.ListCapturedStone = append(schema.Report.ListCapturedStone, &list[index+1], &list[index+2])
	}
	if isCapture = compareNodesSchema(list, schema.Schema[Capture], index, lenRadius, -1); isCapture == true {
		schema.Report.ListCapturedStone = append(schema.Report.ListCapturedStone, &list[index-1], &list[index-2])
	}
	return isCapture
}

func createSliceWinOrLose(list *axis.Radius, start, end int) []*inter.Node {
	newList := make([]*inter.Node, 0, 11)
	for i := start; i < start+end; i++ {
		newList = append(newList, &list[i])
	}
	return newList
}

func checkWin(schema Schema, list *axis.Radius, index, lenRadius int) bool {
	schemaWin := schema.Schema[Win]
	lenSchema := len(schemaWin)
loop:
	for i := 0; i+lenSchema <= lenRadius; i++ {
		for index := 0; index < lenSchema; index++ {
			if list[i+index].Player != schemaWin[index] {
				continue loop
			}
		}
		schema.Report.WinOrLose = append(schema.Report.WinOrLose, createSliceWinOrLose(list, i, lenSchema))
		return true
	}
	return false
}

// ProccessCheckRules ...
func (schema Schema) ProccessCheckRules(list *axis.Radius, index, lenRadius int) {
	var isSuccessChecked bool
	for _, f := range schema.FuncCheck {
		if isSuccessChecked = f(schema, list, index, lenRadius); isSuccessChecked == true {
			return
		}
	}
}

func checkBlock(schema Schema, list *axis.Radius, index, lenRadius int) bool {
	var blockedDown, blockedUp int
	var blankDown, blankUp bool
	for i := index + 1; i < lenRadius; i++ {
		if list[i].Player != 0 && list[i].Player != list[index].Player {
			blockedUp++
			continue
		} else if list[i].Player == 0 && blockedUp > 2 {
			blankUp = true
		}
		break
	}
	for i := index - 1; i > 0; i-- {
		if list[i].Player != 0 && list[i].Player != list[index].Player {
			blockedDown++
			continue
		} else if list[i].Player == 0 && blockedDown > 2 {
			blankDown = true
		}
		break
	}
	if blankUp == true {
		blockedUp += 2
	}
	if blankDown == true {
		blockedDown += 2
	}
	if blockedDown == 3 && blankDown == true {
		blockedDown *= 2
	}
	if blockedUp == 3 && blankUp == true {
		blockedUp *= 2
	}
	blockedUp += blockedDown
	schema.Report.NbBlockStone += (blockedUp * blockedUp)
	return false
}

func ambientScore(schema Schema, list *axis.Radius, index, lenRadius int) bool {
	ambientScore := 0
	if index-1 >= 0 {
		switch list[index-1].Player {
		case 0:
			ambientScore += 2
		default:
			ambientScore++
		}
	}
	if index+1 < lenRadius {
		switch list[index+1].Player {
		case 0:
			ambientScore++
		default:
			ambientScore += 2
		}
	}
	schema.Report.AmbientScore += ambientScore
	return false
}

func checkAlignment(schema Schema, list *axis.Radius, index, lenRadius int) bool {
	var blancAlignmentUp, blancAlignmentDown bool
	var i, alignment int
	for i = index + 1; i < lenRadius; i++ {
		if list[i].Player == list[index].Player {
			alignment++

			continue
		} else if list[i].Player == 0 {
			blancAlignmentUp = true
		}
		break
	}
	for i = index - 1; i > 0; i-- {
		if list[i].Player == list[index].Player {
			alignment++
			continue
		} else if list[i].Player == 0 {
			blancAlignmentDown = true
		}
		break
	}
	if alignment == 3 {
		if blancAlignmentUp == true || blancAlignmentDown == true {
			alignment *= 2
		} else {
			alignment -= int(alignment / 2)
		}
	} else if alignment == 4 {
		alignment *= 4
	}
	schema.Report.SizeAlignment += (alignment * alignment)
	return false
}

func probableCapture(schema Schema, list *axis.Radius, index, lenRadius int) bool {
	var isSuccessChecked bool
	schemaProbableCapture := schema.Schema[ProbableCapture]
	levelCapture := 0
	if index != 0 && index != lenRadius-1 {
		switch list[index+1].Player {
		case list[index].Player:
			if isSuccessChecked = compareNodesSchema(list, schemaProbableCapture, index-1, lenRadius, 1); isSuccessChecked == true {
				levelCapture++
			} else if isSuccessChecked = compareNodesSchema(list, schemaProbableCapture, index+2, lenRadius, -1); isSuccessChecked == true {
				levelCapture++
			}
		case 0:
			if isSuccessChecked = compareNodesSchema(list, schemaProbableCapture, index-2, lenRadius, 1); isSuccessChecked == true {
				levelCapture++
			}
		default:
			if isSuccessChecked = compareNodesSchema(list, schemaProbableCapture, index+1, lenRadius, -1); isSuccessChecked == true {
				levelCapture++
			}
		}
	}
	schema.Report.LevelCapture += levelCapture
	return false
}

// CheckIfPartyIsFinish ...
func (schema Schema) CheckIfPartyIsFinish(list *axis.Radius, index, lenRadius int) {
	var isSuccessChecked bool
	schemaProbableCapture := schema.Schema[ProbableCapture]
	if index != 0 && index != lenRadius-1 {
		switch list[index+1].Player {
		case list[index].Player:
			if isSuccessChecked = compareNodesSchema(list, schemaProbableCapture, index-1, lenRadius, 1); isSuccessChecked == true {
				schema.Report.NextMovesOrLose = append(schema.Report.NextMovesOrLose, &list[index+2])
			} else if isSuccessChecked = compareNodesSchema(list, schemaProbableCapture, index+2, lenRadius, -1); isSuccessChecked == true {
				schema.Report.NextMovesOrLose = append(schema.Report.NextMovesOrLose, &list[index-1])
			}
		case 0:
			if isSuccessChecked = compareNodesSchema(list, schemaProbableCapture, index-2, lenRadius, 1); isSuccessChecked == true {
				schema.Report.NextMovesOrLose = append(schema.Report.NextMovesOrLose, &list[index+1])
			}
		default:
			if isSuccessChecked = compareNodesSchema(list, schemaProbableCapture, index+1, lenRadius, -1); isSuccessChecked == true {
				schema.Report.NextMovesOrLose = append(schema.Report.NextMovesOrLose, &list[index-2])
			}
		}
	}
	return
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

func createProbableCapture(player, opposent int) []int {
	return []int{opposent, player, player, 0}
}
