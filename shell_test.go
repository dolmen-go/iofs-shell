package shell_test

import (
	"reflect"
	"testing"

	"github.com/charmbracelet/bubbles/filepicker"
)

func TestDumpDefaultKeyMap(t *testing.T) {
	km := filepicker.DefaultKeyMap()

	tkm := reflect.TypeOf(km)
	vkm := reflect.ValueOf(km)
	for i := 0; i < tkm.NumField(); i++ {
		t.Logf("%s: %v", tkm.Field(i).Name, vkm.Field(i))
	}
}
