package document

import (
	"passportreader/helper"
	"passportreader/lib"
)

var IndonesiaIDCardDocumentType = 2010

var (
	IndonesiaIDCardFieldReserve              = 0
	IndonesiaIDCardFieldIDCardNumber         = 1
	IndonesiaIDCardFieldName                 = 2
	IndonesiaIDCardFieldBirthDate            = 3
	IndonesiaIDCardFieldBirthPlace           = 4
	IndonesiaIDCardFieldGender               = 5
	IndonesiaIDCardFieldArea                 = 6
	IndonesiaIDCardFieldCountryOfCitizenship = 7
	IndonesiaIDCardFieldNationality          = 8
	IndonesiaIDCardFieldBloodType            = 9
	IndonesiaIDCardFieldAddress              = 10
)

type IndonesiaIDCard struct {
	Reserve              string
	IDCardNumber         string
	Name                 string
	BirthDate            string
	BirthPlace           string
	Gender               string
	Area                 string
	CountryOfCitizenship string
	Nationality          string
	BloodType            string
	Address              string
}

func NewIndonesiaIDCard() Document {
	return &IndonesiaIDCard{}
}

func (indonesiaIDCard *IndonesiaIDCard) Fill(idCard lib.IDCard) error {
	for attr := 0; attr < 2; attr++ {
		for idx := 0; idx < 11; idx++ {
			buff := make([]uint16, 255)
			buffLen := len(buff)
			ret := idCard.GetRecogResultEx(attr, idx, &buff, &buffLen)
			fieldData := helper.UTF16ToString(&buff[0], buffLen)

			if ret == 0 {
				switch idx {
				case IndonesiaIDCardFieldReserve:
					indonesiaIDCard.Reserve = fieldData
				case IndonesiaIDCardFieldIDCardNumber:
					indonesiaIDCard.IDCardNumber = fieldData
				case IndonesiaIDCardFieldName:
					indonesiaIDCard.Name = fieldData
				case IndonesiaIDCardFieldBirthDate:
					indonesiaIDCard.BirthDate = fieldData
				case IndonesiaIDCardFieldBirthPlace:
					indonesiaIDCard.BirthPlace = fieldData
				case IndonesiaIDCardFieldGender:
					indonesiaIDCard.Gender = fieldData
				case IndonesiaIDCardFieldArea:
					indonesiaIDCard.Area = fieldData
				case IndonesiaIDCardFieldCountryOfCitizenship:
					indonesiaIDCard.CountryOfCitizenship = fieldData
				case IndonesiaIDCardFieldNationality:
					indonesiaIDCard.Nationality = fieldData
				case IndonesiaIDCardFieldBloodType:
					indonesiaIDCard.BloodType = fieldData
				case IndonesiaIDCardFieldAddress:
					indonesiaIDCard.Address = fieldData
				}
			}
		}
	}

	return nil
}

func (indonesiaIDCard *IndonesiaIDCard) GetBody() interface{} {
	return indonesiaIDCard
}
