package goutils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/encoding/gcharset"
	"github.com/google/uuid"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"
)

var (
	currentPath = filepath.Dir(os.Args[0])

	ipReTemplate = `(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[1-9])\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)`
	ipAddrRe     = regexp.MustCompile(fmt.Sprintf("^%s$", ipReTemplate))
)

// 获得当前程序所在的目录
func GetCurrentProcessFileDir() string {
	return currentPath
}

// 生成 `UUID` V4 字符串
func UUIDv4() string {
	return uuid.New().String()
}

// 获得当前用户的主目录
func UserHomeDir() string {
	homedir := ""
	user, err := user.Current()
	if nil == err {
		return user.HomeDir
	}

	// cross compile support

	if "windows" == runtime.GOOS {
		homedir, _ = homeWindows()
	} else {
		homedir, _ = homeUnix()
	}
	return homedir

}

// 是否正确的IP格式
func IsValidIP(ip string) bool {
	return ipAddrRe.MatchString(ip)
}

// 中文编码转换
func GBKToUTF8(src string) (dst string) {
	var err error
	dst, err = gcharset.Convert("UTF-8", "GBK", src)
	if err != nil {
		dst = src
	}
	return
}

func UTF8ToGBK(src string) (dst string) {
	var err error
	dst, err = gcharset.Convert("GBK", "UTF-8", src)
	if err != nil {
		dst = src
	}
	return
}

// 检测文件是否存在
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 在时间范围内执行系统命令，并且将输出返回（stdout和stderr）
func ExecCmdWithTimeout(timeout time.Duration, arg ...string) ([]byte, error) {
	if len(arg) == 0 {
		return nil, errors.New("arg is zero")
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var c *exec.Cmd
	if len(arg) == 1 {
		c = exec.CommandContext(ctx, arg[0])
	} else {
		c = exec.CommandContext(ctx, arg[0], arg[1:]...)
	}
	return c.CombinedOutput()
}

func randInt(min, max int) int {
	if min >= max || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

func randStrWithMeta(n int, metaData string) string {
	bytes := []byte(metaData)
	result := []byte{}
	//r :=
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
