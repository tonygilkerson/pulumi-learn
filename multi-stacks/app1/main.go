package main

import (
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		//
		// Get stack reference for infr
		//
		infrStack, err := pulumi.NewStackReference(ctx, "organization/infr/mystack1", nil)
		if err != nil {
			return err
		}
		app1Ns := infrStack.GetOutput(pulumi.String("app1Ns"))


		appLabels := pulumi.StringMap{
			"app1": pulumi.String("nginx"),
		}
		deployment, err := appsv1.NewDeployment(ctx, "app1", &appsv1.DeploymentArgs{
			//
			// Use the namespace set in the infr stack
			//
			Metadata: &metav1.ObjectMetaArgs{
				Namespace: app1Ns.AsStringOutput(),
			},
			Spec: appsv1.DeploymentSpecArgs{
				Selector: &metav1.LabelSelectorArgs{
					MatchLabels: appLabels,
				},
				Replicas: pulumi.Int(1),
				Template: &corev1.PodTemplateSpecArgs{
					Metadata: &metav1.ObjectMetaArgs{
						Labels: appLabels,
					},
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							corev1.ContainerArgs{
								Name:  pulumi.String("nginx"),
								Image: pulumi.String("nginx"),
							}},
					},
				},
			},
		})
		if err != nil {
			return err
		}

		// Looks like this is a bug in the template
		// ctx.Export("name", deployment.Metadata.Elem().Name())
		ctx.Export("name", deployment.Metadata.Name())

		return nil
	})
}
