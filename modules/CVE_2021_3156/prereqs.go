package CVE_2021_3156

import (
	"github.ibm.com/maxwell-ibm/goEnum/utilities/systemInfo"
)

func Prereqs() bool {
	return systemInfo.OS != "windows"
}
