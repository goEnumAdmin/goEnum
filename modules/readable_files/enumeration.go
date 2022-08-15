//go:build !windows

package readable_files

import (
	"strings"

	"github.ibm.com/maxwell-ibm/goEnum/utilities"
	"github.ibm.com/maxwell-ibm/goEnum/utilities/permissions"
	"github.ibm.com/maxwell-ibm/goEnum/utilities/systemInfo"
)

func Enumeration() ([]string, bool) {
	var files []string

	checkWritable("/home/", &files)
	checkWritable("/Users/", &files)

	return files, len(files) != 0
}

func checkWritable(path string, files *[]string) {
	for path := range utilities.IterateOverDirN(path, 10) {
		if systemInfo.OS == "darwin" {
			if strings.Contains(path, ".localized") {
				continue
			} else if strings.Contains(strings.ToLower(path), "plist") {
				continue
			} else if strings.HasSuffix(path, ".pdf") {
				continue
			}
		}
		if strings.HasSuffix(path, "themes") {
			continue
		}
		if permissions.RW(path) {
			*files = append(*files, path)
		}
	}
}
