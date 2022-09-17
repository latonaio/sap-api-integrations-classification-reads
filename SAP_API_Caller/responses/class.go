package responses

type Class struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			DeleteMc                      bool        `json:"Delete_mc"`
			UpdateMc                      bool        `json:"Update_mc"`
			ToClassCharacteristicOc       bool        `json:"to_ClassCharacteristic_oc"`
			ToClassDescriptionOc          bool        `json:"to_ClassDescription_oc"`
			ToClassKeywordOc              bool        `json:"to_ClassKeyword_oc"`
			ToClassTextOc                 bool        `json:"to_ClassText_oc"`
			ClassInternalID               string      `json:"ClassInternalID"`
			ClassType                     string      `json:"ClassType"`
			ClassTypeName                 string      `json:"ClassTypeName"`
			Class                         string      `json:"Class"`
			ClassStatus                   string      `json:"ClassStatus"`
			ClassStatusName               string      `json:"ClassStatusName"`
			ClassGroup                    string      `json:"ClassGroup"`
			ClassGroupName                string      `json:"ClassGroupName"`
			ClassSearchAuthGrp            string      `json:"ClassSearchAuthGrp"`
			ClassClassfctnAuthGrp         string      `json:"ClassClassfctnAuthGrp"`
			ClassMaintAuthGrp             string      `json:"ClassMaintAuthGrp"`
			CreationDate                  string      `json:"CreationDate"`
			LastChangeDate                string      `json:"LastChangeDate"`
			ValidityStartDate             string      `json:"ValidityStartDate"`
			ValidityEndDate               string      `json:"ValidityEndDate"`
			ClassLastChangedDateTime      string      `json:"ClassLastChangedDateTime"`
			KeyDate                       string      `json:"KeyDate"`
			ToCharc                       struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ClassCharacteristic"`
			ToClassDescription struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ClassDescription"`
		} `json:"results"`
	} `json:"d"`
}


