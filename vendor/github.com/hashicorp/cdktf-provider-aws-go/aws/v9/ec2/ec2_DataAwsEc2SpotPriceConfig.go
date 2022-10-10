package ec2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// AWS EC2.
type DataAwsEc2SpotPriceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_spot_price#availability_zone DataAwsEc2SpotPrice#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_spot_price#filter DataAwsEc2SpotPrice#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_spot_price#id DataAwsEc2SpotPrice#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_spot_price#instance_type DataAwsEc2SpotPrice#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_spot_price#timeouts DataAwsEc2SpotPrice#timeouts}
	Timeouts *DataAwsEc2SpotPriceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

