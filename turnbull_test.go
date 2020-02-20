package turnbull

import(
	"testing"
	"reflect"
)

var testConfigPath = "./testing/config"
var testOutputPath = "./testing/output"

// Test Build
func TestBuild(t *testing.T){

	// Build
	err := Build(testConfigPath, testOutputPath)

	// Return
	if err != nil {
		t.Errorf(`loadConfig("` + testConfigPath + `") failed with error %v`, err)
	}
}

// Test Load Config
func TestLoadConfig(t *testing.T){

	// Test Data
	testConfig := &config{
		entities : []entity{
			
			entity{
				Name : "bar",
				Fields : []field{
					field{
						Name : "title",
						Package : "",
						Type : "string",
					},
				},
			},
			entity{
				Name : "foo",
				Fields : []field{
					field{
						Name : "title",
						Package : "",
						Type : "string",
					},
				},
			},
		},
	}

	// Load Config
	configPointer, err := loadConfig(testConfigPath)

	// Return
	if err != nil {
		t.Errorf(`loadConfig("` + testConfigPath + `") failed with error %v`, err)
	} else {
		if configPointer == nil {
			t.Errorf(`loadConfig("` + testConfigPath + `") failed, configPointer is nil`)
		} else {
			if !reflect.DeepEqual(configPointer, testConfig){
				t.Errorf(`loadConfig("` + testConfigPath + `") failed, expected %v, got %v`, testConfig, configPointer)
			} else {
	
			}
		}
		
	}
}

// Test Parse Entity
func TestParseEntityConfig(t *testing.T){

	// Test Data
	testEntityReq := []byte(`
		name = "foo"
		[[fields]]
			name = "title"
			package = ""
			type = "string" 
		[[fields]]
			name = "subtitle"
			package = ""
			type = "string"	 
	`)

	testEntityResp := &entity{
		Name : "foo",
		Fields : []field{
			field{
				Name : "title",
				Package : "",
				Type : "string",
			},
			field{
				Name : "subtitle",
				Package : "",
				Type : "string",
			},
		},
	}

	// Parse Entity Config
	entityPointer, err := parseEntityConfig(testEntityReq)

	// Return
	if err != nil {
		t.Errorf(`parseEntityConfig("` + string(testEntityReq) + `") failed with error %v`, err)
	} else {
		if entityPointer == nil {
			t.Errorf(`parseEntityConfig("` + string(testEntityReq) + `") failed, entityPointer is nil`)
		} else {
			if !reflect.DeepEqual(entityPointer, testEntityResp){
				t.Errorf(`parseEntityConfig("` + string(testEntityReq) + `") failed, expected %v, got %v`, testEntityResp, entityPointer)
			} else {
	
			}
		}
		
	}

}