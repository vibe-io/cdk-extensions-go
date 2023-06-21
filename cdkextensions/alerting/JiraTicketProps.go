package alerting

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Configuration controlling how Jira tickets should be created in response to events.
type JiraTicketProps struct {
	// The AWS account ID this resource belongs to.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// ARN to deduce region and account from.
	//
	// The ARN is parsed and the account and region are taken from the ARN.
	// This should be used for imported resources.
	//
	// Cannot be supplied together with either `account` or `region`.
	EnvironmentFromArn *string `field:"optional" json:"environmentFromArn" yaml:"environmentFromArn"`
	// The value passed in by users to the physical name prop of the resource.
	//
	// - `undefined` implies that a physical name will be allocated by
	//   CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//   by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	Region *string `field:"optional" json:"region" yaml:"region"`
	Logging *StateMachineLogging `field:"optional" json:"logging" yaml:"logging"`
	Credentials awssecretsmanager.ISecret `field:"required" json:"credentials" yaml:"credentials"`
	IssueType *string `field:"required" json:"issueType" yaml:"issueType"`
	JiraUrl *string `field:"required" json:"jiraUrl" yaml:"jiraUrl"`
	PriorityMap *JiraTicketPriorityMap `field:"required" json:"priorityMap" yaml:"priorityMap"`
	Project *string `field:"required" json:"project" yaml:"project"`
	Assignee *string `field:"optional" json:"assignee" yaml:"assignee"`
	EventBus awsevents.IEventBus `field:"optional" json:"eventBus" yaml:"eventBus"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}
