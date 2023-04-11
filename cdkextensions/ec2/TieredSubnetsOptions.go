package ec2


type TieredSubnetsOptions struct {
	Provider ICidrProvider `field:"required" json:"provider" yaml:"provider"`
	TierMask *float64 `field:"optional" json:"tierMask" yaml:"tierMask"`
}

