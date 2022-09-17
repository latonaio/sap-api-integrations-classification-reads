package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-classification-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToClass(raw []byte, l *logger.Logger) ([]Class, error) {
	pm := &responses.Class{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Class. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	class := make([]Class, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		class = append(class, Class{
			DeleteMc:                      data.DeleteMc,
			UpdateMc:                      data.UpdateMc,
			ToClassCharacteristicOc:       data.ToClassCharacteristicOc,
			ToClassDescriptionOc:          data.ToClassDescriptionOc,
			ToClassKeywordOc:              data.ToClassKeywordOc,
			ToClassTextOc:                 data.ToClassTextOc,
			ClassInternalID:               data.ClassInternalID,
			ClassType:                     data.ClassType,
			ClassTypeName:                 data.ClassTypeName,
			Class:                         data.Class,
			ClassStatus:                   data.ClassStatus,
			ClassStatusName:               data.ClassStatusName,
			ClassGroup:                    data.ClassGroup,
			ClassGroupName:                data.ClassGroupName,
			ClassSearchAuthGrp:            data.ClassSearchAuthGrp,
			ClassClassfctnAuthGrp:         data.ClassClassfctnAuthGrp,
			ClassMaintAuthGrp:             data.ClassMaintAuthGrp,
			CreationDate:                  data.CreationDate,
			LastChangeDate:                data.LastChangeDate,
			ValidityStartDate:             data.ValidityStartDate,
			ValidityEndDate:               data.ValidityEndDate,
			ClassLastChangedDateTime:      data.ClassLastChangedDateTime,
			KeyDate:                       data.KeyDate,
			ToCharc:                       data.ToCharc.Deferred.URI,
			ToClassDescription:            data.ToClassDescription.Deferred.URI,
		})
	}

	return class, nil
}

func ConvertToToCharc(raw []byte, l *logger.Logger) ([]ToCharc, error) {
	pm := &responses.ToCharc{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToCharc. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	toCharc := make([]ToCharc, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toCharc = append(toCharc, ToCharc{
			DeleteMc:                 data.DeleteMc,
	UpdateMc:                 data.UpdateMc,
	ClassInternalID:          data.ClassInternalID,
	CharcInternalID:          data.CharcInternalID,
	CharcPositionNumber:      data.CharcPositionNumber,
	Characteristic:           data.Characteristic,
	AncestorClassInternalID:  data.AncestorClassInternalID,
	OriginalCharcInternalID:  data.OriginalCharcInternalID,
	ChangeNumber:             data.ChangeNumber,
	CharcIsPrintRelevant:     data.CharcIsPrintRelevant,
	CharcIsSearchRelevant:    data.CharcIsSearchRelevant,
	CharcIsDisplayRelevant:   data.CharcIsDisplayRelevant,
	ValidityStartDate:        data.ValidityStartDate,
	ValidityEndDate:          data.ValidityEndDate,
	KeyDate:                  data.KeyDate,
	ClassLastChangedDateTime: data.ClassLastChangedDateTime,
		})
	}

	return toCharc, nil
}

func ConvertToToClassDescription(raw []byte, l *logger.Logger) ([]ToClassDescription, error) {
	pm := &responses.ToClassDescription{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToClassDescription. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toClassDescription := make([]ToClassDescription, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toClassDescription = append(toClassDescription, ToClassDescription{
	UpdateMc:                 data.UpdateMc,
	ClassInternalID:          data.ClassInternalID,
	Language:                 data.Language,
	ClassDescription:         data.ClassDescription,
	KeyDate:                  data.KeyDate,
	ClassLastChangedDateTime: data.ClassLastChangedDateTime,
		})
	}

	return toClassDescription, nil
}

