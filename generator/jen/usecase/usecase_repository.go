package usecase

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

type RepositoryGenerator interface{

	File(entity model.Entity) (*jen.File, error)
	ScaffoldFile(entity model.Entity) (*jen.File, error)

	usecaseRepositoryInterface(entity model.Entity) (jen.Statement, error)
	scaffoldUsecaseRepositoryInterface(entity model.Entity) (jen.Statement, error)
	scaffoldUsecaseRepositoryInterfaceMethod(method model.Method, entity model.Entity) (jen.Statement, error)

}

func NewRepositoryGenerator(config *config.Config, formatter formatter.Formatter, helperGenerator helper.Generator) RepositoryGenerator {
	return &repositoryGenerator{
		config : config,
		formatter : formatter,
		helperGenerator :helperGenerator,
	}
}

func (repositoryGenerator *repositoryGenerator) File(entity model.Entity) (*jen.File, error){
	
	// File
	packageName , err := repositoryGenerator.formatter.OutputScaffoldUsecaseRepositoryPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	// Interface
	usecaseRepositoryInterface, err := repositoryGenerator.usecaseRepositoryInterface(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&usecaseRepositoryInterface)

	return f, nil
}

func (repositoryGenerator *repositoryGenerator) ScaffoldFile(entity model.Entity) (*jen.File, error){
	
	
	// File
	packageName , err := repositoryGenerator.formatter.OutputScaffoldUsecaseRepositoryPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	// Interface
	usecaseRepositoryInterface, err := repositoryGenerator.scaffoldUsecaseRepositoryInterface(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&usecaseRepositoryInterface)

	return f, nil
}	

func (repositoryGenerator *repositoryGenerator) usecaseRepositoryInterface(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Type
	resp.Type()

	// ID
	id , err := repositoryGenerator.formatter.OutputUsecaseRepositoryInterfaceId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Scaffold
	scaffoldId , err := repositoryGenerator.formatter.OutputScaffoldUsecaseRepositoryInterfaceId(entity)
	if err != nil {
		return nil, err
	}
	
	// Fields
	fields = append(fields, jen.Id(scaffoldId))


	// Interface
	resp.Interface(fields...)

	return resp, nil

}

func (repositoryGenerator *repositoryGenerator) scaffoldUsecaseRepositoryInterface(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	var methods []jen.Code

	// Loop
	for _, method := range entity.Methods {
		
		// Method
		statement, err := repositoryGenerator.scaffoldUsecaseRepositoryInterfaceMethod(method, entity)
		if err != nil {
			return nil, err
		}

		// Append
		methods = append(methods, &statement)
	}

	// Type
	resp.Type()

	// ID
	id , err := repositoryGenerator.formatter.OutputScaffoldUsecaseRepositoryInterfaceId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Interface
	resp.Interface(methods...)

	return resp, nil

}

func (repositoryGenerator *repositoryGenerator) scaffoldUsecaseRepositoryInterfaceMethod(method model.Method, entity model.Entity) (jen.Statement, error){
	
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
	
	return resp, nil
}