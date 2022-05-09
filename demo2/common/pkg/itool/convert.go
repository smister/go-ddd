package itool

import (
	jsoniter "github.com/json-iterator/go"
	"strconv"
)

// Convert 复制结构体
func Convert(source interface{}, to interface{}) error {
	sourceBytes, err := jsoniter.Marshal(source)
	if err != nil {
		return err
	}

	return jsoniter.Unmarshal(sourceBytes, to)
}

// ConvertUnsafe Convert出现err，则直接panic
func ConvertUnsafe(source interface{}, to interface{}) {
	if err := Convert(source, to); err != nil {
		panic(err)
	}
}

// StringToUint64 将string转成Uint64, 如果转换成为则返回0
func StringToUint64(str string) uint64 {
	udata, err := StringToUint64WithErr(str)
	if err != nil {
		return 0
	}
	return udata
}

// StringToUint64WithErr 将string转成Uint64，并返回错误信息
func StringToUint64WithErr(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

// StringToUint32 将string转成Uint32, 如果转换成为则返回0
func StringToUint32(str string) uint32 {
	udata, err := StringToUint64WithErr(str)
	if err != nil {
		return 0
	}
	return uint32(udata)
}

// StringToFloat64 将string转成float64, 如果转换成为则返回0
func StringToFloat64(str string) float64 {
	fdata, err := StringToFloat64WithErr(str)
	if err != nil {
		return 0
	}
	return fdata
}

// StringToFloat64WithErr 将string转成float64，并返回错误信息
func StringToFloat64WithErr(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

// StringToFloat32 将string转成float32, 如果转换成为则返回0
func StringToFloat32(str string) float32 {
	fdata, err := StringToFloat64WithErr(str)
	if err != nil {
		return 0
	}
	return float32(fdata)
}

// StringToFloat32WithErr 将string转成float64，并返回错误信息
func StringToFloat32WithErr(str string) (float64, error) {
	return strconv.ParseFloat(str, 32)
}
