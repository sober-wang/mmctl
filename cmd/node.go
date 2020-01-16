package cmd

import (
	//		"fmt"
	//	"io/ioutil"
	//	"net/http"
	//	"os"
	//	"strings"
	//	"text/tabwriter"

	"mmctl/common"

	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
	//	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(common.InitConfig)
}

var Node = &cobra.Command{
	Use:   "node",
	Short: "Show node information",
	Run: func(c *cobra.Command, args []string) {
		data := common.GetFwk()
		nl := GetNode(data)
		common.Show(nl, 1)

	},
}

// GetNode 获取 node 信息
func GetNode(data []byte) [][]string {
	var nodeList [][]string
	var h = []string{
		"hostname",
		"ip",
		"allCPUs",
		"allMem",
		"allDisk",
		"usedCPUs",
		"usedMem",
		"usedDisk",
		"attributes",
		"active",
	}
	nodeList = append(nodeList, h)
	for i := 0; i < jsoniter.Get(data, "slaves").Size(); i++ {
		tmp := jsoniter.Get(data, "slaves", i)
		var b = []string{
			tmp.Get("hostname").ToString(),
			tmp.Get("hostname").ToString(),
			tmp.Get("resources").Get("cpus").ToString(),
			tmp.Get("resources").Get("mem").ToString(),
			tmp.Get("resources").Get("disk").ToString(),
			tmp.Get("used_resources").Get("cpus").ToString(),
			tmp.Get("used_resources").Get("mem").ToString(),
			tmp.Get("used_resources").Get("disk").ToString(),
			tmp.Get("attributes").ToString(),
			tmp.Get("active").ToString(),
		}
		nodeList = append(nodeList, b)
		//	fmt.Println(nodeList)
	}
	return nodeList

}
