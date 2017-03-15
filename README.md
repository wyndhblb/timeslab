# Timeslab

just gonna copy and paste the docs string really

NOTE: all input times are converted to UTC first


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
   
   
Get a resolution from a string where the string is defined below

    mi -> Resolution_MIN
    mi5 -> Resolution_MIN5
    mi10 -> Resolution_MIN10
    mi15 -> Resolution_MIN15
    mi20 -> Resolution_MIN20
    mi30 -> Resolution_MIN30
    h -> Resolution_HOUR
    h2 -> Resolution_HOUR2
    h3 -> Resolution_HOUR3
    h6 -> Resolution_HOUR6
    h12 -> Resolution_HOUR12
    d -> Resolution_DAY
    m -> Resolution_MONTH
    m2 -> Resolution_MONTH2
    m3 -> Resolution_MONTH3
    m6 -> Resolution_MONTH6
    y -> Resolution_YEAR
    a -> Resolution_ALL

    ResolutionFromString(string) timeslab.Resolution
    

Get the slab

    ToSlab(res Resolution, t time.Time) string
    
Get a range (inclusive) of a span of time

    ToSlabRange(res Resolution, startTime time.Time, endTime time.Time) []string
    
 
    
    