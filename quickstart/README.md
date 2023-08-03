# Quickstart

```sh
# Run all commands from this folder
cd quickstart

# To use Pulumi without the Pulumi Cloud, log in using pulumi login --local
pulumi login --local

# Create a new Pulumi project that uses k8s manifests
pulumi new kubernetes-go --force --name quickstart --description "My quickstart project" --stack dev
```

Let’s review some of the generated project files:

- `Pulumi.yaml` defines the project.
- `Pulumi.dev.yaml` contains configuration values for the stack we initialized.
- `main.go` is the Pulumi program that defines your stack resources.

Next lets deploy our app

```sh
pulumi up
```
