# lsky-upload
Picture bed upload program of Lsky（Lsky 图床的上传程序）

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/YGXB-net/lsky-upload)](./go.mod) [![GitHub release (latest by date)](https://img.shields.io/github/v/release/YGXB-net/lsky-upload)](../../releases/latest) [![GitHub](https://img.shields.io/github/license/YGXB-net/lsky-upload)](./LICENSE)

[文档](./docs) • [下载](../../releases/latest) • [开始使用](#下载--使用方法)

> 该项目为本人在使用 [PicGo](https://github.com/Molunerfinn/PicGo) 时，觉得 PicGo 虽然功能强大，但过于复杂
>
> 所以就想自己开发一个简单，可以满足自己需求的程序，并且开源出来
>
> 于是就有了此项目 [lsky-upload](https://github.com/YGXB-net/lsky-upload)

------

## 应用概述

**lsky-upload：一个 Typora 的图片上传工具，它可以自动帮你完成图片上传的工作**

**支持的图床为：[兰空图床(Lsky Pro)](https://github.com/lsky-org/lsky-pro)（版本: v2.x）**

lsky-upload 使用 **golang** 语言开发，程序高效，快速，简洁

## 功能

- 上传本地图片到图床
- 获取网络图片并上传到图床
- 与 Typora 结合，实现插入图片自动上传，优化写作流程

## 下载 & 使用方法

1. 前往 [Release](../../releases) 界面，下载最新版本

2. 解压下载的压缩包，得到 `lsky-upload.exe` 文件

3. `lsky-upload.exe` 文件放到一个空文件夹中（例：`C:\ProgramFiles\lsky-upload`）

4. 运行 `lsky-upload.exe` ，在文件夹中会生成 `config.yml` 文件，修改配置（详细信息请参考[相关文档](./docs/README.md)）

5. 修改 Typora 配置，在 "命令" 中填入 `lsky-upload.exe` 的位置（例：`"C:\ProgramFiles\lsky-upload\lsky-upload.exe"`）

   **注意：填入命令时一定要在路径两边加上 `"` 否则可能会因为路径有空格而无法运行**
[![看不到截图点我](https://img.ygxb.net/i/2023/05/14/6460666def259.png)](https://img.ygxb.net/i/2023/05/14/6460666def259.png)

## :hammer: 构建

> 使用 Golang 版本：1.21

克隆本仓库到本地，然后执行：

```shell
go build
```

## :email: 联系我

- E-mail: [me@ygxb.net](mailto:me@ygxb.net)
- QQ: 3040809965

## 其他

本项目的文档和教程还未完善，欢迎大家提出建议和问题到 [Issues](https://github.com/YGXB-net/lsky-upload/issues)

如果不会使用或使用时遇到程序错误，可以加我的 QQ: 3040809965，我会一一解答问题
