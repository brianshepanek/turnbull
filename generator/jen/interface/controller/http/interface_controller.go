package generator

import(
	"github.com/dave/jennifer/jen"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/generator/jen/helper"
)

type controllerGenerator struct{
	config *config.Config
	formatter formatter.Formatter
	helperGenerator helper.Generator
}

// type ControllerGenerator interface{

// 	ScaffoldFile(entity model.Entity) (*jen.File, error)
	
// 	scaffoldInterfaceControllerStructFields(entity model.Entity) ([]jen.Code, error)
// 	scaffoldInterfaceControllerStruct(entity model.Entity) (jen.Statement, error)
// 	scaffoldInterfaceControllerConstructorFunction(entity model.Entity) (jen.Statement, error)
// 	scaffoldInterfaceControllerMethod(method model.Method, entity model.Entity) (jen.Statement, error)

// }

func New(config *config.Config, formatter formatter.Formatter, helperGenerator helper.Generator) *controllerGenerator {
	return &controllerGenerator{
		config : config,
		formatter : formatter,
		helperGenerator : helperGenerator,
	}
}

func (controllerGenerator *controllerGenerator) ScaffoldFile(entity model.Entity) (*jen.File, error){

	// File
	packageName , err := controllerGenerator.formatter.OutputScaffoldInterfaceControllerPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	// Struct
	interfaceControllerStruct, err := controllerGenerator.scaffoldInterfaceControllerStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceControllerStruct)

	// Constructor Function
	interfaceControllerConstructorFunction, err := controllerGenerator.scaffoldInterfaceControllerConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceControllerConstructorFunction)

	// Methods
	for _, entityMethod := range entity.Methods {

		// Function
		method, err := controllerGenerator.scaffoldInterfaceControllerMethod(entityMethod, entity)
		if err != nil {
			return nil, err
		}
		f.Add(&method)

	}

	return f, nil
}	

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerStructFields(entity model.Entity) ([]jen.Code, error){

	// Vars
	var fields []jen.Code

	// Interactor
	interactorId , err := controllerGenerator.formatter.OutputScaffoldUsecaseInteractorInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	interactorImportPath , err := controllerGenerator.formatter.OutputScaffoldUsecaseInteractorDirectoryImportPath()
	if err != nil {
		return nil, err
	}

	interactorPackageName , err := controllerGenerator.formatter.OutputScaffoldUsecaseInteractorPackageName()
	if err != nil {
		return nil, err
	}

	// Interactor
	fields = append(fields, 
		jen.Id(interactorPackageName).
		Qual(interactorImportPath, interactorId), 
	)

	return fields, nil

}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerStruct(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// Fields
	fields, err := controllerGenerator.scaffoldInterfaceControllerStructFields(entity)
	if err != nil {
		return nil, err
	}

	// Type
	resp.Type()

	// ID
	id , err := controllerGenerator.formatter.OutputScaffoldInterfaceControllerStructId("http", entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Struct
	resp.Struct(fields...)

	
	return resp, nil

}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerConstructorFunction(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	
	// Type
	resp.Func()

	// Fields
	fields, err := controllerGenerator.scaffoldInterfaceControllerStructFields(entity)
	if err != nil {
		return nil, err
	}

	interactorPackageName , err := controllerGenerator.formatter.OutputScaffoldUsecaseInteractorPackageName()
	if err != nil {
		return nil, err
	}

	// ID
	id , err := controllerGenerator.formatter.OutputScaffoldInterfaceControllerConstructorFunctionId("http", entity)
	if err != nil {
		return nil, err
	}

	// Struct ID
	structId , err := controllerGenerator.formatter.OutputScaffoldInterfaceControllerStructId("http", entity)
	if err != nil {
		return nil, err
	}

	resp.Id(id)

	// Params
	resp.Params(fields...)

	// Qual
	resp.Op("*")
	resp.Qual("", structId)

	// Block
	resp.Block(
		jen.Return(
			jen.Op("&").
			Id(structId).
			Values(
				jen.Id(interactorPackageName),
			),
		),
	)
	
	return resp, nil

}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerMethod(method model.Method, entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// Config
	configMethod, err := controllerGenerator.config.Method(method)
	
	var arguments, returnValues, repositoryArguments, repositoryReturnValues, presenterArguments []jen.Code
	for _, argument := range configMethod.Repository.Arguments {

		var statement jen.Statement

		// Field
		err = controllerGenerator.helperGenerator.Field(argument.Name, argument, entity, &statement)
		if err != nil {
			return nil, err
		}

		// Append
		arguments = append(arguments, &statement)
		repositoryArguments = append(repositoryArguments, jen.Id(argument.Name))

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
		err = controllerGenerator.helperGenerator.Field("", returnValue, entity, &statement)
		if err != nil {
			return nil, err
		}

		// Append
		returnValues = append(returnValues, &statement)

	}

	// ID
	id , err := controllerGenerator.formatter.OutputScaffoldUsecaseInteractorInterfaceMethodId(method)
	if err != nil {
		return nil, err
	}

	// Struct
	structId , err := controllerGenerator.formatter.OutputScaffoldInterfaceControllerStructId("http", entity)
	if err != nil {
		return nil, err
	}

	// Func
	resp.Func()

	// Params
	resp.Params(
		jen.Id("c").
		Op("*").
		Qual("", structId),
	)

	// ID
	resp.Id(id)

	// Params
	resp.Params(
		jen.List(
			jen.Id("w").
			Qual("net/http", "ResponseWriter"),
			jen.Id("r").
			Qual("net/http", "Request"),
		),
	)

	// Block
	resp.Block(
		jen.Id("w").
		Dot("WriteHeader").
		Call(
			jen.Qual("net/http", "StatusOK"),
		),
		jen.Qual("encoding/json", "json").
		Dot("NewEncoder").
		Call(
			jen.Id("w"),
		).
		Dot("Encode").
		Call(
			jen.Interface(),
		),
	)

	return resp, nil
}
