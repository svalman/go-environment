package environment

import (
	"os"
	"reflect"
	"testing"
)

type actionFunc func()

func TestGetEnv(t *testing.T) {

	type args struct {
		key        string
		defaultVal string
	}

	var tests = []struct {
		name         string
		beforeAction actionFunc
		afterAction  actionFunc
		args         args
		want         string
	}{

		{
			name: "Простое установленное значение",
			beforeAction: func() {
				err := os.Setenv("TEST_VALUE", "OK")
				if err != nil {
					t.Errorf("Cannot setup environment var TEST_VALUE. Error: %s\n", err.Error())
					return
				}
			},
			afterAction: func() {
				err := os.Unsetenv("TEST_VALUE")
				if err != nil {
					t.Errorf("Cannot setup environment var TEST_VALUE. Error %s\n", err.Error())
					return
				}
			},
			args: args{
				key:        "TEST_VALUE",
				defaultVal: "error",
			},
			want: "OK",
		},

		{
			name:         "Значение по умолчанию",
			beforeAction: nil,
			args: args{
				key:        "",
				defaultVal: "default",
			},
			want: "default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeAction != nil {
				tt.beforeAction()
			}
			if got := GetEnv(tt.args.key, tt.args.defaultVal); got != tt.want {
				if tt.afterAction != nil {
					tt.afterAction()
				}
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
			if tt.afterAction != nil {
				tt.afterAction()
			}
		})
	}
}

func TestGetEnvAsBool(t *testing.T) {
	type args struct {
		name       string
		defaultVal bool
	}
	tests := []struct {
		name         string
		beforeAction actionFunc
		afterAction  actionFunc
		args         args
		want         bool
	}{
		{
			name: "Запрос существующей true переменной",
			beforeAction: func() {
				err := os.Setenv("TEST_VALUE", "1")
				if err != nil {
					t.Errorf("Cannot setup environment var TEST_VALUE. Error: %s\n", err.Error())
					return
				}
			},
			afterAction: func() {
				err := os.Unsetenv("TEST_VALUE")
				if err != nil {
					t.Errorf("Cannot setup environment var TEST_VALUE. Error %s\n", err.Error())
					return
				}
			},
			args: args{
				name:       "TEST_VALUE",
				defaultVal: false,
			},
			want: true,
		},

		{
			name: "Запрос существующей false переменной",
			beforeAction: func() {
				err := os.Setenv("TEST_VALUE", "0")
				if err != nil {
					t.Errorf("Cannot setup environment var TEST_VALUE. Error: %s\n", err.Error())
					return
				}
			},
			afterAction: func() {
				err := os.Unsetenv("TEST_VALUE")
				if err != nil {
					t.Errorf("Cannot setup environment var TEST_VALUE. Error %s\n", err.Error())
					return
				}
			},
			args: args{
				name:       "TEST_VALUE",
				defaultVal: true,
			},
			want: false,
		},

		{
			name: "Запрос несуществующей переменной",
			beforeAction: func() {
				err := os.Unsetenv("BOOL_VAL")
				if err != nil {
					t.Errorf("Cannot unset var. Error: %s", err.Error())
					return
				}
			},
			afterAction: nil,
			args: args{
				name:       "BOOL_VAL",
				defaultVal: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeAction != nil {
				tt.beforeAction()
			}
			if got := GetEnvAsBool(tt.args.name, tt.args.defaultVal); got != tt.want {
				if tt.afterAction != nil {
					tt.afterAction()
				}
				t.Errorf("GetEnvAsBool() = %v, want %v", got, tt.want)
			}
			if tt.afterAction != nil {
				tt.afterAction()
			}
		})
	}
}

func TestGetEnvAsInt(t *testing.T) {
	type args struct {
		name       string
		defaultVal int
	}
	tests := []struct {
		name         string
		beforeAction actionFunc
		afterAction  actionFunc
		args         args
		want         int
	}{

		{
			name: "Запрос существующей переменной",
			beforeAction: func() {
				err := os.Setenv("TEST_VALUE", "12345")
				if err != nil {
					t.Errorf("Cannot setup environment var TEST_VALUE. Error: %s\n", err.Error())
					return
				}
			},
			afterAction: func() {
				err := os.Unsetenv("TEST_VALUE")
				if err != nil {
					t.Errorf("Cannot setup environment var TEST_VALUE. Error %s\n", err.Error())
					return
				}
			},
			args: args{
				name:       "TEST_VALUE",
				defaultVal: 10,
			},
			want: 12345,
		},

		{
			name: "Запрос несуществующей переменной",
			beforeAction: func() {
				err := os.Unsetenv("TEST_VAL")
				if err != nil {
					t.Errorf("Cannot unset var. Error: %s", err.Error())
					return
				}
			},
			afterAction: nil,
			args: args{
				name:       "BOOL_VAL",
				defaultVal: 999,
			},
			want: 999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeAction != nil {
				tt.beforeAction()
			}
			if got := GetEnvAsInt(tt.args.name, tt.args.defaultVal); got != tt.want {
				if tt.afterAction != nil {
					tt.afterAction()
				}
				t.Errorf("GetEnvAsInt() = %v, want %v", got, tt.want)
			}
			if tt.afterAction != nil {
				tt.afterAction()
			}
		})
	}
}

func TestGetEnvAsSlice(t *testing.T) {
	type args struct {
		name       string
		defaultVal []string
		sep        string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvAsSlice(tt.args.name, tt.args.defaultVal, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnvAsSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
