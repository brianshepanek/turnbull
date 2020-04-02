package generator

import(
	"bytes"
	"testing"
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/generator"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/generator/jen/helper"
	scribbleRepositoryGenerator "github.com/brianshepanek/turnbull/generator/jen/interface/repository/scribble"
	defaultPresenterGenerator "github.com/brianshepanek/turnbull/generator/jen/interface/presenter/default"
	httpControllerGenerator "github.com/brianshepanek/turnbull/generator/jen/interface/controller/http"
	generatorInterface "github.com/brianshepanek/turnbull/generator/jen/interface"
)

const(
	testConfigPath = "/go/src/github.com/brianshepanek/turnbull/_testing/config"
	testOutputPath = "/go/src/github.com/brianshepanek/turnbull/_testing/output"
)

var(


	testGenerator generator.Generator

	// Test Entity
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


func init(){
	conf, _ := config.New(testConfigPath, testOutputPath)
	formatter := formatter.New(conf)
	testHelperGenerator := helper.New(formatter)
	interfacePresenterGenerator := defaultPresenterGenerator.New(conf, formatter, testHelperGenerator)
	interfaceControllerGenerator := httpControllerGenerator.New(conf, formatter, testHelperGenerator)
	interfaceRepositoryGenerators := make(map[string]generatorInterface.RepositoryGenerator)
	interfaceRepositoryGenerators["scribble"] = scribbleRepositoryGenerator.New(conf, formatter, testHelperGenerator)
	testGenerator = New(conf, formatter, interfaceControllerGenerator, interfacePresenterGenerator, interfaceRepositoryGenerators)
}

// Test Scaffold Interface Repository File
func TestScaffoldEntity(t *testing.T){

	// Build
	buf := &bytes.Buffer{}
	err := testGenerator.ScaffoldEntity(testEntity, buf)

	// Return
	if err != nil {
		t.Errorf(`scaffoldInterfaceRepositoryFile() failed with error %v`, err)
	}
	
}