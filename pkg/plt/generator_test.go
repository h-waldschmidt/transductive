package plt

import (
	"testing"
)

func TestCreateNormalDistributionNormal(t *testing.T) {
	distribution := CreateNormalDistribution(0, 0.1, 10)

	if len(distribution) != 10 {
		t.Error("distribution has unexepected number of points")
	}
}
