package response

type defaultResponse struct {
	Status    int         `json:"status"`
	Message   interface{} `json:"message,omitempty"`
	Parameter interface{} `json:"parameter,omitempty"`
	Body      interface{} `json:"body,omitempty"`
}
