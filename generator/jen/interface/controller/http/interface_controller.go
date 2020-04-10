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

func (controllerGenerator *controllerGenerator) AppFile(entities []model.Entity) (*jen.File, error){

	// File
	packageName , err := controllerGenerator.formatter.OutputScaffoldInterfaceControllerPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)
	
	// Interface
	interfaceControllerInterface, err := controllerGenerator.interfaceAppControllerInterface(entities)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceControllerInterface)

	return f, nil
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

	// // Entity Struct
	// interfaceRepositoryEntityStruct, err := controllerGenerator.scaffoldInterfaceControllerEntityStruct(entity)
	// if err != nil {
	// 	return nil, err
	// }
	// f.Add(&interfaceRepositoryEntityStruct)


	// Methods
	for _, entityMethod := range entity.Methods {

		// Function
		method, err := controllerGenerator.scaffoldInterfaceControllerMethod(entityMethod, entity)
		if err != nil {
			return nil, err
		}
		f.Add(&method)
		f.Line()
	}

	return f, nil
}	

func (controllerGenerator *controllerGenerator) EntityFile(entity model.Entity) (*jen.File, error){

	// File
	packageName , err := controllerGenerator.formatter.OutputScaffoldDomainEntityPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	// Struct
	jsonStruct, err := controllerGenerator.scaffoldInterfaceControllerJSONStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&jsonStruct)
	f.Line()

	// Local Marshal
	localMarshalJSON, err := controllerGenerator.scaffoldInterfaceControllerLocalMarshalJSONFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&localMarshalJSON)
	f.Line()

	// Local Unmarshal
	localUnmarshalJSON, err := controllerGenerator.scaffoldInterfaceControllerLocalUnmarshalJSONFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&localUnmarshalJSON)
	f.Line()

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

	return f, nil
}

func (controllerGenerator *controllerGenerator) RegistryFile(entity model.Entity) (*jen.File, error){

	// File
	packageName , err := controllerGenerator.formatter.OutputRegistryPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)
	
	// Struct
	interfaceControllerStruct, err := controllerGenerator.scaffoldInterfaceControllerRegistryStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceControllerStruct)

	// Constructor Function
	interfaceControllerRegistryConstructorFunction, err := controllerGenerator.interfaceControllerRegistryConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceControllerRegistryConstructorFunction)

	// Local Constructor Function
	interfaceControllerRegistryLocalConstructorFunction, err := controllerGenerator.interfaceControllerRegistryLocalConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceControllerRegistryLocalConstructorFunction)


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

func (controllerGenerator *controllerGenerator) interfaceAppControllerInterface(entities []model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Type
	resp.Type()

	// ID
	id , err := controllerGenerator.formatter.OutputInterfaceControllerInterfaceId("http", model.Entity{Name : "app"})
	if err != nil {
		return nil, err
	}
	resp.Id(id)
	
	for _, entity := range entities {

		var hasMatchingController bool
		for _, controller := range entity.Controllers {
			if controller.Type == "http" {
				hasMatchingController = true
			}
		}

		if hasMatchingController {
			// ID
			var statement jen.Statement

			// Import Path
			importPath, err := controllerGenerator.formatter.OutputInterfaceControllerDirectoryImportPath("http", entity)
			if err != nil {
				return nil, err
			}

			// Constructor
			constructorId, err := controllerGenerator.formatter.OutputRegistryEntityControllerConstructorFunctionId("http", entity)
			if err != nil {
				return nil, err
			}

			// Interface
			interfaceId, err := controllerGenerator.formatter.OutputInterfaceControllerInterfaceId("http", entity)
			if err != nil {
				return nil, err
			}

			// Set
			statement.Id(constructorId)

			// Params
			statement.Params()

			// Field
			err = controllerGenerator.helperGenerator.Field("", model.Field{Package : importPath, Type : interfaceId}, entity, &statement)
			if err != nil {
				return nil, err
			}


			fields = append(fields, &statement)
		}
		

	}
	// // Scaffold
	// scaffoldId , err := controllerGenerator.formatter.OutputScaffoldInterfaceControllerInterfaceId("http", entity)
	// if err != nil {
	// 	return nil, err
	// }

	// // Fields
	// fields = append(fields, jen.Id(scaffoldId))


	// Interface
	resp.Interface(fields...)

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
	id , err := controllerGenerator.formatter.OutputDomainEntityLocalStructId(entity)
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
	if entity.Interface{
		resp.Struct(
			jen.Qual(importPath, interfaceId),
		)
	} else {
		resp.Struct(
			jen.Op("*").Qual(importPath, interfaceId),
		)
	}

	
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
	// jsonSliceId, err := controllerGenerator.formatter.OutputDomainEntitySliceStructId(entity)
	// if err != nil {
	// 	return nil, err
	// }

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

	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.Id("ctx").
		Op(":=").
		Qual("context", "Background").
		Call(),
	)

	block = append(block, 
		jen.Id("req").
		Op(":=").
		Qual(importPath, id).
		Call(),
	)

	// Line
	block = append(block, jen.Line())

	block = append(block, 
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
				jen.Id("req"),
			),
		),
	)

	block = append(block, 
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
				jen.Return(),
			),
		),
	)

	// Line
	block = append(block, jen.Line())

	// block = append(block, 
	// 	jen.Var().
	// 	Id(jsonSliceId).
	// 	Index().
	// 	Op("*").
	// 	Id(jsonId),
	// )

	block = append(block, 
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
	)

	block = append(block, 
		jen.Id("w").
		Dot("WriteHeader").
		Call(
			jen.Qual("net/http", "StatusOK"),
		),
	)

	if entity.Interface {
		// block = append(block, 
		// 	jen.For(
		// 		jen.List(
		// 			jen.Id("_"),
		// 			jen.Id("elem"),
		// 		).
		// 		Op(":=").
		// 		Range().
		// 		Id("resp").
		// 		Dot("Elements").
		// 		Call().
		// 		Block(
		// 			jen.Id(jsonSliceId).
		// 			Op("=").
		// 			Append(
		// 				jen.Id(jsonSliceId),
		// 				jen.Op("&").
		// 				Id("elem"),
		// 			),
		// 		),
		// 	),
		// )
		block = append(block, 
			jen.Qual("encoding/json", "NewEncoder").
			Call(
				jen.Id("w"),
			).
			Dot("Encode").
			Call(
				jen.Id("resp").
				Dot("Elements").
				Call(),
			),
		)
	} else {

		block = append(block, 
			jen.Qual("encoding/json", "NewEncoder").
			Call(
				jen.Id("w"),
			).
			Dot("Encode").
			Call(
				jen.Id("resp"),
			),
		)
		
	}
	

	// Line
	block = append(block, jen.Line())

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

	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.Id("ctx").
		Op(":=").
		Qual("context", "Background").
		Call(),
	)

	block = append(block,
		jen.Id("req").
		Op(":=").
		Qual(importPath, id).
		Call(),
	)

	// Line
	block = append(block, jen.Line())

	// Mux Vars
	block = append(block, 
		jen.Var().
		Id("stringId").
		String(),
	)

	block = append(block, 
		jen.List(
			jen.Id("vars"),
		).
		Op(":=").
		Qual("github.com/gorilla/mux", "Vars").
		Call(
			jen.Id("r"),
		),
	)

	block = append(block, 
		jen.If(
			jen.List(
				jen.Id("val"),
				jen.Id("ok"),
			).
			Op(":=").
			Id("vars").
			Index(
				jen.Lit("id"),
			),
			jen.Id("ok"),
		).
		Block(
			jen.Id("stringId").
			Op("=").
			Id("val"),
		),
	)

	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.List(
			jen.Id("id"),
			jen.Err(),
		).
		Op(":=").
		Id("req").
		Dot("ToPrimary").
		Params(
			jen.Id("ctx"),
			jen.Id("stringId"),
		),
	)

	block = append(block,
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
				jen.Return(),
			),
		),
	)

	// Line
	block = append(block, jen.Line())

	var callList []jen.Code
	callList = append(callList, 
		jen.Id("ctx"),
		jen.Id("id"),
		jen.Id("req"),
	)
	block = append(block, 
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
				callList...,
			),
		),
	)

	block = append(block,
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
				jen.Return(),
			),
		),
	)

	// Line
	block = append(block, jen.Line())

	block = append(block, 
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
	)

	block = append(block, 
		jen.Id("w").
		Dot("WriteHeader").
		Call(
			jen.Qual("net/http", "StatusOK"),
		),
	)

	block = append(block, 
		jen.Qual("encoding/json", "NewEncoder").
		Call(
			jen.Id("w"),
		).
		Dot("Encode").
		Call(
			jen.Op("&").
			Id("resp"),
		),
	)

	// Line
	block = append(block, jen.Line())

	return block, nil
}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerEditMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
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

	// ID
	interactorPackageName , err := controllerGenerator.formatter.OutputScaffoldUsecaseInteractorPackageName()
	if err != nil {
		return nil, err
	}

	// Repository Method ID
	repositoryMethodId , err := controllerGenerator.formatter.OutputScaffoldUsecaseRepositoryInterfaceMethodId(method)
	if err != nil {
		return nil, err
	}

	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.Id("ctx").
		Op(":=").
		Qual("context", "Background").
		Call(),
	)

	block = append(block,
		jen.Id("req").
		Op(":=").
		Qual(importPath, id).
		Call(),
	)

	// Line
	block = append(block, jen.Line())

	// Mux Vars
	block = append(block, 
		jen.Var().
		Id("stringId").
		String(),
	)

	block = append(block, 
		jen.List(
			jen.Id("vars"),
		).
		Op(":=").
		Qual("github.com/gorilla/mux", "Vars").
		Call(
			jen.Id("r"),
		),
	)

	block = append(block, 
		jen.If(
			jen.List(
				jen.Id("val"),
				jen.Id("ok"),
			).
			Op(":=").
			Id("vars").
			Index(
				jen.Lit("id"),
			),
			jen.Id("ok"),
		).
		Block(
			jen.Id("stringId").
			Op("=").
			Id("val"),
		),
	)

	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.List(
			jen.Id("id"),
			jen.Err(),
		).
		Op(":=").
		Id("req").
		Dot("ToPrimary").
		Params(
			jen.Id("ctx"),
			jen.Id("stringId"),
		),
	)

	block = append(block,
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
				jen.Return(),
			),
		),
	)

	// Line
	block = append(block, jen.Line())

	block = append(block,
		jen.Err().
		Op("=").
		Qual("encoding/json", "NewDecoder").
		Call(
			jen.Id("r").
			Dot("Body"),
		).
		Dot("Decode").
		Call(
			jen.Id("req"),
		),
	)

	block = append(block,
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
				jen.Return(),
			),
		),
	)

	// Line
	block = append(block, jen.Line())

	var callList []jen.Code
	callList = append(callList, 
		jen.Id("ctx"),
		jen.Id("id"),
		jen.Id("req"),
	)

	block = append(block, 
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
				callList...,
			),
		),
	)

	block = append(block,
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
				jen.Return(),
			),
		),
	)

	// Line
	block = append(block, jen.Line())

	block = append(block, 
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
	)

	block = append(block, 
		jen.Id("w").
		Dot("WriteHeader").
		Call(
			jen.Qual("net/http", "StatusOK"),
		),
	)

	block = append(block, 
		jen.Qual("encoding/json", "NewEncoder").
		Call(
			jen.Id("w"),
		).
		Dot("Encode").
		Call(
			jen.Op("&").
			Id("resp"),
		),
	)

	// Line
	block = append(block, jen.Line())
	
	return block, nil
}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerAddMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
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

	// Line
	block = append(block, jen.Line())

	block = append(block,
		jen.Id("ctx").
		Op(":=").
		Qual("context", "Background").
		Call(),
	)

	block = append(block,
		jen.Id("req").
		Op(":=").
		Qual(importPath, id).
		Call(),
	)

	// Line
	block = append(block, jen.Line())

	block = append(block,
		jen.Err().
		Op(":=").
		Qual("encoding/json", "NewDecoder").
		Call(
			jen.Id("r").
			Dot("Body"),
		).
		Dot("Decode").
		Call(
			jen.Id("req"),
		),
	)

	block = append(block,
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
				jen.Return(),
			),
		),
	)

	// Line
	block = append(block, jen.Line())

	var callList []jen.Code
	callList = append(callList, 
		jen.Id("ctx"),
		jen.Id("req"),
	)

	block = append(block,
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
				callList...,
			),
		),
	)

	block = append(block,
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
				jen.Return(),
			),
		),
	)

	// Line
	block = append(block, jen.Line())

	block = append(block,
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
	)

	block = append(block,
		jen.Id("w").
		Dot("WriteHeader").
		Call(
			jen.Qual("net/http", "StatusOK"),
		),
	)

	block = append(block,
		jen.Qual("encoding/json", "NewEncoder").
		Call(
			jen.Id("w"),
		).
		Dot("Encode").
		Call(
			jen.Op("&").
			Id("resp"),
		),
	)

	// Line
	block = append(block, jen.Line())

	return block, nil
}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerDeleteMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	
	var block []jen.Code

	// Interface ID

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

	// ID
	interactorPackageName , err := controllerGenerator.formatter.OutputScaffoldUsecaseInteractorPackageName()
	if err != nil {
		return nil, err
	}

	// Repository Method ID
	repositoryMethodId , err := controllerGenerator.formatter.OutputScaffoldUsecaseRepositoryInterfaceMethodId(method)
	if err != nil {
		return nil, err
	}

	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.Id("ctx").
		Op(":=").
		Qual("context", "Background").
		Call(),
	)

	block = append(block,
		jen.Id("req").
		Op(":=").
		Qual(importPath, id).
		Call(),
	)

	// Line
	block = append(block, jen.Line())

	// Mux Vars
	block = append(block, 
		jen.Var().
		Id("stringId").
		String(),
	)

	block = append(block, 
		jen.List(
			jen.Id("vars"),
		).
		Op(":=").
		Qual("github.com/gorilla/mux", "Vars").
		Call(
			jen.Id("r"),
		),
	)

	block = append(block, 
		jen.If(
			jen.List(
				jen.Id("val"),
				jen.Id("ok"),
			).
			Op(":=").
			Id("vars").
			Index(
				jen.Lit("id"),
			),
			jen.Id("ok"),
		).
		Block(
			jen.Id("stringId").
			Op("=").
			Id("val"),
		),
	)

	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.List(
			jen.Id("id"),
			jen.Err(),
		).
		Op(":=").
		Id("req").
		Dot("ToPrimary").
		Params(
			jen.Id("ctx"),
			jen.Id("stringId"),
		),
	)

	block = append(block,
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
				jen.Return(),
			),
		),
	)

	// Line
	block = append(block, jen.Line())

	var callList []jen.Code
	callList = append(callList, 
		jen.Id("ctx"),
		jen.Id("id"),
		jen.Id("req"),
	)

	block = append(block, 
		jen.List(
			jen.Id("_"),
			jen.Err(),
		).
		Op("=").
		Id("c").
		Dot(interactorPackageName).
		Dot(repositoryMethodId).
		Call(
			jen.List(
				callList...,
			),
		),
	)

	block = append(block,
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
				jen.Return(),
			),
		),
	)

	// Line
	block = append(block, jen.Line())

	block = append(block, 
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
	)

	block = append(block, 
		jen.Id("w").
		Dot("WriteHeader").
		Call(
			jen.Qual("net/http", "StatusOK"),
		),
	)

	// Line
	block = append(block, jen.Line())

	return block, nil
}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerCountMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	var block []jen.Code
	
	return block, nil
}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerJSONStruct(entity model.Entity) (jen.Statement, error){

	var statement jen.Statement
	var fields []jen.Code
	for _, field := range entity.Fields {
		if !field.Embedded {
			code, err := controllerGenerator.scaffoldEntityJSONStructField(field, entity)
			if err != nil {
				return nil, err
			}
			
			fields = append(fields, code)
		} else {

			// Struct ID
			structId , err := controllerGenerator.formatter.OutputScaffoldDomainEntityMarshalStructId("json", field.Entity)
			if err != nil {
				return nil, err
			}

			fields = append(fields, jen.Op("*").Id(structId))

		}
		
		
	}

	// Struct ID
	structId , err := controllerGenerator.formatter.OutputScaffoldDomainEntityMarshalStructId("json", entity)
	if err != nil {
		return nil, err
	}

	// Type
	statement.Type()

	// Id
	statement.Id(structId)

	// Struct
	statement.Struct(fields...)
	
	return statement, nil

}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerLocalMarshalJSONFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement

	// Struct ID
	structId , err := controllerGenerator.formatter.OutputScaffoldDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}
	
	// Marshal Struct ID
	marshalStructId , err := controllerGenerator.formatter.OutputScaffoldDomainEntityMarshalStructId("json", entity)
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
	statement.Id("marshalJSON")

	// Params
	statement.Params()

	// Parens
	statement.Parens(
		jen.List(
			jen.Op("*").
			Id(marshalStructId),
		),
	)	


	// Block
	var block []jen.Code

	block = append(block, jen.Line())

	block = append(block,
		jen.Id("jsonStruct").
		Op(":=").
		Id(marshalStructId).
		Values(),
	)

	block = append(block, jen.Line())

	for _, field := range entity.Fields {

		if !field.Embedded {

			// Getter ID
			getterId , err := controllerGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
			if err != nil {
				return nil, err
			}

			if entity.Interface {
				block = append(block,
					jen.Id("jsonStruct").
					Dot(getterId).
					Op("=").
					Id("m").Dot(getterId).Call(),
				)
			} else {
				block = append(block,
					jen.Id("jsonStruct").
					Dot(getterId).
					Op("=").
					Id("m").Dot(getterId),
				)
			}
		} else {

			// Vars
			var entityId string

			// Struct ID
			fieldId , err := controllerGenerator.formatter.OutputScaffoldDomainEntityMarshalStructId("json", field.Entity)
			if err != nil {
				return nil, err
			}

			// Interface ID
			if field.Entity.Interface {
				structId , err := controllerGenerator.formatter.OutputDomainEntityStructId(field.Entity)
				if err != nil {
					return nil, err
				}
				entityId = structId
			} else {
				interfaceId , err := controllerGenerator.formatter.OutputDomainEntityInterfaceId(field.Entity)
				if err != nil {
					return nil, err
				}
				entityId = interfaceId
			}

			block = append(block,
				jen.Id("jsonStruct").
				Dot(fieldId).
				Op("=").
				Id("m").
				Dot(entityId).
				Dot("marshalJSON").
				Call(),
			)

		}	
	}

	block = append(block, jen.Line())

	block = append(block,
		jen.Return(
			jen.Op("&").
			Id("jsonStruct"),
		),
	)
	
	statement.Block(
		block...,
	)
	
	
	return statement, nil
}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerMarshalJSONFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement

	// Struct ID
	structId , err := controllerGenerator.formatter.OutputScaffoldDomainEntityStructId(entity)
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


	// Block
	statement.Block(
		jen.Return(
			jen.Qual("encoding/json", "Marshal").
			Call(
				jen.Op("m").
				Dot("marshalJSON").
				Call(),
			),
		),
	)
	
	
	return statement, nil
}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerLocalUnmarshalJSONFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement

	// Struct ID
	structId , err := controllerGenerator.formatter.OutputScaffoldDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// Marshal Struct ID
	marshalStructId , err := controllerGenerator.formatter.OutputScaffoldDomainEntityMarshalStructId("json", entity)
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
	statement.Id("unmarshalJSON")

	// Params
	statement.Params(
		jen.Id("jsonStruct").
		Op("*").
		Qual("", marshalStructId),
	)

	// Block
	var block []jen.Code

	
	for _, field := range entity.Fields {

		if !field.Embedded {

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
			

			if entity.Interface {
				block = append(block, 
					jen.Id("m").
					Dot(setterId).
					Call(
						jen.Id("jsonStruct").
						Dot(getterId),
					),
				)
			} else {
				block = append(block, 
					jen.Id("m").
					Dot(getterId).
					Op("=").
					Id("jsonStruct").
					Dot(getterId),
				)
			}	
		} else {

			// Vars
			var entityId string

			// Struct ID
			fieldId , err := controllerGenerator.formatter.OutputScaffoldDomainEntityMarshalStructId("json", field.Entity)
			if err != nil {
				return nil, err
			}

			// Interface ID
			if field.Entity.Interface {
				structId , err := controllerGenerator.formatter.OutputDomainEntityStructId(field.Entity)
				if err != nil {
					return nil, err
				}
				entityId = structId
			} else {
				interfaceId , err := controllerGenerator.formatter.OutputDomainEntityInterfaceId(field.Entity)
				if err != nil {
					return nil, err
				}
				entityId = interfaceId
			}

			block = append(block,
				jen.Id("m").
				Dot(entityId).
				Dot("unmarshalJSON").
				Params(
					jen.Id("jsonStruct").
					Dot(fieldId),
				),
			)

			// m.model.unmarshalJSON(jsonStruct.jsonModel)

		} 
	}	
	// block = append(block, 
	// 	jen.Type().
	// 	Id("jsonStructPrivate").
	// 	Struct(jsonStruct...),
	// )
	// block = append(block, 
	// 	jen.Id("jsonStruct").
	// 	Op(":=").
	// 	Qual("", "jsonStructPrivate").
	// 	Values(),
	// )
	// block = append(block, 
	// 	jen.Err().
	// 	Op(":=").
	// 	Qual("encoding/json", "Unmarshal").
	// 	Call(
	// 		jen.Id("data"),
	// 		jen.Op("&").
	// 		Id("jsonStruct"),
	// 	),
	// )
	// block = append(block, 
	// 	jen.If(
	// 		jen.Err().
	// 		Op("!=").
	// 		Nil(),
	// 	).
	// 	Block(
	// 		jen.Return(
	// 			jen.List(
	// 				jen.Err(),
	// 			),
	// 		),
	// 	),
	// )
	// for _, jsonSetterFunction := range jsonSetterFunctions {
	// 	block = append(block, jsonSetterFunction)
	// }
	// block = append(block, jen.Return(
	// 	jen.List(
	// 		jen.Nil(),
	// 	),
	// ),)

	statement.Block(
		block...,
	)
	
	
	return statement, nil
}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerUnmarshalJSONFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement

	// Struct ID
	structId , err := controllerGenerator.formatter.OutputScaffoldDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// Marshal Struct ID
	marshalStructId , err := controllerGenerator.formatter.OutputScaffoldDomainEntityMarshalStructId("json", entity)
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

	// Block
	var block []jen.Code

	block = append(block, jen.Line())

	block = append(block, 
		jen.Id("jsonStruct").
		Op(":=").
		Qual("", marshalStructId).
		Values(),
	)

	block = append(block, jen.Line())

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
	
	block = append(block, jen.Line())

	block = append(block, 
		jen.Id("m").
		Dot("unmarshalJSON").
		Params(
			jen.Op("&").
			Id("jsonStruct"),
		),
	)

	block = append(block, jen.Line())

	block = append(block, 
		jen.Return(
			jen.List(
				jen.Nil(),
			),
		),
	)

	block = append(block, jen.Line())

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
	field.Op = "*"
	err = controllerGenerator.helperGenerator.Field("", field, entity, &statement)
	if err != nil {
		return nil, err
	}

	// Tag
	statement.Tag(map[string]string{"json": strings.Join([]string{tagId, "omitempty"}, ",")})

	return &statement, nil
}

func (controllerGenerator *controllerGenerator) scaffoldInterfaceControllerRegistryStruct(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// Fields
	var fields []jen.Code

	// Type
	resp.Type()

	// ID
	id , err := controllerGenerator.formatter.OutputScaffoldInterfaceControllerRegistryStructId("http", entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Struct
	resp.Struct(fields...)

	
	return resp, nil

}

func (controllerGenerator *controllerGenerator) interfaceControllerRegistryConstructorFunction(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// Registry
	registryName , err := controllerGenerator.formatter.OutputRegistryPackageName()
	if err != nil {
		return nil, err
	}

	// ID
	// registryStructId , err := controllerGenerator.formatter.OutputScaffoldInterfaceControllerRegistryStructId("http", entity)
	// if err != nil {
	// 	return nil, err
	// }

	// Type
	resp.Func()

	resp.Params(
		jen.Id("r").
		Op("*").
		Qual("", registryName),
	)

	// Fields
	var fields []jen.Code

	// ID
	id , err := controllerGenerator.formatter.OutputScaffoldInterfaceControllerRegistryConstructorFunctionId("http", entity)
	if err != nil {
		return nil, err
	}


	resp.Id(id)

	// Params
	resp.Params(fields...)

	// Block
	resp.Block()
	
	return resp, nil

}

func (controllerGenerator *controllerGenerator) interfaceControllerRegistryLocalConstructorFunction(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// Registry
	registryName , err := controllerGenerator.formatter.OutputRegistryPackageName()
	if err != nil {
		return nil, err
	}

	// // ID
	// registryStructId , err := controllerGenerator.formatter.OutputScaffoldInterfaceControllerRegistryStructId("http", entity)
	// if err != nil {
	// 	return nil, err
	// }

	interfaceId , err := controllerGenerator.formatter.OutputInterfaceControllerInterfaceId("http", entity)
	if err != nil {
		return nil, err
	}

	// interfaceImportPath , err := controllerGenerator.formatter.OutputInterfaceControllerDirectoryImportPath("http", entity)
	// if err != nil {
	// 	return nil, err
	// }

	usecaseImportPath , err := controllerGenerator.formatter.OutputInterfaceControllerDirectoryImportPath("http", entity)
	if err != nil {
		return nil, err
	}
	
	// Type
	resp.Func()

	resp.Params(
		jen.Id("r").
		Op("*").
		Qual("", registryName),
	)

	// ID
	id , err := controllerGenerator.formatter.OutputScaffoldInterfaceControllerRegistryLocalConstructorFunctionId("http", entity)
	if err != nil {
		return nil, err
	}

	// Interactor Import Path
	interactorImportPath , err := controllerGenerator.formatter.OutputScaffoldUsecaseInteractorDirectoryImportPath()
	if err != nil {
		return nil, err
	}
	

	// Local Interactor Id
	interactorId , err := controllerGenerator.formatter.OutputUsecaseInteractorInterfaceConstructorFunctionId(entity)
	if err != nil {
		return nil, err
	}

	// Primary Repository Function
	var primaryRepositoryConstructorId string
	for _, repository := range entity.Repositories {
		if repository.Primary {

			// ID
			id, err := controllerGenerator.formatter.OutputScaffoldInterfaceRepositoryRegistryLocalConstructorFunctionId(repository.Type, entity)
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
			id, err := controllerGenerator.formatter.OutputScaffoldInterfacePresenterRegistryLocalConstructorFunctionId(presenter.Type, entity)
			if err != nil {
				return nil, err
			}
			primaryPresenterConstructorId = id

		}
	}
	

	resp.Id(id)

	// Params
	resp.Params()

	resp.Parens(
		jen.List(
			jen.Qual(usecaseImportPath, interfaceId),
		),
	)

	resp.Block(
		jen.Return(
			jen.Qual(usecaseImportPath, "New").
			Params(
				jen.Qual(interactorImportPath, interactorId).
				Params(
					
					jen.Id("r").
					Dot(primaryRepositoryConstructorId).
					Call(),

					jen.Id("r").
					Dot(primaryPresenterConstructorId).
					Call(),

				),
			),
		),
	)
	
	return resp, nil

}