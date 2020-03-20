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
	testPresenterGenerator PresenterGenerator

	testOutputUsecasePresenterFileName = "testOutputUsecasePresenterFile"
	testOutputScaffoldUsecasePresenterFileName = "testOutputScaffoldUsecasePresenterFile"
	testOutputScaffoldUsecasePresenterInterfaceName = "testOutputScaffoldUsecasePresenterInterface"
	testOutputScaffoldUsecasePresenterInterfaceMethodName = "testOutputScaffoldUsecasePresenterInterfaceMethod"

	testOutputUsecasePresenterFile string
	testOutputScaffoldUsecasePresenterFile string
	testOutputScaffoldUsecasePresenterInterface string
	testOutputScaffoldUsecasePresenterInterfaceMethod string

)

func init(){
	conf, _ := config.New(testConfigPath, testOutputPath)
	formatter := formatter.New(conf)
	testHelperGenerator := helper.New(formatter)
	testPresenterGenerator = NewPresenterGenerator(conf, formatter, testHelperGenerator)

	testOutputUsecasePresenterFileFile, _ := ioutil.ReadFile("./testing/usecase/presenter/expected/" + testOutputUsecasePresenterFileName)
	testOutputScaffoldUsecasePresenterFileFile, _ := ioutil.ReadFile("./testing/usecase/presenter/expected/" + testOutputScaffoldUsecasePresenterFileName)
	testOutputScaffoldUsecasePresenterInterfaceFile, _ := ioutil.ReadFile("./testing/usecase/presenter/expected/" + testOutputScaffoldUsecasePresenterInterfaceName)
	testOutputScaffoldUsecasePresenterInterfaceMethodFile, _ := ioutil.ReadFile("./testing/usecase/presenter/expected/" + testOutputScaffoldUsecasePresenterInterfaceMethodName)

	testOutputUsecasePresenterFile = string(testOutputUsecasePresenterFileFile)
	testOutputScaffoldUsecasePresenterFile = string(testOutputScaffoldUsecasePresenterFileFile)
	testOutputScaffoldUsecasePresenterInterface = string(testOutputScaffoldUsecasePresenterInterfaceFile)
	testOutputScaffoldUsecasePresenterInterfaceMethod = string(testOutputScaffoldUsecasePresenterInterfaceMethodFile)

}

// Test Usecase Presenter File
func TestUsecasePresenterFile(t *testing.T){

	// Build
	statement, err := testPresenterGenerator.File(testEntity)

	// Return
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}

	f, err := os.Create("./testing/usecase/presenter/created/" + testOutputUsecasePresenterFileName)
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputUsecasePresenterFile {
		t.Errorf(`File() failed; want "%s", got "%s"`, testOutputUsecasePresenterFile, buf.String())
	}
	
}

// Test Scaffold Usecase Presenter File
func TestScaffoldUsecasePresenterFile(t *testing.T){

	// Build
	statement, err := testPresenterGenerator.ScaffoldFile(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldUsecasePresenterFile() failed with error %v`, err)
	}

	f, err := os.Create("./testing/usecase/presenter/created/" + testOutputScaffoldUsecasePresenterFileName)
	if err != nil {
		t.Errorf(`scaffoldUsecasePresenterFile() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldUsecasePresenterFile() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldUsecasePresenterFile {
		t.Errorf(`scaffoldUsecasePresenterFile() failed; want "%s", got "%s"`, testOutputScaffoldUsecasePresenterFile, buf.String())
	}
	
}

// Test Scaffold Usecase Presenter Interface
func TestScaffoldUsecasePresenterInterface(t *testing.T){

	// Build
	statement, err := testPresenterGenerator.scaffoldUsecasePresenterInterface(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldUsecasePresenterInterface() failed with error %v`, err)
	}

	f, err := os.Create("./testing/usecase/presenter/created/" + testOutputScaffoldUsecasePresenterInterfaceName)
	if err != nil {
		t.Errorf(`scaffoldUsecasePresenterInterface() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldUsecasePresenterInterface() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldUsecasePresenterInterface {
		t.Errorf(`scaffoldEntityStruct() failed; want "%s", got "%s"`, testOutputScaffoldUsecasePresenterInterface, buf.String())
	}
	
}


// // Test Scaffold Usecase Presenter Method
// func TestScaffoldUsecasePresenterInterfaceMethod(t *testing.T){

// 	// Build
// 	statement, err := testPresenterGenerator.scaffoldUsecasePresenterInterfaceMethod(testMethod, testEntity)

// 	// Return
// 	if err != nil {
// 		t.Errorf(`scaffoldUsecasePresenterInterfaceMethod() failed with error %v`, err)
// 	}

// 	f, err := os.Create("./testing/usecase/presenter/created/" + testOutputScaffoldUsecasePresenterInterfaceMethodName)
// 	if err != nil {
// 		t.Errorf(`scaffoldUsecasePresenterInterfaceMethod() failed with error %v`, err)
// 	}
// 	buf := &bytes.Buffer{}
// 	err = statement.Render(buf)
// 	if err != nil {
// 		t.Errorf(`scaffoldUsecasePresenterInterfaceMethod() failed with error %v`, err)
// 	}
// 	_, err = f.Write(buf.Bytes())

// 	if buf.String() != testOutputScaffoldUsecasePresenterInterfaceMethod {
// 		t.Errorf(`scaffoldEntityStruct() failed; want "%s", got "%s"`, testOutputScaffoldUsecasePresenterInterfaceMethod, buf.String())
// 	}
	
// }
