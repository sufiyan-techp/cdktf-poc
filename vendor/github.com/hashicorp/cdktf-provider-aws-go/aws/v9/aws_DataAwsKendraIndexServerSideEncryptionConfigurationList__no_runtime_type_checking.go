//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsKendraIndexServerSideEncryptionConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsKendraIndexServerSideEncryptionConfigurationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsKendraIndexServerSideEncryptionConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsKendraIndexServerSideEncryptionConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsKendraIndexServerSideEncryptionConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsKendraIndexServerSideEncryptionConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

