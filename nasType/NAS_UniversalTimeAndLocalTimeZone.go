package nasType

// UniversalTimeAndLocalTimeZone 9.11.3.53
// Year Row, sBit, len = [0, 0], 8 , 8
// Month Row, sBit, len = [1, 1], 8 , 8
// Day Row, sBit, len = [2, 2], 8 , 8
// Hour Row, sBit, len = [3, 3], 8 , 8
// Minute Row, sBit, len = [4, 4], 8 , 8
// Second Row, sBit, len = [5, 5], 8 , 8
// TimeZone Row, sBit, len = [6, 6], 8 , 8
/*
From 3GPP TS 24.008 Section 10.5.3.9
    The purpose of the timezone part of this information element is to encode 
    the offset between universal time and local time in steps of 15 minutes.
    The purpose of the time part of this information element is to encode the 
    universal time at which this information element may have been sent by the 
    network. The Time Zone and Time information element is coded as shown in 
    figure 10.5.84/3GPP TS 24.008 and table 10.5.97/3GPP TS 24.008. The Time 
    Zone and Time is a type 3 information element with a length of 8 octets.
*/
type UniversalTimeAndLocalTimeZone struct {
	Iei   uint8    `json:"-"`
	Octet [7]uint8 `json:"-"`
    Year uint8 
    Month uint8 
    Day uint8 
    Hour uint8 
    Minute uint8 
    Second uint8 
    TimeZone uint8 
}


func (u *UniversalTimeAndLocalTimeZone ) DecodeNASType() error {

    u.Year = extractDecimal(u.GetYear())
    u.Month= extractDecimal(u.GetMonth())
    u.Day= extractDecimal(u.GetDay())
    u.Hour= extractDecimal(u.GetHour())
    u.Minute= extractDecimal(u.GetMinute())
    u.Second= extractDecimal(u.GetSecond())
    u.TimeZone = u.GetTimeZone()
    return nil
    
}

func extractDecimal(hex uint8) uint8 {
    return (hex & 0xf0 >> 4)  + (hex &0xf ) * 10
}

func NewUniversalTimeAndLocalTimeZone(iei uint8) (universalTimeAndLocalTimeZone *UniversalTimeAndLocalTimeZone) {
	universalTimeAndLocalTimeZone = &UniversalTimeAndLocalTimeZone{}
	universalTimeAndLocalTimeZone.SetIei(iei)
	return universalTimeAndLocalTimeZone
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
// Iei Row, sBit, len = [], 8, 8
func (a *UniversalTimeAndLocalTimeZone) GetIei() (iei uint8) {
	return a.Iei
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
// Iei Row, sBit, len = [], 8, 8
func (a *UniversalTimeAndLocalTimeZone) SetIei(iei uint8) {
	a.Iei = iei
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
// Year Row, sBit, len = [0, 0], 8 , 8
func (a *UniversalTimeAndLocalTimeZone) GetYear() (year uint8) {
	return a.Octet[0]
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
// Year Row, sBit, len = [0, 0], 8 , 8
func (a *UniversalTimeAndLocalTimeZone) SetYear(year uint8) {
	a.Octet[0] = year
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
// Month Row, sBit, len = [1, 1], 8 , 8
func (a *UniversalTimeAndLocalTimeZone) GetMonth() (month uint8) {
	return a.Octet[1]
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
// Month Row, sBit, len = [1, 1], 8 , 8
func (a *UniversalTimeAndLocalTimeZone) SetMonth(month uint8) {
	a.Octet[1] = month
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
// Day Row, sBit, len = [2, 2], 8 , 8
func (a *UniversalTimeAndLocalTimeZone) GetDay() (day uint8) {
	return a.Octet[2]
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
// Day Row, sBit, len = [2, 2], 8 , 8
func (a *UniversalTimeAndLocalTimeZone) SetDay(day uint8) {
	a.Octet[2] = day
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
// Hour Row, sBit, len = [3, 3], 8 , 8
func (a *UniversalTimeAndLocalTimeZone) GetHour() (hour uint8) {
	return a.Octet[3]
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
// Hour Row, sBit, len = [3, 3], 8 , 8
func (a *UniversalTimeAndLocalTimeZone) SetHour(hour uint8) {
	a.Octet[3] = hour
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
// Minute Row, sBit, len = [4, 4], 8 , 8
func (a *UniversalTimeAndLocalTimeZone) GetMinute() (minute uint8) {
	return a.Octet[4]
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
// Minute Row, sBit, len = [4, 4], 8 , 8
func (a *UniversalTimeAndLocalTimeZone) SetMinute(minute uint8) {
	a.Octet[4] = minute
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
// Second Row, sBit, len = [5, 5], 8 , 8
func (a *UniversalTimeAndLocalTimeZone) GetSecond() (second uint8) {
	return a.Octet[5]
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
// Second Row, sBit, len = [5, 5], 8 , 8
func (a *UniversalTimeAndLocalTimeZone) SetSecond(second uint8) {
	a.Octet[5] = second
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
// TimeZone Row, sBit, len = [6, 6], 8 , 8
func (a *UniversalTimeAndLocalTimeZone) GetTimeZone() (timeZone uint8) {
	return a.Octet[6]
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
// TimeZone Row, sBit, len = [6, 6], 8 , 8
func (a *UniversalTimeAndLocalTimeZone) SetTimeZone(timeZone uint8) {
	a.Octet[6] = timeZone
}
