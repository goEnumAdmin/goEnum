package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.ibm.com/maxwell-ibm/goEnum/utilities"
	"github.ibm.com/maxwell-ibm/goEnum/utilities/parameters"
)

var sshModules = &cobra.Command{
	Use:              "modules",
	Short:            "display all available modules",
	Long:             "display all available modules",
	Args:             cobra.NoArgs,
	Run:              sshModulesRun,
	PersistentPreRun: sshModulesParameters,
}

func init() {
	ssh.AddCommand(sshModules)
}

func sshModulesRun(cmd *cobra.Command, args []string) {
	if parameters.Verbose {
		fmt.Fprintln(os.Stderr, utilities.LocalParameters())
	}

	printModules()
}
func sshModulesParameters(cmd *cobra.Command, args []string) {
	sshParameters(cmd, args)
	parameters.Commands = append(parameters.Commands, "modules")

}
