package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2/internal"
)

type IpamAllocation interface {
	awscdk.Resource
	IIpamAllocation
	Allocation() IIpamAllocationConfiguration
	Description() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	IpamAllocationCidr() *string
	IpamAllocationId() *string
	IpamPool() IIpamPool
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
	Resource() awsec2.CfnIPAMAllocation
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
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
}

// The jsii proxy struct for IpamAllocation
type jsiiProxy_IpamAllocation struct {
	internal.Type__awscdkResource
	jsiiProxy_IIpamAllocation
}

func (j *jsiiProxy_IpamAllocation) Allocation() IIpamAllocationConfiguration {
	var returns IIpamAllocationConfiguration
	_jsii_.Get(
		j,
		"allocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamAllocation) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamAllocation) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamAllocation) IpamAllocationCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamAllocationCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamAllocation) IpamAllocationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamAllocationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamAllocation) IpamPool() IIpamPool {
	var returns IIpamPool
	_jsii_.Get(
		j,
		"ipamPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamAllocation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamAllocation) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamAllocation) Resource() awsec2.CfnIPAMAllocation {
	var returns awsec2.CfnIPAMAllocation
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamAllocation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewIpamAllocation(scope constructs.IConstruct, id *string, props *IpamAllocationProps) IpamAllocation {
	_init_.Initialize()

	if err := validateNewIpamAllocationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_IpamAllocation{}

	_jsii_.Create(
		"cdk-extensions.ec2.IpamAllocation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewIpamAllocation_Override(i IpamAllocation, scope constructs.IConstruct, id *string, props *IpamAllocationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.IpamAllocation",
		[]interface{}{scope, id, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func IpamAllocation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIpamAllocation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamAllocation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func IpamAllocation_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateIpamAllocation_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamAllocation",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func IpamAllocation_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateIpamAllocation_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamAllocation",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpamAllocation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IpamAllocation) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpamAllocation) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := i.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpamAllocation) GetResourceNameAttribute(nameAttr *string) *string {
	if err := i.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpamAllocation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

