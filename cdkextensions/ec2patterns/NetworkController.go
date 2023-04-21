package ec2patterns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2patterns/internal"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/networkmanager"
)

type NetworkController interface {
	awscdk.Resource
	AddressManager() IpAddressManager
	DefaultNetmask() *float64
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	FlowLogBucket() awss3.IBucket
	FlowLogFormat() ec2.FlowLogFormat
	GlobalNetwork() networkmanager.GlobalNetwork
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
	RegisteredAccounts() *[]*string
	RegisteredRegions() *[]*string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	AddHub(scope constructs.IConstruct, id *string, options *AddHubOptions) FourTierNetworkHub
	AddSpoke(scope constructs.IConstruct, id *string, options *AddNetworkOptions) FourTierNetworkSpoke
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
	RegisterAccount(account *string)
	RegisterCidr(scope constructs.IConstruct, id *string, cidr *string)
	RegisterRegion(region *string)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for NetworkController
type jsiiProxy_NetworkController struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_NetworkController) AddressManager() IpAddressManager {
	var returns IpAddressManager
	_jsii_.Get(
		j,
		"addressManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkController) DefaultNetmask() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultNetmask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkController) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkController) FlowLogBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"flowLogBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkController) FlowLogFormat() ec2.FlowLogFormat {
	var returns ec2.FlowLogFormat
	_jsii_.Get(
		j,
		"flowLogFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkController) GlobalNetwork() networkmanager.GlobalNetwork {
	var returns networkmanager.GlobalNetwork
	_jsii_.Get(
		j,
		"globalNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkController) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkController) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkController) RegisteredAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"registeredAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkController) RegisteredRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"registeredRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkController) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewNetworkController(scope constructs.IConstruct, id *string, props *NetworkControllerProps) NetworkController {
	_init_.Initialize()

	if err := validateNewNetworkControllerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkController{}

	_jsii_.Create(
		"cdk-extensions.ec2_patterns.NetworkController",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNetworkController_Override(n NetworkController, scope constructs.IConstruct, id *string, props *NetworkControllerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2_patterns.NetworkController",
		[]interface{}{scope, id, props},
		n,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func NetworkController_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkController_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2_patterns.NetworkController",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func NetworkController_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateNetworkController_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2_patterns.NetworkController",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func NetworkController_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateNetworkController_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2_patterns.NetworkController",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkController) AddHub(scope constructs.IConstruct, id *string, options *AddHubOptions) FourTierNetworkHub {
	if err := n.validateAddHubParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns FourTierNetworkHub

	_jsii_.Invoke(
		n,
		"addHub",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkController) AddSpoke(scope constructs.IConstruct, id *string, options *AddNetworkOptions) FourTierNetworkSpoke {
	if err := n.validateAddSpokeParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns FourTierNetworkSpoke

	_jsii_.Invoke(
		n,
		"addSpoke",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkController) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := n.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (n *jsiiProxy_NetworkController) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkController) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (n *jsiiProxy_NetworkController) GetResourceNameAttribute(nameAttr *string) *string {
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

func (n *jsiiProxy_NetworkController) RegisterAccount(account *string) {
	if err := n.validateRegisterAccountParameters(account); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"registerAccount",
		[]interface{}{account},
	)
}

func (n *jsiiProxy_NetworkController) RegisterCidr(scope constructs.IConstruct, id *string, cidr *string) {
	if err := n.validateRegisterCidrParameters(scope, id, cidr); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"registerCidr",
		[]interface{}{scope, id, cidr},
	)
}

func (n *jsiiProxy_NetworkController) RegisterRegion(region *string) {
	if err := n.validateRegisterRegionParameters(region); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"registerRegion",
		[]interface{}{region},
	)
}

func (n *jsiiProxy_NetworkController) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

