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
	testOutputScaffoldDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/scaffold"

	testOutputScaffoldDomainDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/scaffold/domain"
	testOutputScaffoldDomainEntityDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/scaffold/domain/entity"
	testOutputScaffoldDomainEntityDirectoryImportPath = "github.com/brianshepanek/turnbull/_testing/output/scaffold/domain/entity"
	testOutputScaffoldDomainEntityFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/scaffold/domain/entity/foo.go"

	testOutputScaffoldUsecaseDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/scaffold/usecase"
	testOutputScaffoldUsecaseInteractorDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/scaffold/usecase/interactor"
	testOutputScaffoldUsecaseInteractorDirectoryImportPath = "github.com/brianshepanek/turnbull/_testing/output/scaffold/usecase/interactor"
	testOutputScaffoldUsecaseInteractorFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/scaffold/usecase/interactor/foo_interactor.go"
	testOutputScaffoldUsecaseRepositoryDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/scaffold/usecase/repository"
	testOutputScaffoldUsecaseRepositoryDirectoryImportPath = "github.com/brianshepanek/turnbull/_testing/output/scaffold/usecase/repository"
	testOutputScaffoldUsecaseRepositoryFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/scaffold/usecase/repository/foo_repository.go"
	testOutputScaffoldUsecasePresenterDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/scaffold/usecase/presenter"
	testOutputScaffoldUsecasePresenterDirectoryImportPath = "github.com/brianshepanek/turnbull/_testing/output/scaffold/usecase/presenter"
	testOutputScaffoldUsecasePresenterFile = "/go/src/github.com/brianshepanek/turnbull/_testing/output/scaffold/usecase/presenter/foo_presenter.go"

	testOutputScaffoldInterfaceDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/scaffold/interface"
	testOutputScaffoldInterfaceControllerDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/scaffold/interface/controller"
	testOutputScaffoldInterfaceRepositoryDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/scaffold/interface/repository"
	testOutputScaffoldInterfacePresenterDirectory = "/go/src/github.com/brianshepanek/turnbull/_testing/output/scaffold/interface/presenter"

	testOutputScaffoldDomainEntityPackageName = "entity"

	testOutputScaffoldDomainEntityStructId = "FooScaffoldStruct"
	testOutputScaffoldDomainEntitySliceStructId = "FoosScaffoldStruct"

	testOutputScaffoldDomainEntityInterfaceId = "FooScaffoldInterface"
	testOutputScaffoldDomainEntitySliceInterfaceId = "FoosScaffoldInterface"

	testOutputScaffoldDomainEntityStructFieldId = "bar"
	testOutputScaffoldDomainEntityGetterId = "Bar"
	testOutputScaffoldDomainEntityJSONTagId = "bar"
	testOutputScaffoldDomainEntitySetterId = "SetBar"
	testOutputScaffoldDomainEntitySetAllSetterId = "SetAll"

	testOutputScaffoldDomainEntityLenId = "Len"
	testOutputScaffoldDomainEntityAppendId = "Append"
	testOutputScaffoldDomainEntityElementsId = "Elements"

	testOutputScaffoldDomainEntityInterfaceConstructorFunctionId = "NewFooScaffoldStruct"
	testOutputScaffoldDomainEntitySliceInterfaceConstructorFunctionId = "NewFoosScaffoldStruct"

	testOutputScaffoldUsecaseRepositoryPackageName = "repository"
	testOutputScaffoldUsecaseRepositoryInterfaceId = "FooScaffoldRepository"
	testOutputScaffoldUsecaseRepositoryInterfaceMethodId = "Add"

	testOutputScaffoldUsecasePresenterPackageName = "presenter"
	testOutputScaffoldUsecasePresenterInterfaceId = "FooScaffoldPresenter"
	testOutputScaffoldUsecasePresenterInterfaceMethodId = "Add"

	testOutputScaffoldUsecaseInteractorPackageName = "interactor"
	testOutputScaffoldUsecaseInteractorStructId = "fooScaffoldInteractor"
	testOutputScaffoldUsecaseInteractorInterfaceId = "FooScaffoldInteractor"
	testOutputScaffoldUsecaseInteractorInterfaceMethodId = "Add"
	testOutputScaffoldUsecaseInteractorInterfaceConstructorFunctionId = "NewFooScaffoldInteractor"

	testOutputScaffoldInterfaceRepositoryPackageName = "repository"
	testOutputScaffoldInterfaceRepositoryStructId = "scribbleFooScaffoldRepository"
	testOutputScaffoldInterfaceRepositoryConstructorFunctionId = "NewScribbleFooScaffoldRepository"

	testOutputScaffoldInterfacePresenterPackageName = "presenter"
	testOutputScaffoldInterfacePresenterStructId = "defaultFooScaffoldPresenter"
	testOutputScaffoldInterfacePresenterInterfaceId = "DefaultFooScaffoldPresenter"
	testOutputScaffoldInterfacePresenterConstructorFunctionId = "NewDefaultFooScaffoldPresenter"

	testOutputScaffoldInterfaceControllerPackageName = "controller"
	testOutputScaffoldInterfaceControllerStructId = "httpFooScaffoldController"
	testOutputScaffoldInterfaceControllerConstructorFunctionId = "NewHttpFooScaffoldController"

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


// Test Output Scaffold Usecase Interactor Constructor Function ID
func TestOutputScaffoldUsecaseInteractorInterfaceConstructorFunctionId(t *testing.T){

	// Build
	id, err := testFormatter.OutputScaffoldUsecaseInteractorInterfaceConstructorFunctionId(testEntity)

	// Return
	if err != nil {
		t.Errorf(`OutputScaffoldUsecaseInteractorInterfaceConstructorFunctionId() failed with error %v`, err)
	}

	if id != testOutputScaffoldUsecaseInteractorInterfaceConstructorFunctionId {
		t.Errorf(`OutputScaffoldUsecaseInteractorInterfaceConstructorFunctionId() failed; want "%s", got "%s"`, testOutputScaffoldUsecaseInteractorInterfaceConstructorFunctionId, id)
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