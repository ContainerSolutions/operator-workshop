# Solution for Exercise-1

In the [generated operator](https://google.pl) there is a default, hardcoded specification for the container and your role is to:

* [Extend the api structure](https://github.com/ContainerSolutions/operator-workshop/blob/71cdfa9bbb7649405fbf900bfc349d471bfdcdf5/pkg/apis/app/v1alpha1/exampleworkshop_types.go#L16)
* Regenerate the api `operator-sdk generate k8s` and `operator-sdk generate openapi"`
* [Define in the resource definition](https://github.com/ContainerSolutions/operator-workshop/blob/exercise-1/deploy/crds/app_v1alpha1_exampleworkshop_cr.yaml)
* [Write the required go code to reflect your changes inside your controller file](https://github.com/ContainerSolutions/operator-workshop/blob/71cdfa9bbb7649405fbf900bfc349d471bfdcdf5/pkg/controller/exampleworkshop/exampleworkshop_controller.go#L133)
* Deploy your operator `operator-sdk up local --namespace=default`
