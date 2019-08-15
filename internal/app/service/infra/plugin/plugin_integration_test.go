package plugin

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"proctor/internal/app/service/infra/config"
	"testing"
)

type context interface {
	setUp(t *testing.T)
	tearDown()
	instance() *testContext
}

type testContext struct {
	goPlugin GoPlugin
}

func (context *testContext) setUp(t *testing.T) {
	value, available := os.LookupEnv("ENABLE_INTEGRATION_TEST")
	if available != true || value != "true" {
		t.SkipNow()
	}

	println("====================  " + os.Getenv("PROCTOR_AUTH_PLUGIN_BINARY"))
	println("====================  " + os.Getenv("PROCTOR_AUTH_PLUGIN_BINARY"))
	println("====================  " + os.Getenv("PROCTOR_AUTH_PLUGIN_BINARY"))
	println("====================  " + os.Getenv("PROCTOR_AUTH_PLUGIN_BINARY"))
	println("xxxxxxxxxxxxxxxxxxxx  " + config.Config().AuthPluginBinary)
	println("xxxxxxxxxxxxxxxxxxxx  " + config.Config().AuthPluginBinary)
	println("xxxxxxxxxxxxxxxxxxxx  " + config.Config().AuthPluginBinary)
	println("xxxxxxxxxxxxxxxxxxxx  " + config.Config().AuthPluginBinary)

	context.goPlugin = NewGoPlugin()
	assert.NotNil(t, context.goPlugin)
}

func (context *testContext) tearDown() {
}

func (context *testContext) instance() *testContext {
	return context
}

func newContext() context {
	ctx := &testContext{}
	return ctx
}

func TestGoPlugin_LoadPluginFailed(t *testing.T) {
	ctx := newContext()
	ctx.setUp(t)

	binary := "non-existing-binary"
	raw, err := ctx.instance().goPlugin.Load(binary, config.Config().AuthPluginExported)
	assert.EqualError(t, err, fmt.Sprintf("failed to load plugin binary from location: %s", binary))
	assert.Nil(t, raw)
}

func TestGoPlugin_LoadExportedFailed(t *testing.T) {
	ctx := newContext()
	ctx.setUp(t)

	exportedName := "non-existing-exported"
	raw, err := ctx.instance().goPlugin.Load(config.Config().AuthPluginBinary, exportedName)
	assert.EqualError(t, err, fmt.Sprintf("failed to Lookup plugin binary from location: %s with Exported Name: %s", config.Config().AuthPluginBinary, exportedName))
	assert.Nil(t, raw)
}

func TestGoPlugin_LoadSuccessfully(t *testing.T) {
	ctx := newContext()
	ctx.setUp(t)

	raw, err := ctx.instance().goPlugin.Load(config.Config().AuthPluginBinary, config.Config().AuthPluginExported)
	assert.NoError(t, err)
	assert.NotNil(t, raw)
}

func TestGoPlugin_ShitShit(t *testing.T) {
	ctx := newContext()
	ctx.setUp(t)

	assert.True(t, true)
}
