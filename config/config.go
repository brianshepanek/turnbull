package config

import(
	"path/filepath"
	"github.com/brianshepanek/turnbull/domain/model"
)

type Scaffold struct {
	Name string
}

type Layers struct {
	Domain Domain
	Usecase Usecase
	Interface Interface
}

type Domain struct {
	Name string
	Entity Entity
	SetterVerb string
	LenName string
	AppendName string
	ElementsName string
	ToPrimaryName string
}

type Entity struct {
	Name string
}

type Usecase struct {
	Name string
	Interactor UsecaseInteractor
	Repository UsecaseRepository
	Presenter UsecasePresenter
}

type UsecaseInteractor struct {
	Name string
}

type UsecaseRepository struct {
	Name string
}

type UsecasePresenter struct {
	Name string
}

type Interface struct {
	Name string
	Controller InterfaceController
	Repository InterfaceRepository
	Presenter InterfacePresenter
}

type InterfaceController struct {
	Name string
}

type InterfaceRepository struct {
	Name string
}

type InterfacePresenter struct {
	Name string
}



type Registry struct {
	Name string
}


type Config struct {
	AbsOutputPath string
	WorkspaceSourceDirName string
	PathSeparator string
	StringSeparator string
	Scaffold Scaffold
	Layers Layers
	Methods []model.Method
	Registry Registry
	// EntityName string
	// entitiesDirName string
	// interactorName string
	// presenterName string
	// repositoryName string
	// controllerName string
	// scaffoldName string
	// scaffoldDirName string
	// entities []Entity
	// methods []method
	// absOutputPath string
	// workspaceSourceDirName string
}

func New(configPath string, outputPath string) (*Config, error) {

	// Abs Output Path
	absOutputPath, err := filepath.Abs(outputPath)
	if err != nil {
		return nil, err
	}

	return &Config{
		PathSeparator : "/",
		StringSeparator : ".",
		AbsOutputPath : absOutputPath,
		WorkspaceSourceDirName : "src",
		Scaffold : Scaffold {
			Name : "scaffold",
		},
		Registry : Registry {
			Name : "registry",
		},
		Layers : Layers {
			Domain : Domain {
				Name : "domain",
				SetterVerb : "set",
				LenName : "len",
				AppendName : "append",
				ElementsName : "elements",
				ToPrimaryName : "to_primary",
				Entity : Entity {
					Name : "entity",
				},
			},
			Usecase : Usecase {
				Name : "usecase",
				Interactor : UsecaseInteractor {
					Name : "interactor",
				},
				Repository : UsecaseRepository {
					Name : "repository",
				},
				Presenter : UsecasePresenter {
					Name : "presenter",
				},
			},
			Interface : Interface {
				Name : "interface",
				Controller : InterfaceController {
					Name : "controller",
				},
				Repository : InterfaceRepository {
					Name : "repository",
				},
				Presenter : InterfacePresenter {
					Name : "presenter",
				},
			},
		},
		Methods : []model.Method {
			model.Method {
				Type : "add",
				Presenter : model.PresenterMethod {
					Arguments : []model.Field {
						model.Field {
							Name : "ctx",
							Package : "context",
							Type : "Context",
						},
						model.Field {
							Name : "req",
							Type : "self",
						},
					},
					ReturnValues : []model.Field {
						model.Field {
							Name : "req",
							Type : "self",
						},
						model.Field {
							Name : "err",
							Type : "error",
						},
					},
				},
				Repository : model.RepositoryMethod {
					Arguments : []model.Field {
						model.Field {
							Name : "ctx",
							Package : "context",
							Type : "Context",
						},
						model.Field {
							Name : "req",
							Type : "self",
						},
					},
					ReturnValues : []model.Field {
						model.Field {
							Name : "err",
							Type : "error",
						},
					},
				},
			},
			model.Method {
				Type : "browse",
				Presenter : model.PresenterMethod {
					Arguments : []model.Field {
						model.Field {
							Name : "ctx",
							Package : "context",
							Type : "Context",
						},
						model.Field {
							Slice : true,
							Name : "req",
							Type : "self",
						},
					},
					ReturnValues : []model.Field {
						model.Field {
							Slice : true,
							Name : "resp",
							Type : "self",
						},
						model.Field {
							Name : "err",
							Type : "error",
						},
					},
				},
				Repository : model.RepositoryMethod {
					Arguments : []model.Field {
						model.Field {
							Name : "ctx",
							Package : "context",
							Type : "Context",
						},
						model.Field {
							Slice : true,
							Name : "req",
							Type : "self",
						},
					},
					ReturnValues : []model.Field {
						model.Field {
							Name : "err",
							Type : "error",
						},
					},
				},
			},
			model.Method {
				Type : "read",
				Presenter : model.PresenterMethod {
					Arguments : []model.Field {
						model.Field {
							Name : "ctx",
							Package : "context",
							Type : "Context",
						},
						model.Field {
							Name : "req",
							Type : "self",
						},
					},
					ReturnValues : []model.Field {
						model.Field {
							Name : "resp",
							Type : "self",
						},
						model.Field {
							Name : "err",
							Type : "error",
						},
					},
				},
				Repository : model.RepositoryMethod {
					Arguments : []model.Field {
						model.Field {
							Name : "ctx",
							Package : "context",
							Type : "Context",
						},
						model.Field {
							Name : "id",
							Type : "primary",
						},
						model.Field {
							Name : "req",
							Type : "self",
						},
					},
					ReturnValues : []model.Field {
						model.Field {
							Name : "err",
							Type : "error",
						},
					},
				},
			},
			model.Method {
				Type : "delete",
				Presenter : model.PresenterMethod {
					Arguments : []model.Field {
						model.Field {
							Name : "ctx",
							Package : "context",
							Type : "Context",
						},
						model.Field {
							Name : "req",
							Type : "self",
						},
					},
					ReturnValues : []model.Field {
						model.Field {
							Name : "resp",
							Type : "self",
						},
						model.Field {
							Name : "err",
							Type : "error",
						},
					},
				},
				Repository : model.RepositoryMethod {
					Arguments : []model.Field {
						model.Field {
							Name : "ctx",
							Package : "context",
							Type : "Context",
						},
						model.Field {
							Name : "id",
							Type : "primary",
						},
						model.Field {
							Name : "req",
							Type : "self",
						},
					},
					ReturnValues : []model.Field {
						model.Field {
							Name : "err",
							Type : "error",
						},
					},
				},
			},
			model.Method {
				Type : "edit",
				Presenter : model.PresenterMethod {
					Arguments : []model.Field {
						model.Field {
							Name : "ctx",
							Package : "context",
							Type : "Context",
						},
						model.Field {
							Name : "req",
							Type : "self",
						},
					},
					ReturnValues : []model.Field {
						model.Field {
							Name : "req",
							Type : "self",
						},
						model.Field {
							Name : "err",
							Type : "error",
						},
					},
				},
				Repository : model.RepositoryMethod {
					Arguments : []model.Field {
						model.Field {
							Name : "ctx",
							Package : "context",
							Type : "Context",
						},
						model.Field {
							Name : "id",
							Type : "primary",
						},
						model.Field {
							Name : "req",
							Type : "self",
						},
					},
					ReturnValues : []model.Field {
						model.Field {
							Name : "err",
							Type : "error",
						},
					},
				},
			},
			model.Method {
				Type : "count",
				Presenter : model.PresenterMethod {
					Arguments : []model.Field {
						model.Field {
							Name : "ctx",
							Package : "context",
							Type : "Context",
						},
						model.Field {
							Name : "req",
							Type : "int",
						},
					},
					ReturnValues : []model.Field {
						model.Field {
							Name : "resp",
							Type : "int",
						},
						model.Field {
							Name : "err",
							Type : "error",
						},
					},
				},
				Repository : model.RepositoryMethod {
					Arguments : []model.Field {
						model.Field {
							Name : "ctx",
							Package : "context",
							Type : "Context",
						},
						model.Field {
							Name : "req",
							Type : "int",
						},
					},
					ReturnValues : []model.Field {
						model.Field {
							Name : "err",
							Type : "error",
						},
					},
				},
			},
		},
	}, nil
}

func (conf *Config) Method(method model.Method) (model.Method, error){
	for _, confMethod := range conf.Methods {
		if confMethod.Type == method.Type {
			return confMethod, nil
		}
	}
	return method, nil
}