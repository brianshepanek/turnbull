package mongo

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

	// Entity Struct
	interfaceRepositoryEntityStruct, err := repositoryGenerator.scaffoldInterfaceRepositoryEntityStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceRepositoryEntityStruct)

	// Marshal
	marshalBSON, err := repositoryGenerator.scaffoldInterfaceRepositoryMarshalBSONFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&marshalBSON)
	f.Line()

	// Unmarshal
	unmarshalBSON, err := repositoryGenerator.scaffoldInterfaceRepositoryUnmarshalBSONFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&unmarshalBSON)
	f.Line()

	// Methods
	for _, entityMethod := range entity.Methods {

		// Function
		method, err := repositoryGenerator.scaffoldInterfaceRepositoryMethod(entityMethod, entity)
		if err != nil {
			return nil, err
		}
		f.Add(&method)
		f.Line()

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
	id , err := repositoryGenerator.formatter.OutputInterfaceRepositoryStructId("mongo", entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Scaffold
	scaffoldId , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryStructId("mongo", entity)
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
		jen.Id("client").
		Op("*").
		Qual("go.mongodb.org/mongo-driver/mongo", "Client"),
	)

	// DB
	fields = append(fields,
		jen.Id("db").
		String(),
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
	id , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryStructId("mongo", entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Struct
	resp.Struct(fields...)

	
	return resp, nil

}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryEntityStruct(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// Type
	resp.Type()

	// ID
	id , err := repositoryGenerator.formatter.OutputDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Interface ID
	interfaceId , err := repositoryGenerator.formatter.OutputDomainEntityInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	// Import Path
	importPath , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityDirectoryImportPath()
	if err != nil {
		return nil, err
	}

	// Struct
	resp.Struct(
		jen.Qual(importPath, interfaceId),
	)

	
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
	id , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryConstructorFunctionId("mongo", entity)
	if err != nil {
		return nil, err
	}

	// Struct ID
	structId , err := repositoryGenerator.formatter.OutputInterfaceRepositoryStructId("mongo", entity)
	if err != nil {
		return nil, err
	}

	// Scaffold Struct ID
	scaffoldStructId , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryStructId("mongo", entity)
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
						jen.Id("client"):	jen.Id("client"),
						jen.Id("db"):	jen.Id("db"),
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
	structId , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryStructId("mongo", entity)
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

	// Loca ID
	localID, err := repositoryGenerator.formatter.OutputDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}	

	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.Id("collection").
		Op(":=").
		Id("r").
		Dot("client").
		Dot("Database").
		Params(
			jen.Id("r").
			Dot("db"),
		).
		Dot("Collection").
		Params(
			jen.Id("r").
			Dot("collection"),
		),
	)	

	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.List(
			jen.Id("cursor"),
			jen.Err(),
		).
		Op(":=").
		Id("collection").
		Dot("Find").
		Params(
			jen.Id("ctx"),
			jen.Qual("go.mongodb.org/mongo-driver/bson", "D").
			Values(),
		),
	)	

	block = append(block,
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
	)	

	// Line
	block = append(block, jen.Line())

	block = append(block,
		jen.For(
			jen.Id("cursor").
			Dot("Next").
			Params(
				jen.Id("ctx"),
			).
			Block(
				jen.Id("elem").
				Op(":=").
				Op("&").
				Id(localID).
				Values(
					jen.Qual(importPath, id).
					Call(),
				),
				
				jen.Err().
				Op(":=").
				Id("cursor").
				Dot("Decode").
				Params(
					jen.Op("&").
					Id("elem"),
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
					jen.Id("elem"),
				),
			),	
		),
	)	

	// Line
	block = append(block, jen.Line())

	block = append(block,
		jen.Return(
			jen.Nil(),
		),
	)	

	// Line
	block = append(block, jen.Line())

	return block, nil
}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryReadMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	
	var block []jen.Code

	// ID
	id, err := repositoryGenerator.formatter.OutputDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// Primary Field Name
	var primaryFieldName string
	for _, field := range entity.Fields {
		if field.Primary {
			primaryFieldName , err = repositoryGenerator.formatter.OutputScaffoldDomainEntityJSONTagId(field)
			if err != nil {
				return nil, err
			}
		}
	}
	
	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.Id(id).
		Op(":=").
		Op("&").
		Id(id).
		Values(
			jen.Id("req"),
		),
	)	

	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.Id("collection").
		Op(":=").
		Id("r").
		Dot("client").
		Dot("Database").
		Params(
			jen.Id("r").
			Dot("db"),
		).
		Dot("Collection").
		Params(
			jen.Id("r").
			Dot("collection"),
		),
	)	

	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.List(
			jen.Id("filter"),
		).
		Op(":=").
		Qual("go.mongodb.org/mongo-driver/bson", "M").
		Values(
			jen.Dict{
				jen.Lit(primaryFieldName): jen.Id("id"),
			},
		),
	)	

	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.List(
			jen.Err(),
		).
		Op(":=").
		Id("collection").
		Dot("FindOne").
		Params(
			jen.Id("ctx"),
			jen.Id("filter"),
		).
		Dot("Decode").
		Params(
			jen.Id(id),
		),
	)	

	block = append(block,
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
	)	

	// Line
	block = append(block, jen.Line())


	block = append(block,
		jen.Return(
			jen.Nil(),
		),
	)

	// Line
	block = append(block, jen.Line())

	return block, nil
}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryEditMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	var block []jen.Code
	block = append(block,
		jen.Return(
			jen.Nil(),
		),
	)
	return block, nil
}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryAddMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	
	var block []jen.Code

	// ID
	id, err := repositoryGenerator.formatter.OutputDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.Id(id).
		Op(":=").
		Op("&").
		Id(id).
		Values(
			jen.Id("req"),
		),
	)	

	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.Id("collection").
		Op(":=").
		Id("r").
		Dot("client").
		Dot("Database").
		Params(
			jen.Id("r").
			Dot("db"),
		).
		Dot("Collection").
		Params(
			jen.Id("r").
			Dot("collection"),
		),
	)	

	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.List(
			jen.Id("_"),
			jen.Err(),
		).
		Op(":=").
		Id("collection").
		Dot("InsertOne").
		Params(
			jen.Id("ctx"),
			jen.Id(id),
		),
	)	
	
	block = append(block,
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
	)	

	// Line
	block = append(block, jen.Line())

	// block = append(block, 
	// 	jen.List(
	// 		jen.Id("filter"),
	// 	).
	// 	Op(":=").
	// 	Qual("go.mongodb.org/mongo-driver/bson", "M").
	// 	Values(
	// 		jen.Dict{
	// 			jen.Lit("_id"): jen.Id("res").Dot("InsertedID"),
	// 		},
	// 	),
	// )	

	// block = append(block, 
	// 	jen.List(
	// 		jen.Err(),
	// 	).
	// 	Op("=").
	// 	Id("collection").
	// 	Dot("FindOne").
	// 	Params(
	// 		jen.Id("ctx"),
	// 		jen.Id("filter"),
	// 	).
	// 	Dot("Decode").
	// 	Params(
	// 		jen.Id(id),
	// 	),
	// )	
	
	block = append(block,
		jen.Return(
			jen.Nil(),
		),
	)

	// Line
	block = append(block, jen.Line())

	return block, nil
}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryDeleteMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	var block []jen.Code

	// Primary Field Name
	var primaryFieldName string
	var err error
	for _, field := range entity.Fields {
		if field.Primary {
			primaryFieldName , err = repositoryGenerator.formatter.OutputScaffoldDomainEntityJSONTagId(field)
			if err != nil {
				return nil, err
			}
		}
	}
	
	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.Id("collection").
		Op(":=").
		Id("r").
		Dot("client").
		Dot("Database").
		Params(
			jen.Id("r").
			Dot("db"),
		).
		Dot("Collection").
		Params(
			jen.Id("r").
			Dot("collection"),
		),
	)	

	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.List(
			jen.Id("filter"),
		).
		Op(":=").
		Qual("go.mongodb.org/mongo-driver/bson", "M").
		Values(
			jen.Dict{
				jen.Lit(primaryFieldName): jen.Id("id"),
			},
		),
	)	

	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.List(
			jen.Id("_"),
			jen.Err(),
		).
		Op(":=").
		Id("collection").
		Dot("DeleteOne").
		Params(
			jen.Id("ctx"),
			jen.Id("filter"),
		),
	)	

	block = append(block,
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
	)	

	// Line
	block = append(block, jen.Line())


	block = append(block,
		jen.Return(
			jen.Nil(),
		),
	)

	// Line
	block = append(block, jen.Line())

	return block, nil
}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryCountMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	var block []jen.Code

	block = append(block,
		jen.Return(
			jen.Nil(),
		),
	)
	return block, nil
}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryMarshalBSONFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement

	// Struct ID
	structId , err := repositoryGenerator.formatter.OutputDomainEntityStructId(entity)
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
	statement.Id("MarshalBSON")

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

	var bsonStruct []jen.Code
	bsonStructDict := make(jen.Dict)
	for _, field := range entity.Fields {

		// Getter ID
		getterId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
		if err != nil {
			return nil, err
		}
		
		code, err := repositoryGenerator.scaffoldEntityBSONStructField(field, entity)
		if err != nil {
			return nil, err
		}
		
		bsonStruct = append(bsonStruct, code)
		bsonStructDict[jen.Id(getterId)] = jen.Id("m").Dot(getterId).Call()
		
	}

	// Block
	statement.Block(

		jen.Type().Id("bsonStructPrivate").Struct(bsonStruct...),

		jen.Id("bsonStruct").
		Op(":=").
		Qual("", "bsonStructPrivate").
		Values(bsonStructDict),

		jen.Return().
		Qual("go.mongodb.org/mongo-driver/bson", "Marshal").
		Call(
			jen.Op("&").
			Id("bsonStruct"),
		),
	)
	
	
	return statement, nil
}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryUnmarshalBSONFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement

	// Struct ID
	structId , err := repositoryGenerator.formatter.OutputDomainEntityStructId(entity)
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
	statement.Id("UnmarshalBSON")

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

	var bsonStruct []jen.Code
	var bsonSetterFunctions []jen.Code
	for _, field := range entity.Fields {

		// Getter ID
		getterId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
		if err != nil {
			return nil, err
		}

		// Setter ID
		setterId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntitySetterId(field)
		if err != nil {
			return nil, err
		}
		
		code, err := repositoryGenerator.scaffoldEntityBSONStructField(field, entity)
		if err != nil {
			return nil, err
		}

		bsonStruct = append(bsonStruct, code)

		bsonSetterFunctions = append(bsonSetterFunctions, jen.Id("m").
			Dot(setterId).
			Call(
				jen.Id("bsonStruct").
				Dot(getterId),
			),
		)
		
	}

	// Block
	var block []jen.Code


	block = append(block, 
		jen.Type().
		Id("bsonStructPrivate").
		Struct(bsonStruct...),
	)
	block = append(block, 
		jen.Id("bsonStruct").
		Op(":=").
		Qual("", "bsonStructPrivate").
		Values(),
	)
	block = append(block, 
		jen.Err().
		Op(":=").
		Qual("go.mongodb.org/mongo-driver/bson", "Unmarshal").
		Call(
			jen.Id("data"),
			jen.Op("&").
			Id("bsonStruct"),
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
	for _, bsonSetterFunction := range bsonSetterFunctions {
		block = append(block, bsonSetterFunction)
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

func (repositoryGenerator *repositoryGenerator) scaffoldEntityBSONStructField(field model.Field, entity model.Entity) (jen.Code, error){
	
	// ID
	var statement jen.Statement
	id , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
	if err != nil {
		return nil, err
	}
	tagId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityJSONTagId(field)
	if err != nil {
		return nil, err
	}

	// Set
	statement.Id(id)

	// Field
	err = repositoryGenerator.helperGenerator.Field("", field, entity, &statement)
	if err != nil {
		return nil, err
	}

	// Tag
	statement.Tag(map[string]string{"bson": tagId})

	return &statement, nil
}