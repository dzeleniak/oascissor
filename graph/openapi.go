package graph

type OpenAPI struct {
	Components struct {
		Schemas map[string]interface{} `json:"schemas"`
	} `json:"components"`
}
