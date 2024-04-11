// models/epicResponse.go
package models

import "time"

type EpicSearchResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// Add this struct definition to your existing models/responses.go file
type EpicAPIResponse struct {
	Index           string                 `json:"index"`
	ID              string                 `json:"id"`
	Score           float64                `json:"score"`
	SortValues      []interface{}          `json:"sortValues"`
	Content         EpicAPIContent         `json:"content"`
	HighlightFields map[string]interface{} `json:"highlightFields"`
	InnerHits       map[string]interface{} `json:"innerHits"`
	NestedMetaData  interface{}            `json:"nestedMetaData"`
	Routing         interface{}            `json:"routing"`
	Explanation     interface{}            `json:"explanation"`
	MatchedQueries  []interface{}          `json:"matchedQueries"`
}

type EpicAPIContent struct {
	ID                      TrimmedString  `json:"id,omitempty"`
	EpicID                  int64          `json:"epicId,omitempty"`
	EpicNumber              TrimmedString  `json:"epicNumber,omitempty"`
	FormReferenceNo         *TrimmedString `json:"formReferenceNo,omitempty"`
	ApplicantFirstName      TrimmedString  `json:"applicantFirstName,omitempty"`
	ApplicantFirstNameL1    *TrimmedString `json:"applicantFirstNameL1,omitempty"`
	ApplicantFirstNameL2    *TrimmedString `json:"applicantFirstNameL2,omitempty"`
	ApplicantLastName       *TrimmedString `json:"applicantLastName,omitempty"`
	ApplicantLastNameL1     *TrimmedString `json:"applicantLastNameL1,omitempty"`
	ApplicantLastNameL2     *TrimmedString `json:"applicantLastNameL2,omitempty"`
	RelationName            *TrimmedString `json:"relationName,omitempty"`
	RelationNameL1          *TrimmedString `json:"relationNameL1,omitempty"`
	RelationNameL2          *TrimmedString `json:"relationNameL2,omitempty"`
	Age                     int            `json:"age,omitempty"`
	Gender                  TrimmedString  `json:"gender,omitempty"`
	PartNumber              int            `json:"partNumber,omitempty"`
	PartId                  int            `json:"partId,omitempty"`
	PartName                TrimmedString  `json:"partName,omitempty"`
	PartNameL1              TrimmedString  `json:"partNameL1,omitempty"`
	PartSerialNumber        int            `json:"partSerialNumber,omitempty"`
	AsmblyName              TrimmedString  `json:"asmblyName,omitempty"`
	AsmblyNameL1            TrimmedString  `json:"asmblyNameL1,omitempty"`
	AcId                    int            `json:"acId,omitempty"`
	AcNumber                int            `json:"acNumber,omitempty"`
	PrlmntName              TrimmedString  `json:"prlmntName,omitempty"`
	PrlmntNameL1            *TrimmedString `json:"prlmntNameL1,omitempty"`
	PrlmntNo                TrimmedString  `json:"prlmntNo,omitempty"`
	DistrictValue           TrimmedString  `json:"districtValue,omitempty"`
	DistrictValueL1         TrimmedString  `json:"districtValueL1,omitempty"`
	DistrictCd              TrimmedString  `json:"districtCd,omitempty"`
	DistrictId              *int           `json:"districtId,omitempty"`
	DistrictNo              int            `json:"districtNo,omitempty"`
	StateName               TrimmedString  `json:"stateName,omitempty"`
	StateNameL1             TrimmedString  `json:"stateNameL1,omitempty"`
	StateId                 int            `json:"stateId,omitempty"`
	StateCd                 TrimmedString  `json:"stateCd,omitempty"`
	IsActive                bool           `json:"isActive,omitempty"`
	FormType                *TrimmedString `json:"formType,omitempty"`
	RelationType            TrimmedString  `json:"relationType,omitempty"`
	CreatedDttm             *TrimmedString `json:"createdDttm,omitempty"`
	ModifiedDttm            *TrimmedString `json:"modifiedDttm,omitempty"`
	EpicDatetime            *TrimmedString `json:"epicDatetime,omitempty"`
	GenderL1                *TrimmedString `json:"genderL1,omitempty"`
	IsDeleted               *bool          `json:"isDeleted,omitempty"`
	ProcessType             *TrimmedString `json:"processType,omitempty"`
	RelationLName           TrimmedString  `json:"relationLName,omitempty"`
	RelationLNameL1         TrimmedString  `json:"relationLNameL1,omitempty"`
	RelationTypeL1          *TrimmedString `json:"relationTypeL1,omitempty"`
	RevisionId              *int           `json:"revisionId,omitempty"`
	StatusType              *TrimmedString `json:"statusType,omitempty"`
	SectionNo               int            `json:"sectionNo,omitempty"`
	PollingDate             *time.Time     `json:"pollingDate,omitempty"`
	ElectionDate            TrimmedString  `json:"electionDate,omitempty"`
	ElectionTime            *TrimmedString `json:"electionTime,omitempty"`
	PsbuildingName          TrimmedString  `json:"psbuildingName,omitempty"`
	BirthYear               *int           `json:"birthYear,omitempty"`
	PartLatLong             *TrimmedString `json:"partLatLong,omitempty"`
	DisabilityAny           *bool          `json:"disabilityAny,omitempty"`
	DisabilityType          *TrimmedString `json:"disabilityType,omitempty"`
	IsForm8Migration        *bool          `json:"isForm8Migration,omitempty"`
	IsLocomotorDisabled     *bool          `json:"isLocomotorDisabled,omitempty"`
	IsSpeechHearingDisabled *bool          `json:"isSpeechHearingDisabled,omitempty"`
	IsVisuallyImpaired      *bool          `json:"isVisuallyImpaired,omitempty"`
	OtherDisability         *TrimmedString `json:"otherDisability,omitempty"`
	IsValidated             *bool          `json:"isValidated,omitempty"`
	IsWheelchairRequired    *bool          `json:"isWheelchairRequired,omitempty"`
	PwdMarkingFormType      *TrimmedString `json:"pwdMarkingFormType,omitempty"`
	PwdMarkingRefNo         *TrimmedString `json:"pwdMarkingRefNo,omitempty"`
	Pwd                     *bool          `json:"pwd,omitempty"`
	IsVip                   *bool          `json:"isVip,omitempty"`
	FullName                TrimmedString  `json:"fullName,omitempty"`
	FullNameL1              TrimmedString  `json:"fullNameL1,omitempty"`
	RelativeFullName        TrimmedString  `json:"relativeFullName,omitempty"`
	RelativeFullNameL1      TrimmedString  `json:"relativeFullNameL1,omitempty"`
	PsRoomDetails           *TrimmedString `json:"psRoomDetails,omitempty"`
	PsRoomDetailsL1         *TrimmedString `json:"psRoomDetailsL1,omitempty"`
	PsBuildingNameL1        TrimmedString  `json:"psBuildingNameL1,omitempty"`
	BuildingAddress         TrimmedString  `json:"buildingAddress,omitempty"`
	BuildingAddressL1       TrimmedString  `json:"buildingAddressL1,omitempty"`
}
