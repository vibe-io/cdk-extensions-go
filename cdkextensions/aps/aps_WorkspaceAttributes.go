package aps


// Configuration for importing an existing APS workspace.
type WorkspaceAttributes struct {
	// The Amazon Resource Name (ARN) of the APS workspace.
	WorkspaceArn *string `field:"optional" json:"workspaceArn" yaml:"workspaceArn"`
	// The ID generated by AWS for the APS workspace.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
	// The Prometheus endpoint attribute of the workspace.
	//
	// This is the endpoint prefix without the remote_write or query API
	// appended.
	WorkspacePrometheusEndpoint *string `field:"optional" json:"workspacePrometheusEndpoint" yaml:"workspacePrometheusEndpoint"`
}

