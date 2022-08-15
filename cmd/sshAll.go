package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.ibm.com/maxwell-ibm/goEnum/utilities"
	"github.ibm.com/maxwell-ibm/goEnum/utilities/color"
	"github.ibm.com/maxwell-ibm/goEnum/utilities/parameters"
	"github.ibm.com/maxwell-ibm/goEnum/utilities/remote"
)

var sshAll = &cobra.Command{
	Use:              "all",
	Short:            "run all available modules",
	Long:             "run all available modules",
	Args:             cobra.NoArgs,
	Run:              sshAllRun,
	PersistentPreRun: sshAllParameters,
}

func init() {
	ssh.AddCommand(sshAll)
}

func sshAllRun(cmd *cobra.Command, args []string) {
	if parameters.Verbose {
		fmt.Fprintln(os.Stderr, utilities.LocalParameters())
	}

	err := remote.SSH(remote.SSHExecuteGoEnum)
	if err != nil {
		color.Fprintln(color.Red, os.Stderr, err)
	}
}
func sshAllParameters(cmd *cobra.Command, args []string) {
	sshParameters(cmd, args)
	parameters.Commands = append(parameters.Commands, "all")

}
