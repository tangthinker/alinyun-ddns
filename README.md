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