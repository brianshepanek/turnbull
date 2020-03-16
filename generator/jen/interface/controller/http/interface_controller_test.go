package generator

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

var (
	testControllerGenerator *controllerGenerator

	testOutputScaffoldInterfaceControllerFileName = "testOutputScaffoldInterfaceControllerFile"
	testOutputScaffoldInterfaceControllerStructName = "testOutputScaffoldInterfaceControllerStruct"
	testOutputScaffoldInterfaceControllerConstructorFunctionName = "testOutputScaffoldInterfaceControllerConstructorFunction"
	testOutputScaffoldInterfaceControllerMethodName = "testOutputScaffoldInterfaceControllerMethod"
	
	testOutputScaffoldInterfaceControllerFile string
	testOutputScaffoldInterfaceControllerStruct string
	testOutputScaffoldInterfaceControllerConstructorFunction string
	testOutputScaffoldInterfaceControllerMethod string

	testEntity  = model.Entity{
		Name : "foo",
		JSON : true,
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
				Name : "count",
				Type : "count",
			},
			model.Method{
				Name : "browse",
				Type : "browse",
			},
			model.Method{
				Name : "read",
				Type : "read",
			},
			model.Method{
				Name : "edit",
				Type : "edit",
			},
			model.Method{
				Name : "add",
				Type : "add",
			},
			model.Method{
				Name : "delete",
				Type : "delete",
			},
			
		},
	}

	testMethod  = model.Method{
		Name : "add",
		Type : "add",
	}

)

func init(){
	conf, _ := config.New(testConfigPath, testOutputPath)
	formatter := formatter.New(conf)
	testHelperGenerator := helper.New(formatter)
	testControllerGenerator = New(conf, formatter, testHelperGenerator)

	testOutputScaffoldInterfaceControllerFileFile, _ := ioutil.ReadFile("./testing/interface/controller/expected/" + testOutputScaffoldInterfaceControllerFileName)
	testOutputScaffoldInterfaceControllerStructFile, _ := ioutil.ReadFile("./testing/interface/controller/expected/" + testOutputScaffoldInterfaceControllerStructName)
	testOutputScaffoldInterfaceControllerConstructorFunctionFile, _ := ioutil.ReadFile("./testing/interface/controller/expected/" + testOutputScaffoldInterfaceControllerConstructorFunctionName)
	testOutputScaffoldInterfaceControllerMethodFile, _ := ioutil.ReadFile("./testing/interface/controller/expected/" + testOutputScaffoldInterfaceControllerMethodName)
	
	testOutputScaffoldInterfaceControllerFile = string(testOutputScaffoldInterfaceControllerFileFile)
	testOutputScaffoldInterfaceControllerStruct = string(testOutputScaffoldInterfaceControllerStructFile)
	testOutputScaffoldInterfaceControllerConstructorFunction = string(testOutputScaffoldInterfaceControllerConstructorFunctionFile)
	testOutputScaffoldInterfaceControllerMethod = string(testOutputScaffoldInterfaceControllerMethodFile)

}

// Test Scaffold Interface Controller File
func TestScaffoldInterfaceControllerFile(t *testing.T){

	// Build
	statement, err := testControllerGenerator.ScaffoldFile(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldInterfaceControllerFile() failed with error %v`, err)
	}

	f, err := os.Create("./testing/interface/controller/created/" + testOutputScaffoldInterfaceControllerFileName)
	if err != nil {
		t.Errorf(`scaffoldInterfaceControllerFile() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldInterfaceControllerFile() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldInterfaceControllerFile {
		t.Errorf(`scaffoldInterfaceControllerFile() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceControllerFile, buf.String())
	}
	
}

// Test Scaffold Interface Controller Struct
func TestScaffoldInterfaceControllerStruct(t *testing.T){

	// Build
	statement, err := testControllerGenerator.scaffoldInterfaceControllerStruct(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldInterfaceControllerStruct() failed with error %v`, err)
	}

	f, err := os.Create("./testing/interface/controller/created/" + testOutputScaffoldInterfaceControllerStructName)
	if err != nil {
		t.Errorf(`scaffoldInterfaceControllerStruct() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldInterfaceControllerStruct() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldInterfaceControllerStruct {
		t.Errorf(`scaffoldInterfaceControllerStruct() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceControllerStruct, buf.String())
	}
	
}

// Test Scaffold Interface Controller Constructor Function
func TestScaffoldInterfaceControllerConstructorFunction(t *testing.T){

	// Build
	statement, err := testControllerGenerator.scaffoldInterfaceControllerConstructorFunction(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldInterfaceControllerConstructorFunction() failed with error %v`, err)
	}

	f, err := os.Create("./testing/interface/controller/created/" + testOutputScaffoldInterfaceControllerConstructorFunctionName)
	if err != nil {
		t.Errorf(`scaffoldInterfaceControllerConstructorFunction() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldInterfaceControllerConstructorFunction() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldInterfaceControllerConstructorFunction {
		t.Errorf(`scaffoldInterfaceControllerConstructorFunction() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceControllerConstructorFunction, buf.String())
	}
	
}

// Test Scaffold Interface Controller Method
func TestScaffoldInterfaceControllerMethod(t *testing.T){

	// Build
	statement, err := testControllerGenerator.scaffoldInterfaceControllerMethod(testMethod, testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldInterfaceControllerMethod() failed with error %v`, err)
	}

	f, err := os.Create("./testing/interface/controller/created/" + testOutputScaffoldInterfaceControllerMethodName)
	if err != nil {
		t.Errorf(`scaffoldInterfaceControllerMethod() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldInterfaceControllerMethod() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldInterfaceControllerMethod {
		t.Errorf(`scaffoldInterfaceControllerMethod() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceControllerMethod, buf.String())
	}
	
}