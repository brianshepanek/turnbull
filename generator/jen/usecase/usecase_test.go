package usecase

import(
	"github.com/brianshepanek/turnbull/domain/model"
)
const(
	testConfigPath = "/go/src/github.com/brianshepanek/turnbull/_testing/config"
	testOutputPath = "/go/src/github.com/brianshepanek/turnbull/_testing/output"
)

var(

	// Test Entity
	testEntity  = model.Entity{
		Name : "foo",
		Fields : []model.Field{
			model.Field{
				Primary : true,
				Name : "id",
				Type : "int64",
			},
			model.Field{
				Name : "title",
				Type : "string",
			},
		},
		Methods : []model.Method {
			model.Method{
				Name : "add",
				Type : "add",
			},
			model.Method{
				Name : "read",
				Type : "read",
			},
			model.Method{
				Name : "browse",
				Type : "browse",
			},
			model.Method{
				Name : "delete",
				Type : "delete",
			},
		},
	}
	testField  = model.Field{
		Name : "bar",
		Type : "string",
	}
	testMethod  = model.Method{
		Name : "add",
		Type : "add",
	}
)