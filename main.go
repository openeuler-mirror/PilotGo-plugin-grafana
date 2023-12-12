package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

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
	PluginType:  "iframe",
	ReverseDest: "",
}

func main() {
	fmt.Println("hello grafana")

	InitLogger()

	PluginInfo.Url = "http://" + conf.Config().Http.Addr + "/plugin/grafana"
	PluginInfo.ReverseDest = "http://" + conf.Config().Grafana.Addr

	server := gin.Default()

	client := client.DefaultClient(PluginInfo)
	client.RegisterHandlers(server)
	InitRouter(server)

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
	c.Request.URL.Path = strings.Replace(c.Request.URL.Path, "/plugin/grafana", "", 1) //请求API

	// proxy.ModifyResponse = func(r *http.Response) error {
	// 	if setCookie := r.Header.Get("Set-Cookie"); setCookie != "" {
	// 		r.Header.Set("Set-Cookie", setCookie+" ; Secure ")
	// 	}
	// 	return nil
	// }

	proxy.ServeHTTP(c.Writer, c.Request)
}
