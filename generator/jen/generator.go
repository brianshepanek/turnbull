package generator

import(
	"io"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/generator/jen/helper"
	"github.com/brianshepanek/turnbull/generator/jen/domain"
	"github.com/brianshepanek/turnbull/generator/jen/usecase"
	"github.com/brianshepanek/turnbull/generator/jen/registry"
	generatorInterface "github.com/brianshepanek/turnbull/generator/jen/interface"
)

type Generator struct{
	config *config.Config
	formatter formatter.Formatter
	entityGenerator domain.EntityGenerator
	usecaseInteractorGenerator usecase.InteractorGenerator
	usecasePresenterGenerator usecase.PresenterGenerator
	usecaseRepositoryGenerator usecase.RepositoryGenerator
	interfaceControllerGenerators map[string]generatorInterface.ControllerGenerator
	interfacePresenterGenerators map[string]generatorInterface.PresenterGenerator
	interfaceRepositoryGenerators map[string]generatorInterface.RepositoryGenerator
	registryGenerator registry.RegistryGenerator
}

func New(config *config.Config, formatter formatter.Formatter, interfaceControllerGenerators map[string]generatorInterface.ControllerGenerator,interfacePresenterGenerators map[string]generatorInterface.PresenterGenerator, interfaceRepositoryGenerators map[string]generatorInterface.RepositoryGenerator) * Generator{


	helperGenerator := helper.New(formatter)
	entityGenerator := domain.NewEntityGenerator(formatter, helperGenerator)
	usecaseInteractorGenerator := usecase.NewInteractorGenerator(config, formatter, helperGenerator)
	usecasePresenterGenerator := usecase.NewPresenterGenerator(config, formatter, helperGenerator)
	usecaseRepositoryGenerator := usecase.NewRepositoryGenerator(config, formatter, helperGenerator)
	registryGenerator := registry.New(formatter, helperGenerator)

	return &Generator{
		config : config,
		formatter : formatter,
		entityGenerator : entityGenerator,
		usecaseInteractorGenerator : usecaseInteractorGenerator,
		usecasePresenterGenerator : usecasePresenterGenerator,
		usecaseRepositoryGenerator : usecaseRepositoryGenerator,
		interfaceControllerGenerators : interfaceControllerGenerators,
		interfacePresenterGenerators : interfacePresenterGenerators,
		interfaceRepositoryGenerators : interfaceRepositoryGenerators,
		registryGenerator :registryGenerator,
	}
}

func (generator *Generator) Entity(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.entityGenerator.File(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) ScaffoldEntity(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.entityGenerator.ScaffoldFile(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) UsecaseRepository(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.usecaseRepositoryGenerator.File(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) ScaffoldUsecaseRepository(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.usecaseRepositoryGenerator.ScaffoldFile(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) UsecasePresenter(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.usecasePresenterGenerator.File(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) ScaffoldUsecasePresenter(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.usecasePresenterGenerator.ScaffoldFile(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) UsecaseInteractor(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.usecaseInteractorGenerator.File(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) ScaffoldUsecaseInteractor(entity model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.usecaseInteractorGenerator.ScaffoldFile(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) InterfaceRepository(driver string, entity model.Entity, writer io.Writer) (error){

	// Vars
	var interfaceRepositoryGenerator generatorInterface.RepositoryGenerator
	if val, ok := generator.interfaceRepositoryGenerators[driver]; ok {
		interfaceRepositoryGenerator = val
	}

	// File
	file, err := interfaceRepositoryGenerator.File(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) ScaffoldInterfaceRepository(driver string, entity model.Entity, writer io.Writer) (error){
	
	// Vars
	var interfaceRepositoryGenerator generatorInterface.RepositoryGenerator
	if val, ok := generator.interfaceRepositoryGenerators[driver]; ok {
		interfaceRepositoryGenerator = val
	}

	// File
	file, err := interfaceRepositoryGenerator.ScaffoldFile(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) InterfaceRepositoryEntity(driver string, entity model.Entity, writer io.Writer) (error){

	// Vars
	var interfaceRepositoryGenerator generatorInterface.RepositoryGenerator
	if val, ok := generator.interfaceRepositoryGenerators[driver]; ok {
		interfaceRepositoryGenerator = val
	}

	// File
	file, err := interfaceRepositoryGenerator.EntityFile(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) InterfaceRepositoryRegistry(driver string, entity model.Entity, writer io.Writer) (error){

	// Vars
	var interfaceRepositoryGenerator generatorInterface.RepositoryGenerator
	if val, ok := generator.interfaceRepositoryGenerators[driver]; ok {
		interfaceRepositoryGenerator = val
	}

	// File
	file, err := interfaceRepositoryGenerator.RegistryFile(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) InterfacePresenter(driver string, entity model.Entity, writer io.Writer) (error){

	// Vars
	var interfacePresenterGenerator generatorInterface.PresenterGenerator
	if val, ok := generator.interfacePresenterGenerators[driver]; ok {
		interfacePresenterGenerator = val
	}

	// File
	file, err := interfacePresenterGenerator.File(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) ScaffoldInterfacePresenter(driver string, entity model.Entity, writer io.Writer) (error){

	// Vars
	var interfacePresenterGenerator generatorInterface.PresenterGenerator
	if val, ok := generator.interfacePresenterGenerators[driver]; ok {
		interfacePresenterGenerator = val
	}

	// File
	file, err := interfacePresenterGenerator.ScaffoldFile(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) InterfacePresenterEntity(driver string, entity model.Entity, writer io.Writer) (error){

	// Vars
	var interfacePresenterGenerator generatorInterface.PresenterGenerator
	if val, ok := generator.interfacePresenterGenerators[driver]; ok {
		interfacePresenterGenerator = val
	}

	// File
	file, err := interfacePresenterGenerator.EntityFile(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) InterfacePresenterRegistry(driver string, entity model.Entity, writer io.Writer) (error){

	// Vars
	var interfacePresenterGenerator generatorInterface.PresenterGenerator
	if val, ok := generator.interfacePresenterGenerators[driver]; ok {
		interfacePresenterGenerator = val
	}

	// File
	file, err := interfacePresenterGenerator.RegistryFile(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) InterfaceAppController(driver string, entities []model.Entity, writer io.Writer) (error){

	// Vars
	var interfaceControllerGenerator generatorInterface.ControllerGenerator
	if val, ok := generator.interfaceControllerGenerators[driver]; ok {
		interfaceControllerGenerator = val
	}

	// File
	file, err := interfaceControllerGenerator.AppFile(entities)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) InterfaceController(driver string, entity model.Entity, writer io.Writer) (error){

	// Vars
	var interfaceControllerGenerator generatorInterface.ControllerGenerator
	if val, ok := generator.interfaceControllerGenerators[driver]; ok {
		interfaceControllerGenerator = val
	}

	// File
	file, err := interfaceControllerGenerator.File(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) ScaffoldInterfaceController(driver string, entity model.Entity, writer io.Writer) (error){

	// Vars
	var interfaceControllerGenerator generatorInterface.ControllerGenerator
	if val, ok := generator.interfaceControllerGenerators[driver]; ok {
		interfaceControllerGenerator = val
	}

	// File
	file, err := interfaceControllerGenerator.ScaffoldFile(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) InterfaceControllerEntity(driver string, entity model.Entity, writer io.Writer) (error){

	// Vars
	var interfaceControllerGenerator generatorInterface.ControllerGenerator
	if val, ok := generator.interfaceControllerGenerators[driver]; ok {
		interfaceControllerGenerator = val
	}

	// File
	file, err := interfaceControllerGenerator.EntityFile(entity)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) Registry(entities []model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.registryGenerator.File(entities)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}

func (generator *Generator) ScaffoldRegistry(entities []model.Entity, writer io.Writer) (error){

	// File
	file, err := generator.registryGenerator.ScaffoldFile(entities)
	if err != nil {
		return err
	}

	// Render
	if file != nil {
		file.Render(writer)
	}

	return nil
}