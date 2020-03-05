package turnbull

import(
	"github.com/dave/jennifer/jen"
)

func buildScaffoldUsecaseInteractorFile(entity entity) (*jen.File, error){

	// Entities
	f := jen.NewFile(config.interactorName)

	// Struct
	entityStruct, err := entityToInteractorStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(entityStruct)

	// Line
	f.Line()

	// Interface Methods
	var interactorInterface []jen.Code
	var functions []jen.Code
	for _, method := range entity.Methods {
		
		// Build Method
		methodErr := buildScaffoldUsecaseInteractorMethod(entity, method, &interactorInterface)
		if methodErr != nil {
			return nil, methodErr
		}

		// Build Function
		functionErr := buildScaffoldUsecaseInteractorFunction(entity, method, &functions)
		if functionErr != nil {
			return nil, functionErr
		}


	}

	// Interface
	f.Type().Id(interactorId(entity)).Interface(interactorInterface...)

	// Line
	f.Line()

	// Constructor
	f.Func().
	Id(interactorConstructorId(entity)).
	Params(
		jen.Id("r").Qual(scaffoldRepositoryFilePath(), repositoryId(entity)),
		jen.Id("p").Qual(scaffoldPresenterFilePath(), presenterId(entity)),
	).
	Qual("", interactorId(entity)).
	Block(
		jen.Return(
			jen.Op("&").
			Id(interactorStructId(entity)).
			Values(
				jen.Id("r"),
				jen.Id("p"),
			),
		),
	)

	// Line
	f.Line()

	// Functions
	for _, function := range functions {
		f.Add(function)

		// Line
		f.Line()

	}

	return f, nil
}

func entityToInteractorStruct(entity entity) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// Struct
	resp.Type()

	// Fields
	var fields []jen.Code

	// Repository
	fields = append(fields, jen.Id(repositoryId(entity)).Qual(scaffoldRepositoryFilePath(), repositoryId(entity)) )

	// Presenter
	fields = append(fields, jen.Id(presenterId(entity)).Qual(scaffoldPresenterFilePath(), presenterId(entity)) )

	// Type
	if entity.Name != "" {
		resp.Id(interactorStructId(entity))
	}

	// Fields
	resp.Struct(fields...)

	return &resp, nil

}

func buildScaffoldUsecaseInteractorMethod(entity entity, entityMethod entityMethod, interactorInterface *[]jen.Code)(error){
	
	var method method
	for _, configMethod := range config.methods {
		if configMethod.Type == entityMethod.Type {
			method = configMethod
		}
	}

	var arguments []jen.Code
	var returnValues []jen.Code

	for _, argument := range method.Repository.Arguments {
		code, err := fieldToStructField(entity, argument)
		if err != nil {
			return err
		}
		arguments = append(arguments, code)
	}
	for _, returnValue := range method.Presenter.ReturnValues {
		code, err := fieldToStructField(entity, returnValue)
		if err != nil {
			return err
		}
		returnValues = append(returnValues, code)
	}

	*interactorInterface = append(*interactorInterface, jen.Id(interactorMethodName(entityMethod)).Params(
		arguments...,
	).Parens(jen.List(
		returnValues...,
	)))

	return nil
}

func buildScaffoldUsecaseInteractorFunction(entity entity, entityMethod entityMethod, functions *[]jen.Code)(error){
	
	var method method
	for _, configMethod := range config.methods {
		if configMethod.Type == entityMethod.Type {
			method = configMethod
		}
	}

	var arguments []jen.Code
	var returnValues []jen.Code

	var repositoryArguments []jen.Code
	var repositoryReturnValues []jen.Code

	var presenterArguments []jen.Code

	for _, argument := range method.Repository.Arguments {
		code, err := fieldToStructField(entity, argument)
		if err != nil {
			return err
		}
		arguments = append(arguments, code)
		repositoryArguments = append(repositoryArguments, jen.Id(argument.Name))
	}
	for _, returnValue := range method.Repository.ReturnValues {
		repositoryReturnValues = append(repositoryReturnValues, jen.Id(returnValue.Name))
	}

	for _, argument := range method.Presenter.Arguments {
		var resp jen.Statement
		var matched bool
		for _, returnValue := range method.Repository.ReturnValues {
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

	for _, returnValue := range method.Presenter.ReturnValues {
		code, err := fieldToStructField(entity, returnValue)
		if err != nil {
			return err
		}
		returnValues = append(returnValues, code)
	}

	// Funcs
	*functions = append(*functions, 
		jen.Func().
		Params(
			jen.Id("i").Op("*").Qual("", interactorStructId(entity)),
		).
		Id(interactorMethodName(entityMethod)).
		Params(
			arguments...,
		).
		Parens(jen.List(
			returnValues...,
		)).
		Block(
			jen.List(repositoryReturnValues...).
			Op(":=").
			Id("i").
			Dot(repositoryId(entity)).
			Dot(interactorMethodName(entityMethod)).
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
				Dot(presenterId(entity)).
				Dot(interactorMethodName(entityMethod)).
				Call(
					jen.List(presenterArguments...),
				),
			),
		),
	)

	return nil
}