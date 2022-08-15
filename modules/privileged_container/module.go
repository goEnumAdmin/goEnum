package privileged_container

import (
	"github.ibm.com/maxwell-ibm/goEnum/structs"
)

var Module *structs.Module

func init() {
	Module = structs.NewModule(
		"Priviledged Container",
		Prereqs,
		Enumeration,
		Report,
	)
}
