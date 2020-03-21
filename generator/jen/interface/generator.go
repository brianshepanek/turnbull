package generator


import(
	"github.com/dave/jennifer/jen"
	"github.com/brianshepanek/turnbull/domain/model"
)

type RepositoryGenerator interface{
	File(entity model.Entity) (*jen.File, error)
	ScaffoldFile(entity model.Entity) (*jen.File, error)
}

type PresenterGenerator interface{
	File(entity model.Entity) (*jen.File, error)
	ScaffoldFile(entity model.Entity) (*jen.File, error)
}

type ControllerGenerator interface{
	File(entity model.Entity) (*jen.File, error)
	ScaffoldFile(entity model.Entity) (*jen.File, error)
}