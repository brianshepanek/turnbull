package domain

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
		Interface : true,
		Fields : []model.Field{
			model.Field{
				Primary : true,
				Private : true,
				Name : "id",
				Type : "int64",
			},
			model.Field{
				Private : true,
				Name : "title",
				Type : "string",
			},
		},
		Methods : []model.Method {
			model.Method{
				Name : "add",
				Type : "add",
				Callbacks : []model.Callback {
					model.Callback {
						Type : "before",
					},
				},
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
		Private : true,
	}
	testMethod  = model.Method{
		Name : "add",
		Type : "add",
	}
)

var (
	testEntityGenerator EntityGenerator

	// Test File Names
	testOutputDomainEntityFileName = "testOutputDomainEntityFile"
	testOutputScaffoldDomainEntityFileName = "testOutputScaffoldDomainEntityFile"

	testOutputScaffoldDomainEntityStructName = "testOutputScaffoldDomainEntityStruct"
	testOutputScaffoldDomainEntitySliceStructName = "testOutputScaffoldDomainEntitySliceStruct"
	testOutputScaffoldDomainEntityInterfaceName = "testOutputScaffoldDomainEntityInterface"
	testOutputScaffoldDomainEntitySliceInterfaceName = "testOutputScaffoldDomainEntitySliceInterface"

	testOutputScaffoldDomainEntityInterfaceConstructorFunctionName = "testOutputScaffoldDomainEntityInterfaceConstructorFunction"
	testOutputScaffoldDomainEntitySliceInterfaceConstructorFunctionName = "testOutputScaffoldDomainEntitySliceInterfaceConstructorFunction"

	testOutputScaffoldDomainEntitySliceInterfaceLenFunctionName = "testOutputScaffoldDomainEntitySliceInterfaceLenFunction"
	testOutputScaffoldDomainEntitySliceInterfaceAppendFunctionName = "testOutputScaffoldDomainEntitySliceInterfaceAppendFunction"
	testOutputScaffoldDomainEntitySliceInterfaceElementsFunctionName = "testOutputScaffoldDomainEntitySliceInterfaceElementsFunction"

	testOutputScaffoldDomainEntityInterfaceGetterFunctionName = "testOutputScaffoldDomainEntityInterfaceGetterFunction"
	testOutputScaffoldDomainEntityInterfaceSetterFunctionName = "testOutputScaffoldDomainEntityInterfaceSetterFunction"
	testOutputScaffoldDomainEntityInterfaceSetAllSetterFunctionName = "testOutputScaffoldDomainEntityInterfaceSetAllSetterFunction"

	testOutputScaffoldDomainEntityInterfaceMarshalJSONFunctionName = "testOutputScaffoldDomainEntityInterfaceMarshalJSONFunction"
	testOutputScaffoldDomainEntityInterfaceUnmarshalJSONFunctionName = "testOutputScaffoldDomainEntityInterfaceUnmarshalJSONFunction"

	// Test File Strings
	testOutputDomainEntityFile string
	testOutputScaffoldDomainEntityFile string

	testOutputScaffoldDomainEntityStruct string
	testOutputScaffoldDomainEntitySliceStruct string
	testOutputScaffoldDomainEntityInterface string
	testOutputScaffoldDomainEntitySliceInterface string

	testOutputScaffoldDomainEntityInterfaceConstructorFunction string
	testOutputScaffoldDomainEntitySliceInterfaceConstructorFunction string

	testOutputScaffoldDomainEntitySliceInterfaceLenFunction string
	testOutputScaffoldDomainEntitySliceInterfaceAppendFunction string
	testOutputScaffoldDomainEntitySliceInterfaceElementsFunction string

	testOutputScaffoldDomainEntityInterfaceGetterFunction string
	testOutputScaffoldDomainEntityInterfaceSetterFunction string
	testOutputScaffoldDomainEntityInterfaceSetAllSetterFunction string

	testOutputScaffoldDomainEntityInterfaceMarshalJSONFunction string
	testOutputScaffoldDomainEntityInterfaceUnmarshalJSONFunction string

)


func init(){
	conf, _ := config.New(testConfigPath, testOutputPath)
	formatter := formatter.New(conf)
	testHelperGenerator := helper.New(formatter)
	testEntityGenerator = NewEntityGenerator(formatter, testHelperGenerator)

	// Test
	testOutputDomainEntityFileFile, _ := ioutil.ReadFile("./testing/domain/model/expected/" + testOutputDomainEntityFileName)
	testOutputScaffoldDomainEntityFileFile, _ := ioutil.ReadFile("./testing/domain/model/expected/" + testOutputScaffoldDomainEntityFileName)

	testOutputScaffoldDomainEntityStructFile, _ := ioutil.ReadFile("./testing/domain/model/expected/" + testOutputScaffoldDomainEntityStructName)
	testOutputScaffoldDomainEntitySliceStructFile, _ := ioutil.ReadFile("./testing/domain/model/expected/" + testOutputScaffoldDomainEntitySliceStructName)
	testOutputScaffoldDomainEntityInterfaceFile, _ := ioutil.ReadFile("./testing/domain/model/expected/" + testOutputScaffoldDomainEntityInterfaceName)
	testOutputScaffoldDomainEntitySliceInterfaceFile, _ := ioutil.ReadFile("./testing/domain/model/expected/" + testOutputScaffoldDomainEntitySliceInterfaceName)
	
	testOutputScaffoldDomainEntityInterfaceConstructorFunctionFile, _ := ioutil.ReadFile("./testing/domain/model/expected/" + testOutputScaffoldDomainEntityInterfaceConstructorFunctionName)
	testOutputScaffoldDomainEntitySliceInterfaceConstructorFunctionFile, _ := ioutil.ReadFile("./testing/domain/model/expected/" + testOutputScaffoldDomainEntitySliceInterfaceConstructorFunctionName)
	
	testOutputScaffoldDomainEntitySliceInterfaceLenFunctionFile, _ := ioutil.ReadFile("./testing/domain/model/expected/" + testOutputScaffoldDomainEntitySliceInterfaceLenFunctionName)
	testOutputScaffoldDomainEntitySliceInterfaceAppendFunctionFile, _ := ioutil.ReadFile("./testing/domain/model/expected/" + testOutputScaffoldDomainEntitySliceInterfaceAppendFunctionName)
	testOutputScaffoldDomainEntitySliceInterfaceElementsFunctionFile, _ := ioutil.ReadFile("./testing/domain/model/expected/" + testOutputScaffoldDomainEntitySliceInterfaceElementsFunctionName)

	testOutputScaffoldDomainEntityInterfaceGetterFunctionFile, _ := ioutil.ReadFile("./testing/domain/model/expected/" + testOutputScaffoldDomainEntityInterfaceGetterFunctionName)
	testOutputScaffoldDomainEntityInterfaceSetterFunctionFile, _ := ioutil.ReadFile("./testing/domain/model/expected/" + testOutputScaffoldDomainEntityInterfaceSetterFunctionName)
	testOutputScaffoldDomainEntityInterfaceSetAllSetterFunctionFile, _ := ioutil.ReadFile("./testing/domain/model/expected/" + testOutputScaffoldDomainEntityInterfaceSetAllSetterFunctionName)

	testOutputScaffoldDomainEntityInterfaceMarshalJSONFunctionFile, _ := ioutil.ReadFile("./testing/domain/model/expected/" + testOutputScaffoldDomainEntityInterfaceMarshalJSONFunctionName)
	testOutputScaffoldDomainEntityInterfaceUnmarshalJSONFunctionFile, _ := ioutil.ReadFile("./testing/domain/model/expected/" + testOutputScaffoldDomainEntityInterfaceUnmarshalJSONFunctionName)

	testOutputDomainEntityFile = string(testOutputDomainEntityFileFile)
	testOutputScaffoldDomainEntityFile = string(testOutputScaffoldDomainEntityFileFile)
	
	testOutputScaffoldDomainEntityStruct = string(testOutputScaffoldDomainEntityStructFile)
	testOutputScaffoldDomainEntitySliceStruct = string(testOutputScaffoldDomainEntitySliceStructFile)
	testOutputScaffoldDomainEntityInterface = string(testOutputScaffoldDomainEntityInterfaceFile)
	testOutputScaffoldDomainEntitySliceInterface = string(testOutputScaffoldDomainEntitySliceInterfaceFile)
	
	testOutputScaffoldDomainEntityInterfaceConstructorFunction = string(testOutputScaffoldDomainEntityInterfaceConstructorFunctionFile)
	testOutputScaffoldDomainEntitySliceInterfaceConstructorFunction = string(testOutputScaffoldDomainEntitySliceInterfaceConstructorFunctionFile)

	testOutputScaffoldDomainEntitySliceInterfaceLenFunction = string(testOutputScaffoldDomainEntitySliceInterfaceLenFunctionFile)
	testOutputScaffoldDomainEntitySliceInterfaceAppendFunction = string(testOutputScaffoldDomainEntitySliceInterfaceAppendFunctionFile)
	testOutputScaffoldDomainEntitySliceInterfaceElementsFunction = string(testOutputScaffoldDomainEntitySliceInterfaceElementsFunctionFile)

	testOutputScaffoldDomainEntityInterfaceGetterFunction = string(testOutputScaffoldDomainEntityInterfaceGetterFunctionFile)
	testOutputScaffoldDomainEntityInterfaceSetterFunction = string(testOutputScaffoldDomainEntityInterfaceSetterFunctionFile)
	testOutputScaffoldDomainEntityInterfaceSetAllSetterFunction = string(testOutputScaffoldDomainEntityInterfaceSetAllSetterFunctionFile)

	testOutputScaffoldDomainEntityInterfaceMarshalJSONFunction = string(testOutputScaffoldDomainEntityInterfaceMarshalJSONFunctionFile)
	testOutputScaffoldDomainEntityInterfaceUnmarshalJSONFunction = string(testOutputScaffoldDomainEntityInterfaceUnmarshalJSONFunctionFile)

}

// Test Entity File
func TestEntityFile(t *testing.T){

	// Build
	statement, err := testEntityGenerator.File(testEntity)

	// Return
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}

	f, err := os.Create("./testing/domain/model/created/" + testOutputDomainEntityFileName)
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputDomainEntityFile {
		t.Errorf(`File() failed; want "%s", got "%s"`, testOutputDomainEntityFile, buf.String())
	}
	
}

// Test Scaffold Entity File
func TestScaffoldEntityFile(t *testing.T){

	// Build
	statement, err := testEntityGenerator.ScaffoldFile(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldEntityFile() failed with error %v`, err)
	}

	f, err := os.Create("./testing/domain/model/created/" + testOutputScaffoldDomainEntityFileName)
	if err != nil {
		t.Errorf(`scaffoldEntityFile() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldEntityFile() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldDomainEntityFile {
		t.Errorf(`scaffoldEntityFile() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityFile, buf.String())
	}
	
}

// Test Scaffold Entity Struct
func TestScaffoldEntityStruct(t *testing.T){

	// Build
	statement, err := testEntityGenerator.scaffoldEntityStruct(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldEntityStruct() failed with error %v`, err)
	}

	f, err := os.Create("./testing/domain/model/created/" + testOutputScaffoldDomainEntityStructName)
	if err != nil {
		t.Errorf(`scaffoldEntityStruct() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldEntityStruct() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldDomainEntityStruct {
		t.Errorf(`scaffoldEntityStruct() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityStruct, buf.String())
	}
	
}

// Test Scaffold Entity Slice Struct
func TestScaffoldEntitySliceStruct(t *testing.T){

	// Build
	statement, err := testEntityGenerator.scaffoldEntitySliceStruct(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldEntitySliceStruct() failed with error %v`, err)
	}

	f, err := os.Create("./testing/domain/model/created/" + testOutputScaffoldDomainEntitySliceStructName)
	if err != nil {
		t.Errorf(`scaffoldEntityStruct() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldEntityStruct() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldDomainEntitySliceStruct {
		t.Errorf(`scaffoldEntitySliceStruct() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntitySliceStruct, buf.String())
	}
	
}

// Test Scaffold Entity Interface
func TestScaffoldEntityInterface(t *testing.T){

	// Build
	statement, err := testEntityGenerator.scaffoldEntityInterface(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldEntityInterface() failed with error %v`, err)
	}

	f, err := os.Create("./testing/domain/model/created/" + testOutputScaffoldDomainEntityInterfaceName)
	if err != nil {
		t.Errorf(`scaffoldEntityStruct() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldEntityStruct() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldDomainEntityInterface {
		t.Errorf(`scaffoldEntityInterface() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityInterface, buf.String())
	}
	
}

// Test Scaffold Entity Slice Interface
func TestScaffoldEntitySliceInterface(t *testing.T){

	// Build
	statement, err := testEntityGenerator.scaffoldEntitySliceInterface(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldEntitySliceInterface() failed with error %v`, err)
	}

	f, err := os.Create("./testing/domain/model/created/" + testOutputScaffoldDomainEntitySliceInterfaceName)
	if err != nil {
		t.Errorf(`scaffoldEntityStruct() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldEntityStruct() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldDomainEntitySliceInterface {
		t.Errorf(`scaffoldEntitySliceInterface() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntitySliceInterface, buf.String())
	}
	
}

// Test Scaffold Entity Interface Constructor Function
func TestScaffoldEntityInterfaceConstructorFunction(t *testing.T){

	// Build
	statement, err := testEntityGenerator.scaffoldEntityInterfaceConstructorFunction(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldEntityInterfaceConstructorFunction() failed with error %v`, err)
	}

	f, err := os.Create("./testing/domain/model/created/" + testOutputScaffoldDomainEntityInterfaceConstructorFunctionName)
	if err != nil {
		t.Errorf(`scaffoldEntityInterfaceConstructorFunction() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldEntityInterfaceConstructorFunction() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldDomainEntityInterfaceConstructorFunction {
		t.Errorf(`scaffoldEntityInterfaceConstructorFunction() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityInterfaceConstructorFunction, buf.String())
	}
	
}

// Test Scaffold Entity Slice Interface Constructor Function
func TestScaffoldEntitySliceInterfaceConstructorFunction(t *testing.T){

	// Build
	statement, err := testEntityGenerator.scaffoldEntitySliceInterfaceConstructorFunction(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldEntitySliceInterfaceConstructorFunction() failed with error %v`, err)
	}

	f, err := os.Create("./testing/domain/model/created/" + testOutputScaffoldDomainEntitySliceInterfaceConstructorFunctionName)
	if err != nil {
		t.Errorf(`scaffoldEntitySliceInterfaceConstructorFunction() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldEntitySliceInterfaceConstructorFunction() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldDomainEntitySliceInterfaceConstructorFunction {
		t.Errorf(`scaffoldEntitySliceInterfaceConstructorFunction() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntitySliceInterfaceConstructorFunction, buf.String())
	}
	
}

// Test Scaffold Entity Slice Interface Len Function
func TestScaffoldEntitySliceInterfaceLenFunction(t *testing.T){

	// Build
	statement, err := testEntityGenerator.scaffoldEntitySliceInterfaceLenFunction(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldEntitySliceInterfaceLenFunction() failed with error %v`, err)
	}

	f, err := os.Create("./testing/domain/model/created/" + testOutputScaffoldDomainEntitySliceInterfaceLenFunctionName)
	if err != nil {
		t.Errorf(`scaffoldEntitySliceInterfaceLenFunction() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldEntitySliceInterfaceLenFunction() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldDomainEntitySliceInterfaceLenFunction {
		t.Errorf(`scaffoldEntitySliceInterfaceLenFunction() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntitySliceInterfaceLenFunction, buf.String())
	}
	
}

// Test Scaffold Entity Slice Interface Append Function
func TestScaffoldEntitySliceInterfaceAppendFunction(t *testing.T){

	// Build
	statement, err := testEntityGenerator.scaffoldEntitySliceInterfaceAppendFunction(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldEntitySliceInterfaceAppendFunction() failed with error %v`, err)
	}

	f, err := os.Create("./testing/domain/model/created/" + testOutputScaffoldDomainEntitySliceInterfaceAppendFunctionName)
	if err != nil {
		t.Errorf(`scaffoldEntitySliceInterfaceAppendFunction() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldEntitySliceInterfaceAppendFunction() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldDomainEntitySliceInterfaceAppendFunction {
		t.Errorf(`scaffoldEntitySliceInterfaceAppendFunction() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntitySliceInterfaceAppendFunction, buf.String())
	}
	
}

// Test Scaffold Entity Slice Interface Elements Function
func TestScaffoldEntitySliceInterfaceElementsFunction(t *testing.T){

	// Build
	statement, err := testEntityGenerator.scaffoldEntitySliceInterfaceElementsFunction(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldEntitySliceInterfaceElementsFunction() failed with error %v`, err)
	}

	f, err := os.Create("./testing/domain/model/created/" + testOutputScaffoldDomainEntitySliceInterfaceElementsFunctionName)
	if err != nil {
		t.Errorf(`scaffoldEntitySliceInterfaceElementsFunction() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldEntitySliceInterfaceElementsFunction() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldDomainEntitySliceInterfaceElementsFunction {
		t.Errorf(`scaffoldEntitySliceInterfaceElementsFunction() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntitySliceInterfaceElementsFunction, buf.String())
	}
	
}

// Test Scaffold Entity Interface Getter Function
func TestScaffoldEntityInterfaceGetterFunction(t *testing.T){

	// Build
	statement, err := testEntityGenerator.scaffoldEntityInterfaceGetterFunction(testField, testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldEntityInterfaceGetterFunction() failed with error %v`, err)
	}

	f, err := os.Create("./testing/domain/model/created/" + testOutputScaffoldDomainEntityInterfaceGetterFunctionName)
	if err != nil {
		t.Errorf(`scaffoldEntityInterfaceGetterFunction() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldEntityInterfaceGetterFunction() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldDomainEntityInterfaceGetterFunction {
		t.Errorf(`scaffoldEntityInterfaceGetterFunction() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityInterfaceGetterFunction, buf.String())
	}
	
}

// Test Scaffold Entity Interface Setter Function
func TestScaffoldEntityInterfaceSetterFunction(t *testing.T){

	// Build
	statement, err := testEntityGenerator.scaffoldEntityInterfaceSetterFunction(testField, testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldEntityInterfaceSetterFunction() failed with error %v`, err)
	}

	f, err := os.Create("./testing/domain/model/created/" + testOutputScaffoldDomainEntityInterfaceSetterFunctionName)
	if err != nil {
		t.Errorf(`scaffoldEntityInterfaceSetterFunction() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldEntityInterfaceSetterFunction() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldDomainEntityInterfaceSetterFunction {
		t.Errorf(`scaffoldEntityInterfaceSetterFunction() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityInterfaceSetterFunction, buf.String())
	}
	
}

// Test Scaffold Entity Interface Set All Setter Function
func TestScaffoldEntityInterfaceSetAllSetterFunction(t *testing.T){

	// Build
	statement, err := testEntityGenerator.scaffoldEntityInterfaceSetAllSetterFunction(testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldEntityInterfaceSetAllSetterFunction() failed with error %v`, err)
	}

	f, err := os.Create("./testing/domain/model/created/" + testOutputScaffoldDomainEntityInterfaceSetAllSetterFunctionName)
	if err != nil {
		t.Errorf(`scaffoldEntityInterfaceSetAllSetterFunction() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`scaffoldEntityInterfaceSetAllSetterFunction() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputScaffoldDomainEntityInterfaceSetAllSetterFunction {
		t.Errorf(`scaffoldEntityInterfaceSetAllSetterFunction() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityInterfaceSetAllSetterFunction, buf.String())
	}
	
}

// // Test Scaffold Entity Interface Marshal JSON Function
// func TestScaffoldEntityInterfaceMarshalJSONFunction(t *testing.T){

// 	// Build
// 	statement, err := testEntityGenerator.scaffoldEntityInterfaceMarshalJSONFunction(testEntity)

// 	// Return
// 	if err != nil {
// 		t.Errorf(`scaffoldEntityInterfaceMarshalJSONFunction() failed with error %v`, err)
// 	}

// 	f, err := os.Create("./testing/domain/model/created/" + testOutputScaffoldDomainEntityInterfaceMarshalJSONFunctionName)
// 	if err != nil {
// 		t.Errorf(`scaffoldEntityInterfaceMarshalJSONFunction() failed with error %v`, err)
// 	}
// 	buf := &bytes.Buffer{}
// 	err = statement.Render(buf)
// 	if err != nil {
// 		t.Errorf(`scaffoldEntityInterfaceMarshalJSONFunction() failed with error %v`, err)
// 	}
// 	_, err = f.Write(buf.Bytes())

// 	if buf.String() != testOutputScaffoldDomainEntityInterfaceMarshalJSONFunction {
// 		t.Errorf(`scaffoldEntityInterfaceMarshalJSONFunction() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityInterfaceMarshalJSONFunction, buf.String())
// 	}
	
// }

// // Test Scaffold Entity Interface Unmarshal JSON Function
// func TestScaffoldEntityInterfaceUnmarshalJSONFunction(t *testing.T){

// 	// Build
// 	statement, err := testEntityGenerator.scaffoldEntityInterfaceUnmarshalJSONFunction(testEntity)

// 	// Return
// 	if err != nil {
// 		t.Errorf(`scaffoldEntityInterfaceUnmarshalJSONFunction() failed with error %v`, err)
// 	}

// 	f, err := os.Create("./testing/domain/model/created/" + testOutputScaffoldDomainEntityInterfaceUnmarshalJSONFunctionName)
// 	if err != nil {
// 		t.Errorf(`scaffoldEntityInterfaceUnmarshalJSONFunction() failed with error %v`, err)
// 	}
// 	buf := &bytes.Buffer{}
// 	err = statement.Render(buf)
// 	if err != nil {
// 		t.Errorf(`scaffoldEntityInterfaceUnmarshalJSONFunction() failed with error %v`, err)
// 	}
// 	_, err = f.Write(buf.Bytes())

// 	if buf.String() != testOutputScaffoldDomainEntityInterfaceUnmarshalJSONFunction {
// 		t.Errorf(`scaffoldEntityInterfaceUnmarshalJSONFunction() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityInterfaceUnmarshalJSONFunction, buf.String())
// 	}
	
// }

// Test Scaffold Entity Struct Field
func TestScaffoldEntityStructField(t *testing.T){

	// Build
	_, err := testEntityGenerator.scaffoldEntityStructField(testField, testEntity)

	// Return
	if err != nil {
		t.Errorf(`scaffoldEntityStructField() failed with error %v`, err)
	}

	// if code.GoString() != testOutputScaffoldDomainEntityStruct {
	// 	t.Errorf(`generateStructField() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityStruct, code.GoString())
	// }
	
}

// // Test Scaffold Entity Getter
// func TestScaffoldEntityInterfaceSetter(t *testing.T){

// 	// Build
// 	code, err := testEntityGenerator.scaffoldEntityInterfaceGetter(testField, testEntity)

// 	// Return
// 	if err != nil {
// 		t.Errorf(`scaffoldEntityInterfaceGetter() failed with error %v`, err)
// 	}

// 	if code.GoString() != testOutputScaffoldDomainEntityInterfaceGetter {
// 		t.Errorf(`scaffoldEntityInterfaceGetter() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityInterfaceGetter, code.GoString())
// 	}
	
// }