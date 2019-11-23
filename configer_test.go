package configer

import (
	"path/filepath"
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
	name, err := c.getstring("name")
	if name != "sasukebo" || err != nil {
		t.Errorf("get string config error; want: sasukebo, <nil>, get %s, %v", name, err)
	}

	age, err := c.getint("age")
	if age != 25 || err != nil {
		t.Errorf("get int config error; want: 25, <nil>, get %d, %v", age, err)
	}

	man, err := c.getbool("man")
	if !man || err != nil {
		t.Errorf("get bool config error; want: true, <nil>, get %v, %v", man, err)
	}

	vf, err := c.getfloat("afloat")
	if vf != 3.1415 || err != nil {
		t.Errorf("get float64 config error; want: true, <nil>, get %f, %v", vf, err)
	}

	ep, err := c.getstring("password")
	if ep != "e23dsa3c9a7" || err != nil {
		t.Errorf("get development config error; want: e23dsa3c9a7, <nil>, get %f, %v", vf, err)
	}
}

func TestReloadConfig(t *testing.T) {
	SetConfigFileDir("./conf")
	err := ReloadConfig()
	if err != nil {
		t.Errorf("reload config error; %v", err)
	}

	hobby, err := c.getstring("hobby")
	if hobby != "coding" || err != nil {
		t.Errorf("get reload config error; want: coding, <nil>, get %s, %v", hobby, err)
	}
}
