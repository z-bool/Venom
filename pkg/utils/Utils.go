package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"os"
	"reflect"
	"regexp"
	"unsafe"
)

func FileExist(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		return false
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func GetLastTimeFrame(conn *tls.Conn, property string) []byte {
	rawInputPtr := reflect.ValueOf(conn).Elem().FieldByName(property)
	if rawInputPtr.Kind() != reflect.Struct {
		return []byte{}
	}
	val, _ := reflect.NewAt(rawInputPtr.Type(), unsafe.Pointer(rawInputPtr.UnsafeAddr())).Elem().Interface().(bytes.Buffer)
	return val.Bytes()
}

func RegxResult(regx string, body []byte) []string {
	return regexp.MustCompile(regx).FindAllString(string(body), -1)
}

func AddToList(totalList []string, dealList []string) []string {
	if len(dealList) == 0 {
		return totalList
	}
	for _, val := range dealList {
		totalList = append(totalList, val)
	}
	return totalList
}

func Md5To32(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5To16(str string) string {
	return Md5To32(str)[8:24]
}

// 去重
func RemoveDuplicateElement(arr []string) []string {
	result := make([]string, 0, len(arr))
	temp := map[string]struct{}{}
	for _, item := range arr {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
