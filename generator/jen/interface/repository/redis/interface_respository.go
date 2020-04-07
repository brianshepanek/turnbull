package redis

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

	f.ImportName("github.com/go-redis/redis/v7", "redis")

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

	f.ImportName("github.com/go-redis/redis/v7", "redis")

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
	return nil, nil
}

func (repositoryGenerator *repositoryGenerator) RegistryFile(entity model.Entity) (*jen.File, error){

	// File
	packageName , err := repositoryGenerator.formatter.OutputRegistryPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	f.ImportName("github.com/go-redis/redis/v7", "redis")
	
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
	id , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryRegistryStructId("redis", entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Struct
	resp.Struct(fields...)

	
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
	registryStructId , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryRegistryStructId("redis", entity)
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
	id , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryRegistryConstructorFunctionId("redis", entity)
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
		Dot("namespace").
		Op("=").
		Id("namespace"),

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
	registryStructId , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryRegistryStructId("redis", entity)
	if err != nil {
		return nil, err
	}

	interfaceId , err := repositoryGenerator.formatter.OutputUsecaseRepositoryInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	interfaceImportPath , err := repositoryGenerator.formatter.OutputInterfaceRepositoryDirectoryImportPath("redis", entity)
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
	id , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryRegistryLocalConstructorFunctionId("redis", entity)
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
				Dot("namespace"),

			),
		),
	)
	
	return resp, nil

}

func (repositoryGenerator *repositoryGenerator) interfaceRepositoryStruct(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Type
	resp.Type()

	// ID
	id , err := repositoryGenerator.formatter.OutputInterfaceRepositoryStructId("redis", entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Scaffold
	scaffoldId , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryStructId("redis", entity)
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
		Qual("github.com/go-redis/redis/v7", "Client"),
	)

	// Namespace
	fields = append(fields,
		jen.Id("namespace").
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
	id , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryStructId("redis", entity)
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
	id , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryConstructorFunctionId("redis", entity)
	if err != nil {
		return nil, err
	}

	// Struct ID
	structId , err := repositoryGenerator.formatter.OutputInterfaceRepositoryStructId("redis", entity)
	if err != nil {
		return nil, err
	}

	// Scaffold Struct ID
	scaffoldStructId , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryStructId("redis", entity)
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
						jen.Id("namespace"):	jen.Id("namespace"),
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
	structId , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryStructId("redis", entity)
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
	interfaceId, err := repositoryGenerator.formatter.OutputDomainEntityInterfaceConstructorFunctionId(entity)
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

	// Match
	block = append(block, 
		jen.Id("match").
		Op(":=").
		Id("r").
		Dot("namespace").
		Op("+").
		Lit(":*"),
	)

	// Keys
	block = append(block, 
		jen.List(
			jen.Id("keys"),
			jen.Id("_"),
			jen.Err(),
		).
		Op(":=").
		Id("r").
		Dot("client").
		Dot("Scan").
		Params(
			jen.Lit(0),
			jen.Id("match"),
			jen.Lit(-1),
		).
		Dot("Result").
		Call(),
	)	

	// Err
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

	// For
	var forBlock []jen.Code

	// Line
	forBlock = append(forBlock, jen.Line())

	// Elem
	forBlock = append(forBlock,
		jen.Id("elem").
		Op(":=").
		Qual(importPath, interfaceId).
		Call(),
	)

	// Val
	forBlock = append(forBlock, 
		jen.List(
			jen.Id("val"),
			jen.Err(),
		).
		Op(":=").
		Id("r").
		Dot("client").
		Dot("Get").
		Params(
			jen.Id("key"),
		).
		Dot("Result").
		Call(),
	)	

	// Err
	forBlock = append(forBlock,
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
	forBlock = append(forBlock, jen.Line())

	// Unmarshal
	forBlock = append(forBlock,
		jen.Err().
		Op("=").
		Qual("encoding/json", "Unmarshal").
		Params(
			jen.Index().
			Byte().
			Parens(
				jen.Id("val"),
			),
			jen.Id("elem"),
		),
	)	

	// Err
	forBlock = append(forBlock,
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
	forBlock = append(forBlock, jen.Line())

	if entity.Interface {
		forBlock = append(forBlock,
			jen.Id("req").
			Dot("Append").
			Call(
				jen.Id("elem"),
			),
		)
	} else {
		forBlock = append(forBlock,
			jen.Id("*req").
			Op("=").
			Append(
				jen.Id("*req"),
				jen.Id("*elem"),
			),
		)
	}

	// Line
	forBlock = append(forBlock, jen.Line())

	// For
	block = append(block,
		jen.For(
			jen.List(
				jen.Id("_"),
				jen.Id("key"),
			).
			Op(":=").
			Range().
			Id("keys").
			Block(forBlock...),
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

	// Line
	block = append(block, jen.Line())

	// Key
	block = append(block, 
		jen.Id("key").
		Op(":=").
		Id("r").
		Dot("namespace").
		Op("+").
		Lit(":").
		Op("+").
		Qual("github.com/spf13/cast", "ToString").
		Params(
			jen.Id("id"),
		),
	)

	// Val
	block = append(block, 
		jen.List(
			jen.Id("val"),
			jen.Err(),
		).
		Op(":=").
		Id("r").
		Dot("client").
		Dot("Get").
		Params(
			jen.Id("key"),
		).
		Dot("Result").
		Call(),
	)	

	// Err
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

	// Unmarshal
	block = append(block,
		jen.Err().
		Op("=").
		Qual("encoding/json", "Unmarshal").
		Params(
			jen.Index().
			Byte().
			Parens(
				jen.Id("val"),
			),
			jen.Id("req"),
		),
	)	

	// Err
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

	// Return
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
	
	// Line
	block = append(block, jen.Line())

	block = append(block, 
		jen.Id("current").
		Op(":=").
		Qual(importPath, interfaceId).
		Call(),
	)	

	// Key
	block = append(block, 
		jen.Id("key").
		Op(":=").
		Id("r").
		Dot("namespace").
		Op("+").
		Lit(":").
		Op("+").
		Qual("github.com/spf13/cast", "ToString").
		Params(
			jen.Id("id"),
		),
	)

	// Val
	block = append(block, 
		jen.List(
			jen.Id("currentVal"),
			jen.Err(),
		).
		Op(":=").
		Id("r").
		Dot("client").
		Dot("Get").
		Params(
			jen.Id("key"),
		).
		Dot("Result").
		Call(),
	)	

	// Err
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

	// Unmarshal
	block = append(block,
		jen.Err().
		Op("=").
		Qual("encoding/json", "Unmarshal").
		Params(
			jen.Index().
			Byte().
			Parens(
				jen.Id("currentVal"),
			),
			jen.Id("current"),
		),
	)	

	// Err
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

	// Fields
	for _, field := range entity.Fields {
		
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

	// Marshal
	block = append(block,
		jen.List(
			jen.Id("val"),
			jen.Err(),
		).
		Op(":=").
		Qual("encoding/json", "Marshal").
		Params(
			jen.Id("current"),
		),
	)	

	// Err
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

	// Set
	block = append(block, 
		jen.List(
			jen.Err(),
		).
		Op("=").
		Id("r").
		Dot("client").
		Dot("Set").
		Params(
			jen.Id("key"),
			jen.Id("val"),
			jen.Lit(0),
		).
		Dot("Err").
		Call(),
	)	

	// Err
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

	// Return
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

	// Var
	var primaryField model.Field
	for _, field := range entity.Fields {
		if field.Primary {
			primaryField = field
		}
	}

	getterId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityGetterId(primaryField)
	if err != nil {
		return nil, err
	}

	// Line
	block = append(block, jen.Line())
	
	// Key
	if entity.Interface {
		block = append(block, 
			jen.Id("key").
			Op(":=").
			Id("r").
			Dot("namespace").
			Op("+").
			Lit(":").
			Op("+").
			Qual("github.com/spf13/cast", "ToString").
			Params(
				jen.Id("req").
				Dot(getterId).
				Call(),
			),
		)
	} else {
		block = append(block, 
			jen.Id("key").
			Op(":=").
			Id("r").
			Dot("namespace").
			Op("+").
			Lit(":").
			Op("+").
			Qual("github.com/spf13/cast", "ToString").
			Params(
				jen.Id("req").
				Dot(getterId),
			),
		)
	}
	
	// Marshal
	block = append(block,
		jen.List(
			jen.Id("val"),
			jen.Err(),
		).
		Op(":=").
		Qual("encoding/json", "Marshal").
		Params(
			jen.Id("req"),
		),
	)	

	// Err
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

	// Set
	block = append(block, 
		jen.List(
			jen.Err(),
		).
		Op("=").
		Id("r").
		Dot("client").
		Dot("Set").
		Params(
			jen.Id("key"),
			jen.Id("val"),
			jen.Lit(0),
		).
		Dot("Err").
		Call(),
	)	

	// Err
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
	
	// Return
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
	
	// Line
	block = append(block, jen.Line())

	// Key
	block = append(block, 
		jen.Id("key").
		Op(":=").
		Id("r").
		Dot("namespace").
		Op("+").
		Lit(":").
		Op("+").
		Qual("github.com/spf13/cast", "ToString").
		Params(
			jen.Id("id"),
		),
	)

	// Val
	block = append(block, 
		jen.List(
			jen.Err(),
		).
		Op(":=").
		Id("r").
		Dot("client").
		Dot("Del").
		Params(
			jen.Id("key"),
		).
		Dot("Err").
		Call(),
	)	

	// Err
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

	// Return
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