package types

import "context"

type SnowflakeService interface {
	GenerateID(context.Context) int64
}
