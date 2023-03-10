package utils

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 *
 * 生成指定长度的随机字符串
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param   length int16 随机字符串长度
 * @return  string
 *
 */
func RandString(length int16) string {
	var _letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	_randStr := make([]rune, length)

	for i := range _randStr {
		_randStr[i] = _letters[rand.Intn(len(_letters))]
	}

	return string(_randStr)
}

/**
 *
 * 获取当前日期
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param   void
 * @return  string
 *
 */
func CurrentDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

/**
 *
 * 获取当前时间戳
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param   void
 * @return  int64
 *
 */
func CurrentUnixStamp() int64 {
	return (time.Now().Unix())
}

/**
 *
 * 响应JSON数据
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param   c gin.Context
 * @param	data interface 响应数据结果
 * @param	errno int64
 * @param	errmsg string
 * @return  void
 *
 */
func ReturnJson(c *gin.Context, data interface{}, errno uint16, errmsg string) {
	result := map[string]interface{}{
		"errno":  errno,
		"errmsg": errmsg,
		"data":   data,
	}

	c.JSON(http.StatusOK, result)
}

/**
 *
 * interface转成string类型
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	value interface
 * @return  string
 *
 */
func TypeInterface2String(value interface{}) string {
	var newValue string
	if value == nil {
		return newValue
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		newValue = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		newValue = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		newValue = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		newValue = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		newValue = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		newValue = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		newValue = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		newValue = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		newValue = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		newValue = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		newValue = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		newValue = strconv.FormatUint(it, 10)
	case string:
		newValue = value.(string)
	case []byte:
		newValue = string(value.([]byte))
	default:
		jsonValue, _ := json.Marshal(value)
		newValue = string(jsonValue)
	}
	return newValue
}
