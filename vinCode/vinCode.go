// 汽车vin码格式校验规则
//	1、vin17个字符，在vinMap中转换成相应的数字
// 	2、vin17个字符的系数值：vinCoefficient
// 	3、将转换的17个数字和系数相乘的结果相加，用加出来和除以11，得到余数
//	4、余数作为位置值，在数组vinCode中找到对应的值，就是vin码的第9位数值
package vinCode

import "strings"

var (
	vinCoefficient []int32 = []int32{8, 7, 6, 5, 4, 3, 2, 10, 0, 9, 8, 7, 6, 5, 4, 3, 2}
	vinCode        []byte  = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'X'}

	vinMap = map[byte]int32{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7,
		'8': 8, '9': 9, 'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5, 'F': 6, 'G': 7, 'H': 8,
		'J': 1, 'K': 2, 'L': 3, 'M': 4, 'N': 5, 'P': 7, 'R': 9, 'S': 2, 'T': 3, 'U': 4,
		'V': 5, 'W': 6, 'X': 7, 'Y': 8, 'Z': 9}
)

func Verification(vin string) bool {
	if len(vin) != 17 {
		return false
	}
	vinBs := []byte(strings.ToUpper(vin))
	var value int32
	for i, b := range vinBs {
		if v, ok := vinMap[b]; ok {
			value += v * vinCoefficient[i]
		} else {
			return false
		}
	}
	return vinCode[value%11] == vinBs[8]
}
