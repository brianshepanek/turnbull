package scribble

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
	testRepositoryGenerator *repositoryGenerator

	testOutputInterfaceRepositoryFileName = "testOutputInterfaceRepositoryFile"
	testOutputScaffoldInterfaceRepositoryFileName = "testOutputScaffoldInterfaceRepositoryFile"
	testOutputScaffoldInterfaceRepositoryStructName = "testOutputScaffoldInterfaceRepositoryStruct"
	testOutputScaffoldInterfaceRepositoryConstructorFunctionName = "testOutputScaffoldInterfaceRepositoryConstructorFunction"
	testOutputScaffoldInterfaceRepositoryMethodName = "testOutputScaffoldInterfaceRepositoryMethod"
	
	testOutputInterfaceRepositoryFile string
	testOutputScaffoldInterfaceRepositoryFile string
	testOutputScaffoldInterfaceRepositoryStruct string
	testOutputScaffoldInterfaceRepositoryConstructorFunction string
	testOutputScaffoldInterfaceRepositoryMethod string

	testEntity  = model.Entity{
		Name : "foo",
		JSON : true,
		Fields : []model.Field{
			model.Field{
				Primary : true,
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
	testRepositoryGenerator = New(conf, formatter, testHelperGenerator)

	testOutputInterfaceRepositoryFileFile, _ := ioutil.ReadFile("./testing/interface/repository/expected/" + testOutputInterfaceRepositoryFileName)
	testOutputScaffoldInterfaceRepositoryFileFile, _ := ioutil.ReadFile("./testing/interface/repository/expected/" + testOutputScaffoldInterfaceRepositoryFileName)
	testOutputScaffoldInterfaceRepositoryStructFile, _ := ioutil.ReadFile("./testing/interface/repository/expected/" + testOutputScaffoldInterfaceRepositoryStructName)
	testOutputScaffoldInterfaceRepositoryConstructorFunctionFile, _ := ioutil.ReadFile("./testing/interface/repository/expected/" + testOutputScaffoldInterfaceRepositoryConstructorFunctionName)
	testOutputScaffoldInterfaceRepositoryMethodFile, _ := ioutil.ReadFile("./testing/interface/repository/expected/" + testOutputScaffoldInterfaceRepositoryMethodName)
	
	testOutputInterfaceRepositoryFile = string(testOutputInterfaceRepositoryFileFile)
	testOutputScaffoldInterfaceRepositoryFile = string(testOutputScaffoldInterfaceRepositoryFileFile)
	testOutputScaffoldInterfaceRepositoryStruct = string(testOutputScaffoldInterfaceRepositoryStructFile)
	testOutputScaffoldInterfaceRepositoryConstructorFunction = string(testOutputScaffoldInterfaceRepositoryConstructorFunctionFile)
	testOutputScaffoldInterfaceRepositoryMethod = string(testOutputScaffoldInterfaceRepositoryMethodFile)

}

// Test Interface Repository File
func TestInterfaceRepositoryFile(t *testing.T){

	// Build
	statement, err := testRepositoryGenerator.File(testEntity)

	// Return
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}

	f, err := os.Create("./testing/interface/repository/created/" + testOutputInterfaceRepositoryFileName)
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputInterfaceRepositoryFile {
		t.Errorf(`File() failed; want "%s", got "%s"`, testOutputInterfaceRepositoryFile, buf.String())
	}
	
}

// Test Scaffold Interface Repository File
func TestScaffoldInterfaceRepositoryFile(t *testing.T){

	// Build
	statement, err := testRepositoryGenerator.ScaffoldFile(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldInterfaceRepositoryFile() failed with error %v`, err)
	}

	f, err := os.Create("./testing/interface/repository/created/" + testOutputScaffoldInterfaceRepositoryFileName)
	if err != nil {
		t.Errorf(`scaffoldInterfaceRepositoryFile() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldInterfaceRepositoryFile() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldInterfaceRepositoryFile {
		t.Errorf(`scaffoldInterfaceRepositoryFile() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceRepositoryFile, buf.String())
	}
	
}

// Test Scaffold Interface Repository Struct
func TestScaffoldInterfaceRepositoryStruct(t *testing.T){

	// Build
	statement, err := testRepositoryGenerator.scaffoldInterfaceRepositoryStruct(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldInterfaceRepositoryStruct() failed with error %v`, err)
	}

	f, err := os.Create("./testing/interface/repository/created/" + testOutputScaffoldInterfaceRepositoryStructName)
	if err != nil {
		t.Errorf(`scaffoldInterfaceRepositoryStruct() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldInterfaceRepositoryStruct() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldInterfaceRepositoryStruct {
		t.Errorf(`scaffoldInterfaceRepositoryStruct() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceRepositoryStruct, buf.String())
	}
	
}

// Test Interface Repository Constructor Function
func TestInterfaceRepositoryConstructorFunction(t *testing.T){

	// Build
	statement, err := testRepositoryGenerator.interfaceRepositoryConstructorFunction(testEntity)

	// Return
	if err != nil {
		t.Errorf(`interfaceRepositoryConstructorFunction() failed with error %v`, err)
	}

	f, err := os.Create("./testing/interface/repository/created/" + testOutputScaffoldInterfaceRepositoryConstructorFunctionName)
	if err != nil {
		t.Errorf(`interfaceRepositoryConstructorFunction() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`interfaceRepositoryConstructorFunction() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldInterfaceRepositoryConstructorFunction {
		t.Errorf(`interfaceRepositoryConstructorFunction() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceRepositoryConstructorFunction, buf.String())
	}
	
}

// Test Scaffold Interface Repository Method
func TestScaffoldInterfaceRepositoryMethod(t *testing.T){

	// Build
	statement, err := testRepositoryGenerator.scaffoldInterfaceRepositoryMethod(testMethod, testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldInterfaceRepositoryMethod() failed with error %v`, err)
	}

	f, err := os.Create("./testing/interface/repository/created/" + testOutputScaffoldInterfaceRepositoryMethodName)
	if err != nil {
		t.Errorf(`scaffoldInterfaceRepositoryMethod() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldInterfaceRepositoryMethod() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldInterfaceRepositoryMethod {
		t.Errorf(`scaffoldInterfaceRepositoryMethod() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceRepositoryMethod, buf.String())
	}
	
}