package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"mmctl/common"

	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	fwk    string
	allApp bool
	AppID  = &cobra.Command{
		Use:   "appid",
		Short: "Show Marathon appids",
		Run: func(cmd *cobra.Command, args []string) {
			run(args)
		},
	}
)

// init 初始化配置文件
func init() {
	cobra.OnInitialize(common.InitConfig)
	AppID.Flags().StringVarP(&fwk, "framework", "m", "", "Select a FrameWork show appid")
	AppID.Flags().BoolVarP(&allApp, "all-appid", "", false, "Show all framework all appid")
}

// getOneM 只获取 -m 参数中的 appid
func getOneM(murl string) {
	data := GetMarthonAppID(murl)
	tmp := matchAppID(data)
	common.Show(tmp, 0)
}

// showAllAppid 展示集群下所有 appid
func showAllAppid() {
	tmp := common.GetFwk()
	fwkList := common.FwkInfo(tmp)
	for _, v := range fwkList {
		if v[3] == "false" {
			continue
		}

		if ok, err := regexp.MatchString("^http.*", v[2]); err != nil && !ok {
			fmt.Println("Not")
		} else if ok {
			fmt.Printf("\nMarathon Name is [ %v ],Marathon URL is %v\n", v[0], v[2])
			url := fmt.Sprintf("%v/v2/apps", v[2])
			getOneM(url)
		}
	}
}

// run 运行命令
func run(args []string) {
	switch {
	// 展示指定的 appid
	case fwk != "" && len(args) != 0:
		fmt.Println(args[0])
		murl := fmt.Sprintf("%v/%v", makeMURL(fwk), args[0])
		oneData := GetMarthonAppID(murl)
		oneApp := matchOneID(oneData, args[0])
		common.Show(oneApp, 1)
		return
	// 展示指定 marathons 下的 appid
	case fwk != "":
		murl := makeMURL(fwk)
		getOneM(murl)
		return
	// 展示Mesos 上所有的 appid
	case allApp == true:
		showAllAppid()
		return
	default:
		fmt.Println(`
		You should give me a framework name
		1. mmctl get framework 
		2. mmctl get appid -m myfwk
		or 
		mmctl get appid --all-appid
		`)
		return

	}
}

// getMURL 构建 marathon URI
func makeMURL(fwkName string) string {
	data := common.GetFwk()
	for i := 0; i < jsoniter.Get(data, "frameworks").Size(); i++ {
		tmp := jsoniter.Get(data, "frameworks", i)
		if tmp.Get("name").ToString() == fwkName {
			return fmt.Sprintf("%v/v2/apps", tmp.Get("webui_url").ToString())
		}
	}
	return ""
}

// matchOneID 从一个 appid 中获取数据
func matchOneID(d []byte, appid string) [][]string {
	var al [][]string
	var h = []string{
		"APPID",
		"TASKID",
		"HOSTNAME",
		"STATE",
		"PROTOCOL",
		"HEALTH",
	}
	al = append(al, h)
	for i := 0; i < jsoniter.Get(d, "app", "tasks").Size(); i++ {
		tmp := jsoniter.Get(d, "app", "tasks", i)
		var a = []string{
			appid,
			tmp.Get("id").ToString(),
			tmp.Get("host").ToString(),
			tmp.Get("state").ToString(),
			tmp.Get("ipAddresses", 0).Get("protocol").ToString(),
			tmp.Get("healthCheckResults", 0).Get("alive").ToString(),
		}
		al = append(al, a)
	}
	return al
}

// matchAppID 从 Marathon RES API 接口中提取出 appid 列表
func matchAppID(b []byte) [][]string {
	var al [][]string
	head := []string{"NAME", "CPUs", "Mem", "Running", "Staged", "Healthy", "Unhealthy"}
	al = append(al, head)
	for i := 0; i < jsoniter.Get(b, "apps").Size(); i++ {
		tmp := jsoniter.Get(b, "apps", i)
		var a = []string{
			tmp.Get("id").ToString(),
			tmp.Get("cpus").ToString(),
			tmp.Get("mem").ToString(),
			tmp.Get("tasksRunning").ToString(),
			tmp.Get("tasksStaged").ToString(),
			tmp.Get("tasksHealthy").ToString(),
			tmp.Get("tasksUnhealthy").ToString(),
		}
		al = append(al, a)
	}
	return al
}

// getMarathonAppID 获取 Marathon appid
func GetMarthonAppID(murl string) []byte {
	username := viper.Get("mmctl.marathon.user").(string)
	passwd := viper.Get("mmctl.marathon.passwd").(string)

	c := &http.Client{}
	req, err := http.NewRequest("GET", murl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth(username, passwd)

	resp, err := c.Do(req)
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
