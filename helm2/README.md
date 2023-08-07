# helm2

Exercises for learning how to use pulumi to deploy a helm chart

```sh
# Run all commands from this folder
cd helm2

# To use Pulumi without the Pulumi Cloud, log in using pulumi login --local
pulumi login --local
# or to use Pulumi cloud
pulumi logout
pulumi login
# or
pulumi login file:///$PWD

# Create a new Pulumi project that deployes a helm chart
pulumi new helm-kubernetes-go --force --name helm2 --description "My pulumi helm project" --stack dev

# Create chart
mkdir charts
helm create charts/foo
```

Not update the `charts/foo/values.yaml` to use `httpbin`

```yaml
image:
  repository: kennethreitz/httpbin
  pullPolicy: IfNotPresent
  # Overrides the image tag whose default is the chart appVersion.
  tag: "latest"
```

To use a local chart you need to comment out the `RepositoryOpts` in `main.go` in the `helmv3.NewRelease` section.  Also, you need to point to the local chart.

```go
Chart:     pulumi.String("./charts/foo"),
```

Then finally deploy it.

```sh
pulumi up
```

Things to try:

```sh
pulumi stack ls
pulumi stack output --json
pulumi stack export  | jq .

# see Pulumi.<stack>.yaml
pulumi config set FOO bar
pulumi config set MY_PASS bla --secret

pulumi preview
```
