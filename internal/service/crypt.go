package service

import (
	"math"
	"strconv"
	"strings"
)

func DecodePassword(Prl ParamPass) string {
	decPass := ""
	if len(Prl.TmK) < 8 {
		return ""
	}

	mMM := Prl.TmK[6:8]
	mMMRRTemp, _ := strconv.ParseInt(mMM, 10, 32)
	mMMRR := int(mMMRRTemp)

	if mMMRR < 10 {
		mMMRR = mMMRR + 10
	}
	Prl.PrP = Prl.PrP[1:len(Prl.PrP)]
	aa := strings.Split(Prl.PrP, "|")
	mmmm2 := Trimascii(unicodeToWin1251_UrlEncoded_str(aa))
	ss := ""
	for i := 0; i < len(aa); i++ {
		j, _ := strconv.ParseInt(aa[i], 10, 32)
		ss += string(j)
	}
	for ii := 2; ii <= 2; ii++ {
		mMM1 := mMMRR + ii
		mmmR := mmmm2
		mmmP := make([]string, 0)
		for i := 0; i < len(mmmR); i++ {

			mMM1 = mMM1 + i + 1
			Ros := mMM1 % 8
			Roi := math.Floor(float64(mMM1 / 8))
			var mMM2 int
			if Ros > 0 {
				mMM2 = int(Ros)
			} else {
				mMM2 = int(Roi)
			}
			i1 := mMM2
			i2 := 8 - mMM2
			R10A := 1
			for j := 1; j <= i1; j++ {
				R10A = R10A * 2
			}
			R10B := 1
			for j := 1; j <= i2; j++ {
				R10B = R10B * 2
			}
			//R1 := mmmR[i]
			R2, _ := strconv.ParseInt(mmmR[i], 10, 32)

			RR1 := int(math.Floor(float64(R2 / int64(R10B))))
			RR2 := RR1 * R10B
			RR3 := R2 - int64(RR2)
			RR4 := RR3 * int64(R10A)
			RR5 := RR4 + int64(RR1)
			RR8 := RR5

			for RR8 > 255 {
				RR8 = int64(math.Floor(float64(RR8 / 2)))
			}
			//R6 := string(RR8)
			mmmP = append(mmmP, strconv.Itoa(int(RR8)))
		}
		mmmm2 = mmmP

	}
	PrP := mmmm2
	n1 := searchCharIndex(PrP, 16)
	if n1 < 1 {
		n1 = searchCharIndex(PrP, 128)
	}
	n := len(PrP)
	for i := 1; i < n; i++ {
		c1, _ := strconv.ParseInt(PrP[i], 10, 32)
		if int(c1) < 33 {
			n1 = i
			break
		}
	}
	if n1 > 0 {
		i := n1
		PrP = PrP[0:i]
		if i >= 0 {
			for j := 1; j <= 16; j++ {
				PrP = append(PrP, "32")
			}
			PrP = PrP[0:16]
		}
	}
	for p := 0; p < len(PrP); p++ {
		c, _ := strconv.ParseInt(PrP[p], 10, 32)
		decPass += string(c)
	}
	return strings.TrimSpace(decPass)
}

func unicodeToWin1251_UrlEncoded_str(prp []string) []string {

	dmap := map[int64]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9, 10: 10,
		11: 11, 12: 12, 13: 13, 14: 14, 15: 15, 16: 16, 17: 17, 18: 18, 19: 19, 20: 20, 21: 21,
		22: 22, 23: 23, 24: 24, 25: 25, 26: 26, 27: 27, 28: 28, 29: 29, 30: 30, 31: 31, 32: 32,
		33: 33, 34: 34, 35: 35, 36: 36, 37: 37, 38: 38, 39: 39, 40: 40, 41: 41, 42: 42, 43: 43,
		44: 44, 45: 45, 46: 46, 47: 47, 48: 48, 49: 49, 50: 50, 51: 51, 52: 52, 53: 53, 54: 54,
		55: 55, 56: 56, 57: 57, 58: 58, 59: 59, 60: 60, 61: 61, 62: 62, 63: 63, 64: 64, 65: 65,
		66: 66, 67: 67, 68: 68, 69: 69, 70: 70, 71: 71, 72: 72, 73: 73, 74: 74, 75: 75, 76: 76,
		77: 77, 78: 78, 79: 79, 80: 80, 81: 81, 82: 82, 83: 83, 84: 84, 85: 85, 86: 86, 87: 87,
		88: 88, 89: 89, 90: 90, 91: 91, 92: 92, 93: 93, 94: 94, 95: 95, 96: 96, 97: 97, 98: 98,
		99: 99, 100: 100, 101: 101, 102: 102, 103: 103, 104: 104, 105: 105, 106: 106, 107: 107,
		108: 108, 109: 109, 110: 110, 111: 111, 112: 112, 113: 113, 114: 114, 115: 115, 116: 116,
		117: 117, 118: 118, 119: 119, 120: 120, 121: 121, 122: 122, 123: 123, 124: 124, 125: 125,
		126: 126, 127: 127, 1027: 129, 8225: 135, 1046: 198, 8222: 132, 1047: 199, 1168: 165, 1048: 200,
		1113: 154, 1049: 201, 1045: 197, 1050: 202, 1028: 170, 160: 160, 1040: 192, 1051: 203, 164: 164,
		166: 166, 167: 167, 169: 169, 171: 171, 172: 172, 173: 173, 174: 174, 1053: 205, 176: 176, 177: 177,
		1114: 156, 181: 181, 182: 182, 183: 183, 8221: 148, 187: 187, 1029: 189, 1056: 208, 1057: 209, 1058: 210,
		8364: 136, 1112: 188, 1115: 158, 1059: 211, 1060: 212, 1030: 178, 1061: 213, 1062: 214, 1063: 215, 1116: 157,
		1064: 216, 1065: 217, 1031: 175, 1066: 218, 1067: 219, 1068: 220, 1069: 221, 1070: 222, 1032: 163, 8226: 149,
		1071: 223, 1072: 224, 8482: 153, 1073: 225, 8240: 137, 1118: 162, 1074: 226, 1110: 179, 8230: 133, 1075: 227,
		1033: 138, 1076: 228, 1077: 229, 8211: 150, 1078: 230, 1119: 159, 1079: 231, 1042: 194, 1080: 232, 1034: 140,
		1025: 168, 1081: 233, 1082: 234, 8212: 151, 1083: 235, 1169: 180, 1084: 236, 1052: 204, 1085: 237, 1035: 142,
		1086: 238, 1087: 239, 1088: 240, 1089: 241, 1090: 242, 1036: 141, 1041: 193, 1091: 243, 1092: 244, 8224: 134,
		1093: 245, 8470: 185, 1094: 246, 1054: 206, 1095: 247, 1096: 248, 8249: 139, 1097: 249, 1098: 250, 1044: 196,
		1099: 251, 1111: 191, 1055: 207, 1100: 252, 1038: 161, 8220: 147, 1101: 253, 8250: 155, 1102: 254, 8216: 145,
		1103: 255, 1043: 195, 1105: 184, 1039: 143, 1026: 128, 1106: 144, 8218: 130, 1107: 131, 8217: 146, 1108: 186,
		1109: 190}

	ss := make([]string, 0)
	for i := 0; i < len(prp); i++ {
		j, _ := strconv.ParseInt(prp[i], 10, 32)
		_, ok := dmap[j]
		if !ok {
			ss = append(ss, strconv.Itoa(int(j)))
		} else {
			ss = append(ss, strconv.Itoa(int(dmap[j])))
		}

	}
	return ss
}

func Trimascii(str []string) []string {

	end := len(str)
	for j := len(str) - 1; j >= 0; j-- {
		k, _ := strconv.ParseInt(str[j], 10, 32)
		if k == 32 {
			end--
		} else {
			break
		}
	}
	return str[:end]
}

func searchCharIndex(str []string, charCode int) int {

	for i := 0; i < len(str); i++ {
		j, _ := strconv.ParseInt(str[i], 10, 32)
		if int(j) == charCode {
			return i
		}

	}
	return 0
}
