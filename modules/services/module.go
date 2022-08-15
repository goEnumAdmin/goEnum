package services

import (
	"github.ibm.com/maxwell-ibm/goEnum/structs"
)

var Module *structs.Module

func init() {
	Module = structs.NewModule(
		"Insecure Services and Utilized Binaries",
		Prereqs,
		Enumeration,
		Report,
	)
}
