package athena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsathena"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/athena/internal"
)

type WorkGroup interface {
	awscdk.Resource
	IWorkGroup
	Description() *string
	Engine() IAnalyticsEngine
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
	RecursiveDelete() *bool
	Resource() awsathena.CfnWorkGroup
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	State() WorkGroupState
	WorkGroupArn() *string
	WorkGroupCreationTime() *string
	WorkGroupEffectiveEngineVersion() *string
	WorkGroupName() *string
	AddNamedQuery(id *string, options *AddNamedQueryOptions) NamedQuery
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

// The jsii proxy struct for WorkGroup
type jsiiProxy_WorkGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IWorkGroup
}

func (j *jsiiProxy_WorkGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkGroup) Engine() IAnalyticsEngine {
	var returns IAnalyticsEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkGroup) RecursiveDelete() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"recursiveDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkGroup) Resource() awsathena.CfnWorkGroup {
	var returns awsathena.CfnWorkGroup
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkGroup) State() WorkGroupState {
	var returns WorkGroupState
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkGroup) WorkGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkGroup) WorkGroupCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workGroupCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkGroup) WorkGroupEffectiveEngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workGroupEffectiveEngineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkGroup) WorkGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workGroupName",
		&returns,
	)
	return returns
}


func NewWorkGroup(scope constructs.IConstruct, id *string, props *WorkGroupProps) WorkGroup {
	_init_.Initialize()

	if err := validateNewWorkGroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkGroup{}

	_jsii_.Create(
		"cdk-extensions.athena.WorkGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewWorkGroup_Override(w WorkGroup, scope constructs.IConstruct, id *string, props *WorkGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.athena.WorkGroup",
		[]interface{}{scope, id, props},
		w,
	)
}

func WorkGroup_FromWorkGroupArn(scope constructs.IConstruct, id *string, arn *string) IWorkGroup {
	_init_.Initialize()

	if err := validateWorkGroup_FromWorkGroupArnParameters(scope, id, arn); err != nil {
		panic(err)
	}
	var returns IWorkGroup

	_jsii_.StaticInvoke(
		"cdk-extensions.athena.WorkGroup",
		"fromWorkGroupArn",
		[]interface{}{scope, id, arn},
		&returns,
	)

	return returns
}

func WorkGroup_FromWorkGroupAttributes(scope constructs.IConstruct, id *string, attrs *WorkGroupAttributes) IWorkGroup {
	_init_.Initialize()

	if err := validateWorkGroup_FromWorkGroupAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IWorkGroup

	_jsii_.StaticInvoke(
		"cdk-extensions.athena.WorkGroup",
		"fromWorkGroupAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

func WorkGroup_FromWorkGroupName(scope constructs.IConstruct, id *string, name *string) IWorkGroup {
	_init_.Initialize()

	if err := validateWorkGroup_FromWorkGroupNameParameters(scope, id, name); err != nil {
		panic(err)
	}
	var returns IWorkGroup

	_jsii_.StaticInvoke(
		"cdk-extensions.athena.WorkGroup",
		"fromWorkGroupName",
		[]interface{}{scope, id, name},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func WorkGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.athena.WorkGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func WorkGroup_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateWorkGroup_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.athena.WorkGroup",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func WorkGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateWorkGroup_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.athena.WorkGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func WorkGroup_ARN_FORMAT() awscdk.ArnFormat {
	_init_.Initialize()
	var returns awscdk.ArnFormat
	_jsii_.StaticGet(
		"cdk-extensions.athena.WorkGroup",
		"ARN_FORMAT",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WorkGroup) AddNamedQuery(id *string, options *AddNamedQueryOptions) NamedQuery {
	if err := w.validateAddNamedQueryParameters(id, options); err != nil {
		panic(err)
	}
	var returns NamedQuery

	_jsii_.Invoke(
		w,
		"addNamedQuery",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := w.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (w *jsiiProxy_WorkGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := w.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkGroup) GetResourceNameAttribute(nameAttr *string) *string {
	if err := w.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
