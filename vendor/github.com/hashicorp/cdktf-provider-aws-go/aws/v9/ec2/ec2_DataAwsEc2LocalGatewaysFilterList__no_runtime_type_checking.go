//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsEc2LocalGatewaysFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsEc2LocalGatewaysFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsEc2LocalGatewaysFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataAwsEc2LocalGatewaysFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsEc2LocalGatewaysFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsEc2LocalGatewaysFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsEc2LocalGatewaysFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

