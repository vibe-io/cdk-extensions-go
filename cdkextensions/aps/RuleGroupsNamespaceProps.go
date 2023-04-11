package aps


// Configuration for the RuleGroupsNamespace resource.
type RuleGroupsNamespaceProps struct {
	// The AWS account ID this resource belongs to.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// ARN to deduce region and account from.
	//
	// The ARN is parsed and the account and region are taken from the ARN.
	// This should be used for imported resources.
	//
	// Cannot be supplied together with either `account` or `region`.
	EnvironmentFromArn *string `field:"optional" json:"environmentFromArn" yaml:"environmentFromArn"`
	// The value passed in by users to the physical name prop of the resource.
	//
	// - `undefined` implies that a physical name will be allocated by
	//   CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//   by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The APS workspace that contains this rule groups namespace.
	// See: [RuleGroupsNamespace Workspace](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-rulegroupsnamespace.html#cfn-aps-rulegroupsnamespace-workspace)
	//
	Workspace IWorkspace `field:"required" json:"workspace" yaml:"workspace"`
	// The rules definition file for this namespace.
	// See: [RuleGroupsNamespace Data](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-rulegroupsnamespace.html#cfn-aps-rulegroupsnamespace-data)
	//
	Configuration IRuleGroupConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// The name of the rule groups namespace.
	// See: [RuleGroupsNamespace Name](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-rulegroupsnamespace.html#cfn-aps-rulegroupsnamespace-name)
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}
