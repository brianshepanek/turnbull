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

	testOutputUsecaseInteractorFileName = "testOutputUsecaseInteractorFile"
	testOutputScaffoldUsecaseInteractorFileName = "testOutputScaffoldUsecaseInteractorFile"
	testOutputScaffoldUsecaseInteractorStructName = "testOutputScaffoldUsecaseInteractorStruct"
	testOutputScaffoldUsecaseInteractorInterfaceName = "testOutputScaffoldUsecaseInteractorInterface"
	testOutputScaffoldUsecaseInteractorInterfaceMethodName = "testOutputScaffoldUsecaseInteractorInterfaceMethod"
	testOutputUsecaseInteractorConstructorFunctionName = "testOutputUsecaseInteractorConstructorFunction"
	testOutputScaffoldUsecaseInteractorInterfaceFunctionName = "testOutputScaffoldUsecaseInteractorInterfaceFunction"

	testOutputUsecaseInteractorFile string
	testOutputScaffoldUsecaseInteractorFile string
	testOutputScaffoldUsecaseInteractorStruct string
	testOutputScaffoldUsecaseInteractorInterface string
	testOutputScaffoldUsecaseInteractorInterfaceMethod string
	testOutputUsecaseInteractorConstructorFunction string
	testOutputScaffoldUsecaseInteractorInterfaceFunction string

)

func init(){
	conf, _ := config.New(testConfigPath, testOutputPath)
	formatter := formatter.New(conf)
	testHelperGenerator := helper.New(formatter)
	testInteractorGenerator = NewInteractorGenerator(conf, formatter, testHelperGenerator)

	testOutputUsecaseInteractorFileFile, _ := ioutil.ReadFile("./testing/usecase/interactor/expected/" + testOutputUsecaseInteractorFileName)
	testOutputScaffoldUsecaseInteractorFileFile, _ := ioutil.ReadFile("./testing/usecase/interactor/expected/" + testOutputScaffoldUsecaseInteractorFileName)
	testOutputScaffoldUsecaseInteractorStructFile, _ := ioutil.ReadFile("./testing/usecase/interactor/expected/" + testOutputScaffoldUsecaseInteractorStructName)
	testOutputScaffoldUsecaseInteractorInterfaceFile, _ := ioutil.ReadFile("./testing/usecase/interactor/expected/" + testOutputScaffoldUsecaseInteractorInterfaceName)
	testOutputScaffoldUsecaseInteractorInterfaceMethodFile, _ := ioutil.ReadFile("./testing/usecase/interactor/expected/" + testOutputScaffoldUsecaseInteractorInterfaceMethodName)
	testOutputUsecaseInteractorConstructorFunctionFile, _ := ioutil.ReadFile("./testing/usecase/interactor/expected/" + testOutputUsecaseInteractorConstructorFunctionName)
	testOutputScaffoldUsecaseInteractorInterfaceFunctionFile, _ := ioutil.ReadFile("./testing/usecase/interactor/expected/" + testOutputScaffoldUsecaseInteractorInterfaceFunctionName)

	testOutputUsecaseInteractorFile = string(testOutputUsecaseInteractorFileFile)
	testOutputScaffoldUsecaseInteractorFile = string(testOutputScaffoldUsecaseInteractorFileFile)
	testOutputScaffoldUsecaseInteractorStruct = string(testOutputScaffoldUsecaseInteractorStructFile)
	testOutputScaffoldUsecaseInteractorInterface = string(testOutputScaffoldUsecaseInteractorInterfaceFile)
	testOutputScaffoldUsecaseInteractorInterfaceMethod = string(testOutputScaffoldUsecaseInteractorInterfaceMethodFile)
	testOutputUsecaseInteractorConstructorFunction = string(testOutputUsecaseInteractorConstructorFunctionFile)
	testOutputScaffoldUsecaseInteractorInterfaceFunction = string(testOutputScaffoldUsecaseInteractorInterfaceFunctionFile)

}

// Test Usecase Interactor File
func TestUsecaseInteractorFile(t *testing.T){

	// Build
	statement, err := testInteractorGenerator.File(testEntity)

	// Return
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}

	f, err := os.Create("./testing/usecase/interactor/created/" + testOutputUsecaseInteractorFileName)
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputUsecaseInteractorFile {
		t.Errorf(`File() failed; want "%s", got "%s"`, testOutputUsecaseInteractorFile, buf.String())
	}
	
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

// Test  Usecase Interactor Constructor Function
func TestUsecaseInteractorConstructorFunction(t *testing.T){

	// Build
	statement, err := testInteractorGenerator.usecaseInteractorConstructorFunction(testEntity)

	// Return
	if err != nil {
		t.Errorf(`usecaseInteractorConstructorFunction() failed with error %v`, err)
	}

	f, err := os.Create("./testing/usecase/interactor/created/" + testOutputUsecaseInteractorConstructorFunctionName)
	if err != nil {
		t.Errorf(`usecaseInteractorConstructorFunction() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`usecaseInteractorConstructorFunction() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputUsecaseInteractorConstructorFunction {
		t.Errorf(`usecaseInteractorConstructorFunction() failed; want "%s", got "%s"`, testOutputUsecaseInteractorConstructorFunction, buf.String())
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
