package idcard

import (
	"strconv"
	"time"
	"unicode"

	age "github.com/bearbin/go-age"
)

var weight = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var code = []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

// Verify 身份证 ID 是否有效
func Verify(id string) (ok bool) {
	if len(id) != 18 {
		return
	}

	sum := 0
	for i := 0; i < 17; i++ {
		if unicode.IsDigit(rune(id[i])) {
			n, err := strconv.Atoi(string(id[i]))
			if err != nil {
				return
			}
			sum += n * weight[i]
		} else {
			return
		}
	}

	if string(id[17]) == code[sum%11] {
		ok = true
	}
	return
}

func getBirthDay(id string) (time.Time, error) {
	return time.Parse("20060102", id[6:14])
}

// IDCard 身份证信息
type IDCard struct {
	Valid    bool      `json:"valid"`
	Birthday time.Time `json:"birthday"`
}

// Info 给出身份证上的信息
func Info(id string) *IDCard {
	var info = &IDCard{}

	info.Valid = Verify(id)
	info.Birthday, _ = getBirthDay(id)

	return info
}

// Age 计算年龄
func (i *IDCard) Age() int {
	return age.Age(i.Birthday)
}
