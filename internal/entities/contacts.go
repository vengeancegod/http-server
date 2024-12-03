package entities

type Contacts struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ContactAPIResponse struct {
	Embedded struct {
		Contacts []struct {
			Name               string `json:"name"`
			CustomFieldsValues []struct {
				FieldName string `json:"field_name"`
				Values    []struct {
					Value string `json:"value"`
				} `json:"values"`
			} `json:"custom_fields_values"`
		} `json:"contacts"`
	} `json:"_embedded"`
}
