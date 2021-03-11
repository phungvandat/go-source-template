package redisutil

import (
	"context"
	"encoding/json"
	"reflect"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/phungvandat/source-template/utils/errs"
)

// SetRedisKeys with multi keys
func SetRedisKeys(ctx context.Context, client *redis.Client, mKeyVal map[string]interface{}, exp time.Duration) error {
	if len(mKeyVal) == 0 {
		return nil
	}

	var pairs []interface{}
	for key := range mKeyVal {
		var (
			in = mKeyVal[key]

			val, err = getValueToSet(in)
		)
		if err != nil {
			return err
		}

		pairs = append(pairs, key, val)
	}

	err := client.MSet(ctx, pairs...).Err()
	if err != nil {
		return err
	}

	return nil
}

// SetWithKey set a key
func SetWithKey(ctx context.Context, client *redis.Client, key string, val interface{}, exp time.Duration) error {
	setVal, err := getValueToSet(val)
	if err != nil {
		return err
	}

	setStatus := client.Set(ctx, key, setVal, exp)
	if setStatus == nil {
		return errs.ErrSomethingWentWrong
	}

	if setStatus.Err() != nil {
		return setStatus.Err()
	}

	return nil
}

func getValueToSet(in interface{}) (interface{}, error) {
	var (
		valKind = reflect.TypeOf(in).Kind()
		val     interface{}
		iVal    = in
	)
	if valKind == reflect.Ptr {
		iVal = reflect.ValueOf(in).Elem().Interface()
		valKind = reflect.TypeOf(iVal).Kind()
	}

	if valKind == reflect.Struct || valKind == reflect.Map {
		valByte, err := json.Marshal(iVal)
		if err != nil {
			return nil, err
		}
		val = string(valByte)
	} else {
		val = in
	}

	return val, nil
}
