# pulumi-learn

Exercises for learning [pulumi](https://www.pulumi.com/)

## What is Pulumi

### IaC Tool

* Pulumi is an Infrastructure as Code (IaC) tool

### General purpose programming languages vs domain-specific language (DSL)

* A core point of differentiation between Pulumi and other IaC offerings has been the ability to use popular general purpose programming languages 
* TypeScript/JavaScript, Python, Go, C#, Java, and YAML
* and their rich software engineering ecosystems, IDE and test frameworks

### Pulumi CLI

* **CLI** - Primarily driven via the `pulumi` [CLI](https://www.pulumi.com/docs/cli/) but also support a Pulumi Operator

### Pulumi Operator

* Exposes the Pulumi Stack as a first-class Kubernetes API resource
* Use the StackController to drive the updates.
* It allows users to adopt a GitOps workflow

## What I Like

* **Documentation** - The Doc is good
* **Easy things are easy** - [Pulumi YAML](https://www.pulumi.com/docs/languages-sdks/yaml/) is about as simple as using `kubectl apply` with k8s manifests. For more on the Pulumi YAML support see this [blog post](https://www.pulumi.com/blog/pulumi-yaml/)
* **Headroom** - If your use case is complex Pulumi supports that by allowing you to use a general purpose programming language such as Go, CUE and code reuse via [Pulumi Components](https://www.pulumi.com/docs/concepts/resources/components/) for example
* **Strongly typed configuration** - Pulumi supports [type specifications](https://www.pulumi.com/docs/concepts/config/#strongly-typed-configuration) for configuration, including setting defaults, but it is optional
* **Helm** - Deploying a Helm chart with Pulumi is very straightforward and easy

## What I find interesting

Things I liked but would need to look into in more detail.

* **Secrets Management** - Pulumi can manage [secret values](https://www.pulumi.com/docs/concepts/secrets/)
* **Components** - I would be interested to learn more about [Pulumi Components](https://www.pulumi.com/docs/concepts/resources/components/) and how to share code
* **Pulumi Operator** - I would like to know more about how to implement GitOps via the Pulumi Operator
* **CUE support** - languages that are designed to compile down to YAML/JSON can be used directly with [Pulumi YAML](https://www.pulumi.com/blog/pulumi-yaml/#yaml-as-a-compilation-target-and-cue-support)

## What I Don't Like

### The relationship between Stacks and Projects

The relationship between stacks and projects is such that we need <u>one stack per project, per environment</u>. I would like the option to share a single stack for a given environment.

This [GitHub Issue](https://github.com/pulumi/pulumi/issues/8402) describes the problem well

>Currently, Pulumi maintains a 1:1 relationship between stacks and projects. This means that if I have multiple projects targeting the same "environment", instead of being able to define one stack and one corresponding configuration file; I'm on the hook for one stack per project, per environment.

* You can use the output of one stack as input to other stacks but the I don't see a clean way to deploy multiple projects at once.
* Pulumi does support [Project Level Configuration](https://www.pulumi.com/docs/concepts/config/#project-level-configuration) which allows setting configuration at the project level instead of having to repeat the configuration setting in each stack’s configuration file. But this is per project and does not address the multiple project issue.

We will see this in the `multi-stacks` example. Notice there are multiple `Pulumi.mystack1.yaml` files

```sh
.
├── README.md
├── allUp.sh
├── app1
│   ├── Pulumi.mystack1.yaml
│   ├── Pulumi.yaml
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── app2
│   ├── Pulumi.mystack1.yaml
│   ├── Pulumi.yaml
│   ├── go.mod
│   ├── go.sum
│   └── main.go
└── infr
    ├── Pulumi.mystack1.yaml
    └── Pulumi.yaml
```

## Hello World

Lets look at a hello world Pulumi example first, then talk about the concepts then finish up with more demos.

```sh
#
# Install the Pulumi CLI
#
brew install pulumi/tap/pulumi

#
# Create a folder to hold our Pulumi project
#
mkdir cd hello-world
cd hello-world

#
# To use Pulumi without the Pulumi Cloud
# Save state in the current directory
#
pulumi logout
pulumi login file:///$PWD

#
# A folder named ".pulumi" is created to hold our state
#
$ tree -a
.
└── .pulumi
    ├── meta.yaml
    └── meta.yaml.attrs

#
# Create a new Pulumi project that uses k8s manifests
#
pulumi new kubernetes-go --name hi --description hi --stack demo --dir go
Created project 'hi'

Created stack 'demo'
Enter your passphrase to protect config/secrets: 
Re-enter your passphrase to confirm: 

Installing dependencies...

Finished installing dependencies

Your new project is ready to go! ✨

To perform an initial deployment, run 'cd go', then, run `pulumi up`

#
# We can see that a new stack is created in our state
# And we have a go project that we can use to generate a k8s deployment
#
$ tree -a
.
├── .pulumi
│   ├── locks
│   │   └── organization
│   │       └── hi
│   │           └── demo
│   ├── meta.yaml
│   ├── meta.yaml.attrs
│   └── stacks
│       └── hi
│           ├── demo.json
│           └── demo.json.attrs
└── go
    ├── Pulumi.demo.yaml
    ├── Pulumi.yaml
    ├── go.mod
    ├── go.sum
    └── main.go

# 
# Example config var
# You will see a new entry in go/Pulumi.demo.yaml 
#
pulumi config set FOO  bar --cwd ./go
pulumi config set PASSWORD opensesame --secret --cwd ./go

$ cat go/Pulumi.demo.yaml 
encryptionsalt: v1:nApVG/BPNYo=:v1:z9nHwmyDsdP6VFnP:L7VCw7Vi7030vC906zmeQ/RlTnjnjg==
config:
  hi:FOO: bar
  hi:PASSWORD:
    secure: v1:bzTfJl4IRVFb9hqI:pWFaDaug9O6jdGibQ2VtnJZHr7WTjjpYFsQ=

#
# Deploy the stack
#
pulumi up --cwd ./go

$ kubectl get deployments
NAME               READY   UP-TO-DATE   AVAILABLE   AGE
app-dep-a21fb72b   1/1     1            1           5s

#
# Destroy the stack
#
pulumi destroy --cwd ./go


##########################################################################
# Let's do the same thing but with YAML
##########################################################################


#
# Create a new Pulumi project that uses k8s manifests
#
$ pulumi new kubernetes-yaml --name hi --description hi --stack demo --dir yaml
error: a project with this name already exists: hi

pulumi new kubernetes-yaml --name hiyaml --description hiyaml --stack demo --dir yaml

#
# We can see 
# - new hiyaml stack added to our state 
# - new hiyaml organization
# - and our new yaml project files 
# 
#
$ tree -a
.
├── .pulumi
│   ├── backups
│   │   └── hi
│   │       └── demo
│   │           ├── demo.1691528747666666000.json
│   │           ├── demo.1691528747666666000.json.attrs
│   │           ├── demo.1691528974507474000.json
│   │           └── demo.1691528974507474000.json.attrs
│   ├── history
│   │   └── hi
│   │       └── demo
│   │           ├── demo-1691528747664742000.checkpoint.json
│   │           ├── demo-1691528747664742000.checkpoint.json.attrs
│   │           ├── demo-1691528747664742000.history.json
│   │           ├── demo-1691528747664742000.history.json.attrs
│   │           ├── demo-1691528974506170000.checkpoint.json
│   │           ├── demo-1691528974506170000.checkpoint.json.attrs
│   │           ├── demo-1691528974506170000.history.json
│   │           └── demo-1691528974506170000.history.json.attrs
│   ├── locks
│   │   └── organization
│   │       ├── hi
│   │       │   └── demo
│   │       └── hiyaml
│   │           └── demo
│   ├── meta.yaml
│   ├── meta.yaml.attrs
│   └── stacks
│       ├── hi
│       │   ├── demo.json
│       │   ├── demo.json.attrs
│       │   ├── demo.json.bak
│       │   └── demo.json.bak.attrs
│       └── hiyaml
│           ├── demo.json
│           └── demo.json.attrs
├── go
│   ├── Pulumi.demo.yaml
│   ├── Pulumi.yaml
│   ├── go.mod
│   ├── go.sum
│   └── main.go
└── yaml
    ├── Pulumi.demo.yaml
    └── Pulumi.yaml
# 
# Example config var
# You will see a new entry in go/Pulumi.demo.yaml 
#
pulumi config set FOO  bar --cwd ./yaml
pulumi config set PASSWORD opensesame --secret --cwd ./yaml

$ cat yaml/Pulumi.demo.yaml 
encryptionsalt: v1:JIfxkMIVjCo=:v1:lRDUUtKrkFj0KLGV:UqRRj8Qymem9Z/Tyo5GnqI+8BY7zkA==
config:
  hiyaml:FOO: bar
  hiyaml:PASSWORD:
    secure: v1:LNiNJTKKJv2hclII:XZnuD7J9GVrDdzU/FsnGVIbsZ1m2FOsKtHk=
    
#
# Let's look at the Pulumi.yaml file
# Note: everything nested inside "properties" below is directly copy-pasted from a Kubernetes YAML specification
#
$ cat yaml/Pulumi.yaml
name: hiyaml
runtime: yaml
description: hiyaml
outputs:
  name: ${deployment.metadata.name}
resources:
  deployment:
    properties:
      spec:
        replicas: 1
        selector:
          matchLabels: ${appLabels}
        template:
          metadata:
            labels: ${appLabels}
          spec:
            containers:
              - image: nginx
                name: nginx
    type: kubernetes:apps/v1:Deployment
variables:
  appLabels:
    app: nginx

#
# Deploy the stack
#
pulumi up --cwd ./yaml

$  kubectl get deployments
NAME                  READY   UP-TO-DATE   AVAILABLE   AGE
deployment-af8c6159   1/1     1            1           71s

#
# Destroy the stack
#
pulumi destroy --cwd ./yaml

```

Now that we have a folder with several apps we can see

1. The stacks are not shared for a given target environment
2. We need to create our own tooling to act multiple applications. For example the following script could be used to deploy all apps at once.

```sh
#! /bin/bash

pulumi up --cwd ./go
pulumi up --cwd ./yaml
```

## Pulumi Concepts

### Overview

![overview](img/overview.png)

### How Pulumi works

![how-works](img/how-works.png)

### Projects

* A [Pulumi project](https://www.pulumi.com/docs/concepts/projects/) is any folder which contains a `Pulumi.yaml`
* A new project can be created with `pulumi new`
* A typical Pulumi.yaml file looks like the following:

```yaml
name: webserver
runtime: nodejs
description: Basic example of an AWS web server accessible over HTTP.
```

### Stacks

* A [stacks](https://www.pulumi.com/docs/concepts/stack/#stacks) is an isolated, independently configurable instance of a Pulumi program. 
* Stacks are commonly used to denote different phases of development (such as `development`, `staging`, and `production`)
* To initilize a new stack

```sh
pulumi stack init staging
```

* A stack can export values as stack outputs. They can be used for important values like resource IDs, computed IP addresses, and DNS names.

### State and Backends

* Each stack has its own state, and state is how Pulumi knows when and how to create, read, delete, or update cloud resources.
* Pulumi supports two classes of state backends for storing your infrastructure state:
  * **Service**: a managed cloud experience using the online or self-hosted Pulumi Cloud application
  * **Self-Managed**: a manually managed object store, including AWS S3, Azure Blob Storage, Google Cloud Storage, any AWS S3 compatible server such as Minio or Ceph, or your local filesystem

* The login command logs you into a backend:

```sh
pulumi login <URL or Path>
```

### Configuration

* Pulumi stores [configuration](https://www.pulumi.com/docs/concepts/config/#configuration) as key-value pairs for each stack in the project stack setting file `Pulumi.<stack-name>.yaml`

```sh
$ pulumi config set name Foo
$ pulumi config get name
Foo
```

Examples in different languages:

yaml

```yaml
config:
  name:
    type: string
  lucky:
    default: 42
  secret:
    type: string
    secret: true
```

Go

```go
package main

import (
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)
func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        conf := config.New(ctx, "")
        name := conf.Require("name")
        lucky, err := conf.TryInt("lucky")
        if err != nil {
            lucky = 42
        }
        secret := conf.RequireSecret("secret")
        return nil
    }
}
```

JavaScript

```javascript
let config = new pulumi.Config();
let name = config.require("name");
let lucky = config.getNumber("lucky") || 42;
let secret = config.requireSecret("secret");
```

### Secrets

Pulumi supports encrypting specific values as “[secrets](https://www.pulumi.com/docs/concepts/secrets/)” and stores them in your stack settings file.

In the [stack setting file](https://www.pulumi.com/docs/concepts/projects/#stack-settings-file) section of the doc it says:

>For stacks that are actively developed by multiple members of a team, the recommended practice is to check your stack settings file into source control as a means of collaboration.

We will see this in the `helm2` example

```sh
pulumi config set MY_PASS bla --secret
```

My stack settings file - `Pulumi.dev.yaml`

```yaml
encryptionsalt: v1:GlKk87t1m9k=:v1:jqGSNfyuIt3OgUz9:13FSoA4gn4qBD+OlozIDdbUaNBQcyw==
config:
  helm2:FOO: bar
  helm2:MY_PASS:
    secure: v1:uvdz4bigX4DPchxL:rebKmW2PLMaePdhE/Ye+r/HG5Q==
  helm2:k8sNamespace: helm2ns
```

See also the [Pulumi new provider options](https://www.pulumi.com/docs/cli/commands/pulumi_new/#synopsis)

## Demos

