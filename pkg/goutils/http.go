package goutils

import (
	"fmt"
	"net/url"
	"strings"
)

// 安全添加URI，如果uri已存在则不追加
func SafeAddUri(target, uri string, check string) (result string) {
	u, err := url.Parse(target)
	if err != nil {
		return ""
	}
	if check == "" && strings.TrimRight(u.RequestURI(), "/") == "" {
		result = AppendUri(target, uri)
	} else if CompareIgnoreCase(u.RequestURI(), check) {
		result = target
	} else {
		result = AppendUri(target, uri)
	}

	return
}

// 追加URL，基于当前目录
func AppendUri(target, uri string) (result string) {
	return strings.TrimRight(target, "/") + uri
}

func GetHostname(target string) (hostname string) {
	u, err := url.Parse(target)
	if err != nil {
		return
	}
	hostname = fmt.Sprintf("%s://%s", u.Scheme, u.Host)
	return
}
