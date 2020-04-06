package turnbull

import(
	"testing"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/structure"
	"github.com/brianshepanek/turnbull/generator/jen"
	"github.com/brianshepanek/turnbull/generator/jen/helper"
	generatorInterface "github.com/brianshepanek/turnbull/generator/jen/interface"
	mysqlRepositoryGenerator "github.com/brianshepanek/turnbull/generator/jen/interface/repository/mysql"
	mongoRepositoryGenerator "github.com/brianshepanek/turnbull/generator/jen/interface/repository/mongo"
	redisRepositoryGenerator "github.com/brianshepanek/turnbull/generator/jen/interface/repository/redis"
	defaultPresenterGenerator "github.com/brianshepanek/turnbull/generator/jen/interface/presenter/default"
	httpControllerGenerator "github.com/brianshepanek/turnbull/generator/jen/interface/controller/http"
)

const(
	testConfigPath = "./config"
	testOutputPath = "./output"
)

var (
	testTurnbull *turnbull
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
		},	
	}
)

func init(){
	conf, _ := config.New(testConfigPath, testOutputPath)
	formatter := formatter.New(conf)
	structure := structure.New(formatter)

	testHelperGenerator := helper.New(formatter)

	// Controller
	interfaceControllerGenerators := make(map[string]generatorInterface.ControllerGenerator)
	interfaceControllerGenerators["http"] =  httpControllerGenerator.New(conf, formatter, testHelperGenerator)

	// Presenter
	interfacePresenterGenerators := make(map[string]generatorInterface.PresenterGenerator)
	interfacePresenterGenerators["default"] =  defaultPresenterGenerator.New(conf, formatter, testHelperGenerator)

	// Repository
	interfaceRepositoryGenerators := make(map[string]generatorInterface.RepositoryGenerator)
	interfaceRepositoryGenerators["mongo"] = mongoRepositoryGenerator.New(conf, formatter, testHelperGenerator)
	interfaceRepositoryGenerators["mysql"] = mysqlRepositoryGenerator.New(conf, formatter, testHelperGenerator)
	interfaceRepositoryGenerators["redis"] = redisRepositoryGenerator.New(conf, formatter, testHelperGenerator)

	generator := generator.New(conf, formatter, interfaceControllerGenerators, interfacePresenterGenerators, interfaceRepositoryGenerators)

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

// Test Build Domain Entity
func TestBuildDomainEntity(t *testing.T){

	// Build
	for _, testEntity := range testEntities {
		err := testTurnbull.buildDomainEntity(testEntity)

		// Return
		if err != nil {
			t.Errorf(`buildDomainEntity() failed with error %v`, err)
		}
	}
	
}

// Test Build Scaffold Domain Entity
func TestBuildScaffoldDomainEntity(t *testing.T){

	// Build
	for _, testEntity := range testEntities {
		err := testTurnbull.buildScaffoldDomainEntity(testEntity)

		// Return
		if err != nil {
			t.Errorf(`buildStructure() failed with error %v`, err)
		}
	}	
}

// Test Build Usecase Repository
func TestBuildUsecaseRepository(t *testing.T){

	// Build
	for _, testEntity := range testEntities {
		err := testTurnbull.buildUsecaseRepository(testEntity)

		// Return
		if err != nil {
			t.Errorf(`buildUsecaseRepository() failed with error %v`, err)
		}
	}	
}

// Test Build Scaffold Usecase Repository
func TestBuildScaffoldUsecaseRepository(t *testing.T){

	// Build
	for _, testEntity := range testEntities {
		err := testTurnbull.buildScaffoldUsecaseRepository(testEntity)

		// Return
		if err != nil {
			t.Errorf(`buildScaffoldUsecaseRepository() failed with error %v`, err)
		}
	}	
}

// Test Build Usecase Presenter
func TestBuildUsecasePresenter(t *testing.T){

	// Build
	for _, testEntity := range testEntities {
		err := testTurnbull.buildUsecasePresenter(testEntity)

		// Return
		if err != nil {
			t.Errorf(`buildUsecasePresenter() failed with error %v`, err)
		}
	}	
}

// Test Build Scaffold Usecase Presenter
func TestBuildScaffoldUsecasePresenter(t *testing.T){

	// Build
	for _, testEntity := range testEntities {
		err := testTurnbull.buildScaffoldUsecasePresenter(testEntity)

		// Return
		if err != nil {
			t.Errorf(`buildScaffoldUsecasePresenter() failed with error %v`, err)
		}
	}	
}

// Test Build Usecase Interactor
func TestBuildUsecaseInteractor(t *testing.T){

	// Build
	for _, testEntity := range testEntities {
		err := testTurnbull.buildUsecaseInteractor(testEntity)

		// Return
		if err != nil {
			t.Errorf(`buildUsecaseInteractor() failed with error %v`, err)
		}
	}	
}

// Test Build Scaffold Usecase Interactor
func TestBuildScaffoldUsecaseInteractor(t *testing.T){

	// Build
	for _, testEntity := range testEntities {
		err := testTurnbull.buildScaffoldUsecaseInteractor(testEntity)

		// Return
		if err != nil {
			t.Errorf(`buildScaffoldUsecaseInteractor() failed with error %v`, err)
		}
	}	
}

// Test Build Interface Repository
func TestBuildInterfaceRepository(t *testing.T){

	// Build
	for _, testEntity := range testEntities {

		for _, repository := range testEntity.Repositories {
			
			err := testTurnbull.buildInterfaceRepository(repository.Type, testEntity)

			// Return
			if err != nil {
				t.Errorf(`buildInterfaceRepository() failed with error %v`, err)
			}

		}
		
	}	
}

// Test Build Scaffold Interface Repository
func TestBuildScaffoldInterfaceRepository(t *testing.T){

	// Build
	for _, testEntity := range testEntities {

		for _, repository := range testEntity.Repositories {

			err := testTurnbull.buildScaffoldInterfaceRepository(repository.Type, testEntity)

			// Return
			if err != nil {
				t.Errorf(`buildScaffoldInterfaceRepository() failed with error %v`, err)
			}

		}

	}	
}

// Test Build Interface Repository Entity
func TestBuildInterfaceRepositoryEntity(t *testing.T){

	// Build
	for _, testEntity := range testEntities {

		for _, repository := range testEntity.Repositories {

			err := testTurnbull.buildEntityInterfaceRepository(repository.Type, testEntity)

			// Return
			if err != nil {
				t.Errorf(`buildEntityInterfaceRepository() failed with error %v`, err)
			}

		}

	}	
}

// Test Build Interface Presenter
func TestBuildInterfacePresenter(t *testing.T){

	// Build
	for _, testEntity := range testEntities {
		err := testTurnbull.buildInterfacePresenter("default", testEntity)

		// Return
		if err != nil {
			t.Errorf(`buildInterfacePresenter() failed with error %v`, err)
		}
	}	
}

// Test Build Scaffold Interface Presenter
func TestBuildScaffoldInterfacePresenter(t *testing.T){

	// Build
	for _, testEntity := range testEntities {
		err := testTurnbull.buildScaffoldInterfacePresenter("default", testEntity)

		// Return
		if err != nil {
			t.Errorf(`buildScaffoldInterfacePresenter() failed with error %v`, err)
		}
	}	
}

// Test Build Interface Controller
func TestBuildInterfaceController(t *testing.T){

	// Build
	for _, testEntity := range testEntities {
		err := testTurnbull.buildInterfaceController("http", testEntity)

		// Return
		if err != nil {
			t.Errorf(`buildInterfaceController() failed with error %v`, err)
		}
	}	
}

// Test Build Scaffold Interface Controller
func TestBuildScaffoldInterfaceController(t *testing.T){

	// Build
	for _, testEntity := range testEntities {
		err := testTurnbull.buildScaffoldInterfaceController("http", testEntity)

		// Return
		if err != nil {
			t.Errorf(`buildScaffoldInterfaceController() failed with error %v`, err)
		}
	}	
}

// Test Build Interface Controller Entity
func TestBuildInterfaceControllerEntity(t *testing.T){

	// Build
	for _, testEntity := range testEntities {
		err := testTurnbull.buildInterfaceControllerEntity("http", testEntity)

		// Return
		if err != nil {
			t.Errorf(`buildScaffoldInterfaceController() failed with error %v`, err)
		}
	}	
}