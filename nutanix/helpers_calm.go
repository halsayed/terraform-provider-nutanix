package nutanix

import (
	"strconv"
	v3 "terraform-provider-nutanix/client/v3"
	"terraform-provider-nutanix/utils"
)

func setCalmEntityMetadata(v *v3.CalmMetadata) (map[string]interface{}, []interface{}) {
	metadata := make(map[string]interface{})
	metadata["kind"] = utils.StringValue(v.Kind)
	metadata["uuid"] = utils.StringValue(v.UUID)
	metadata["spec_version"] = strconv.Itoa(int(utils.Int64Value(v.SpecVersion)))
	metadata["name"] = utils.StringValue(v.Name)

	return metadata, flattenCategories(v.Categories)
}