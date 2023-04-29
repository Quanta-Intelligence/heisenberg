package internal

import (
	"fmt"
)

type errorCode uint

const (
	NullError errorCode = iota

	ConnectionErrorCode
	InvalidServiceErrorCode
	IncorrectServiceErrorCode

	InvalidCollectionErrorCode
	InvalidKeyErrorCode
	CorruptValueErrorCode
)

func ConnectionError(addr string, args ...interface{}) error {
	return fmt.Errorf(
		"%d: connection to address %s failed, trace: %v",
		ConnectionErrorCode,
		addr,
		args,
	)
}

func InvalidServiceError(args ...interface{}) error {
	return fmt.Errorf(
		"%d: service is invalid, trace: %v",
		InvalidServiceErrorCode,
		args,
	)
}

func IncorrectServiceError(expected Service, recieved Service, args ...interface{}) error {
	return fmt.Errorf(
		"%d: connected to incorrect service, expected %s, recieved %s, trace: %v",
		IncorrectServiceErrorCode,
		expected.String(),
		recieved.String(),
		args,
	)
}

func InvalidCollectionError(collection []byte, args ...interface{}) error {
	return fmt.Errorf(
		"%d: collection %s does not exist, trace: %v",
		InvalidCollectionErrorCode,
		string(collection),
		args,
	)
}

func InvalidKeyError(key []byte, collection []byte, args ...interface{}) error {
	return fmt.Errorf(
		"%d: key %s does not exist in collection %s, trace: %v",
		InvalidKeyErrorCode,
		string(key),
		string(collection),
		args,
	)
}

func CorruptValueError(key []byte, args ...interface{}) error {
	return fmt.Errorf(
		"%d: corrupt value at key %s, trace: %v",
		CorruptValueErrorCode,
		string(key),
		args,
	)
}
