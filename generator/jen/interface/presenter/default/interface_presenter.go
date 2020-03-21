package generator

import(
	"github.com/dave/jennifer/jen"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/generator/jen/helper"
)

type presenterGenerator struct{
	config *config.Config
	formatter formatter.Formatter
	helperGenerator helper.Generator
}

// type PresenterGenerator interface{

// 	ScaffoldFile(entity model.Entity) (*jen.File, error)

// 	scaffoldInterfacePresenterStruct(entity model.Entity) (jen.Statement, error)
// 	scaffoldInterfacePresenterInterface(entity model.Entity) (jen.Statement, error)
// 	scaffoldInterfacePresenterInterfaceMethod(method model.Method, entity model.Entity) (jen.Statement, error)
// 	scaffoldInterfacePresenterConstructorFunction(entity model.Entity) (jen.Statement, error)
// 	scaffoldInterfacePresenterMethod(method model.Method, entity model.Entity) (jen.Statement, error)

// }

func New(config *config.Config, formatter formatter.Formatter, helperGenerator helper.Generator) *presenterGenerator {
	return &presenterGenerator{
		config : config,
		formatter : formatter,
		helperGenerator :helperGenerator,
	}
}

func (presenterGenerator *presenterGenerator) File(entity model.Entity) (*jen.File, error){
	
	
	// File
	packageName , err := presenterGenerator.formatter.OutputScaffoldInterfacePresenterPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	// Struct
	interfacePresenterStruct, err := presenterGenerator.interfacePresenterStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfacePresenterStruct)

	// Constructor Function
	interfacePresenterConstructorFunction, err := presenterGenerator.interfacePresenterConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfacePresenterConstructorFunction)

	return f, nil
}	

func (presenterGenerator *presenterGenerator) ScaffoldFile(entity model.Entity) (*jen.File, error){
	
	
	// File
	packageName , err := presenterGenerator.formatter.OutputScaffoldInterfacePresenterPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	// Struct
	interfacePresenterStruct, err := presenterGenerator.scaffoldInterfacePresenterStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfacePresenterStruct)

	// Functions
	for _, method := range entity.Methods {

		// Function
		function, err := presenterGenerator.scaffoldInterfacePresenterMethod(method, entity)
		if err != nil {
			return nil, err
		}
		f.Add(&function)

	}

	return f, nil
}	

func (presenterGenerator *presenterGenerator) interfacePresenterStruct(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Type
	resp.Type()

	// ID
	id , err := presenterGenerator.formatter.OutputInterfacePresenterStructId("default", entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Scaffold
	scaffoldId , err := presenterGenerator.formatter.OutputScaffoldInterfacePresenterStructId("default", entity)
	if err != nil {
		return nil, err
	}
	
	// Fields
	fields = append(fields, jen.Id(scaffoldId))


	// Struct
	resp.Struct(fields...)

	return resp, nil

}

func (presenterGenerator *presenterGenerator) scaffoldInterfacePresenterStruct(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Type
	resp.Type()

	// ID
	id , err := presenterGenerator.formatter.OutputScaffoldInterfacePresenterStructId("default", entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)


	// Struct
	resp.Struct(fields...)

	return resp, nil

}

func (presenterGenerator *presenterGenerator) scaffoldInterfacePresenterInterface(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	var methods []jen.Code

	// Loop
	for _, method := range entity.Methods {
		
		// Method
		statement, err := presenterGenerator.scaffoldInterfacePresenterInterfaceMethod(method, entity)
		if err != nil {
			return nil, err
		}

		// Append
		methods = append(methods, &statement)
	}

	// Type
	resp.Type()

	// ID
	id , err := presenterGenerator.formatter.OutputScaffoldInterfacePresenterInterfaceId("default", entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Interface
	resp.Interface(methods...)

	return resp, nil

}

func (presenterGenerator *presenterGenerator) interfacePresenterConstructorFunction(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// Type
	resp.Func()

	// ID
	id , err := presenterGenerator.formatter.OutputScaffoldInterfacePresenterConstructorFunctionId("default", entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)


	// Struct ID
	structId , err := presenterGenerator.formatter.OutputScaffoldInterfacePresenterStructId("default", entity)
	if err != nil {
		return nil, err
	}

	// Params
	resp.Params()

	// Qual
	resp.Op("*")
	resp.Qual("", structId)

	// Block
	resp.Block(
		jen.Return(
			jen.Op("&").
			Id(structId).
			Values(),
		),
	)
	
	return resp, nil

}

func (presenterGenerator *presenterGenerator) scaffoldInterfacePresenterMethod(method model.Method, entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	configMethod, err := presenterGenerator.config.Method(method)

	var arguments, returnValues []jen.Code
	for _, argument := range configMethod.Presenter.Arguments {

		var statement jen.Statement

		// Field
		err = presenterGenerator.helperGenerator.Field(argument.Name, argument, entity, &statement)
		if err != nil {
			return nil, err
		}

		// Append
		arguments = append(arguments, &statement)

	}

	for _, returnValue := range configMethod.Presenter.ReturnValues {

		var statement jen.Statement

		// Field
		err = presenterGenerator.helperGenerator.Field("", returnValue, entity, &statement)
		if err != nil {
			return nil, err
		}

		// Append
		returnValues = append(returnValues, &statement)

	}

	// Type
	resp.Func()

	// Struct ID
	structId , err := presenterGenerator.formatter.OutputScaffoldInterfacePresenterStructId("default", entity)
	if err != nil {
		return nil, err
	}
	resp.Params(
		jen.Id("r").
		Op("*").
		Qual("", structId),
	)

	// ID
	id , err := presenterGenerator.formatter.OutputScaffoldUsecasePresenterInterfaceMethodId(method)
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

	// Block
	resp.Block(
		jen.Return(
			jen.List(
				jen.Id("req"),
				jen.Nil(),
			),
		),
	)

	return resp, nil

}

func (presenterGenerator *presenterGenerator) scaffoldInterfacePresenterInterfaceMethod(method model.Method, entity model.Entity) (jen.Statement, error){
	
	// Vars
	var resp jen.Statement

	configMethod, err := presenterGenerator.config.Method(method)

	var arguments, returnValues []jen.Code
	for _, argument := range configMethod.Presenter.Arguments {

		var statement jen.Statement

		// Field
		err = presenterGenerator.helperGenerator.Field(argument.Name, argument, entity, &statement)
		if err != nil {
			return nil, err
		}

		// Append
		arguments = append(arguments, &statement)

	}

	for _, returnValue := range configMethod.Presenter.ReturnValues {

		var statement jen.Statement

		// Field
		err = presenterGenerator.helperGenerator.Field("", returnValue, entity, &statement)
		if err != nil {
			return nil, err
		}

		// Append
		returnValues = append(returnValues, &statement)

	}

	// ID
	id , err := presenterGenerator.formatter.OutputScaffoldUsecasePresenterInterfaceMethodId(method)
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