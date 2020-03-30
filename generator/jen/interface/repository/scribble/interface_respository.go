package scribble

import(
	"github.com/dave/jennifer/jen"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/generator/jen/helper"
)

type repositoryGenerator struct{
	config *config.Config
	formatter formatter.Formatter
	helperGenerator helper.Generator
}

func New(config *config.Config, formatter formatter.Formatter, helperGenerator helper.Generator) *repositoryGenerator {
	return &repositoryGenerator{
		config : config,
		formatter : formatter,
		helperGenerator : helperGenerator,
	}
}

func (repositoryGenerator *repositoryGenerator) File(entity model.Entity) (*jen.File, error){

	// File
	packageName , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	// Struct
	interfaceRepositoryStruct, err := repositoryGenerator.interfaceRepositoryStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceRepositoryStruct)

	// Constructor Function
	interfaceRepositoryConstructorFunction, err := repositoryGenerator.interfaceRepositoryConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceRepositoryConstructorFunction)

	return f, nil
}

func (repositoryGenerator *repositoryGenerator) ScaffoldFile(entity model.Entity) (*jen.File, error){

	// File
	packageName , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	// Struct
	interfaceRepositoryStruct, err := repositoryGenerator.scaffoldInterfaceRepositoryStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceRepositoryStruct)

	// // Constructor Function
	// interfaceRepositoryConstructorFunction, err := repositoryGenerator.scaffoldInterfaceRepositoryConstructorFunction(entity)
	// if err != nil {
	// 	return nil, err
	// }
	// f.Add(&interfaceRepositoryConstructorFunction)

	// Methods
	for _, entityMethod := range entity.Methods {

		// Function
		method, err := repositoryGenerator.scaffoldInterfaceRepositoryMethod(entityMethod, entity)
		if err != nil {
			return nil, err
		}
		f.Add(&method)

	}

	return f, nil
}	

func (repositoryGenerator *repositoryGenerator) interfaceRepositoryStruct(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Type
	resp.Type()

	// ID
	id , err := repositoryGenerator.formatter.OutputInterfaceRepositoryStructId("scribble", entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Scaffold
	scaffoldId , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryStructId("scribble", entity)
	if err != nil {
		return nil, err
	}
	
	// Fields
	fields = append(fields, jen.Id(scaffoldId))


	// Struct
	resp.Struct(fields...)

	return resp, nil

}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryStructFields() ([]jen.Code, error){

	// Vars
	var fields []jen.Code

	// Driver
	fields = append(fields,
		jen.Id("driver").
		Op("*").
		Qual("github.com/nanobox-io/golang-scribble", "Driver"),
	)

	// Collection
	fields = append(fields,
		jen.Id("collection").
		String(),
	)

	return fields, nil

}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryStruct(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// Fields
	fields, err := repositoryGenerator.scaffoldInterfaceRepositoryStructFields()
	if err != nil {
		return nil, err
	}
	// Type
	resp.Type()

	// ID
	id , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryStructId("scribble", entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Struct
	resp.Struct(fields...)

	
	return resp, nil

}

func (repositoryGenerator *repositoryGenerator) interfaceRepositoryConstructorFunction(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// Type
	resp.Func()

	// Fields
	fields, err := repositoryGenerator.scaffoldInterfaceRepositoryStructFields()
	if err != nil {
		return nil, err
	}

	// ID
	id , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryConstructorFunctionId("scribble", entity)
	if err != nil {
		return nil, err
	}

	// Struct ID
	structId , err := repositoryGenerator.formatter.OutputInterfaceRepositoryStructId("scribble", entity)
	if err != nil {
		return nil, err
	}

	// Scaffold Struct ID
	scaffoldStructId , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryStructId("scribble", entity)
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
				jen.Id(scaffoldStructId).
				Values(
					jen.Dict{
						jen.Id("driver"):	jen.Id("driver"),
						jen.Id("collection"):	jen.Id("collection"),
					},
				),
			),
		),
	)
	
	return resp, nil

}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryMethod(method model.Method, entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	configMethod, err := repositoryGenerator.config.Method(method)

	var arguments, returnValues []jen.Code
	for _, argument := range configMethod.Repository.Arguments {

		var statement jen.Statement

		// Field
		err = repositoryGenerator.helperGenerator.Field(argument.Name, argument, entity, &statement)
		if err != nil {
			return nil, err
		}

		// Append
		arguments = append(arguments, &statement)

	}

	for _, returnValue := range configMethod.Repository.ReturnValues {

		var statement jen.Statement

		// Field
		err = repositoryGenerator.helperGenerator.Field("", returnValue, entity, &statement)
		if err != nil {
			return nil, err
		}

		// Append
		returnValues = append(returnValues, &statement)

	}

	// Type
	resp.Func()

	// Struct ID
	structId , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryStructId("scribble", entity)
	if err != nil {
		return nil, err
	}
	resp.Params(
		jen.Id("r").
		Op("*").
		Qual("", structId),
	)

	// ID
	id , err := repositoryGenerator.formatter.OutputScaffoldUsecaseRepositoryInterfaceMethodId(method)
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
	var block []jen.Code
	switch method.Type {
	case "browse":
		block, err = repositoryGenerator.scaffoldInterfaceRepositoryBrowseMethodBlock(method, entity)
	case "read":
		block, err = repositoryGenerator.scaffoldInterfaceRepositoryReadMethodBlock(method, entity)
	case "edit":
		block, err = repositoryGenerator.scaffoldInterfaceRepositoryEditMethodBlock(method, entity)	
	case "add":
		block, err = repositoryGenerator.scaffoldInterfaceRepositoryAddMethodBlock(method, entity)
	case "delete":
		block, err = repositoryGenerator.scaffoldInterfaceRepositoryDeleteMethodBlock(method, entity)
	case "count":
		block, err = repositoryGenerator.scaffoldInterfaceRepositoryCountMethodBlock(method, entity)

	default:
		
	}

	resp.Block(block...)

	return resp, nil
}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryBrowseMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	var block []jen.Code

	// Interface ID
	id , err := repositoryGenerator.formatter.OutputDomainEntityInterfaceConstructorFunctionId(entity)
	if err != nil {
		return block, err
	}

	// Import Path
	importPath , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityDirectoryImportPath()
	if err != nil {
		return block, err
	}

	block = append(block, 
		jen.List(
			jen.Id("records"),
			jen.Err(),
		).
		Op(":=").
		Id("r").
		Dot("driver").
		Dot("ReadAll").
		Call(
			jen.Id("r").
			Dot("collection"),
		),
		jen.If(
			jen.Err().
			Op("!=").
			Nil().
			Block(
				jen.Return(
					jen.Err(),
				),
			),
		),
		jen.For(
			jen.List(
				jen.Id("_"),
				jen.Id("record"),
			).
			Op(":=").
			Range().
			Id("records").
			Block(
				jen.
				Id("resp").
				Op(":=").
				Qual(importPath, id).
				Call(),
				jen.Err().
				Op(":=").
				Qual("encoding/json", "Unmarshal").
				Call(
					jen.Index().
					Byte().
					Call(
						jen.Id("record"),
					),
					jen.Id("resp"),
				),
				jen.If(
					jen.Err().
					Op("!=").
					Nil().
					Block(
						jen.Return(
							jen.Err(),
						),
					),
				),
				jen.Id("req").
				Dot("Append").
				Call(
					jen.Id("resp"),
				),
			),
		),
		jen.Return(
			jen.Nil(),
		),
	)	
	return block, nil
}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryReadMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	var block []jen.Code
	block = append(block, 
		jen.Return(
			jen.Id("r").
			Dot("driver").
			Dot("Read").
			Call(
				jen.Id("r").Dot("collection"),
				jen.Id("req").
				Dot("Id").
				Call(),
				jen.Id("req"),
			),
		),
	)	
	return block, nil
}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryEditMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	var block []jen.Code
	block = append(block, 
		jen.Return(
			jen.Id("r").
			Dot("driver").
			Dot("Read").
			Call(
				jen.Id("r").Dot("collection"),
				jen.Id("req").
				Dot("Id").
				Call(),
				jen.Id("req"),
			),
		),
	)	
	return block, nil
}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryAddMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	
	var block []jen.Code
	block = append(block, 
		jen.Return(
			jen.Id("r").
			Dot("driver").
			Dot("Write").
			Call(
				jen.Id("r").Dot("collection"),
				jen.Id("req").
				Dot("Id").
				Call(),
				jen.Id("req"),
			),
		),
	)	

	return block, nil
}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryDeleteMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	var block []jen.Code
	block = append(block, 
		jen.Return(
			jen.Id("r").
			Dot("driver").
			Dot("Delete").
			Call(
				jen.Id("r").Dot("collection"),
				jen.Id("req").
				Dot("Id").
				Call(),
			),
		),
	)	
	return block, nil
}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryCountMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	var block []jen.Code

	block = append(block, 
		jen.List(
			jen.Id("records"),
			jen.Err(),
		).
		Op(":=").
		Id("r").
		Dot("driver").
		Dot("ReadAll").
		Call(
			jen.Id("r").
			Dot("collection"),
		),
		jen.If(
			jen.Err().
			Op("!=").
			Nil().
			Block(
				jen.Return(
					jen.Err(),
				),
			),
		),
		jen.Id("count").
		Op(":=").
		Len(
			jen.Id("records"),
		),
		jen.Id("req").
		Op("=").
		Op("&").

		Id("count"),
		jen.Return(
			jen.Nil(),
		),
	)	
	return block, nil
}