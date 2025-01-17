package webp

/*
#cgo CFLAGS: -I./internal/libwebp-1.1/
#cgo CFLAGS: -I./internal/libwebp-1.1/src/
#cgo CFLAGS: -Wno-pointer-sign -DWEBP_USE_THREAD
#cgo !windows LDFLAGS: -lm

#include <webp/encode.h>
#include <webp/mux.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

type WebPMuxError int

const (
	WebpMuxAbiVersion     = 0x0108
	WebpEncoderAbiVersion = 0x020f
)

const (
	WebpMuxOk              = WebPMuxError(C.WEBP_MUX_OK)
	WebpMuxNotFound        = WebPMuxError(C.WEBP_MUX_NOT_FOUND)
	WebpMuxInvalidArgument = WebPMuxError(C.WEBP_MUX_INVALID_ARGUMENT)
	WebpMuxBadData         = WebPMuxError(C.WEBP_MUX_BAD_DATA)
	WebpMuxMemoryError     = WebPMuxError(C.WEBP_MUX_MEMORY_ERROR)
	WebpMuxNotEnoughData   = WebPMuxError(C.WEBP_MUX_NOT_ENOUGH_DATA)
)

type WebPPicture C.WebPPicture
type WebPAnimEncoder C.WebPAnimEncoder
type WebPAnimEncoderOptions C.WebPAnimEncoderOptions
type WebPData C.WebPData
type WebPMux C.WebPMux
type WebPMuxAnimParams C.WebPMuxAnimParams
type webPConfig C.WebPConfig

func WebPDataClear(webPData *WebPData) {
	C.WebPDataClear((*C.WebPData)(unsafe.Pointer(webPData)))
}

func WebPMuxDelete(webPMux *WebPMux) {
	C.WebPMuxDelete((*C.WebPMux)(unsafe.Pointer(webPMux)))
}

func WebPPictureFree(webPPicture *WebPPicture) {
	C.WebPPictureFree((*C.WebPPicture)(unsafe.Pointer(webPPicture)))
}

func WebPAnimEncoderDelete(webPAnimEncoder *WebPAnimEncoder) {
	C.WebPAnimEncoderDelete((*C.WebPAnimEncoder)(unsafe.Pointer(webPAnimEncoder)))
}

func (wmap *WebPMuxAnimParams) SetBgcolor(v uint32) {
	((*C.WebPMuxAnimParams)(wmap)).bgcolor = (C.uint32_t)(v)
}

func (wmap *WebPMuxAnimParams) SetLoopCount(v int) {
	(*C.WebPMuxAnimParams)(wmap).loop_count = (C.int)(v)
}

func WebPPictureInit(webPPicture *WebPPicture) int {
	return int(C.WebPPictureInit((*C.WebPPicture)(unsafe.Pointer(webPPicture))))
}

func (wpp *WebPPicture) SetWidth(v int) {
	((*C.WebPPicture)(wpp)).width = (C.int)(v)
}

func (wpp *WebPPicture) SetHeight(v int) {
	((*C.WebPPicture)(wpp)).height = (C.int)(v)
}

func (wpp WebPPicture) GetWidth() int {
	return int(((C.WebPPicture)(wpp)).width)
}

func (wpp WebPPicture) GetHeight() int {
	return int(((C.WebPPicture)(wpp)).height)
}

func (wpp *WebPPicture) SetUseArgb(v int) {
	((*C.WebPPicture)(wpp)).use_argb = (C.int)(v)
}

func (wpd WebPData) GetBytes() []byte {
	return C.GoBytes(unsafe.Pointer(((C.WebPData)(wpd)).bytes), (C.int)(((C.WebPData)(wpd)).size))
}

func WebPDataInit(webPData *WebPData) {
	C.WebPDataInit((*C.WebPData)(unsafe.Pointer(webPData)))
}

func WebPConfigInitInternal(config *webPConfig) int {
	return int(C.WebPConfigInitInternal(
		(*C.WebPConfig)(unsafe.Pointer(config)),
		(C.WebPPreset)(0),
		(C.float)(75.0),
		(C.int)(WebpEncoderAbiVersion),
	))
}

func (webpCfg *webPConfig) SetLossless(v int) {
	((*C.WebPConfig)(webpCfg)).lossless = (C.int)(v)
}

func (webpCfg webPConfig) GetLossless() int {
	return int(((C.WebPConfig)(webpCfg)).lossless)
}

func (webpCfg *webPConfig) SetMethod(v int) {
	((*C.WebPConfig)(webpCfg)).method = (C.int)(v)
}

func (webpCfg *webPConfig) SetImageHint(v int) {
	((*C.WebPConfig)(webpCfg)).image_hint = (C.WebPImageHint)(v)
}

func (webpCfg *webPConfig) SetTargetSize(v int) {
	((*C.WebPConfig)(webpCfg)).target_size = (C.int)(v)
}

func (webpCfg *webPConfig) SetTargetPSNR(v float32) {
	((*C.WebPConfig)(webpCfg)).target_PSNR = (C.float)(v)
}

func (webpCfg *webPConfig) SetSegments(v int) {
	((*C.WebPConfig)(webpCfg)).segments = (C.int)(v)
}

func (webpCfg *webPConfig) SetSnsStrength(v int) {
	((*C.WebPConfig)(webpCfg)).sns_strength = (C.int)(v)
}

func (webpCfg *webPConfig) SetFilterStrength(v int) {
	((*C.WebPConfig)(webpCfg)).filter_strength = (C.int)(v)
}

func (webpCfg *webPConfig) SetFilterSharpness(v int) {
	((*C.WebPConfig)(webpCfg)).filter_sharpness = (C.int)(v)
}

func (webpCfg *webPConfig) SetAutofilter(v int) {
	((*C.WebPConfig)(webpCfg)).autofilter = (C.int)(v)
}

func (webpCfg *webPConfig) SetAlphaCompression(v int) {
	((*C.WebPConfig)(webpCfg)).alpha_compression = (C.int)(v)
}

func (webpCfg *webPConfig) SetAlphaFiltering(v int) {
	((*C.WebPConfig)(webpCfg)).alpha_filtering = (C.int)(v)
}

func (webpCfg *webPConfig) SetPass(v int) {
	((*C.WebPConfig)(webpCfg)).pass = (C.int)(v)
}

func (webpCfg *webPConfig) SetShowCompressed(v int) {
	((*C.WebPConfig)(webpCfg)).show_compressed = (C.int)(v)
}

func (webpCfg *webPConfig) SetPreprocessing(v int) {
	((*C.WebPConfig)(webpCfg)).preprocessing = (C.int)(v)
}

func (webpCfg *webPConfig) SetPartitions(v int) {
	((*C.WebPConfig)(webpCfg)).partitions = (C.int)(v)
}

func (webpCfg *webPConfig) SetPartitionLimit(v int) {
	((*C.WebPConfig)(webpCfg)).partition_limit = (C.int)(v)
}

func (webpCfg *webPConfig) SetEmulateJpegSize(v int) {
	((*C.WebPConfig)(webpCfg)).emulate_jpeg_size = (C.int)(v)
}

func (webpCfg *webPConfig) SetThreadLevel(v int) {
	((*C.WebPConfig)(webpCfg)).thread_level = (C.int)(v)
}

func (webpCfg *webPConfig) SetLowMemory(v int) {
	((*C.WebPConfig)(webpCfg)).low_memory = (C.int)(v)
}

func (webpCfg *webPConfig) SetNearLossless(v int) {
	((*C.WebPConfig)(webpCfg)).near_lossless = (C.int)(v)
}

func (webpCfg *webPConfig) SetExact(v int) {
	((*C.WebPConfig)(webpCfg)).exact = (C.int)(v)
}

func (webpCfg *webPConfig) SetUseDeltaPalette(v int) {
	((*C.WebPConfig)(webpCfg)).use_delta_palette = (C.int)(v)
}

func (webpCfg *webPConfig) SetUseSharpYuv(v int) {
	((*C.WebPConfig)(webpCfg)).use_sharp_yuv = (C.int)(v)
}

func (webpCfg *webPConfig) SetAlphaQuality(v int) {
	((*C.WebPConfig)(webpCfg)).alpha_quality = (C.int)(v)
}

func (webpCfg *webPConfig) SetFilterType(v int) {
	((*C.WebPConfig)(webpCfg)).filter_type = (C.int)(v)
}

func (webpCfg *webPConfig) SetQuality(v float32) {
	((*C.WebPConfig)(webpCfg)).quality = (C.float)(v)
}

func (encOptions *WebPAnimEncoderOptions) GetAnimParams() WebPMuxAnimParams {
	return WebPMuxAnimParams(((*C.WebPAnimEncoderOptions)(encOptions)).anim_params)
}

func (encOptions *WebPAnimEncoderOptions) SetAnimParams(v WebPMuxAnimParams) {
	((*C.WebPAnimEncoderOptions)(encOptions)).anim_params = (C.WebPMuxAnimParams)(v)
}

func (encOptions *WebPAnimEncoderOptions) SetMinimizeSize(v int) {
	((*C.WebPAnimEncoderOptions)(encOptions)).minimize_size = (C.int)(v)
}

func (encOptions *WebPAnimEncoderOptions) SetKmin(v int) {
	((*C.WebPAnimEncoderOptions)(encOptions)).kmin = (C.int)(v)
}

func (encOptions *WebPAnimEncoderOptions) SetKmax(v int) {
	((*C.WebPAnimEncoderOptions)(encOptions)).kmax = (C.int)(v)
}

func (encOptions *WebPAnimEncoderOptions) SetAllowMixed(v int) {
	((*C.WebPAnimEncoderOptions)(encOptions)).allow_mixed = (C.int)(v)
}

func (encOptions *WebPAnimEncoderOptions) SetVerbose(v int) {
	((*C.WebPAnimEncoderOptions)(encOptions)).verbose = (C.int)(v)
}

func WebPAnimEncoderOptionsInitInternal(webPAnimEncoderOptions *WebPAnimEncoderOptions) int {
	return int(C.WebPAnimEncoderOptionsInitInternal(
		(*C.WebPAnimEncoderOptions)(unsafe.Pointer(webPAnimEncoderOptions)),
		(C.int)(WebpMuxAbiVersion),
	))
}

func WebPPictureImportRGBA(data []byte, stride int, webPPicture *WebPPicture) error {
	res := int(C.WebPPictureImportRGBA(
		(*C.WebPPicture)(unsafe.Pointer(webPPicture)),
		(*C.uint8_t)(unsafe.Pointer(&data[0])),
		(C.int)(stride),
	))
	if res == 0 {
		return errors.New("error: WebPPictureImportBGRA")
	}
	return nil
}

func WebPAnimEncoderNewInternal(width, height int, webPAnimEncoderOptions *WebPAnimEncoderOptions) *WebPAnimEncoder {
	return (*WebPAnimEncoder)(C.WebPAnimEncoderNewInternal(
		(C.int)(width),
		(C.int)(height),
		(*C.WebPAnimEncoderOptions)(unsafe.Pointer(webPAnimEncoderOptions)),
		(C.int)(WebpMuxAbiVersion),
	))
}

func WebPAnimEncoderAdd(webPAnimEncoder *WebPAnimEncoder, webPPicture *WebPPicture, timestamp int, webPConfig *webPConfig) int {
	return int(C.WebPAnimEncoderAdd(
		(*C.WebPAnimEncoder)(unsafe.Pointer(webPAnimEncoder)),
		(*C.WebPPicture)(unsafe.Pointer(webPPicture)),
		(C.int)(timestamp),
		(*C.WebPConfig)(unsafe.Pointer(webPConfig)),
	))
}

func WebPAnimEncoderAssemble(webPAnimEncoder *WebPAnimEncoder, webPData *WebPData) int {
	return int(C.WebPAnimEncoderAssemble(
		(*C.WebPAnimEncoder)(unsafe.Pointer(webPAnimEncoder)),
		(*C.WebPData)(unsafe.Pointer(webPData)),
	))
}

func WebPMuxCreateInternal(webPData *WebPData, copyData int) *WebPMux {
	return (*WebPMux)(C.WebPMuxCreateInternal(
		(*C.WebPData)(unsafe.Pointer(webPData)),
		(C.int)(copyData),
		(C.int)(WebpMuxAbiVersion),
	))
}

func WebPMuxSetAnimationParams(webPMux *WebPMux, webPMuxAnimParams *WebPMuxAnimParams) WebPMuxError {
	return (WebPMuxError)(C.WebPMuxSetAnimationParams(
		(*C.WebPMux)(unsafe.Pointer(webPMux)),
		(*C.WebPMuxAnimParams)(unsafe.Pointer(webPMuxAnimParams)),
	))
}

func WebPMuxGetAnimationParams(webPMux *WebPMux, webPMuxAnimParams *WebPMuxAnimParams) WebPMuxError {
	return (WebPMuxError)(C.WebPMuxGetAnimationParams(
		(*C.WebPMux)(unsafe.Pointer(webPMux)),
		(*C.WebPMuxAnimParams)(unsafe.Pointer(webPMuxAnimParams)),
	))
}

func WebPMuxAssemble(webPMux *WebPMux, webPData *WebPData) WebPMuxError {
	return (WebPMuxError)(C.WebPMuxAssemble(
		(*C.WebPMux)(unsafe.Pointer(webPMux)),
		(*C.WebPData)(unsafe.Pointer(webPData)),
	))
}
