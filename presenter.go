package turnbull

import(
	"github.com/dave/jennifer/jen"
)

func buildScaffoldUsecasePresenterFile(entity entity) (*jen.File, error){

	// Entities
	f := jen.NewFile(config.presenterName)

	// Interface Methods
	var presenterInterface []jen.Code
	for _, method := range entity.Methods {
		
		// Build Method
		err := buildScaffoldUsecasePresenterMethod(entity, method, &presenterInterface)
		if err != nil {
			return nil, err
		}
	}

	// Interface
	f.Type().Id(presenterId(entity)).Interface(presenterInterface...)

	return f, nil
}

func buildScaffoldInterfacePresenterFile(entity entity) (*jen.File, error){

	// Entities
	f := jen.NewFile(config.presenterName)

	// Interface Methods
	var presenterInterface []jen.Code
	// for _, method := range entity.Presenter.Methods {
	// 	var arguments []jen.Code
	// 	var returnValues []jen.Code
	// 	for _, argument := range method.Arguments {
	// 		code, err := fieldToStructField(entity, argument)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		arguments = append(arguments, code)
	// 	}
	// 	for _, returnValue := range method.ReturnValues {
	// 		code, err := fieldToStructField(entity, returnValue)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		returnValues = append(returnValues, code)
	// 	}

	// 	presenterInterface = append(presenterInterface, jen.Id(presenterMethodName(method)).Params(
	// 		arguments...,
	// 	).Parens(jen.List(
	// 		returnValues...,
	// 	)))
	// }

	// Interface
	f.Type().Id(presenterId(entity)).Interface(presenterInterface...)

	return f, nil
}

func buildScaffoldUsecasePresenterMethod(entity entity, entityMethod entityMethod, presenterInterface *[]jen.Code)(error){
	
	var method method
	for _, configMethod := range config.methods {
		if configMethod.Type == entityMethod.Type {
			method = configMethod
		}
	}

	var arguments []jen.Code
	var returnValues []jen.Code
	for _, argument := range method.Presenter.Arguments {
		code, err := fieldToStructField(entity, argument)
		if err != nil {
			return err
		}
		arguments = append(arguments, code)
	}
	for _, returnValue := range method.Presenter.ReturnValues {
		code, err := fieldToStructField(entity, returnValue)
		if err != nil {
			return err
		}
		returnValues = append(returnValues, code)
	}

	*presenterInterface = append(*presenterInterface, jen.Id(presenterMethodName(entityMethod)).Params(
		arguments...,
	).Parens(jen.List(
		returnValues...,
	)))

	return nil
}
