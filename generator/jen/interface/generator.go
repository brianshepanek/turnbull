package generator


import(
	"github.com/dave/jennifer/jen"
	"github.com/brianshepanek/turnbull/domain/model"
)

type RepositoryGenerator interface{
	File(entity model.Entity) (*jen.File, error)
	ScaffoldFile(entity model.Entity) (*jen.File, error)
	EntityFile(entity model.Entity) (*jen.File, error)
}

type PresenterGenerator interface{
	File(entity model.Entity) (*jen.File, error)
	ScaffoldFile(entity model.Entity) (*jen.File, error)
}

type ControllerGenerator interface{
	AppFile(entities []model.Entity) (*jen.File, error)
	File(entity model.Entity) (*jen.File, error)
	ScaffoldFile(entity model.Entity) (*jen.File, error)
	EntityFile(entity model.Entity) (*jen.File, error)
}