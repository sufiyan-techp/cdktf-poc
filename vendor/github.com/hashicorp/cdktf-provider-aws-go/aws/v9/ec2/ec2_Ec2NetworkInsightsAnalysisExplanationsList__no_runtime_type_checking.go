//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2NetworkInsightsAnalysisExplanationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

