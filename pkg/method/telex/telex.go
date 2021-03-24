package telex

// This package implement Telex input method
// https://en.wikipedia.org/wiki/Telex_(input_method)
// https://www.unikey.org/support/ukmanual.html#telex

var (
	diacritics = map[string]string{
		"aw": "ă",
		"aa": "â",
		"dd": "đ",
		"ee": "ê",
		"oo": "ô",
		"ow": "ơ",
		"uw": "ư",
		"w":  "ư",
	}

	tones = map[string][]string{
		"a": []string{"a", "à", "á", "ả", "ã", "ạ"},
		"ă": []string{"ă", "ằ", "ắ", "ẳ", "ẵ", "ặ"},
		"â": []string{"â", "ầ", "ấ", "ẩ", "ẫ", "ậ"},
		"e": []string{"e", "è", "é", "ẻ", "ẽ", "ẹ"},
		"ê": []string{"ê", "ề", "ế", "ể", "ễ", "ệ"},
		"i": []string{"i", "ì", "í", "ỉ", "ĩ", "ị"},
		"o": []string{"o", "ò", "ó", "ỏ", "õ", "ọ"},
		"ô": []string{"ô", "ồ", "ố", "ổ", "ỗ", "ộ"},
		"ơ": []string{"ơ", "ờ", "ớ", "ở", "ỡ", "ợ"},
		"u": []string{"u", "ù", "ú", "ủ", "ũ", "ụ"},
		"ư": []string{"ư", "ừ", "ứ", "ử", "ữ", "ự"},
	}

	toneKeys = map[string]int{
		"z": 0, // clear previous: toansz -> toan
		"f": 1, // huyeenf -> huyền
		"s": 2, // sawcs -> sắc
		"r": 3, // hoir -> 	hỏi
		"x": 4, // ngax -> ngã
		"j": 5, // nawngj -> nặng
	}
)

func Convert(raw string) string {
	return ""
}

// Check diacritics first, then tones later
func combine(base, extra string) string {
	result, ok := diacritics[base+extra]
	if ok {
		return result
	}

	baseTones, ok := tones[base]
	if !ok {
		return base + extra
	}

	toneKey, ok := toneKeys[extra]
	if !ok {
		return base + extra
	}

	return baseTones[toneKey]
}
