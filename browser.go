package uaparser

var (
	ie = &itemSpec{
		name:         "IE",
		mustContains: []string{"MSIE", "rv:11.0", "Edge/12.0", "IEMobile"},
		mustNotContains: []string{
			"360SE",
			"Maxthon",
			"qihu",
			"QIHU",
			"QQBrowser",
			"QQDownload",
			"SE ",
			"MetaSr",
			"TencentTraveler",
			"Firefox",
			"Opera",
			"MAXTHON",
		},
		versionSplitters: [][]string{[]string{"MSIE ", ";"}},
	}

	firefox = &itemSpec{
		name:             "Firefox",
		mustContains:     []string{"Firefox", "firefox"},
		mustNotContains:  []string{"Seamonkey", "Opera"},
		versionSplitters: [][]string{[]string{"Firefox/", " "}},
	}

	safari = &itemSpec{
		name:         "Safari",
		mustContains: []string{"Safari", "AppleWebKit", "safari"},
		mustNotContains: []string{
			"Chrome",
			"Chromium",
			"CoolNovo",
			"Maxthon",
			"LBBROWSER",
			"QIHU",
			"QQBrowser",
			"SE ",
			"CriOS",
			"MetaSr",
			"TaoBrowser",
			"PlayStation",
			"PLAYSTATION",
			"Opera",
			"FBAN",
			"FB_IAB",
			"Android",
			"Nokia",
			"Instagram",
			"WhatsApp",
			"YaBrowser",
		},
		versionSplitters: [][]string{
			[]string{"Version/", " "},
		},
	}

	chrome = &itemSpec{
		name:         "Chrome",
		mustContains: []string{"Chrome", "CriOS"},
		mustNotContains: []string{
			"CoolNovo",
			"LBBROWSER",
			"Maxthon",
			"QIHU",
			"QQBrowser",
			"SE ",
			"Edge",
			"MetaSr",
			"TaoBrowser",
			"FBAN",
			"FB_IAB",
			"Instagram",
			"WhatsApp",
			"YaBrowser",
		},
		versionSplitters: [][]string{
			[]string{"Chrome/", " "},
			[]string{"CriOS/", " "},
		},
	}

	edge = &itemSpec{
		name:         "Edge",
		mustContains: []string{"Chrome", "Edge"},
		mustNotContains: []string{
			"CoolNovo",
			"LBBROWSER",
			"Maxthon",
			"QIHU",
			"QQBrowser",
			"SE ",
			"MetaSr",
			"TaoBrowser",
			"FBAN",
			"FB_IAB",
			"Instagram",
			"WhatsApp",
			"YaBrowser",
		},
		versionSplitters: [][]string{[]string{"Edge/", " "}},
	}

	opera = &itemSpec{
		name:            "Opera",
		mustContains:    []string{"Opera"},
		mustNotContains: []string{},
		versionSplitters: [][]string{
			[]string{"Version/", " "},
			[]string{"Opera/", " "},
		},
	}

	_360se = &itemSpec{
		name:            "360SE",
		mustContains:    []string{"360SE", "qihu", "QIHU"},
		mustNotContains: []string{},
	}

	sougou = &itemSpec{
		name:             "Sougou",
		mustContains:     []string{"SE ", "MetaSr"},
		mustNotContains:  []string{},
		versionSplitters: [][]string{[]string{"SE ", " "}},
	}

	tencent = &itemSpec{
		name:             "Tencent Traveler",
		mustContains:     []string{"TencentTraveler"},
		mustNotContains:  []string{"SE ", "MetaSr"},
		versionSplitters: [][]string{[]string{"TencentTraveler ", " "}},
	}

	qq = &itemSpec{
		name:             "QQ Browser",
		mustContains:     []string{"QQBrowser"},
		mustNotContains:  []string{},
		versionSplitters: [][]string{[]string{"QQBrowser/", " "}},
	}

	maxthon = &itemSpec{
		name:            "Maxthon",
		mustContains:    []string{"Maxthon", "MAXTHON"},
		mustNotContains: []string{},
		versionSplitters: [][]string{
			[]string{"Maxthon/", " "},
			[]string{"Maxthon ", " "},
		},
	}

	playstation = &itemSpec{
		name:         "PlayStation",
		mustContains: []string{"PlayStation", "PLAYSTATION"},
		versionSplitters: [][]string{
			[]string{"PLAYSTATION 3", ")"},
			[]string{"PlayStation 4", ""},
		},
	}

	instagram = &itemSpec{
		name: 	"Instagram",
		mustContains: []string{"Instagram"},
		versionSplitters: [][]string{
			[]string{"Instagram ", " "},
		},
	}

	facebook = &itemSpec{
		name: "Facebook",
		mustContains: []string{"FB_IAB", "FBAN", "FBAN/FBIOS"},
		versionSplitters: [][]string{
			[]string{"FBAV/", ";"},
		},
	}

	yahoo = &itemSpec{
		name: "Yahoo Link Preview",
		mustContains: []string{"Yahoo Link Preview"},

	}

	androidBrowser = &itemSpec{
		name: "Android Browser",
		mustContains: []string{"Dalvik", "AppleWebKit"},
		mustNotContains: []string{
			"Chrome",
			"Chromium",
			"CoolNovo",
			"Maxthon",
			"LBBROWSER",
			"QIHU",
			"QQBrowser",
			"SE ",
			"MetaSr",
			"TaoBrowser",
			"PlayStation",
			"PLAYSTATION",
			"Opera",
			"FBAN",
			"FB_IAB",
			"WhatsApp",
		},
		versionSplitters: [][]string{
			[]string{"Dalvik/", " "},
			[]string{"Version/", " "},
		},
	}

	whatsApp = &itemSpec{
		name: "WhatsApp",
		mustContains: []string{"WhatsApp"},
		versionSplitters: [][]string{
			[]string{"WhatsApp/", " "},
		},
	}

	yandex = &itemSpec{
		name: "Yandex Browser",
		mustContains: []string{"YaBrowser"},
		versionSplitters: [][]string{
			[]string{"YaBrowser/", " "},
		},
	}


	_BROWSERS = []*itemSpec{
		ie,
		firefox,
		safari,
		chrome,
		edge,
		opera,
		_360se,
		sougou,
		tencent,
		qq,
		maxthon,
		playstation,
		instagram,
		facebook,
		yahoo,
		androidBrowser,
		whatsApp,
		yandex,
	}
)
