package main

import (
	"fmt"

	"github.com/eflowcr/sap-eintegra/database"
	logiclayer "github.com/eflowcr/sap-eintegra/logic_layer"
)

func main() {

	database.DSN = fmt.Sprintf("sqlserver://%s:%s@%s:%v?database=%s",
		"usreprac",
		"usreprac",
		"20.118.233.183",
		1433,
		"EINTEGRA_SL")

	//cll := logiclayer.SapCompanyLL{}

	// companies, err := cll.GetCompanies()

	// if err != nil {
	// 	panic(err)
	// }

	// for _, company := range *companies {
	// 	println(company.CompanyDB)
	// }

	// company, err := cll.GetCompany("SB1LD_DTL_PRD")

	// if err != nil {
	// 	panic(err)
	// }

	// println(company.CompanyDB)

	// }

	// company, err := cll.GetCompany("SB1LD_DTL_PRD")

	// if err != nil {
	// 	panic(err)
	// }

	// cll := logiclayer.SapIncommingInterfacesLL{}

	// a, err := cll.GetIncommingInterfacesByProcess("EXPERP")

	// if err != nil {
	// 	panic(err)
	// }

	// for _, b := range *a {
	// 	println(b.InterfaceId)
	// }

	cll := logiclayer.SapIncommingInterfacesLL{}

	a, err := cll.GetQueriesByInterface(2)

	if err != nil {
		panic(err)
	}

	for _, b := range *a {
		println(b.QueryDescription)
	}

}
