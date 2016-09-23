package holidays

import "time"

func GregorianEasterSunday(year int) time.Time {
	X := year

	// Comments from german Wikipedia: https://de.wikipedia.org/wiki/Osterzyklus

	//  1. die Säkularzahl:                              K = X div 100
	K := X / 100
	//  2. die säkulare Mondschaltung:                   M = 15 + (3K + 3) div 4 − (8K + 13) div 25
	M := 15 + (3*K+3)/4 - (8*K+13)/25
	//  3. die säkulare Sonnenschaltung:                 S = 2 − (3K + 3) div 4
	S := 2 - (3*K+3)/4
	//  4. den Mondparameter:                            A = X mod 19
	A := X % 19
	//  5. den Keim für den ersten Frühlingsvollmond:    D = (19A + M) mod 30
	D := (19*A + M) % 30
	//  6. die kalendarische Korrekturgröße:             R = D div 29 + (D div 28 − D div 29) (A div 11)
	R := D/29 + (D/28-D/29)*(A/11)
	//  7. die Ostergrenze:                             OG = 21 + D − R
	OG := 21 + D - R
	//  8. den ersten Sonntag im März:                  SZ = 7 − (X + X div 4 + S) mod 7
	SZ := 7 - (X+X/4+S)%7
	//  9. die Entfernung des Ostersonntags von der
	//     Ostergrenze (Osterentfernung in Tagen):      OE = 7 − (OG − SZ) mod 7
	OE := 7 - (OG-SZ)%7
	// 10. das Datum des Ostersonntags als Märzdatum
	//     (32. März = 1. April usw.):                  OS = OG + OE
	OS := OG + OE

	var month int = 3 // March
	var day int = OS

	if day > 31 {
		day = day % 31
		month++
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}
