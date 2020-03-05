package turnbull

import(
	"github.com/dave/jennifer/jen"
)

func buildScaffoldInterfaceRepositoryFile(entity entity) (*jen.File, error){

	// Entities
	f := jen.NewFile(config.repositoryName)

	// Struct Model
	var scribbleRepositoryStruct []jen.Code
	scribbleRepositoryStruct = append(scribbleRepositoryStruct, jen.Id("driver").Op("*").Qual("github.com/nanobox-io/golang-scribble", "Driver"))
	scribbleRepositoryStruct = append(scribbleRepositoryStruct, jen.Id("collection").String())
	f.Type().Id(interfaceRepositoryStructId(entity, "scribble")).Struct(scribbleRepositoryStruct...)

	// New
	f.Func().Id(interfaceRepositoryConstructorId(entity, "scribble")).Params(
		scribbleRepositoryStruct...,
	).Op("*").Qual("", interfaceRepositoryStructId(entity, "scribble")).Block(
		jen.Return(
			jen.Op("&").Id(interfaceRepositoryStructId(entity, "scribble")).Values(
				jen.Dict{
					jen.Id("driver"):	jen.Id("driver"),
					jen.Id("collection"):	jen.Id("collection"),
				},
			),
		),
	)

	var functions []jen.Code
	for _, entityMethod := range entity.Methods {

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
				return nil, err
			}
			arguments = append(arguments, code)
		}
		for _, returnValue := range method.Repository.ReturnValues {
			code, err := fieldToStructField(entity, returnValue)
			if err != nil {
				return nil, err
			}
			returnValues = append(returnValues, code)
		}

		// Block
		var block []jen.Code

		if entityMethod.Type == "add" {
			block = append(block, 
				jen.Err().
				Op("=").
				Id("r").
				Dot("driver").
				Dot("Write").
				Call(
					jen.Id("r").Dot("collection"),
					jen.Id("req").
					Dot("Primary").
					Call(),
					jen.Id("req"),
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
			)	
		}
		block = append(block, jen.Return(
			jen.Nil(),
			jen.Nil(),
		))

		functions = append(functions, 
			jen.Func().
			Params(
				jen.Id("r").Op("*").Qual("", interfaceRepositoryStructId(entity, "scribble")),
			).
			Id(repositoryMethodName(entityMethod)).
			Params(
				arguments...,
			).
			Parens(jen.List(
				returnValues...,
			)).
			Block(block...),
		)

	}

	// Functions
	for _, function := range functions {
		f.Add(function)

		// Line
		f.Line()

	}

	return f, nil
}

func buildScaffoldUsecaseRepositoryFile(entity entity) (*jen.File, error){

	// Entities
	f := jen.NewFile(config.repositoryName)

	// Interface Methods
	var repositoryInterface []jen.Code
	for _, method := range entity.Methods {
		
		// Build Method
		err := buildScaffoldUsecaseRepositoryMethod(entity, method, &repositoryInterface)
		if err != nil {
			return nil, err
		}
	}

	// Interface
	f.Type().Id(repositoryId(entity)).Interface(repositoryInterface...)

	return f, nil
}

func buildScaffoldUsecaseRepositoryMethod(entity entity, entityMethod entityMethod, repositoryInterface *[]jen.Code)(error){
	
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
	for _, returnValue := range method.Repository.ReturnValues {
		code, err := fieldToStructField(entity, returnValue)
		if err != nil {
			return err
		}
		returnValues = append(returnValues, code)
	}

	*repositoryInterface = append(*repositoryInterface, jen.Id(repositoryMethodName(entityMethod)).Params(
		arguments...,
	).Parens(jen.List(
		returnValues...,
	)))

	return nil
}

