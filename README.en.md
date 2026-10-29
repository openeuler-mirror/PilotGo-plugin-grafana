English | [简体中文](./README.md)

# PilotGo-plugin-grafana

#### Introduction
PilotGo grafana plugin maintains grafana to provide beautiful visual monitoring interface.

#### Usage Instructions
Please make sure that the pilotgo-server application service is already running properly before using the grafana plugin. If it is not installed, deploy it by following the instructions at https://atomgit.com/src-openeuler/PilotGo .


#### Development and Installation Tutorial

1.  Install grafana;
2.  Modify the configuration file:
   >vim /etc/grafana/grafana.ini
   >root_url = http://real_ip:9999/plugin/grafana
   >serve_from_sub_path = true
   >allow_embedding = true
3.  Restart the grafana service:
   >systemctl restart grafana
4.  Clone the PilotGo-plugin-grafana repository and modify the configuration file address:
   >git clone https://atomgit.com/openeuler/PilotGo-plugin-grafana.git

   >vim /opt/PilotGo/plugin/grafana/config.yaml

   http_server:addr is the address where PilotGo-plugin-grafana is installed;

   grafana_server:addr is the address where grafana is installed;
   >cd PilotGo-plugin-grafana
   >go run main.go
5.  Register the grafana plugin service with the PilotGo platform: [Plugin Management] -> [Add Plugin] -> plugin name: grafana; plugin address: http://ip:port/plugin/grafana (9999 by default) -> [Enable Plugin]. Once enabled successfully, you can view the grafana plugin dashboard on the PilotGo platform; if it fails to start, check the plugin error in /opt/PilotGo/plugin/grafana/log.

#### Usage Instructions

1.  xxxx
2.  xxxx
3.  xxxx

#### Supplementary Links
1.  [PilotGo User Manual](https://atomgit.com/openeuler/PilotGo/tree/master/docs/zh/usage_instructions.md)
2.  PilotGo [Code Repository](https://atomgit.com/openeuler/PilotGo)
3.  PilotGo [Package Repository](https://atomgit.com/src-openeuler/PilotGo)
4.  PilotGo-plugin-grafana [Package Repository](https://atomgit.com/src-openeuler/PilotGo-plugin-grafana)

#### Contributing

1.  Fork this repository
2.  Create a Feat_xxx branch
3.  Commit your code
4.  Create a Pull Request


