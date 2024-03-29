package securityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/securityhub/internal"
)

type Hub interface {
	awscdk.Resource
	IHub
	AutoEnableControls() *bool
	ConsolidatedFindings() *bool
	ControlFindingGenerator() ControlFindingGenerator
	EnableDefaultStandards() *bool
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	HubArn() *string
	HubName() *string
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
	Resource() awssecurityhub.CfnHub
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

// The jsii proxy struct for Hub
type jsiiProxy_Hub struct {
	internal.Type__awscdkResource
	jsiiProxy_IHub
}

func (j *jsiiProxy_Hub) AutoEnableControls() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoEnableControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Hub) ConsolidatedFindings() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"consolidatedFindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Hub) ControlFindingGenerator() ControlFindingGenerator {
	var returns ControlFindingGenerator
	_jsii_.Get(
		j,
		"controlFindingGenerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Hub) EnableDefaultStandards() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableDefaultStandards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Hub) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Hub) HubArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Hub) HubName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Hub) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Hub) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Hub) Resource() awssecurityhub.CfnHub {
	var returns awssecurityhub.CfnHub
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Hub) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewHub(scope constructs.IConstruct, id *string, props *HubProps) Hub {
	_init_.Initialize()

	if err := validateNewHubParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Hub{}

	_jsii_.Create(
		"cdk-extensions.securityhub.Hub",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewHub_Override(h Hub, scope constructs.IConstruct, id *string, props *HubProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.securityhub.Hub",
		[]interface{}{scope, id, props},
		h,
	)
}

func Hub_FromHubArn(scope constructs.IConstruct, id *string, arn *string) IHub {
	_init_.Initialize()

	if err := validateHub_FromHubArnParameters(scope, id, arn); err != nil {
		panic(err)
	}
	var returns IHub

	_jsii_.StaticInvoke(
		"cdk-extensions.securityhub.Hub",
		"fromHubArn",
		[]interface{}{scope, id, arn},
		&returns,
	)

	return returns
}

func Hub_FromHubAttributes(scope constructs.IConstruct, id *string, attrs *HubAttributes) IHub {
	_init_.Initialize()

	if err := validateHub_FromHubAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IHub

	_jsii_.StaticInvoke(
		"cdk-extensions.securityhub.Hub",
		"fromHubAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

func Hub_FromHubName(scope constructs.IConstruct, id *string, name *string) IHub {
	_init_.Initialize()

	if err := validateHub_FromHubNameParameters(scope, id, name); err != nil {
		panic(err)
	}
	var returns IHub

	_jsii_.StaticInvoke(
		"cdk-extensions.securityhub.Hub",
		"fromHubName",
		[]interface{}{scope, id, name},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Hub_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHub_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.securityhub.Hub",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Hub_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateHub_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.securityhub.Hub",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Hub_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateHub_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.securityhub.Hub",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func Hub_ARN_FORMAT() awscdk.ArnFormat {
	_init_.Initialize()
	var returns awscdk.ArnFormat
	_jsii_.StaticGet(
		"cdk-extensions.securityhub.Hub",
		"ARN_FORMAT",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_Hub) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := h.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (h *jsiiProxy_Hub) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_Hub) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := h.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_Hub) GetResourceNameAttribute(nameAttr *string) *string {
	if err := h.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_Hub) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

