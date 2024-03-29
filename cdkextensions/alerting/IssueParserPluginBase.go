package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

type IssueParserPluginBase interface {
	IssuePluginBase
	IIssueParser
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	Handler() awsstepfunctions.IStateMachine
	Logging() *StateMachineLogging
	MatchType() *string
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
	Timeout() awscdk.Duration
	Triggers() *[]IssueTrigger
	AddDefaultTrigger() IssueTrigger
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
	Bind(_node constructs.IConstruct) *[]IssueTrigger
	BuildLogging() *awsstepfunctions.LogOptions
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

// The jsii proxy struct for IssueParserPluginBase
type jsiiProxy_IssueParserPluginBase struct {
	jsiiProxy_IssuePluginBase
	jsiiProxy_IIssueParser
}

func (j *jsiiProxy_IssueParserPluginBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueParserPluginBase) Handler() awsstepfunctions.IStateMachine {
	var returns awsstepfunctions.IStateMachine
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueParserPluginBase) Logging() *StateMachineLogging {
	var returns *StateMachineLogging
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueParserPluginBase) MatchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueParserPluginBase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueParserPluginBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueParserPluginBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueParserPluginBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueParserPluginBase) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueParserPluginBase) Triggers() *[]IssueTrigger {
	var returns *[]IssueTrigger
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}


func NewIssueParserPluginBase_Override(i IssueParserPluginBase, scope constructs.IConstruct, id *string, props *IssueParserPluginBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.alerting.IssueParserPluginBase",
		[]interface{}{scope, id, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func IssueParserPluginBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIssueParserPluginBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.IssueParserPluginBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func IssueParserPluginBase_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateIssueParserPluginBase_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.IssueParserPluginBase",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func IssueParserPluginBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateIssueParserPluginBase_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.IssueParserPluginBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueParserPluginBase) AddDefaultTrigger() IssueTrigger {
	var returns IssueTrigger

	_jsii_.Invoke(
		i,
		"addDefaultTrigger",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueParserPluginBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IssueParserPluginBase) Bind(_node constructs.IConstruct) *[]IssueTrigger {
	if err := i.validateBindParameters(_node); err != nil {
		panic(err)
	}
	var returns *[]IssueTrigger

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_node},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueParserPluginBase) BuildLogging() *awsstepfunctions.LogOptions {
	var returns *awsstepfunctions.LogOptions

	_jsii_.Invoke(
		i,
		"buildLogging",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueParserPluginBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueParserPluginBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (i *jsiiProxy_IssueParserPluginBase) GetResourceNameAttribute(nameAttr *string) *string {
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

func (i *jsiiProxy_IssueParserPluginBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

