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

type Method struct {
	Type string
	Repository RepositoryMethod
	Presenter PresenterMethod
}

type RepositoryMethod struct {
	Name string
	Arguments []model.Field
	ReturnValues []model.Field
}

type PresenterMethod struct {
	Name string
	Arguments []model.Field
	ReturnValues []model.Field
}



type Config struct {
	AbsOutputPath string
	WorkspaceSourceDirName string
	PathSeparator string
	StringSeparator string
	Scaffold Scaffold
	Layers Layers
	Methods []Method
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
		Layers : Layers {
			Domain : Domain {
				Name : "domain",
				SetterVerb : "set",
				LenName : "len",
				AppendName : "append",
				ElementsName : "elements",
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
		Methods : []Method {
			Method {
				Type : "add",
				Presenter : PresenterMethod {
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
							Op : "*",
							Type : "self",
						},
						model.Field {
							Name : "err",
							Type : "error",
						},
					},
				},
				Repository : RepositoryMethod {
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
			Method {
				Type : "browse",
				Presenter : PresenterMethod {
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
							Op : "*",
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
				Repository : RepositoryMethod {
					Arguments : []model.Field {
						model.Field {
							Name : "ctx",
							Package : "context",
							Type : "Context",
						},
						model.Field {
							Name : "query",
							Type : "interface{}",
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
			Method {
				Type : "read",
				Presenter : PresenterMethod {
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
							Op : "*",
							Type : "self",
						},
						model.Field {
							Name : "err",
							Type : "error",
						},
					},
				},
				Repository : RepositoryMethod {
					Arguments : []model.Field {
						model.Field {
							Name : "ctx",
							Package : "context",
							Type : "Context",
						},
						model.Field {
							Name : "query",
							Type : "interface{}",
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
			Method {
				Type : "delete",
				Presenter : PresenterMethod {
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
							Op : "*",
							Type : "self",
						},
						model.Field {
							Name : "err",
							Type : "error",
						},
					},
				},
				Repository : RepositoryMethod {
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
			Method {
				Type : "edit",
				Presenter : PresenterMethod {
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
							Op : "*",
							Type : "self",
						},
						model.Field {
							Name : "err",
							Type : "error",
						},
					},
				},
				Repository : RepositoryMethod {
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
			Method {
				Type : "count",
				Presenter : PresenterMethod {
					Arguments : []model.Field {
						model.Field {
							Name : "ctx",
							Package : "context",
							Type : "Context",
						},
						model.Field {
							Name : "req",
							Op : "*",
							Type : "int",
						},
					},
					ReturnValues : []model.Field {
						model.Field {
							Name : "resp",
							Op : "*",
							Type : "int",
						},
						model.Field {
							Name : "err",
							Type : "error",
						},
					},
				},
				Repository : RepositoryMethod {
					Arguments : []model.Field {
						model.Field {
							Name : "ctx",
							Package : "context",
							Type : "Context",
						},
						model.Field {
							Name : "query",
							Type : "interface{}",
						},
						model.Field {
							Name : "req",
							Op : "*",
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

func (conf *Config) Method(method model.Method) (Method, error){
	var resp Method
	for _, confMethod := range conf.Methods {
		if confMethod.Type == method.Type {
			resp = confMethod
		}
	}
	return resp, nil
}