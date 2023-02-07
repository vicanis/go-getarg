package getarg

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
)

const tagName = "getarg"

func Decode(uv url.Values, ptr any) error {
	if ptr == nil {
		return errors.New("no object")
	}

	obj := reflect.ValueOf(ptr)

	if obj.Kind() != reflect.Ptr {
		return fmt.Errorf("expected pointer to struct, given: %s", obj.Kind())
	}

	obj = obj.Elem()

	if obj.Kind() != reflect.Struct {
		return fmt.Errorf("object type (%s) is not struct", obj.Kind())
	}

	t := obj.Type()

	if t.Kind() != reflect.Struct {
		return fmt.Errorf("object type (%s) is not struct", t.Kind())
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		v := obj.Field(i)

		if !v.IsValid() {
			return fmt.Errorf("struct field %s is not valid", f.Name)
		}

		if !v.CanSet() {
			return fmt.Errorf("struct field %s is unassignable", f.Name)
		}

		if f.Type.Kind() != reflect.String {
			return fmt.Errorf("struct field %s type (%s) should be string", f.Name, f.Type.Kind())
		}

		tag := f.Tag.Get(tagName)
		if tag == "" {
			continue
		}

		val := uv.Get(tag)
		if val == "" {
			continue
		}

		v.SetString(val)
	}

	return nil
}

func Encode(data any) (uv url.Values, err error) {
	val := reflect.ValueOf(data)
	t := val.Type()

	if t.Kind() == reflect.Ptr {
		// dereference pointer
		val = reflect.ValueOf(val)
		t = val.Type()
	}

	if t.Kind() != reflect.Struct {
		err = fmt.Errorf("expected struct, given: %s", t.Kind())
		return
	}

	uv = make(url.Values)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if f.Type.Kind() != reflect.String {
			err = fmt.Errorf("struct field %s type (%s) should be string", f.Name, f.Type.Kind())
			uv = nil
			return
		}

		tag := f.Tag.Get(tagName)
		if tag == "" {
			continue
		}

		text := val.Field(i).String()
		if text != "" {
			uv.Set(tag, text)
		}
	}

	return
}
