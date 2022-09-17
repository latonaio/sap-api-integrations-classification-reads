package responses

type ToClassDescription struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			UpdateMc                 bool        `json:"Update_mc"`
			ClassInternalID          string      `json:"ClassInternalID"`
			Language                 string      `json:"Language"`
			ClassDescription         string      `json:"ClassDescription"`
			KeyDate                  string      `json:"KeyDate"`
			ClassLastChangedDateTime string      `json:"ClassLastChangedDateTime"`
		} `json:"results"`
	} `json:"d"`
}

