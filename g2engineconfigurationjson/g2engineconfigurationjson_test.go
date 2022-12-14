package g2engineconfigurationjson

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func testError(test *testing.T, err error) {
	if err != nil {
		assert.FailNow(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestBuildSimpleSystemConfigurationJson(test *testing.T) {
	_, err := BuildSimpleSystemConfigurationJson("postgresql://postgres:postgres@$10.0.0.1:5432/G2")
	testError(test, err)
}

func TestBuildSimpleSystemConfigurationJsonFailure(test *testing.T) {
	_, err := BuildSimpleSystemConfigurationJson("")
	if err == nil {
		assert.FailNow(test, "There should be an error.")
	}
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleBuildSimpleSystemConfigurationJson() {
	result, err := BuildSimpleSystemConfigurationJson("postgresql://username:password@hostname:5432/G2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"postgresql://username:password@hostname:5432:G2/"}}
}

func ExampleBuildSimpleSystemConfigurationJson_db2() {
	result, err := BuildSimpleSystemConfigurationJson("db2://username:password@hostname:50000/G2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"db2://username:password@G2"}}
}

func ExampleBuildSimpleSystemConfigurationJson_db2WithSchema() {
	result, err := BuildSimpleSystemConfigurationJson("db2://username:password@hostname:50000/G2/?schema=schemaname")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"db2://username:password@G2/?schema=schemaname"}}
}

func ExampleBuildSimpleSystemConfigurationJson_oci() {
	result, err := BuildSimpleSystemConfigurationJson("oci://username:password@hostname:1521/G2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"oci://username:password@G2"}}
}

func ExampleBuildSimpleSystemConfigurationJson_mssql() {
	result, err := BuildSimpleSystemConfigurationJson("mssql://username:password@hostname:1433/G2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"mssql://username:password@G2"}}
}

func ExampleBuildSimpleSystemConfigurationJson_mysql() {
	result, err := BuildSimpleSystemConfigurationJson("mysql://username:password@hostname:3306/G2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"mysql://username:password@hostname:3306/?schema=G2"}}
}

func ExampleBuildSimpleSystemConfigurationJson_postgresql() {
	result, err := BuildSimpleSystemConfigurationJson("postgresql://username:password@hostname:5432/G2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"postgresql://username:password@hostname:5432:G2/"}}
}

func ExampleBuildSimpleSystemConfigurationJson_postgresqlWithSchema() {
	result, err := BuildSimpleSystemConfigurationJson("postgresql://username:password@hostname:5432/G2/?schema=schemaname")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"postgresql://username:password@hostname:5432:G2/?schema=schemaname"}}
}
func ExampleBuildSimpleSystemConfigurationJson_sqlite() {
	result, err := BuildSimpleSystemConfigurationJson("sqlite3://na:na@/var/opt/senzing/sqlite/G2C.db")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"sqlite3://na:na@/var/opt/senzing/sqlite/G2C.db"}}
}
