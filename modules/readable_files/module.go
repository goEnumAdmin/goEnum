package readable_files

import (
	"github.ibm.com/maxwell-ibm/goEnum/structs"
)

var Module *structs.Module

func init() {
	Module = structs.NewModule(
		"Mispermissioned Files (readable)",
		Prereqs,
		Enumeration,
		Report,
	)
}
