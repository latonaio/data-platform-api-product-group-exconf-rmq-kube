package dpfm_api_input_reader

import (
	"data-platform-api-product-group-exconf-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToProductGroup() *requests.ProductGroup {
	data := sdc.ProductGroup
	return &requests.ProductGroup{
		ProductGroup: data.ProductGroup,
	}
}
