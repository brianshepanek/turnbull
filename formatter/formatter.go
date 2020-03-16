package formatter

import(
	"os"
	"strings"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/domain/model"
)

type formatter struct{
	config *config.Config
}

type Formatter interface{

	OutputDirectory() (string, error)
	OutputScaffoldDirectory() (string, error)

	OutputScaffoldDomainDirectory() (string, error)
	OutputScaffoldDomainEntityDirectory() (string, error)
	OutputScaffoldDomainEntityDirectoryImportPath() (string, error)
	OutputScaffoldDomainEntityFile(entity model.Entity) (string, error)

	OutputScaffoldUsecaseDirectory() (string, error)
	OutputScaffoldUsecaseInteractorDirectory() (string, error)
	OutputScaffoldUsecaseInteractorDirectoryImportPath() (string, error)
	OutputScaffoldUsecaseInteractorFile(entity model.Entity) (string, error)
	OutputScaffoldUsecaseRepositoryDirectory() (string, error)
	OutputScaffoldUsecaseRepositoryDirectoryImportPath() (string, error)
	OutputScaffoldUsecaseRepositoryFile(entity model.Entity) (string, error)
	OutputScaffoldUsecasePresenterDirectory() (string, error)
	OutputScaffoldUsecasePresenterDirectoryImportPath() (string, error)
	OutputScaffoldUsecasePresenterFile(entity model.Entity) (string, error)

	OutputScaffoldInterfaceDirectory() (string, error)
	OutputScaffoldInterfaceControllerDirectory() (string, error)
	OutputScaffoldInterfaceRepositoryDirectory() (string, error)
	OutputScaffoldInterfacePresenterDirectory() (string, error)

	OutputScaffoldDomainEntityPackageName() (string, error)
	OutputScaffoldDomainEntityStructId(entity model.Entity) (string, error)
	OutputScaffoldDomainEntitySliceStructId(entity model.Entity) (string, error)
	OutputScaffoldDomainEntityInterfaceId(entity model.Entity) (string, error)
	OutputScaffoldDomainEntitySliceInterfaceId(entity model.Entity) (string, error)
	OutputScaffoldDomainEntityFieldId(field model.Field) (string,  error)
	OutputScaffoldDomainEntityGetterId(field model.Field) (string,  error)
	OutputScaffoldDomainEntitySetterId(field model.Field) (string,  error)
	OutputScaffoldDomainEntitySetAllSetterId() (string,  error)
	OutputScaffoldDomainEntityLenId() (string,  error)
	OutputScaffoldDomainEntityAppendId() (string,  error)
	OutputScaffoldDomainEntityElementsId() (string,  error)
	OutputScaffoldDomainEntityJSONTagId(field model.Field) (string,  error)
	
	OutputScaffoldDomainEntityInterfaceConstructorFunctionId(entity model.Entity) (string, error)
	OutputScaffoldDomainEntitySliceInterfaceConstructorFunctionId(entity model.Entity) (string, error)

	OutputScaffoldUsecaseRepositoryPackageName() (string, error)
	OutputScaffoldUsecaseRepositoryInterfaceId(entity model.Entity) (string, error)
	OutputScaffoldUsecaseRepositoryInterfaceMethodId(method model.Method) (string, error)

	OutputScaffoldUsecasePresenterPackageName() (string, error)
	OutputScaffoldUsecasePresenterInterfaceId(entity model.Entity) (string, error)
	OutputScaffoldUsecasePresenterInterfaceMethodId(method model.Method) (string, error)

	OutputScaffoldUsecaseInteractorPackageName() (string, error)
	OutputScaffoldUsecaseInteractorStructId(entity model.Entity) (string, error)
	OutputScaffoldUsecaseInteractorInterfaceId(entity model.Entity) (string, error)
	OutputScaffoldUsecaseInteractorInterfaceMethodId(method model.Method) (string, error)
	OutputScaffoldUsecaseInteractorInterfaceConstructorFunctionId(entity model.Entity) (string, error)

	OutputScaffoldInterfaceRepositoryPackageName() (string, error)
	OutputScaffoldInterfaceRepositoryStructId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfaceRepositoryConstructorFunctionId(driver string, entity model.Entity) (string, error)

	OutputScaffoldInterfacePresenterPackageName() (string, error)
	OutputScaffoldInterfacePresenterStructId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfacePresenterInterfaceId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfacePresenterConstructorFunctionId(driver string, entity model.Entity) (string, error)

	OutputScaffoldInterfaceControllerPackageName() (string, error)
	OutputScaffoldInterfaceControllerStructId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfaceControllerConstructorFunctionId(driver string, entity model.Entity) (string, error)

}

func New(config *config.Config) Formatter {
	return &formatter{
		config : config,
	}
}

// Output

func (formatter *formatter) OutputDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath}, formatter.config.PathSeparator), nil
}

// Domain

func (formatter *formatter) OutputScaffoldDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Scaffold.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldDomainDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Scaffold.Name, formatter.config.Layers.Domain.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Scaffold.Name, formatter.config.Layers.Domain.Name, formatter.config.Layers.Domain.Entity.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityDirectoryImportPath() (string, error) {
	path, err  := formatter.OutputScaffoldDomainEntityDirectory()
	if err != nil {
		return "", nil
	}
	return strings.TrimLeft(strings.Replace(path, strings.Join([]string{os.Getenv("GOPATH"), formatter.config.WorkspaceSourceDirName}, formatter.config.PathSeparator), "", 1), formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityFile(entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldDomainEntityDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(entity.Name), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}

// Usecase

func (formatter *formatter) OutputScaffoldUsecaseDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Scaffold.Name, formatter.config.Layers.Usecase.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecaseInteractorDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Scaffold.Name, formatter.config.Layers.Usecase.Name, formatter.config.Layers.Usecase.Interactor.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecaseInteractorDirectoryImportPath() (string, error) {
	path, err  := formatter.OutputScaffoldUsecaseInteractorDirectory()
	if err != nil {
		return "", nil
	}
	return strings.TrimLeft(strings.Replace(path, strings.Join([]string{os.Getenv("GOPATH"), formatter.config.WorkspaceSourceDirName}, formatter.config.PathSeparator), "", 1), formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecaseInteractorFile(entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldUsecaseInteractorDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name, formatter.config.Layers.Usecase.Interactor.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecaseRepositoryDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Scaffold.Name, formatter.config.Layers.Usecase.Name, formatter.config.Layers.Usecase.Repository.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecaseRepositoryDirectoryImportPath() (string, error) {
	path, err  := formatter.OutputScaffoldUsecaseRepositoryDirectory()
	if err != nil {
		return "", nil
	}
	return strings.TrimLeft(strings.Replace(path, strings.Join([]string{os.Getenv("GOPATH"), formatter.config.WorkspaceSourceDirName}, formatter.config.PathSeparator), "", 1), formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecaseRepositoryFile(entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldUsecaseRepositoryDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name, formatter.config.Layers.Usecase.Repository.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecasePresenterDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Scaffold.Name, formatter.config.Layers.Usecase.Name, formatter.config.Layers.Usecase.Presenter.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecasePresenterDirectoryImportPath() (string, error) {
	path, err  := formatter.OutputScaffoldUsecasePresenterDirectory()
	if err != nil {
		return "", nil
	}
	return strings.TrimLeft(strings.Replace(path, strings.Join([]string{os.Getenv("GOPATH"), formatter.config.WorkspaceSourceDirName}, formatter.config.PathSeparator), "", 1), formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecasePresenterFile(entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldUsecasePresenterDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name, formatter.config.Layers.Usecase.Presenter.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}

// Interface

func (formatter *formatter) OutputScaffoldInterfaceDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Scaffold.Name, formatter.config.Layers.Interface.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldInterfaceControllerDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Scaffold.Name, formatter.config.Layers.Interface.Name, formatter.config.Layers.Interface.Controller.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldInterfaceRepositoryDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Scaffold.Name, formatter.config.Layers.Interface.Name, formatter.config.Layers.Interface.Repository.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldInterfacePresenterDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Scaffold.Name, formatter.config.Layers.Interface.Name, formatter.config.Layers.Interface.Presenter.Name}, formatter.config.PathSeparator), nil
}

// Domain Entity

func (formatter *formatter) OutputScaffoldDomainEntityPackageName() (string, error) {
	return strcase.ToSnake(formatter.config.Layers.Domain.Entity.Name), nil
}


func (formatter *formatter) OutputScaffoldDomainEntityStructId(entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{entity.Name, formatter.config.Scaffold.Name, "struct"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntitySliceStructId(entity model.Entity) (string, error) {
	pluralize := pluralize.NewClient()
	return strcase.ToCamel(strings.Join([]string{pluralize.Plural(entity.Name), formatter.config.Scaffold.Name, "struct"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityInterfaceId(entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{entity.Name, formatter.config.Scaffold.Name, "interface"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntitySliceInterfaceId(entity model.Entity) (string, error) {
	pluralize := pluralize.NewClient()
	return strcase.ToCamel(strings.Join([]string{pluralize.Plural(entity.Name), formatter.config.Scaffold.Name, "interface"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityGetterId(field model.Field) (string, error) {
	return strcase.ToCamel(field.Name), nil
}

func (formatter *formatter) OutputScaffoldDomainEntitySetterId(field model.Field) (string, error) {
	return strcase.ToCamel(strings.Join([]string{formatter.config.Layers.Domain.SetterVerb, field.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntitySetAllSetterId() (string, error) {
	return strcase.ToCamel(strings.Join([]string{formatter.config.Layers.Domain.SetterVerb, "all"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityLenId() (string, error) {
	return strcase.ToCamel(strings.Join([]string{formatter.config.Layers.Domain.LenName}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityAppendId() (string, error) {
	return strcase.ToCamel(strings.Join([]string{formatter.config.Layers.Domain.AppendName}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityElementsId() (string, error) {
	return strcase.ToCamel(strings.Join([]string{formatter.config.Layers.Domain.ElementsName}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityFieldId(field model.Field) (string, error) {
	return strcase.ToLowerCamel(field.Name), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityInterfaceConstructorFunctionId(entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"new", entity.Name, formatter.config.Scaffold.Name, "struct"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntitySliceInterfaceConstructorFunctionId(entity model.Entity) (string, error) {
	pluralize := pluralize.NewClient()
	return strcase.ToCamel(strings.Join([]string{"new", pluralize.Plural(entity.Name), formatter.config.Scaffold.Name, "struct"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityJSONTagId(field model.Field) (string, error) {
	return strcase.ToSnake(field.Name), nil
}


// Usecase

func (formatter *formatter) OutputScaffoldUsecaseRepositoryPackageName() (string, error) {
	return strcase.ToSnake(formatter.config.Layers.Usecase.Repository.Name), nil
}

func (formatter *formatter) OutputScaffoldUsecaseRepositoryInterfaceId(entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{entity.Name, formatter.config.Scaffold.Name, formatter.config.Layers.Usecase.Repository.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldUsecaseRepositoryInterfaceMethodId(method model.Method) (string, error) {
	return strcase.ToCamel(method.Name), nil
}

func (formatter *formatter) OutputScaffoldUsecasePresenterPackageName() (string, error) {
	return strcase.ToSnake(formatter.config.Layers.Usecase.Presenter.Name), nil
}

func (formatter *formatter) OutputScaffoldUsecasePresenterInterfaceId(entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{entity.Name, formatter.config.Scaffold.Name, formatter.config.Layers.Usecase.Presenter.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldUsecasePresenterInterfaceMethodId(method model.Method) (string, error) {
	return strcase.ToCamel(method.Name), nil
}

func (formatter *formatter) OutputScaffoldUsecaseInteractorPackageName() (string, error) {
	return strcase.ToSnake(formatter.config.Layers.Usecase.Interactor.Name), nil
}

func (formatter *formatter) OutputScaffoldUsecaseInteractorStructId(entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{entity.Name, formatter.config.Scaffold.Name, formatter.config.Layers.Usecase.Interactor.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldUsecaseInteractorInterfaceId(entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{entity.Name, formatter.config.Scaffold.Name, formatter.config.Layers.Usecase.Interactor.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldUsecaseInteractorInterfaceMethodId(method model.Method) (string, error) {
	return strcase.ToCamel(method.Name), nil
}

func (formatter *formatter) OutputScaffoldUsecaseInteractorInterfaceConstructorFunctionId(entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"new", entity.Name, formatter.config.Scaffold.Name, formatter.config.Layers.Usecase.Interactor.Name}, formatter.config.StringSeparator)), nil
}

// Interface

func (formatter *formatter) OutputScaffoldInterfaceRepositoryPackageName() (string, error) {
	return strcase.ToSnake(formatter.config.Layers.Interface.Repository.Name), nil
}

func (formatter *formatter) OutputScaffoldInterfaceRepositoryStructId(driver string, entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{driver, entity.Name, formatter.config.Scaffold.Name, formatter.config.Layers.Interface.Repository.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfaceRepositoryConstructorFunctionId(driver string, entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"new", driver, entity.Name, formatter.config.Scaffold.Name, formatter.config.Layers.Interface.Repository.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfacePresenterPackageName() (string, error) {
	return strcase.ToSnake(formatter.config.Layers.Interface.Presenter.Name), nil
}

func (formatter *formatter) OutputScaffoldInterfacePresenterStructId(driver string, entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{driver, entity.Name, formatter.config.Scaffold.Name, formatter.config.Layers.Interface.Presenter.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfacePresenterInterfaceId(driver string, entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{driver, entity.Name, formatter.config.Scaffold.Name, formatter.config.Layers.Interface.Presenter.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfacePresenterConstructorFunctionId(driver string, entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"new", driver, entity.Name, formatter.config.Scaffold.Name, formatter.config.Layers.Interface.Presenter.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfaceControllerPackageName() (string, error) {
	return strcase.ToSnake(formatter.config.Layers.Interface.Controller.Name), nil
}

func (formatter *formatter) OutputScaffoldInterfaceControllerStructId(driver string, entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{driver, entity.Name, formatter.config.Scaffold.Name, formatter.config.Layers.Interface.Controller.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfaceControllerConstructorFunctionId(driver string, entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"new", driver, entity.Name, formatter.config.Scaffold.Name, formatter.config.Layers.Interface.Controller.Name}, formatter.config.StringSeparator)), nil
}