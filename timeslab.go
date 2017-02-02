/*
   Time Slab objects

   define: a time slab is basically a string representation of resolutions on a time

   a "year" is YYYY
   a "every half year" is YYYYM6{month/6}
   a "every quarter" is YYYYM3{month/3}
   a "every bimonthly" is YYYYM2{month/2}
   a "month" is YYYYMM
   a "day" is YYYYMMDD
   a "hour" is YYYYMMDDHH
   a "every 30 min" is YYYYMMDDHHM30{min/30}
   a "every 20 min" is YYYYMMDDHHM20{min/20}
   a "every 15 min" is YYYYMMDDHHM15{min/15}
   a "every 10 min" is YYYYMMDDHHM10{min/10}
   a "every 5 min" is YYYYMMDDHH{M5min/5}
   a "every min" is YYYYMMDDHHMM

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
			outStr = append(outStr, onT.Format("2006010215")+"I5"+strconv.Itoa(m))
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
