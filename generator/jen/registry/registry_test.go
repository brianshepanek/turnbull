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

	testEntities = []model.Entity{
		model.Entity{
			Name : "post",
			Fields : []model.Field{
				model.Field{
					Name : "id",
					Type : "int64",
					Primary : true,
				},
				model.Field{
					Name : "user_id",
					Type : "int64",
				},
				model.Field{
					Name : "title",
					Type : "string",
				},
				model.Field{
					Name : "subtitle",
					Type : "string",
				},
				model.Field{
					Name : "views",
					Type : "int",
				},
				model.Field{
					Name : "tags",
					Slice : true,
					Type : "string",
				},
				model.Field{
					Name : "created",
					Package : "time",
					Type : "Time",
				},
				model.Field{
					Name : "modified",
					Package : "time",
					Type : "Time",
				},
			},
			Methods : []model.Method {
				model.Method{
					Name : "browse",
					Type : "browse",
				},
				model.Method{
					Name : "read",
					Type : "read",
					Callbacks : []model.Callback {
						model.Callback {
							Type : "before",
						},
					},
				},
				model.Method{
					Name : "edit",
					Type : "edit",
				},
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
					Name : "delete",
					Type : "delete",
				},
			},
			Repositories : []model.Repository {
				model.Repository {
					Type : "mongo",
				},
				model.Repository {
					Type : "mysql",
				},
				model.Repository {
					Type : "redis",
				},
			},
			Controllers : []model.Controller {
				model.Controller {
					Type : "http",
				},
			},
		},	
		model.Entity{
			Name : "comment",
			Interface : true,
			Fields : []model.Field{
				model.Field{
					Name : "id",
					Type : "int64",
					Primary : true,
					Private : true,
				},
				model.Field{
					Name : "post_id",
					Type : "int64",
					Private : true,
				},
				model.Field{
					Name : "user_id",
					Type : "int64",
					Private : true,
				},
				model.Field{
					Name : "title",
					Type : "string",
					Private : true,
				},
				model.Field{
					Name : "body",
					Type : "string",
					Private : true,
				},
				model.Field{
					Name : "created",
					Package : "time",
					Type : "Time",
					Private : true,
				},
				model.Field{
					Name : "modified",
					Package : "time",
					Type : "Time",
					Private : true,
				},
			},
			Methods : []model.Method {
				model.Method{
					Name : "browse",
					Type : "browse",
				},
				model.Method{
					Name : "read",
					Type : "read",
					Callbacks : []model.Callback {
						model.Callback {
							Type : "before",
						},
					},
				},
				model.Method{
					Name : "edit",
					Type : "edit",
				},
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
					Name : "delete",
					Type : "delete",
				},
			},
			Repositories : []model.Repository {
				model.Repository {
					Type : "mongo",
				},
				model.Repository {
					Type : "mysql",
				},
				model.Repository {
					Type : "redis",
				},
			},
			Controllers : []model.Controller {
				model.Controller {
					Type : "http",
				},
			},
		},	
	}
)

var (

	testRegistryGenerator RegistryGenerator

	// Test File Names
	testOutputRegistryFileName = "testOutputRegistryFile"
	testOutputRegistryScaffoldFileName = "testOutputRegistryScaffolFile"

	// Test File Strings
	testOutputRegistryFile string
	testOutputRegistryScaffoldFile string

)


func init(){

	conf, _ := config.New(testConfigPath, testOutputPath)
	formatter := formatter.New(conf)
	testHelperGenerator := helper.New(formatter)
	testRegistryGenerator = New(formatter, testHelperGenerator)

	// Test
	testOutputRegistryFileFile, _ := ioutil.ReadFile("./testing/registry/expected/" + testOutputRegistryFileName)
	testOutputRegistryScaffoldFileFile, _ := ioutil.ReadFile("./testing/registry/expected/" + testOutputRegistryScaffoldFileName)

	testOutputRegistryFile = string(testOutputRegistryFileFile)
	testOutputRegistryScaffoldFile = string(testOutputRegistryScaffoldFileFile)

}

// Test Registry File
func TestRegistryFile(t *testing.T){

	// Build
	statement, err := testRegistryGenerator.File(testEntities)

	// Return
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}

	f, err := os.Create("./testing/registry/created/" + testOutputRegistryFileName)
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`File() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputRegistryFile {
		t.Errorf(`File() failed; want "%s", got "%s"`, testOutputRegistryFile, buf.String())
	}
	
}

// Test Registry Scaffold File
func TestRegistryScaffoldFile(t *testing.T){

	// Build
	statement, err := testRegistryGenerator.ScaffoldFile(testEntities)

	// Return
	if err != nil {
		t.Errorf(`ScaffoldFile() failed with error %v`, err)
	}

	f, err := os.Create("./testing/registry/created/" + testOutputRegistryScaffoldFileName)
	if err != nil {
		t.Errorf(`ScaffoldFile() failed with error %v`, err)
	}
	buf := &bytes.Buffer{}
	err = statement.Render(buf)
	if err != nil {
		t.Errorf(`ScaffoldFile() failed with error %v`, err)
	}
	_, err = f.Write(buf.Bytes())

	if buf.String() != testOutputRegistryScaffoldFile {
		t.Errorf(`ScaffoldFile() failed; want "%s", got "%s"`, testOutputRegistryScaffoldFile, buf.String())
	}
	
}