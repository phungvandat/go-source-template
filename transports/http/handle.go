type DecodeHttpFunc func(context.Context, *http.Request) (interface{}, error)
type EndcodeHttpFunc func(ctx context.Context, res http.ResponseWriter, resData interface{}, err error) error
type Endpoint func(ctx context.Context, req interface{}) (interface{}, error)

func RouteHdl(
	decode DecodeHttpFunc,
	endpointHandler Endpoint,
	encode EndcodeHttpFunc,
) func(*http.Request, http.ResponseWriter, httprouter.Params) {
	ctx := context.Background()
	return func(req *http.Request, res http.ResponseWriter, _ httprouter.Params) {
		var (
			startMillis = time.Now()
			errMsg      = new(string)
			err         error
		)

		// Panic Recovery
		defer func() {
			if r := recover(); r != nil {
				msg := fmt.Sprintf("%v\n%v", r, string(debug.Stack()))
				log.Println(msg)
			}
		}()

		defer func() {
			err = req.Body.Close()
		}()
		requestInfoFunc := func() {
			endMillis := core.GetNowMillis()
			args := []interface{}{
				req.Method,
				req.RequestURI,
				start.Format("2006/01/02-15:04:05"),
				endMillis - core.TimeToMillis(start),
			}
		}
		// TODO: add cors when run in production
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		res.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Language")

		if req.Method == http.MethodOptions {
			res.WriteHeader(http.StatusNoContent)
			return
		}

		if req.Method != http.MethodPost {
			res.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		in, err := decode(ctx, req)
		if err != nil {
			return nil, err
		}

		out, err := endpointHandler(ctx, in)
		encode(ctx, res, out, err)
	}
}

func LogRequestInfo(req *http.Request, start time.Time, errMsg *string) {
	template := "Method: %v, Route: %v, Time: %v Duration: %v ms"
	endMillis := core.GetNowMillis()
	args := []interface{}{
		req.Method,
		req.RequestURI,
		start.Format("2006/01/02-15:04:05"),
		endMillis - core.TimeToMillis(start),
	}

	if errMsg != nil && *errMsg != "" {
		template += " Err: %v"
		args = append(args, *errMsg)
		logger.Error(template, args...)
		return
	}
	logger.Info(template, args...)
}