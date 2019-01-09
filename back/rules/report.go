package rules

import (
	"github.com/Salibert/Gomoku/back/server/inter"
)

// ReportCheckRules create a report of this move
type ReportCheckRules struct {
	// UPDATE RESET AND CLONE IF MODIF REPORTCHECKRULES
	ListCapturedStone []*inter.Node
	ItIsAValidMove    bool
	PartyFinish       bool
	WinOrLose         [][]*inter.Node
	NextMovesOrLose   []*inter.Node
	NbFreeThree       int
	SizeAlignment     int
	NbBlockStone      int
	LevelCapture      int
	AmbientScore      int
}

// Clone create a uniform object, Dont clone slicer but init
func (report *ReportCheckRules) Clone() *ReportCheckRules {
	clone := &ReportCheckRules{}
	clone.ListCapturedStone = make([]*inter.Node, 0, 16)
	clone.ItIsAValidMove = report.ItIsAValidMove
	clone.PartyFinish = report.PartyFinish
	clone.WinOrLose = make([][]*inter.Node, 0, 8)

	clone.NextMovesOrLose = make([]*inter.Node, 0, 16)
	clone.NbFreeThree = report.NbFreeThree
	clone.SizeAlignment = report.SizeAlignment
	clone.NbBlockStone = report.NbBlockStone
	clone.LevelCapture = report.LevelCapture
	clone.AmbientScore = report.AmbientScore
	return clone
}

// Reset all report value
func (report *ReportCheckRules) Reset() {
	report.ListCapturedStone = report.ListCapturedStone[:0]
	report.WinOrLose = report.WinOrLose[:0]
	report.NextMovesOrLose = report.NextMovesOrLose[:0]
	report.PartyFinish = false
	report.ItIsAValidMove = false
	report.NbFreeThree = 0
	report.SizeAlignment = 0
	report.NbBlockStone = 0
	report.LevelCapture = 0
	report.AmbientScore = 0
}
