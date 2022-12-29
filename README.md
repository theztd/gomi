# **Go**mi

[![Release](https://github.com/theztd/gomi/actions/workflows/release.yml/badge.svg)](https://github.com/theztd/gomi/actions/workflows/release.yml)

Lightweight GOlang utility for auto-commiting uncommited changes in given repositories with reporting in node_exporter (tex-file) format

## Use case

Maybe you are somewhere on the way to GitOps or IaaC, but world is not perfect. You still need to keep some configuration in the git, but there is someone/something changing files directly on the server.


## Use

To see actual list of flags use flag **-h** 
```bash
Usage of gomi:
  -check
        Only show changes but don't do anything.
  -path string
        Absolute path to the git repo root, default is curent dir.
  -prom-path string
        Absolute path to node_exporter tex-file path, when is not defined, nothing will be generated.
```

### Example

Gomi wil check the dir **/etc/nginx/sites-enabled/** and report state to file **/tmp/metrics/gomi_nginx-sites.txt**
```bash
gomi -path /etc/nginx/sites-enabled/ -prom-path /tmp/metrics/gomi_nginx-sites.txt
```

**Example how the commit looks like**

```bash
commit 69ce2a794ba7b5f152e60ffa24e71d65cd4f0950
Author: Automat <automat@gin05>
Date:   Thu Dec 29 12:24:08 2022 +0000

----- Detail -----
 M README.md
?? Makefile
?? git/
?? go.mod
?? main.go
?? node_exporter.go
```

## Install

### Requirements

 * Linux OS with x86-64 / arm64
 * Git
 * Working ssh key and config for commits / ssh config with base auth

### SSH-KEY Git

**TODO:** add example how to configure whole env to be able run gomi from cron


**Don't forget to set required privileges in gitlab / github /.. .**

### Baseauth Git

**.git/config**
```toml
[core]
	repositoryformatversion = 0
	filemode = true
	bare = false
	logallrefupdates = true
[remote "origin"]
	url = https://automat:paSSSSSSSSwOOOOOOOOOrd@gitlab.com/theztd/example.git
	fetch = +refs/heads/*:refs/remotes/origin/*
[branch "master"]
	remote = origin
	merge = refs/heads/master
```

**Don't forget to set required privileges in gitlab / github /.. .**

Cron job example

```cron
50  *  *  *  *       /usr/local/bin/gomi -path /opt/project_dir -prom-path /var/metrics/gomi-project_dir.prom &>> /var/log/gomi.log
```