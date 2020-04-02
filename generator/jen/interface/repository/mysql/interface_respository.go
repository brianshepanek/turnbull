package mysql

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

	// Entity Struct
	interfaceRepositoryEntityStruct, err := repositoryGenerator.scaffoldInterfaceRepositoryEntityStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceRepositoryEntityStruct)

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
	id , err := repositoryGenerator.formatter.OutputInterfaceRepositoryStructId("mysql", entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Scaffold
	scaffoldId , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryStructId("mysql", entity)
	if err != nil {
		return nil, err
	}
	
	// Fields
	fields = append(fields, jen.Id(scaffoldId))


	// Struct
	resp.Struct(fields...)

	return resp, nil

}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryEntityStruct(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Type
	resp.Type()

	// ID
	id , err := repositoryGenerator.formatter.OutputDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	for _, field := range entity.Fields {
		
		code, err := repositoryGenerator.scaffoldEntityStructField(field, entity)
		if err != nil {
			return nil, err
		}
		
		fields = append(fields, code)
	}

	// Struct
	resp.Struct(fields...)
	
	return resp, nil

}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryStructFields() ([]jen.Code, error){

	// Vars
	var fields []jen.Code

	// Driver
	fields = append(fields,
		jen.Id("db").
		Op("*").
		Qual("database/sql", "DB"),
	)

	// Collection
	fields = append(fields,
		jen.Id("table").
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
	id , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryStructId("mysql", entity)
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
	id , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryConstructorFunctionId("mysql", entity)
	if err != nil {
		return nil, err
	}

	// Struct ID
	structId , err := repositoryGenerator.formatter.OutputInterfaceRepositoryStructId("mysql", entity)
	if err != nil {
		return nil, err
	}

	// Scaffold Struct ID
	scaffoldStructId , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryStructId("mysql", entity)
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
						jen.Id("db"):	jen.Id("db"),
						jen.Id("table"):	jen.Id("table"),
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
	structId , err := repositoryGenerator.formatter.OutputScaffoldInterfaceRepositoryStructId("mysql", entity)
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
	
	var block, nextBlock []jen.Code

	// ID
	localId , err := repositoryGenerator.formatter.OutputDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// Interface ID
	constructorFunctionId , err := repositoryGenerator.formatter.OutputDomainEntityInterfaceConstructorFunctionId(entity)
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

	var fields []string
	var scanFields []jen.Code
	for _, field := range entity.Fields {

		getterId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
		if err != nil {
			return nil, err
		}

		tagId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityJSONTagId(field)
		if err != nil {
			return nil, err
		}

		fields = append(fields, tagId)
		scanFields = append(scanFields, 
			jen.Op("&").
			Id("res").
			Dot(getterId),
		)

	}

	// Vars
	block = append(block,
		jen.Var().
		List(
			jen.Id("sqlStatement"),
		).
		String(),
	)

	block = append(block,
		jen.Id("sqlStatement").
		Op("+=").
		Lit(`SELECT ` + strings.Join(fields, ", ") + ` FROM `).
		Op("+").
		Id("r").
		Dot("table"),
	)

	// Line
	block = append(block, jen.Line())

	// Prepare
	block = append(block,
		jen.List(
			jen.Id("stmt"),
			jen.Err(),
		).
		Op(":=").
		Id("r").
		Dot("db").
		Dot("Prepare").
		Params(
			jen.Id("sqlStatement"),
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
	block = append(block,
		jen.Defer().
		Id("stmt").
		Dot("Close").
		Call(),
	)

	// Line
	block = append(block, jen.Line())

	// Rows
	block = append(block,
		jen.List(
			jen.Id("rows"),
			jen.Err(),
		).
		Op(":=").
		Id("stmt").
		Dot("Query").
		Call(),
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
	block = append(block,
		jen.Defer().
		Id("rows").
		Dot("Close").
		Call(),
	)

	// Line
	block = append(block, jen.Line())

	// Scan
	nextBlock = append(nextBlock,
		jen.Line(),
	)

	nextBlock = append(nextBlock,
		jen.Var().
		Id("res").
		Qual("", localId),
	)

	nextBlock = append(nextBlock,
		jen.Line(),
	)

	nextBlock = append(nextBlock,
		jen.Err().
		Op(":=").
		Id("rows").
		Dot("Scan").
		Params(
			scanFields...
		),
	)

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

	nextBlock = append(nextBlock,
		jen.Line(),
	)

	nextBlock = append(nextBlock,
		jen.Id("elem").
		Op(":=").
		Qual(importPath, constructorFunctionId).
		Call(),
	)	

	nextBlock = append(nextBlock,
		jen.Line(),
	)

	for _, field := range entity.Fields {

		getterId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
		if err != nil {
			return nil, err
		}

		setterId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntitySetterId(field)
		if err != nil {
			return nil, err
		}

		nextBlock = append(nextBlock,
			jen.If().
			Id("res").
			Dot(getterId).
			Dot("Valid").
			Block(

				jen.Var().
				Id("val").
				Qual(field.Package, field.Type),

				jen.Err().
				Op(":=").
				Id("res").
				Dot(getterId).
				Dot("Scan").
				Params(
					jen.Op("&").
					Id("val"),
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

				jen.Id("elem").
				Dot(setterId).
				Params(
					jen.Op("&").
					Id("val"),
				),
				
			),
		)	

		nextBlock = append(nextBlock,
			jen.Line(),
		)
	}

	nextBlock = append(nextBlock,
		jen.Id("req").
		Dot("Append").
		Call(
			jen.Id("elem"),
		),
	)

	nextBlock = append(nextBlock,
		jen.Line(),
	)

	block = append(block,
		jen.For().
		Id("rows").
		Dot("Next").
		Call().
		Block(nextBlock...),
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
	localId , err := repositoryGenerator.formatter.OutputDomainEntityStructId(entity)
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

	var fields []string
	var scanFields []jen.Code
	for _, field := range entity.Fields {

		getterId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
		if err != nil {
			return nil, err
		}

		tagId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityJSONTagId(field)
		if err != nil {
			return nil, err
		}

		fields = append(fields, tagId)
		scanFields = append(scanFields, 
			jen.Op("&").
			Id("res").
			Dot(getterId),
		)

	}

	// Vars
	block = append(block,
		jen.Var().
		List(
			jen.Id("sqlStatement"),
		).
		String(),
	)

	block = append(block,
		jen.Id("sqlStatement").
		Op("+=").
		Lit(`SELECT ` + strings.Join(fields, ", ") + ` FROM `).
		Op("+").
		Id("r").
		Dot("table").
		Op("+").
		Lit(` WHERE ` + primaryFieldName + ` = ?`),
	)

	// Line
	block = append(block, jen.Line())

	// Prepare
	block = append(block,
		jen.List(
			jen.Id("stmt"),
			jen.Err(),
		).
		Op(":=").
		Id("r").
		Dot("db").
		Dot("Prepare").
		Params(
			jen.Id("sqlStatement"),
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
	block = append(block,
		jen.Defer().
		Id("stmt").
		Dot("Close").
		Call(),
	)

	// Line
	block = append(block, jen.Line())

	// Rows
	block = append(block,
		jen.List(
			jen.Id("row"),
		).
		Op(":=").
		Id("stmt").
		Dot("QueryRow").
		Params(
			jen.Id("id"),
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

	// Scan
	block = append(block,
		jen.Line(),
	)

	block = append(block,
		jen.Var().
		Id("res").
		Qual("", localId),
	)

	block = append(block,
		jen.Line(),
	)

	block = append(block,
		jen.Err().
		Op("=").
		Id("row").
		Dot("Scan").
		Params(
			scanFields...
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

	block = append(block,
		jen.Line(),
	)

	for _, field := range entity.Fields {

		getterId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
		if err != nil {
			return nil, err
		}

		setterId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntitySetterId(field)
		if err != nil {
			return nil, err
		}

		block = append(block,
			jen.If().
			Id("res").
			Dot(getterId).
			Dot("Valid").
			Block(

				jen.Var().
				Id("val").
				Qual(field.Package, field.Type),

				jen.Err().
				Op(":=").
				Id("res").
				Dot(getterId).
				Dot("Scan").
				Params(
					jen.Op("&").
					Id("val"),
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
				Dot(setterId).
				Params(
					jen.Op("&").
					Id("val"),
				),
				
			),
		)	

		block = append(block,
			jen.Line(),
		)
	}


	block = append(block,
		jen.Line(),
	)

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

	// Primary Field Name
	var err error
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

	// Vars
	block = append(block,
		jen.Var().
		List(
			jen.Id("set"),
		).
		Index().
		String(),
	)

	block = append(block,
		jen.Var().
		List(
			jen.Id("vals"),
		).
		Index().
		Interface(),
	)

	for _, field := range entity.Fields {
		
		// Getter ID
		getterId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
		if err != nil {
			return nil, err
		}

		tagId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityJSONTagId(field)
		if err != nil {
			return nil, err
		}

		block = append(block,
			jen.If().
			Id("req").
			Dot(getterId).
			Call().
			Op("!=").
			Nil().
			Block(

				jen.Id("set").
				Op("=").
				Append(
					jen.Id("set"),
					jen.Lit(tagId + ` = ?`),
				),

				jen.Id("vals").
				Op("=").
				Append(
					jen.Id("vals"),
					jen.Id("req").
					Dot(getterId).
					Call(),
				),

			),
		)

		// Line
		block = append(block, jen.Line())

	}

	block = append(block, 
		jen.Id("vals").
		Op("=").
		Append(
			jen.Id("vals"),
			jen.Id("id"),
		),
	)	

	// Line
	block = append(block, jen.Line())

	// Vars
	block = append(block,
		jen.Var().
		List(
			jen.Id("sqlStatement"),
		).
		String(),
	)

	block = append(block,
		jen.Id("sqlStatement").
		Op("+=").
		Lit(`UPDATE `).
		Op("+").
		Id("r").
		Dot("table").
		Op("+").
		Lit(" "),
	)

	block = append(block,
		jen.If().
		Len(
			jen.Id("set"),
		).
		Op(">").
		Lit(0).
		Block(
			jen.Id("sqlStatement").
			Op("+=").
			Lit(`SET `).
			Op("+").
			Qual("strings", "Join").
			Params(
				jen.Id("set"),
				jen.Lit(`, `),
			).
			Op("+").
			Lit(" "),
		),
	)

	block = append(block,
		jen.Id("sqlStatement").
		Op("+=").
		Lit(` WHERE ` + primaryFieldName + ` = ?`),
	)	

	// Line
	block = append(block, jen.Line())

	// Prepare
	block = append(block,
		jen.List(
			jen.Id("stmt"),
			jen.Err(),
		).
		Op(":=").
		Id("r").
		Dot("db").
		Dot("Prepare").
		Params(
			jen.Id("sqlStatement"),
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
	block = append(block,
		jen.Defer().
		Id("stmt").
		Dot("Close").
		Call(),
	)
	
	// Line
	block = append(block, jen.Line())

	// Exec
	block = append(block,
		jen.List(
			jen.Id("_"),
			jen.Err(),
		).
		Op("=").
		Id("stmt").
		Dot("Exec").
		Params(
			jen.Id("vals").
			Op("..."),
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

	block = append(block,
		jen.Return(
			jen.Nil(),
		),
	)
	return block, nil
}

func (repositoryGenerator *repositoryGenerator) scaffoldInterfaceRepositoryAddMethodBlock(method model.Method, entity model.Entity) ([]jen.Code, error){
	
	var block []jen.Code

	// Line
	block = append(block, jen.Line())

	// Vars
	block = append(block,
		jen.Var().
		List(
			jen.Id("set"),
			jen.Id("vars"),
		).
		Index().
		String(),
	)

	block = append(block,
		jen.Var().
		List(
			jen.Id("vals"),
		).
		Index().
		Interface(),
	)

	for _, field := range entity.Fields {
		
		// Getter ID
		getterId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
		if err != nil {
			return nil, err
		}

		tagId , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityJSONTagId(field)
		if err != nil {
			return nil, err
		}

		block = append(block,
			jen.If().
			Id("req").
			Dot(getterId).
			Call().
			Op("!=").
			Nil().
			Block(

				jen.Id("set").
				Op("=").
				Append(
					jen.Id("set"),
					jen.Lit(tagId),
				),

				jen.Id("vars").
				Op("=").
				Append(
					jen.Id("vars"),
					jen.Lit("?"),
				),

				jen.Id("vals").
				Op("=").
				Append(
					jen.Id("vals"),
					jen.Id("req").
					Dot(getterId).
					Call(),
				),

			),
		)

		// Line
		block = append(block, jen.Line())

	}

	// Vars
	block = append(block,
		jen.Var().
		List(
			jen.Id("sqlStatement"),
		).
		String(),
	)

	block = append(block,
		jen.Id("sqlStatement").
		Op("+=").
		Lit(`INSERT INTO `).
		Op("+").
		Id("r").
		Dot("table").
		Op("+").
		Lit(" "),
	)

	block = append(block,
		jen.If().
		Len(
			jen.Id("set"),
		).
		Op(">").
		Lit(0).
		Block(
			jen.Id("sqlStatement").
			Op("+=").
			Lit(`(`).
			Op("+").
			Qual("strings", "Join").
			Params(
				jen.Id("set"),
				jen.Lit(`, `),
			).
			Op("+").
			Lit(") "),
		),
	)

	block = append(block,
		jen.If().
		Len(
			jen.Id("vars"),
		).
		Op(">").
		Lit(0).
		Block(
			jen.Id("sqlStatement").
			Op("+=").
			Lit(`VALUES(`).
			Op("+").
			Qual("strings", "Join").
			Params(
				jen.Id("vars"),
				jen.Lit(`, `),
			).
			Op("+").
			Lit(") "),
		),
	)

	// Line
	block = append(block, jen.Line())

	// Prepare
	block = append(block,
		jen.List(
			jen.Id("stmt"),
			jen.Err(),
		).
		Op(":=").
		Id("r").
		Dot("db").
		Dot("Prepare").
		Params(
			jen.Id("sqlStatement"),
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
	block = append(block,
		jen.Defer().
		Id("stmt").
		Dot("Close").
		Call(),
	)
	
	// Line
	block = append(block, jen.Line())

	// Exec
	block = append(block,
		jen.List(
			jen.Id("_"),
			jen.Err(),
		).
		Op("=").
		Id("stmt").
		Dot("Exec").
		Params(
			jen.Id("vals").
			Op("..."),
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
	var err error

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


	// Vars
	block = append(block,
		jen.Var().
		List(
			jen.Id("sqlStatement"),
		).
		String(),
	)

	block = append(block,
		jen.Id("sqlStatement").
		Op("+=").
		Lit(`DELETE FROM `).
		Op("+").
		Id("r").
		Dot("table").
		Op("+").
		Lit(` WHERE ` + primaryFieldName + ` = ?`),
	)

	// Line
	block = append(block, jen.Line())

	// Prepare
	block = append(block,
		jen.List(
			jen.Id("stmt"),
			jen.Err(),
		).
		Op(":=").
		Id("r").
		Dot("db").
		Dot("Prepare").
		Params(
			jen.Id("sqlStatement"),
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
	block = append(block,
		jen.Defer().
		Id("stmt").
		Dot("Close").
		Call(),
	)

	// Line
	block = append(block, jen.Line())

	// Rows
	block = append(block,
		jen.List(
			jen.Id("_"),
			jen.Err(),
		).
		Op("=").
		Id("stmt").
		Dot("Exec").
		Params(
			jen.Id("id"),
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

	// Scan
	block = append(block,
		jen.Line(),
	)

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

func (repositoryGenerator *repositoryGenerator) scaffoldEntityStructField(field model.Field, entity model.Entity) (jen.Code, error){
	
	// ID
	var statement jen.Statement
	id , err := repositoryGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
	if err != nil {
		return nil, err
	}
	
	// Set
	statement.Id(id)

	sqlField := field
	switch sqlField.Type {
	case "bool":
		sqlField.Package = "database/sql"
		sqlField.Type = "NullBool"
	case "string":
		sqlField.Package = "database/sql"
		sqlField.Type = "NullString"		
	case "int":
		sqlField.Package = "database/sql"
		sqlField.Type = "NullInt32"
	case "int32":
		sqlField.Package = "database/sql"
		sqlField.Type = "NullInt32"	
	case "int64":
		sqlField.Package = "database/sql"
		sqlField.Type = "NullInt64"
	case "float32":
		sqlField.Package = "database/sql"
		sqlField.Type = "NullFloat64"	
	case "float64":
		sqlField.Package = "database/sql"
		sqlField.Type = "NullFloat64"
	case "Time":
		sqlField.Package = "database/sql"
		sqlField.Type = "NullTime"				
	default:
		
	}


	// Field
	err = repositoryGenerator.helperGenerator.Field("", sqlField, entity, &statement)
	if err != nil {
		return nil, err
	}


	return &statement, nil
}