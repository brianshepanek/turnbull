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
	testPresenterGenerator *presenterGenerator

	testOutputInterfacePresenterFileName = "testOutputInterfacePresenterFile"
	testOutputScaffoldInterfacePresenterFileName = "testOutputScaffoldInterfacePresenterFile"
	testOutputScaffoldInterfacePresenterStructName = "testOutputScaffoldInterfacePresenterStruct"
	testOutputScaffoldInterfacePresenterInterfaceName = "testOutputScaffoldInterfacePresenterInterface"
	testOutputScaffoldInterfacePresenterInterfaceMethodName = "testOutputScaffoldInterfacePresenterInterfaceMethod"
	testOutputScaffoldInterfacePresenterConstructorFunctionName = "testOutputScaffoldInterfacePresenterConstructorFunction"
	testOutputScaffoldInterfacePresenterInterfaceFunctionName = "testOutputScaffoldInterfacePresenterInterfaceFunction"

	testOutputInterfacePresenterFile string
	testOutputScaffoldInterfacePresenterFile string
	testOutputScaffoldInterfacePresenterStruct string
	testOutputScaffoldInterfacePresenterInterface string
	testOutputScaffoldInterfacePresenterInterfaceMethod string
	testOutputScaffoldInterfacePresenterConstructorFunction string
	testOutputScaffoldInterfacePresenterInterfaceFunction string

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
	testPresenterGenerator = New(conf, formatter, testHelperGenerator)

	testOutputInterfacePresenterFileFile, _ := ioutil.ReadFile("./testing/interface/presenter/expected/" + testOutputInterfacePresenterFileName)
	testOutputScaffoldInterfacePresenterFileFile, _ := ioutil.ReadFile("./testing/interface/presenter/expected/" + testOutputScaffoldInterfacePresenterFileName)
	testOutputScaffoldInterfacePresenterStructFile, _ := ioutil.ReadFile("./testing/interface/presenter/expected/" + testOutputScaffoldInterfacePresenterStructName)
	testOutputScaffoldInterfacePresenterInterfaceFile, _ := ioutil.ReadFile("./testing/interface/presenter/expected/" + testOutputScaffoldInterfacePresenterInterfaceName)
	testOutputScaffoldInterfacePresenterInterfaceMethodFile, _ := ioutil.ReadFile("./testing/interface/presenter/expected/" + testOutputScaffoldInterfacePresenterInterfaceMethodName)
	testOutputScaffoldInterfacePresenterConstructorFunctionFile, _ := ioutil.ReadFile("./testing/interface/presenter/expected/" + testOutputScaffoldInterfacePresenterConstructorFunctionName)
	testOutputScaffoldInterfacePresenterInterfaceFunctionFile, _ := ioutil.ReadFile("./testing/interface/presenter/expected/" + testOutputScaffoldInterfacePresenterInterfaceFunctionName)

	testOutputInterfacePresenterFile = string(testOutputInterfacePresenterFileFile)
	testOutputScaffoldInterfacePresenterFile = string(testOutputScaffoldInterfacePresenterFileFile)
	testOutputScaffoldInterfacePresenterStruct = string(testOutputScaffoldInterfacePresenterStructFile)
	testOutputScaffoldInterfacePresenterInterface = string(testOutputScaffoldInterfacePresenterInterfaceFile)
	testOutputScaffoldInterfacePresenterInterfaceMethod = string(testOutputScaffoldInterfacePresenterInterfaceMethodFile)
	testOutputScaffoldInterfacePresenterConstructorFunction = string(testOutputScaffoldInterfacePresenterConstructorFunctionFile)
	testOutputScaffoldInterfacePresenterInterfaceFunction = string(testOutputScaffoldInterfacePresenterInterfaceFunctionFile)

}

// Test Interface Presenter File
func TestInterfacePresenterFile(t *testing.T){

	// Build
	statement, err := testPresenterGenerator.File(testEntity)

	// Return
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}

	f, err := os.Create("./testing/interface/presenter/created/" + testOutputInterfacePresenterFileName)
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputInterfacePresenterFile {
		t.Errorf(`File() failed; want "%s", got "%s"`, testOutputInterfacePresenterFile, buf.String())
	}
	
}

// Test Scaffold Interface Presenter File
func TestScaffoldInterfacePresenterFile(t *testing.T){

	// Build
	statement, err := testPresenterGenerator.ScaffoldFile(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldInterfacePresenterFile() failed with error %v`, err)
	}

	f, err := os.Create("./testing/interface/presenter/created/" + testOutputScaffoldInterfacePresenterFileName)
	if err != nil {
		t.Errorf(`scaffoldInterfacePresenterFile() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldInterfacePresenterFile() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldInterfacePresenterFile {
		t.Errorf(`scaffoldInterfacePresenterFile() failed; want "%s", got "%s"`, testOutputScaffoldInterfacePresenterFile, buf.String())
	}
	
}

// Test Scaffold Interface Presenter Struct
func TestScaffoldInterfacePresenterStruct(t *testing.T){

	// Build
	statement, err := testPresenterGenerator.scaffoldInterfacePresenterStruct(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldInterfacePresenterStruct() failed with error %v`, err)
	}

	f, err := os.Create("./testing/interface/presenter/created/" + testOutputScaffoldInterfacePresenterStructName)
	if err != nil {
		t.Errorf(`scaffoldInterfacePresenterStruct() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldInterfacePresenterStruct() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldInterfacePresenterStruct {
		t.Errorf(`scaffoldInterfacePresenterStruct() failed; want "%s", got "%s"`, testOutputScaffoldInterfacePresenterStruct, buf.String())
	}
	
}

// Test Scaffold Interface Presenter Interface
func TestScaffoldInterfacePresenterInterface(t *testing.T){

	// Build
	statement, err := testPresenterGenerator.scaffoldInterfacePresenterInterface(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldInterfacePresenterInterface() failed with error %v`, err)
	}

	f, err := os.Create("./testing/interface/presenter/created/" + testOutputScaffoldInterfacePresenterInterfaceName)
	if err != nil {
		t.Errorf(`scaffoldInterfacePresenterInterface() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldInterfacePresenterInterface() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldInterfacePresenterInterface {
		t.Errorf(`scaffoldInterfacePresenterInterface() failed; want "%s", got "%s"`, testOutputScaffoldInterfacePresenterInterface, buf.String())
	}
	
}

// Test Scaffold Interface Presenter Constructor Function
func TestScaffoldInterfacePresenterConstructorFunction(t *testing.T){

	// Build
	statement, err := testPresenterGenerator.interfacePresenterConstructorFunction(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldInterfacePresenterConstructorFunction() failed with error %v`, err)
	}

	f, err := os.Create("./testing/interface/presenter/created/" + testOutputScaffoldInterfacePresenterConstructorFunctionName)
	if err != nil {
		t.Errorf(`scaffoldInterfacePresenterConstructorFunction() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldInterfacePresenterConstructorFunction() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldInterfacePresenterConstructorFunction {
		t.Errorf(`scaffoldInterfacePresenterConstructorFunction() failed; want "%s", got "%s"`, testOutputScaffoldInterfacePresenterConstructorFunction, buf.String())
	}
	
}

// // Test Scaffold Interface Presenter Interface Function
// func TestScaffoldInterfacePresenterInterfaceFunction(t *testing.T){

// 	// Build
// 	statement, err := testPresenterGenerator.scaffoldInterfacePresenterInterfaceMethod(testMethod, testEntity)

// 	// Return
// 	if err != nil {
// 		t.Errorf(`scaffoldInterfacePresenterInterfaceFunction() failed with error %v`, err)
// 	}

// 	f, err := os.Create("./testing/interface/presenter/created/" + testOutputScaffoldInterfacePresenterInterfaceFunctionName)
// 	if err != nil {
// 		t.Errorf(`scaffoldInterfacePresenterInterfaceFunction() failed with error %v`, err)
// 	}
// 	buf := &bytes.Buffer{}
// 	err = statement.Render(buf)
// 	if err != nil {
// 		t.Errorf(`scaffoldInterfacePresenterInterfaceFunction() failed with error %v`, err)
// 	}
// 	_, err = f.Write(buf.Bytes())

// 	if buf.String() != testOutputScaffoldInterfacePresenterInterfaceFunction {
// 		t.Errorf(`scaffoldInterfacePresenterInterfaceFunction() failed; want "%s", got "%s"`, testOutputScaffoldInterfacePresenterInterfaceFunction, buf.String())
// 	}
	
// }

// // Test Scaffold Interface Presenter Method
// func TestScaffoldInterfacePresenterInterfaceMethod(t *testing.T){

// 	// Build
// 	statement, err := testPresenterGenerator.scaffoldInterfacePresenterInterfaceMethod(testMethod, testEntity)

// 	// Return
// 	if err != nil {
// 		t.Errorf(`scaffoldInterfacePresenterInterfaceMethod() failed with error %v`, err)
// 	}

// 	f, err := os.Create("./testing/interface/presenter/created/" + testOutputScaffoldInterfacePresenterInterfaceMethodName)
// 	if err != nil {
// 		t.Errorf(`scaffoldInterfacePresenterInterfaceMethod() failed with error %v`, err)
// 	}
// 	buf := &bytes.Buffer{}
// 	err = statement.Render(buf)
// 	if err != nil {
// 		t.Errorf(`scaffoldInterfacePresenterInterfaceMethod() failed with error %v`, err)
// 	}
// 	_, err = f.Write(buf.Bytes())

// 	if buf.String() != testOutputScaffoldInterfacePresenterInterfaceMethod {
// 		t.Errorf(`scaffoldEntityStruct() failed; want "%s", got "%s"`, testOutputScaffoldInterfacePresenterInterfaceMethod, buf.String())
// 	}
	
// }
