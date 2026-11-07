# 私有构建说明（wbteve 自用）

本仓库是 ehang-io/nps 的私有分支，只在我们自己的机器上使用，不对外分发。

## 与原版的差异

1. 去掉 `lib/common/util.go` 中向 `myexternalip.com` 查询公网 IP 的外联，改由配置项 `external_ip` 提供，未配置则为空。
2. 禁用 `lib/install/install.go` 的在线升级：原版会访问 `api.github.com` 与 `ehang.io` 拉取升级包，现在直接拒绝并提示用内部构建产物替换。
3. 服务端新增高危端口默认拦截：`allow_ports` 为空时，拒绝 22、23、135、137、138、139、445、1433、3306、3389、5432、5900、6379、9200、11211、27017 的隧道映射。可用配置项 `deny_dangerous_port = false` 关闭。
4. 版本号改为 `0.26.10-private.1`，便于日志识别。

## 构建

GitHub Actions 会自动编译 linux_amd64、linux_arm64、linux_armv7、windows_amd64、darwin_amd64、darwin_arm64 六个平台的服务端与客户端，产物为 tar.gz。

打 tag 形如 `v0.26.10-private.1` 会自动创建 Release 并附上全部产物。

本地构建注意：`cmd/npc/sdk.go` 与 `npc.go` 都有 main 函数，编译普通二进制前需要先把 sdk.go 移出目录。

## 部署约定

- 替换前先备份旧二进制（例如 `nps.bak-YYYYMMDD`）。
- 服务端与客户端尽量使用同一版本，协议核心版本 0.26.0 兼容。
