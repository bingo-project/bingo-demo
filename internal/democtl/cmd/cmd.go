package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"

	"demo/internal/apiserver"
	"demo/internal/democtl/cmd/user"
	"demo/internal/democtl/cmd/version"
	"demo/internal/democtl/util/templates"
	"demo/internal/pkg/log"
	"demo/pkg/cli/genericclioptions"
)

func NewDefaultDemoCtlCommand() *cobra.Command {
	return NewDemoCtlCommand(os.Stdin, os.Stdout, os.Stderr)
}

func NewDemoCtlCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	var cmds = &cobra.Command{
		Use:   "democtl",
		Short: "democtl is the demo startup client",
		Long:  `democtl is the client side tool for demo startup.`,
		Run:   runHelp,
	}

	// Load config
	cobra.OnInitialize(initConfig)

	ioStreams := genericclioptions.IOStreams{In: in, Out: out, ErrOut: err}

	groups := templates.CommandGroups{
		{
			Message:  "Basic Commands:",
			Commands: []*cobra.Command{},
		},
		{
			Message: "Advanced Commands:",
			Commands: []*cobra.Command{
				user.NewCmdUser(ioStreams),
			},
		},
	}
	groups.Add(cmds)

	filters := []string{""}
	templates.ActsAsRootCommand(cmds, filters, groups...)

	// Config file
	cmds.PersistentFlags().StringVarP(&apiserver.CfgFile, "config", "c", "", "The path to the configuration file. Empty string for no configuration file.")

	// Add commands
	cmds.AddCommand(version.NewCmdVersion(ioStreams))

	return cmds
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	apiserver.InitConfig()

	// Init store
	if err := apiserver.InitStore(); err != nil {
		log.Fatalw(err.Error())
	}
}

func runHelp(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}
