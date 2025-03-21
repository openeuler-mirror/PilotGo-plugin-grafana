package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"os"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/grafana-plugin/conf"
)

const Version = "1.0.1"

var PluginInfo = &client.PluginInfo{
	Name:        "grafana",
	Version:     Version,
	Description: "grafana可视化工具支持",
	Author:      "guozhengxin",
	Email:       "guozhengxin@kylinos.cn",
	Url:         "",
	PluginType:  "",
	ReverseDest: "",
}

func main() {
	fmt.Println("hello grafana")

	InitLogger()

	PluginInfo.Url = "http://" + conf.Config().Http.Addr
	PluginInfo.ReverseDest = "http://" + conf.Config().Grafana.Addr
	PluginInfo.PluginType = "iframe"

	server := gin.Default()

	client := client.DefaultClient(PluginInfo)
	client.RegisterHandlers(server)
	InitRouter(server)

	// 添加扩展点
	var ex []common.Extention
	me1 := &common.PageExtention{
		Type:       common.ExtentionPage,
		Name:       "plugin-grafana",
		URL:        "/plugin/grafana",
		Permission: "plugin.grafana.page/menu",
	}
	ex = append(ex, me1)
	client.RegisterExtention(ex)

	if err := server.Run(conf.Config().Http.Addr); err != nil {
		logger.Fatal("failed to run server")
	}
}

func InitLogger() {
	if err := logger.Init(conf.Config().Logopts); err != nil {
		fmt.Printf("logger init failed, please check the config file: %s", err)
		os.Exit(-1)
	}
}

func InitRouter(router *gin.Engine) {
	// 所有grafana请求反向代理到grafana服务器
	pg := router.Group("/plugin/" + PluginInfo.Name)
	{
		pg.Any("/*any", func(c *gin.Context) {
			c.Set("__internal__reverse_dest", PluginInfo.ReverseDest)
			ReverseProxyHandler(c)
		})
	}
}

func ReverseProxyHandler(c *gin.Context) {
	remote := c.GetString("__internal__reverse_dest")
	if remote == "" {
		fmt.Println("get reverse dest failed!")
		return
	}

	target, err := url.Parse(remote)
	if err != nil {
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	// proxy.ModifyResponse = func(r *http.Response) error {
	// 	if setCookie := r.Header.Get("Set-Cookie"); setCookie != "" {
	// 		r.Header.Set("Set-Cookie", setCookie+" ; Secure ")
	// 	}
	// 	return nil
	// }

	proxy.ServeHTTP(c.Writer, c.Request)
}
