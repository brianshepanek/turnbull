package registry

import(
	"github.com/dave/jennifer/jen"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/generator/jen/helper"
)

type registryGenerator struct{
	formatter formatter.Formatter
	helperGenerator helper.Generator
}

type RegistryGenerator interface{

	// File
	File(entity model.Entity) (*jen.File, error)

}

func New(formatter formatter.Formatter, helperGenerator helper.Generator) RegistryGenerator {
	return &registryGenerator{
		formatter : formatter,
		helperGenerator : helperGenerator,
	}
}

func (registryGenerator *registryGenerator) File(entity model.Entity) (*jen.File, error){
	
	
	// File
	packageName , err := registryGenerator.formatter.OutputRegistryPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	// Presenter Constructor Function
	presenterConstructorFunction, err := registryGenerator.presenterConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&presenterConstructorFunction)
	f.Line()

	// Repository Constructor Function
	repositoryConstructorFunction, err := registryGenerator.repositoryConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&repositoryConstructorFunction)
	f.Line()
	
	
	return f, nil
}

func (registryGenerator *registryGenerator) presenterConstructorFunction(entity model.Entity) (jen.Statement, error){

	// ID
	var statement jen.Statement
	packageName , err := registryGenerator.formatter.OutputRegistryPackageName()
	if err != nil {
		return nil, err
	}

	id , err := registryGenerator.formatter.OutputRegistryEntityPresenterConstructorFunctionId(entity)
	if err != nil {
		return nil, err
	}

	usecaseImportPath , err := registryGenerator.formatter.OutputScaffoldUsecasePresenterDirectoryImportPath()
	if err != nil {
		return nil, err
	}

	usecaseInterfaceId , err := registryGenerator.formatter.OutputUsecasePresenterInterfaceId(entity)
	if err != nil {
		return nil, err
	}
	

	// Func
	statement.Func()

	// Params
	statement.Params(
		jen.Op("*").
		Id(packageName),
	)

	// ID
	statement.Id(id)

	// Params
	statement.Params()

	// List
	statement.List(
		jen.Qual(usecaseImportPath, usecaseInterfaceId),
	)

	return statement, nil
}

func (registryGenerator *registryGenerator) repositoryConstructorFunction(entity model.Entity) (jen.Statement, error){

	// ID
	var statement jen.Statement
	packageName , err := registryGenerator.formatter.OutputRegistryPackageName()
	if err != nil {
		return nil, err
	}

	id , err := registryGenerator.formatter.OutputRegistryEntityRepositoryConstructorFunctionId(entity)
	if err != nil {
		return nil, err
	}

	usecaseImportPath , err := registryGenerator.formatter.OutputScaffoldUsecaseRepositoryDirectoryImportPath()
	if err != nil {
		return nil, err
	}

	usecaseInterfaceId , err := registryGenerator.formatter.OutputUsecaseRepositoryInterfaceId(entity)
	if err != nil {
		return nil, err
	}
	

	// Func
	statement.Func()

	// Params
	statement.Params(
		jen.Op("*").
		Id(packageName),
	)

	// ID
	statement.Id(id)

	// Params
	statement.Params()

	// List
	statement.List(
		jen.Qual(usecaseImportPath, usecaseInterfaceId),
	)

	return statement, nil
}