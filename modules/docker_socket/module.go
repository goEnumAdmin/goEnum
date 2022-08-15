package docker_socket

import (
	"github.ibm.com/maxwell-ibm/goEnum/structs"
)

var Module *structs.Module

func init() {
	Module = structs.NewModule(
		"Container with Docker Socket",
		Prereqs,
		Enumeration,
		Report,
	)
}
