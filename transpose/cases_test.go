package transpose

// Source: exercism/x-common
// Commit: cda8f98 Create new exercises structure

var testCases = []struct {
	description string
	input       []string
	expected    []string
}{
	{
		"empty string",
		[]string{},
		[]string{},
	},
	{
		"two characters in a row",
		[]string{
			"A1",
		},
		[]string{
			"A",
			"1",
		},
	},
	{
		"two characters in a column",
		[]string{
			"A",
			"1",
		},
		[]string{
			"A1",
		},
	},
	{
		"simple",
		[]string{
			"ABC",
			"123",
		},
		[]string{
			"A1",
			"B2",
			"C3",
		},
	},
	{
		"single line",
		[]string{
			"Single line.",
		},
		[]string{
			"S",
			"i",
			"n",
			"g",
			"l",
			"e",
			" ",
			"l",
			"i",
			"n",
			"e",
			".",
		},
	},
	{
		"first line longer than second line",
		[]string{
			"The fourth line.",
			"The fifth line.",
		},
		[]string{
			"TT",
			"hh",
			"ee",
			"  ",
			"ff",
			"oi",
			"uf",
			"rt",
			"th",
			"h ",
			" l",
			"li",
			"in",
			"ne",
			"e.",
			".",
		},
	},
	{
		"second line longer than first line",
		[]string{
			"The first line.",
			"The second line.",
		},
		[]string{
			"TT",
			"hh",
			"ee",
			"  ",
			"fs",
			"ie",
			"rc",
			"so",
			"tn",
			" d",
			"l ",
			"il",
			"ni",
			"en",
			".e",
			" .",
		},
	},
	{
		"square",
		[]string{
			"HEART",
			"EMBER",
			"ABUSE",
			"RESIN",
			"TREND",
		},
		[]string{
			"HEART",
			"EMBER",
			"ABUSE",
			"RESIN",
			"TREND",
		},
	},
	{
		"rectangle",
		[]string{
			"FRACTURE",
			"OUTLINED",
			"BLOOMING",
			"SEPTETTE",
		},
		[]string{
			"FOBS",
			"RULE",
			"ATOP",
			"CLOT",
			"TIME",
			"UNIT",
			"RENT",
			"EDGE",
		},
	},
	{
		"triangle",
		[]string{
			"T",
			"EE",
			"AAA",
			"SSSS",
			"EEEEE",
			"RRRRRR",
		},
		[]string{
			"TEASER",
			" EASER",
			"  ASER",
			"   SER",
			"    ER",
			"     R",
		},
	},
	{
		"many lines",
		[]string{
			"Chor. Two households, both alike in dignity,",
			"In fair Verona, where we lay our scene,",
			"From ancient grudge break to new mutiny,",
			"Where civil blood makes civil hands unclean.",
			"From forth the fatal loins of these two foes",
			"A pair of star-cross'd lovers take their life;",
			"Whose misadventur'd piteous overthrows",
			"Doth with their death bury their parents' strife.",
			"The fearful passage of their death-mark'd love,",
			"And the continuance of their parents' rage,",
			"Which, but their children's end, naught could remove,",
			"Is now the two hours' traffic of our stage;",
			"The which if you with patient ears attend,",
			"What here shall miss, our toil shall strive to mend.",
		},
		[]string{
			"CIFWFAWDTAWITW",
			"hnrhr hohnhshh",
			"o oeopotedi ea",
			"rfmrmash  cn t",
			".a e ie fthow ",
			" ia fr weh,whh",
			"Trnco miae  ie",
			"w ciroitr btcr",
			"oVivtfshfcuhhe",
			" eeih a uote  ",
			"hrnl sdtln  is",
			"oot ttvh tttfh",
			"un bhaeepihw a",
			"saglernianeoyl",
			"e,ro -trsui ol",
			"h uofcu sarhu ",
			"owddarrdan o m",
			"lhg to'egccuwi",
			"deemasdaeehris",
			"sr als t  ists",
			",ebk 'phool'h,",
			"  reldi ffd   ",
			"bweso tb  rtpo",
			"oea ileutterau",
			"t kcnoorhhnatr",
			"hl isvuyee'fi ",
			" atv es iisfet",
			"ayoior trr ino",
			"l  lfsoh  ecti",
			"ion   vedpn  l",
			"kuehtteieadoe ",
			"erwaharrar,fas",
			"   nekt te  rh",
			"ismdsehphnnosa",
			"ncuse ra-tau l",
			" et  tormsural",
			"dniuthwea'g t ",
			"iennwesnr hsts",
			"g,ycoi tkrttet",
			"n ,l r s'a anr",
			"i  ef  'dgcgdi",
			"t  aol   eoe,v",
			"y  nei sl,u; e",
			",  .sf to l   ",
			"     e rv d  t",
			"     ; ie    o",
			"       f, r   ",
			"       e  e  m",
			"       .  m  e",
			"          o  n",
			"          v  d",
			"          e  .",
			"          ,",
		},
	},
}
