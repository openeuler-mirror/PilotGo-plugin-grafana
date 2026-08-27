# PilotGo-plugin-grafana

#### 介绍
PilotGo grafana plugin maintains grafana to provide beautiful visual monitoring interface.

#### 使用说明
请确保使用grafana插件之前已经正常运行pilotgo-server应用服务，若未安装，请通过链接 https://atomgit.com/src-openeuler/PilotGo ，按照使用说明进行部署。


#### 开发安装教程

1.  安装grafana；
2.  修改配置文件：
   >vim /etc/grafana/grafana.ini
   >root_url = http://真实ip:9999/plugin/grafana
   >serve_from_sub_path = true
   >allow_embedding = true
3.  重启grafana服务：
   >systemctl restart grafana
4.  克隆PilotGo-plugin-grafana仓库，修改配置文件地址：
   >git clone https://atomgit.com/openeuler/PilotGo-plugin-grafana.git

   >vim /opt/PilotGo/plugin/grafana/config.yaml

   http_server:addr为安装PilotGo-plugin-grafana地址；

   grafana_server:addr为安装grafana地址；
   >cd PilotGo-plugin-grafana
   >go run main.go
5.  将grafana插件服务注册到PilotGo平台 [插件管理] -> [添加插件] -> 插件名称：grafana；插件地址：http://ip:port/plugin/grafana（默认9999） -> [启用插件] 启用成功，即可在PilotGo平台查看grafana插件看板；若启动失败，通过 /opt/PilotGo/plugin/grafana/log 查看插件报错原因。

#### 使用说明

1.  xxxx
2.  xxxx
3.  xxxx

#### 补充链接
1.  [PilotGo使用手册](https://atomgit.com/openeuler/docs/tree/master/docs/zh/docs/PilotGo/使用手册.md)
2.  PilotGo[代码仓](https://atomgit.com/openeuler/PilotGo)
3.  PilotGo[软件包仓](https://atomgit.com/src-openeuler/PilotGo)
4.  PilotGo-plugin-grafana[软件包仓](https://atomgit.com/src-openeuler/PilotGo-plugin-grafana)

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


