package lib

import (
	"errors"
	"path"
	"syscall"
	"unsafe"
)

type IDCardSO struct {
	so                           syscall.Handle
	initIDCard                   uintptr
	freeIDCard                   uintptr
	detectDocument               uintptr
	autoProcessIDCard            uintptr
	getFieldNameEx               uintptr
	getRecogResultEx             uintptr
	getResultTypeEx              uintptr
	getFieldConfEx               uintptr
	getIDCardName                uintptr
	getSubID                     uintptr
	saveImageEx                  uintptr
	setBarCodeMode               uintptr
	setRecogCodeTypes            uintptr
	recogCellPhoneBarCode        uintptr
	getBarcodeCount              uintptr
	getBarcodeResult             uintptr
	setConfigByFile              uintptr
	setLanguage                  uintptr
	setSaveImageType             uintptr
	setRecogVIZ                  uintptr
	setRecogDG                   uintptr
	setAnalyseMRZ                uintptr
	resetIDCardID                uintptr
	addIDCardID                  uintptr
	getDeviceSN                  uintptr
	getCurrentDevice             uintptr
	getVersionInfo               uintptr
	checkDeviceOnlineEx          uintptr
	setIOStatus                  uintptr
	buzzerAlarm                  uintptr
	acquireImage                 uintptr
	initIDCardEx                 uintptr
	initIDCardSN                 uintptr
	setIDCardRejectType          uintptr
	setRecogChipCardAttribute    uintptr
	reRecogMRZbyVI               uintptr
	setBGSubtraction             uintptr
	setAcquireImageResolution    uintptr
	settAcquireImageExposureTime uintptr
	checkUVDull                  uintptr
	fibreDetect                  uintptr
	getFibrePos                  uintptr
	getImageSourceType           uintptr
	getAnalyseMRZFieldName       uintptr
	getAnalyseMRZResult          uintptr
	recogIDCardEX                uintptr
	getDataGroupContent          uintptr
	classifyIDCard               uintptr
}

func NewIDCardSO(soPath string) (IDCard, error) {
	if path.Ext(soPath) != ".so" {
		return nil, errors.New("invalid SO file extension")
	}

	so, err := syscall.LoadLibrary(soPath)
	if err != nil {
		return nil, err
	}

	initIDCard, err := syscall.GetProcAddress(so, "InitIDCard")
	if err != nil {
		return nil, err
	}

	freeIDCard, err := syscall.GetProcAddress(so, "FreeIDCard")
	if err != nil {
		return nil, err
	}

	detectDocument, err := syscall.GetProcAddress(so, "DetectDocument")
	if err != nil {
		return nil, err
	}

	autoProcessIDCard, err := syscall.GetProcAddress(so, "AutoProcessIDCard")
	if err != nil {
		return nil, err
	}

	getFieldNameEx, err := syscall.GetProcAddress(so, "GetFieldNameEx")
	if err != nil {
		return nil, err
	}

	getRecogResultEx, err := syscall.GetProcAddress(so, "GetRecogResultEx")
	if err != nil {
		return nil, err
	}

	getResultTypeEx, err := syscall.GetProcAddress(so, "GetResultTypeEx")
	if err != nil {
		return nil, err
	}

	getFieldConfEx, err := syscall.GetProcAddress(so, "GetFieldConfEx")
	if err != nil {
		return nil, err
	}

	getIDCardName, err := syscall.GetProcAddress(so, "GetIDCardName")
	if err != nil {
		return nil, err
	}

	getSubID, err := syscall.GetProcAddress(so, "GetSubID")
	if err != nil {
		return nil, err
	}

	saveImageEx, err := syscall.GetProcAddress(so, "SaveImageEx")
	if err != nil {
		return nil, err
	}

	setBarCodeMode, err := syscall.GetProcAddress(so, "SetBarCodeMode")
	if err != nil {
		return nil, err
	}

	setRecogCodeTypes, err := syscall.GetProcAddress(so, "SetRecogCodeTypes")
	if err != nil {
		return nil, err
	}

	recogCellPhoneBarCode, err := syscall.GetProcAddress(so, "RecogCellPhoneBarCode")
	if err != nil {
		return nil, err
	}

	getBarcodeCount, err := syscall.GetProcAddress(so, "GetBarcodeCount")
	if err != nil {
		return nil, err
	}

	getBarcodeRecogResult, err := syscall.GetProcAddress(so, "GetBarcodeRecogResult")
	if err != nil {
		return nil, err
	}

	setConfigByFile, err := syscall.GetProcAddress(so, "SetConfigByFile")
	if err != nil {
		return nil, err
	}

	setLanguage, err := syscall.GetProcAddress(so, "SetLanguage")
	if err != nil {
		return nil, err
	}

	setRecogVIZ, err := syscall.GetProcAddress(so, "SetRecogVIZ")
	if err != nil {
		return nil, err
	}

	setRecogDG, err := syscall.GetProcAddress(so, "SetRecogDG")
	if err != nil {
		return nil, err
	}

	setAnalyseMRZ, err := syscall.GetProcAddress(so, "SetAnalyseMRZ")
	if err != nil {
		return nil, err
	}

	resetIDCardID, err := syscall.GetProcAddress(so, "ResetIDCardID")
	if err != nil {
		return nil, err
	}

	addIDCardID, err := syscall.GetProcAddress(so, "AddIDCardID")
	if err != nil {
		return nil, err
	}

	getDeviceSN, err := syscall.GetProcAddress(so, "GetDeviceSN")
	if err != nil {
		return nil, err
	}

	getCurrentDevice, err := syscall.GetProcAddress(so, "GetCurrentDevice")
	if err != nil {
		return nil, err
	}

	getVersionInfo, err := syscall.GetProcAddress(so, "GetVersionInfo")
	if err != nil {
		return nil, err
	}

	checkDeviceOnlineEx, err := syscall.GetProcAddress(so, "CheckDeviceOnlineEx")
	if err != nil {
		return nil, err
	}

	setIOStatus, err := syscall.GetProcAddress(so, "SetIOStatus")
	if err != nil {
		return nil, err
	}

	buzzerAlarm, err := syscall.GetProcAddress(so, "BuzzerAlarm")
	if err != nil {
		return nil, err
	}

	acquireImage, err := syscall.GetProcAddress(so, "AcquireImage")
	if err != nil {
		return nil, err
	}

	initIDCardEx, err := syscall.GetProcAddress(so, "InitIDCardEx")
	if err != nil {
		return nil, err
	}

	initIDCardSN, err := syscall.GetProcAddress(so, "InitIDCardSN")
	if err != nil {
		return nil, err
	}

	setIDCardRejectType, err := syscall.GetProcAddress(so, "SetIDCardRejectType")
	if err != nil {
		return nil, err
	}

	setRecogChipCardAttribute, err := syscall.GetProcAddress(so, "SetRecogChipCardAttribute")
	if err != nil {
		return nil, err
	}

	reRecogMRZbyVI, err := syscall.GetProcAddress(so, "ReRecogMRZbyVI")
	if err != nil {
		return nil, err
	}

	setBGSubtraction, err := syscall.GetProcAddress(so, "SetBGSubtraction")
	if err != nil {
		return nil, err
	}

	setAcquireImageResolution, err := syscall.GetProcAddress(so, "SetAcquireImageResolution")
	if err != nil {
		return nil, err
	}

	setAcquireImageExposureTime, err := syscall.GetProcAddress(so, "SetAcquireImageExposureTime")
	if err != nil {
		return nil, err
	}

	checkUVDull, err := syscall.GetProcAddress(so, "CheckUVDull")
	if err != nil {
		return nil, err
	}

	fibreDetect, err := syscall.GetProcAddress(so, "FibreDetect")
	if err != nil {
		return nil, err
	}

	getFibrePos, err := syscall.GetProcAddress(so, "GetFibrePos")
	if err != nil {
		return nil, err
	}

	getImageSourceType, err := syscall.GetProcAddress(so, "GetImageSourceType")
	if err != nil {
		return nil, err
	}

	getAnalyseMRZFieldName, err := syscall.GetProcAddress(so, "GetAnalyseMRZFieldName")
	if err != nil {
		return nil, err
	}

	getAnalyseMRZResult, err := syscall.GetProcAddress(so, "GetAnalyseMRZResult")
	if err != nil {
		return nil, err
	}

	recogIDCardEX, err := syscall.GetProcAddress(so, "RecogIDCardEX")
	if err != nil {
		return nil, err
	}

	getDataGroupContent, err := syscall.GetProcAddress(so, "GetDataGroupContent")
	if err != nil {
		return nil, err
	}

	classifyIDCard, err := syscall.GetProcAddress(so, "ClassifyIDCard")
	if err != nil {
		return nil, err
	}

	return &IDCardSO{
		so:                           so,
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

func (w *IDCardSO) Free() error {
	return syscall.FreeLibrary(w.so)
}

func (w *IDCardSO) InitIDCard(lpUserID string, nType int, lpDirectory string) (int, error) {
	userID, err := syscall.UTF16PtrFromString(lpUserID)
	if err != nil {
		return -1, err
	}

	directory, err := syscall.UTF16PtrFromString(lpDirectory)
	if err != nil {
		return -1, err
	}

	ret, _, _ := syscall.SyscallN(
		w.initIDCard,
		uintptr(unsafe.Pointer(userID)),
		uintptr(nType),
		uintptr(unsafe.Pointer(directory)),
	)

	return int(ret), nil
}

func (w *IDCardSO) FreeIDCard() {
	syscall.SyscallN(w.freeIDCard)
}

func (w *IDCardSO) DetectDocument() int {
	ret, _, _ := syscall.SyscallN(w.detectDocument)
	return int(ret)
}

func (w *IDCardSO) AutoProcessIDCard(nType *int) int {
	ret, _, _ := syscall.SyscallN(
		w.autoProcessIDCard,
		uintptr(unsafe.Pointer(nType)),
	)
	return int(ret)
}

func (w *IDCardSO) GetFieldNameEx(nAttribute int, nIndex int, lpBuffer *[]uint16, nBufferLen *int) int {
	ret, _, _ := syscall.SyscallN(
		w.getFieldNameEx,
		uintptr(nAttribute),
		uintptr(nIndex),
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(unsafe.Pointer(nBufferLen)),
	)

	return int(ret)
}

func (w *IDCardSO) GetRecogResultEx(nAttribute int, nIndex int, lpBuffer *[]uint16, nBufferLen *int) int {
	ret, _, _ := syscall.SyscallN(
		w.getRecogResultEx,
		uintptr(nAttribute),
		uintptr(nIndex),
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(unsafe.Pointer(nBufferLen)),
	)

	return int(ret)
}

func (w *IDCardSO) GetResultTypeEx(nAttribute int, nIndex int) int {
	ret, _, _ := syscall.SyscallN(
		w.getResultTypeEx,
		uintptr(nAttribute),
		uintptr(nIndex),
	)

	return int(ret)
}

func (w *IDCardSO) GetFieldConfEx(nAttribute int, nIndex int) int {
	ret, _, _ := syscall.SyscallN(
		w.getFieldConfEx,
		uintptr(nAttribute),
		uintptr(nIndex),
	)

	return int(ret)
}

func (w *IDCardSO) GetIDCardName(lpBuffer *[]uint16, nBufferLen *int) int {
	ret, _, _ := syscall.SyscallN(
		w.getIDCardName,
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(unsafe.Pointer(nBufferLen)),
	)

	return int(ret)
}

func (w *IDCardSO) GetSubID() int {
	ret, _, _ := syscall.SyscallN(w.getSubID)
	return int(ret)
}

func (w *IDCardSO) SaveImageEx(lpFileName string, nType int) (int, error) {
	fileName, err := syscall.UTF16PtrFromString(lpFileName)
	if err != nil {
		return -1, err
	}

	ret, _, _ := syscall.SyscallN(
		w.saveImageEx,
		uintptr(unsafe.Pointer(fileName)),
		uintptr(nType),
	)

	return int(ret), nil
}

func (w *IDCardSO) SetBarCodeMode(bBarCodeMode, bCellPhoneBarCodeCheck bool) {
	bbcm := uintptr(0)
	if bBarCodeMode {
		bbcm = 1
	}

	bcpbcc := uintptr(0)
	if bCellPhoneBarCodeCheck {
		bcpbcc = 1
	}

	syscall.SyscallN(w.setBarCodeMode, bbcm, bcpbcc)
}

func (w *IDCardSO) SetRecogCodeTypes(nMode int, nTypes uint) {
	syscall.SyscallN(
		w.setRecogCodeTypes,
		uintptr(nMode),
		uintptr(nTypes),
	)
}

func (w *IDCardSO) RecogCellPhoneBarCode() int {
	ret, _, _ := syscall.SyscallN(w.recogCellPhoneBarCode)
	return int(ret)
}

func (w *IDCardSO) GetBarcodeCount() int {
	ret, _, _ := syscall.SyscallN(w.getBarcodeCount)
	return int(ret)
}

func (w *IDCardSO) GetBarcodeRecogResult(nIndex int, lpBuffer *[]uint16, nBufferLen *int, lpResultType *[]uint16, nResultTypeLen *int) int {
	ret, _, _ := syscall.SyscallN(
		w.getBarcodeResult,
		uintptr(nIndex),
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(unsafe.Pointer(nBufferLen)),
		uintptr(unsafe.Pointer(&(*lpResultType)[0])),
		uintptr(unsafe.Pointer(nResultTypeLen)),
	)

	return int(ret)
}

func (w *IDCardSO) SetConfigByFile(lpConfigFile string) (int, error) {
	configFile, err := syscall.UTF16PtrFromString(lpConfigFile)
	if err != nil {
		return -1, err
	}

	ret, _, _ := syscall.SyscallN(w.setConfigByFile, uintptr(unsafe.Pointer(configFile)))
	return int(ret), nil
}

func (w *IDCardSO) SetLanguage(nType int) int {
	ret, _, _ := syscall.SyscallN(w.setLanguage, uintptr(nType))
	return int(ret)
}

func (w *IDCardSO) SetSaveImageType(nImageType int) {
	syscall.SyscallN(w.setSaveImageType, uintptr(nImageType))
}

func (w *IDCardSO) SetRecogVIZ(bRecogViz bool) {
	brv := uintptr(0)
	if bRecogViz {
		brv = 1
	}

	syscall.SyscallN(w.setRecogVIZ, brv)
}

func (w *IDCardSO) SetRecogDG(nDG int) {
	syscall.SyscallN(w.setRecogDG, uintptr(nDG))
}

func (w *IDCardSO) SetAnalyseMRZ(bAnalysis bool) {
	ba := uintptr(0)
	if bAnalysis {
		ba = 1
	}

	syscall.SyscallN(w.setAnalyseMRZ, ba)
}

func (w *IDCardSO) ResetIDCardID() {
	syscall.SyscallN(w.resetIDCardID)
}

func (w *IDCardSO) AddIDCardID(nMainID int, nSubID []int, nSubIDCount int) int {
	ret, _, _ := syscall.SyscallN(
		w.addIDCardID,
		uintptr(nMainID),
		uintptr(unsafe.Pointer(&nSubID[0])),
		uintptr(nSubIDCount),
	)

	return int(ret)
}

func (w *IDCardSO) GetDeviceSN(lpBuffer *[]uint16, nBufferLen int) int {
	ret, _, _ := syscall.SyscallN(
		w.getDeviceSN,
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(nBufferLen),
	)

	return int(ret)
}

func (w *IDCardSO) GetCurrentDevice(lpBuffer *[]uint16, nBufferLen int) bool {
	ret, _, _ := syscall.SyscallN(
		w.getCurrentDevice,
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(nBufferLen),
	)

	return ret == 1
}

func (w *IDCardSO) GetVersionInfo(lpBuffer *[]uint16, nBufferLen int) bool {
	ret, _, _ := syscall.SyscallN(
		w.getVersionInfo,
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(nBufferLen),
	)

	return ret == 1
}

func (w *IDCardSO) CheckDeviceOnlineEx() int {
	ret, _, _ := syscall.SyscallN(w.checkDeviceOnlineEx)
	return int(ret)
}

func (w *IDCardSO) SetIOStatus(nIOType int, bOpen bool) int {
	bo := uintptr(0)
	if bOpen {
		bo = 1
	}

	ret, _, _ := syscall.SyscallN(w.setIOStatus, uintptr(nIOType), bo)
	return int(ret)
}

func (w *IDCardSO) BuzzerAlarm(nDuration int) int {
	ret, _, _ := syscall.SyscallN(w.buzzerAlarm, uintptr(nDuration))
	return int(ret)
}

func (w *IDCardSO) AcquireImage(nImageSizeType int) int {
	ret, _, _ := syscall.SyscallN(w.acquireImage, uintptr(nImageSizeType))
	return int(ret)
}

func (w *IDCardSO) InitIDCardEx(lpUserID string, nType int, lpDirectory string, lpDeviceName string) (int, error) {
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

	ret, _, _ := syscall.SyscallN(
		w.initIDCardEx,
		uintptr(unsafe.Pointer(userID)),
		uintptr(nType),
		uintptr(unsafe.Pointer(directory)),
		uintptr(unsafe.Pointer(deviceName)),
	)

	return int(ret), nil
}

func (w *IDCardSO) InitIDCardSN(lpUserID string, nType int, lpDirectory string, lpDeviceSN string) (int, error) {
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

	ret, _, _ := syscall.SyscallN(
		w.initIDCardSN,
		uintptr(unsafe.Pointer(userID)),
		uintptr(nType),
		uintptr(unsafe.Pointer(directory)),
		uintptr(unsafe.Pointer(deviceSN)),
	)

	return int(ret), nil
}

func (w *IDCardSO) SetIDCardRejectType(nMainID int, bSet bool) int {
	bs := uintptr(0)
	if bSet {
		bs = 1
	}

	ret, _, _ := syscall.SyscallN(w.setIDCardRejectType, uintptr(nMainID), bs)
	return int(ret)
}

func (w *IDCardSO) SetRecogChipCardAttribute(nReadCard int) int {
	ret, _, _ := syscall.SyscallN(w.setRecogChipCardAttribute, uintptr(nReadCard))
	return int(ret)
}

func (w *IDCardSO) ReRecogMRZbyVI(bFlag bool) {
	bf := uintptr(0)
	if bFlag {
		bf = 1
	}

	syscall.SyscallN(w.reRecogMRZbyVI, bf)
}

func (w *IDCardSO) SetBGSubtraction(nBGSub int) {
	syscall.SyscallN(w.setBGSubtraction, uintptr(nBGSub))
}

func (w *IDCardSO) SetAcquireImageResolution(nResolutionX int, nResolutionY int) bool {
	ret, _, _ := syscall.SyscallN(w.setAcquireImageResolution, uintptr(nResolutionX), uintptr(nResolutionY))
	return ret == 1
}

func (w *IDCardSO) SetAcquireImageExposureTime(nLightType int, nModel int) int {
	ret, _, _ := syscall.SyscallN(w.settAcquireImageExposureTime, uintptr(nLightType), uintptr(nModel))
	return int(ret)
}

func (w *IDCardSO) CheckUVDull(bForceAcquire bool, nReserve int) int {
	bfa := uintptr(0)
	if bForceAcquire {
		bfa = 1
	}

	ret, _, _ := syscall.SyscallN(w.checkUVDull, bfa, uintptr(nReserve))
	return int(ret)
}

func (w *IDCardSO) FibreDetect(bAcquireImage bool) int {
	bai := uintptr(0)
	if bAcquireImage {
		bai = 1
	}

	ret, _, _ := syscall.SyscallN(w.fibreDetect, bai)
	return int(ret)
}

func (w *IDCardSO) GetFibrePos(nIndex int, nLeft *int, nTop *int, nRight *int, nBottom *int) int {
	ret, _, _ := syscall.SyscallN(
		w.getFibrePos,
		uintptr(nIndex),
		uintptr(unsafe.Pointer(nLeft)),
		uintptr(unsafe.Pointer(nTop)),
		uintptr(unsafe.Pointer(nRight)),
		uintptr(unsafe.Pointer(nBottom)),
	)

	return int(ret)
}

func (w *IDCardSO) GetImageSourceType(nMainID int, nScale int, bAcquireImage bool) int {
	bai := uintptr(0)
	if bAcquireImage {
		bai = 1
	}

	ret, _, _ := syscall.SyscallN(
		w.getImageSourceType,
		uintptr(nMainID),
		uintptr(nScale),
		bai,
	)

	return int(ret)
}

func (w *IDCardSO) GetAnalyseMRZFieldName(nIndex int, lpBuffer *[]uint16, nBufferLen *int) int {
	ret, _, _ := syscall.SyscallN(
		w.getAnalyseMRZFieldName,
		uintptr(nIndex),
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(unsafe.Pointer(nBufferLen)),
	)

	return int(ret)
}

func (w *IDCardSO) GetAnalyseMRZResult(nIndex int, lpBuffer *[]uint16, nBufferLen *int) int {
	ret, _, _ := syscall.SyscallN(
		w.getAnalyseMRZResult,
		uintptr(nIndex),
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(unsafe.Pointer(nBufferLen)),
	)

	return int(ret)
}

func (w *IDCardSO) RecogIDCardEX(nManID int, nSUBID int) int {
	ret, _, _ := syscall.SyscallN(
		w.recogIDCardEX,
		uintptr(nManID),
		uintptr(nSUBID),
	)

	return int(ret)
}

func (w *IDCardSO) GetDataGroupContent(nDGIndex int, bRawData bool, lpBuffer *[]byte, nLength *int) int {
	brd := uintptr(0)
	if bRawData {
		brd = 1
	}

	ret, _, _ := syscall.SyscallN(
		w.getDataGroupContent,
		uintptr(nDGIndex),
		brd,
		uintptr(unsafe.Pointer(&(*lpBuffer)[0])),
		uintptr(unsafe.Pointer(nLength)),
	)

	return int(ret)
}

func (w *IDCardSO) ClassifyIDCard(nCardType *int) int {
	ret, _, _ := syscall.SyscallN(w.classifyIDCard, uintptr(unsafe.Pointer(nCardType)))
	return int(ret)
}
