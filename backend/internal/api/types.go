package api

type CreateHelloWorldRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

type GetHelloWorldRequest struct {
	Name string `json:"name"`
}

type HelloWorldResponse struct {
	Message string `json:"message"`
}
