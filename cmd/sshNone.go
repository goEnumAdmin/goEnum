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

var sshNone = &cobra.Command{
	Use:              "none",
	Short:            "runs no modules",
	Long:             "runs no modules. Gives quick fingerprinting output for system",
	Args:             cobra.NoArgs,
	Run:              sshNoneRun,
	PersistentPreRun: sshNoneParameters,
}

func init() {
	ssh.AddCommand(sshNone)
}

func sshNoneRun(cmd *cobra.Command, args []string) {
	if parameters.Verbose {
		fmt.Fprintln(os.Stderr, utilities.LocalParameters())
	}

	err := remote.SSH(remote.SSHExecuteGoEnum)
	if err != nil {
		color.Fprintln(color.Red, os.Stderr, err)
	}
}
func sshNoneParameters(cmd *cobra.Command, args []string) {
	sshParameters(cmd, args)
	parameters.Commands = append(parameters.Commands, "none")

}
