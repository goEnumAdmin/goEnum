package unquoted_service_path

import (
	"github.ibm.com/maxwell-ibm/goEnum/structs"
)

var Module *structs.Module

func init() {
	Module = structs.NewModule(
		"Unquoted Service Paths",
		Prereqs,
		Enumeration,
		Report,
	)
}
