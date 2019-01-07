package rules

import (
	"github.com/Salibert/Gomoku/back/server/inter"
)

type ListStone [16]inter.Node
type WinListStone [5]inter.Node

// ReportCheckRules create a report of this move
type ReportCheckRules struct {
	ListCapturedStone      *ListStone
	IndexListCapturedStone int
	ItIsAValidMove         bool
	PartyFinish            bool
	NextMovesOrLose        *ListStone
	WinOrLose              [][]*inter.Node
	IndexNextMovesOrLose   int
	NbFreeThree            int
	SizeAlignment          int
	NbBlockStone           int
	LevelCapture           int
	AmbientScore           int
}

// Clone create a uniform object, Dont clone slicer but init
func (report *ReportCheckRules) Clone() *ReportCheckRules {
	clone := &ReportCheckRules{}
	clone.ListCapturedStone = &ListStone{}
	clone.ItIsAValidMove = report.ItIsAValidMove
	clone.PartyFinish = report.PartyFinish
	clone.WinOrLose = report.WinOrLose[:0]
	clone.NextMovesOrLose = &ListStone{}
	clone.NbFreeThree = report.NbFreeThree
	clone.SizeAlignment = report.SizeAlignment
	clone.NbBlockStone = report.NbBlockStone
	clone.LevelCapture = report.LevelCapture
	clone.AmbientScore = report.AmbientScore
	return clone
}

// Reset all report value
func (report *ReportCheckRules) Reset() {
	report.ListCapturedStone = &ListStone{}
	report.WinOrLose = report.WinOrLose[:0]
	report.NextMovesOrLose = &ListStone{}
	report.PartyFinish = false
	report.ItIsAValidMove = false
	report.NbFreeThree = 0
	report.SizeAlignment = 0
	report.NbBlockStone = 0
	report.LevelCapture = 0
	report.AmbientScore = 0
}
