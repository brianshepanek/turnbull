package usecase

import(
	"github.com/dave/jennifer/jen"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/generator/jen/helper"
)

type interactorGenerator struct{
	config *config.Config
	formatter formatter.Formatter
	helperGenerator helper.Generator
}

type InteractorGenerator interface{

	File(entity model.Entity) (*jen.File, error)
	ScaffoldFile(entity model.Entity) (*jen.File, error)
	RegistryFile(entity model.Entity) (*jen.File, error)

	scaffoldUsecaseInteractorStruct(entity model.Entity) (jen.Statement, error)
	scaffoldUsecaseInteractorInterface(entity model.Entity) (jen.Statement, error)
	scaffoldUsecaseInteractorInterfaceMethod(method model.Method, entity model.Entity) (jen.Statement, error)
	usecaseInteractorConstructorFunction(entity model.Entity) (jen.Statement, error)
	scaffoldUsecaseInteractorMethod(method model.Method, entity model.Entity) (jen.Statement, error)

}

func NewInteractorGenerator(config *config.Config, formatter formatter.Formatter, helperGenerator helper.Generator) InteractorGenerator {
	return &interactorGenerator{
		config : config,
		formatter : formatter,
		helperGenerator :helperGenerator,
	}
}

func (interactorGenerator *interactorGenerator) File(entity model.Entity) (*jen.File, error){
	
	
	// File
	packageName , err := interactorGenerator.formatter.OutputScaffoldUsecaseInteractorPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	// Struct
	usecaseInteractorStruct, err := interactorGenerator.usecaseInteractorStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&usecaseInteractorStruct)

	// Interface
	usecaseInteractorInterface, err := interactorGenerator.usecaseInteractorInterface(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&usecaseInteractorInterface)

	// Constructor Function
	usecaseInteractorConstructorFunction, err := interactorGenerator.usecaseInteractorConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&usecaseInteractorConstructorFunction)

	return f, nil
}	

func (interactorGenerator *interactorGenerator) ScaffoldFile(entity model.Entity) (*jen.File, error){
	
	
	// File
	packageName , err := interactorGenerator.formatter.OutputScaffoldUsecaseInteractorPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	// Struct
	usecaseInteractorStruct, err := interactorGenerator.scaffoldUsecaseInteractorStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&usecaseInteractorStruct)

	// Interface
	usecaseInteractorInterface, err := interactorGenerator.scaffoldUsecaseInteractorInterface(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&usecaseInteractorInterface)

	// Functions
	for _, method := range entity.Methods {

		// Function
		function, err := interactorGenerator.scaffoldUsecaseInteractorMethod(method, entity)
		if err != nil {
			return nil, err
		}
		f.Add(&function)

	}

	return f, nil
}	

func (interactorGenerator *interactorGenerator) RegistryFile(entity model.Entity) (*jen.File, error){

	// File
	packageName , err := interactorGenerator.formatter.OutputRegistryPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)
	
	// Local Constructor Function
	usecaseInteractorRegistryLocalConstructorFunction, err := interactorGenerator.usecaseInteractorRegistryLocalConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&usecaseInteractorRegistryLocalConstructorFunction)


	return f, nil
}

func (interactorGenerator *interactorGenerator) usecaseInteractorRegistryLocalConstructorFunction(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// Registry
	registryName , err := interactorGenerator.formatter.OutputRegistryPackageName()
	if err != nil {
		return nil, err
	}

	interfaceId , err := interactorGenerator.formatter.OutputUsecaseInteractorInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	// Interactor Import Path
	importPath , err := interactorGenerator.formatter.OutputScaffoldUsecaseInteractorDirectoryImportPath()
	if err != nil {
		return nil, err
	}
	
	// ID
	id , err := interactorGenerator.formatter.OutputScaffoldUsecaseInteractorRegistryLocalConstructorFunctionId(entity)
	if err != nil {
		return nil, err
	}

	constructorFunctionId , err := interactorGenerator.formatter.OutputUsecaseInteractorInterfaceConstructorFunctionId(entity)
	if err != nil {
		return nil, err
	}
	


	// Primary Repository Function
	var primaryRepositoryConstructorId string
	for _, repository := range entity.Repositories {
		if repository.Primary {

			// ID
			id, err := interactorGenerator.formatter.OutputScaffoldInterfaceRepositoryRegistryLocalConstructorFunctionId(repository.Type, entity)
			if err != nil {
				return nil, err
			}
			primaryRepositoryConstructorId = id

		}
	}

	// Primary Presenter Function
	var primaryPresenterConstructorId string
	for _, presenter := range entity.Presenters {
		if presenter.Primary {

			// ID
			id, err := interactorGenerator.formatter.OutputScaffoldInterfacePresenterRegistryLocalConstructorFunctionId(presenter.Type, entity)
			if err != nil {
				return nil, err
			}
			primaryPresenterConstructorId = id

		}
	}
	

	// Type
	resp.Func()

	resp.Params(
		jen.Id("r").
		Op("*").
		Qual("", registryName),
	)

	// ID
	resp.Id(id)

	// Params
	resp.Params()

	resp.Parens(
		jen.List(
			jen.Qual(importPath, interfaceId),
		),
	)

	// Fields
	var fields []jen.Code

	fields = append(fields, 
		jen.Id("r").
		Dot(primaryRepositoryConstructorId).
		Call(),
	)

	fields = append(fields, 
		jen.Id("r").
		Dot(primaryPresenterConstructorId).
		Call(),
	)

	for _, interactor := range entity.Interactors {
		
		// ID
		injectedInteractorConstructorFunctionId , err := interactorGenerator.formatter.OutputScaffoldUsecaseInteractorRegistryLocalConstructorFunctionId(model.Entity{Name : interactor})
		if err != nil {
			return nil, err
		}

		fields = append(fields, 
			jen.Id("r").
			Dot(injectedInteractorConstructorFunctionId).
			Call(),
		)

	}

	

	resp.Block(
		jen.Return(
			jen.Qual(importPath, constructorFunctionId).
			Params(fields...),
		),
	)
	
	return resp, nil

}

func (interactorGenerator *interactorGenerator) usecaseInteractorStruct(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Type
	resp.Type()

	// ID
	id , err := interactorGenerator.formatter.OutputUsecaseInteractorStructId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Scaffold
	scaffoldId , err := interactorGenerator.formatter.OutputScaffoldUsecaseInteractorStructId(entity)
	if err != nil {
		return nil, err
	}
	
	// Fields
	fields = append(fields, jen.Id(scaffoldId))


	// Struct
	resp.Struct(fields...)

	return resp, nil

}

func (interactorGenerator *interactorGenerator) usecaseInteractorInterface(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Type
	resp.Type()

	// ID
	id , err := interactorGenerator.formatter.OutputUsecaseInteractorInterfaceId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Scaffold
	scaffoldId , err := interactorGenerator.formatter.OutputScaffoldUsecaseInteractorInterfaceId(entity)
	if err != nil {
		return nil, err
	}
	
	// Fields
	fields = append(fields, jen.Id(scaffoldId))


	// Interface
	resp.Interface(fields...)

	return resp, nil

}

func (interactorGenerator *interactorGenerator) scaffoldUsecaseInteractorStruct(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// Type
	resp.Type()

	// ID
	id , err := interactorGenerator.formatter.OutputScaffoldUsecaseInteractorStructId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	fields, err := interactorGenerator.scaffoldUsecaseInteractorStructFields(entity)
	if err != nil {
		return nil, err
	}
	
	// Struct
	resp.Struct(fields...)

	return resp, nil

}

func (interactorGenerator *interactorGenerator) scaffoldUsecaseInteractorInterface(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	var methods []jen.Code

	// Loop
	for _, method := range entity.Methods {
		
		// Method
		statement, err := interactorGenerator.scaffoldUsecaseInteractorInterfaceMethod(method, entity)
		if err != nil {
			return nil, err
		}

		// Append
		methods = append(methods, &statement)
	}

	// Type
	resp.Type()

	// ID
	id , err := interactorGenerator.formatter.OutputScaffoldUsecaseInteractorInterfaceId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Interface
	resp.Interface(methods...)

	return resp, nil

}

func (interactorGenerator *interactorGenerator) usecaseInteractorConstructorFunction(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// Type
	resp.Func()

	// ID
	id , err := interactorGenerator.formatter.OutputUsecaseInteractorInterfaceConstructorFunctionId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// // Repository
	// repositoryId , err := interactorGenerator.formatter.OutputUsecaseRepositoryInterfaceId(entity)
	// if err != nil {
	// 	return nil, err
	// }

	// repositoryImportPath , err := interactorGenerator.formatter.OutputScaffoldUsecaseRepositoryDirectoryImportPath()
	// if err != nil {
	// 	return nil, err
	// }

	// // Presenter
	// presenterId , err := interactorGenerator.formatter.OutputUsecasePresenterInterfaceId(entity)
	// if err != nil {
	// 	return nil, err
	// }

	// presenterImportPath , err := interactorGenerator.formatter.OutputScaffoldUsecasePresenterDirectoryImportPath()
	// if err != nil {
	// 	return nil, err
	// }

	// Struct ID
	structId , err := interactorGenerator.formatter.OutputUsecaseInteractorStructId(entity)
	if err != nil {
		return nil, err
	}

	// Scaffold Struct ID
	scaffoldStructId , err := interactorGenerator.formatter.OutputScaffoldUsecaseInteractorStructId(entity)
	if err != nil {
		return nil, err
	}

	// Interfacr ID
	interfaceId , err := interactorGenerator.formatter.OutputUsecaseInteractorInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	fields, err := interactorGenerator.scaffoldUsecaseInteractorStructFields(entity)
	if err != nil {
		return nil, err
	}

	fieldNames, err := interactorGenerator.scaffoldUsecaseInteractorStructFieldNames(entity)
	if err != nil {
		return nil, err
	}

	// Params
	resp.Params(
		fields...,
	).
	Qual("", interfaceId).
	Block(
		jen.Return(
			jen.Op("&").
			Id(structId).
			Values(
				jen.Id(scaffoldStructId).
				Values(
					fieldNames...,
				),
			),
		),
	)
	
	return resp, nil

}

func (interactorGenerator *interactorGenerator) scaffoldUsecaseInteractorMethod(method model.Method, entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// Config
	configMethod, err := interactorGenerator.config.Method(method)
	
	var entityVar, contextVar string
	var arguments, returnValues, repositoryArguments, repositoryReturnValues, presenterArguments []jen.Code
	for _, argument := range configMethod.Repository.Arguments {

		var statement jen.Statement

		// Field
		err = interactorGenerator.helperGenerator.Field(argument.Name, argument, entity, &statement)
		if err != nil {
			return nil, err
		}

		// Append
		arguments = append(arguments, &statement)
		repositoryArguments = append(repositoryArguments, jen.Id(argument.Name))

		// Entity Var
		if argument.Type == "self" {
			entityVar = argument.Name
		}

		// Context Var
		if argument.Type == "Context" {
			contextVar = argument.Name
		}

	}

	for _, returnValue := range configMethod.Repository.ReturnValues {

		repositoryReturnValues = append(repositoryReturnValues, jen.Id(returnValue.Name))

	}

	for _, argument := range configMethod.Presenter.Arguments {
		var resp jen.Statement
		var matched bool
		for _, returnValue := range configMethod.Repository.ReturnValues {
			if returnValue.Package == argument.Package && returnValue.Type == argument.Type {
				matched = true
				if returnValue.Op == "*" && argument.Op != "*" {
					resp.Op("*")
				}
				resp.Id(returnValue.Name)
			}
		}
		if !matched {
			resp.Id(argument.Name)
		}
		presenterArguments = append(presenterArguments, &resp)
	}

	for _, returnValue := range configMethod.Presenter.ReturnValues {

		var statement jen.Statement

		// Field
		err = interactorGenerator.helperGenerator.Field("", returnValue, entity, &statement)
		if err != nil {
			return nil, err
		}

		// Append
		returnValues = append(returnValues, &statement)

	}

	// Callbacks
	var hasBefore, hasAfter bool
	var beforeId, afterId string
	for _, callback := range method.Callbacks {
		
		// Before
		if callback.Type == "before" {
			hasBefore = true

			// Struct ID
			beforeId , err = interactorGenerator.formatter.OutputScaffoldDomainEntityCallbackId(callback, method)
			if err != nil {
				return nil, err
			}

		}

		// After
		if callback.Type == "after" {
			hasAfter = true

			// Struct ID
			afterId , err = interactorGenerator.formatter.OutputScaffoldDomainEntityCallbackId(callback, method)
			if err != nil {
				return nil, err
			}

		}

	}

	// ID
	id , err := interactorGenerator.formatter.OutputScaffoldUsecaseInteractorInterfaceMethodId(method)
	if err != nil {
		return nil, err
	}

	// Repository
	repositoryInterfaceMethodId , err := interactorGenerator.formatter.OutputScaffoldUsecaseRepositoryInterfaceMethodId(method)
	if err != nil {
		return nil, err
	}

	repositoryPackageName , err := interactorGenerator.formatter.OutputScaffoldUsecaseRepositoryPackageName()
	if err != nil {
		return nil, err
	}

	// Presenter
	presenterInterfaceMethodId , err := interactorGenerator.formatter.OutputScaffoldUsecasePresenterInterfaceMethodId(method)
	if err != nil {
		return nil, err
	}

	presenterPackageName , err := interactorGenerator.formatter.OutputScaffoldUsecasePresenterPackageName()
	if err != nil {
		return nil, err
	}

	// Struct
	structId , err := interactorGenerator.formatter.OutputScaffoldUsecaseInteractorStructId(entity)
	if err != nil {
		return nil, err
	}

	// Func
	resp.Func()

	// Params
	resp.Params(
		jen.Id("i").
		Op("*").
		Qual("", structId),
	)

	// ID
	resp.Id(id)

	// Params
	resp.Params(arguments...)

	// Parens
	resp.Parens(
		jen.List(returnValues...),
	)	

	// Block
	var block []jen.Code

	// Var
	block = append(block,
		jen.Var().
		Id("err").
		Error(),
	)

	// Has Before
	if hasBefore{
		
		block = append(block, 
			jen.Err().
			Op("=").
			Id(entityVar).
			Dot(beforeId).
			Params(
				jen.Id(contextVar),
			),
		)
		block = append(block, 
			jen.If(
				jen.Err().
				Op("!=").
				Nil(),
			).
			Block(
				jen.Return(
					jen.List(
						jen.Nil(),
						jen.Err(),
					),
				),
			),
		)
	}
	

	block = append(block,
		jen.List(repositoryReturnValues...).
		Op("=").
		Id("i").
		Dot(repositoryPackageName).
		Dot(repositoryInterfaceMethodId).
		Call(
			jen.List(repositoryArguments...),
		),
	)

	block = append(block,
		jen.If(
			jen.Err().
			Op("!=").
			Nil(),
		).
		Block(
			jen.Return(
				jen.List(
					jen.Nil(),
					jen.Err(),
				),
			),
		),
	)

	// Has After
	if hasAfter {
		block = append(block, 
			jen.Err().
			Op("=").
			Id(entityVar).
			Dot(afterId).
			Params(
				jen.Id(contextVar),
			),
		)
	
		block = append(block, 
			jen.If(
				jen.Err().
				Op("!=").
				Nil(),
			).
			Block(
				jen.Return(
					jen.List(
						jen.Nil(),
						jen.Err(),
					),
				),
			),
		)
	}
	

	block = append(block,
		jen.Return(
			jen.Id("i").
			Dot(presenterPackageName).
			Dot(presenterInterfaceMethodId).
			Call(
				jen.List(presenterArguments...),
			),
		),
	)
	
	resp.Block(block...)

	return resp, nil

}

func (interactorGenerator *interactorGenerator) scaffoldUsecaseInteractorInterfaceMethod(method model.Method, entity model.Entity) (jen.Statement, error){
	
	// Vars
	var resp jen.Statement

	// Config
	configMethod, err := interactorGenerator.config.Method(method)

	var arguments, returnValues []jen.Code
	for _, argument := range configMethod.Repository.Arguments {

		var statement jen.Statement

		// Field
		err = interactorGenerator.helperGenerator.Field(argument.Name, argument, entity, &statement)
		if err != nil {
			return nil, err
		}

		// Append
		arguments = append(arguments, &statement)

	}

	for _, returnValue := range configMethod.Presenter.ReturnValues {

		var statement jen.Statement

		// Field
		err = interactorGenerator.helperGenerator.Field("", returnValue, entity, &statement)
		if err != nil {
			return nil, err
		}

		// Append
		returnValues = append(returnValues, &statement)

	}

	// ID
	id , err := interactorGenerator.formatter.OutputScaffoldUsecaseInteractorInterfaceMethodId(method)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Params
	resp.Params(arguments...)

	// Parens
	resp.Parens(
		jen.List(returnValues...),
	)	
	
	return resp, nil
}

func (interactorGenerator *interactorGenerator) scaffoldUsecaseInteractorStructFields(entity model.Entity) ([]jen.Code, error){

	// Vars
	var fields []jen.Code

	// Repository
	repositoryId , err := interactorGenerator.formatter.OutputUsecaseRepositoryInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	repositoryImportPath , err := interactorGenerator.formatter.OutputScaffoldUsecaseRepositoryDirectoryImportPath()
	if err != nil {
		return nil, err
	}

	repositoryPackageName , err := interactorGenerator.formatter.OutputScaffoldUsecaseRepositoryPackageName()
	if err != nil {
		return nil, err
	}

	// Presenter
	presenterId , err := interactorGenerator.formatter.OutputUsecasePresenterInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	presenterImportPath , err := interactorGenerator.formatter.OutputScaffoldUsecasePresenterDirectoryImportPath()
	if err != nil {
		return nil, err
	}

	presenterPackageName , err := interactorGenerator.formatter.OutputScaffoldUsecasePresenterPackageName()
	if err != nil {
		return nil, err
	}

	// Repository
	fields = append(fields, 
		jen.Id(repositoryPackageName).
		Qual(repositoryImportPath, repositoryId), 
	)

	// Presenter
	fields = append(fields, 
		jen.Id(presenterPackageName).
		Qual(presenterImportPath, presenterId), 
	)

	for _, interactor := range entity.Interactors {
		
		injectedInteractorStructId , err := interactorGenerator.formatter.OutputUsecaseInteractorStructId(model.Entity{Name : interactor})
		if err != nil {
			return nil, err
		}

		injectedInteractorInterfaceId , err := interactorGenerator.formatter.OutputUsecaseInteractorInterfaceId(model.Entity{Name : interactor})
		if err != nil {
			return nil, err
		}
		
		// Presenter
		fields = append(fields, 
			jen.Id(injectedInteractorStructId).
			Qual("", injectedInteractorInterfaceId), 
		)

	}

	return fields, nil

}

func (interactorGenerator *interactorGenerator) scaffoldUsecaseInteractorStructFieldNames(entity model.Entity) ([]jen.Code, error){

	// Vars
	var fields []jen.Code

	repositoryPackageName , err := interactorGenerator.formatter.OutputScaffoldUsecaseRepositoryPackageName()
	if err != nil {
		return nil, err
	}

	presenterPackageName , err := interactorGenerator.formatter.OutputScaffoldUsecasePresenterPackageName()
	if err != nil {
		return nil, err
	}

	// Repository
	fields = append(fields, 
		jen.Id(repositoryPackageName),
	)

	// Presenter
	fields = append(fields, 
		jen.Id(presenterPackageName),
	)

	for _, interactor := range entity.Interactors {
		
		injectedInteractorStructId , err := interactorGenerator.formatter.OutputUsecaseInteractorStructId(model.Entity{Name : interactor})
		if err != nil {
			return nil, err
		}

		// Presenter
		fields = append(fields, 
			jen.Id(injectedInteractorStructId),
		)

	}

	return fields, nil

}