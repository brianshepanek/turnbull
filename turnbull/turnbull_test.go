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
				Name : "id",
				Type : "string",
			},
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
)

func init(){
	conf, _ := config.New(testConfigPath, testOutputPath)
	formatter := formatter.New(conf)
	structure := structure.New(formatter)

	testHelperGenerator := helper.New(formatter)
	interfaceRepositoryGenerator := scribbleRepositoryGenerator.New(conf, formatter, testHelperGenerator)
	interfacePresenterGenerator := defaultPresenterGenerator.New(conf, formatter, testHelperGenerator)
	interfaceControllerGenerator := httpControllerGenerator.New(conf, formatter, testHelperGenerator)
	generator := generator.New(conf, formatter, interfaceControllerGenerator, interfacePresenterGenerator, interfaceRepositoryGenerator)

	testTurnbull = New(formatter, structure, generator)
}

// Test Build Structure
func TestBuildStructure(t *testing.T){

	// Build
	err := testTurnbull.buildStructure()

	// Return
	if err != nil {
		t.Errorf(`buildStructure() failed with error %v`, err)
	}
}

// Test Build Scaffold Domain Entity
func TestBuildScaffoldDomainEntity(t *testing.T){

	// Build
	err := testTurnbull.buildScaffoldDomainEntity(testEntity)

	// Return
	if err != nil {
		t.Errorf(`buildStructure() failed with error %v`, err)
	}
}

// Test Build Scaffold Usecase Repository
func TestBuildScaffoldUsecaseRepository(t *testing.T){

	// Build
	err := testTurnbull.buildScaffoldUsecaseRepository(testEntity)

	// Return
	if err != nil {
		t.Errorf(`buildScaffoldUsecaseRepository() failed with error %v`, err)
	}
}

// Test Build Scaffold Usecase Presenter
func TestBuildScaffoldUsecasePresenter(t *testing.T){

	// Build
	err := testTurnbull.buildScaffoldUsecasePresenter(testEntity)

	// Return
	if err != nil {
		t.Errorf(`buildScaffoldUsecasePresenter() failed with error %v`, err)
	}
}

// Test Build Scaffold Usecase Interactor
func TestBuildScaffoldUsecaseInteractor(t *testing.T){

	// Build
	err := testTurnbull.buildScaffoldUsecaseInteractor(testEntity)

	// Return
	if err != nil {
		t.Errorf(`buildScaffoldUsecaseInteractor() failed with error %v`, err)
	}
}

// Test Build Scaffold Interface Repository
func TestBuildScaffoldInterfaceRepository(t *testing.T){

	// Build
	err := testTurnbull.buildScaffoldInterfaceRepository("scribble", testEntity)

	// Return
	if err != nil {
		t.Errorf(`buildScaffoldInterfaceRepository() failed with error %v`, err)
	}
}

// Test Build Scaffold Interface Presenter
func TestBuildScaffoldInterfacePresenter(t *testing.T){

	// Build
	err := testTurnbull.buildScaffoldInterfacePresenter("default", testEntity)

	// Return
	if err != nil {
		t.Errorf(`buildScaffoldInterfacePresenter() failed with error %v`, err)
	}
}

// Test Build Scaffold Interface Controller
func TestBuildScaffoldInterfaceController(t *testing.T){

	// Build
	err := testTurnbull.buildScaffoldInterfaceController("http", testEntity)

	// Return
	if err != nil {
		t.Errorf(`buildScaffoldInterfaceController() failed with error %v`, err)
	}
}