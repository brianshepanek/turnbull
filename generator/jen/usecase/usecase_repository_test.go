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
	testRepositoryGenerator RepositoryGenerator

	testOutputScaffoldUsecaseRepositoryFileName = "testOutputScaffoldUsecaseRepositoryFile"
	testOutputScaffoldUsecaseRepositoryInterfaceName = "testOutputScaffoldUsecaseRepositoryInterface"
	testOutputScaffoldUsecaseRepositoryInterfaceMethodName = "testOutputScaffoldUsecaseRepositoryInterfaceMethod"

	testOutputScaffoldUsecaseRepositoryFile string
	testOutputScaffoldUsecaseRepositoryInterface string
	testOutputScaffoldUsecaseRepositoryInterfaceMethod string

)

func init(){
	conf, _ := config.New(testConfigPath, testOutputPath)
	formatter := formatter.New(conf)
	testHelperGenerator := helper.New(formatter)
	testRepositoryGenerator = NewRepositoryGenerator(conf, formatter, testHelperGenerator)

	testOutputScaffoldUsecaseRepositoryFileFile, _ := ioutil.ReadFile("./testing/usecase/repository/expected/" + testOutputScaffoldUsecaseRepositoryFileName)
	testOutputScaffoldUsecaseRepositoryInterfaceFile, _ := ioutil.ReadFile("./testing/usecase/repository/expected/" + testOutputScaffoldUsecaseRepositoryInterfaceName)
	testOutputScaffoldUsecaseRepositoryInterfaceMethodFile, _ := ioutil.ReadFile("./testing/usecase/repository/expected/" + testOutputScaffoldUsecaseRepositoryInterfaceMethodName)

	testOutputScaffoldUsecaseRepositoryFile = string(testOutputScaffoldUsecaseRepositoryFileFile)
	testOutputScaffoldUsecaseRepositoryInterface = string(testOutputScaffoldUsecaseRepositoryInterfaceFile)
	testOutputScaffoldUsecaseRepositoryInterfaceMethod = string(testOutputScaffoldUsecaseRepositoryInterfaceMethodFile)

}

// Test Scaffold Usecase Repository File
func TestScaffoldUsecaseRepositoryFile(t *testing.T){

	// Build
	statement, err := testRepositoryGenerator.ScaffoldFile(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldUsecaseRepositoryFile() failed with error %v`, err)
	}

	f, err := os.Create("./testing/usecase/repository/created/" + testOutputScaffoldUsecaseRepositoryFileName)
	if err != nil {
		t.Errorf(`scaffoldUsecaseRepositoryFile() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldUsecaseRepositoryFile() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldUsecaseRepositoryFile {
		t.Errorf(`scaffoldUsecaseRepositoryFile() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseRepositoryFile, buf.String())
	}
	
}

// Test Scaffold Usecase Repository Interface
func TestScaffoldUsecaseRepositoryInterface(t *testing.T){

	// Build
	statement, err := testRepositoryGenerator.scaffoldUsecaseRepositoryInterface(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldUsecaseRepositoryInterface() failed with error %v`, err)
	}

	f, err := os.Create("./testing/usecase/repository/created/" + testOutputScaffoldUsecaseRepositoryInterfaceName)
	if err != nil {
		t.Errorf(`scaffoldUsecaseRepositoryInterface() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldUsecaseRepositoryInterface() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldUsecaseRepositoryInterface {
		t.Errorf(`scaffoldEntityStruct() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseRepositoryInterface, buf.String())
	}
	
}


// // Test Scaffold Usecase Repository Method
// func TestScaffoldUsecaseRepositoryInterfaceMethod(t *testing.T){

// 	// Build
// 	statement, err := testRepositoryGenerator.scaffoldUsecaseRepositoryInterfaceMethod(testMethod, testEntity)

// 	// Return
// 	if err != nil {
// 		t.Errorf(`scaffoldUsecaseRepositoryInterfaceMethod() failed with error %v`, err)
// 	}

// 	f, err := os.Create("./testing/usecase/repository/created/" + testOutputScaffoldUsecaseRepositoryInterfaceMethodName)
// 	if err != nil {
// 		t.Errorf(`scaffoldUsecaseRepositoryInterfaceMethod() failed with error %v`, err)
// 	}
// 	buf := &bytes.Buffer{}
// 	err = statement.Render(buf)
// 	if err != nil {
// 		t.Errorf(`scaffoldUsecaseRepositoryInterfaceMethod() failed with error %v`, err)
// 	}
// 	_, err = f.Write(buf.Bytes())

// 	if buf.String() != testOutputScaffoldUsecaseRepositoryInterfaceMethod {
// 		t.Errorf(`scaffoldEntityStruct() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseRepositoryInterfaceMethod, buf.String())
// 	}
	
// }
