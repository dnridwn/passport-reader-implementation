package lib

import (
	"errors"
	"path"
	"syscall"
	"unsafe"
)

type IDCardDLL struct {
	dll                          *syscall.DLL
	initIDCard                   *syscall.Proc
	freeIDCard                   *syscall.Proc
	detectDocument               *syscall.Proc
	autoProcessIDCard            *syscall.Proc
	getFieldNameEx               *syscall.Proc
	getRecogResultEx             *syscall.Proc
	getResultTypeEx              *syscall.Proc
	getFieldConfEx               *syscall.Proc
	getIDCardName                *syscall.Proc
	getSubID                     *syscall.Proc
	saveImageEx                  *syscall.Proc
	setBarCodeMode               *syscall.Proc
	setRecogCodeTypes            *syscall.Proc
	recogCellPhoneBarCode        *syscall.Proc
	getBarcodeCount              *syscall.Proc
	getBarcodeResult             *syscall.Proc
	setConfigByFile              *syscall.Proc
	setLanguage                  *syscall.Proc
	setSaveImageType             *syscall.Proc
	setRecogVIZ                  *syscall.Proc
	setRecogDG                   *syscall.Proc
	setAnalyseMRZ                *syscall.Proc
	resetIDCardID                *syscall.Proc
	addIDCardID                  *syscall.Proc
	getDeviceSN                  *syscall.Proc
	getCurrentDevice             *syscall.Proc
	getVersionInfo               *syscall.Proc
	checkDeviceOnlineEx          *syscall.Proc
	setIOStatus                  *syscall.Proc
	buzzerAlarm                  *syscall.Proc
	acquireImage                 *syscall.Proc
	initIDCardEx                 *syscall.Proc
	initIDCardSN                 *syscall.Proc
	setIDCardRejectType          *syscall.Proc
	setRecogChipCardAttribute    *syscall.Proc
	reRecogMRZbyVI               *syscall.Proc
	setBGSubtraction             *syscall.Proc
	setAcquireImageResolution    *syscall.Proc
	settAcquireImageExposureTime *syscall.Proc
	checkUVDull                  *syscall.Proc
	fibreDetect                  *syscall.Proc
	getFibrePos                  *syscall.Proc
	getImageSourceType           *syscall.Proc
	getAnalyseMRZFieldName       *syscall.Proc
	getAnalyseMRZResult          *syscall.Proc
	recogIDCardEX                *syscall.Proc
	getDataGroupContent          *syscall.Proc
	classifyIDCard               *syscall.Proc
}

func NewIDCardDLL(dllPath string) (IDCard, error) {
	if path.Ext(dllPath) != ".dll" {
		return nil, errors.New("invalid DLL file extension")
	}

	dll, err := syscall.LoadDLL(dllPath)
	if err != nil {
		return nil, err
	}

	initIDCard, err := dll.FindProc("InitIDCard")
	if err != nil {
		return nil, err
	}

	freeIDCard, err := dll.FindProc("FreeIDCard")
	if err != nil {
		return nil, err
	}

	detectDocument, err := dll.FindProc("DetectDocument")
	if err != nil {
		return nil, err
	}

	autoProcessIDCard, err := dll.FindProc("AutoProcessIDCard")
	if err != nil {
		return nil, err
	}

	getFieldNameEx, err := dll.FindProc("GetFieldNameEx")
	if err != nil {
		return nil, err
	}

	getRecogResultEx, err := dll.FindProc("GetRecogResultEx")
	if err != nil {
		return nil, err
	}

	getResultTypeEx, err := dll.FindProc("GetResultTypeEx")
	if err != nil {
		return nil, err
	}

	getFieldConfEx, err := dll.FindProc("GetFieldConfEx")
	if err != nil {
		return nil, err
	}

	getIDCardName, err := dll.FindProc("GetIDCardName")
	if err != nil {
		return nil, err
	}

	getSubID, err := dll.FindProc("GetSubID")
	if err != nil {
		return nil, err
	}

	saveImageEx, err := dll.FindProc("SaveImageEx")
	if err != nil {
		return nil, err
	}

	setBarCodeMode, err := dll.FindProc("SetBarCodeMode")
	if err != nil {
		return nil, err
	}

	setRecogCodeTypes, err := dll.FindProc("SetRecogCodeTypes")
	if err != nil {
		return nil, err
	}

	recogCellPhoneBarCode, err := dll.FindProc("RecogCellPhoneBarCode")
	if err != nil {
		return nil, err
	}

	getBarcodeCount, err := dll.FindProc("GetBarcodeCount")
	if err != nil {
		return nil, err
	}

	getBarcodeRecogResult, err := dll.FindProc("GetBarcodeRecogResult")
	if err != nil {
		return nil, err
	}

	setConfigByFile, err := dll.FindProc("SetConfigByFile")
	if err != nil {
		return nil, err
	}

	setLanguage, err := dll.FindProc("SetLanguage")
	if err != nil {
		return nil, err
	}

	setRecogVIZ, err := dll.FindProc("SetRecogVIZ")
	if err != nil {
		return nil, err
	}

	setRecogDG, err := dll.FindProc("SetRecogDG")
	if err != nil {
		return nil, err
	}

	setAnalyseMRZ, err := dll.FindProc("SetAnalyseMRZ")
	if err != nil {
		return nil, err
	}

	resetIDCardID, err := dll.FindProc("ResetIDCardID")
	if err != nil {
		return nil, err
	}

	addIDCardID, err := dll.FindProc("AddIDCardID")
	if err != nil {
		return nil, err
	}

	getDeviceSN, err := dll.FindProc("GetDeviceSN")
	if err != nil {
		return nil, err
	}

	getCurrentDevice, err := dll.FindProc("GetCurrentDevice")
	if err != nil {
		return nil, err
	}

	getVersionInfo, err := dll.FindProc("GetVersionInfo")
	if err != nil {
		return nil, err
	}

	checkDeviceOnlineEx, err := dll.FindProc("CheckDeviceOnlineEx")
	if err != nil {
		return nil, err
	}

	setIOStatus, err := dll.FindProc("SetIOStatus")
	if err != nil {
		return nil, err
	}

	buzzerAlarm, err := dll.FindProc("BuzzerAlarm")
	if err != nil {
		return nil, err
	}

	acquireImage, err := dll.FindProc("AcquireImage")
	if err != nil {
		return nil, err
	}

	initIDCardEx, err := dll.FindProc("InitIDCardEx")
	if err != nil {
		return nil, err
	}

	initIDCardSN, err := dll.FindProc("InitIDCardSN")
	if err != nil {
		return nil, err
	}

	setIDCardRejectType, err := dll.FindProc("SetIDCardRejectType")
	if err != nil {
		return nil, err
	}

	setRecogChipCardAttribute, err := dll.FindProc("SetRecogChipCardAttribute")
	if err != nil {
		return nil, err
	}

	reRecogMRZbyVI, err := dll.FindProc("ReRecogMRZbyVI")
	if err != nil {
		return nil, err
	}

	setBGSubtraction, err := dll.FindProc("SetBGSubtraction")
	if err != nil {
		return nil, err
	}

	setAcquireImageResolution, err := dll.FindProc("SetAcquireImageResolution")
	if err != nil {
		return nil, err
	}

	setAcquireImageExposureTime, err := dll.FindProc("SetAcquireImageExposureTime")
	if err != nil {
		return nil, err
	}

	checkUVDull, err := dll.FindProc("CheckUVDull")
	if err != nil {
		return nil, err
	}

	fibreDetect, err := dll.FindProc("FibreDetect")
	if err != nil {
		return nil, err
	}

	getFibrePos, err := dll.FindProc("GetFibrePos")
	if err != nil {
		return nil, err
	}

	getImageSourceType, err := dll.FindProc("GetImageSourceType")
	if err != nil {
		return nil, err
	}

	getAnalyseMRZFieldName, err := dll.FindProc("GetAnalyseMRZFieldName")
	if err != nil {
		return nil, err
	}

	getAnalyseMRZResult, err := dll.FindProc("GetAnalyseMRZResult")
	if err != nil {
		return nil, err
	}

	recogIDCardEX, err := dll.FindProc("RecogIDCardEX")
	if err != nil {
		return nil, err
	}

	getDataGroupContent, err := dll.FindProc("GetDataGroupContent")
	if err != nil {
		return nil, err
	}

	classifyIDCard, err := dll.FindProc("ClassifyIDCard")
	if err != nil {
		return nil, err
	}

	return &IDCardDLL{
		dll:                          dll,
		initIDCard:                   initIDCard,
		freeIDCard:                   freeIDCard,
		detectDocument:               detectDocument,
		autoProcessIDCard:            autoProcessIDCard,
		getFieldNameEx:               getFieldNameEx,
		getRecogResultEx:             getRecogResultEx,
		getResultTypeEx:              getResultTypeEx,
		getFieldConfEx:               getFieldConfEx,
		getIDCardName:                getIDCardName,
		getSubID:                     getSubID,
		saveImageEx:                  saveImageEx,
		setBarCodeMode:               setBarCodeMode,
		setRecogCodeTypes:            setRecogCodeTypes,
		recogCellPhoneBarCode:        recogCellPhoneBarCode,
		getBarcodeCount:              getBarcodeCount,
		getBarcodeResult:             getBarcodeRecogResult,
		setConfigByFile:              setConfigByFile,
		setLanguage:                  setLanguage,
		setRecogVIZ:                  setRecogVIZ,
		setRecogDG:                   setRecogDG,
		setAnalyseMRZ:                setAnalyseMRZ,
		resetIDCardID:                resetIDCardID,
		addIDCardID:                  addIDCardID,
		getDeviceSN:                  getDeviceSN,
		getCurrentDevice:             getCurrentDevice,
		getVersionInfo:               getVersionInfo,
		checkDeviceOnlineEx:          checkDeviceOnlineEx,
		setIOStatus:                  setIOStatus,
		buzzerAlarm:                  buzzerAlarm,
		acquireImage:                 acquireImage,
		initIDCardEx:                 initIDCardEx,
		initIDCardSN:                 initIDCardSN,
		setIDCardRejectType:          setIDCardRejectType,
		setRecogChipCardAttribute:    setRecogChipCardAttribute,
		reRecogMRZbyVI:               reRecogMRZbyVI,
		setBGSubtraction:             setBGSubtraction,
		setAcquireImageResolution:    setAcquireImageResolution,
		settAcquireImageExposureTime: setAcquireImageExposureTime,
		checkUVDull:                  checkUVDull,
		fibreDetect:                  fibreDetect,
		getFibrePos:                  getFibrePos,
		getImageSourceType:           getImageSourceType,
		getAnalyseMRZFieldName:       getAnalyseMRZFieldName,
		getAnalyseMRZResult:          getAnalyseMRZResult,
		recogIDCardEX:                recogIDCardEX,
		getDataGroupContent:          getDataGroupContent,
		classifyIDCard:               classifyIDCard,
	}, nil
}

func (w *IDCardDLL) Free() error {
	return w.dll.Release()
}

func (w *IDCardDLL) InitIDCard(lpUserID string, nType int, lpDirectory string) (int, error) {
	userID, err := syscall.UTF16PtrFromString(lpUserID)
	if err != nil {
		return -1, err
	}

	directory, err := syscall.UTF16PtrFromString(lpDirectory)
	if err != nil {
		return -1, err
	}

	ret, _, _ := w.initIDCard.Call(
		uintptr(unsafe.Pointer(userID)),
		uintptr(nType),
		uintptr(unsafe.Pointer(directory)),
	)

	return int(ret), nil
}

func (w *IDCardDLL) FreeIDCard() {
	w.freeIDCard.Call()
}

func (w *IDCardDLL) DetectDocument() int {
	ret, _, _ := w.detectDocument.Call()
	return int(ret)
}

func (w *IDCardDLL) AutoProcessIDCard(nType *int) int {
	ret, _, _ := w.autoProcessIDCard.Call(uintptr(unsafe.Pointer(nType)))
	return int(ret)
}

func (w *IDCardDLL) GetFieldNameEx(nAttribute int, nIndex int, lpBuffer *[]uint16, nBufferLen *int) int {
	ret, _, _ := w.getFieldNameEx.Call(
		uintptr(nAttribute),
		uintptr(nIndex),
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(unsafe.Pointer(nBufferLen)),
	)

	return int(ret)
}

func (w *IDCardDLL) GetRecogResultEx(nAttribute int, nIndex int, lpBuffer *[]uint16, nBufferLen *int) int {
	ret, _, _ := w.getRecogResultEx.Call(
		uintptr(nAttribute),
		uintptr(nIndex),
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(unsafe.Pointer(nBufferLen)),
	)

	return int(ret)
}

func (w *IDCardDLL) GetResultTypeEx(nAttribute int, nIndex int) int {
	ret, _, _ := w.getResultTypeEx.Call(
		uintptr(nAttribute),
		uintptr(nIndex),
	)

	return int(ret)
}

func (w *IDCardDLL) GetFieldConfEx(nAttribute int, nIndex int) int {
	ret, _, _ := w.getFieldConfEx.Call(
		uintptr(nAttribute),
		uintptr(nIndex),
	)

	return int(ret)
}

func (w *IDCardDLL) GetIDCardName(lpBuffer *[]uint16, nBufferLen *int) int {
	ret, _, _ := w.getIDCardName.Call(
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(unsafe.Pointer(nBufferLen)),
	)

	return int(ret)
}

func (w *IDCardDLL) GetSubID() int {
	ret, _, _ := w.getSubID.Call()
	return int(ret)
}

func (w *IDCardDLL) SaveImageEx(lpFileName string, nType int) (int, error) {
	fileName, err := syscall.UTF16PtrFromString(lpFileName)
	if err != nil {
		return -1, err
	}

	ret, _, _ := w.saveImageEx.Call(
		uintptr(unsafe.Pointer(fileName)),
		uintptr(nType),
	)

	return int(ret), nil
}

func (w *IDCardDLL) SetBarCodeMode(bBarCodeMode, bCellPhoneBarCodeCheck bool) {
	bbcm := uintptr(0)
	if bBarCodeMode {
		bbcm = 1
	}

	bcpbcc := uintptr(0)
	if bCellPhoneBarCodeCheck {
		bcpbcc = 1
	}

	w.setBarCodeMode.Call(bbcm, bcpbcc)
}

func (w *IDCardDLL) SetRecogCodeTypes(nMode int, nTypes uint) {
	w.setRecogCodeTypes.Call(
		uintptr(nMode),
		uintptr(nTypes),
	)
}

func (w *IDCardDLL) RecogCellPhoneBarCode() int {
	ret, _, _ := w.recogCellPhoneBarCode.Call()
	return int(ret)
}

func (w *IDCardDLL) GetBarcodeCount() int {
	ret, _, _ := w.getBarcodeCount.Call()
	return int(ret)
}

func (w *IDCardDLL) GetBarcodeRecogResult(nIndex int, lpBuffer *[]uint16, nBufferLen *int, lpResultType *[]uint16, nResultTypeLen *int) int {
	ret, _, _ := w.getBarcodeResult.Call(
		uintptr(nIndex),
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(unsafe.Pointer(nBufferLen)),
		uintptr(unsafe.Pointer(&(*lpResultType)[0])),
		uintptr(unsafe.Pointer(nResultTypeLen)),
	)

	return int(ret)
}

func (w *IDCardDLL) SetConfigByFile(lpConfigFile string) (int, error) {
	configFile, err := syscall.UTF16PtrFromString(lpConfigFile)
	if err != nil {
		return -1, err
	}

	ret, _, _ := w.setConfigByFile.Call(uintptr(unsafe.Pointer(configFile)))
	return int(ret), nil
}

func (w *IDCardDLL) SetLanguage(nType int) int {
	ret, _, _ := w.setLanguage.Call(uintptr(nType))
	return int(ret)
}

func (w *IDCardDLL) SetSaveImageType(nImageType int) {
	w.setSaveImageType.Call(uintptr(nImageType))
}

func (w *IDCardDLL) SetRecogVIZ(bRecogViz bool) {
	brv := uintptr(0)
	if bRecogViz {
		brv = 1
	}

	w.setRecogVIZ.Call(brv)
}

func (w *IDCardDLL) SetRecogDG(nDG int) {
	w.setRecogDG.Call(uintptr(nDG))
}

func (w *IDCardDLL) SetAnalyseMRZ(bAnalysis bool) {
	ba := uintptr(0)
	if bAnalysis {
		ba = 1
	}

	w.setAnalyseMRZ.Call(ba)
}

func (w *IDCardDLL) ResetIDCardID() {
	w.resetIDCardID.Call()
}

func (w *IDCardDLL) AddIDCardID(nMainID int, nSubID []int, nSubIDCount int) int {
	ret, _, _ := w.addIDCardID.Call(
		uintptr(nMainID),
		uintptr(unsafe.Pointer(&nSubID[0])),
		uintptr(nSubIDCount),
	)

	return int(ret)
}

func (w *IDCardDLL) GetDeviceSN(lpBuffer *[]uint16, nBufferLen int) int {
	ret, _, _ := w.getDeviceSN.Call(
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(nBufferLen),
	)

	return int(ret)
}

func (w *IDCardDLL) GetCurrentDevice(lpBuffer *[]uint16, nBufferLen int) bool {
	ret, _, _ := w.getCurrentDevice.Call(
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(nBufferLen),
	)

	return ret == 1
}

func (w *IDCardDLL) GetVersionInfo(lpBuffer *[]uint16, nBufferLen int) bool {
	ret, _, _ := w.getVersionInfo.Call(
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(nBufferLen),
	)

	return ret == 1
}

func (w *IDCardDLL) CheckDeviceOnlineEx() int {
	ret, _, _ := w.checkDeviceOnlineEx.Call()
	return int(ret)
}

func (w *IDCardDLL) SetIOStatus(nIOType int, bOpen bool) int {
	bo := uintptr(0)
	if bOpen {
		bo = 1
	}

	ret, _, _ := w.setIOStatus.Call(uintptr(nIOType), bo)
	return int(ret)
}

func (w *IDCardDLL) BuzzerAlarm(nDuration int) int {
	ret, _, _ := w.buzzerAlarm.Call(uintptr(nDuration))
	return int(ret)
}

func (w *IDCardDLL) AcquireImage(nImageSizeType int) int {
	ret, _, _ := w.acquireImage.Call(uintptr(nImageSizeType))
	return int(ret)
}

func (w *IDCardDLL) InitIDCardEx(lpUserID string, nType int, lpDirectory string, lpDeviceName string) (int, error) {
	userID, err := syscall.UTF16PtrFromString(lpUserID)
	if err != nil {
		return -1, err
	}

	directory, err := syscall.UTF16PtrFromString(lpDirectory)
	if err != nil {
		return -1, err
	}

	deviceName, err := syscall.UTF16PtrFromString(lpDeviceName)
	if err != nil {
		return -1, err
	}

	ret, _, _ := w.initIDCardEx.Call(
		uintptr(unsafe.Pointer(userID)),
		uintptr(nType),
		uintptr(unsafe.Pointer(directory)),
		uintptr(unsafe.Pointer(deviceName)),
	)

	return int(ret), nil
}

func (w *IDCardDLL) InitIDCardSN(lpUserID string, nType int, lpDirectory string, lpDeviceSN string) (int, error) {
	userID, err := syscall.UTF16PtrFromString(lpUserID)
	if err != nil {
		return -1, err
	}

	directory, err := syscall.UTF16PtrFromString(lpDirectory)
	if err != nil {
		return -1, err
	}

	deviceSN, err := syscall.UTF16PtrFromString(lpDeviceSN)
	if err != nil {
		return -1, err
	}

	ret, _, _ := w.initIDCardSN.Call(
		uintptr(unsafe.Pointer(userID)),
		uintptr(nType),
		uintptr(unsafe.Pointer(directory)),
		uintptr(unsafe.Pointer(deviceSN)),
	)

	return int(ret), nil
}

func (w *IDCardDLL) SetIDCardRejectType(nMainID int, bSet bool) int {
	bs := uintptr(0)
	if bSet {
		bs = 1
	}

	ret, _, _ := w.setIDCardRejectType.Call(uintptr(nMainID), bs)
	return int(ret)
}

func (w *IDCardDLL) SetRecogChipCardAttribute(nReadCard int) int {
	ret, _, _ := w.setRecogChipCardAttribute.Call(uintptr(nReadCard))
	return int(ret)
}

func (w *IDCardDLL) ReRecogMRZbyVI(bFlag bool) {
	bf := uintptr(0)
	if bFlag {
		bf = 1
	}

	w.reRecogMRZbyVI.Call(bf)
}

func (w *IDCardDLL) SetBGSubtraction(nBGSub int) {
	w.setBGSubtraction.Call(uintptr(nBGSub))
}

func (w *IDCardDLL) SetAcquireImageResolution(nResolutionX int, nResolutionY int) bool {
	ret, _, _ := w.setAcquireImageResolution.Call(uintptr(nResolutionX), uintptr(nResolutionY))
	return ret == 1
}

func (w *IDCardDLL) SetAcquireImageExposureTime(nLightType int, nModel int) int {
	ret, _, _ := w.settAcquireImageExposureTime.Call(uintptr(nLightType), uintptr(nModel))
	return int(ret)
}

func (w *IDCardDLL) CheckUVDull(bForceAcquire bool, nReserve int) int {
	bfa := uintptr(0)
	if bForceAcquire {
		bfa = 1
	}

	ret, _, _ := w.checkUVDull.Call(bfa, uintptr(nReserve))
	return int(ret)
}

func (w *IDCardDLL) FibreDetect(bAcquireImage bool) int {
	bai := uintptr(0)
	if bAcquireImage {
		bai = 1
	}

	ret, _, _ := w.fibreDetect.Call(bai)
	return int(ret)
}

func (w *IDCardDLL) GetFibrePos(nIndex int, nLeft *int, nTop *int, nRight *int, nBottom *int) int {
	ret, _, _ := w.getFibrePos.Call(
		uintptr(nIndex),
		uintptr(unsafe.Pointer(nLeft)),
		uintptr(unsafe.Pointer(nTop)),
		uintptr(unsafe.Pointer(nRight)),
		uintptr(unsafe.Pointer(nBottom)),
	)

	return int(ret)
}

func (w *IDCardDLL) GetImageSourceType(nMainID int, nScale int, bAcquireImage bool) int {
	bai := uintptr(0)
	if bAcquireImage {
		bai = 1
	}

	ret, _, _ := w.getImageSourceType.Call(
		uintptr(nMainID),
		uintptr(nScale),
		bai,
	)

	return int(ret)
}

func (w *IDCardDLL) GetAnalyseMRZFieldName(nIndex int, lpBuffer *[]uint16, nBufferLen *int) int {
	ret, _, _ := w.getAnalyseMRZFieldName.Call(
		uintptr(nIndex),
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(unsafe.Pointer(nBufferLen)),
	)

	return int(ret)
}

func (w *IDCardDLL) GetAnalyseMRZResult(nIndex int, lpBuffer *[]uint16, nBufferLen *int) int {
	ret, _, _ := w.getAnalyseMRZResult.Call(
		uintptr(nIndex),
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(unsafe.Pointer(nBufferLen)),
	)

	return int(ret)
}

func (w *IDCardDLL) RecogIDCardEX(nManID int, nSUBID int) int {
	ret, _, _ := w.recogIDCardEX.Call(
		uintptr(nManID),
		uintptr(nSUBID),
	)

	return int(ret)
}

func (w *IDCardDLL) GetDataGroupContent(nDGIndex int, bRawData bool, lpBuffer *[]byte, nLength *int) int {
	brd := uintptr(0)
	if bRawData {
		brd = 1
	}

	ret, _, _ := w.getDataGroupContent.Call(
		uintptr(nDGIndex),
		brd,
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(unsafe.Pointer(nLength)),
	)

	return int(ret)
}

func (w *IDCardDLL) ClassifyIDCard(nCardType *int) int {
	ret, _, _ := w.classifyIDCard.Call(uintptr(unsafe.Pointer(nCardType)))
	return int(ret)
}
