package lib

type IDCard interface {
	Free() error
	InitIDCard(lpUserID string, nType int, lpDirectory string) (int, error)
	FreeIDCard()
	DetectDocument() int
	AutoProcessIDCard(nType *int) int
	GetFieldNameEx(nAttribute int, nIndex int, lpBuffer *[]uint16, nBufferLen *int) int
	GetRecogResultEx(nAttribute int, nIndex int, lpBuffer *[]uint16, nBufferLen *int) int
	GetResultTypeEx(nAttribute int, nIndex int) int
	GetFieldConfEx(nAttribute int, nIndex int) int
	GetIDCardName(lpBuffer *[]uint16, nBufferLen *int) int
	GetSubID() int
	SaveImageEx(lpFileName string, nType int) (int, error)
	SetBarCodeMode(bBarCodeMode, bCellPhoneBarCodeCheck bool)
	SetRecogCodeTypes(nMode int, nTypes uint)
	RecogCellPhoneBarCode() int
	GetBarcodeCount() int
	GetBarcodeRecogResult(nIndex int, lpBuffer *[]uint16, nBufferLen *int, lpResultType *[]uint16, nResultTypeLen *int) int
	SetConfigByFile(lpConfigFile string) (int, error)
	SetLanguage(nType int) int
	SetSaveImageType(nImageType int)
	SetRecogVIZ(bRecogViz bool)
	SetRecogDG(nDG int)
	SetAnalyseMRZ(bAnalysis bool)
	ResetIDCardID()
	AddIDCardID(nMainID int, nSubID []int, nSubIDCount int) int
	GetDeviceSN(lpBuffer *[]uint16, nBufferLen int) int
	GetCurrentDevice(lpBuffer *[]uint16, nBufferLen int) bool
	GetVersionInfo(buffer *[]uint16, bufferLen int) bool
	CheckDeviceOnlineEx() int
	SetIOStatus(nIOType int, bOpen bool) int
	BuzzerAlarm(nDuration int) int
	AcquireImage(nImageSizeType int) int
	InitIDCardEx(lpUserID string, nType int, lpDirectory string, lpDeviceName string) (int, error)
	InitIDCardSN(lpUserID string, nType int, lpDirectory string, lpDeviceSN string) (int, error)
	SetIDCardRejectType(nMainID int, bSet bool) int
	SetRecogChipCardAttribute(nReadCard int) int
	ReRecogMRZbyVI(bFlag bool)
	SetBGSubtraction(nBGSub int)
	SetAcquireImageResolution(nResolutionX int, nResolutionY int) bool
	SetAcquireImageExposureTime(nLightType int, nModel int) int
	CheckUVDull(bForceAcquire bool, nReserve int) int
	FibreDetect(bAcquireImage bool) int
	GetFibrePos(nIndex int, nLeft *int, nTop *int, nRight *int, nBottom *int) int
	GetImageSourceType(nMainID int, nScale int, bAcquireImage bool) int
	GetAnalyseMRZFieldName(nIndex int, lpBuffer *[]uint16, nBufferLen *int) int
	GetAnalyseMRZResult(nIndex int, lpBuffer *[]uint16, nBufferLen *int) int
	RecogIDCardEX(nManID int, nSUBID int) int
	GetDataGroupContent(nDGIndex int, bRawData bool, lpBuffer *[]byte, nLength *int) int
	ClassifyIDCard(nCardType *int) int
}
