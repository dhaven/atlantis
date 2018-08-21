package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"
	logging "github.com/runatlantis/atlantis/server/logging"
)

func AnyPtrToLoggingSimpleLogger() *logging.SimpleLogger {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*logging.SimpleLogger))(nil)).Elem()))
	var nullValue *logging.SimpleLogger
	return nullValue
}

func EqPtrToLoggingSimpleLogger(value *logging.SimpleLogger) *logging.SimpleLogger {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *logging.SimpleLogger
	return nullValue
}