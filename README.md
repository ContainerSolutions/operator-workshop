# Solution for Exercise-2

In the [generated operator](https://github.com/ContainerSolutions/operator-workshop/blob/exercise-2/pkg/controller/exampleworkshop/exampleworkshop_controller.go) there is a default, hardcoded specification for the container and your role is to:

* [Extend the api structure](https://github.com/ContainerSolutions/operator-workshop/blob/6c3ee071b61b55c74c123fff77cdfdb9a881148c/pkg/apis/app/v1alpha1/exampleworkshop_types.go#L12)
* Regenerate the api `operator-sdk generate k8s` and `operator-sdk generate openapi"`
* [Define in the resource definition](https://github.com/ContainerSolutions/operator-workshop/blob/exercise-2/deploy/crds/app_v1alpha1_exampleworkshop_cr.yaml)
* [Write the required go code to reflect your changes inside your controller file](https://github.com/ContainerSolutions/operator-workshop/blob/exercise-2/pkg/controller/exampleworkshop/exampleworkshop_controller.go)
* Deploy your operator `operator-sdk up local --namespace=default`
