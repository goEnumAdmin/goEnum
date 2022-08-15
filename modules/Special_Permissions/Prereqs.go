package special_permissions

import (
	"github.ibm.com/maxwell-ibm/goEnum/utilities/systemInfo"
)

func Prereqs() bool {
	return systemInfo.OS == "darwin" || systemInfo.OS == "linux"
}
