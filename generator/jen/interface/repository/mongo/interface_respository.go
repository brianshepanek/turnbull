package mongo

import(
	"strings"
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

func (repositoryGenerator *repositoryGenerator) EntityFile(entity model.Entity) (*jen.File, error){

	// File
	packageName , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)


	// Struct
	jsonStruct, err := repositoryGenerator.scaffoldInterfaceControllerBSONStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&jsonStruct)
	f.Line()

	// Local Marshal
	localMarshalBSON, err := repositoryGenerator.scaffoldInterfaceControllerLocalMarshalBSONFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&localMarshalBSON)
	f.Line()

	// Local Unmarshal
	localUnmarshalBSON, err := repositoryGenerator.scaffoldInterfaceControllerLocalUnmarshalBSONFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&localUnmarshalBSON)
	f.Line()

	// Marshal
	marshalBSON, err := repositoryGenerator.scaffoldInterfaceControllerMarshalBSONFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&marshalBSON)
	f.Line()

	// Unmarshal
	unmarshalBSON, err := repositoryGenerator.scaffoldInterfaceControllerUnmarshalBSONFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&unmarshalBSON)
	f.Line()

	return f, nil
}

func (repositoryGenerator *repositoryGenerator) RegistryFile(entity model.Entity) (*jen.File, error){

	// File
	packageName , err := repositoryGenerator.formatter.OutputRegistryPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)
	
	// Struct
	interfaceRepositoryStruct, err := repositoryGenerator.scaffoldInterfaceRepositoryRegistryStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceRepositoryStruct)

	// Constructor Function
	interfaceRepositoryRegistryConstructorFunction, err := repositoryGenerator.interfaceRepositoryRegistryConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceRepositoryRegistryConstructorFunction)

	// Local Constructor Function
	interfaceRepositoryRegistryLocalConstructorFunction, err := repositoryGenerator.interfaceRepositoryRegistryLocalConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceRepositoryRegistryLocalConstructorFunction)


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

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryRegistryStruct(entity model.Entity) (jen.Statement, error){

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
	id , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryRegistryStructId("mongo", entity)
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
	id , err := repositoryGenerator.formatter.OutputDomainEntityLocalStructId(entity)
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

func (repositoryGenerator *repositoryGenerator) interfaceRepositoryRegistryConstructorFunction(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// Registry
	registryName , err := repositoryGenerator.formatter.OutputRegistryPackageName()
	if err != nil {
		return nil, err
	}

	// ID
	registryStructId , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryRegistryStructId("mongo", entity)
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

	// Fields
	fields, err := repositoryGenerator.scaffoldInterfaceRepositoryStructFields()
	if err != nil {
		return nil, err
	}

	// ID
	id , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryRegistryConstructorFunctionId("mongo", entity)
	if err != nil {
		return nil, err
	}


	resp.Id(id)

	// Params
	resp.Params(fields...)

	// Block
	resp.Block(

		jen.Id("r").
		Dot(registryStructId).
		Dot("client").
		Op("=").
		Id("client"),

		jen.Id("r").
		Dot(registryStructId).
		Dot("db").
		Op("=").
		Id("db"),

		jen.Id("r").
		Dot(registryStructId).
		Dot("collection").
		Op("=").
		Id("collection"),

	)
	
	return resp, nil

}

func (repositoryGenerator *repositoryGenerator) interfaceRepositoryRegistryLocalConstructorFunction(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement

	// Registry
	registryName , err := repositoryGenerator.formatter.OutputRegistryPackageName()
	if err != nil {
		return nil, err
	}

	// ID
	registryStructId , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryRegistryStructId("mongo", entity)
	if err != nil {
		return nil, err
	}

	interfaceId , err := repositoryGenerator.formatter.OutputUsecaseRepositoryInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	interfaceImportPath , err := repositoryGenerator.formatter.OutputInterfaceRepositoryDirectoryImportPath("mongo", entity)
	if err != nil {
		return nil, err
	}

	usecaseImportPath , err := repositoryGenerator.formatter.OutputScaffoldUsecaseRepositoryDirectoryImportPath()
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
	id , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryRegistryLocalConstructorFunctionId("mongo", entity)
	if err != nil {
		return nil, err
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
			jen.Qual(interfaceImportPath, "New").
			Params(

				jen.Id("r").
				Dot(registryStructId).
				Dot("client"),

				jen.Id("r").
				Dot(registryStructId).
				Dot("db"),

				jen.Id("r").
				Dot(registryStructId).
				Dot("collection"),

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
				jen.If(
					jen.Err().
					Op("!=").
					Qual("go.mongodb.org/mongo-driver/mongo", "ErrNoDocuments").
					Block(
						jen.Return(
							jen.Err(),
						),
					),
				),
			),
		),
	)	

	// Line
	block = append(block, jen.Line())
	
	var nextBlock []jen.Code

	nextBlock = append(nextBlock,
		jen.Id("elem").
		Op(":=").
		Qual(importPath, id).
		Call(),
	)

	if entity.Interface {
		nextBlock = append(nextBlock,
			jen.Err().
			Op(":=").
			Id("cursor").
			Dot("Decode").
			Params(
				jen.Id("elem"),
			),
		)
	} else {
		nextBlock = append(nextBlock,
			jen.Err().
			Op(":=").
			Id("cursor").
			Dot("Decode").
			Params(
				jen.Op("&").
				Id("elem"),
			),
		)
	}	

	nextBlock = append(nextBlock,
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

	if entity.Interface {
		nextBlock = append(nextBlock,
			jen.Id("req").
			Dot("Append").
			Call(
				jen.Id("elem"),
			),
		)
	} else {
		nextBlock = append(nextBlock,
			jen.Id("*req").
			Op("=").
			Append(
				jen.Id("*req"),
				jen.Id("*elem"),
			),
		)
	}
	

	block = append(block,
		jen.For(
			jen.Id("cursor").
			Dot("Next").
			Params(
				jen.Id("ctx"),
			).
			Block(
				nextBlock...,
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

	// Primary Field Name
	var primaryFieldName string
	primaryField, err := repositoryGenerator.helperGenerator.PrimaryField(entity)
	if err != nil {
		return nil, err
	}

	if primaryField != nil {
		primaryFieldName, err = repositoryGenerator.formatter.OutputScaffoldDomainEntityJSONTagId(*primaryField)
		if err != nil {
			return nil, err
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
			jen.Id("req"),
		),
	)	

	block = append(block,
		jen.If(
			jen.Err().
			Op("!=").
			Nil().
			Block(
				jen.If(
					jen.Err().
					Op("!=").
					Qual("go.mongodb.org/mongo-driver/mongo", "ErrNoDocuments").
					Block(
						jen.Return(
							jen.Err(),
						),
					),
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

	// Interface ID
	interfaceId, err := repositoryGenerator.formatter.OutputDomainEntityInterfaceConstructorFunctionId(entity)
	if err != nil {
		return block, err
	}

	// Import Path
	importPath , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityDirectoryImportPath()
	if err != nil {
		return block, err
	}

	// Primary Field Name
	var primaryFieldName string
	primaryField, err := repositoryGenerator.helperGenerator.PrimaryField(entity)
	if err != nil {
		return nil, err
	}

	if primaryField != nil {
		primaryFieldName, err = repositoryGenerator.formatter.OutputScaffoldDomainEntityJSONTagId(*primaryField)
		if err != nil {
			return nil, err
		}
	}
	
	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.Id("current").
		Op(":=").
		Qual(importPath, interfaceId).
		Call(),
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
			jen.Id("current"),
		),
	)	

	block = append(block,
		jen.If(
			jen.Err().
			Op("!=").
			Nil().
			Block(
				jen.If(
					jen.Err().
					Op("!=").
					Qual("go.mongodb.org/mongo-driver/mongo", "ErrNoDocuments").
					Block(
						jen.Return(
							jen.Err(),
						),
					),
				),
			),
		),
	)	

	// Line
	block = append(block, jen.Line())

	// Expanded Fields
	expandedFieldsPointer, err := repositoryGenerator.helperGenerator.ExpandedFields(entity)
	if err != nil {
		return nil, err
	}

	// Fields
	if expandedFieldsPointer != nil {

		for _, field := range *expandedFieldsPointer {
		
			// Getter
			getterId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
			if err != nil {
				return block, err
			}
	
			// Setter
			setterId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntitySetterId(field)
			if err != nil {
				return block, err
			}
	
			// Check
			if entity.Interface {
				block = append(block,
					jen.If(
						jen.Id("req").
						Dot(getterId).
						Call().
						Op("!=").
						Nil().
						Block(
							jen.Id("current").
							Dot(setterId).
							Params(
								jen.Id("req").
								Dot(getterId).
								Call(),
							),
						),
					),
				)
			} else {
				block = append(block,
					jen.If(
						jen.Id("req").
						Dot(getterId).
						Op("!=").
						Nil().
						Block(
							jen.Id("current").
							Dot(getterId).
							Op("=").
							Id("req").
							Dot(getterId),
						),
					),
				)
			}
				
	
			// Line
			block = append(block, jen.Line())
	
		}
	}
	

	block = append(block, 
		jen.List(
			jen.Err(),
		).
		Op("=").
		Id("collection").
		Dot("FindOneAndReplace").
		Params(
			jen.Id("ctx"),
			jen.Id("filter"),
			jen.Id("current"),
		).
		Dot("Decode").
		Params(
			jen.Id("current"),
		),
	)	

	block = append(block,
		jen.If(
			jen.Err().
			Op("!=").
			Nil().
			Block(
				jen.If(
					jen.Err().
					Op("!=").
					Qual("go.mongodb.org/mongo-driver/mongo", "ErrNoDocuments").
					Block(
						jen.Return(
							jen.Err(),
						),
					),
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

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryAddMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	
	var block []jen.Code

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
			jen.Id("req"),
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

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryDeleteMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	var block []jen.Code

	// Primary Field Name
	var primaryFieldName string
	primaryField, err := repositoryGenerator.helperGenerator.PrimaryField(entity)
	if err != nil {
		return nil, err
	}

	if primaryField != nil {
		primaryFieldName, err = repositoryGenerator.formatter.OutputScaffoldDomainEntityJSONTagId(*primaryField)
		if err != nil {
			return nil, err
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


func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceControllerBSONStruct(entity model.Entity) (jen.Statement, error){

	var statement jen.Statement
	var fields []jen.Code
	for _, field := range entity.Fields {
		if !field.Embedded {
			code, err := repositoryGenerator.scaffoldEntityBSONStructField(field, entity)
			if err != nil {
				return nil, err
			}
			
			fields = append(fields, code)
		} else {


			
			// Field ID
			fieldId , err := repositoryGenerator.formatter.OutputDomainEntityInterfaceId(field.Entity)
			if err != nil {
				return nil, err
			}

			// Struct ID
			structId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityMarshalStructId("bson", field.Entity)
			if err != nil {
				return nil, err
			}
			fields = append(fields, 
				jen.Id(fieldId).
				Op("*").
				Qual("", structId).
				Tag(map[string]string{"bson": strings.Join([]string{"inline"}, ",")}),
			)

		}
		
		
	}

	// Struct ID
	structId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityMarshalStructId("bson", entity)
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

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceControllerLocalMarshalBSONFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement

	// Struct ID
	structId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}
	
	// Marshal Struct ID
	marshalStructId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityMarshalStructId("bson", entity)
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
	statement.Id("marshalBSON")

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
		jen.Id("bsonStruct").
		Op(":=").
		Id(marshalStructId).
		Values(),
	)

	block = append(block, jen.Line())

	for _, field := range entity.Fields {

		if !field.Embedded {

			// Getter ID
			getterId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
			if err != nil {
				return nil, err
			}

			if entity.Interface {
				block = append(block,
					jen.Id("bsonStruct").
					Dot(getterId).
					Op("=").
					Id("m").Dot(getterId).Call(),
				)
			} else {
				block = append(block,
					jen.Id("bsonStruct").
					Dot(getterId).
					Op("=").
					Id("m").Dot(getterId),
				)
			}
		} else {

			// Vars
			var entityId string

			// Struct ID
			fieldId , err := repositoryGenerator.formatter.OutputDomainEntityInterfaceId(field.Entity)
			if err != nil {
				return nil, err
			}

			// Interface ID
			if field.Entity.Interface {
				structId , err := repositoryGenerator.formatter.OutputDomainEntityStructId(field.Entity)
				if err != nil {
					return nil, err
				}
				entityId = structId
			} else {
				interfaceId , err := repositoryGenerator.formatter.OutputDomainEntityInterfaceId(field.Entity)
				if err != nil {
					return nil, err
				}
				entityId = interfaceId
			}

			block = append(block,
				jen.Id("bsonStruct").
				Dot(fieldId).
				Op("=").
				Id("m").
				Dot(entityId).
				Dot("marshalBSON").
				Call(),
			)

		}	
	}

	block = append(block, jen.Line())

	block = append(block,
		jen.Return(
			jen.Op("&").
			Id("bsonStruct"),
		),
	)
	
	statement.Block(
		block...,
	)
	
	
	return statement, nil
}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceControllerMarshalBSONFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement

	// Struct ID
	structId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityStructId(entity)
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


	// Block
	statement.Block(
		jen.Return(
			jen.Qual("go.mongodb.org/mongo-driver/bson", "Marshal").
			Call(
				jen.Op("m").
				Dot("marshalBSON").
				Call(),
			),
		),
	)
	
	
	return statement, nil
}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceControllerLocalUnmarshalBSONFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement

	// Struct ID
	structId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// Marshal Struct ID
	marshalStructId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityMarshalStructId("bson", entity)
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
	statement.Id("unmarshalBSON")

	// Params
	statement.Params(
		jen.Id("bsonStruct").
		Op("*").
		Qual("", marshalStructId),
	)

	// Block
	var block []jen.Code

	
	for _, field := range entity.Fields {

		if !field.Embedded {

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
			

			if entity.Interface {
				block = append(block, 
					jen.Id("m").
					Dot(setterId).
					Call(
						jen.Id("bsonStruct").
						Dot(getterId),
					),
				)
			} else {
				block = append(block, 
					jen.Id("m").
					Dot(getterId).
					Op("=").
					Id("bsonStruct").
					Dot(getterId),
				)
			}	
		} else {

			// Vars
			var entityId string

			// Struct ID
			fieldId , err := repositoryGenerator.formatter.OutputDomainEntityInterfaceId(field.Entity)
			if err != nil {
				return nil, err
			}

			// Interface ID
			if field.Entity.Interface {
				structId , err := repositoryGenerator.formatter.OutputDomainEntityStructId(field.Entity)
				if err != nil {
					return nil, err
				}
				entityId = structId
			} else {
				interfaceId , err := repositoryGenerator.formatter.OutputDomainEntityInterfaceId(field.Entity)
				if err != nil {
					return nil, err
				}
				entityId = interfaceId
			}

			block = append(block,
				jen.Id("m").
				Dot(entityId).
				Dot("unmarshalBSON").
				Params(
					jen.Id("bsonStruct").
					Dot(fieldId),
				),
			)

		} 
	}	

	statement.Block(
		block...,
	)
	
	
	return statement, nil
}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceControllerUnmarshalBSONFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement

	// Struct ID
	structId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// Marshal Struct ID
	marshalStructId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityMarshalStructId("bson", entity)
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

	// Block
	var block []jen.Code

	block = append(block, jen.Line())

	block = append(block, 
		jen.Id("bsonStruct").
		Op(":=").
		Qual("", marshalStructId).
		Values(),
	)

	for _, field := range entity.Fields {

		if field.Embedded {

			// Field ID
			fieldId , err := repositoryGenerator.formatter.OutputDomainEntityInterfaceId(field.Entity)
			if err != nil {
				return nil, err
			}

			// Struct ID
			structId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityMarshalStructId("bson", field.Entity)
			if err != nil {
				return nil, err
			}

			block = append(block, 
				jen.Id("bsonStruct").
				Dot(fieldId).
				Op("=").
				Op("&").
				Id(structId).
				Values(),
			)

		}
		

	}

	block = append(block, jen.Line())

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
	
	block = append(block, jen.Line())

	block = append(block, 
		jen.Id("m").
		Dot("unmarshalBSON").
		Params(
			jen.Op("&").
			Id("bsonStruct"),
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
	field.Op = "*"
	err = repositoryGenerator.helperGenerator.Field("", field, entity, &statement)
	if err != nil {
		return nil, err
	}

	// Tag
	statement.Tag(map[string]string{"bson": strings.Join([]string{tagId, "omitempty"}, ",")})

	return &statement, nil
}