/*
   Time Slab objects

   define: a time slab is basically a string representation of resolutions on a time

    a "year" is YYYY -> 2016
    a "every half year" is YYYYM6{month/6} -> 2016M6[0-1]
    a "every quarter" is YYYYM3{month/3} -> 2016M3[0-3]
    a "every bimonthly" is YYYYM2{month/2} -> 2016M2[0-6]
    a "month" is YYYYMM -> 201601
    a "day" is YYYYMMDD -> 20160123
    a "hour" is YYYYMMDDHH -> 2016012317
    a "every 2 hours" is YYYYMMDDH02{hour/2} -> 20160123H02[00-11]
    a "every 3 hours" is YYYYMMDDH03{hour/3} -> 20160123H03[0-8]
    a "every 6 hours" is YYYYMMDDH06{hour/6} -> 20160123H06[0-3]
    a "every 12 hours" is YYYYMMDDH12{hour/12} -> 20160123H12[0-1]
    a "every 30 min" is YYYYMMDDHHI30{min/30} -> 2016012317I30[0-1]
    a "every 20 min" is YYYYMMDDHHI20{min/20} -> 2016012317I20[0-2]
    a "every 15 min" is YYYYMMDDHHI15{min/15} -> 2016012317I15[0-3]
    a "every 10 min" is YYYYMMDDHHI10{min/10} -> 2016012317I10[0-5]
    a "every 5 min" is YYYYMMDDHHI5{min/5} -> 2016012317I5[00-12]
    a "every min" is YYYYMMDDHHMM -> 201601231745

*/

//go:generate protoc --go_out=. timeslab.proto
//go:generate msgp -tests -o timeslab_mspg.go --file timeslab.pb.go
//go:generate easyjson timeslab.pb.go
package timeslab

import (
	"fmt"
	"strconv"
	"time"
)

// ToResolution takes a single string and converts it to the proper ENUM value
// mi -> Resolution_MIN
// mi5 -> Resolution_MIN5
// mi10 -> Resolution_MIN10
// mi15 -> Resolution_MIN15
// mi20 -> Resolution_MIN20
// mi30 -> Resolution_MIN30
// h -> Resolution_HOUR
// h2 -> Resolution_HOUR2
// h3 -> Resolution_HOUR3
// h6 -> Resolution_HOUR6
// h12 -> Resolution_HOUR12
// d -> Resolution_DAY
// m -> Resolution_MONTH
// m2 -> Resolution_MONTH2
// m3 -> Resolution_MONTH3
// m6 -> Resolution_MONTH6
// y -> Resolution_YEAR
// a -> Resolution_ALL
//
// if not matched the default will be Resolution_HOUR
//
func ResolutionFromString(res string) Resolution {
	switch res {
	case "mi":
		return Resolution_MIN
	case "mi5":
		return Resolution_MIN5
	case "mi10":
		return Resolution_MIN10
	case "mi15":
		return Resolution_MIN15
	case "mi20":
		return Resolution_MIN20
	case "mi30":
		return Resolution_MIN30
	case "h":
		return Resolution_HOUR
	case "h2":
		return Resolution_HOUR2
	case "h3":
		return Resolution_HOUR3
	case "h6":
		return Resolution_HOUR6
	case "h12":
		return Resolution_HOUR12
	case "d":
		return Resolution_DAY
	case "w":
		return Resolution_WEEK
	case "m":
		return Resolution_MONTH
	case "m2":
		return Resolution_MONTH2
	case "m3":
		return Resolution_MONTH3
	case "m6":
		return Resolution_MONTH6
	case "y":
		return Resolution_YEAR
	case "a":
		return Resolution_ALL
	}
	return Resolution_HOUR
}

// ToSlab take a resolution and time and make it the slab the time is converted to UTC first
//
// MIN YYYYMMDDHHMM
// MIN5 YYYYMMDDHHI5{min/5}
// MIN10 YYYYMMDDHHI10{min/10}
// MIN15 YYYYMMDDHHI15{min/15}
// MIN20 YYYYMMDDHHI20{min/20}
// MIN30 YYYYMMDDHHI30{min/30}
// HOUR YYYYMMDDHH
// HOUR2 YYYYMMDDH02{hour/2}
// HOUR3 YYYYMMDDH03{hour/3}
// HOUR6 YYYYMMDDH06{hour/6}
// HOUR12 YYYYMMDDH12{hour/12}
// DAY YYYYMMDD
// MONTH YYYYMM
// MONTH2 -> YYYYM2{month / 2}
// MONTH3 -> YYYYM3{month / 3}
// MONTH6 -> YYYYM6{month / 6}
// YEAR YYYY
// ALL ALL
//
func ToSlab(res Resolution, t time.Time) string {
	useT := t.UTC()
	switch res {
	case Resolution_MIN:
		return useT.Format("200601021504")
	case Resolution_MIN5:
		m := useT.Minute() / 5
		return useT.Format("2006010215") + "I5" + fmt.Sprintf("%02d", m)
	case Resolution_MIN10:
		m := useT.Minute() / 10
		return useT.Format("2006010215") + "I10" + strconv.Itoa(m)
	case Resolution_MIN15:
		m := useT.Minute() / 15
		return useT.Format("2006010215") + "I15" + strconv.Itoa(m)
	case Resolution_MIN20:
		m := useT.Minute() / 20
		return useT.Format("2006010215") + "I20" + strconv.Itoa(m)
	case Resolution_MIN30:
		m := useT.Minute() / 30
		return useT.Format("2006010215") + "I30" + strconv.Itoa(m)
	case Resolution_HOUR:
		return useT.Format("2006010215")
	case Resolution_HOUR2:
		m := useT.Hour() / 2
		return useT.Format("20060102") + "H02" + strconv.Itoa(m)
	case Resolution_HOUR3:
		m := useT.Hour() / 3
		return useT.Format("20060102") + "H03"+ strconv.Itoa(m)
	case Resolution_HOUR6:
		m := useT.Hour() / 6
		return useT.Format("20060102") + "H06"+ strconv.Itoa(m)
	case Resolution_HOUR12:
		m := useT.Hour() / 12
		return useT.Format("20060102") + "H12" + strconv.Itoa(m)
	case Resolution_DAY:
		return useT.Format("20060102")
	case Resolution_WEEK:
		ynum, wnum := t.ISOWeek()
		return fmt.Sprintf("%04d%02d", ynum, wnum)
	case Resolution_MONTH:
		return useT.Format("200601")
	case Resolution_MONTH2:
		m := (int(useT.Month()) / 2)
		return useT.Format("2006") + "M2" + strconv.Itoa(m)
	case Resolution_MONTH3:
		m := (int(useT.Month()) / 3)
		return useT.Format("2006") + "M3" + strconv.Itoa(m)
	case Resolution_MONTH6:
		m := (int(useT.Month()) / 6)
		return useT.Format("2006") + "M6" + strconv.Itoa(m)
	case Resolution_YEAR:
		return useT.Format("2006")
	case Resolution_ALL:
		return "ALL"
	default:
		return useT.Format("2006010215")
	}
}

// ToSlabRange given a resolution and a start/end time return the list of slabs that are in the time range
// the end slab is inclusive of the slab the end time falls in
// both the start and end times will be converted to UTC
func ToSlabRange(res Resolution, sTime time.Time, eTime time.Time) []string {
	outStr := []string{}
	onT := sTime.UTC()
	useEnd := eTime.UTC()
	switch res {
	case Resolution_MIN:
		useEnd = useEnd.Add(time.Minute * 1) // need to include the end
		for onT.Before(useEnd) {
			outStr = append(outStr, onT.Format("200601021504"))
			onT = onT.Add(time.Minute * 1)
		}
		return outStr
	case Resolution_MIN5:
		useEnd = useEnd.Add(time.Minute * 5) // need to include the end
		for onT.Before(useEnd) {
			m := onT.Minute() / 5
			outStr = append(outStr, onT.Format("2006010215")+"I5"+fmt.Sprintf("%0d", m))
			onT = onT.Add(time.Minute * 5)
		}
		return outStr
	case Resolution_MIN10:
		useEnd = useEnd.Add(time.Minute * 10) // need to include the end
		for onT.Before(useEnd) {
			m := onT.Minute() / 10
			outStr = append(outStr, onT.Format("2006010215")+"I10"+strconv.Itoa(m))
			onT = onT.Add(time.Minute * 10)
		}
		return outStr
	case Resolution_MIN15:
		useEnd = useEnd.Add(time.Minute * 15) // need to include the end
		for onT.Before(useEnd) {
			m := onT.Minute() / 15
			outStr = append(outStr, onT.Format("2006010215")+"I15"+strconv.Itoa(m))
			onT = onT.Add(time.Minute * 15)
		}
		return outStr
	case Resolution_MIN20:
		useEnd = useEnd.Add(time.Minute * 20) // need to include the end
		for onT.Before(useEnd) {
			m := onT.Minute() / 20
			outStr = append(outStr, onT.Format("2006010215")+"I20"+strconv.Itoa(m))
			onT = onT.Add(time.Minute * 20)
		}
		return outStr
	case Resolution_MIN30:
		useEnd = useEnd.Add(time.Minute * 30) // need to include the end
		for onT.Before(useEnd) {
			m := onT.Minute() / 30
			outStr = append(outStr, onT.Format("2006010215")+"I30"+strconv.Itoa(m))
			onT = onT.Add(time.Minute * 30)
		}
		return outStr
	case Resolution_HOUR2:
		useEnd = useEnd.Add(time.Hour * 2) // need to include the end
		for onT.Before(useEnd) {
			m := onT.Hour() / 2
			outStr = append(outStr, onT.Format("20060102")+"H02"+fmt.Sprintf("%0d", m))
			onT = onT.Add(time.Hour * 2)
		}
		return outStr
	case Resolution_HOUR3:
		useEnd = useEnd.Add(time.Hour * 3) // need to include the end
		for onT.Before(useEnd) {
			m := onT.Hour() / 3
			outStr = append(outStr, onT.Format("20060102")+"H03"+strconv.Itoa(m))
			onT = onT.Add(time.Hour * 3)
		}
		return outStr
	case Resolution_HOUR6:
		useEnd = useEnd.Add(time.Hour * 6) // need to include the end
		for onT.Before(useEnd) {
			m := onT.Hour() / 6
			outStr = append(outStr, onT.Format("20060102")+"H06"+strconv.Itoa(m))
			onT = onT.Add(time.Hour * 6)
		}
		return outStr
	case Resolution_HOUR12:
		useEnd = useEnd.Add(time.Hour * 12) // need to include the end
		for onT.Before(useEnd) {
			m := onT.Hour() / 12
			outStr = append(outStr, onT.Format("20060102")+"H12"+strconv.Itoa(m))
			onT = onT.Add(time.Hour * 12)
		}
		return outStr
	case Resolution_DAY:
		useEnd = useEnd.AddDate(0, 0, 1) // need to include the end
		for onT.Before(useEnd) {
			outStr = append(outStr, onT.Format("20060102"))
			onT = onT.AddDate(0, 0, 1)
		}
		return outStr
	case Resolution_WEEK:
		useEnd = useEnd.AddDate(0, 0, 7) // need to include the end
		for onT.Before(useEnd) {
			ynum, wnum := onT.ISOWeek()
			outStr = append(outStr, fmt.Sprintf("%04d%02d", ynum, wnum))
			onT = onT.AddDate(0, 0, 7)
		}
		return outStr
	case Resolution_MONTH:
		useEnd = useEnd.AddDate(0, 1, 0) // need to include the end
		for onT.Before(useEnd) {
			outStr = append(outStr, onT.Format("200601"))
			onT = onT.AddDate(0, 1, 0)
		}
		return outStr
	case Resolution_MONTH2:
		useEnd = useEnd.AddDate(0, 2, 0) // need to include the end
		for onT.Before(useEnd) {
			m := (int(onT.Month()) / 2)
			outStr = append(outStr, onT.Format("2006")+"M2"+strconv.Itoa(m))
			onT = onT.AddDate(0, 2, 0)
		}
		return outStr
	case Resolution_MONTH3:
		useEnd = useEnd.AddDate(0, 3, 0) // need to include the end
		for onT.Before(eTime) {
			m := (int(onT.Month()) / 3)
			outStr = append(outStr, onT.Format("2006")+"M3"+strconv.Itoa(m))
			onT = onT.AddDate(0, 3, 0)
		}
		return outStr
	case Resolution_MONTH6:
		useEnd = useEnd.AddDate(0, 6, 0) // need to include the end
		for onT.Before(useEnd) {
			m := (int(onT.Month()) / 6)
			outStr = append(outStr, onT.Format("2006")+"M6"+strconv.Itoa(m))
			onT = onT.AddDate(0, 6, 0)
		}
		return outStr
	case Resolution_YEAR:
		useEnd = useEnd.AddDate(1, 0, 0) // need to include the end
		for onT.Before(useEnd) {
			outStr = append(outStr, onT.Format("2006"))
			onT = onT.AddDate(1, 0, 0)
		}
		return outStr
	case Resolution_ALL:
		return []string{"ALL"}

	//default is hourly
	default:
		useEnd = useEnd.Add(time.Hour) // need to include the end
		for onT.Before(eTime) {
			outStr = append(outStr, onT.Format("2006010215"))
			onT = onT.Add(time.Hour)
		}
		return outStr
	}
}
