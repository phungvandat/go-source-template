package user

import (
	"context"

	iom "github.com/phungvandat/source-template/model/service/user"
	"github.com/phungvandat/source-template/utils/errs"
)

func (mw validMw) FindOne(ctx context.Context, in *iom.FindOneSvcIn) (*iom.FindOneSvcOut, error) {
	if in.ID.IsZero() {
		return nil, mw.eTracer.Trace(errs.ErrIDIsInvalid)
	}
	return mw.UseCase.FindOne(ctx, in)
}
