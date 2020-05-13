package configer

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestGetRealConfigDir(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	workdir := filepath.Dir(filename)
	want := filepath.Join(workdir, configFileDir)
	result := getRealConfigDir()

	if result != want {
		t.Errorf("get real config file dir error; want %s, get %s", want, result)
	}
}

func TestConfigGet(t *testing.T) {
	name := GetEnv("name")
	fmt.Println(name)
	if _, ok := name.(string); !ok {
		vt := reflect.TypeOf(name)
		t.Errorf("expect string but got %s\n", vt.Name())
	}

	boolV := GetEnv("bool")
	fmt.Println(boolV)
	if _, ok := boolV.(bool); !ok {
		vt := reflect.TypeOf(boolV)
		t.Errorf("expect bool but got %s\n", vt.Name())
	}

	intV := GetEnv("int")
	fmt.Println(intV)
	if _, ok := intV.(int); !ok {
		vt := reflect.TypeOf(intV)
		t.Errorf("expect int but got %s\n", vt.Name())
	}

	floatV := GetEnv("float")
	fmt.Println(floatV)
	if _, ok := floatV.(float64); !ok {
		vt := reflect.TypeOf(floatV)
		t.Errorf("expect int but got %s\n", vt.Name())
	}

	GetEnv("hello")
}
