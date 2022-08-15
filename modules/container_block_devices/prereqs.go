package container_block_devices

import (
	"github.ibm.com/maxwell-ibm/goEnum/utilities/systemInfo"
)

func Prereqs() bool {
	return systemInfo.Container
}
