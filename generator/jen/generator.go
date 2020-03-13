package generator

import(
	"io"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/generator/jen/helper"
	"github.com/brianshepanek/turnbull/generator/jen/domain"
	"github.com/brianshepanek/turnbull/generator/jen/usecase"
)

type Generator struct{
	config *config.Config
	formatter formatter.Formatter
	entityGenerator domain.EntityGenerator
	usecaseInteractorGenerator usecase.InteractorGenerator
	usecasePresenterGenerator usecase.PresenterGenerator
	usecaseRepositoryGenerator usecase.RepositoryGenerator
}

func NewGenerator(config *config.Config, formatter formatter.Formatter) * Generator{


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
	}
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


