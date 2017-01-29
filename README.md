# Timeslab

just gonna copy and paste the docs string really


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
   
   
Get a resolution from a string where the string is defined below

    mi -> Resolution_MIN
    mi5 -> Resolution_MIN5
    mi10 -> Resolution_MIN10
    mi15 -> Resolution_MIN15
    mi20 -> Resolution_MIN20
    mi30 -> Resolution_MIN30
    h -> Resolution_HOUR
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

    ToSlabRange(res Resolution, statTime time.Time, endTime time.Time) []string
    
 
    
    