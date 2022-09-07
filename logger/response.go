package logger

type (
	apiError struct {
		Error apiMsg `json:"error"`
	}

	apiField struct {
		Name string `json:"name"`
	}

	apiMsg struct {
		Message string     `json:"message"`
		Fields  []apiField `json:"fields"`
	}
)
