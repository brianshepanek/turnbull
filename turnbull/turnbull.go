package turnbull

import(
	"os"
	"path/filepath"
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

// Build Domain Entity
func (turnbull *turnbull) buildDomainEntity(entity model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.Entity(entity, buf)
	if err != nil {
		return err
	}

	// File Name
	fileName, err := turnbull.formatter.OutputDomainEntityFile(entity)
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

// Build Usecase Repository
func (turnbull *turnbull) buildUsecaseRepository(entity model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.UsecaseRepository(entity, buf)
	if err != nil {
		return err
	}

	// File Name
	fileName, err := turnbull.formatter.OutputUsecaseRepositoryFile(entity)
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

// Build Usecase Presenter
func (turnbull *turnbull) buildUsecasePresenter(entity model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.UsecasePresenter(entity, buf)
	if err != nil {
		return err
	}

	// File Name
	fileName, err := turnbull.formatter.OutputUsecasePresenterFile(entity)
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

// Build Usecase Interactor
func (turnbull *turnbull) buildUsecaseInteractor(entity model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.UsecaseInteractor(entity, buf)
	if err != nil {
		return err
	}

	// File Name
	fileName, err := turnbull.formatter.OutputUsecaseInteractorFile(entity)
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

// Build Interface Repository
func (turnbull *turnbull) buildInterfaceRepository(driver string, entity model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.InterfaceRepository(driver, entity, buf)
	if err != nil {
		return err
	}

	// File Name
	fileName, err := turnbull.formatter.OutputInterfaceRepositoryFile(driver, entity)
	if err != nil {
		return err
	}

	// Ensure
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			return err
		}
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

// Build Scaffold Interface Repository
func (turnbull *turnbull) buildScaffoldInterfaceRepository(driver string, entity model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.ScaffoldInterfaceRepository(driver, entity, buf)
	if err != nil {
		return err
	}

	// File Name
	fileName, err := turnbull.formatter.OutputScaffoldInterfaceRepositoryFile(driver, entity)
	if err != nil {
		return err
	}

	// Ensure
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			return err
		}
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

// Build Entity Interface Repository
func (turnbull *turnbull) buildEntityInterfaceRepository(driver string, entity model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.InterfaceRepositoryEntity(driver, entity, buf)
	if err != nil {
		return err
	}

	// Exists
	if len(buf.String()) > 0 {

		// File Name
		fileName, err := turnbull.formatter.OutputInterfaceRepositoryEntityFile(driver, entity)
		if err != nil {
			return err
		}

		// Ensure
		dirName := filepath.Dir(fileName)
		if _, serr := os.Stat(dirName); serr != nil {
			merr := os.MkdirAll(dirName, os.ModePerm)
			if merr != nil {
				return err
			}
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
		
	}

	

	return nil
}

// Build Interface Presenter
func (turnbull *turnbull) buildInterfacePresenter(driver string, entity model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.InterfacePresenter(driver, entity, buf)
	if err != nil {
		return err
	}

	// File Name
	fileName, err := turnbull.formatter.OutputInterfacePresenterFile(driver, entity)
	if err != nil {
		return err
	}

	// Ensure
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			return err
		}
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

// Build Scaffold Interface Presenter
func (turnbull *turnbull) buildScaffoldInterfacePresenter(driver string, entity model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.ScaffoldInterfacePresenter(driver, entity, buf)
	if err != nil {
		return err
	}

	// File Name
	fileName, err := turnbull.formatter.OutputScaffoldInterfacePresenterFile(driver, entity)
	if err != nil {
		return err
	}

	// Ensure
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			return err
		}
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

// Build Interface App
func (turnbull *turnbull) buildInterfaceAppController(driver string, entities []model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.InterfaceAppController(driver, entities, buf)
	if err != nil {
		return err
	}

	// File Name

	fileName, err := turnbull.formatter.OutputInterfaceControllerFile(driver, model.Entity{Name : "app"})
	if err != nil {
		return err
	}

	// Ensure
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			return err
		}
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

// Build Interface Controller
func (turnbull *turnbull) buildInterfaceController(driver string, entity model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.InterfaceController(driver,entity, buf)
	if err != nil {
		return err
	}

	// File Name
	fileName, err := turnbull.formatter.OutputInterfaceControllerFile(driver, entity)
	if err != nil {
		return err
	}

	// Ensure
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			return err
		}
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

// Build Scaffold Interface Controller
func (turnbull *turnbull) buildScaffoldInterfaceController(driver string, entity model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.ScaffoldInterfaceController(driver, entity, buf)
	if err != nil {
		return err
	}

	// File Name
	fileName, err := turnbull.formatter.OutputScaffoldInterfaceControllerFile(driver, entity)
	if err != nil {
		return err
	}

	// Ensure
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			return err
		}
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

// Build Interface Controller Entitty
func (turnbull *turnbull) buildInterfaceControllerEntity(driver string, entity model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.InterfaceControllerEntity(driver, entity, buf)
	if err != nil {
		return err
	}

	// File Name
	fileName, err := turnbull.formatter.OutputInterfaceControllerEntityFile(driver, entity)
	if err != nil {
		return err
	}

	// Ensure
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			return err
		}
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