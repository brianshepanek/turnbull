package registry

import(
	"os"
	"bytes"
	"io/ioutil"
	"testing"
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/generator/jen/helper"
	
)

const(
	testConfigPath = "/go/src/github.com/brianshepanek/turnbull/_testing/config"
	testOutputPath = "/go/src/github.com/brianshepanek/turnbull/_testing/output"
)

var(

	// Test Entity
	testEntity  = model.Entity{
		Name : "foo",
		Fields : []model.Field{
			model.Field{
				Name : "string",
				Type : "string",
			},
			model.Field{
				Name : "int",
				Type : "int",
			},
		},
		Methods : []model.Method {
			model.Method{
				Name : "add",
				Type : "add",
			},
			model.Method{
				Name : "read",
				Type : "read",
			},
			model.Method{
				Name : "browse",
				Type : "browse",
			},
		},
	}
	testField  = model.Field{
		Name : "bar",
		Type : "string",
	}
	testMethod  = model.Method{
		Name : "add",
		Type : "add",
	}
)

var (
	testRegistryGenerator RegistryGenerator

	// Test File Names
	testOutputRegistryEntityFileName = "testOutputRegistryEntityFile"

	// Test File Strings
	testOutputRegistryEntityFile string

)


func init(){

	conf, _ := config.New(testConfigPath, testOutputPath)
	formatter := formatter.New(conf)
	testHelperGenerator := helper.New(formatter)
	testRegistryGenerator = New(formatter, testHelperGenerator)

	// Test
	testOutputRegistryEntityFileFile, _ := ioutil.ReadFile("./testing/registry/expected/" + testOutputRegistryEntityFileName)

	testOutputRegistryEntityFile = string(testOutputRegistryEntityFileFile)

}

// Test Registry Entity File
func TestRegistryFile(t *testing.T){

	// Build
	statement, err := testRegistryGenerator.File(testEntity)

	// Return
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}

	f, err := os.Create("./testing/registry/created/" + testOutputRegistryEntityFileName)
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputRegistryEntityFile {
		t.Errorf(`File() failed; want "%s", got "%s"`, testOutputRegistryEntityFile, buf.String())
	}
	
}