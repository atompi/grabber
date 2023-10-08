# Grabber

Grabber is a simple multi-web file parallel download tool, can be used for web file batch download, multi-configuration file batch synchronization and other scenarios.

## Usage

1. Create a project configuration file.

- ./grabber.yaml

```
threads: 5    # download work threads
sources:    # multiple configuration projects sources
- url: https://git.atompi.com/atompi/conf/raw/branch/main/nginx    # configuration project root url
  auth: 123123    # access token (gitea)
  files:    # files in current project to download
  - src: nginx.conf    # file path based on project root
    dest: /tmp/nginx/nginx.conf    # destination path
  - src: proxy.conf
    dest: /tmp/nginx/proxy.conf
```

2. Run the program.

```
./grabber -c ./grabber.yaml
```
