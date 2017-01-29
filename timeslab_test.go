package timeslab

import (
	"testing"
	"time"
)

func Test_Slab_Formatting(t *testing.T) {

	ti := time.Date(2009, time.November, 10, 23, 1, 2, 0, time.UTC)

	tData := make(map[Resolution]string)
	tData[Resolution_MIN] = "200911102301"
	tData[Resolution_MIN5] = "2009111023I500"
	tData[Resolution_MIN10] = "2009111023I100"
	tData[Resolution_MIN20] = "2009111023I200"
	tData[Resolution_MIN30] = "2009111023I300"
	tData[Resolution_HOUR] = "2009111023"
	tData[Resolution_DAY] = "20091110"
	tData[Resolution_WEEK] = "200946"
	tData[Resolution_MONTH] = "200911"
	tData[Resolution_MONTH2] = "2009M25"
	tData[Resolution_MONTH3] = "2009M33"
	tData[Resolution_MONTH6] = "2009M61"
	tData[Resolution_YEAR] = "2009"

	for res, st := range tData {
		onSl := ToSlab(res, ti)
		if onSl != st {
			t.Fatalf("Invalid time slab: got: %s, wanted: %s for resolution %s", onSl, st, Resolution_name[int32(res)])
		}
	}

	ti = time.Date(2009, time.May, 30, 6, 46, 2, 0, time.UTC)
	tData = make(map[Resolution]string)
	tData[Resolution_MIN] = "200905300646"
	tData[Resolution_MIN5] = "2009053006I509"
	tData[Resolution_MIN10] = "2009053006I104"
	tData[Resolution_MIN20] = "2009053006I202"
	tData[Resolution_MIN30] = "2009053006I301"
	tData[Resolution_HOUR] = "2009053006"
	tData[Resolution_DAY] = "20090530"
	tData[Resolution_WEEK] = "200922"
	tData[Resolution_MONTH] = "200905"
	tData[Resolution_MONTH2] = "2009M22"
	tData[Resolution_MONTH3] = "2009M31"
	tData[Resolution_MONTH6] = "2009M60"
	tData[Resolution_YEAR] = "2009"

	for res, st := range tData {
		onSl := ToSlab(res, ti)
		if onSl != st {
			t.Fatalf("Invalid time slab: got: %s, wanted: %s for resolution %s", onSl, st, Resolution_name[int32(res)])
		}
	}
}
