package responses

type ToCharc struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			DeleteMc                 bool        `json:"Delete_mc"`
			UpdateMc                 bool        `json:"Update_mc"`
			ClassInternalID          string      `json:"ClassInternalID"`
			CharcInternalID          string      `json:"CharcInternalID"`
			CharcPositionNumber      string      `json:"CharcPositionNumber"`
			Characteristic           string      `json:"Characteristic"`
			AncestorClassInternalID  string      `json:"AncestorClassInternalID"`
			OriginalCharcInternalID  string      `json:"OriginalCharcInternalID"`
			ChangeNumber             string      `json:"ChangeNumber"`
			CharcIsPrintRelevant     string      `json:"CharcIsPrintRelevant"`
			CharcIsSearchRelevant    string      `json:"CharcIsSearchRelevant"`
			CharcIsDisplayRelevant   string      `json:"CharcIsDisplayRelevant"`
			ValidityStartDate        string      `json:"ValidityStartDate"`
			ValidityEndDate          string      `json:"ValidityEndDate"`
			KeyDate                  string      `json:"KeyDate"`
			ClassLastChangedDateTime string      `json:"ClassLastChangedDateTime"`
		} `json:"results"`
	} `json:"d"`
}
