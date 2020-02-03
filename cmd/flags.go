package cmd

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type Flags struct {
	ConfigFile string
}

func ReadFlags(flags interface{}) {
	var values = make(map[string]*string)
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
			values[typeField.Name] = fset.String(name, defaultValue, usage)
		}
	}
	_ = fset.Parse(os.Args[1:])
	for name, value := range values {
		valueField := v.FieldByName(name)
		typeField, _ := t.FieldByName(name)
		isRequired, hasRequired := typeField.Tag.Lookup("required")
		if hasRequired && isRequired == "true" && *value == "" {
			panic(fmt.Sprintf("%s is missing", name))
		}
		valueField.SetString(*value)
	}
}
