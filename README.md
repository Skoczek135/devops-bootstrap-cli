**D**ev**O**ps**C****L****I** is a project written in golang to enable bootstrapping daily-usage things in DevOps world.
It's personal project to simplify my daily work.

## Installation
```bash
go get -u github.com/skoczewd/devopscli
```

## Usage
```bash
docli --help
```
## Pre-requisites
1. fzf
2. kubectl node-shell 

## Features
### Kubernetes
    [x] Deploying debug pod
    [x] SSHing into K8S node via node-shell 
    [x] Bash Exec into the pod using not full name but its prefix

### Terraform
    [x] Adding terraform providers (aws, helm, kubernetes)
    [x] Adding helm release and default values.yaml is being saved in the same directory
    [x] Adding terraform module which pulls the module from the source and put default values along with a description.
    [x] Bootstraping terraform workspace
    [x] Bootstraping terraform module

