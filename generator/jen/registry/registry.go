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
	File(entities []model.Entity) (*jen.File, error)
	ScaffoldFile(entities []model.Entity) (*jen.File, error)

}

func New(formatter formatter.Formatter, helperGenerator helper.Generator) RegistryGenerator {
	return &registryGenerator{
		formatter : formatter,
		helperGenerator : helperGenerator,
	}
}

func (registryGenerator *registryGenerator) File(entities []model.Entity) (*jen.File, error){
	
	
	// File
	packageName , err := registryGenerator.formatter.OutputRegistryPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	return f, nil
}

func (registryGenerator *registryGenerator) ScaffoldFile(entities []model.Entity) (*jen.File, error){
	
	// File
	packageName , err := registryGenerator.formatter.OutputRegistryPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	// Struct
	buildStruct, err := registryGenerator.buildStruct(entities)
	if err != nil {
		return nil, err
	}
	f.Add(&buildStruct)


	// Constructor
	constructorFunction, err := registryGenerator.buildConstructorFunction(entities)
	if err != nil {
		return nil, err
	}
	f.Add(&constructorFunction)

	// App Controllers
	appControllers := make(map[string][]model.Entity)
	for _, entity := range entities {
		for _, controller := range entity.Controllers {
			appControllers[controller.Type] = append(appControllers[controller.Type], entity)
		}
	}

	for appController, appControllerEntities := range appControllers {

		appControllerConstructorFunction, err := registryGenerator.buildAppControllerConstructorFunction(appController, appControllerEntities)
		if err != nil {
			return nil, err
		}
		f.Add(&appControllerConstructorFunction)
	}

	return f, nil
}

func (registryGenerator *registryGenerator) buildStruct(entities []model.Entity) (jen.Statement, error){

	var resp jen.Statement
	var fields []jen.Code

	for _, entity := range entities {
		for _, repository := range entity.Repositories {
			
			// ID
			registryStructId , err := registryGenerator.formatter.OutputScaffoldInterfaceRepositoryRegistryStructId(repository.Type, entity)
			if err != nil {
				return nil, err
			}

			fields = append(fields, jen.Id(registryStructId))

		}

		for _, presenter := range entity.Presenters {
			
			// ID
			registryStructId , err := registryGenerator.formatter.OutputScaffoldInterfacePresenterRegistryStructId(presenter.Type, entity)
			if err != nil {
				return nil, err
			}

			fields = append(fields, jen.Id(registryStructId))

		}
		
	}

	// Type
	resp.Type()

	// ID
	resp.Id("registry")

	// Struct
	resp.Struct(fields...)

	return resp, nil
}

func (registryGenerator *registryGenerator) buildConstructorFunction(entities []model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement


	// Type
	resp.Func()


	resp.Id("New")

	resp.Params()

	// Qual
	resp.List(
		jen.Op("*").
		Id("registry"),
	)

	// Block
	resp.Block(
		jen.Return(
			jen.Op("&").
			Id("registry").
			Values(),
		),
	)
	
	return resp, nil

}

func (registryGenerator *registryGenerator) buildAppControllerConstructorFunction(driver string, entities []model.Entity) (jen.Statement, error){

	// ID
	id , err := registryGenerator.formatter.OutputScaffoldInterfaceControllerRegistryLocalConstructorFunctionId(driver, model.Entity{Name : "app"})
	if err != nil {
		return nil, err
	}

	// Import Path
	importPath , err := registryGenerator.formatter.OutputInterfaceControllerDirectoryImportPath(driver, model.Entity{Name : "app"})
	if err != nil {
		return nil, err
	}

	// Interface ID
	interfaceId , err := registryGenerator.formatter.OutputInterfaceControllerInterfaceId(driver, model.Entity{Name : "app"})
	if err != nil {
		return nil, err
	}
	
	

	// Vars
	var resp jen.Statement
	
	// Type
	resp.Func()

	resp.Params(
		jen.Id("r").
		Op("*").
		Id("registry"),
	)


	resp.Id(id)

	resp.Params()

	// Qual
	resp.List(
		jen.Qual(importPath, interfaceId),
	)

	// Block
	var block []jen.Code

	for _, entity := range entities {
		
		// ID
		id , err := registryGenerator.formatter.OutputScaffoldInterfaceControllerRegistryLocalConstructorFunctionId(driver, entity)
		if err != nil {
			return nil, err
		}

		block = append(block,
			jen.Id("r").
			Dot(id).
			Call(),
		)

	}

	// Return
	block = append(block, 
		jen.Return(
			jen.Id("r"),
		),
	)

	resp.Block(block...)
	
	return resp, nil

}