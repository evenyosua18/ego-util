package db

import (
	"context"
	"github.com/evenyosua18/ego-util/tracing"
)

func CheckIdExist(ctx context.Context, id string) error {
	//tracing
	sp := tracing.StartChild(ctx, id)
	defer tracing.Close(sp)

	//all process here

	return nil
}
