package protected_files

import (
	"github.ibm.com/maxwell-ibm/goEnum/utilities/systemInfo"
)

func Prereqs() bool {
	return systemInfo.OS != "windows"
}
