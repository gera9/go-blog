package utils

import "time"

func ExponentialBackOff[T any](attemps int, operation func() (T, error)) (data T, err error) {
	for i := range attemps {
		data, err = operation()
		if err == nil {
			return
		}

		time.Sleep(time.Second<<i + 1)
	}

	return
}

func Retry[T any](operation func() (T, error)) (T, error) {
	return ExponentialBackOff(3, operation)
}
