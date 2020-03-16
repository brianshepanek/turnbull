package turnbull

import(
	"testing"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/structure"
	"github.com/brianshepanek/turnbull/generator/jen"
	"github.com/brianshepanek/turnbull/generator/jen/helper"
	scribbleRepositoryGenerator "github.com/brianshepanek/turnbull/generator/jen/interface/repository/scribble"
	defaultPresenterGenerator "github.com/brianshepanek/turnbull/generator/jen/interface/presenter/default"
	httpControllerGenerator "github.com/brianshepanek/turnbull/generator/jen/interface/controller/http"
)

const(
	testConfigPath = "./config"
	testOutputPath = "./output"
)

var (
	testTurnbull *turnbull
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
)

func init(){
	conf, _ := config.New(testConfigPath, testOutputPath)
	formatter := formatter.New(conf)
	structure := structure.New(formatter)

	testHelperGenerator := helper.New(formatter)
	interfaceRepositoryGenerator := scribbleRepositoryGenerator.New(conf, formatter, testHelperGenerator)
	interfacePresenterGenerator := defaultPresenterGenerator.New(conf, formatter, testHelperGenerator)
	interfaceControllerGenerator := httpControllerGenerator.New(conf, formatter, testHelperGenerator)
	generator := generator.New(conf, formatter, interfaceRepositoryGenerator, interfacePresenterGenerator, interfaceControllerGenerator)

	testTurnbull = New(formatter, structure, generator)
}

// Test Build Structture
func TestBuildStructture(t *testing.T){

	// Build
	err := testTurnbull.buildStructure()

	// Return
	if err != nil {
		t.Errorf(`buildStructure() failed with error %v`, err)
	}
}

// Test Build Domain Entity
func TestBuildDomainEntity(t *testing.T){

	// Build
	err := testTurnbull.buildDomainEntity(testEntity)

	// Return
	if err != nil {
		t.Errorf(`buildStructure() failed with error %v`, err)
	}
}