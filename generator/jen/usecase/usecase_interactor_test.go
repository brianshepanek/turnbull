package usecase

import(
	"os"
	"bytes"
	"io/ioutil"
	"testing"
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/generator/jen/helper"
)


var (
	testInteractorGenerator InteractorGenerator

	testOutputScaffoldUsecaseInteractorFileName = "testOutputScaffoldUsecaseInteractorFile"
	testOutputScaffoldUsecaseInteractorStructName = "testOutputScaffoldUsecaseInteractorStruct"
	testOutputScaffoldUsecaseInteractorInterfaceName = "testOutputScaffoldUsecaseInteractorInterface"
	testOutputScaffoldUsecaseInteractorInterfaceMethodName = "testOutputScaffoldUsecaseInteractorInterfaceMethod"
	testOutputScaffoldUsecaseInteractorConstructorFunctionName = "testOutputScaffoldUsecaseInteractorConstructorFunction"
	testOutputScaffoldUsecaseInteractorInterfaceFunctionName = "testOutputScaffoldUsecaseInteractorInterfaceFunction"

	testOutputScaffoldUsecaseInteractorFile string
	testOutputScaffoldUsecaseInteractorStruct string
	testOutputScaffoldUsecaseInteractorInterface string
	testOutputScaffoldUsecaseInteractorInterfaceMethod string
	testOutputScaffoldUsecaseInteractorConstructorFunction string
	testOutputScaffoldUsecaseInteractorInterfaceFunction string

)

func init(){
	conf, _ := config.New(testConfigPath, testOutputPath)
	formatter := formatter.New(conf)
	testHelperGenerator := helper.New(formatter)
	testInteractorGenerator = NewInteractorGenerator(conf, formatter, testHelperGenerator)

	testOutputScaffoldUsecaseInteractorFileFile, _ := ioutil.ReadFile("./testing/usecase/interactor/expected/" + testOutputScaffoldUsecaseInteractorFileName)
	testOutputScaffoldUsecaseInteractorStructFile, _ := ioutil.ReadFile("./testing/usecase/interactor/expected/" + testOutputScaffoldUsecaseInteractorStructName)
	testOutputScaffoldUsecaseInteractorInterfaceFile, _ := ioutil.ReadFile("./testing/usecase/interactor/expected/" + testOutputScaffoldUsecaseInteractorInterfaceName)
	testOutputScaffoldUsecaseInteractorInterfaceMethodFile, _ := ioutil.ReadFile("./testing/usecase/interactor/expected/" + testOutputScaffoldUsecaseInteractorInterfaceMethodName)
	testOutputScaffoldUsecaseInteractorConstructorFunctionFile, _ := ioutil.ReadFile("./testing/usecase/interactor/expected/" + testOutputScaffoldUsecaseInteractorConstructorFunctionName)
	testOutputScaffoldUsecaseInteractorInterfaceFunctionFile, _ := ioutil.ReadFile("./testing/usecase/interactor/expected/" + testOutputScaffoldUsecaseInteractorInterfaceFunctionName)

	testOutputScaffoldUsecaseInteractorFile = string(testOutputScaffoldUsecaseInteractorFileFile)
	testOutputScaffoldUsecaseInteractorStruct = string(testOutputScaffoldUsecaseInteractorStructFile)
	testOutputScaffoldUsecaseInteractorInterface = string(testOutputScaffoldUsecaseInteractorInterfaceFile)
	testOutputScaffoldUsecaseInteractorInterfaceMethod = string(testOutputScaffoldUsecaseInteractorInterfaceMethodFile)
	testOutputScaffoldUsecaseInteractorConstructorFunction = string(testOutputScaffoldUsecaseInteractorConstructorFunctionFile)
	testOutputScaffoldUsecaseInteractorInterfaceFunction = string(testOutputScaffoldUsecaseInteractorInterfaceFunctionFile)

}

// Test Scaffold Usecase Interactor File
func TestScaffoldUsecaseInteractorFile(t *testing.T){

	// Build
	statement, err := testInteractorGenerator.ScaffoldFile(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldUsecaseInteractorFile() failed with error %v`, err)
	}

	f, err := os.Create("./testing/usecase/interactor/created/" + testOutputScaffoldUsecaseInteractorFileName)
	if err != nil {
		t.Errorf(`scaffoldUsecaseInteractorFile() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldUsecaseInteractorFile() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldUsecaseInteractorFile {
		t.Errorf(`scaffoldUsecaseInteractorFile() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseInteractorFile, buf.String())
	}
	
}

// Test Scaffold Usecase Interactor Struct
func TestScaffoldUsecaseInteractorStruct(t *testing.T){

	// Build
	statement, err := testInteractorGenerator.scaffoldUsecaseInteractorStruct(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldUsecaseInteractorStruct() failed with error %v`, err)
	}

	f, err := os.Create("./testing/usecase/interactor/created/" + testOutputScaffoldUsecaseInteractorStructName)
	if err != nil {
		t.Errorf(`scaffoldUsecaseInteractorStruct() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldUsecaseInteractorStruct() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldUsecaseInteractorStruct {
		t.Errorf(`scaffoldUsecaseInteractorStruct() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseInteractorStruct, buf.String())
	}
	
}

// Test Scaffold Usecase Interactor Interface
func TestScaffoldUsecaseInteractorInterface(t *testing.T){

	// Build
	statement, err := testInteractorGenerator.scaffoldUsecaseInteractorInterface(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldUsecaseInteractorInterface() failed with error %v`, err)
	}

	f, err := os.Create("./testing/usecase/interactor/created/" + testOutputScaffoldUsecaseInteractorInterfaceName)
	if err != nil {
		t.Errorf(`scaffoldUsecaseInteractorInterface() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldUsecaseInteractorInterface() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldUsecaseInteractorInterface {
		t.Errorf(`scaffoldUsecaseInteractorInterface() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseInteractorInterface, buf.String())
	}
	
}

// Test Scaffold Usecase Interactor Constructor Function
func TestScaffoldUsecaseInteractorConstructorFunction(t *testing.T){

	// Build
	statement, err := testInteractorGenerator.scaffoldUsecaseInteractorConstructorFunction(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldUsecaseInteractorConstructorFunction() failed with error %v`, err)
	}

	f, err := os.Create("./testing/usecase/interactor/created/" + testOutputScaffoldUsecaseInteractorConstructorFunctionName)
	if err != nil {
		t.Errorf(`scaffoldUsecaseInteractorConstructorFunction() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldUsecaseInteractorConstructorFunction() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldUsecaseInteractorConstructorFunction {
		t.Errorf(`scaffoldUsecaseInteractorConstructorFunction() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseInteractorConstructorFunction, buf.String())
	}
	
}

// Test Scaffold Usecase Interactor Interface Function
func TestScaffoldUsecaseInteractorInterfaceFunction(t *testing.T){

	// Build
	statement, err := testInteractorGenerator.scaffoldUsecaseInteractorMethod(testMethod, testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldUsecaseInteractorMethod() failed with error %v`, err)
	}

	f, err := os.Create("./testing/usecase/interactor/created/" + testOutputScaffoldUsecaseInteractorInterfaceFunctionName)
	if err != nil {
		t.Errorf(`scaffoldUsecaseInteractorMethod() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldUsecaseInteractorMethod() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldUsecaseInteractorInterfaceFunction {
		t.Errorf(`scaffoldUsecaseInteractorMethod() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseInteractorInterfaceFunction, buf.String())
	}
	
}

// // Test Scaffold Usecase Interactor Method
// func TestScaffoldUsecaseInteractorInterfaceMethod(t *testing.T){

// 	// Build
// 	statement, err := testInteractorGenerator.scaffoldUsecaseInteractorInterfaceMethod(testMethod, testEntity)

// 	// Return
// 	if err != nil {
// 		t.Errorf(`scaffoldUsecaseInteractorInterfaceMethod() failed with error %v`, err)
// 	}

// 	f, err := os.Create("./testing/usecase/interactor/created/" + testOutputScaffoldUsecaseInteractorInterfaceMethodName)
// 	if err != nil {
// 		t.Errorf(`scaffoldUsecaseInteractorInterfaceMethod() failed with error %v`, err)
// 	}
// 	buf := &bytes.Buffer{}
// 	err = statement.Render(buf)
// 	if err != nil {
// 		t.Errorf(`scaffoldUsecaseInteractorInterfaceMethod() failed with error %v`, err)
// 	}
// 	_, err = f.Write(buf.Bytes())

// 	if buf.String() != testOutputScaffoldUsecaseInteractorInterfaceMethod {
// 		t.Errorf(`scaffoldEntityStruct() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseInteractorInterfaceMethod, buf.String())
// 	}
	
// }
