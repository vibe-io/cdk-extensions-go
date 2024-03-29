package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

type SecurityHubFinding interface {
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
	Bind(_scope constructs.IConstruct) *[]IssueTrigger
	BuildDescription() awsstepfunctions.Chain
	BuildLogging() *awsstepfunctions.LogOptions
	BuildRemediation() awsstepfunctions.Chain
	BuildResources() awsstepfunctions.Chain
	BuildSeverityMap() awsstepfunctions.Chain
	BuildUrl() awsstepfunctions.Chain
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
	RegisterIssueTrigger(id *string, options *SecurityHubFindingEventOptions) IssueTrigger
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for SecurityHubFinding
type jsiiProxy_SecurityHubFinding struct {
	jsiiProxy_IssuePluginBase
	jsiiProxy_IIssueParser
}

func (j *jsiiProxy_SecurityHubFinding) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityHubFinding) Handler() awsstepfunctions.IStateMachine {
	var returns awsstepfunctions.IStateMachine
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityHubFinding) Logging() *StateMachineLogging {
	var returns *StateMachineLogging
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityHubFinding) MatchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityHubFinding) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityHubFinding) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityHubFinding) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityHubFinding) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityHubFinding) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityHubFinding) Triggers() *[]IssueTrigger {
	var returns *[]IssueTrigger
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}


func NewSecurityHubFinding(scope constructs.IConstruct, id *string, props *SecurityHubFindingProps) SecurityHubFinding {
	_init_.Initialize()

	if err := validateNewSecurityHubFindingParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityHubFinding{}

	_jsii_.Create(
		"cdk-extensions.alerting.SecurityHubFinding",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSecurityHubFinding_Override(s SecurityHubFinding, scope constructs.IConstruct, id *string, props *SecurityHubFindingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.alerting.SecurityHubFinding",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func SecurityHubFinding_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecurityHubFinding_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.SecurityHubFinding",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func SecurityHubFinding_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateSecurityHubFinding_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.SecurityHubFinding",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func SecurityHubFinding_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateSecurityHubFinding_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.SecurityHubFinding",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func SecurityHubFinding_MATCH_TYPE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk-extensions.alerting.SecurityHubFinding",
		"MATCH_TYPE",
		&returns,
	)
	return returns
}

func SecurityHubFinding_SEVERITIES() *[]SecurityHubSeverity {
	_init_.Initialize()
	var returns *[]SecurityHubSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.SecurityHubFinding",
		"SEVERITIES",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SecurityHubFinding) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_SecurityHubFinding) Bind(_scope constructs.IConstruct) *[]IssueTrigger {
	if err := s.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *[]IssueTrigger

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityHubFinding) BuildDescription() awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"buildDescription",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityHubFinding) BuildLogging() *awsstepfunctions.LogOptions {
	var returns *awsstepfunctions.LogOptions

	_jsii_.Invoke(
		s,
		"buildLogging",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityHubFinding) BuildRemediation() awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"buildRemediation",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityHubFinding) BuildResources() awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"buildResources",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityHubFinding) BuildSeverityMap() awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"buildSeverityMap",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityHubFinding) BuildUrl() awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"buildUrl",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityHubFinding) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityHubFinding) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (s *jsiiProxy_SecurityHubFinding) GetResourceNameAttribute(nameAttr *string) *string {
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

func (s *jsiiProxy_SecurityHubFinding) RegisterIssueTrigger(id *string, options *SecurityHubFindingEventOptions) IssueTrigger {
	if err := s.validateRegisterIssueTriggerParameters(id, options); err != nil {
		panic(err)
	}
	var returns IssueTrigger

	_jsii_.Invoke(
		s,
		"registerIssueTrigger",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityHubFinding) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

