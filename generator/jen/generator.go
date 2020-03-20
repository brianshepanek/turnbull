package generator

import(
	"io"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/generator/jen/helper"
	"github.com/brianshepanek/turnbull/generator/jen/domain"
	"github.com/brianshepanek/turnbull/generator/jen/usecase"
	generatorInterface "github.com/brianshepanek/turnbull/generator/jen/interface"
)

type Generator struct{
	config *config.Config
	formatter formatter.Formatter
	entityGenerator domain.EntityGenerator
	usecaseInteractorGenerator usecase.InteractorGenerator
	usecasePresenterGenerator usecase.PresenterGenerator
	usecaseRepositoryGenerator usecase.RepositoryGenerator
	interfaceControllerGenerator generatorInterface.ControllerGenerator
	interfacePresenterGenerator generatorInterface.PresenterGenerator
	interfaceRepositoryGenerator generatorInterface.RepositoryGenerator
}

func New(config *config.Config, formatter formatter.Formatter, interfaceControllerGenerator generatorInterface.ControllerGenerator, interfacePresenterGenerator generatorInterface.PresenterGenerator, interfaceRepositoryGenerator generatorInterface.RepositoryGenerator) * Generator{


	helperGenerator := helper.New(formatter)
	entityGenerator := domain.NewEntityGenerator(formatter, helperGenerator)
	usecaseInteractorGenerator := usecase.NewInteractorGenerator(config, formatter, helperGenerator)
	usecasePresenterGenerator := usecase.NewPresenterGenerator(config, formatter, helperGenerator)
	usecaseRepositoryGenerator := usecase.NewRepositoryGenerator(config, formatter, helperGenerator)

	return &Generator{
		config : config,
		formatter : formatter,
		entityGenerator : entityGenerator,
		usecaseInteractorGenerator : usecaseInteractorGenerator,
		usecasePresenterGenerator : usecasePresenterGenerator,
		usecaseRepositoryGenerator : usecaseRepositoryGenerator,
		interfaceControllerGenerator : interfaceControllerGenerator,
		interfacePresenterGenerator : interfacePresenterGenerator,
		interfaceRepositoryGenerator : interfaceRepositoryGenerator,
	}
}

func (generator *Generator) Entity(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.entityGenerator.File(entity)
	if err != nil {
		return err
	}

	// Render
	file.Render(writer)

	return nil
}

func (generator *Generator) ScaffoldEntity(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.entityGenerator.ScaffoldFile(entity)
	if err != nil {
		return err
	}

	// Render
	file.Render(writer)

	return nil
}

func (generator *Generator) UsecaseRepository(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.usecaseRepositoryGenerator.File(entity)
	if err != nil {
		return err
	}

	// Render
	file.Render(writer)

	return nil
}

func (generator *Generator) ScaffoldUsecaseRepository(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.usecaseRepositoryGenerator.ScaffoldFile(entity)
	if err != nil {
		return err
	}

	// Render
	file.Render(writer)

	return nil
}

func (generator *Generator) UsecasePresenter(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.usecasePresenterGenerator.File(entity)
	if err != nil {
		return err
	}

	// Render
	file.Render(writer)

	return nil
}

func (generator *Generator) ScaffoldUsecasePresenter(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.usecasePresenterGenerator.ScaffoldFile(entity)
	if err != nil {
		return err
	}

	// Render
	file.Render(writer)

	return nil
}

func (generator *Generator) UsecaseInteractor(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.usecaseInteractorGenerator.File(entity)
	if err != nil {
		return err
	}

	// Render
	file.Render(writer)

	return nil
}

func (generator *Generator) ScaffoldUsecaseInteractor(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.usecaseInteractorGenerator.ScaffoldFile(entity)
	if err != nil {
		return err
	}

	// Render
	file.Render(writer)

	return nil
}

func (generator *Generator) InterfaceRepository(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.interfaceRepositoryGenerator.File(entity)
	if err != nil {
		return err
	}

	// Render
	file.Render(writer)

	return nil
}

func (generator *Generator) ScaffoldInterfaceRepository(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.interfaceRepositoryGenerator.ScaffoldFile(entity)
	if err != nil {
		return err
	}

	// Render
	file.Render(writer)

	return nil
}

func (generator *Generator) ScaffoldInterfacePresenter(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.interfacePresenterGenerator.ScaffoldFile(entity)
	if err != nil {
		return err
	}

	// Render
	file.Render(writer)

	return nil
}

func (generator *Generator) ScaffoldInterfaceController(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.interfaceControllerGenerator.ScaffoldFile(entity)
	if err != nil {
		return err
	}

	// Render
	file.Render(writer)

	return nil
}