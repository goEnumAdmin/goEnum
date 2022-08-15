package modules

import (
	"github.ibm.com/maxwell-ibm/goEnum/modules/CVE_2021_3156"
	"github.ibm.com/maxwell-ibm/goEnum/modules/container_block_devices"
	"github.ibm.com/maxwell-ibm/goEnum/modules/cronjobs"
	"github.ibm.com/maxwell-ibm/goEnum/modules/docker_socket"
	"github.ibm.com/maxwell-ibm/goEnum/modules/privileged_container"
	"github.ibm.com/maxwell-ibm/goEnum/modules/protected_files"
	"github.ibm.com/maxwell-ibm/goEnum/modules/readable_files"
	"github.ibm.com/maxwell-ibm/goEnum/modules/services"
	"github.ibm.com/maxwell-ibm/goEnum/modules/special_permissions"
	"github.ibm.com/maxwell-ibm/goEnum/modules/unquoted_service_path"
	"github.ibm.com/maxwell-ibm/goEnum/modules/writable_files"
	"github.ibm.com/maxwell-ibm/goEnum/structs"
)

var (
	Modules map[string]*structs.Module
	Padding int
)

func init() {
	Modules = make(map[string]*structs.Module)

	Modules["protected-files"] = protected_files.Module
	Modules["cve-2021-3156"] = CVE_2021_3156.Module
	Modules["cronjobs"] = cronjobs.Module
	Modules["readable-files"] = readable_files.Module
	Modules["writable-files"] = writable_files.Module
	Modules["special-perms"] = special_permissions.Module
	Modules["priv-container"] = privileged_container.Module
	Modules["docker-sock"] = docker_socket.Module
	Modules["block-devices"] = container_block_devices.Module
	Modules["services"] = services.Module
	Modules["unquoted-service-path"] = unquoted_service_path.Module

	padding()
}

func padding() {
	for key := range Modules {
		if len(key) > Padding {
			Padding = len(key)
		}
	}
}
