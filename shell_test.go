package shell_test

import (
	"reflect"
	"testing"

	"charm.land/bubbles/v2/filepicker"
)

func TestDumpDefaultKeyMap(t *testing.T) {
	km := filepicker.DefaultKeyMap()

	tkm := reflect.TypeOf(km)
	vkm := reflect.ValueOf(km)
	for i := range tkm.NumField() {
		t.Logf("%s: %v", tkm.Field(i).Name, vkm.Field(i))
	}
}
