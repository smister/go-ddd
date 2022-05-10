package integral

import (
	"context"
)

type IntegralRepoIFace interface {
	GetIntegral(ctx context.Context, integralID uint64) (*Integral, error)
	Update(ctx context.Context, integral *Integral) error
}
