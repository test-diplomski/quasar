package validators

import (
	"errors"
	"strings"

	pb "github.com/jtomic1/config-schema-service/proto"
	"github.com/xeipuuv/gojsonschema"
	"golang.org/x/mod/semver"
	"sigs.k8s.io/yaml"
)

func IsSchemaValid(schema string) (bool, error) {
	if schema == "" {
		return false, errors.New("schema cannot be empty")
	}
	schemaJson, err := yaml.YAMLToJSON([]byte(schema))
	if err != nil {
		return false, err
	}
	loader := gojsonschema.NewStringLoader(string(schemaJson))
	_, schemaErr := gojsonschema.NewSchema(loader)
	if schemaErr != nil {
		return false, schemaErr
	}
	return true, nil
}

func IsConfigurationValid(configuration string) (bool, error) {
	if configuration == "" {
		return false, errors.New("configuration cannot be empty")
	}
	return true, nil
}

func AreSchemaDetailsValid(schemaDetails *pb.ConfigSchemaDetails, isVersionRequired bool) (bool, error) {
	if schemaDetails == nil {
		return false, errors.New("schema details cannot be empty")
	} else if schemaDetails.GetSchemaName() == "" {
		return false, errors.New("schema name cannot be empty")
	} else if isVersionRequired && schemaDetails.GetVersion() == "" {
		return false, errors.New("schema version cannot be empty")
	} else if isVersionRequired && !semver.IsValid(schemaDetails.GetVersion()) {
		return false, errors.New("schema version must be a valid SemVer string with 'v' prefix")
	} else if strings.Contains(schemaDetails.GetSchemaName(), "/") || strings.Contains(schemaDetails.GetVersion(), "/") {
		return false, errors.New("schema details must not contain '/'")
	}
	return true, nil
}

func IsSaveSchemaRequestValid(saveRequest *pb.SaveConfigSchemaRequest) (bool, error) {
	schemaDetailsValid, schemaDetailsErr := AreSchemaDetailsValid(saveRequest.GetSchemaDetails(), true)
	if schemaDetailsErr != nil {
		return false, schemaDetailsErr
	}
	schemaValid, schemaErr := IsSchemaValid(saveRequest.GetSchema())
	if schemaErr != nil {
		return false, schemaErr
	}
	requestValid := schemaDetailsValid && schemaValid
	return requestValid, nil
}

func IsGetSchemaRequestValid(getRequest *pb.GetConfigSchemaRequest) (bool, error) {
	schemaDetailsValid, schemaDetailsErr := AreSchemaDetailsValid(getRequest.GetSchemaDetails(), true)
	if schemaDetailsErr != nil {
		return false, schemaDetailsErr
	}
	return schemaDetailsValid, nil
}

func IsDeleteSchemaRequestValid(deleteRequest *pb.DeleteConfigSchemaRequest) (bool, error) {
	schemaDetailsValid, schemaDetailsErr := AreSchemaDetailsValid(deleteRequest.GetSchemaDetails(), true)
	if schemaDetailsErr != nil {
		return false, schemaDetailsErr
	}
	return schemaDetailsValid, nil
}

func IsValidateConfigurationRequestValid(validateRequest *pb.ValidateConfigurationRequest) (bool, error) {
	schemaDetailsValid, schemaDetailsErr := AreSchemaDetailsValid(validateRequest.GetSchemaDetails(), true)
	if schemaDetailsErr != nil {
		return false, schemaDetailsErr
	}
	configurationValid, configurationErr := IsConfigurationValid(validateRequest.GetConfiguration())
	if configurationErr != nil {
		return false, configurationErr
	}
	requestValid := schemaDetailsValid && configurationValid
	return requestValid, nil
}

func IsGetConfigSchemaVersionsValid(versionsRequest *pb.ConfigSchemaVersionsRequest) (bool, error) {
	schemaDetailsValid, schemaDetailsErr := AreSchemaDetailsValid(versionsRequest.GetSchemaDetails(), false)
	if schemaDetailsErr != nil {
		return false, schemaDetailsErr
	}
	return schemaDetailsValid, nil
}
