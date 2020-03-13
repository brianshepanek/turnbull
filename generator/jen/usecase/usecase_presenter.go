package usecase

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

type PresenterGenerator interface{

	ScaffoldFile(entity model.Entity) (*jen.File, error)

	scaffoldUsecasePresenterInterface(entity model.Entity) (jen.Statement, error)
	scaffoldUsecasePresenterInterfaceMethod(method model.Method, entity model.Entity) (jen.Statement, error)

}

func NewPresenterGenerator(config *config.Config, formatter formatter.Formatter, helperGenerator helper.Generator) PresenterGenerator {
	return &presenterGenerator{
		config : config,
		formatter : formatter,
		helperGenerator :helperGenerator,
	}
}
func (presenterGenerator *presenterGenerator) ScaffoldFile(entity model.Entity) (*jen.File, error){
	
	
	// File
	packageName , err := presenterGenerator.formatter.OutputScaffoldUsecasePresenterPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	// Interface
	usecasePresenterInterface, err := presenterGenerator.scaffoldUsecasePresenterInterface(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&usecasePresenterInterface)

	return f, nil
}	

func (presenterGenerator *presenterGenerator) scaffoldUsecasePresenterInterface(entity model.Entity) (jen.Statement, error){

	// Vars
	var resp jen.Statement
	var methods []jen.Code

	// Loop
	for _, method := range entity.Methods {
		
		// Method
		statement, err := presenterGenerator.scaffoldUsecasePresenterInterfaceMethod(method, entity)
		if err != nil {
			return nil, err
		}

		// Append
		methods = append(methods, &statement)
	}

	// Type
	resp.Type()

	// ID
	id , err := presenterGenerator.formatter.OutputScaffoldUsecasePresenterInterfaceId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Interface
	resp.Interface(methods...)

	return resp, nil

}

func (presenterGenerator *presenterGenerator) scaffoldUsecasePresenterInterfaceMethod(method model.Method, entity model.Entity) (jen.Statement, error){
	
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