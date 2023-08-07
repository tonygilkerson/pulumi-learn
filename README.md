# pulumi-learn

Exercises for learning pulumi

```sh
brew install pulumi/tap/pulumi
```

## Of Interest

* https://www.pulumi.com/docs/using-pulumi/organizing-projects-stacks/
* [patterns - the centralized platform infrastructure repository](https://www.pulumi.com/blog/organizational-patterns-infra-repo/)

## What is Pulumi

* Pulumi is an Infrastructure as Code tool (IaC)
* A core point of differentiation between Pulumi and other IaC offerings has been the ability to use popular general purpose programming languages (TypeScript/JavaScript, Python, Go, C#, Java, and YAML) and their rich software engineering ecosystems, IDE and test frameworks
* Pulumi has a lot in common with Terraform
* Pulumi can manage secret values

## What I Like

https://www.pulumi.com/blog/pulumi-yaml/#yaml-as-a-compilation-target-and-cue-support

* The Doc is good
* **Easy things are easy** - [Pulumi Yaml](https://www.pulumi.com/docs/languages-sdks/yaml/) is about as simple as using `kubectl apply` with k8s manifests
* **Can handle high complex use cases** - If your use case is complex Pulumi supports that by allowing you to use a general purpose programming language such as Go, CUE and code [Pulumi Components](https://www.pulumi.com/docs/concepts/resources/components/)for example
* The different approaches are interoperable so you can use the easy 
* Pulumi component packages are reusable infrastructure components built in one Pulumi language, and exposed via Pulumi Schema to all other Pulumi languages
* Support for Helm

## What I Don't Like

* I have not found a way to deploy multiple projects at once. I think this is done with [stack-tags](https://www.pulumi.com/docs/concepts/stack/#stack-tags) and Stack tags are only supported with the Pulumi Cloud backend.
  * or maybe you do it with [stack references](https://www.pulumi.com/docs/concepts/stack/#stackreferences)
  * or maybe this[micro stacks](https://blog.bitsrc.io/managing-micro-stacks-using-pulumi-87053eeb8678)

see: https://github.com/pulumi/pulumi/issues/8402

## Thing to look at

* [Project vs Stack Config](https://www.pulumi.com/docs/concepts/config/#project-and-stack-configuration-scope)
* 
* 