package gscript

import (
	"testing"
	"time"
  "fmt"
	"github.com/stretchr/testify/assert"
	"github.com/davecgh/go-spew/spew"
)

var g_file_1 = fmt.Sprintf("/tmp/%s", RandString(6))
var g_file_2 = fmt.Sprintf("/tmp/%s", RandString(6))
var g_file_3 = fmt.Sprintf("/tmp/%s", RandString(6))

func TestVMMD5(t *testing.T) {
	testScript := `
    var hash_val = "helloworld";
    var return_value = MD5(hash_val);
  `
	// "helloworld" = fc5e038d38a57032085441e7fe7010b0

	e := New()
	e.EnableLogging()
	e.CreateVM()

	e.VM.Run(testScript)
	retVal, err := e.VM.Get("return_value")
	assert.Nil(t, err)

	retValAsString, err := retVal.ToString()
	assert.Nil(t, err)

	assert.Equal(t, "fc5e038d38a57032085441e7fe7010b0", retValAsString)
}

func TestVMCopyFile(t *testing.T) {
	file_2 := g_file_1
	testScript := fmt.Sprintf(`
    var file_1 = "/etc/passwd";
    var file_2 = "%s";
    var return_value = CopyFile(file_1, file_2);
  `, file_2)

	e := New()
	e.EnableLogging()
	e.CreateVM()

	e.VM.Run(testScript)
	retVal, err := e.VM.Get("return_value")
	assert.Nil(t, err)
	retValAsString, err := retVal.ToString()
	assert.Nil(t, err)
	assert.Equal(t, "true", retValAsString)
}

func TestVMAppendFile(t *testing.T) {
	bytes := "60,104,116,109,108,62,10,32,32,60,98,111,100,121,62,10,32,32,32,32"
	testScript := fmt.Sprintf(`
    var file_1 = "%s";
		var file_2 = "%s";
    var bytes = [%s];
    var return_value1 = AppendFile(file_1, bytes);
		var return_value2 = AppendFile(file_2, bytes);
  `, g_file_1, g_file_2, bytes)

	e := New()
	e.EnableLogging()
	e.CreateVM()

	e.VM.Run(testScript)
	e.LogInfof("Function: function=%s msg='Appended local file at: %s'", CalledBy(), spew.Sdump(g_file_1))
	e.LogInfof("Function: function=%s msg='Appended local file at: %s'", CalledBy(), spew.Sdump(g_file_2))
	retVal, err := e.VM.Get("return_value1")
	assert.Nil(t, err)
	retValAsString, err := retVal.ToString()
	assert.Nil(t, err)
	assert.Equal(t, "true", retValAsString)
	retVal2, err := e.VM.Get("return_value2")
	assert.Nil(t, err)
	retValAsString2, err := retVal2.ToString()
	assert.Nil(t, err)
	assert.Equal(t, "true", retValAsString2)
}

func TestVMReplaceInFile(t *testing.T) {
	string01 := "root"
	string02 := "lol"
	testScript := fmt.Sprintf(`
    var file_1 = "%s";
    var string01 = "%s";
		var string02 = "%s";
    var return_value1 = ReplaceInFile(file_1, string01, string02);
  `, g_file_1, string01, string02)

	e := New()
	e.EnableLogging()
	e.CreateVM()

	e.VM.Run(testScript)
	retVal, err := e.VM.Get("return_value1")
	assert.Nil(t, err)
	retValAsString, err := retVal.ToString()
	assert.Nil(t, err)
	assert.Equal(t, "true", retValAsString)
}

func TestVMRetrieveFileFromURL(t *testing.T) {
  url := "https://alexlevinson.com/"
	file_3 := g_file_3
  testScript2 := fmt.Sprintf(`
	  var url = "%s";
		var file_3 = "%s";
	  var response2 = RetrieveFileFromURL(url);
	  var return_value2 = response2;
		var response3 = WriteFile(file_3, response2);
  `, url, file_3)
  e := New()
  e.EnableLogging()
  e.CreateVM()

  e.VM.Run(testScript2)
	e.LogInfof("Function: function=%s msg='wrote local file at: %s'", CalledBy(), spew.Sdump(file_3))
  retVal, err := e.VM.Get("return_value2")
  assert.Nil(t, err)
  retValAsString, err := retVal.ToString()
  assert.Nil(t, err)
  assert.Equal(t, "60,104,116,109,108,62,10,32,32,60,98,111,100,121,62,10,32,32,32,32,60,99,101,110,116,101,114,62,10,32,32,32,32,32,32,60,105,109,103,32,115,114,99,61,34,114,111,111,116,46,106,112,103,34,32,47,62,10,32,32,32,32,60,47,99,101,110,116,101,114,62,10,32,32,60,47,98,111,100,121,62,10,60,47,104,116,109,108,62,10", retValAsString)
}

func TestVMDeleteFile(t *testing.T) {
	testScript := fmt.Sprintf(`
		var file_1 = "%s";
    var return_value1 = DeleteFile(file_1);
    var file_2 = "%s";
    var return_value2 = DeleteFile(file_2);
		var file_3 = "%s";
    var return_value3 = DeleteFile(file_3);
  `, g_file_1, g_file_2, g_file_3)

	e := New()
	e.EnableLogging()
	e.CreateVM()

	e.VM.Run(testScript)
	retVal, err := e.VM.Get("return_value1")
	assert.Nil(t, err)
	retValAsString, err := retVal.ToString()
	assert.Nil(t, err)
	assert.Equal(t, "true", retValAsString)
	retVal2, err := e.VM.Get("return_value2")
	assert.Nil(t, err)
	retValAsString2, err := retVal2.ToString()
	assert.Nil(t, err)
	assert.Equal(t, "true", retValAsString2)
	retVal3, err := e.VM.Get("return_value3")
	assert.Nil(t, err)
	retValAsString3, err := retVal3.ToString()
	assert.Nil(t, err)
	assert.Equal(t, "true", retValAsString3)
}

func TestVMTimestamp(t *testing.T) {
	currTime := time.Now().Unix()

	testScript := `
    var test_time = Timestamp();
  `
	e := New()
	e.EnableLogging()
	e.CreateVM()

	e.VM.Run(testScript)
	retVal, err := e.VM.Get("test_time")
	assert.Nil(t, err)
	assert.True(t, retVal.IsNumber())
	retValAsNumber, err := retVal.ToInteger()
	assert.Nil(t, err)
	assert.True(t, (retValAsNumber >= currTime))
}

func TestExec(t *testing.T) {
	testCmd := ExecuteCommand("ls", "-lah")

	testScript := `
      var test_exec = Exec("ls", ["-lah"]);
    `
	e := New()
	e.EnableLogging()
	e.CreateVM()

	e.VM.Run(testScript)
	retVal, err := e.VM.Get("test_exec")
	assert.Nil(t, err)
	assert.True(t, retVal.IsObject())
	retValAsInterface, err := retVal.Export()
	assert.Nil(t, err)
	realRetVal := retValAsInterface.(VMExecResponse)

	assert.Equal(t, testCmd.Stdout, realRetVal.Stdout)
}

func TestCPUStats(t *testing.T) {
	//resultz := CPUStats()
	testScript := `
      var results = CPUStats();
    `
	e := New()
	e.EnableLogging()
	e.CreateVM()

	e.VM.Run(testScript)
	retVal, err := e.VM.Get("results")
	assert.Nil(t, err)
	retValAsString, err := retVal.ToString()
	assert.Nil(t, err)
	assert.Equal(t, "true", retValAsString)
}
