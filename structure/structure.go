package structure

import(
	"os"
	"github.com/brianshepanek/turnbull/formatter"
)

type structure struct{
	formatter formatter.Formatter
}

type Structure interface{
	Build() (error)
// 	buildOutputDirectory() (error)
// 	buildOutputScaffoldDirectory() (error)

// 	buildOutputScaffoldDomainDirectory() (error)
// 	buildOutputScaffoldDomainEntityDirectory() (error)

// 	buildOutputScaffoldUsecaseDirectory() (error)
// 	buildOutputScaffoldUsecaseInteractorDirectory() (error)
// 	buildOutputScaffoldUsecaseRepositoryDirectory() (error)
// 	buildOutputScaffoldUsecasePresenterDirectory() (error)

// 	buildOutputScaffoldInterfaceDirectory() (error)
// 	buildOutputScaffoldInterfaceControllerDirectory() (error)
// 	buildOutputScaffoldInterfaceRepositoryDirectory() (error)
// 	buildOutputScaffoldInterfacePresenterDirectory() (error)

// 	makeDir(path string) (error)
}

func New(formatter formatter.Formatter) *structure {
	return &structure{
		formatter : formatter,
	}
}

// Build
func (structure *structure) Build() (error){

	var err error
	err = structure.buildOutputDirectory()
	if err != nil {
		return err
	}
	err = structure.buildRegistryDirectory()
	if err != nil {
		return err
	}
	err = structure.buildOutputScaffoldDomainDirectory()
	if err != nil {
		return err
	}
	err = structure.buildOutputScaffoldDirectory()
	if err != nil {
		return err
	}
	err = structure.buildOutputScaffoldDomainEntityDirectory()
	if err != nil {
		return err
	}
	err = structure.buildOutputScaffoldUsecaseDirectory()
	if err != nil {
		return err
	}
	err = structure.buildOutputScaffoldUsecaseInteractorDirectory()
	if err != nil {
		return err
	}
	err = structure.buildOutputScaffoldUsecaseRepositoryDirectory()
	if err != nil {
		return err
	}
	err = structure.buildOutputScaffoldUsecasePresenterDirectory()
	if err != nil {
		return err
	}
	err = structure.buildOutputScaffoldInterfaceDirectory()
	if err != nil {
		return err
	}
	err = structure.buildOutputScaffoldInterfaceControllerDirectory()
	if err != nil {
		return err
	}
	err = structure.buildOutputScaffoldInterfaceRepositoryDirectory()
	if err != nil {
		return err
	}
	err = structure.buildOutputScaffoldInterfacePresenterDirectory()
	if err != nil {
		return err
	}

	return nil
}

// Output

func (structure *structure) buildOutputDirectory() (error){
	dir, err := structure.formatter.OutputDirectory()
	if err != nil {
		return err
	}
	return structure.makeDir(dir)
}

// Scaffold

func (structure *structure) buildOutputScaffoldDirectory() (error){
	dir, err := structure.formatter.OutputScaffoldDirectory()
	if err != nil {
		return err
	}
	return structure.makeDir(dir)
}

// Domain

func (structure *structure) buildOutputScaffoldDomainDirectory() (error){
	dir, err := structure.formatter.OutputScaffoldDomainDirectory()
	if err != nil {
		return err
	}
	return structure.makeDir(dir)
}

func (structure *structure) buildOutputScaffoldDomainEntityDirectory() (error){
	dir, err := structure.formatter.OutputScaffoldDomainEntityDirectory()
	if err != nil {
		return err
	}
	return structure.makeDir(dir)
}


// Usecase

func (structure *structure) buildOutputScaffoldUsecaseDirectory() (error){
	dir, err := structure.formatter.OutputScaffoldUsecaseDirectory()
	if err != nil {
		return err
	}
	return structure.makeDir(dir)
}

func (structure *structure) buildOutputScaffoldUsecaseInteractorDirectory() (error){
	dir, err := structure.formatter.OutputScaffoldUsecaseInteractorDirectory()
	if err != nil {
		return err
	}
	return structure.makeDir(dir)
}

func (structure *structure) buildOutputScaffoldUsecaseRepositoryDirectory() (error){
	dir, err := structure.formatter.OutputScaffoldUsecaseRepositoryDirectory()
	if err != nil {
		return err
	}
	return structure.makeDir(dir)
}

func (structure *structure) buildOutputScaffoldUsecasePresenterDirectory() (error){
	dir, err := structure.formatter.OutputScaffoldUsecasePresenterDirectory()
	if err != nil {
		return err
	}
	return structure.makeDir(dir)
}

// Interface

func (structure *structure) buildOutputScaffoldInterfaceDirectory() (error){
	dir, err := structure.formatter.OutputScaffoldInterfaceDirectory()
	if err != nil {
		return err
	}
	return structure.makeDir(dir)
}

func (structure *structure) buildOutputScaffoldInterfaceControllerDirectory() (error){
	dir, err := structure.formatter.OutputScaffoldInterfaceControllerDirectory()
	if err != nil {
		return err
	}
	return structure.makeDir(dir)
}

func (structure *structure) buildOutputScaffoldInterfaceRepositoryDirectory() (error){
	dir, err := structure.formatter.OutputScaffoldInterfaceRepositoryDirectory()
	if err != nil {
		return err
	}
	return structure.makeDir(dir)
}

func (structure *structure) buildOutputScaffoldInterfacePresenterDirectory() (error){
	dir, err := structure.formatter.OutputScaffoldInterfacePresenterDirectory()
	if err != nil {
		return err
	}
	return structure.makeDir(dir)
}

func (structure *structure) buildRegistryDirectory() (error){
	dir, err := structure.formatter.OutputRegistryDirectory()
	if err != nil {
		return err
	}
	return structure.makeDir(dir)
}

// Make Dir

func (structure *structure) makeDir(path string) (error){
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}