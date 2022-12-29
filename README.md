# **Go**mi

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

## Install

### Requirements

 * Linux OS with x86-64 / arm64
 * Git
 * Working ssh key and config for commits / ssh config with base auth

### SSH-KEY Git

**TODO:** add example how to configure whole env to be able run gomi from cron

### Baseauth Git

**TODO:** add example how to configure whole env to be able run gomi from cron


