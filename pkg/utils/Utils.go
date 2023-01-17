package utils

import (
	"bytes"
	"crypto/tls"
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

// 思路是如果传入的操作为空数组的话直接返回源数组，否则遍历追加并返回新数组
func AddToList(totalList []string, dealList []string) []string {
	if len(dealList) == 0 {
		return totalList
	}
	for _, val := range dealList {
		totalList = append(totalList, val)
	}
	return totalList
}
