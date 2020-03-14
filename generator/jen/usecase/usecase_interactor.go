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

	ScaffoldFile(entity model.Entity) (*jen.File, error)

	scaffoldUsecaseInteractorStruct(entity model.Entity) (jen.Statement, error)
	scaffoldUsecaseInteractorInterface(entity model.Entity) (jen.Statement, error)
	scaffoldUsecaseInteractorInterfaceMethod(method model.Method, entity model.Entity) (jen.Statement, error)
	scaffoldUsecaseInteractorConstructorFunction(entity model.Entity) (jen.Statement, error)
	scaffoldUsecaseInteractorMethod(method model.Method, entity model.Entity) (jen.Statement, error)

}

func NewInteractorGenerator(config *config.Config, formatter formatter.Formatter, helperGenerator helper.Generator) InteractorGenerator {
	return &interactorGenerator{
		config : config,
		formatter : formatter,
		helperGenerator :helperGenerator,
	}
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

	// Constructor Function
	usecaseInteractorConstructorFunction, err := interactorGenerator.scaffoldUsecaseInteractorConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&usecaseInteractorConstructorFunction)

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

func (interactorGenerator *interactorGenerator) scaffoldUsecaseInteractorStruct(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Type
	resp.Type()

	// ID
	id , err := interactorGenerator.formatter.OutputScaffoldUsecaseInteractorStructId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Repository
	repositoryId , err := interactorGenerator.formatter.OutputScaffoldUsecaseRepositoryInterfaceId(entity)
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
	presenterId , err := interactorGenerator.formatter.OutputScaffoldUsecasePresenterInterfaceId(entity)
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

func (interactorGenerator *interactorGenerator) scaffoldUsecaseInteractorConstructorFunction(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// Type
	resp.Func()

	// ID
	id , err := interactorGenerator.formatter.OutputScaffoldUsecaseInteractorInterfaceConstructorFunctionId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Repository
	repositoryId , err := interactorGenerator.formatter.OutputScaffoldUsecaseRepositoryInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	repositoryImportPath , err := interactorGenerator.formatter.OutputScaffoldUsecaseRepositoryDirectoryImportPath()
	if err != nil {
		return nil, err
	}

	// Presenter
	presenterId , err := interactorGenerator.formatter.OutputScaffoldUsecasePresenterInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	presenterImportPath , err := interactorGenerator.formatter.OutputScaffoldUsecasePresenterDirectoryImportPath()
	if err != nil {
		return nil, err
	}

	// Struct ID
	structId , err := interactorGenerator.formatter.OutputScaffoldUsecaseInteractorStructId(entity)
	if err != nil {
		return nil, err
	}

	// Interfacr ID
	interfaceId , err := interactorGenerator.formatter.OutputScaffoldUsecaseInteractorInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	// Params
	resp.Params(
		jen.Id("r").
		Qual(repositoryImportPath, repositoryId),
		jen.Id("p").
		Qual(presenterImportPath, presenterId),
	).
	Qual("", interfaceId).
	Block(
		jen.Return(
			jen.Op("&").
			Id(structId).
			Values(
				jen.Id("r"),
				jen.Id("p"),
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
	resp.Block(
		jen.List(repositoryReturnValues...).
		Op(":=").
		Id("i").
		Dot(repositoryPackageName).
		Dot(repositoryInterfaceMethodId).
		Call(
			jen.List(repositoryArguments...),
		),
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
		jen.Return(
			jen.Id("i").
			Dot(presenterPackageName).
			Dot(presenterInterfaceMethodId).
			Call(
				jen.List(presenterArguments...),
			),
		),
	)

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