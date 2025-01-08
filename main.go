package main

import (
	"log"
	"os"
	"os/signal"
	"passportreader/document"
	"passportreader/helper"
	"passportreader/lib"
	"syscall"
)

func main() {
	// recover from panic
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	// listen for signals
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		done <- true
	}()

	// load configuration from environment variables
	config := LoadFromEnv()

	// initialize IDCardDLL
	idCard, err := lib.NewIDCardDLL(config.DllPath)
	if err != nil {
		panic(err)
	}
	defer idCard.Free()
	log.Println("IDCardDLL initialized")

	// initialize IDCard
	initIDCardRet, err := idCard.InitIDCard(config.UserID, 0, config.DirPath)
	if err != nil {
		panic(err)
	}
	defer idCard.FreeIDCard()
	if initIDCardRet != 0 {
		panic("InitIDCard failed")
	}

	// add support for Indonesia ID Card
	idCard.AddIDCardID(document.IndonesiaIDCardDocumentType, []int{0}, 1)

	for {
		select {
		case <-sigs:
			return
		default:
			// detect document
			detectDocRet := idCard.DetectDocument()
			if detectDocRet < 1 {
				continue
			}

			switch detectDocRet {
			case 1:
				log.Println("Document is placed in")

				var nType int

				// classify document
				cardType := idCard.ClassifyIDCard(&nType)
				if cardType <= 0 {
					log.Println("Document classification failed")
					continue
				}

				// auto process document
				autoProcessRet := idCard.AutoProcessIDCard(&nType)
				if autoProcessRet <= 0 {
					log.Println("Auto process document is failed")
					continue
				}

				switch cardType {
				case document.IndonesiaIDCardDocumentType:
					log.Println("Document Type: Indonesia ID Card")

					// create Indonesia ID Card object
					indonesiaIDCard := document.NewIndonesiaIDCard()

					// fill Indonesia ID Card object with data
					if err := indonesiaIDCard.Fill(idCard); err != nil {
						log.Println(err)
						continue
					}

					body := indonesiaIDCard.GetBody().(*document.IndonesiaIDCard)
					log.Println("=== Information ===")
					log.Println(">>> Reserve:", body.Reserve)
					log.Println(">>> ID Card Number:", body.IDCardNumber)
					log.Println(">>> Name:", body.Name)
					log.Println(">>> Birth Date:", body.BirthDate)
					log.Println(">>> Birth Place:", body.BirthPlace)
					log.Println(">>> Gender", body.Gender)
					log.Println(">>> Area:", body.Area)
					log.Println(">>> Country Of Citizenship:", body.CountryOfCitizenship)
					log.Println(">>> Nationality", body.Nationality)
					log.Println(">>> Blood Type:", body.BloodType)
					log.Println(">>> Address:", body.Address)

				case document.PassportDocumentType:
					log.Println("Document Type: Passport")

					// create Passport object
					passport := document.NewPassport()

					// fill Passport object with data
					if err := passport.Fill(idCard); err != nil {
						log.Println(err)
						continue
					}

					body := passport.GetBody().(*document.Passport)
					log.Println("=== Information ===")
					log.Println(">>> Passport Type:", body.PassportType)
					log.Println(">>> Passport Number MRZ:", body.PassportNumberMRZ)
					log.Println(">>> Domestic Name:", body.DomesticName)
					log.Println(">>> English Name:", body.EnglishName)
					log.Println(">>> Gender", body.Gender)
					log.Println(">>> Birth Day:", body.BirthDay)
					log.Println(">>> Birth Place:", body.BirthPlace)
					log.Println(">>> Expiry Date:", body.ExpiryDate)
					log.Println(">>> Issuing Country Code:", body.IssuingCountryCode)
					log.Println(">>> English Surname:", body.EnglishSurname)
					log.Println(">>> English First Name:", body.EnglishFirstName)
					log.Println(">>> MRZ1:", body.MRZ1)
					log.Println(">>> MRZ2:", body.MRZ2)
					log.Println(">>> Holder Nationality Code:", body.HolderNationalityCode)
					log.Println(">>> Passport Number:", body.PassportNumber)
					log.Println(">>> Birth Place:", body.BirthPlace)
					log.Println(">>> Issue Place:", body.IssuePlace)
					log.Println(">>> Issue Date:", body.IssueDate)
					log.Println(">>> RFID MRZ:", body.RFIDMRZ)
					log.Println(">>> OCR MRZ:", body.OCRMRZ)
					log.Println(">>> Birth Place Pinyin:", body.BirthPlacePinyin)
					log.Println(">>> Issue Place Pinyin:", body.IssuePlacePinyin)
					log.Println(">>> ID Number:", body.IDNumber)
					log.Println(">>> National Name Pinyin OCR:", body.NationalNamePinyinOCR)
					log.Println(">>> Gender OCR", body.GenderOCR)
					log.Println(">>> License Nationality Code OCR:", body.LicenseNationalityCodeOCR)
					log.Println(">>> ID Card Number OCR:", body.IDCardNumberOCR)
					log.Println(">>> Birth Date OCR:", body.BirthDateOCR)
					log.Println(">>> Valid Until OCR:", body.ValidUntilOCR)
					log.Println(">>> Issuing Authority OCR:", body.IssuingAuthorityOCR)
					log.Println(">>> Domestic Surname:", body.DomesticSurname)
					log.Println(">>> Domestic First Name:", body.DomesticFirstName)

				default:
					log.Println("Unknown document type")
					continue
				}

				// generate random file name
				fn, err := helper.GenerateRandomString(8)
				if err != nil {
					log.Println(err)
					continue
				}
				fn = "C:\\Users\\asus_\\Test\\gov2\\images\\" + fn + ".jpg"

				// perform save image
				saveImgRet, err := idCard.SaveImageEx(fn, 1)
				if err != nil {
					log.Println(err)
					continue
				}
				if saveImgRet != 0 {
					log.Println("Save image failed")
					continue
				}

				log.Println("Image saved to", fn)

			case 2:
				log.Println("Document is taken out")
			}
		}
	}
}
