package leetcode

func intToRoman(num int) string {
	chs := []byte{}

	for num != 0 {
		switch {
		case num >= 1000:
			chs = append(chs, 'M')
			num -= 1000
		case num >= 900:
			chs = append(chs, []byte("CM")...)
			num -= 900
		case num >= 500:
			chs = append(chs, 'D')
			num -= 500
		case num >= 400:
			chs = append(chs, []byte("CD")...)
			num -= 400
		case num >= 100:
			chs = append(chs, 'C')
			num -= 100
		case num >= 90:
			chs = append(chs, []byte("XC")...)
			num -= 90
		case num >= 50:
			chs = append(chs, 'L')
			num -= 50
		case num >= 40:
			chs = append(chs, []byte("XL")...)
			num -= 40
		case num >= 10:
			chs = append(chs, 'X')
			num -= 10
		case num == 9:
			chs = append(chs, []byte("IX")...)
			num -= 9
		case num >= 5:
			chs = append(chs, 'V')
			num -= 5
		case num == 4:
			chs = append(chs, []byte("IV")...)
			num -= 4
		default:
			for i := 0; i < num; i++ {
				chs = append(chs, 'I')
			}
			num = 0
		}
	}

	return string(chs)
}
