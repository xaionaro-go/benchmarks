package bench

import (
	"testing"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
)

var (
	testMap2 = map[string]string{
		"a": "A",
		"b": "B",
	}
	testMap64 = map[string]string{
		"00": "))",
		"01": ")!",
		"02": ")@",
		"03": ")#",
		"04": ")$",
		"05": ")%",
		"06": ")^",
		"07": ")&",
		"08": ")*",
		"09": ")(",
		"0a": ")A",
		"0b": ")B",
		"0c": ")C",
		"0d": ")D",
		"0e": ")E",
		"0f": ")F",
		"10": "!)",
		"11": "!!",
		"12": "!@",
		"13": "!#",
		"14": "!$",
		"15": "!%",
		"16": "!^",
		"17": "!&",
		"18": "!*",
		"19": "!(",
		"1a": "!A",
		"1b": "!B",
		"1c": "!C",
		"1d": "!D",
		"1e": "!E",
		"1f": "!F",
		"20": "@)",
		"21": "@!",
		"22": "@@",
		"23": "@#",
		"24": "@$",
		"25": "@%",
		"26": "@^",
		"27": "@&",
		"28": "@*",
		"29": "@(",
		"2a": "@A",
		"2b": "@B",
		"2c": "@C",
		"2d": "@D",
		"2e": "@E",
		"2f": "@F",
		"30": "#)",
		"31": "#!",
		"32": "#@",
		"33": "##",
		"34": "#$",
		"35": "#%",
		"36": "#^",
		"37": "#&",
		"38": "#*",
		"39": "#(",
		"3a": "#A",
		"3b": "#B",
		"3c": "#C",
		"3d": "#D",
		"3e": "#E",
		"3f": "#F",
	}
)

func testSwitch2(in string) string {
	switch in {
	case "a": return "A"
	case "b": return "B"
	}
	return ""
}

func testSwitch64(in string) string {
	switch in {
	case "00": return "))"
	case "01": return ")!"
	case "02": return ")@"
	case "03": return ")#"
	case "04": return ")$"
	case "05": return ")%"
	case "06": return ")^"
	case "07": return ")&"
	case "08": return ")*"
	case "09": return ")("
	case "0a": return ")A"
	case "0b": return ")B"
	case "0c": return ")C"
	case "0d": return ")D"
	case "0e": return ")E"
	case "0f": return ")F"
	case "10": return "!)"
	case "11": return "!!"
	case "12": return "!@"
	case "13": return "!#"
	case "14": return "!$"
	case "15": return "!%"
	case "16": return "!^"
	case "17": return "!&"
	case "18": return "!*"
	case "19": return "!("
	case "1a": return "!A"
	case "1b": return "!B"
	case "1c": return "!C"
	case "1d": return "!D"
	case "1e": return "!E"
	case "1f": return "!F"
	case "20": return "@)"
	case "21": return "@!"
	case "22": return "@@"
	case "23": return "@#"
	case "24": return "@$"
	case "25": return "@%"
	case "26": return "@^"
	case "27": return "@&"
	case "28": return "@*"
	case "29": return "@("
	case "2a": return "@A"
	case "2b": return "@B"
	case "2c": return "@C"
	case "2d": return "@D"
	case "2e": return "@E"
	case "2f": return "@F"
	case "30": return "#)"
	case "31": return "#!"
	case "32": return "#@"
	case "33": return "##"
	case "34": return "#$"
	case "35": return "#%"
	case "36": return "#^"
	case "37": return "#&"
	case "38": return "#*"
	case "39": return "#("
	case "3a": return "#A"
	case "3b": return "#B"
	case "3c": return "#C"
	case "3d": return "#D"
	case "3e": return "#E"
	case "3f": return "#F"
	}
	return ""
}

func BenchmarkMap2(b *testing.B) {
	var r string

	for n := 0; n < b.N; n++ {
		r = testMap2[string(n)]
	}

	helpers.Dummy(r)
}

func BenchmarkMap64(b *testing.B) {
	var r string

	for n := 0; n < b.N; n++ {
		r = testMap64[string(n)]
	}

	helpers.Dummy(r)
}

func BenchmarkSwitch2(b *testing.B) {
	var r string

	for n := 0; n < b.N; n++ {
		r = testSwitch2(string(n))
	}

	helpers.Dummy(r)
}

func BenchmarkSwitch64(b *testing.B) {
	var r string

	for n := 0; n < b.N; n++ {
		r = testSwitch64(string(n))
	}

	helpers.Dummy(r)
}
