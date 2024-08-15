# alinyun-ddns

# Get-Started

设置配置文件即可启动

```toml

access_key_id = ""          # 阿里云AccessKeyID
access_key_secret = ""      # 阿里云AccessKeySecret
region_id = "cn-beijing"    # 阿里云region_id

domain_name = "xiaoyusb.fun"    # 域名 
RR = "local"                    # 子域名/主机名
record_type = "AAAA"            # 记录类型 (A, AAAA, CNAME, TXT, MX, etc.) 支持AAA(ipv6) A(ipv4)

ddns_interval = "30s"           # 本机ip检查并更新DNS记录间隔

```

# Appendix

以下内容来自豆包：

**DDNS** 即动态域名系统（Dynamic Domain Name System）。

简单来说，它是一种将动态分配的 IP 地址与固定的域名关联起来的技术。

在网络中，很多情况下用户接入互联网时所获得的 IP 地址是动态变化的，比如家庭宽带用户每次重新连接网络时，可能被分配到不同的 IP 地址。

而 DDNS 可以解决这个问题。它能让用户通过一个固定不变的域名，始终能够访问到具有动态 IP 地址的网络设备或服务器。

例如，您在家中搭建了一个私人服务器用于远程访问，但您的家庭网络 IP 地址会变化。通过使用 DDNS 服务，您注册一个域名（如：myhome.server.com），即使您家的 IP 地址改变了，您依然可以通过这个固定的域名来访问您的私人服务器。

再比如，一些小型企业使用网络摄像头进行监控，通过 DDNS 可以确保在 IP 地址变动的情况下，仍然能够通过固定域名远程查看监控画面。