package securityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/securityhub/internal"
)

type Standard interface {
	awscdk.Resource
	IStandard
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
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
	Resource() awssecurityhub.CfnStandard
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	StandardArn() *string
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
	DisableControl(control *string, options *DisableControlOptions)
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

// The jsii proxy struct for Standard
type jsiiProxy_Standard struct {
	internal.Type__awscdkResource
	jsiiProxy_IStandard
}

func (j *jsiiProxy_Standard) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Standard) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Standard) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Standard) Resource() awssecurityhub.CfnStandard {
	var returns awssecurityhub.CfnStandard
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Standard) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Standard) StandardArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standardArn",
		&returns,
	)
	return returns
}


func NewStandard(scope constructs.IConstruct, id *string, props *StandardProps) Standard {
	_init_.Initialize()

	if err := validateNewStandardParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Standard{}

	_jsii_.Create(
		"cdk-extensions.securityhub.Standard",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewStandard_Override(s Standard, scope constructs.IConstruct, id *string, props *StandardProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.securityhub.Standard",
		[]interface{}{scope, id, props},
		s,
	)
}

func Standard_FromStandardArn(scope constructs.IConstruct, id *string, arn *string) IStandard {
	_init_.Initialize()

	if err := validateStandard_FromStandardArnParameters(scope, id, arn); err != nil {
		panic(err)
	}
	var returns IStandard

	_jsii_.StaticInvoke(
		"cdk-extensions.securityhub.Standard",
		"fromStandardArn",
		[]interface{}{scope, id, arn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Standard_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStandard_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.securityhub.Standard",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Standard_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateStandard_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.securityhub.Standard",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Standard_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateStandard_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.securityhub.Standard",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func Standard_ARN_FORMAT() awscdk.ArnFormat {
	_init_.Initialize()
	var returns awscdk.ArnFormat
	_jsii_.StaticGet(
		"cdk-extensions.securityhub.Standard",
		"ARN_FORMAT",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_Standard) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_Standard) DisableControl(control *string, options *DisableControlOptions) {
	if err := s.validateDisableControlParameters(control, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"disableControl",
		[]interface{}{control, options},
	)
}

func (s *jsiiProxy_Standard) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Standard) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := s.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Standard) GetResourceNameAttribute(nameAttr *string) *string {
	if err := s.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Standard) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
