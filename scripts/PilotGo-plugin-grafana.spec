%define         debug_package %{nil}

Name:           PilotGo-plugin-grafana
Version:        1.0.1
Release:        1
Summary:        PilotGo grafana plugin maintains grafana to provide beautiful visual monitoring interface.
License:        MulanPSL-2.0
URL:            https://gitee.com/openeuler/PilotGo-plugin-grafana
Source0:        https://gitee.com/openeuler/PilotGo-plugin-grafana/%{name}-%{version}.tar.gz
BuildRequires:  systemd
BuildRequires:  golang
Requires:       grafana
Provides:       pilotgo-plugin-grafana = %{version}-%{release}

%description
PilotGo grafana plugin maintains grafana to provide beautiful visual monitoring interface.


%prep
%autosetup -p1 -n %{name}-%{version}

%build
GO111MODULE=on go build -mod=vendor -o PilotGo-plugin-grafana main.go

%install
mkdir -p %{buildroot}/opt/PilotGo/plugin/grafana/log
install -D -m 0755 PilotGo-plugin-grafana %{buildroot}/opt/PilotGo/plugin/grafana
install -D -m 0644 config.yaml.templete %{buildroot}/opt/PilotGo/plugin/grafana/config.yaml
install -D -m 0644 scripts/PilotGo-plugin-grafana.service %{buildroot}%{_unitdir}/PilotGo-plugin-grafana.service

%post
%systemd_post PilotGo-plugin-grafana.service

%preun
%systemd_preun PilotGo-plugin-grafana.service

%postun
%systemd_postun PilotGo-plugin-grafana.service

%files
%dir /opt/PilotGo
%dir /opt/PilotGo/plugin
%dir /opt/PilotGo/plugin/grafana
%dir /opt/PilotGo/plugin/grafana/log
/opt/PilotGo/plugin/grafana/PilotGo-plugin-grafana
/opt/PilotGo/plugin/grafana/config.yaml
%{_unitdir}/PilotGo-plugin-grafana.service


%changelog
* Fri Sep 01 2023 jianxinyu <jiangxinyu@kylinos.cn> - 1.0.1-1
- Package init

