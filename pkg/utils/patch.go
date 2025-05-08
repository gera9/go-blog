package utils

import (
	"errors"
	"reflect"
)

var (
	ErrInvalidPatchType = errors.New("invalid patch type")
	ErrModelNotPtr      = errors.New("model is not a pointer")
	ErrInvalidModelType = errors.New("invalid model type")
	ErrDiffStructs      = errors.New("both are different structs")
)

func PatchModel(patch, model any) error {
	patchV := reflect.ValueOf(patch)
	patchT := patchV.Type()

	if patchT.Kind() != reflect.Struct {
		return ErrInvalidPatchType
	}

	modelV := reflect.ValueOf(model)
	if modelV.Kind() != reflect.Pointer || modelV.IsNil() {
		return ErrModelNotPtr
	}

	modelV = modelV.Elem()
	modelT := modelV.Type()

	if modelT.Kind() != reflect.Struct {
		return ErrInvalidModelType
	}

	// Ensure structs are of the same type
	if patchT != modelT {
		return ErrDiffStructs
	}

	// Iterate over fields
	for i := range patchT.NumField() {
		patchField := patchT.Field(i)
		if !patchField.IsExported() {
			continue
		}

		patchValue := patchV.Field(i)
		if patchValue.IsZero() {
			continue
		}

		modelValue := modelV.Field(i)

		// Handle pointer fields
		if patchField.Type.Kind() == reflect.Pointer {
			patchValue = patchValue.Elem()

			if modelValue.IsNil() {
				newVal := reflect.New(patchValue.Type())
				modelV.Field(i).Set(newVal)
				modelValue = newVal.Elem()
			} else {
				modelValue = modelValue.Elem()
			}
		}

		// Handle nested structs (except time.Time)
		if patchValue.Kind() == reflect.Struct && patchField.Type.String() != "time.Time" {
			PatchModel(patchValue.Interface(), modelValue.Addr().Interface())
			continue
		}

		// Only set if values are different
		if !patchValue.Comparable() || !patchValue.Equal(modelValue) {
			modelValue.Set(patchValue)
		}
	}

	return nil
}
