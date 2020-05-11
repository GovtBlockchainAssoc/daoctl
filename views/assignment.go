package views

import (
	"strconv"

	"github.com/alexeyco/simpletable"
	"github.com/eoscanada/eos-go"
	"github.com/hypha-dao/daoctl/models"
)

func assignmentHeader() *simpletable.Header {
	return &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Assigned"},
			{Align: simpletable.AlignCenter, Text: "Role"},
			{Align: simpletable.AlignCenter, Text: "Role Annually"},
			{Align: simpletable.AlignCenter, Text: "Time %"},
			{Align: simpletable.AlignCenter, Text: "Deferred %"},
			{Align: simpletable.AlignCenter, Text: "HUSD %"},
			{Align: simpletable.AlignCenter, Text: "HUSD"},
			{Align: simpletable.AlignCenter, Text: "HYPHA"},
			{Align: simpletable.AlignCenter, Text: "HVOICE"},
			{Align: simpletable.AlignCenter, Text: "Escrow SEEDS"},
			{Align: simpletable.AlignCenter, Text: "Liquid SEEDS"},
			{Align: simpletable.AlignCenter, Text: "Start Date"},
			{Align: simpletable.AlignCenter, Text: "End Date"},
		},
	}
}

// AssignmentTable returns a string representing a table of the assignnments
func AssignmentTable(assignments []models.Assignment) *simpletable.Table {

	table := simpletable.New()
	table.Header = assignmentHeader()

	husdTotal, _ := eos.NewAssetFromString("0.00 HUSD")
	hvoiceTotal, _ := eos.NewAssetFromString("0.00 HVOICE")
	hyphaTotal, _ := eos.NewAssetFromString("0.00 HYPHA")
	seedsLiquidTotal, _ := eos.NewAssetFromString("0.0000 SEEDS")
	seedsEscrowTotal, _ := eos.NewAssetFromString("0.0000 SEEDS")

	for index := range assignments {

		husdTotal = husdTotal.Add(assignments[index].HusdPerPhase)
		hyphaTotal = hyphaTotal.Add(assignments[index].HyphaPerPhase)
		hvoiceTotal = hvoiceTotal.Add(assignments[index].HvoicePerPhase)
		seedsLiquidTotal = seedsLiquidTotal.Add(assignments[index].SeedsLiquidPerPhase)
		seedsEscrowTotal = seedsEscrowTotal.Add(assignments[index].SeedsEscrowPerPhase)

		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: strconv.Itoa(int(assignments[index].ID))},
			{Align: simpletable.AlignRight, Text: string(assignments[index].Assigned)},
			{Align: simpletable.AlignLeft, Text: string(assignments[index].Role.Title)},
			{Align: simpletable.AlignRight, Text: FormatAsset(&assignments[index].Role.AnnualUSDSalary)},
			{Align: simpletable.AlignRight, Text: strconv.FormatFloat(assignments[index].TimeShare*100, 'f', -1, 64)},
			{Align: simpletable.AlignRight, Text: strconv.FormatFloat(assignments[index].DeferredPay*100, 'f', -1, 64)},
			{Align: simpletable.AlignRight, Text: strconv.FormatFloat(assignments[index].InstantHusdPerc*100, 'f', -1, 64)},
			{Align: simpletable.AlignRight, Text: FormatAsset(&assignments[index].HusdPerPhase)},
			{Align: simpletable.AlignRight, Text: FormatAsset(&assignments[index].HyphaPerPhase)},
			{Align: simpletable.AlignRight, Text: FormatAsset(&assignments[index].HvoicePerPhase)},
			{Align: simpletable.AlignRight, Text: FormatAsset(&assignments[index].SeedsEscrowPerPhase)},
			{Align: simpletable.AlignRight, Text: FormatAsset(&assignments[index].SeedsLiquidPerPhase)},
			{Align: simpletable.AlignRight, Text: assignments[index].StartPeriod.StartTime.Time.Format("2006 Jan 02")},
			{Align: simpletable.AlignRight, Text: assignments[index].EndPeriod.EndTime.Time.Format("2006 Jan 02")},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{},
			{},
			{}, {}, {}, {},
			{Align: simpletable.AlignRight, Text: "Subtotal"},
			{Align: simpletable.AlignRight, Text: FormatAsset(&husdTotal)},
			{Align: simpletable.AlignRight, Text: FormatAsset(&hyphaTotal)},
			{Align: simpletable.AlignRight, Text: FormatAsset(&hvoiceTotal)},
			{Align: simpletable.AlignRight, Text: FormatAsset(&seedsEscrowTotal)},
			{Align: simpletable.AlignRight, Text: FormatAsset(&seedsLiquidTotal)},
			{}, {},
		},
	}

	return table
}