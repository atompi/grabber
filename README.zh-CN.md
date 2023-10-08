# Grabber

Grabber 是一个简单的多 web 文件并行下载工具，可用于 web 文件批量下载、多配置文件批量同步等场景。

## 用法

1. 创建一个配置文件，如 ./grabber.yaml，该配置文件描述了需要下载的 web 文件列表。

- ./grabber.yaml

```
threads: 5    # 下载并发线程数
sources:    # 下载源列表，可以包含多个源
- url: https://git.atompi.com/atompi/conf/raw/branch/main/nginx    # 下载源地址，该地址为需要下载的文件的根路径，此次以 gitea 代码仓库为下载源
  auth: 123123    # access token，可选，用于访问 gitea 仓库
  files:    # 需要下载的文件列表，可以包含多个文件
  - src: nginx.conf    # 下载源地址下需要下载的文件，该地址为文件相对路径
    dest: /tmp/nginx/nginx.conf    # 下载后保存的文件路径
  - src: proxy.conf
    dest: /tmp/nginx/proxy.conf
```

2. 运行 Grabber。

```
./grabber -c ./grabber.yaml
```
