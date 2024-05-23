package minecraft_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/raian621/minecraft-server-controller/minecraft"
)

func TestSaveToFile(t *testing.T) {
	serverProps := minecraft.NewServerPropertiesMonitor()
	var buffer bytes.Buffer
	err := serverProps.SaveToServerFile(
		"templates/server.properties.tmpl",
		&buffer,
	)
	if err != nil {
		t.Error(err)
	}

	fixture, err := os.ReadFile("fixtures/server.properties")
	if err != nil {
		t.Error(err)
	}

	if string(fixture) != buffer.String() {
		t.Fail()
	}
}
