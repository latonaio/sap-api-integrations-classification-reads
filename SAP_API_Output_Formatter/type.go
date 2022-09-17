package sap_api_output_formatter

type Classification struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	ClassCode     string `json:"class_code"`
	Deleted       bool   `json:"deleted"`
}

type Class struct {
	DeleteMc                 bool   `json:"Delete_mc"`
	UpdateMc                 bool   `json:"Update_mc"`
	ToClassCharacteristicOc  bool   `json:"to_ClassCharacteristic_oc"`
	ToClassDescriptionOc     bool   `json:"to_ClassDescription_oc"`
	ToClassKeywordOc         bool   `json:"to_ClassKeyword_oc"`
	ToClassTextOc            bool   `json:"to_ClassText_oc"`
	ClassInternalID          string `json:"ClassInternalID"`
	ClassType                string `json:"ClassType"`
	ClassTypeName            string `json:"ClassTypeName"`
	Class                    string `json:"Class"`
	ClassStatus              string `json:"ClassStatus"`
	ClassStatusName          string `json:"ClassStatusName"`
	ClassGroup               string `json:"ClassGroup"`
	ClassGroupName           string `json:"ClassGroupName"`
	ClassSearchAuthGrp       string `json:"ClassSearchAuthGrp"`
	ClassClassfctnAuthGrp    string `json:"ClassClassfctnAuthGrp"`
	ClassMaintAuthGrp        string `json:"ClassMaintAuthGrp"`
	CreationDate             string `json:"CreationDate"`
	LastChangeDate           string `json:"LastChangeDate"`
	ValidityStartDate        string `json:"ValidityStartDate"`
	ValidityEndDate          string `json:"ValidityEndDate"`
	ClassLastChangedDateTime string `json:"ClassLastChangedDateTime"`
	KeyDate                  string `json:"KeyDate"`
	ToCharc                  string `json:"to_ClassCharacteristic"`
	ToClassDescription       string `json:"to_ClassDescription"`
}

type ToCharc struct {
	DeleteMc                 bool   `json:"Delete_mc"`
	UpdateMc                 bool   `json:"Update_mc"`
	ClassInternalID          string `json:"ClassInternalID"`
	CharcInternalID          string `json:"CharcInternalID"`
	CharcPositionNumber      string `json:"CharcPositionNumber"`
	Characteristic           string `json:"Characteristic"`
	AncestorClassInternalID  string `json:"AncestorClassInternalID"`
	OriginalCharcInternalID  string `json:"OriginalCharcInternalID"`
	ChangeNumber             string `json:"ChangeNumber"`
	CharcIsPrintRelevant     string `json:"CharcIsPrintRelevant"`
	CharcIsSearchRelevant    string `json:"CharcIsSearchRelevant"`
	CharcIsDisplayRelevant   string `json:"CharcIsDisplayRelevant"`
	ValidityStartDate        string `json:"ValidityStartDate"`
	ValidityEndDate          string `json:"ValidityEndDate"`
	KeyDate                  string `json:"KeyDate"`
	ClassLastChangedDateTime string `json:"ClassLastChangedDateTime"`
}

type ToClassDescription struct {
	UpdateMc                 bool   `json:"Update_mc"`
	ClassInternalID          string `json:"ClassInternalID"`
	Language                 string `json:"Language"`
	ClassDescription         string `json:"ClassDescription"`
	KeyDate                  string `json:"KeyDate"`
	ClassLastChangedDateTime string `json:"ClassLastChangedDateTime"`
}
