//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ITerraformResource) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_ITerraformResource) validateSetLifecycleParameters(val *TerraformResourceLifecycle) error {
	return nil
}
