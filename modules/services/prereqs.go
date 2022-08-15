package services

import (
	"github.ibm.com/maxwell-ibm/goEnum/utilities/systemInfo"
)

func Prereqs() bool {
	return systemInfo.OS == "linux"
}
