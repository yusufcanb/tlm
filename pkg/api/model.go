package api

type pullModelRequestPayload struct {
	Name     string `json:"name"`
	Insecure bool   `json:"insecure"`
	Stream   bool   `json:"stream"`
}
