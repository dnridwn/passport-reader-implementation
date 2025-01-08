package document

import (
	"errors"
	"passportreader/helper"
	"passportreader/lib"
)

var PassportDocumentType = 13

var (
	PassportFieldPassportType              = 0
	PassportFieldPassportNumberMRZ         = 1 // MRZ output
	PassportFieldDomesticName              = 2
	PassportFieldEnglishName               = 3
	PassportFieldGender                    = 4
	PassportFieldBirthDay                  = 5
	PassportFieldExpiryDate                = 6
	PassportFieldIssuingCountryCode        = 7
	PassportFieldEnglishSurname            = 8
	PassportFieldEnglishFirstName          = 9
	PassportFieldMRZ1                      = 10
	PassportFieldMRZ2                      = 11
	PassportFieldHolderNationalityCode     = 12
	PassportFieldPassportNumber            = 13 // Direct identification
	PassportFieldBirthPlace                = 14 // Chinese passport only
	PassportFieldIssuePlace                = 15 // Chinese passport only
	PassportFieldIssueDate                 = 16 // Chinese passport only
	PassportFieldRFIDMRZ                   = 17
	PassportFieldOCRMRZ                    = 18
	PassportFieldBirthPlacePinyin          = 19 // Chinese passport only
	PassportFieldIssuePlacePinyin          = 20 // Chinese passport only
	PassportFieldIDNumber                  = 21 // Taiwan and Korean passports only
	PassportFieldNationalNamePinyinOCR     = 22
	PassportFieldGenderOCR                 = 23
	PassportFieldLicenseNationalityCodeOCR = 24
	PassportFieldIDCardNumberOCR           = 25
	PassportFieldBirthDateOCR              = 26
	PassportFieldValidUntilOCR             = 27
	PassportFieldIssuingAuthorityOCR       = 28
	PassportFieldDomesticSurname           = 29
	PassportFieldDomesticFirstName         = 30
)

type Passport struct {
	PassportType              string
	PassportNumberMRZ         string
	DomesticName              string
	EnglishName               string
	Gender                    string
	BirthDay                  string
	ExpiryDate                string
	IssuingCountryCode        string
	EnglishSurname            string
	EnglishFirstName          string
	MRZ1                      string
	MRZ2                      string
	HolderNationalityCode     string
	PassportNumber            string
	BirthPlace                string
	IssuePlace                string
	IssueDate                 string
	RFIDMRZ                   string
	OCRMRZ                    string
	BirthPlacePinyin          string
	IssuePlacePinyin          string
	IDNumber                  string
	NationalNamePinyinOCR     string
	GenderOCR                 string
	LicenseNationalityCodeOCR string
	IDCardNumberOCR           string
	BirthDateOCR              string
	ValidUntilOCR             string
	IssuingAuthorityOCR       string
	DomesticSurname           string
	DomesticFirstName         string
}

func NewPassport() Document {
	return &Passport{}
}

func (passport *Passport) Fill(idCard lib.IDCard) error {
	nType := 1
	ret := idCard.AutoProcessIDCard(&nType)
	if ret <= 0 {
		return errors.New("get passport information failed")
	}

	for attr := 0; attr < 2; attr++ {
		for idx := 0; idx < 31; idx++ {
			buff := make([]uint16, 255)
			buffLen := len(buff)
			ret := idCard.GetRecogResultEx(attr, idx, &buff, &buffLen)
			fieldData := helper.UTF16ToString(&buff[0], buffLen)

			if ret == 0 {
				switch idx {
				case PassportFieldPassportType:
					passport.PassportType = fieldData
				case PassportFieldPassportNumberMRZ:
					passport.PassportNumberMRZ = fieldData
				case PassportFieldDomesticName:
					passport.DomesticName = fieldData
				case PassportFieldEnglishName:
					passport.EnglishName = fieldData
				case PassportFieldGender:
					passport.Gender = fieldData
				case PassportFieldBirthDay:
					passport.BirthDay = fieldData
				case PassportFieldExpiryDate:
					passport.ExpiryDate = fieldData
				case PassportFieldIssuingCountryCode:
					passport.IssuingCountryCode = fieldData
				case PassportFieldEnglishSurname:
					passport.EnglishSurname = fieldData
				case PassportFieldEnglishFirstName:
					passport.EnglishFirstName = fieldData
				case PassportFieldMRZ1:
					passport.MRZ1 = fieldData
				case PassportFieldMRZ2:
					passport.MRZ2 = fieldData
				case PassportFieldHolderNationalityCode:
					passport.HolderNationalityCode = fieldData
				case PassportFieldPassportNumber:
					passport.PassportNumber = fieldData
				case PassportFieldBirthPlace:
					passport.BirthPlace = fieldData
				case PassportFieldIssuePlace:
					passport.IssuePlace = fieldData
				case PassportFieldIssueDate:
					passport.IssueDate = fieldData
				case PassportFieldRFIDMRZ:
					passport.RFIDMRZ = fieldData
				case PassportFieldOCRMRZ:
					passport.OCRMRZ = fieldData
				case PassportFieldBirthPlacePinyin:
					passport.BirthPlacePinyin = fieldData
				case PassportFieldIssuePlacePinyin:
					passport.IssuePlacePinyin = fieldData
				case PassportFieldIDNumber:
					passport.IDNumber = fieldData
				case PassportFieldNationalNamePinyinOCR:
					passport.NationalNamePinyinOCR = fieldData
				case PassportFieldGenderOCR:
					passport.GenderOCR = fieldData
				case PassportFieldLicenseNationalityCodeOCR:
					passport.LicenseNationalityCodeOCR = fieldData
				case PassportFieldIDCardNumberOCR:
					passport.IDCardNumberOCR = fieldData
				case PassportFieldBirthDateOCR:
					passport.BirthDateOCR = fieldData
				case PassportFieldValidUntilOCR:
					passport.ValidUntilOCR = fieldData
				case PassportFieldIssuingAuthorityOCR:
					passport.IssuingAuthorityOCR = fieldData
				case PassportFieldDomesticSurname:
					passport.DomesticSurname = fieldData
				case PassportFieldDomesticFirstName:
					passport.DomesticFirstName = fieldData
				}
			}
		}
	}

	return nil
}

func (passport *Passport) GetBody() interface{} {
	return passport
}
