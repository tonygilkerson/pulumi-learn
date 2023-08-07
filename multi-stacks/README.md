# Multi Stacks

Micro stacks demo

```sh
# Run all commands from here
cd multi-stacks

# Storage - this will create a .pulumi subfolder
pulumi logout
pulumi login file:///$PWD
```

## infr

Create project for infrastructure stuff using a simple YAML provider

```sh
mkdir infr

# Create project from template
pulumi new kubernetes-yaml --cwd ./infr --stack mystack1 --name infr --description infr 

# Replace default content
echo '
name: infr
runtime: yaml
description: infr
config:
  hostname:
    default: example.com
    type: string
variables:
  myLabels:
    foo: bar
outputs:
  app1Ns: ${app1Ns.metadata.name}
  app2Ns: ${app2Ns.metadata.name}
resources:
  app1Ns:
    type: kubernetes:core/v1:Namespace
    properties:
      metadata:
        name: app1
  app2Ns:
    type: kubernetes:core/v1:Namespace
    properties:
      metadata:
        name: app2
' > infr/Pulumi.yaml

# Example config var
# You will see a new entry in infr/Pulumi.mystack1.yaml
pulumi config set FOO bar --cwd ./infr

# Preview the changes
pulumi preview --cwd ./infr 

# Deploy
pulumi up --cwd ./infr 
```

>The [YAML provider supports CUE](https://www.pulumi.com/blog/pulumi-yaml/#yaml-as-a-compilation-target-and-cue-support) and can be [converted to a Pulumi Language](https://www.pulumi.com/blog/pulumi-yaml/#convert-to-other-pulumi-languages)

## app1

```sh
mkdir app1

# Create project from template
pulumi new kubernetes-go --cwd ./app1 --stack mystack1 --name app1 --description app1

# Update the app to use the namespace from the infr stack (not shown here)

# Preview the changes
pulumi preview --cwd ./app1

# Deploy
pulumi up --cwd ./app1
```

## app2

```sh
mkdir app2

# Create project from template
pulumi new kubernetes-go --cwd ./app2 --stack mystack1 --name app2 --description app2

# Update the app to use the namespace from the infr stack (not shown here)

# Preview the changes
pulumi preview --cwd ./app2

# Deploy
pulumi up --cwd ./app2
```
