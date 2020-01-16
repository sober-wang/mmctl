package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"

	"mmctl/types"

	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(types.ConfPath)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	} else {
		viper.ConfigFileUsed()
	}
}

// fwkName 获取 framework name 列表
func FwkInfo(d []byte) [][]string {
	//	var fwkList []string
	var h = []string{
		"Name",
		"HostName",
		"Webui_url",
		"Active",
	}
	//var fwkList []types.FrameWork
	var fwkList [][]string
	fwkList = append(fwkList, h)
	for i := 0; i < jsoniter.Get(d, "frameworks").Size(); i++ {
		/*
			var fwk types.FrameWork
			fwk.Name = jsoniter.Get(d, "frameworks", i).Get("name").ToString()
			fwk.HostName = jsoniter.Get(d, "frameworks", i).Get("hostname").ToString()
			fwk.WebeUi = jsoniter.Get(d, "frameworks", i).Get("webui_url").ToString()
			fwk.Active = jsoniter.Get(d, "frameworks", i).Get("active").ToString()
		*/
		var fwk = []string{
			jsoniter.Get(d, "frameworks", i).Get("name").ToString(),
			jsoniter.Get(d, "frameworks", i).Get("hostname").ToString(),
			jsoniter.Get(d, "frameworks", i).Get("webui_url").ToString(),
			jsoniter.Get(d, "frameworks", i).Get("active").ToString(),
		}

		fwkList = append(fwkList, fwk)
	}
	return fwkList
}

// GetFwk 获取 framework name
func GetFwk() []byte {
	url := fmt.Sprintf("http://%v:%v/state-summary",
		viper.Get("mmctl.mesos.ip"),
		viper.Get("mmctl.mesos.port"))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return d
}

// Show 展示数据
func Show(data [][]string, adjust int) {
	//	fmt.Println(data)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, adjust, '\t', 0)
	for _, v := range data {
		fmt.Fprintf(w, "%v\n", strings.Join(v, "\t"))
	}
	w.Flush()

}
