package glue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Job bookmarks encryption configuration.
type JobBookmarksEncryption struct {
	// Encryption mode.
	Mode JobBookmarksEncryptionMode `field:"required" json:"mode" yaml:"mode"`
	// The KMS key to be used to encrypt the data.
	// Default: A key will be created if one is not provided.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

