package app

import (
	"flag"
	"flying-castle/validation"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"time"
)

type Flags struct {
	ConfigFile string
}

type AppFlags interface {
	// Implement this method to check complex flag constraints
	// such as when 2 fields are mutually exclusive but one of them is required,
	// when you want to match values against a regex
	// or if you want to compare several flag values
	Validate()
}

type ValueSetter func(fieldValue reflect.Value, value reflect.Value)

var typeMap = make(map[reflect.Kind]ValueSetter)

func init() {
	typeMap[reflect.Bool] = func(fieldValue reflect.Value, value reflect.Value) {
		fieldValue.SetBool(value.Bool())
	}
	typeMap[reflect.String] = func(fieldValue reflect.Value, value reflect.Value) {
		fieldValue.SetString(value.String())
	}
	var intSetter = func(fieldValue reflect.Value, value reflect.Value) {
		fieldValue.SetInt(value.Int())
	}
	typeMap[reflect.Int] = intSetter
	typeMap[reflect.Int8] = intSetter
	typeMap[reflect.Int16] = intSetter
	typeMap[reflect.Int32] = intSetter
	typeMap[reflect.Int64] = intSetter
	var uintSetter = func(fieldValue reflect.Value, value reflect.Value) {
		fieldValue.SetUint(value.Uint())
	}
	typeMap[reflect.Uint] = uintSetter
	typeMap[reflect.Uint8] = uintSetter
	typeMap[reflect.Uint16] = uintSetter
	typeMap[reflect.Uint32] = uintSetter
	typeMap[reflect.Uint64] = uintSetter
}

func parseFlags(fset *flag.FlagSet, flags AppFlags, values map[string]interface{}) {
	var t = reflect.TypeOf(flags).Elem()
	var v = reflect.Indirect(reflect.ValueOf(flags))
	_ = fset.Parse(os.Args[1:])
	for name, value := range values {
		valueField := v.FieldByName(name)
		typeField, _ := t.FieldByName(name)
		isRequired, hasRequired := typeField.Tag.Lookup("required")
		realValue := reflect.Indirect(reflect.ValueOf(value))
		if hasRequired && isRequired == "true" && (!realValue.IsValid() || realValue.String() == "") {
			panic(fmt.Sprintf("-%s is missing", typeField.Tag.Get("flag")))
		}
		valueField.Set(realValue)
	}
	flags.Validate()
}

func readFlags(fset *flag.FlagSet, flags AppFlags, values map[string]interface{}) {
	var v = reflect.Indirect(reflect.ValueOf(flags))
	var t = reflect.TypeOf(v)
	for i := 0; i < t.NumField(); i++ {
		typeField := t.Field(i)
		valueField := v.Field(i)
		tag := typeField.Tag
		if name, ok := tag.Lookup("flag"); ok {
			var usage = tag.Get("usage")
			var defaultValue = tag.Get("default")
			switch typeField.Type.Kind() {
			case reflect.Struct:
				readFlags(fset, valueField.Interface().(AppFlags), values)
				break
			case reflect.String:
				values[typeField.Name] = fset.String(name, defaultValue, usage)
				break
			case reflect.Bool:
				d, err := strconv.ParseBool(defaultValue)
				if err != nil {
					d = false
				}
				values[typeField.Name] = fset.Bool(name, d, usage)
				break
			case reflect.Int:
			case reflect.Int8:
			case reflect.Int16:
			case reflect.Int32:
				d, err := strconv.ParseInt(defaultValue, 10, 0)
				if err != nil {
					d = 0
				}
				values[typeField.Name] = fset.Int(name, int(d), usage)
				break
			case reflect.Int64:
				d, err := strconv.ParseInt(defaultValue, 10, 64)
				if err != nil {
					panic(err)
				}
				values[typeField.Name] = fset.Int64(name, d, usage)
				break
			case reflect.Float64:
				d, err := strconv.ParseFloat(defaultValue, 64)
				if err != nil {
					panic(err)
				}
				values[typeField.Name] = fset.Float64(name, d, usage)
				break
			}
		}
	}
}

func ReadFlags(flags AppFlags) {
	var values = make(map[string]interface{})
	var t = reflect.TypeOf(flags).Elem()
	var v = reflect.ValueOf(flags).Elem()
	var fset = flag.NewFlagSet(t.Name(), flag.ContinueOnError)
	fset.SetOutput(ioutil.Discard)
	for i := 0; i < t.NumField(); i++ {
		typeField := t.Field(i)
		tag := typeField.Tag
		if name, ok := tag.Lookup("flag"); ok {
			var usage = tag.Get("usage")
			var defaultValue = tag.Get("default")
			switch typeField.Type.Kind() {
			case reflect.String:
				values[typeField.Name] = fset.String(name, defaultValue, usage)
				break
			case reflect.Bool:
				d, err := strconv.ParseBool(defaultValue)
				if err != nil {
					d = false
				}
				values[typeField.Name] = fset.Bool(name, d, usage)
				break
			case reflect.Int:
			case reflect.Int8:
			case reflect.Int16:
			case reflect.Int32:
				d, err := strconv.ParseInt(defaultValue, 10, 0)
				if err != nil {
					d = 0
				}
				values[typeField.Name] = fset.Int(name, int(d), usage)
				break
			case reflect.Int64:
				if typeField.Type.String() == "time.Duration" {
					d, err := time.ParseDuration(defaultValue)
					if err != nil {
						panic(err)
					}
					values[typeField.Name] = fset.Duration(name, d, usage)
					break
				}
				d, err := strconv.ParseInt(defaultValue, 10, 64)
				if err != nil {
					panic(err)
				}
				values[typeField.Name] = fset.Int64(name, d, usage)
				break
			case reflect.Float64:
				d, err := strconv.ParseFloat(defaultValue, 64)
				if err != nil {
					panic(err)
				}
				values[typeField.Name] = fset.Float64(name, d, usage)
				break
			}
		}
	}
	_ = fset.Parse(os.Args[1:])
	for name, value := range values {
		valueField := v.FieldByName(name)
		typeField, _ := t.FieldByName(name)
		if realValue, ok := validation.ValidateField(valueField, typeField, value); ok {
			valueField.Set(realValue)
		} else {
			panic(fmt.Sprintf("%s: missing or invalid value", name))
		}
	}
	flags.Validate()
}
