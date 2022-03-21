// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"

	logging "github.com/runatlantis/atlantis/server/logging"
)

func AnyLoggingLogLevel() logging.LogLevel {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(logging.LogLevel))(nil)).Elem()))
	var nullValue logging.LogLevel
	return nullValue
}

func EqLoggingLogLevel(value logging.LogLevel) logging.LogLevel {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue logging.LogLevel
	return nullValue
}

func NotEqLoggingLogLevel(value logging.LogLevel) logging.LogLevel {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue logging.LogLevel
	return nullValue
}

func LoggingLogLevelThat(matcher pegomock.ArgumentMatcher) logging.LogLevel {
	pegomock.RegisterMatcher(matcher)
	var nullValue logging.LogLevel
	return nullValue
}
