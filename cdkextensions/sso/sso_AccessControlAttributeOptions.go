package sso


// Configuration options for adding an ABAC attribute to IAM Identity Center.
type AccessControlAttributeOptions struct {
	// The name of the attribute associated with your identities in your identity source.
	//
	// This is used to map a specified attribute in your
	// identity source with an attribute in IAM Identity Center.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of identity sources to use when mapping a specified attribute to IAM Identity Center.
	// See: [AWS::SSO::InstanceAccessControlAttributeConfiguration AccessControlAttributeValue](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributevalue.html#cfn-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributevalue-source)
	//
	Sources *[]*string `field:"optional" json:"sources" yaml:"sources"`
}

