package cmd

import (

	"mmctl/common"

	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize(common.InitConfig)
}


// FrameWork 定义 framework commend
var FrameWork = &cobra.Command{
	Use:   "framework",
	Short: "Show framework name list",
	Run: func(fwcmd *cobra.Command, args []string) {
		httpData := common.GetFwk()
		fwkList := common.FwkInfo(httpData)
		//showFwk(fwkList)
		common.Show(fwkList,4)
	},
}


