package rules

import pb "github.com/Salibert/Gomoku/back/server/pb"

// ReportCheckRules create a report of this move
type ReportCheckRules struct {
	// UPDATE RESET AND CLONE IF MODIF REPORTCHECKRULES
	ListCapturedStone []*pb.Node
	ItIsAValidMove    bool
	PartyFinish       bool
	WinOrLose         [][]*pb.Node
	NextMovesOrLose   []*pb.Node
	NbFreeThree       int
	SizeAlignment     int
}

// Clone create a uniform object, Dont clone slicer but init
func (report *ReportCheckRules) Clone() *ReportCheckRules {
	clone := &ReportCheckRules{}
	clone.ListCapturedStone = report.ListCapturedStone[:0]
	clone.ItIsAValidMove = report.ItIsAValidMove
	clone.PartyFinish = report.PartyFinish
	clone.WinOrLose = report.WinOrLose[:0]
	clone.NextMovesOrLose = report.NextMovesOrLose[:0]
	clone.NbFreeThree = report.NbFreeThree
	clone.SizeAlignment = report.SizeAlignment
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
}
