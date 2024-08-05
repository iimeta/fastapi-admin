package util

import "strings"

// 脱敏
func Desensitize(input string, keep ...int) string {

	if len(input) == 0 {
		return input
	}

	keepStart := 10
	keepEnd := 5

	if len(keep) > 0 {
		keepStart = keep[0]
	}

	if len(keep) > 1 {
		keepEnd = keep[1]
	}

	// 获取字符串长度
	length := len(input)

	// 如果字符串长度小于等于需要保留的字符数之和，除2后重试脱敏
	if length <= keepStart+keepEnd {
		return Desensitize(input, keepStart/2, keepEnd/2)
	}

	// 获取开头和结尾部分
	start := input[:keepStart]
	end := input[length-keepEnd:]

	// 中间部分用 '*' 替换
	middle := strings.Repeat("*", length-keepStart-keepEnd)

	// 拼接结果
	return start + middle + end
}
