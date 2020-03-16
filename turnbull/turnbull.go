package turnbull

import(
	"os"
	"bytes"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/structure"
	"github.com/brianshepanek/turnbull/generator"
)

type turnbull struct{
	formatter formatter.Formatter
	structure structure.Structure
	generator generator.Generator
}

func New(formatter formatter.Formatter, structure structure.Structure, generator generator.Generator) *turnbull {
	return &turnbull{
		formatter : formatter,
		structure : structure,
		generator : generator,
	}
}

// Build
func (turnbull *turnbull) buildStructure() (error){

	var err error
	err = turnbull.structure.Build()
	if err != nil {
		return err
	}

	return nil
}

// Build Scaffold Domain Entity
func (turnbull *turnbull) buildScaffoldDomainEntity(entity model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.ScaffoldEntity(entity, buf)
	if err != nil {
		return err
	}

	// File Name
	fileName, err := turnbull.formatter.OutputScaffoldDomainEntityFile(entity)
	if err != nil {
		return err
	}

	// File
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write
	_, err = file.WriteString(buf.String())
	if err != nil {
		return err
	}

	return nil
}

// Build Scaffold Usecase Repository
func (turnbull *turnbull) buildScaffoldUsecaseRepository(entity model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.ScaffoldUsecaseRepository(entity, buf)
	if err != nil {
		return err
	}

	// File Name
	fileName, err := turnbull.formatter.OutputScaffoldUsecaseRepositoryFile(entity)
	if err != nil {
		return err
	}

	// File
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write
	_, err = file.WriteString(buf.String())
	if err != nil {
		return err
	}

	return nil
}

// Build Scaffold Usecase Presenter
func (turnbull *turnbull) buildScaffoldUsecasePresenter(entity model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.ScaffoldUsecasePresenter(entity, buf)
	if err != nil {
		return err
	}

	// File Name
	fileName, err := turnbull.formatter.OutputScaffoldUsecasePresenterFile(entity)
	if err != nil {
		return err
	}

	// File
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write
	_, err = file.WriteString(buf.String())
	if err != nil {
		return err
	}

	return nil
}

// Build Scaffold Usecase Interactor
func (turnbull *turnbull) buildScaffoldUsecaseInteractor(entity model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.ScaffoldUsecaseInteractor(entity, buf)
	if err != nil {
		return err
	}

	// File Name
	fileName, err := turnbull.formatter.OutputScaffoldUsecaseInteractorFile(entity)
	if err != nil {
		return err
	}

	// File
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write
	_, err = file.WriteString(buf.String())
	if err != nil {
		return err
	}

	return nil
}