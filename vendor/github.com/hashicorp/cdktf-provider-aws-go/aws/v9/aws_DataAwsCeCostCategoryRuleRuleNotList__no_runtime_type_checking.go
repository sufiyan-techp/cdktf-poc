//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsCeCostCategoryRuleRuleNotList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsCeCostCategoryRuleRuleNotList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsCeCostCategoryRuleRuleNotList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsCeCostCategoryRuleRuleNotList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsCeCostCategoryRuleRuleNotList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsCeCostCategoryRuleRuleNotListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

