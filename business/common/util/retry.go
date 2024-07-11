package util

import (
	"github.com/avast/retry-go"
	"github.com/dubbogo/gost/log/logger"
)

func Retry(fun func() error) {
	err := retry.Do(fun, RetryStrategy...)
	if err != nil {
		logger.Error("retry failed", err)
	}
}
