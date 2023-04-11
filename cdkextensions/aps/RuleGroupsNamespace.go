package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsaps"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/aps/internal"
)

// Amazon Managed Service for Prometheus allows for the configuration of rules that configure alerting and precomputation of frequently needed expressions.
//
// These rules are added to a workspace using configurations that define one or
// more resource groups. Eache group can contain one or more rules and
// configures the frequency that its rules should be evaluated.
//
// You can have multiple configurations per workspace. Each configuration is
// contained in a separate _namespace_. Having multiple configuration lets you
// import existing Prometheus rules files to a workspace without having to
// change or combine them. Different rule group namespaces can also have
// different tags.
// See: [Recording rules and alerting rules](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-Ruler.html)
//
type RuleGroupsNamespace interface {
	awscdk.Resource
	IRuleGroupsNamespace
	// The rules definition file for this namespace.
	// See: [RuleGroupsNamespace Data](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-rulegroupsnamespace.html#cfn-aps-rulegroupsnamespace-data)
	//
	Configuration() IRuleGroupConfiguration
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The name of the rule groups namespace.
	// See: [RuleGroupsNamespace Name](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-rulegroupsnamespace.html#cfn-aps-rulegroupsnamespace-name)
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
	// The underlying RuleGroupsNamespace CloudFormation resource.
	// See: [AWS::APS::RuleGroupsNamespace](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-rulegroupsnamespace.html)
	//
	Resource() awsaps.CfnRuleGroupsNamespace
	// The Amazon Resource Name (ARN) of the APS rule groups namespace.
	RulesGroupsNamespaceArn() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The APS workspace that contains this rule groups namespace.
	// See: [RuleGroupsNamespace Workspace](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-rulegroupsnamespace.html#cfn-aps-rulegroupsnamespace-workspace)
	//
	Workspace() IWorkspace
	// Adds a new rule group to the configuration.
	//
	// Returns: The rule group that was added.
	AddRuleGroup(id *string, options *RuleGroupProps) RuleGroup
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

// The jsii proxy struct for RuleGroupsNamespace
type jsiiProxy_RuleGroupsNamespace struct {
	internal.Type__awscdkResource
	jsiiProxy_IRuleGroupsNamespace
}

func (j *jsiiProxy_RuleGroupsNamespace) Configuration() IRuleGroupConfiguration {
	var returns IRuleGroupConfiguration
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleGroupsNamespace) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleGroupsNamespace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleGroupsNamespace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleGroupsNamespace) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleGroupsNamespace) Resource() awsaps.CfnRuleGroupsNamespace {
	var returns awsaps.CfnRuleGroupsNamespace
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleGroupsNamespace) RulesGroupsNamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rulesGroupsNamespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleGroupsNamespace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleGroupsNamespace) Workspace() IWorkspace {
	var returns IWorkspace
	_jsii_.Get(
		j,
		"workspace",
		&returns,
	)
	return returns
}


// Creates a new instance of the RuleGroupsNamespace class.
func NewRuleGroupsNamespace(scope constructs.IConstruct, id *string, props *RuleGroupsNamespaceProps) RuleGroupsNamespace {
	_init_.Initialize()

	if err := validateNewRuleGroupsNamespaceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RuleGroupsNamespace{}

	_jsii_.Create(
		"cdk-extensions.aps.RuleGroupsNamespace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new instance of the RuleGroupsNamespace class.
func NewRuleGroupsNamespace_Override(r RuleGroupsNamespace, scope constructs.IConstruct, id *string, props *RuleGroupsNamespaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.aps.RuleGroupsNamespace",
		[]interface{}{scope, id, props},
		r,
	)
}

// Imports an existing APS rule groups namespace by specifying its Amazon Resource Name (ARN).
//
// Returns: An object representing the imported APS rule groups namespace.
func RuleGroupsNamespace_FromRuleGroupsNamespaceArn(scope constructs.IConstruct, id *string, arn *string) IRuleGroupsNamespace {
	_init_.Initialize()

	if err := validateRuleGroupsNamespace_FromRuleGroupsNamespaceArnParameters(scope, id, arn); err != nil {
		panic(err)
	}
	var returns IRuleGroupsNamespace

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.RuleGroupsNamespace",
		"fromRuleGroupsNamespaceArn",
		[]interface{}{scope, id, arn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func RuleGroupsNamespace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRuleGroupsNamespace_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.RuleGroupsNamespace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func RuleGroupsNamespace_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateRuleGroupsNamespace_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.RuleGroupsNamespace",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func RuleGroupsNamespace_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateRuleGroupsNamespace_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.RuleGroupsNamespace",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RuleGroupsNamespace) AddRuleGroup(id *string, options *RuleGroupProps) RuleGroup {
	if err := r.validateAddRuleGroupParameters(id, options); err != nil {
		panic(err)
	}
	var returns RuleGroup

	_jsii_.Invoke(
		r,
		"addRuleGroup",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RuleGroupsNamespace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := r.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_RuleGroupsNamespace) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RuleGroupsNamespace) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := r.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RuleGroupsNamespace) GetResourceNameAttribute(nameAttr *string) *string {
	if err := r.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RuleGroupsNamespace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
