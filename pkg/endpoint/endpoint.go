package endpoint

import "context"

type Endpoint func(ctx context.Context, req interface{}) (interface{}, error)
