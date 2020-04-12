package helper

import(
	"testing"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/formatter"
)


const(
	testConfigPath = "./config"
	testOutputPath = "./output"
)

var(
	testHelperGenerator Generator
	testEntities = []model.Entity{
		model.Entity{
			Name : "a",
			Fields : []model.Field{
				model.Field{
					Type : "b",
					Embedded : true,
				},
				model.Field{
					Name : "aa",
					Type : "string",
				},
			},
		},
		model.Entity{
			Name : "b",
			Fields : []model.Field{
				model.Field{
					Type : "c",
					Embedded : true,
				},
				model.Field{
					Name : "bb",
					Type : "string",
				},
			},
		},
		model.Entity{
			Name : "c",
			Fields : []model.Field{
				model.Field{
					Type : "d",
					Embedded : true,
				},
				model.Field{
					Name : "cc",
					Type : "string",
				},
			},
		},
		model.Entity{
			Name : "d",
			Fields : []model.Field{
				model.Field{
					Type : "e",
					Embedded : true,
				},
				model.Field{
					Name : "dd",
					Type : "string",
				},
			},
		},
		model.Entity{
			Name : "e",
			Fields : []model.Field{
				model.Field{
					Name : "primary",
					Type : "int",
					Primary : true,
				},
			},
		},
	}

)


func init(){

	conf, _ := config.New(testConfigPath, testOutputPath)
	formatter := formatter.New(conf)
	testHelperGenerator = New(formatter)

	err := testHelperGenerator.FormatDomainEntities(&testEntities)

	// Return
	if err != nil {
		
	}

}	

// Test Primary Field
func TestPrimaryField(t *testing.T){

	for _, testEntity := range testEntities {
		
		// Build
		fieldPointer, err := testHelperGenerator.PrimaryField(testEntity)
		
		// Return
		if err != nil {
			t.Errorf(`PrimaryField(` + testEntity.Name + `) failed with error %v`, err)
		}
		if fieldPointer == nil {
			t.Errorf(`PrimaryField(` + testEntity.Name + `) failed;`)
		} else {
			field := *fieldPointer
			if field.Name != "primary" {
				t.Errorf(`PrimaryField(` + testEntity.Name + `) failed; want "%s", got "%s"`, "primary", field.Name)
			}
		}

	}

}

// Test Expanded Fields
func TestExpandedFields(t *testing.T){

	for testEntityKey, testEntity := range testEntities {
		
		// Build
		fieldsPointer, err := testHelperGenerator.ExpandedFields(testEntity)
		
		// Return
		if err != nil {
			t.Errorf(`ExpandedFields(` + testEntity.Name + `) failed with error %v`, err)
		}
		if fieldsPointer == nil {
			t.Errorf(`ExpandedFields(` + testEntity.Name + `) failed;`)
		} else {
			fields := *fieldsPointer
			if len(fields) != len(testEntities) - testEntityKey {
				t.Errorf(`ExpandedFields(` + testEntity.Name + `) failed; want "%s", got "%s"`, len(testEntities) - testEntityKey, len(fields))
			}
		}

	}

}