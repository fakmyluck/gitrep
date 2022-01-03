package tim

import "fmt"

type Time struct {
	Year  uint16
	Mon   uint16
	Day   uint16
	Hour  uint16
	Min   uint16
	Sec   uint16
	Units int32
}

func Subtime(s, e Time) int32 {

	var sub Time

	sub.Year = e.Year - s.Year
	e.Mon += sub.Year * 12
	if e.Year < s.Year {
		fmt.Printf("\n! Time Error ![e.Year < s.Year]\n")
		//return sub.Min //vernut' ERROR!?!?!?
	}
	/////////////////////	DOBAVIT' "-1" (teper' uint16 [net "-"])
	sub.Mon = e.Mon - s.Mon
	if e.Mon > s.Mon {
		e.Day += s.returnDays()
	} else if e.Mon < s.Mon {
		fmt.Printf("\n! Time Error ! [e.Mon < s.Mon]\n")
		//return sub.Min
	}

	sub.Day = e.Day - s.Day
	if e.Day > s.Day {
		e.Hour += 24 * sub.Day
	} else if e.Day < s.Day {
		fmt.Printf("\n! Time Error ! [e.Day < s.Day]\n")
		//return sub.Min
	}

	sub.Hour = e.Hour - s.Hour
	if e.Hour > s.Hour {
		e.Min = e.Min + 60*sub.Hour
	} else if e.Hour < s.Hour {
		fmt.Printf("\n! Time Error ! [e.Hour < s.Hour] %v-%v\n", s, e)
		//return sub.Min
	}

	sub.Min = e.Min - s.Min
	if e.Min < s.Min {
		fmt.Printf("\n! Time Error ! [e.Min < s.Min]\n")
		//return sub.Min
	}
	sub.Sec = e.Sec - s.Sec

	return int32(sub.Min)
}

// func main() {
// 	start := "2.01.2021 3:30:31"
// 	end := "2.01.2021 4:37:31"

// 	startInt := toTime(start)
// 	endtInt := toTime(end)

// 	fmt.Println(startInt)
// 	fmt.Println(endtInt)

// 	sub := subtime(startInt, endtInt)
// 	fmt.Println(sub)
// 	fmt.Printf("часы: %.2f\n", float32(sub.Min)/60)
// }

func retTimeStr(n byte) string {
	switch n {
	case 0:
		return "секунды"
	case 1:
		return "минуты"
	case 2:
		return "часы"
	case 3:
		return "года"
	case 4:
		return "месяц"
	case 5:
		return "день"
	}
	return "ОШИБКА ВОЗВРАТА"
}

func checktime(tm string) {
	var e, n byte
	i := len(tm) - 1
	for ; i >= 0; i-- {
		if tm[i] != ':' && tm[i] != '.' && tm[i] != ' ' {
			e++
		} else {
			if e != 2 {
				//fmt.Printf("%v  %v\n\n", retTimeStr(n), tm[i:i+int(e)+1])
				if n == 2 && e == 1 {
				} else if n == 3 && e == 4 {
				} else {
					fmt.Printf("! Time Error ! [func checktime]\n(%v) %v\n\n", tm[(i):(i+int(e)+1)], retTimeStr(n))
				}
			}
			e = 0
			n++
		}

	}

	// if e == 0 || e > 2 {
	// 	fmt.Printf("! Time Err !\n(%v)\n\n", tm[0:e+1])
	// }
}

func ToTime(s string) Time {
	checktime(s)
	T := Time{
		Sec:  loopstring(&s),
		Min:  loopstring(&s),
		Hour: loopstring(&s),
		//	Fix dlya 00:00 (24:00)
		Year: loopstring(&s),
		Mon:  loopstring(&s),
		Day:  loopstring(&s),
	}
		if T.Hour==0{
			T.Hour=24
		}
	return T
}

func loopstring(s *string) uint16 {
	var sum uint16
	i := len(*s) - 1
	var mult uint16 = 1
	for ; i >= 0 && (*s)[i] != ' ' && (*s)[i] != ':' && (*s)[i] != '.'; i-- {
		sum += (uint16((*s)[i]) - '0') * mult
		mult *= 10
	}
	if i >= 0 {
		*s = (*s)[:i]
	}

	return sum
}

func (tm Time) returnDays() uint16 {
	var maxdays uint16
	switch tm.Mon {
	case 1:
		maxdays = 31
	case 2:
		if tm.Year%400 == 0 {
			maxdays = 29
		} else if tm.Year%100 == 0 {
			maxdays = 28
		} else if tm.Year%4 == 0 {
			maxdays = 29
		} else {
			maxdays = 28
		}
	case 3:
		maxdays = 31
	case 4:
		maxdays = 30
	case 5:
		maxdays = 31
	case 6:
		maxdays = 30
	case 7:
		maxdays = 31
	case 8:
		maxdays = 31
	case 9:
		maxdays = 30
	case 10:
		maxdays = 31
	case 11:
		maxdays = 30
	case 12:
		maxdays = 31
	default:
		maxdays = 404
		fmt.Println("Error! returndays")
	}
	return maxdays
}

func adddays(mon uint16) uint16 {
	var maxdays uint16
	if (mon+11)%48 == 0 {
		maxdays = 1
	}
	mon = mon % 12
	switch mon {
	case 0:
		maxdays += 31
	case 1:
		maxdays += 28
		// if tm.Year%400 == 0 {
		// 	maxdays = 29
		// } else if tm.Year%100 == 0 {
		// 	maxdays = 28
		// } else if tm.Year%4 == 0 {
		// 	maxdays = 29
		// } else {
		// 	maxdays = 28
		// }
	case 2:
		maxdays += 31
	case 3:
		maxdays += 30
	case 4:
		maxdays += 31
	case 5:
		maxdays += 30
	case 6:
		maxdays += 31
	case 7:
		maxdays += 31
	case 8:
		maxdays += 30
	case 9:
		maxdays += 31
	case 10:
		maxdays += 30
	case 11:
		maxdays += 31
	default:
		fmt.Println("Error! ReturnDays")
	}
	return maxdays
}

func TimeToUnits(today Time) int32 {
	start := Time{
		Day: 1, // DAY nachinaetsa s 1 a ne s 0
		Mon: 1,

		Year: 2021,
	}
	var bufferDay uint16 = today.Day - 1
	today.Mon += (today.Year - start.Year) * 12
	monvar := today.Mon - start.Mon
	for i := uint16(0); i < monvar; i++ {
		bufferDay += adddays(i)
	}
	fmt.Println("Debug Days:", bufferDay)
	units := ((int32(bufferDay)*24+int32(today.Hour))*60+int32(today.Min))*60 + int32(today.Sec)

	return units
}

//2147483648 	= 2^31
//307584000		=356*10*24*60*60
