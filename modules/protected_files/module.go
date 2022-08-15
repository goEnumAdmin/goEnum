package protected_files

import (
	"github.ibm.com/maxwell-ibm/goEnum/structs"
)

var Module *structs.Module

func init() {
	Module = structs.NewModule(
		"Protected Files",
		Prereqs,
		Enumeration,
		Report,
	)
}
