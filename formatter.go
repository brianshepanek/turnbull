package turnbull

import(
	"strings"
	"os"
	// "github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

func packageName()(string){
	return strcase.ToSnake(config.EntityName)
}

func scaffoldEntityFileName(outputPath string, entity entity)(string){
	return strings.Join([]string{outputPath, config.scaffoldDirName, config.domainLayerName, config.entitiesDirName, strings.Join([]string{strcase.ToSnake(entity.Name), "go"}, config.stringSeparator)}, config.pathSeparator)
}

func scaffoldEntitiesFilePath()(string){
	relativePackagePath := config.absOutputPath
	relativePackagePath = strings.TrimLeft(strings.Replace(relativePackagePath, strings.Join([]string{os.Getenv("GOPATH"), config.workspaceSourceDirName}, config.pathSeparator), "", 1), config.pathSeparator)
	return strings.Join([]string{relativePackagePath, config.scaffoldDirName, config.domainLayerName, config.entitiesDirName}, config.pathSeparator)
}

func scaffoldRepositoryFilePath()(string){
	relativePackagePath := config.absOutputPath
	relativePackagePath = strings.TrimLeft(strings.Replace(relativePackagePath, strings.Join([]string{os.Getenv("GOPATH"), config.workspaceSourceDirName}, config.pathSeparator), "", 1), config.pathSeparator)
	return strings.Join([]string{relativePackagePath, config.scaffoldDirName, config.usecaseLayerName, config.repositoryName}, config.pathSeparator)
}

func scaffoldPresenterFilePath()(string){
	relativePackagePath := config.absOutputPath
	relativePackagePath = strings.TrimLeft(strings.Replace(relativePackagePath, strings.Join([]string{os.Getenv("GOPATH"), config.workspaceSourceDirName}, config.pathSeparator), "", 1), config.pathSeparator)
	return strings.Join([]string{relativePackagePath, config.scaffoldDirName, config.usecaseLayerName, config.presenterName}, config.pathSeparator)
}

func scaffoldUsecaseRepositoryFileName(outputPath string, entity entity)(string){
	return strings.Join([]string{outputPath, config.scaffoldDirName, config.usecaseLayerName, config.repositoryName, strings.Join([]string{ strcase.ToSnake( strings.Join([]string{entity.Name, config.repositoryName}, " ") )  , "go"}, config.stringSeparator)}, config.pathSeparator)
}

func scaffoldUsecasePresenterFileName(outputPath string, entity entity)(string){
	return strings.Join([]string{outputPath, config.scaffoldDirName, config.usecaseLayerName, config.presenterName, strings.Join([]string{ strcase.ToSnake( strings.Join([]string{entity.Name, config.presenterName}, " ") )  , "go"}, config.stringSeparator)}, config.pathSeparator)
}

func scaffoldUsecaseInteractorFileName(outputPath string, entity entity)(string){
	return strings.Join([]string{outputPath, config.scaffoldDirName, config.usecaseLayerName, config.interactorName, strings.Join([]string{ strcase.ToSnake( strings.Join([]string{entity.Name, config.interactorName}, " ") )  , "go"}, config.stringSeparator)}, config.pathSeparator)
}

func scaffoldInterfacePresenterFileName(outputPath string, entity entity)(string){
	return strings.Join([]string{outputPath, config.scaffoldDirName, config.interfaceLayerName, config.presenterName, strings.Join([]string{ strcase.ToSnake( strings.Join([]string{entity.Name, config.presenterName}, " ") )  , "go"}, config.stringSeparator)}, config.pathSeparator)
}

func scaffoldInterfaceRepositoryFileName(outputPath string, entity entity)(string){
	return strings.Join([]string{outputPath, config.scaffoldDirName, config.interfaceLayerName, config.repositoryName, strings.Join([]string{ strcase.ToSnake( strings.Join([]string{entity.Name, config.repositoryName}, " ") )  , "go"}, config.stringSeparator)}, config.pathSeparator)
}


func structId(entity entity)(string){
	return strcase.ToCamel(strings.Join([]string{entity.Name, config.scaffoldName, "struct"}, config.stringSeparator))
}

func interactorStructId(entity entity)(string){
	return strcase.ToLowerCamel(strings.Join([]string{entity.Name, config.scaffoldName, config.interactorName}, config.stringSeparator))
}


func structFieldId(field field)(string){
	return strcase.ToLowerCamel(field.Name)
}

func interfaceId(entity entity)(string){
	return strcase.ToCamel(strings.Join([]string{entity.Name, config.scaffoldName, "interface"}, config.stringSeparator))
}

func getterId(field field)(string){
	return strcase.ToCamel(field.Name)
}

func tagId(field field)(string){
	return strcase.ToSnake(field.Name)
}

func setterId(field field)(string){
	return strcase.ToCamel(strings.Join([]string{config.setterVerb, field.Name}, config.stringSeparator))
}

func repositoryMethodName(entityMethod entityMethod)(string){
	return strcase.ToCamel(entityMethod.Name)
}

func repositoryId(entity entity)(string){
	return strcase.ToCamel(strings.Join([]string{entity.Name, config.scaffoldName, config.repositoryName}, config.stringSeparator))
}

func interfaceRepositoryStructId(entity entity, driver string)(string){
	return strcase.ToLowerCamel(strings.Join([]string{driver, entity.Name, config.scaffoldName, config.repositoryName}, config.stringSeparator))
}

func interfaceRepositoryConstructorId(entity entity, driver string)(string){
	return strcase.ToCamel(strings.Join([]string{"new", driver, entity.Name, config.scaffoldName, config.repositoryName}, config.stringSeparator))
}

func presenterMethodName(entityMethod entityMethod)(string){
	return strcase.ToCamel(entityMethod.Name)
}

func presenterId(entity entity)(string){
	return strcase.ToCamel(strings.Join([]string{entity.Name, config.scaffoldName, config.presenterName}, config.stringSeparator))
}

func interactorMethodName(entityMethod entityMethod)(string){
	return strcase.ToCamel(entityMethod.Name)
}

func interactorId(entity entity)(string){
	return strcase.ToCamel(strings.Join([]string{entity.Name, config.scaffoldName, config.interactorName}, config.stringSeparator))
}

func interactorConstructorId(entity entity)(string){
	return strcase.ToCamel(strings.Join([]string{"new", entity.Name, config.scaffoldName, config.interactorName}, config.stringSeparator))
}