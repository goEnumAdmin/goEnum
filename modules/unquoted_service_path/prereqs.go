package unquoted_service_path

import (
	"github.ibm.com/maxwell-ibm/goEnum/utilities/systemInfo"
)

func Prereqs() bool {
	return systemInfo.OS == "windows"
}
