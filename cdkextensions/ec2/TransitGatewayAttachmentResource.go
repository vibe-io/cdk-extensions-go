package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Attaches a VPC to a transit gateway.
//
// If you attach a VPC with a CIDR range that overlaps the CIDR range of a VPC
// that is already attached, the new VPC CIDR range is not propagated to the
// default propagation route table.
type TransitGatewayAttachmentResource interface {
	TransitGatewayAttachmentBase
	// Enables appliance mode on the attachment.
	//
	// When appliance mode is enabled, all traffic flowing between attachments is
	// forwarded to an appliance in a shared VPC to be inspected and processed.
	// See: [Appliance in a shared services VPC](https://docs.aws.amazon.com/vpc/latest/tgw/transit-gateway-appliance-scenario.html)
	//
	ApplianceModeSupport() *bool
	// Enables DNS support for the attachment.
	//
	// With DNS Support enabled public DNS names that resolve to a connected VPC
	// will be translated to private IP addresses when resolved in a connected VPC.
	// See: [TransitGatewayVpcAttachment DnsSupport](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayvpcattachment-options.html#cfn-ec2-transitgatewayvpcattachment-options-dnssupport)
	//
	DnsSupport() *bool
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// Enables DNS support for the attachment.
	//
	// With DNS Support enabled public DNS names that resolve to a connected VPC
	// will be translated to private IP addresses when resolved in a connected VPC.
	// See: [IPv6 connectivity with TransitGateway](https://docs.aws.amazon.com/whitepapers/latest/ipv6-on-aws/amazon-vpc-connectivity-options-for-ipv6.html#ipv6-connectivity-with-transit-gateway)
	//
	Ipv6Support() *bool
	// The name of the Transit Gateway Attachment.
	//
	// Used to tag the attachment with a name that will be displayed in the AWS
	// EC2 console.
	// See: [TransitGatewayVpcAttachment Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-tags)
	//
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The subnets where the attachment should be created.
	//
	// Can select up to one subnet per Availability Zone.
	// See: [TransitGatewayVpcAttachment SubnetIds](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-subnetids)
	//
	Subnets() *awsec2.SubnetSelection
	// The transit gateway for which the attachment should be created.
	TransitGateway() ITransitGateway
	// The ARN of this Transit Gateway Attachment.
	TransitGatewayAttachmentArn() *string
	// The ID of this Transit Gateway Attachment.
	TransitGatewayAttachmentId() *string
	// The VPC where the attachment should be created.
	// See: [TransitGatewayVpcAttachment VpcId](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-vpcid)
	//
	Vpc() awsec2.IVpc
	// Adds a route that directs traffic to this transit gateway attachment.
	//
	// Returns: The TransitGatewayRoute that was added.
	AddRoute(id *string, cidr *string, routeTable ITransitGatewayRouteTable) ITransitGatewayRoute
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
	// Translates a boolean input into the strings used by the transit gateway attachment resource to implement boolean values.
	//
	// Returns: The string used to reprersent the input boolean or undefined if
	// the input boolean is undefined.
	TranslateBoolean(val *bool) *string
}

// The jsii proxy struct for TransitGatewayAttachmentResource
type jsiiProxy_TransitGatewayAttachmentResource struct {
	jsiiProxy_TransitGatewayAttachmentBase
}

func (j *jsiiProxy_TransitGatewayAttachmentResource) ApplianceModeSupport() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"applianceModeSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGatewayAttachmentResource) DnsSupport() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"dnsSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGatewayAttachmentResource) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGatewayAttachmentResource) Ipv6Support() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ipv6Support",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGatewayAttachmentResource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGatewayAttachmentResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGatewayAttachmentResource) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGatewayAttachmentResource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGatewayAttachmentResource) Subnets() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGatewayAttachmentResource) TransitGateway() ITransitGateway {
	var returns ITransitGateway
	_jsii_.Get(
		j,
		"transitGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGatewayAttachmentResource) TransitGatewayAttachmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayAttachmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGatewayAttachmentResource) TransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGatewayAttachmentResource) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Creates a new instance of the TransitGatewayAttachment class.
func NewTransitGatewayAttachmentResource_Override(t TransitGatewayAttachmentResource, scope constructs.Construct, id *string, props *TransitGatewayAttachmentResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.TransitGatewayAttachmentResource",
		[]interface{}{scope, id, props},
		t,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func TransitGatewayAttachmentResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTransitGatewayAttachmentResource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.TransitGatewayAttachmentResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func TransitGatewayAttachmentResource_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateTransitGatewayAttachmentResource_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.TransitGatewayAttachmentResource",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func TransitGatewayAttachmentResource_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateTransitGatewayAttachmentResource_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.TransitGatewayAttachmentResource",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitGatewayAttachmentResource) AddRoute(id *string, cidr *string, routeTable ITransitGatewayRouteTable) ITransitGatewayRoute {
	if err := t.validateAddRouteParameters(id, cidr, routeTable); err != nil {
		panic(err)
	}
	var returns ITransitGatewayRoute

	_jsii_.Invoke(
		t,
		"addRoute",
		[]interface{}{id, cidr, routeTable},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitGatewayAttachmentResource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := t.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (t *jsiiProxy_TransitGatewayAttachmentResource) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitGatewayAttachmentResource) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := t.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitGatewayAttachmentResource) GetResourceNameAttribute(nameAttr *string) *string {
	if err := t.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitGatewayAttachmentResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitGatewayAttachmentResource) TranslateBoolean(val *bool) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"translateBoolean",
		[]interface{}{val},
		&returns,
	)

	return returns
}

