package generator

import(
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/generator"
	"github.com/brianshepanek/turnbull/domain/model"
)

const(
	testConfigPath = "/go/src/github.com/brianshepanek/turnbull/_testing/config"
	testOutputPath = "/go/src/github.com/brianshepanek/turnbull/_testing/output"
)

var(


	testGenerator generator.UsecaseGenerator

	// Test Entity
	testEntity  = model.Entity{
		Name : "foo",
		JSON : true,
		Fields : []model.Field{
			model.Field{
				Name : "string",
				Type : "string",
			},
			model.Field{
				Name : "int",
				Type : "int",
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


func init(){
	conf, _ := config.New(testConfigPath, testOutputPath)
	formatter := formatter.New(conf)
	testGenerator = NewGenerator(conf, formatter)
}