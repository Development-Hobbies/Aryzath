package builder

import "github.com/vindecodex/Aryzath/builder/builder/product"

type CrystalBuilder struct {
	RightArm  string
	LeftArm   string
	Legs      string
	Processor string
}

func (c *CrystalBuilder) SetRightArm() {
	c.RightArm = "Right Crystal Arm"
}
func (c *CrystalBuilder) SetLeftArm() {
	c.LeftArm = "Left Crystal Arm"
}
func (c *CrystalBuilder) SetLegs() {
	c.Legs = "Crystal Legs"
}
func (c *CrystalBuilder) SetProcessor() {
	c.Processor = "AMD Ryzeen"
}
func (c *CrystalBuilder) BuildCyborg() product.Cyborg {
	return product.Cyborg{
		RightArm:  c.RightArm,
		LeftArm:   c.LeftArm,
		Legs:      c.Legs,
		Processor: c.Processor,
	}
}
