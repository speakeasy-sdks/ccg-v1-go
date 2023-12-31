// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Security struct {
	HTTPCCG string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *Security) GetHTTPCCG() string {
	if o == nil {
		return ""
	}
	return o.HTTPCCG
}
