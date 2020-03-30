package formatter

import(
	"testing"
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/domain/model"
)

const(

	testConfigPath = "/go/src/github.com/brianshepanek/turnbull/_testing/config"
	testOutputPath = "/go/src/github.com/brianshepanek/turnbull/_testing/output"

	testOutputDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output"
	testOutputScaffoldDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output"

	
	testOutputRegistryDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/registry"
	testOutputRegistryEntityFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/registry/foo_registry.go"
	testOutputRegistryEntityRepositoryConstructorFunctionId = "NewFooRepository"
	testOutputRegistryEntityPresenterConstructorFunctionId = "NewFooPresenter"

	testOutputScaffoldDomainDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/domain"
	testOutputScaffoldDomainEntityDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/domain/entity"
	testOutputScaffoldDomainEntityDirectoryImportPath = "github.com/brianshepanek/turnbull/_testing/output/domain/entity"
	testOutputDomainEntityFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/domain/entity/foo.go"
	testOutputScaffoldDomainEntityFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/domain/entity/foo_scaffold.go"

	testOutputScaffoldUsecaseDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/usecase"
	testOutputScaffoldUsecaseInteractorDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/usecase/interactor"
	testOutputScaffoldUsecaseInteractorDirectoryImportPath = "github.com/brianshepanek/turnbull/_testing/output/usecase/interactor"
	testOutputUsecaseInteractorFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/usecase/interactor/foo_interactor.go"
	testOutputScaffoldUsecaseInteractorFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/usecase/interactor/foo_interactor_scaffold.go"
	testOutputScaffoldUsecaseRepositoryDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/usecase/repository"
	testOutputScaffoldUsecaseRepositoryDirectoryImportPath = "github.com/brianshepanek/turnbull/_testing/output/usecase/repository"
	testOutputUsecaseRepositoryFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/usecase/repository/foo_repository.go"
	testOutputScaffoldUsecaseRepositoryFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/usecase/repository/foo_repository_scaffold.go"
	testOutputScaffoldUsecasePresenterDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/usecase/presenter"
	testOutputScaffoldUsecasePresenterDirectoryImportPath = "github.com/brianshepanek/turnbull/_testing/output/usecase/presenter"
	testOutputUsecasePresenterFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/usecase/presenter/foo_presenter.go"
	testOutputScaffoldUsecasePresenterFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/usecase/presenter/foo_presenter_scaffold.go"

	testOutputScaffoldInterfaceDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/interface"
	testOutputScaffoldInterfaceControllerDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/interface/controller"
	testOutputInterfaceControllerFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/interface/controller/foo_http_controller.go"
	testOutputScaffoldInterfaceControllerFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/interface/controller/foo_http_controller_scaffold.go"
	testOutputScaffoldInterfaceRepositoryDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/interface/repository"
	testOutputInterfaceRepositoryFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/interface/repository/foo_scribble_repository.go"
	testOutputScaffoldInterfaceRepositoryFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/interface/repository/foo_scribble_repository_scaffold.go"
	testOutputScaffoldInterfacePresenterDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/interface/presenter"
	testOutputInterfacePresenterFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/interface/presenter/foo_default_presenter.go"
	testOutputScaffoldInterfacePresenterFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/interface/presenter/foo_default_presenter_scaffold.go"

	testOutputRegistryPackageName = "registry"

	testOutputScaffoldDomainEntityPackageName = "entity"

	testOutputDomainEntityStructId = "foo"
	testOutputDomainEntitySliceStructId = "foos"

	testOutputDomainEntityInterfaceId = "Foo"
	testOutputDomainEntitySliceInterfaceId = "Foos"

	testOutputScaffoldDomainEntityStructId = "fooStruct"
	testOutputScaffoldDomainEntitySliceStructId = "foosStruct"

	testOutputScaffoldDomainEntityInterfaceId = "fooInterface"
	testOutputScaffoldDomainEntitySliceInterfaceId = "foosInterface"

	testOutputScaffoldDomainEntityStructFieldId = "bar"
	testOutputScaffoldDomainEntityGetterId = "Bar"
	testOutputScaffoldDomainEntityJSONTagId = "bar"
	testOutputScaffoldDomainEntitySetterId = "SetBar"
	testOutputScaffoldDomainEntitySetAllSetterId = "SetAll"

	testOutputScaffoldDomainEntityLenId = "Len"
	testOutputScaffoldDomainEntityAppendId = "Append"
	testOutputScaffoldDomainEntityElementsId = "Elements"

	testOutputDomainEntityInterfaceConstructorFunctionId = "NewFoo"
	testOutputDomainEntitySliceInterfaceConstructorFunctionId = "NewFoos"
	testOutputScaffoldDomainEntityInterfaceConstructorFunctionId = "NewFooScaffoldStruct"
	testOutputScaffoldDomainEntitySliceInterfaceConstructorFunctionId = "NewFoosScaffoldStruct"

	testOutputScaffoldUsecaseRepositoryPackageName = "repository"
	testOutputScaffoldUsecaseRepositoryInterfaceId = "fooRepository"
	testOutputUsecaseRepositoryInterfaceId = "FooRepository"
	testOutputScaffoldUsecaseRepositoryInterfaceMethodId = "Add"

	testOutputScaffoldUsecasePresenterPackageName = "presenter"
	testOutputScaffoldUsecasePresenterInterfaceId = "fooPresenter"
	testOutputUsecasePresenterInterfaceId = "FooPresenter"
	testOutputScaffoldUsecasePresenterInterfaceMethodId = "Add"

	testOutputScaffoldUsecaseInteractorPackageName = "interactor"
	testOutputUsecaseInteractorStructId = "fooInteractor"
	testOutputScaffoldUsecaseInteractorStructId = "fooInteractorStruct"
	testOutputUsecaseInteractorInterfaceId = "FooInteractor"
	testOutputScaffoldUsecaseInteractorInterfaceId = "fooInteractorInterface"
	testOutputScaffoldUsecaseInteractorInterfaceMethodId = "Add"
	testOutputUsecaseInteractorInterfaceConstructorFunctionId = "NewFooInteractor"

	testOutputScaffoldInterfaceRepositoryPackageName = "repository"
	testOutputInterfaceRepositoryStructId = "scribbleFooRepository"
	testOutputScaffoldInterfaceRepositoryStructId = "scribbleFooRepositoryStruct"
	testOutputScaffoldInterfaceRepositoryConstructorFunctionId = "NewScribbleFooRepository"

	testOutputScaffoldInterfacePresenterPackageName = "presenter"
	testOutputInterfacePresenterStructId = "defaultFooPresenter"
	testOutputScaffoldInterfacePresenterStructId = "defaultFooPresenterStruct"
	testOutputScaffoldInterfacePresenterInterfaceId = "DefaultFooPresenter"
	testOutputScaffoldInterfacePresenterConstructorFunctionId = "NewDefaultFooPresenter"

	testOutputScaffoldInterfaceControllerPackageName = "controller"
	testOutputInterfaceControllerStructId = "httpFooController"
	testOutputScaffoldInterfaceControllerStructId = "httpFooControllerStruct"
	testOutputInterfaceControllerInterfaceId = "HttpFooController"
	testOutputScaffoldInterfaceControllerInterfaceId = "httpFooControllerInterface"
	testOutputScaffoldInterfaceControllerConstructorFunctionId = "NewHttpFooController"

)

var (
	testFormatter Formatter
	testEntity  = model.Entity{
		Name : "foo",
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
	testFormatter = New(conf)
	
}

// Test Output Directory
func TestOutputDirectory(t *testing.T){

	// Build

	dir, err := testFormatter.OutputDirectory()

	// Return
	if err != nil {
		t.Errorf(`OutputDirectory() failed with error %v`, err)
	}

	if dir != testOutputDirectory {
		t.Errorf(`OutputDirectory() failed; want "%s", got "%s"`, testOutputDirectory, dir)
	}
}

// Test Output Registry Directory
func TestOutputRegistryDirectory(t *testing.T){

	// Build

	dir, err := testFormatter.OutputRegistryDirectory()

	// Return
	if err != nil {
		t.Errorf(`OutputRegistryDirectory() failed with error %v`, err)
	}

	if dir != testOutputRegistryDirectory {
		t.Errorf(`OutputRegistryDirectory() failed; want "%s", got "%s"`, testOutputRegistryDirectory, dir)
	}
}

// Test Output Registry Entity File
func TestOutputRegistryEntityFile(t *testing.T){

	// Build

	file, err := testFormatter.OutputRegistryEntityFile(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputRegistryEntityFile() failed with error %v`, err)
	}

	if file != testOutputRegistryEntityFile {
		t.Errorf(`OutputRegistryEntityFile() failed; want "%s", got "%s"`, testOutputRegistryEntityFile, file)
	}
}

// Test Output Registry Entity Presenter Constructor Function ID
func TestOutputRegistryEntityPresenterConstructorFunctionId(t *testing.T){

	// Build

	file, err := testFormatter.OutputRegistryEntityPresenterConstructorFunctionId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputRegistryEntityPresenterConstructorFunctionId() failed with error %v`, err)
	}

	if file != testOutputRegistryEntityPresenterConstructorFunctionId {
		t.Errorf(`OutputRegistryEntityPresenterConstructorFunctionId() failed; want "%s", got "%s"`, testOutputRegistryEntityPresenterConstructorFunctionId, file)
	}
}

// Test Output Registry Entity Repository Constructor Function ID
func TestOutputRegistryEntityRepositoryConstructorFunctionId(t *testing.T){

	// Build

	file, err := testFormatter.OutputRegistryEntityRepositoryConstructorFunctionId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputRegistryEntityRepositoryConstructorId() failed with error %v`, err)
	}

	if file != testOutputRegistryEntityRepositoryConstructorFunctionId {
		t.Errorf(`OutputRegistryEntityRepositoryConstructorId() failed; want "%s", got "%s"`, testOutputRegistryEntityRepositoryConstructorFunctionId, file)
	}
}

// Test Output Scaffold Directory
func TestOutputScaffoldDirectory(t *testing.T){

	// Build

	dir, err := testFormatter.OutputScaffoldDirectory()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDirectory() failed with error %v`, err)
	}

	if dir != testOutputScaffoldDirectory {
		t.Errorf(`OutputScaffoldDirectory() failed; want "%s", got "%s"`, testOutputScaffoldDirectory, dir)
	}
}

// Test Output Scaffold Domain Directory
func TestOutputScaffoldDomainDirectory(t *testing.T){

	// Build

	dir, err := testFormatter.OutputScaffoldDomainDirectory()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainDirectory() failed with error %v`, err)
	}

	if dir != testOutputScaffoldDomainDirectory {
		t.Errorf(`OutputScaffoldDomainDirectory() failed; want "%s", got "%s"`, testOutputScaffoldDomainDirectory, dir)
	}
}

// Test Output Scaffold Domain Entity Directory
func TestOutputScaffoldDomainEntityDirectory(t *testing.T){

	// Build

	dir, err := testFormatter.OutputScaffoldDomainEntityDirectory()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntityDirectory() failed with error %v`, err)
	}

	if dir != testOutputScaffoldDomainEntityDirectory {
		t.Errorf(`OutputScaffoldDomainEntityDirectory() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityDirectory, dir)
	}
}

// Test Output Scaffold Domain Entity Directory Import Path
func TestOutputScaffoldDomainEntityDirectoryImportPath(t *testing.T){

	// Build

	dir, err := testFormatter.OutputScaffoldDomainEntityDirectoryImportPath()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntityDirectoryImportPath() failed with error %v`, err)
	}

	if dir != testOutputScaffoldDomainEntityDirectoryImportPath {
		t.Errorf(`OutputScaffoldDomainEntityDirectoryImportPath() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityDirectoryImportPath, dir)
	}
}

// Test Output Domain Entity File
func TestOutputDomainEntityFile(t *testing.T){

	// Build

	file, err := testFormatter.OutputDomainEntityFile(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputDomainEntityFile() failed with error %v`, err)
	}

	if file != testOutputDomainEntityFile {
		t.Errorf(`OutputDomainEntityFile() failed; want "%s", got "%s"`, testOutputDomainEntityFile, file)
	}
}

// Test Output Scaffold Domain Entity File
func TestOutputScaffoldDomainEntityFile(t *testing.T){

	// Build

	file, err := testFormatter.OutputScaffoldDomainEntityFile(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntityFile() failed with error %v`, err)
	}

	if file != testOutputScaffoldDomainEntityFile {
		t.Errorf(`OutputScaffoldDomainEntityFile() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityFile, file)
	}
}

// Test Output Scaffold Usecase Directory
func TestOutputScaffoldUsecaseDirectory(t *testing.T){

	// Build

	dir, err := testFormatter.OutputScaffoldUsecaseDirectory()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecaseDirectory() failed with error %v`, err)
	}

	if dir != testOutputScaffoldUsecaseDirectory {
		t.Errorf(`OutputScaffoldUsecaseDirectory() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseDirectory, dir)
	}
}

// Test Output Scaffold Usecase Interactor Directory
func TestOutputScaffoldUsecaseInteractorDirectory(t *testing.T){

	// Build

	dir, err := testFormatter.OutputScaffoldUsecaseInteractorDirectory()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecaseInteractorDirectory() failed with error %v`, err)
	}

	if dir != testOutputScaffoldUsecaseInteractorDirectory {
		t.Errorf(`OutputScaffoldUsecaseInteractorDirectory() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseInteractorDirectory, dir)
	}
}

// Test Output Scaffold Usecase Interactor Directory Import Path
func TestOutputScaffoldUsecaseInteractorDirectoryImportPath(t *testing.T){

	// Build

	dir, err := testFormatter.OutputScaffoldUsecaseInteractorDirectoryImportPath()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecOutputScaffoldUsecaseInteractorDirectoryImportPathaseRepositoryDirectoryImportPath() failed with error %v`, err)
	}

	if dir != testOutputScaffoldUsecaseInteractorDirectoryImportPath {
		t.Errorf(`OutputScaffoldUsecaseInteractorDirectoryImportPath() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseInteractorDirectoryImportPath, dir)
	}
}

// Test Output Usecase Interactor File
func TestOutputUsecaseInteractorFile(t *testing.T){

	// Build

	file, err := testFormatter.OutputUsecaseInteractorFile(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputUsecaseInteractorFile() failed with error %v`, err)
	}

	if file != testOutputUsecaseInteractorFile {
		t.Errorf(`OutputUsecaseInteractorFile() failed; want "%s", got "%s"`, testOutputUsecaseInteractorFile, file)
	}
}

// Test Output Scaffold Usecase Interactor File
func TestOutputScaffoldUsecaseInteractorFile(t *testing.T){

	// Build

	file, err := testFormatter.OutputScaffoldUsecaseInteractorFile(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecaseInteractorFile() failed with error %v`, err)
	}

	if file != testOutputScaffoldUsecaseInteractorFile {
		t.Errorf(`OutputScaffoldUsecaseInteractorFile() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseInteractorFile, file)
	}
}

// Test Output Scaffold Usecase Repository Directory
func TestOutputScaffoldUsecaseRepositoryDirectory(t *testing.T){

	// Build

	dir, err := testFormatter.OutputScaffoldUsecaseRepositoryDirectory()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecaseRepositoryDirectory() failed with error %v`, err)
	}

	if dir != testOutputScaffoldUsecaseRepositoryDirectory {
		t.Errorf(`OutputScaffoldUsecaseRepositoryDirectory() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseRepositoryDirectory, dir)
	}
}

// Test Output Scaffold Usecase Repository Directory Import Path
func TestOutputScaffoldUsecaseRepositoryDirectoryImportPath(t *testing.T){

	// Build

	dir, err := testFormatter.OutputScaffoldUsecaseRepositoryDirectoryImportPath()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecaseRepositoryDirectoryImportPath() failed with error %v`, err)
	}

	if dir != testOutputScaffoldUsecaseRepositoryDirectoryImportPath {
		t.Errorf(`OutputScaffoldUsecaseRepositoryDirectoryImportPath() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseRepositoryDirectoryImportPath, dir)
	}
}

// Test Output Usecase Repository File
func TestOutputUsecaseRepositoryFile(t *testing.T){

	// Build

	file, err := testFormatter.OutputUsecaseRepositoryFile(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputUsecaseRepositoryFile() failed with error %v`, err)
	}

	if file != testOutputUsecaseRepositoryFile {
		t.Errorf(`OutputUsecaseRepositoryFile() failed; want "%s", got "%s"`, testOutputUsecaseRepositoryFile, file)
	}
}

// Test Output Scaffold Usecase Repository File
func TestOutputScaffoldUsecaseRepositoryFile(t *testing.T){

	// Build

	file, err := testFormatter.OutputScaffoldUsecaseRepositoryFile(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecaseRepositoryFile() failed with error %v`, err)
	}

	if file != testOutputScaffoldUsecaseRepositoryFile {
		t.Errorf(`OutputScaffoldUsecaseRepositoryFile() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseRepositoryFile, file)
	}
}

// Test Output Scaffold Usecase Presenter Directory
func TestOutputScaffoldUsecasePresenterDirectory(t *testing.T){

	// Build

	dir, err := testFormatter.OutputScaffoldUsecasePresenterDirectory()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecasePresenterDirectory() failed with error %v`, err)
	}

	if dir != testOutputScaffoldUsecasePresenterDirectory {
		t.Errorf(`OutputScaffoldUsecasePresenterDirectory() failed; want "%s", got "%s"`, testOutputScaffoldUsecasePresenterDirectory, dir)
	}
}

// Test Output Scaffold Usecase Presenter Directory Import Path
func TestOutputScaffoldUsecasePresenterDirectoryImportPath(t *testing.T){

	// Build

	dir, err := testFormatter.OutputScaffoldUsecasePresenterDirectoryImportPath()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecasePresenterDirectoryImportPath() failed with error %v`, err)
	}

	if dir != testOutputScaffoldUsecasePresenterDirectoryImportPath {
		t.Errorf(`OutputScaffoldUsecasePresenterDirectoryImportPath() failed; want "%s", got "%s"`, testOutputScaffoldUsecasePresenterDirectoryImportPath, dir)
	}
}

// Test Output Usecase Presenter File
func TestOutputUsecasePresenterFile(t *testing.T){

	// Build

	file, err := testFormatter.OutputUsecasePresenterFile(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputUsecasePresenterFile() failed with error %v`, err)
	}

	if file != testOutputUsecasePresenterFile {
		t.Errorf(`OutputUsecasePresenterFile() failed; want "%s", got "%s"`, testOutputUsecasePresenterFile, file)
	}
}

// Test Output Scaffold Usecase Presenter File
func TestOutputScaffoldUsecasePresenterFile(t *testing.T){

	// Build

	file, err := testFormatter.OutputScaffoldUsecasePresenterFile(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecasePresenterFile() failed with error %v`, err)
	}

	if file != testOutputScaffoldUsecasePresenterFile {
		t.Errorf(`OutputScaffoldUsecasePresenterFile() failed; want "%s", got "%s"`, testOutputScaffoldUsecasePresenterFile, file)
	}
}

// Test Output Scaffold Interface Directory
func TestOutputScaffoldInterfaceDirectory(t *testing.T){

	// Build

	dir, err := testFormatter.OutputScaffoldInterfaceDirectory()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfaceDirectory() failed with error %v`, err)
	}

	if dir != testOutputScaffoldInterfaceDirectory {
		t.Errorf(`OutputScaffoldInterfaceDirectory() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceDirectory, dir)
	}
}

// Test Output Scaffold Interface Controller Directory
func TestOutputScaffoldInterfaceControllerDirectory(t *testing.T){

	// Build

	dir, err := testFormatter.OutputScaffoldInterfaceControllerDirectory()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfaceControllerDirectory() failed with error %v`, err)
	}

	if dir != testOutputScaffoldInterfaceControllerDirectory {
		t.Errorf(`OutputScaffoldInterfaceControllerDirectory() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceControllerDirectory, dir)
	}
}

// Test Output Interface Controller File
func TestOutputInterfaceRepositoryController(t *testing.T){

	// Build

	file, err := testFormatter.OutputInterfaceControllerFile("http", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputInterfaceControllerFile() failed with error %v`, err)
	}

	if file != testOutputInterfaceControllerFile {
		t.Errorf(`OutputInterfaceControllerFile() failed; want "%s", got "%s"`, testOutputInterfaceControllerFile, file)
	}
}

// Test Output Scaffold Interface Controller File
func TestOutputScaffoldInterfaceRepositoryController(t *testing.T){

	// Build

	file, err := testFormatter.OutputScaffoldInterfaceControllerFile("http", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfaceControllerFile() failed with error %v`, err)
	}

	if file != testOutputScaffoldInterfaceControllerFile {
		t.Errorf(`OutputScaffoldInterfaceControllerFile() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceControllerFile, file)
	}
}

// Test Output Scaffold Interface Repository Directory
func TestOutputScaffoldInterfaceRepositoryDirectory(t *testing.T){

	// Build

	dir, err := testFormatter.OutputScaffoldInterfaceRepositoryDirectory()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfaceRepositoryDirectory() failed with error %v`, err)
	}

	if dir != testOutputScaffoldInterfaceRepositoryDirectory {
		t.Errorf(`OutputScaffoldInterfaceRepositoryDirectory() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceRepositoryDirectory, dir)
	}
}

// Test Output Interface Repository File
func TestOutputInterfaceRepositoryFile(t *testing.T){

	// Build

	file, err := testFormatter.OutputInterfaceRepositoryFile("scribble", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputInterfaceRepositoryFile() failed with error %v`, err)
	}

	if file != testOutputInterfaceRepositoryFile {
		t.Errorf(`OutputInterfaceRepositoryFile() failed; want "%s", got "%s"`, testOutputInterfaceRepositoryFile, file)
	}
}

// Test Output Scaffold Interface Repository File
func TestOutputScaffoldInterfaceRepositoryFile(t *testing.T){

	// Build

	file, err := testFormatter.OutputScaffoldInterfaceRepositoryFile("scribble", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfaceRepositoryFile() failed with error %v`, err)
	}

	if file != testOutputScaffoldInterfaceRepositoryFile {
		t.Errorf(`OutputScaffoldInterfaceRepositoryFile() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceRepositoryFile, file)
	}
}

// Test Output Scaffold Interface Presenter Directory
func TestOutputScaffoldInterfacePresenterDirectory(t *testing.T){

	// Build

	dir, err := testFormatter.OutputScaffoldInterfacePresenterDirectory()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfacePresenterDirectory() failed with error %v`, err)
	}

	if dir != testOutputScaffoldInterfacePresenterDirectory {
		t.Errorf(`OutputScaffoldInterfacePresenterDirectory() failed; want "%s", got "%s"`, testOutputScaffoldInterfacePresenterDirectory, dir)
	}
}

// Test Output Interface Presenter File
func TestOutputInterfaceRepositoryPresenter(t *testing.T){

	// Build

	file, err := testFormatter.OutputInterfacePresenterFile("default", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputInterfacePresenterFile() failed with error %v`, err)
	}

	if file != testOutputInterfacePresenterFile {
		t.Errorf(`OutputInterfacePresenterFile() failed; want "%s", got "%s"`, testOutputInterfacePresenterFile, file)
	}
}

// Test Output Scaffold Interface Presenter File
func TestOutputScaffoldInterfaceRepositoryPresenter(t *testing.T){

	// Build

	file, err := testFormatter.OutputScaffoldInterfacePresenterFile("default", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfacePresenterFile() failed with error %v`, err)
	}

	if file != testOutputScaffoldInterfacePresenterFile {
		t.Errorf(`OutputScaffoldInterfacePresenterFile() failed; want "%s", got "%s"`, testOutputScaffoldInterfacePresenterFile, file)
	}
}

// Test Output Registry Package Name
func TestOutputRegistryPackageName(t *testing.T){

	// Build
	id, err := testFormatter.OutputRegistryPackageName()

	// Return
	if err != nil {
		t.Errorf(`OutputRegistryPackageName() failed with error %v`, err)
	}

	if id != testOutputRegistryPackageName {
		t.Errorf(`OutputRegistryPackageName() failed; want "%s", got "%s"`, testOutputRegistryPackageName, id)
	}
}

// Test Output Scaffold Domain Entity Package Name
func TestOutputScaffoldDomainEntityPackageName(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldDomainEntityPackageName()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntityPackageName() failed with error %v`, err)
	}

	if id != testOutputScaffoldDomainEntityPackageName {
		t.Errorf(`OutputScaffoldDomainEntityPackageName() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityPackageName, id)
	}
}

// Test Output Domain Entity Struct ID
func TestOutputDomainEntityStructId(t *testing.T){

	// Build
	id, err := testFormatter.OutputDomainEntityStructId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputDomainEntityStructId() failed with error %v`, err)
	}

	if id != testOutputDomainEntityStructId {
		t.Errorf(`OutputDomainEntityStructId() failed; want "%s", got "%s"`, testOutputDomainEntityStructId, id)
	}
}

// Test Output Domain Entity Slice Struct ID
func TestOutputDomainEntitySliceStructId(t *testing.T){

	// Build
	id, err := testFormatter.OutputDomainEntitySliceStructId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputDomainEntitySliceStructId() failed with error %v`, err)
	}

	if id != testOutputDomainEntitySliceStructId {
		t.Errorf(`OutputDomainEntitySliceStructId() failed; want "%s", got "%s"`, testOutputDomainEntitySliceStructId, id)
	}
}


// Test Output Domain Entity Interface ID
func TestOutputDomainEntityInterfaceId(t *testing.T){

	// Build
	id, err := testFormatter.OutputDomainEntityInterfaceId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputDomainEntityInterfaceId() failed with error %v`, err)
	}

	if id != testOutputDomainEntityInterfaceId {
		t.Errorf(`OutputDomainEntityInterfaceId() failed; want "%s", got "%s"`, testOutputDomainEntityInterfaceId, id)
	}
}

// Test Output Domain Entity Slice InterfaceID
func TestOutputDomainEntitySliceInterfaceId(t *testing.T){

	// Build
	id, err := testFormatter.OutputDomainEntitySliceInterfaceId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputDomainEntitySliceInterfaceId() failed with error %v`, err)
	}

	if id != testOutputDomainEntitySliceInterfaceId {
		t.Errorf(`OutputDomainEntitySliceInterfaceId() failed; want "%s", got "%s"`, testOutputDomainEntitySliceInterfaceId, id)
	}
}

// Test Output Scaffold Domain Entity Struct ID
func TestOutputScaffoldDomainEntityStructId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldDomainEntityStructId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntityStructId() failed with error %v`, err)
	}

	if id != testOutputScaffoldDomainEntityStructId {
		t.Errorf(`OutputScaffoldDomainEntityStructId() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityStructId, id)
	}
}

// Test Output Scaffold Domain Entity Slice Struct ID
func TestOutputScaffoldDomainEntitySliceStructId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldDomainEntitySliceStructId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntitySliceStructId() failed with error %v`, err)
	}

	if id != testOutputScaffoldDomainEntitySliceStructId {
		t.Errorf(`OutputScaffoldDomainEntitySliceStructId() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntitySliceStructId, id)
	}
}


// Test Output Scaffold Domain Entity Interface ID
func TestOutputScaffoldDomainEntityInterfaceId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldDomainEntityInterfaceId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntityInterfaceId() failed with error %v`, err)
	}

	if id != testOutputScaffoldDomainEntityInterfaceId {
		t.Errorf(`OutputScaffoldDomainEntityInterfaceId() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityInterfaceId, id)
	}
}

// Test Output Scaffold Domain Entity Slice InterfaceID
func TestOutputScaffoldDomainEntitySliceInterfaceId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldDomainEntitySliceInterfaceId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntitySliceInterfaceId() failed with error %v`, err)
	}

	if id != testOutputScaffoldDomainEntitySliceInterfaceId {
		t.Errorf(`OutputScaffoldDomainEntitySliceInterfaceId() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntitySliceInterfaceId, id)
	}
}

// Test Output Scaffold Domain Entity Field ID
func TestOutputScaffoldDomainEntityFieldId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldDomainEntityFieldId(testField)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntityFieldId() failed with error %v`, err)
	}

	if id != testOutputScaffoldDomainEntityStructFieldId {
		t.Errorf(`OutputScaffoldDomainEntityFieldId() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityStructFieldId, id)
	}
}

// Test Output Scaffold Domain Entity Getter ID
func TestOutputScaffoldDomainEntityGetterId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldDomainEntityGetterId(testField)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntityGetterId() failed with error %v`, err)
	}

	if id != testOutputScaffoldDomainEntityGetterId {
		t.Errorf(`OutputScaffoldDomainEntityGetterId() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityGetterId, id)
	}
}

// Test Output Scaffold Domain Entity JSON Tag
func TestOutputScaffoldDomainEntityJSONTagId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldDomainEntityJSONTagId(testField)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntityJSONTagId() failed with error %v`, err)
	}

	if id != testOutputScaffoldDomainEntityJSONTagId {
		t.Errorf(`OutputScaffoldDomainEntityJSONTagId() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityJSONTagId, id)
	}
}

// Test Output Scaffold Domain Entity Setter ID
func TestOutputScaffoldDomainEntitySetterId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldDomainEntitySetterId(testField)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntitySetterId() failed with error %v`, err)
	}

	if id != testOutputScaffoldDomainEntitySetterId {
		t.Errorf(`OutputScaffoldDomainEntitySetterId() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntitySetterId, id)
	}
}

// Test Output Scaffold Domain Entity Set All Setter ID
func TestOutputScaffoldDomainEntitySetAllSetterId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldDomainEntitySetAllSetterId()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntitySetAllSetterId() failed with error %v`, err)
	}

	if id != testOutputScaffoldDomainEntitySetAllSetterId {
		t.Errorf(`OutputScaffoldDomainEntitySetAllSetterId() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntitySetAllSetterId, id)
	}
}

// Test Output Scaffold Domain Entity Len ID
func TestOutputScaffoldDomainEntityLenId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldDomainEntityLenId()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntityLenId() failed with error %v`, err)
	}

	if id != testOutputScaffoldDomainEntityLenId {
		t.Errorf(`OutputScaffoldDomainEntityLenId() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityLenId, id)
	}
}

// Test Output Scaffold Domain Entity Append ID
func TestOutputScaffoldDomainEntityAppendId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldDomainEntityAppendId()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntityAppendId() failed with error %v`, err)
	}

	if id != testOutputScaffoldDomainEntityAppendId {
		t.Errorf(`OutputScaffoldDomainEntityAppendId() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityAppendId, id)
	}
}

// Test Output Scaffold Domain Entity Elements ID
func TestOutputScaffoldDomainEntityElementsId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldDomainEntityElementsId()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntityElementsId() failed with error %v`, err)
	}

	if id != testOutputScaffoldDomainEntityElementsId {
		t.Errorf(`OutputScaffoldDomainEntityElementsId() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityElementsId, id)
	}
}

// Test Output Domain Entity Interface Constructor ID
func TestOutputDomainEntityInterfaceConstructorFunctionId(t *testing.T){

	// Build
	id, err := testFormatter.OutputDomainEntityInterfaceConstructorFunctionId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputDomainEntityInterfaceConstructorFunctionId() failed with error %v`, err)
	}

	if id != testOutputDomainEntityInterfaceConstructorFunctionId {
		t.Errorf(`OutputDomainEntityInterfaceConstructorFunctionId() failed; want "%s", got "%s"`, testOutputDomainEntityInterfaceConstructorFunctionId, id)
	}
}

// Test Output Scaffold Domain Entity Interface Constructor ID
func TestOutputScaffoldDomainEntityInterfaceConstructorFunctionId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldDomainEntityInterfaceConstructorFunctionId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntityInterfaceConstructorFunctionId() failed with error %v`, err)
	}

	if id != testOutputScaffoldDomainEntityInterfaceConstructorFunctionId {
		t.Errorf(`OutputScaffoldDomainEntityInterfaceConstructorFunctionId() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntityInterfaceConstructorFunctionId, id)
	}
}

// Test Output Domain Entity Slice Interface Constructor ID
func TestOutputDomainEntitySliceInterfaceConstructorFunctionId(t *testing.T){

	// Build
	id, err := testFormatter.OutputDomainEntitySliceInterfaceConstructorFunctionId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputDomainEntitySliceInterfaceConstructorFunctionId() failed with error %v`, err)
	}

	if id != testOutputDomainEntitySliceInterfaceConstructorFunctionId {
		t.Errorf(`OutputDomainEntitySliceInterfaceConstructorFunctionId() failed; want "%s", got "%s"`, testOutputDomainEntitySliceInterfaceConstructorFunctionId, id)
	}
}

// Test Output Scaffold Domain Entity Slice Interface Constructor ID
func TestOutputScaffoldDomainEntitySliceInterfaceConstructorFunctionId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldDomainEntitySliceInterfaceConstructorFunctionId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldDomainEntitySliceInterfaceConstructorFunctionId() failed with error %v`, err)
	}

	if id != testOutputScaffoldDomainEntitySliceInterfaceConstructorFunctionId {
		t.Errorf(`OutputScaffoldDomainEntitySliceInterfaceConstructorFunctionId() failed; want "%s", got "%s"`, testOutputScaffoldDomainEntitySliceInterfaceConstructorFunctionId, id)
	}
}

// Test Output Usecase Repository Interface ID
func TestOutputUsecaseRepositoryInterfaceId(t *testing.T){

	// Build
	id, err := testFormatter.OutputUsecaseRepositoryInterfaceId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputUsecaseRepositoryInterfaceId() failed with error %v`, err)
	}

	if id != testOutputUsecaseRepositoryInterfaceId {
		t.Errorf(`OutputUsecaseRepositoryInterfaceId() failed; want "%s", got "%s"`, testOutputUsecaseRepositoryInterfaceId, id)
	}
}

// Test Output Scaffold Usecase Repository Interface ID
func TestOutputScaffoldUsecaseRepositoryInterfaceId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldUsecaseRepositoryInterfaceId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecaseRepositoryInterfaceId() failed with error %v`, err)
	}

	if id != testOutputScaffoldUsecaseRepositoryInterfaceId {
		t.Errorf(`OutputScaffoldUsecaseRepositoryInterfaceId() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseRepositoryInterfaceId, id)
	}
}

// Test Output Scaffold Usecase Repository Interface Method ID
func TestOutputScaffoldUsecaseRepositoryInterfaceMethodId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldUsecaseRepositoryInterfaceMethodId(testMethod)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecaseRepositoryInterfaceMethodId() failed with error %v`, err)
	}

	if id != testOutputScaffoldUsecaseRepositoryInterfaceMethodId {
		t.Errorf(`OutputScaffoldUsecaseRepositoryInterfaceMethodId() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseRepositoryInterfaceMethodId, id)
	}
}

// Test Output Scaffold Usecase Repository Package Name
func TestOutputScaffoldUsecaseRepositoryPackageName(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldUsecaseRepositoryPackageName()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecaseRepositoryPackageName() failed with error %v`, err)
	}

	if id != testOutputScaffoldUsecaseRepositoryPackageName {
		t.Errorf(`OutputScaffoldUsecaseRepositoryPackageName() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseRepositoryPackageName, id)
	}
}

// Test Output Usecase Presenter Interface ID
func TestOutputUsecasePresenterInterfaceId(t *testing.T){

	// Build
	id, err := testFormatter.OutputUsecasePresenterInterfaceId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputUsecasePresenterInterfaceId() failed with error %v`, err)
	}

	if id != testOutputUsecasePresenterInterfaceId {
		t.Errorf(`OutputUsecasePresenterInterfaceId() failed; want "%s", got "%s"`, testOutputUsecasePresenterInterfaceId, id)
	}
}

// Test Output Scaffold Usecase Presenter Interface ID
func TestOutputScaffoldUsecasePresenterInterfaceId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldUsecasePresenterInterfaceId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecasePresenterInterfaceId() failed with error %v`, err)
	}

	if id != testOutputScaffoldUsecasePresenterInterfaceId {
		t.Errorf(`OutputScaffoldUsecasePresenterInterfaceId() failed; want "%s", got "%s"`, testOutputScaffoldUsecasePresenterInterfaceId, id)
	}
}

// Test Output Scaffold Usecase Presenter Interface Method ID
func TestOutputScaffoldUsecasePresenterInterfaceMethodId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldUsecasePresenterInterfaceMethodId(testMethod)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecasePresenterInterfaceMethodId() failed with error %v`, err)
	}

	if id != testOutputScaffoldUsecasePresenterInterfaceMethodId {
		t.Errorf(`OutputScaffoldUsecasePresenterInterfaceMethodId() failed; want "%s", got "%s"`, testOutputScaffoldUsecasePresenterInterfaceMethodId, id)
	}
}

// Test Output Scaffold Usecase Presenter Package Name
func TestOutputScaffoldUsecasePresenterPackageName(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldUsecasePresenterPackageName()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecasePresenterPackageName() failed with error %v`, err)
	}

	if id != testOutputScaffoldUsecasePresenterPackageName {
		t.Errorf(`OutputScaffoldUsecasePresenterPackageName() failed; want "%s", got "%s"`, testOutputScaffoldUsecasePresenterPackageName, id)
	}
}

// Test Output Usecase Interactor Struct ID
func TestOutputUsecaseInteractorStructId(t *testing.T){

	// Build
	id, err := testFormatter.OutputUsecaseInteractorStructId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputUsecaseInteractorStructId() failed with error %v`, err)
	}

	if id != testOutputUsecaseInteractorStructId {
		t.Errorf(`OutputUsecaseInteractorStructId() failed; want "%s", got "%s"`, testOutputUsecaseInteractorStructId, id)
	}
}

// Test Output Scaffold Usecase Interactor Struct ID
func TestOutputScaffoldUsecaseInteractorStructId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldUsecaseInteractorStructId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecaseInteractorStructId() failed with error %v`, err)
	}

	if id != testOutputScaffoldUsecaseInteractorStructId {
		t.Errorf(`OutputScaffoldUsecaseInteractorStructId() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseInteractorStructId, id)
	}
}

// Test Output Usecase Interactor Interface ID
func TestOutputUsecaseInteractorInterfaceId(t *testing.T){

	// Build
	id, err := testFormatter.OutputUsecaseInteractorInterfaceId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputUsecaseInteractorInterfaceId() failed with error %v`, err)
	}

	if id != testOutputUsecaseInteractorInterfaceId {
		t.Errorf(`OutputUsecaseInteractorInterfaceId() failed; want "%s", got "%s"`, testOutputUsecaseInteractorInterfaceId, id)
	}
}

// Test Output Scaffold Usecase Interactor Interface ID
func TestOutputScaffoldUsecaseInteractorInterfaceId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldUsecaseInteractorInterfaceId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecaseInteractorInterfaceId() failed with error %v`, err)
	}

	if id != testOutputScaffoldUsecaseInteractorInterfaceId {
		t.Errorf(`OutputScaffoldUsecaseInteractorInterfaceId() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseInteractorInterfaceId, id)
	}
}

// Test Output Scaffold Usecase Interactor Interface Method ID
func TestOutputScaffoldUsecaseInteractorInterfaceMethodId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldUsecaseInteractorInterfaceMethodId(testMethod)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecaseInteractorInterfaceMethodId() failed with error %v`, err)
	}

	if id != testOutputScaffoldUsecaseInteractorInterfaceMethodId {
		t.Errorf(`OutputScaffoldUsecaseInteractorInterfaceMethodId() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseInteractorInterfaceMethodId, id)
	}
}

// Test Output Scaffold Usecase Interactor Package Name
func TestOutputScaffoldUsecaseInteractorPackageName(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldUsecaseInteractorPackageName()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecaseInteractorPackageName() failed with error %v`, err)
	}

	if id != testOutputScaffoldUsecaseInteractorPackageName {
		t.Errorf(`OutputScaffoldUsecaseInteractorPackageName() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseInteractorPackageName, id)
	}
}


// Test Output Usecase Interactor Constructor Function ID
func TestOutputUsecaseInteractorInterfaceConstructorFunctionId(t *testing.T){

	// Build
	id, err := testFormatter.OutputUsecaseInteractorInterfaceConstructorFunctionId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputUsecaseInteractorInterfaceConstructorFunctionId() failed with error %v`, err)
	}

	if id != testOutputUsecaseInteractorInterfaceConstructorFunctionId {
		t.Errorf(`OutputUsecaseInteractorInterfaceConstructorFunctionId() failed; want "%s", got "%s"`, testOutputUsecaseInteractorInterfaceConstructorFunctionId, id)
	}
}

// Test Output Scaffold Interface Repository Package Name
func TestOutputScaffoldInterfaceRepositoryPackageName(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldInterfaceRepositoryPackageName()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfaceRepositoryPackageName() failed with error %v`, err)
	}

	if id != testOutputScaffoldInterfaceRepositoryPackageName {
		t.Errorf(`OutputScaffoldInterfaceRepositoryPackageName() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceRepositoryPackageName, id)
	}
}

// Test Output Interface Repository Struct ID
func TestOutputInterfaceRepositoryStructId(t *testing.T){

	// Build
	id, err := testFormatter.OutputInterfaceRepositoryStructId("scribble", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputInterfaceRepositoryStructId() failed with error %v`, err)
	}

	if id != testOutputInterfaceRepositoryStructId {
		t.Errorf(`OutputInterfaceRepositoryStructId() failed; want "%s", got "%s"`, testOutputInterfaceRepositoryStructId, id)
	}
}

// Test Output Scaffold Interface Repository Struct ID
func TestOutputScaffoldInterfaceRepositoryStructId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldInterfaceRepositoryStructId("scribble", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfaceRepositoryStructId() failed with error %v`, err)
	}

	if id != testOutputScaffoldInterfaceRepositoryStructId {
		t.Errorf(`OutputScaffoldInterfaceRepositoryStructId() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceRepositoryStructId, id)
	}
}


// Test Output Scaffold Interface Repository Constructor Function ID
func TestOutputScaffoldInterfaceRepositoryConstructorFunctionId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldInterfaceRepositoryConstructorFunctionId("scribble", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfaceRepositoryConstructorFunctionId() failed with error %v`, err)
	}

	if id != testOutputScaffoldInterfaceRepositoryConstructorFunctionId {
		t.Errorf(`OutputScaffoldInterfaceRepositoryConstructorFunctionId() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceRepositoryConstructorFunctionId, id)
	}
}

// Test Output Scaffold Interface Presenter Package Name
func TestOutputScaffoldInterfacePresenterPackageName(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldInterfacePresenterPackageName()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfacePresenterPackageName() failed with error %v`, err)
	}

	if id != testOutputScaffoldInterfacePresenterPackageName {
		t.Errorf(`OutputScaffoldInterfacePresenterPackageName() failed; want "%s", got "%s"`, testOutputScaffoldInterfacePresenterPackageName, id)
	}
}

// Test Output Interface Presenter Struct ID
func TestOutputInterfacePresenterStructId(t *testing.T){

	// Build
	id, err := testFormatter.OutputInterfacePresenterStructId("default", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputInterfacePresenterStructId() failed with error %v`, err)
	}

	if id != testOutputInterfacePresenterStructId {
		t.Errorf(`OutputInterfacePresenterStructId() failed; want "%s", got "%s"`, testOutputInterfacePresenterStructId, id)
	}
}

// Test Output Scaffold Interface Presenter Struct ID
func TestOutputScaffoldInterfacePresenterStructId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldInterfacePresenterStructId("default", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfacePresenterStructId() failed with error %v`, err)
	}

	if id != testOutputScaffoldInterfacePresenterStructId {
		t.Errorf(`OutputScaffoldInterfacePresenterStructId() failed; want "%s", got "%s"`, testOutputScaffoldInterfacePresenterStructId, id)
	}
}

// Test Output Scaffold Interface Presenter Interface ID
func TestOutputScaffoldInterfacePresenterInterfaceId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldInterfacePresenterInterfaceId("default", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfacePresenterInterfaceId() failed with error %v`, err)
	}

	if id != testOutputScaffoldInterfacePresenterInterfaceId {
		t.Errorf(`OutputScaffoldInterfacePresenterInterfaceId() failed; want "%s", got "%s"`, testOutputScaffoldInterfacePresenterInterfaceId, id)
	}
}


// Test Output Scaffold Interface Presenter Constructor Function ID
func TestOutputScaffoldInterfacePresenterConstructorFunctionId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldInterfacePresenterConstructorFunctionId("default", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfacePresenterConstructorFunctionId() failed with error %v`, err)
	}

	if id != testOutputScaffoldInterfacePresenterConstructorFunctionId {
		t.Errorf(`OutputScaffoldInterfacePresenterConstructorFunctionId() failed; want "%s", got "%s"`, testOutputScaffoldInterfacePresenterConstructorFunctionId, id)
	}
}


// Test Output Scaffold Interface Controller Package Name
func TestOutputScaffoldInterfaceControllerPackageName(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldInterfaceControllerPackageName()

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfaceControllerPackageName() failed with error %v`, err)
	}

	if id != testOutputScaffoldInterfaceControllerPackageName {
		t.Errorf(`OutputScaffoldInterfaceControllerPackageName() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceControllerPackageName, id)
	}
}

// Test Output Interface Controller Struct ID
func TestOutputInterfaceControllerStructId(t *testing.T){

	// Build
	id, err := testFormatter.OutputInterfaceControllerStructId("http", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputInterfaceControllerStructId() failed with error %v`, err)
	}

	if id != testOutputInterfaceControllerStructId {
		t.Errorf(`OutputInterfaceControllerStructId() failed; want "%s", got "%s"`, testOutputInterfaceControllerStructId, id)
	}
}

// Test Output Scaffold Interface Controller Struct ID
func TestOutputScaffoldInterfaceControllerStructId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldInterfaceControllerStructId("http", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfaceROutputScaffoldInterfaceControllerStructIdepositoryStructId() failed with error %v`, err)
	}

	if id != testOutputScaffoldInterfaceControllerStructId {
		t.Errorf(`OutputScaffoldInterfaceControllerStructId() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceControllerStructId, id)
	}
}

// Test Output Interface Controller Interface ID
func TestOutputInterfaceControllerInterfaceId(t *testing.T){

	// Build
	id, err := testFormatter.OutputInterfaceControllerInterfaceId("http", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputInterfaceControllerInterfaceId() failed with error %v`, err)
	}

	if id != testOutputInterfaceControllerInterfaceId {
		t.Errorf(`OutputInterfaceControllerInterfaceId() failed; want "%s", got "%s"`, testOutputInterfaceControllerInterfaceId, id)
	}
}

// Test Output Scaffold Interface Controller Interface ID
func TestOutputScaffoldInterfaceControllerInterfaceId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldInterfaceControllerInterfaceId("http", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfaceControllerInterfaceId() failed with error %v`, err)
	}

	if id != testOutputScaffoldInterfaceControllerInterfaceId {
		t.Errorf(`OutputScaffoldInterfaceControllerInterfaceId() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceControllerInterfaceId, id)
	}
}

// Test Output Scaffold Interface Controller Constructor Function ID
func TestOutputScaffoldInterfaceControllerConstructorFunctionId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldInterfaceControllerConstructorFunctionId("http", testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldInterfaceControllerConstructorFunctionId() failed with error %v`, err)
	}

	if id != testOutputScaffoldInterfaceControllerConstructorFunctionId {
		t.Errorf(`OutputScaffoldInterfaceControllerConstructorFunctionId() failed; want "%s", got "%s"`, testOutputScaffoldInterfaceControllerConstructorFunctionId, id)
	}
}