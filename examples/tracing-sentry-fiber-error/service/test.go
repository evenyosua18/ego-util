package service

import (
	"context"
	"github.com/evenyosua18/ego-util/codes"
	"github.com/evenyosua18/ego-util/examples/tracing-sentry-fiber-error/db"
	"github.com/evenyosua18/ego-util/tracing"
)

func TestService(ctx context.Context, id string) error {
	//tracing
	sp := tracing.StartChild(ctx, id)
	defer tracing.Close(sp)

	//call db
	if err := db.CheckIdExist(tracing.Context(sp), id); err != nil {
		return tracing.LogError(sp, codes.Wrap(err))
	}

	return nil
}
