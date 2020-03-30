package generator

import(
	"strings"
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

func (controllerGenerator *controllerGenerator) File(entity model.Entity) (*jen.File, error){

	// File
	packageName , err := controllerGenerator.formatter.OutputScaffoldInterfaceControllerPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	// Struct
	interfaceControllerStruct, err := controllerGenerator.interfaceControllerStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceControllerStruct)

	// Interface
	interfaceControllerInterface, err := controllerGenerator.interfaceControllerInterface(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceControllerInterface)

	// Constructor Function
	interfaceControllerConstructorFunction, err := controllerGenerator.interfaceControllerConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceControllerConstructorFunction)

	return f, nil
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

	// Interface
	interfaceControllerInterface, err := controllerGenerator.scaffoldInterfaceControllerInterface(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceControllerInterface)

	// Entity Struct
	interfaceRepositoryEntityStruct, err := controllerGenerator.scaffoldInterfaceControllerEntityStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceRepositoryEntityStruct)

	// Marshal
	marshalJSON, err := controllerGenerator.scaffoldInterfaceControllerMarshalJSONFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&marshalJSON)
	f.Line()

	// Unmarshal
	unmarshalJSON, err := controllerGenerator.scaffoldInterfaceControllerUnmarshalJSONFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&unmarshalJSON)
	f.Line()

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
	interactorId , err := controllerGenerator.formatter.OutputUsecaseInteractorInterfaceId(entity)
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

func (controllerGenerator *controllerGenerator) interfaceControllerStruct(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Type
	resp.Type()

	// ID
	id , err := controllerGenerator.formatter.OutputInterfaceControllerStructId("http", entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Scaffold
	scaffoldId , err := controllerGenerator.formatter.OutputScaffoldInterfaceControllerStructId("http", entity)
	if err != nil {
		return nil, err
	}

	// Fields
	fields = append(fields, jen.Id(scaffoldId))


	// Struct
	resp.Struct(fields...)

	return resp, nil

}

func (controllerGenerator *controllerGenerator) interfaceControllerInterface(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Type
	resp.Type()

	// ID
	id , err := controllerGenerator.formatter.OutputInterfaceControllerInterfaceId("http", entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Scaffold
	scaffoldId , err := controllerGenerator.formatter.OutputScaffoldInterfaceControllerInterfaceId("http", entity)
	if err != nil {
		return nil, err
	}

	// Fields
	fields = append(fields, jen.Id(scaffoldId))


	// Interface
	resp.Interface(fields...)

	return resp, nil
	
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

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerInterface(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	var methods []jen.Code

	// Loop
	for _, method := range entity.Methods {
		
		// Method
		statement, err := controllerGenerator.scaffoldInterfaceControllerInterfaceMethod(method, entity)
		if err != nil {
			return nil, err
		}

		// Append
		methods = append(methods, &statement)
	}

	// Type
	resp.Type()

	// ID
	id , err := controllerGenerator.formatter.OutputScaffoldInterfaceControllerInterfaceId("http", entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Interface
	resp.Interface(methods...)

	return resp, nil

}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerEntityStruct(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// Type
	resp.Type()

	// ID
	id , err := controllerGenerator.formatter.OutputDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Interface ID
	interfaceId , err := controllerGenerator.formatter.OutputDomainEntityInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	// Import Path
	importPath , err := controllerGenerator.formatter.OutputScaffoldDomainEntityDirectoryImportPath()
	if err != nil {
		return nil, err
	}

	// Struct
	resp.Struct(
		jen.Qual(importPath, interfaceId),
	)

	
	return resp, nil

}

func (controllerGenerator *controllerGenerator) interfaceControllerConstructorFunction(entity model.Entity) (jen.Statement, error){

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
	structId , err := controllerGenerator.formatter.OutputInterfaceControllerStructId("http", entity)
	if err != nil {
		return nil, err
	}

	// Scaffold Struct ID
	scaffoldStructId , err := controllerGenerator.formatter.OutputScaffoldInterfaceControllerStructId("http", entity)
	if err != nil {
		return nil, err
	}

	// Interface ID
	interfaceId , err := controllerGenerator.formatter.OutputInterfaceControllerInterfaceId("http", entity)
	if err != nil {
		return nil, err
	}

	resp.Id(id)

	// Params
	resp.Params(fields...)

	// Qual
	resp.Qual("", interfaceId)

	// Block
	resp.Block(
		jen.Return(
			jen.Op("&").
			Id(structId).
			Values(
				jen.Id(scaffoldStructId).
				Values(
					jen.Id(interactorPackageName),
				),
			),
		),
	)
	
	return resp, nil

}



func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerInterfaceMethod(method model.Method, entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// ID
	id , err := controllerGenerator.formatter.OutputScaffoldUsecaseInteractorInterfaceMethodId(method)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Params
	resp.Params(
		jen.List(
			jen.Id("w").
			Qual("net/http", "ResponseWriter"),
			jen.Id("r").
			Op("*").
			Qual("net/http", "Request"),
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
			Op("*").
			Qual("net/http", "Request"),
		),
	)

	// Block
	var block []jen.Code
	switch method.Type {
	case "browse":
		block, err = controllerGenerator.scaffoldInterfaceControllerBrowseMethodBlock(method, entity)
	case "read":
		block, err = controllerGenerator.scaffoldInterfaceControllerReadMethodBlock(method, entity)
	case "edit":
		block, err = controllerGenerator.scaffoldInterfaceControllerEditMethodBlock(method, entity)	
	case "add":
		block, err = controllerGenerator.scaffoldInterfaceControllerAddMethodBlock(method, entity)
	case "delete":
		block, err = controllerGenerator.scaffoldInterfaceControllerDeleteMethodBlock(method, entity)
	case "count":
		block, err = controllerGenerator.scaffoldInterfaceControllerCountMethodBlock(method, entity)

	default:
		
	}

	resp.Block(block...)

	return resp, nil
}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerBrowseMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	var block []jen.Code

	// ID
	jsonId, err := controllerGenerator.formatter.OutputDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// ID
	jsonSliceId, err := controllerGenerator.formatter.OutputDomainEntitySliceStructId(entity)
	if err != nil {
		return nil, err
	}

	// Interface ID
	id , err := controllerGenerator.formatter.OutputDomainEntitySliceInterfaceConstructorFunctionId(entity)
	if err != nil {
		return block, err
	}

	// Import Path
	importPath , err := controllerGenerator.formatter.OutputScaffoldDomainEntityDirectoryImportPath()
	if err != nil {
		return block, err
	}


	interactorPackageName , err := controllerGenerator.formatter.OutputScaffoldUsecaseInteractorPackageName()
	if err != nil {
		return nil, err
	}

	// Repository Method ID
	repositoryMethodId , err := controllerGenerator.formatter.OutputScaffoldUsecaseRepositoryInterfaceMethodId(method)
	if err != nil {
		return nil, err
	}

	// Context
	block = append(block,

		jen.Id("ctx").
		Op(":=").
		Qual("context", "Background").
		Call(),

		jen.Id("req").
		Op(":=").
		Qual(importPath, id).
		Call(),

		jen.List(
			jen.Id("resp"),
			jen.Err(),
		).
		Op(":=").
		Id("c").
		Dot(interactorPackageName).
		Dot(repositoryMethodId).
		Call(
			jen.List(
				jen.Id("ctx"),
				jen.Nil(),
				jen.Id("req"),
			),
		),

		jen.If(
			jen.Err().
			Op("!=").
			Nil().
			Block(),
		),

		jen.Var().
		Id(jsonSliceId).
		Index().
		Op("*").
		Id(jsonId),

		jen.For(
			jen.List(
				jen.Id("_"),
				jen.Id("elem"),
			).
			Op(":=").
			Range().
			Id("resp").
			Dot("Elements").
			Call().
			Block(
				jen.Id(jsonSliceId).
				Op("=").
				Append(
					jen.Id(jsonSliceId),
					jen.Op("&").
					Id(jsonId).
					Values(
						jen.Id("elem"),
					),
				),
			),
		),

		jen.Id("w").
		Dot("Header").
		Call().
		Dot("Set").
		Params(
			jen.List(
				jen.Lit("Content-Type"),
				jen.Lit("application/json"),
			),
		),

		jen.Id("w").
		Dot("WriteHeader").
		Call(
			jen.Qual("net/http", "StatusOK"),
		),
		jen.Qual("encoding/json", "NewEncoder").
		Call(
			jen.Id("w"),
		).
		Dot("Encode").
		Call(
			jen.Id(jsonSliceId),
		),
	)

	return block, nil
}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerReadMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	
	var block []jen.Code
	
	// Interface ID
	id , err := controllerGenerator.formatter.OutputDomainEntityInterfaceConstructorFunctionId(entity)
	if err != nil {
		return block, err
	}

	// Import Path
	importPath , err := controllerGenerator.formatter.OutputScaffoldDomainEntityDirectoryImportPath()
	if err != nil {
		return block, err
	}


	interactorPackageName , err := controllerGenerator.formatter.OutputScaffoldUsecaseInteractorPackageName()
	if err != nil {
		return nil, err
	}

	// Repository Method ID
	repositoryMethodId , err := controllerGenerator.formatter.OutputScaffoldUsecaseRepositoryInterfaceMethodId(method)
	if err != nil {
		return nil, err
	}

	// Context
	block = append(block,

		jen.Id("ctx").
		Op(":=").
		Qual("context", "Background").
		Call(),

		jen.Id("req").
		Op(":=").
		Qual(importPath, id).
		Call(),

		jen.List(
			jen.Id("resp"),
			jen.Err(),
		).
		Op(":=").
		Id("c").
		Dot(interactorPackageName).
		Dot(repositoryMethodId).
		Call(
			jen.List(
				jen.Id("ctx"),
				jen.Nil(),
				jen.Id("req"),
			),
		),

		jen.If(
			jen.Err().
			Op("!=").
			Nil().
			Block(
				jen.Id("w").
				Dot("WriteHeader").
				Call(
					jen.Qual("net/http", "StatusBadRequest"),
				),

				jen.Qual("encoding/json", "NewEncoder").
				Call(
					jen.Id("w"),
				).
				Dot("Encode").
				Call(
					jen.Err().
					Dot("Error").
					Call(),
				),
			),
		),

		

		jen.Id("w").
		Dot("Header").
		Call().
		Dot("Set").
		Params(
			jen.List(
				jen.Lit("Content-Type"),
				jen.Lit("application/json"),
			),
		),

		jen.Id("w").
		Dot("WriteHeader").
		Call(
			jen.Qual("net/http", "StatusOK"),
		),
		jen.Qual("encoding/json", "NewEncoder").
		Call(
			jen.Id("w"),
		).
		Dot("Encode").
		Call(
			jen.Id("resp"),
		),
	)
	return block, nil
}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerEditMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	var block []jen.Code
	
	return block, nil
}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerAddMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	var block []jen.Code
	
	// ID
	jsonId, err := controllerGenerator.formatter.OutputDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// Interface ID
	id , err := controllerGenerator.formatter.OutputDomainEntityInterfaceConstructorFunctionId(entity)
	if err != nil {
		return block, err
	}

	// Import Path
	importPath , err := controllerGenerator.formatter.OutputScaffoldDomainEntityDirectoryImportPath()
	if err != nil {
		return block, err
	}


	interactorPackageName , err := controllerGenerator.formatter.OutputScaffoldUsecaseInteractorPackageName()
	if err != nil {
		return nil, err
	}

	// Repository Method ID
	repositoryMethodId , err := controllerGenerator.formatter.OutputScaffoldUsecaseRepositoryInterfaceMethodId(method)
	if err != nil {
		return nil, err
	}

	// Context
	block = append(block,

		jen.Id("ctx").
		Op(":=").
		Qual("context", "Background").
		Call(),

		jen.Id(jsonId).
		Op(":=").
		Op("&").
		Id(jsonId).
		Values(
			jen.Qual(importPath, id).
			Call(),
		),

		jen.Err().
		Op(":=").
		Qual("encoding/json", "NewDecoder").
		Call(
			jen.Id("r").
			Dot("Body"),
		).
		Dot("Decode").
		Call(
			jen.Id(jsonId),
		),

		jen.If(
			jen.Err().
			Op("!=").
			Nil().
			Block(
				jen.Id("w").
				Dot("WriteHeader").
				Call(
					jen.Qual("net/http", "StatusBadRequest"),
				),

				jen.Qual("encoding/json", "NewEncoder").
				Call(
					jen.Id("w"),
				).
				Dot("Encode").
				Call(
					jen.Err().
					Dot("Error").
					Call(),
				),
			),
		),

		jen.List(
			jen.Id("resp"),
			jen.Err(),
		).
		Op(":=").
		Id("c").
		Dot(interactorPackageName).
		Dot(repositoryMethodId).
		Call(
			jen.List(
				jen.Id("ctx"),
				jen.Id(jsonId),
			),
		),

		jen.If(
			jen.Err().
			Op("!=").
			Nil().
			Block(
				jen.Id("w").
				Dot("WriteHeader").
				Call(
					jen.Qual("net/http", "StatusBadRequest"),
				),

				jen.Qual("encoding/json", "NewEncoder").
				Call(
					jen.Id("w"),
				).
				Dot("Encode").
				Call(
					jen.Err().
					Dot("Error").
					Call(),
				),
			),
		),

		jen.Id("w").
		Dot("Header").
		Call().
		Dot("Set").
		Params(
			jen.List(
				jen.Lit("Content-Type"),
				jen.Lit("application/json"),
			),
		),

		jen.Id("w").
		Dot("WriteHeader").
		Call(
			jen.Qual("net/http", "StatusOK"),
		),
		jen.Qual("encoding/json", "NewEncoder").
		Call(
			jen.Id("w"),
		).
		Dot("Encode").
		Call(
			jen.Id("resp"),
		),
	)
	return block, nil
}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerDeleteMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	var block []jen.Code
	
	return block, nil
}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerCountMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	var block []jen.Code
	
	return block, nil
}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerMarshalJSONFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement

	// Struct ID
	structId , err := controllerGenerator.formatter.OutputDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// Func
	statement.Func()

	// Params
	statement.Params(
		jen.Id("m").
		Op("*").
		Qual("", structId),
	)

	// ID
	statement.Id("MarshalJSON")

	// Params
	statement.Params()

	// Parens
	statement.Parens(
		jen.List(
			jen.Index().
			Byte(),
			jen.Error(),
		),
	)	

	var jsonStruct []jen.Code
	jsonStructDict := make(jen.Dict)
	for _, field := range entity.Fields {

		// Getter ID
		getterId , err := controllerGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
		if err != nil {
			return nil, err
		}
		
		code, err := controllerGenerator.scaffoldEntityJSONStructField(field, entity)
		if err != nil {
			return nil, err
		}
		
		jsonStruct = append(jsonStruct, code)
		jsonStructDict[jen.Id(getterId)] = jen.Id("m").Dot(getterId).Call()
		
	}

	// Block
	statement.Block(

		jen.Type().Id("jsonStructPrivate").Struct(jsonStruct...),

		jen.Id("jsonStruct").
		Op(":=").
		Qual("", "jsonStructPrivate").
		Values(jsonStructDict),

		jen.Return().
		Qual("encoding/json", "Marshal").
		Call(
			jen.Op("&").
			Id("jsonStruct"),
		),
	)
	
	
	return statement, nil
}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerUnmarshalJSONFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement

	// Struct ID
	structId , err := controllerGenerator.formatter.OutputDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// Func
	statement.Func()

	// Params
	statement.Params(
		jen.Id("m").
		Op("*").
		Qual("", structId),
	)

	// ID
	statement.Id("UnmarshalJSON")

	// Params
	statement.Params(
		jen.Id("data").
		Index().
		Byte(),
	)

	// Parens
	statement.Parens(
		jen.List(
			jen.Error(),
		),
	)	

	var jsonStruct []jen.Code
	var jsonSetterFunctions []jen.Code
	for _, field := range entity.Fields {

		// Getter ID
		getterId , err := controllerGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
		if err != nil {
			return nil, err
		}

		// Setter ID
		setterId , err := controllerGenerator.formatter.OutputScaffoldDomainEntitySetterId(field)
		if err != nil {
			return nil, err
		}
		
		code, err := controllerGenerator.scaffoldEntityJSONStructField(field, entity)
		if err != nil {
			return nil, err
		}

		jsonStruct = append(jsonStruct, code)

		jsonSetterFunctions = append(jsonSetterFunctions, jen.Id("m").
			Dot(setterId).
			Call(
				jen.Id("jsonStruct").
				Dot(getterId),
			),
		)
		
	}

	// Block
	var block []jen.Code


	block = append(block, 
		jen.Type().
		Id("jsonStructPrivate").
		Struct(jsonStruct...),
	)
	block = append(block, 
		jen.Id("jsonStruct").
		Op(":=").
		Qual("", "jsonStructPrivate").
		Values(),
	)
	block = append(block, 
		jen.Err().
		Op(":=").
		Qual("encoding/json", "Unmarshal").
		Call(
			jen.Id("data"),
			jen.Op("&").
			Id("jsonStruct"),
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
					jen.Err(),
				),
			),
		),
	)
	for _, jsonSetterFunction := range jsonSetterFunctions {
		block = append(block, jsonSetterFunction)
	}
	block = append(block, jen.Return(
		jen.List(
			jen.Nil(),
		),
	),)

	statement.Block(
		block...,
	)
	
	
	return statement, nil
}

func (controllerGenerator *controllerGenerator) scaffoldEntityJSONStructField(field model.Field, entity model.Entity) (jen.Code, error){
	
	// ID
	var statement jen.Statement
	id , err := controllerGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
	if err != nil {
		return nil, err
	}
	tagId , err := controllerGenerator.formatter.OutputScaffoldDomainEntityJSONTagId(field)
	if err != nil {
		return nil, err
	}

	// Set
	statement.Id(id)

	// Field
	err = controllerGenerator.helperGenerator.Field("", field, entity, &statement)
	if err != nil {
		return nil, err
	}

	// Tag
	statement.Tag(map[string]string{"json": strings.Join([]string{tagId, "omitempty"}, ",")})

	return &statement, nil
}