package entities

type Contacts struct {
	ID        int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	AccountID int64  `gorm:"index"` 
	Name      string `json:"name"`
	Email     string `json:"email"`
}

type ContactAPIResponse struct {
	Embedded struct {
		Contacts []struct {
			ID int64 `json:"id"`
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
