package structure

import(
	"testing"
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/formatter"
)

const(
	testConfigPath = "./config"
	testOutputPath = "./output"
)

var (
	testStructure Structure
)

func init(){
	conf, _ := config.New(testConfigPath, testOutputPath)
	formatter := formatter.New(conf)
	testStructure = New(formatter)
	
}


// Test Build Output Directory
func TestBuildOutputDirectory(t *testing.T){

	// Build
	err := testStructure.buildOutputDirectory()

	// Return
	if err != nil {
		t.Errorf(`buildOutputDirectory() failed with error %v`, err)
	}
}

// Test Build Output Scaffold Directory
func TestBuildOutputScaffoldDirectory(t *testing.T){

	// Build
	err := testStructure.buildOutputScaffoldDirectory()

	// Return
	if err != nil {
		t.Errorf(`buildOutputScaffoldDirectory() failed with error %v`, err)
	}
}

// Test Build Output Scaffold Domain Directory
func TestBuildOutputScaffoldDomainDirectory(t *testing.T){

	// Build
	err := testStructure.buildOutputScaffoldDomainDirectory()

	// Return
	if err != nil {
		t.Errorf(`buildOutputScaffoldDomainDirectory() failed with error %v`, err)
	}
}

// Test Build Output Scaffold Domain Entity Directory
func TestBuildOutputScaffoldDomainEntityDirectory(t *testing.T){

	// Build
	err := testStructure.buildOutputScaffoldDomainEntityDirectory()

	// Return
	if err != nil {
		t.Errorf(`buildOutputScaffoldDomainEntityDirectory() failed with error %v`, err)
	}
}

// Test Build Output Scaffold Usecase Directory
func TestBuildOutputScaffoldUsecaseDirectory(t *testing.T){

	// Build
	err := testStructure.buildOutputScaffoldUsecaseDirectory()

	// Return
	if err != nil {
		t.Errorf(`buildOutputScaffoldUsecaseDirectory() failed with error %v`, err)
	}
}

// Test Build Output Scaffold Usecase Interactor Directory
func TestBuildOutputScaffoldUsecaseInteractorDirectory(t *testing.T){

	// Build
	err := testStructure.buildOutputScaffoldUsecaseInteractorDirectory()

	// Return
	if err != nil {
		t.Errorf(`buildOutputScaffoldUsecaseInteractorDirectory() failed with error %v`, err)
	}
}

// Test Build Output Scaffold Usecase Repository Directory
func TestBuildOutputScaffoldUsecaseRepositoryDirectory(t *testing.T){

	// Build
	err := testStructure.buildOutputScaffoldUsecaseRepositoryDirectory()

	// Return
	if err != nil {
		t.Errorf(`buildOutputScaffoldUsecaseRepositoryDirectory() failed with error %v`, err)
	}
}

// Test Build Output Scaffold Usecase Presenter Directory
func TestBuildOutputScaffoldUsecasePresenterDirectory(t *testing.T){

	// Build
	err := testStructure.buildOutputScaffoldUsecasePresenterDirectory()

	// Return
	if err != nil {
		t.Errorf(`buildOutputScaffoldUsecasePresenterDirectory() failed with error %v`, err)
	}
}

// Test Build Output Scaffold Interface Directory
func TestBuildOutputScaffoldInterfaceDirectory(t *testing.T){

	// Build
	err := testStructure.buildOutputScaffoldInterfaceDirectory()

	// Return
	if err != nil {
		t.Errorf(`buildOutputScaffoldInterfaceDirectory() failed with error %v`, err)
	}
}

// Test Build Output Scaffold Interface Controller Directory
func TestBuildOutputScaffoldInterfaceControllerDirectory(t *testing.T){

	// Build
	err := testStructure.buildOutputScaffoldInterfaceControllerDirectory()

	// Return
	if err != nil {
		t.Errorf(`buildOutputScaffoldInterfaceControllerDirectory() failed with error %v`, err)
	}
}

// Test Build Output Scaffold Interface Repository Directory
func TestBuildOutputScaffoldInterfaceRepositoryDirectory(t *testing.T){

	// Build
	err := testStructure.buildOutputScaffoldInterfaceRepositoryDirectory()

	// Return
	if err != nil {
		t.Errorf(`buildOutputScaffoldInterfaceRepositoryDirectory() failed with error %v`, err)
	}
}

// Test Build Output Scaffold Interface Presenter Directory
func TestBuildOutputScaffoldInterfacePresenterDirectory(t *testing.T){

	// Build
	err := testStructure.buildOutputScaffoldInterfacePresenterDirectory()

	// Return
	if err != nil {
		t.Errorf(`buildOutputScaffoldInterfacePresenterDirectory() failed with error %v`, err)
	}
}