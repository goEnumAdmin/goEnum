package cronjobs

import (
	"github.ibm.com/maxwell-ibm/goEnum/structs"
)

var Module *structs.Module

func init() {
	Module = structs.NewModule(
		"Cronjobs with Writable Executable",
		Prereqs,
		Enumeration,
		Report,
	)
}
