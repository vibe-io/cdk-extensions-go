package ec2patterns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2patterns/internal"
)

type NetworkIsolatedClientVpnEndpoint interface {
	awscdk.Resource
	awsec2.IClientVpnEndpoint
	awsec2.IConnectable
	AuthorizeAllUsersToVpcCidr() *bool
	ClientCertificate() awscertificatemanager.ICertificate
	ClientConnectionHandler() awsec2.IClientVpnConnectionHandler
	ClientLoginBanner() *string
	ClientVpnEndpoint() awsec2.ClientVpnEndpoint
	// The network connections associated with this resource.
	Connections() awsec2.Connections
	Description() *string
	DnsServers() *[]*string
	// The endpoint ID.
	EndpointId() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	Logging() *bool
	LogGroup() awslogs.ILogGroup
	LogStream() awslogs.ILogStream
	MaxAzs() *float64
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
	Port() awsec2.VpnPort
	SecurityGroups() *[]awsec2.ISecurityGroup
	SelfServicePortal() *bool
	ServerCertificate() awscertificatemanager.ICertificate
	SplitTunnel() *bool
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	Subnets() *[]awsec2.ISubnet
	// Dependable that can be depended upon to force target networks associations.
	TargetNetworksAssociated() constructs.IDependable
	TransitGateway() ec2.ITransitGateway
	TransportProtocol() awsec2.TransportProtocol
	UserBasedAuthentication() awsec2.ClientVpnUserBasedAuthentication
	Vpc() awsec2.IVpc
	VpcCidrBlock() ec2.VpcCidrBlock
	VpnCidr() ec2.IIpv4CidrAssignment
	AddAuthorizationRule(id *string, options *AddAuthorizationRuleOptions) awsec2.ClientVpnAuthorizationRule
	AddMultiSubnetRoute(id *string, options *AddMultiSubnetRouteOptions) interface{}
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
	RegisterTransitGateway(transitGateway ec2.ITransitGateway)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for NetworkIsolatedClientVpnEndpoint
type jsiiProxy_NetworkIsolatedClientVpnEndpoint struct {
	internal.Type__awscdkResource
	internal.Type__awsec2IClientVpnEndpoint
	internal.Type__awsec2IConnectable
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) AuthorizeAllUsersToVpcCidr() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"authorizeAllUsersToVpcCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) ClientCertificate() awscertificatemanager.ICertificate {
	var returns awscertificatemanager.ICertificate
	_jsii_.Get(
		j,
		"clientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) ClientConnectionHandler() awsec2.IClientVpnConnectionHandler {
	var returns awsec2.IClientVpnConnectionHandler
	_jsii_.Get(
		j,
		"clientConnectionHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) ClientLoginBanner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientLoginBanner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) ClientVpnEndpoint() awsec2.ClientVpnEndpoint {
	var returns awsec2.ClientVpnEndpoint
	_jsii_.Get(
		j,
		"clientVpnEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) EndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) Logging() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) LogStream() awslogs.ILogStream {
	var returns awslogs.ILogStream
	_jsii_.Get(
		j,
		"logStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) MaxAzs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAzs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) Port() awsec2.VpnPort {
	var returns awsec2.VpnPort
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) SecurityGroups() *[]awsec2.ISecurityGroup {
	var returns *[]awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) SelfServicePortal() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"selfServicePortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) ServerCertificate() awscertificatemanager.ICertificate {
	var returns awscertificatemanager.ICertificate
	_jsii_.Get(
		j,
		"serverCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) SplitTunnel() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"splitTunnel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) Subnets() *[]awsec2.ISubnet {
	var returns *[]awsec2.ISubnet
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) TargetNetworksAssociated() constructs.IDependable {
	var returns constructs.IDependable
	_jsii_.Get(
		j,
		"targetNetworksAssociated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) TransitGateway() ec2.ITransitGateway {
	var returns ec2.ITransitGateway
	_jsii_.Get(
		j,
		"transitGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) TransportProtocol() awsec2.TransportProtocol {
	var returns awsec2.TransportProtocol
	_jsii_.Get(
		j,
		"transportProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) UserBasedAuthentication() awsec2.ClientVpnUserBasedAuthentication {
	var returns awsec2.ClientVpnUserBasedAuthentication
	_jsii_.Get(
		j,
		"userBasedAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) VpcCidrBlock() ec2.VpcCidrBlock {
	var returns ec2.VpcCidrBlock
	_jsii_.Get(
		j,
		"vpcCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedClientVpnEndpoint) VpnCidr() ec2.IIpv4CidrAssignment {
	var returns ec2.IIpv4CidrAssignment
	_jsii_.Get(
		j,
		"vpnCidr",
		&returns,
	)
	return returns
}


func NewNetworkIsolatedClientVpnEndpoint(scope constructs.IConstruct, id *string, props *NetworkIsolatedClientVpnEndpointProps) NetworkIsolatedClientVpnEndpoint {
	_init_.Initialize()

	if err := validateNewNetworkIsolatedClientVpnEndpointParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkIsolatedClientVpnEndpoint{}

	_jsii_.Create(
		"cdk-extensions.ec2_patterns.NetworkIsolatedClientVpnEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNetworkIsolatedClientVpnEndpoint_Override(n NetworkIsolatedClientVpnEndpoint, scope constructs.IConstruct, id *string, props *NetworkIsolatedClientVpnEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2_patterns.NetworkIsolatedClientVpnEndpoint",
		[]interface{}{scope, id, props},
		n,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func NetworkIsolatedClientVpnEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkIsolatedClientVpnEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2_patterns.NetworkIsolatedClientVpnEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func NetworkIsolatedClientVpnEndpoint_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateNetworkIsolatedClientVpnEndpoint_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2_patterns.NetworkIsolatedClientVpnEndpoint",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func NetworkIsolatedClientVpnEndpoint_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateNetworkIsolatedClientVpnEndpoint_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2_patterns.NetworkIsolatedClientVpnEndpoint",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func NetworkIsolatedClientVpnEndpoint_DEFAULT_VPN_CIDR() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk-extensions.ec2_patterns.NetworkIsolatedClientVpnEndpoint",
		"DEFAULT_VPN_CIDR",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkIsolatedClientVpnEndpoint) AddAuthorizationRule(id *string, options *AddAuthorizationRuleOptions) awsec2.ClientVpnAuthorizationRule {
	if err := n.validateAddAuthorizationRuleParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsec2.ClientVpnAuthorizationRule

	_jsii_.Invoke(
		n,
		"addAuthorizationRule",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIsolatedClientVpnEndpoint) AddMultiSubnetRoute(id *string, options *AddMultiSubnetRouteOptions) interface{} {
	if err := n.validateAddMultiSubnetRouteParameters(id, options); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"addMultiSubnetRoute",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIsolatedClientVpnEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := n.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (n *jsiiProxy_NetworkIsolatedClientVpnEndpoint) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIsolatedClientVpnEndpoint) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := n.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIsolatedClientVpnEndpoint) GetResourceNameAttribute(nameAttr *string) *string {
	if err := n.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIsolatedClientVpnEndpoint) RegisterTransitGateway(transitGateway ec2.ITransitGateway) {
	if err := n.validateRegisterTransitGatewayParameters(transitGateway); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"registerTransitGateway",
		[]interface{}{transitGateway},
	)
}

func (n *jsiiProxy_NetworkIsolatedClientVpnEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

