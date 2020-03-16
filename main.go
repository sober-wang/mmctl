package main

import (
	"fmt"

	mema "mmctl/cmd"

	"github.com/spf13/cobra"
)

func version() *cobra.Command {

	v := &cobra.Command{
		Use:   "version",
		Short: "show version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("v0.3")
		},
	}
	return v
}

func main() {
	//	cmd.Execute()
//	var fwk string
	mainCmd := &cobra.Command{
		Use:   "mmctl",
		Short: "Mesos Marathon ctl",
	}
	get := &cobra.Command{
			Use: "get",
			Short: "Get Mesos and Marathon information",
	}
	inspect := &cobra.Command{
			Use: "inspect",
			Short: "Inspect a appid details information",
	}
	mainCmd.AddCommand(get)
	mainCmd.AddCommand(inspect)
//	mema.AppID.Flags().StringVarP(&fwk,"framework","f","","Select a FrameWork show appid")
	get.AddCommand(
		version(),
		mema.AppID,
		mema.FrameWork,
		mema.Node,
	)
	
	if err := mainCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
