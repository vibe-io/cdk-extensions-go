package rds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/rds/internal"
)

type DatabaseProxyEndpoint interface {
	awscdk.Resource
	awsec2.IConnectable
	Access() DatabaseProxyEndpointAccess
	// The network connections associated with this resource.
	Connections() awsec2.Connections
	DatabaseProxy() awsrds.IDatabaseProxy
	DatabaseProxyEndpointArn() *string
	DatabaseProxyEndpointHost() *string
	DatabaseProxyEndpointIsDefault() awscdk.IResolvable
	DatabaseProxyEndpointVpcId() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
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
	Resource() awsrds.CfnDBProxyEndpoint
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	Vpc() awsec2.IVpc
	VpcSubnets() *awsec2.SubnetSelection
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

// The jsii proxy struct for DatabaseProxyEndpoint
type jsiiProxy_DatabaseProxyEndpoint struct {
	internal.Type__awscdkResource
	internal.Type__awsec2IConnectable
}

func (j *jsiiProxy_DatabaseProxyEndpoint) Access() DatabaseProxyEndpointAccess {
	var returns DatabaseProxyEndpointAccess
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxyEndpoint) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxyEndpoint) DatabaseProxy() awsrds.IDatabaseProxy {
	var returns awsrds.IDatabaseProxy
	_jsii_.Get(
		j,
		"databaseProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxyEndpoint) DatabaseProxyEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseProxyEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxyEndpoint) DatabaseProxyEndpointHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseProxyEndpointHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxyEndpoint) DatabaseProxyEndpointIsDefault() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"databaseProxyEndpointIsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxyEndpoint) DatabaseProxyEndpointVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseProxyEndpointVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxyEndpoint) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxyEndpoint) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxyEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxyEndpoint) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxyEndpoint) Resource() awsrds.CfnDBProxyEndpoint {
	var returns awsrds.CfnDBProxyEndpoint
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxyEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxyEndpoint) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxyEndpoint) VpcSubnets() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"vpcSubnets",
		&returns,
	)
	return returns
}


func NewDatabaseProxyEndpoint(scope constructs.Construct, id *string, props *DatabaseProxyEndpointProps) DatabaseProxyEndpoint {
	_init_.Initialize()

	if err := validateNewDatabaseProxyEndpointParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseProxyEndpoint{}

	_jsii_.Create(
		"cdk-extensions.rds.DatabaseProxyEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDatabaseProxyEndpoint_Override(d DatabaseProxyEndpoint, scope constructs.Construct, id *string, props *DatabaseProxyEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.rds.DatabaseProxyEndpoint",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func DatabaseProxyEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseProxyEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.rds.DatabaseProxyEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func DatabaseProxyEndpoint_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDatabaseProxyEndpoint_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.rds.DatabaseProxyEndpoint",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DatabaseProxyEndpoint_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDatabaseProxyEndpoint_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.rds.DatabaseProxyEndpoint",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseProxyEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := d.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_DatabaseProxyEndpoint) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseProxyEndpoint) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := d.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseProxyEndpoint) GetResourceNameAttribute(nameAttr *string) *string {
	if err := d.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseProxyEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
