# Yaml

```sh
cd yaml # run all commands from here

pulumi logout
pulumi login file:///$PWD

pulumi new kubernetes-yaml --force

pulumi up

kubectl -n default get deployment
```
