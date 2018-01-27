// 根据手机号获取该手机号归属地、运营商信息
package mobile

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// 手机号信息
type MInfo struct {
	No7          string `json:"no7"`          // 手机号码前7位
	Province     string `json:"province"`     // 省份
	City         string `json:"city"`         // 城市
	CityAreaCode string `json:"cityAreaCode"` // 城市区域编码
	CityZipCode  string `json:"cityZipCode"`  // 城市邮政编码
	Company      string `json:"company"`      // 运营商公司
}

var (
	ErrInvalidMobileNo = errors.New("invalid mobile no")
	ErrNotFound        = errors.New("not found")
)

var mis map[string]*MInfo

// 初始化函数，全局只需要执行一次
func Init(file string) error {
	mis = make(map[string]*MInfo)
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	firstLine := true
	for scanner.Scan() {
		if firstLine {
			firstLine = false
			continue
		}

		line := strings.TrimSpace(scanner.Text())
		strArray := strings.Split(line, "\t")
		mi := MInfo{
			No7:          strArray[0],
			Province:     strArray[1],
			City:         strArray[2],
			CityAreaCode: strArray[3],
			CityZipCode:  strArray[4],
			Company:      strArray[5],
		}
		mis[mi.No7] = &mi
	}
	return nil
}

// 根据手机号获取手机归属地信息
func Search(no string) (*MInfo, error) {
	if !validMobileNo(no) {
		return nil, ErrInvalidMobileNo
	}

	k := no[:7]
	if v, ok := mis[k]; ok {
		return v, nil
	}
	return nil, ErrNotFound
}

// 简单验证下手机号码格式：
//		11位数字，即认为是格式正确
func validMobileNo(no string) bool {
	if len(no) != 11 {
		return false
	}
	if _, err := strconv.Atoi(no); err != nil {
		return false
	}
	return true
}
