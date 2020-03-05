package turnbull

import(
	"os"
	"io/ioutil"
	"testing"
	"reflect"
)

// Test Data
var testFieldReq = field{
	Name : "title",
	Package : "",
	Type : "string",
}

// Test Data
var testEntityReq = entity{
	Name : "foo",
	JSON : true,
	Fields : []field{
		field{
			Name : "title",
			Package : "",
			Type : "string",
		},
		field{
			Name : "subtitle",
			Package : "",
			Slice : true,
			Type : "string",
		},
	},
}


// Test Build Entity
func TestBuildEntity(t *testing.T){

	// Test Data
	entityFile, _ := os.Open("./testing/test_output/foo.go")
	defer entityFile.Close()
	entityFileBytes, _ := ioutil.ReadAll(entityFile)

	testEntityResp := string(entityFileBytes)

	// Parse Entity Config
	entityFilePointer, err := buildScaffoldEntityFile(testEntityReq)

	// Return
	if err != nil {
		t.Errorf(`buildEbuildScaffoldEntityFilentity() failed with error %v`, err)
	} else {
		if entityFilePointer == nil {
			t.Errorf(`buildScaffoldEntityFile() failed, entityFilePointer is nil`)
		} else {
			if !reflect.DeepEqual(entityFilePointer.GoString(), testEntityResp){
				t.Errorf(`buildScaffoldEntityFile() failed, expected %v, got %v`, testEntityResp, entityFilePointer.GoString())
			} else {
	
			}
		}
		
	}

}

// Test Entity to Struct
func TestEntityToStruct(t *testing.T){

	// Parse Entity Config
	structPointer, err := entityToStruct(testEntityReq)

	// Return
	if err != nil {
		t.Errorf(`entityToStruct() failed with error %v`, err)
	} else {
		if structPointer == nil {
			t.Errorf(`entityToStruct() failed, entityFilePointer is nil`)
		} else {
			
		}
		
	}

}

// Test Field to Struct Field
func TestFieldToStructField(t *testing.T){

	// Parse Entity Config
	structFieldPointer, err := fieldToStructField(testEntityReq, testFieldReq)

	// Return
	if err != nil {
		t.Errorf(`fieldToStructField() failed with error %v`, err)
	} else {
		if structFieldPointer == nil {
			t.Errorf(`fieldToStructField() failed, entityFilePointer is nil`)
		} else {
			
		}
		
	}

}

// Test Entity to Interface
func TestEntityToInterface(t *testing.T){

	// Parse Entity Config
	interfacePointer, err := entityToInterface(testEntityReq)

	// Return
	if err != nil {
		t.Errorf(`entityToInterface() failed with error %v`, err)
	} else {
		if interfacePointer == nil {
			t.Errorf(`entityToInterface() failed, entityFilePointer is nil`)
		} else {
			
		}
		
	}

}

// Test Field to Interface Getter
func TestFieldToInterfaceGetter(t *testing.T){

	// Parse Entity Config
	interfaceGetterPointer, err := fieldToInterfaceGetter(testFieldReq)

	// Return
	if err != nil {
		t.Errorf(`interfaceGetterPointer() failed with error %v`, err)
	} else {
		if interfaceGetterPointer == nil {
			t.Errorf(`interfaceGetterPointer() failed, entityFilePointer is nil`)
		} else {
			
		}
		
	}

}

// Test Field to Interface Setter
func TestFieldToInterfaceSetter(t *testing.T){

	// Parse Entity Config
	interfaceSetterPointer, err := fieldToInterfaceSetter(testFieldReq)

	// Return
	if err != nil {
		t.Errorf(`fieldToInterfaceSetter() failed with error %v`, err)
	} else {
		if interfaceSetterPointer == nil {
			t.Errorf(`fieldToInterfaceSetter() failed, entityFilePointer is nil`)
		} else {
			
		}
		
	}

}

// Test Field to Getter
func TestFieldToGetter(t *testing.T){

	// Parse Entity Config
	getterPointer, err := fieldToGetter(testEntityReq, testFieldReq)

	// Return
	if err != nil {
		t.Errorf(`fieldToGetter() failed with error %v`, err)
	} else {
		if getterPointer == nil {
			t.Errorf(`fieldToGetter() failed, entityFilePointer is nil`)
		} else {
			
		}
		
	}

}

// Test Field to Setter
func TestFieldToSetter(t *testing.T){

	// Parse Entity Config
	setterPointer, err := fieldToSetter(testEntityReq, testFieldReq)

	// Return
	if err != nil {
		t.Errorf(`fieldToSetter() failed with error %v`, err)
	} else {
		if setterPointer == nil {
			t.Errorf(`fieldToSetter() failed, entityFilePointer is nil`)
		} else {
			
		}
		
	}

}