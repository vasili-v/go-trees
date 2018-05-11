package dltree

import (
	"testing"

	"github.com/infobloxopen/go-trees/domain"
)

type testLabelPair struct {
	s string
	n int
}

var (
	strs = []string{
		"mwjgsapga",
		"lmodylu",
		"guq",
		"iwwmnp",
		"qimge",
		"mgkbnptni",
		"dyjd",
		"voyafnndf",
		"kfgqdlci",
		"nbcsw",
		"cnvuxdm",
		"qjisqc",
		"lvhpdcsfx",
		"bphmkkkph",
		"dqhjwrm",
		"tbc",
		"equsltg",
		"akkmyd",
		"jdywmhqo",
		"lbbpb",
		"ofibpcir",
		"dew",
		"rissqkkpe",
		"hpvof",
		"ajuiujw",
		"pme",
		"xguqa",
		"ergcl",
		"ymwrn",
		"dnsgth",
		"cag",
		"urs",
		"vlhukwe",
		"kueeokb",
		"hqaomile",
		"jpto",
		"mpajx",
		"svpyt",
		"lnwtuy",
		"ysg",
		"cuwvr",
		"pheregvf",
		"sdq",
		"rii",
		"dyjaygm",
		"roijvxym",
		"yog",
		"idqb",
		"mlbfmox",
		"fjdjlcsuv",
		"pugf",
		"jqguf",
		"crdasfig",
		"ijrxmxwdi",
		"vxn",
		"ygpghc",
		"eiflm",
		"evsxgp",
		"utgyyut",
		"hyhdyap",
		"ouoggbked",
		"rwacmxrjh",
		"fgytxt",
		"iykfuox",
		"jcuvixj",
		"evm",
		"kvrbynuyb",
		"xvqakwnoj",
		"vvtnlmn",
		"abtaof",
		"abjtq",
		"uqrotvyb",
		"knokqthq",
		"qwbfcnri",
		"eaipcj",
		"dmwtmvj",
		"gooxro",
		"ftd",
		"tkvnbdhjm",
		"vapp",
		"rqx",
		"hmuq",
		"iftptxprh",
		"uhsbbv",
		"smsbvn",
		"sumlqpx",
		"bis",
		"kpg",
		"gryyxblkr",
		"osnkxapcp",
		"yktkbqb",
		"kjn",
		"xrspds",
		"gnvcr",
		"hjahw",
		"kxf",
		"mcao",
		"dbxb",
		"aucyn",
		"paspk",
		"sbmdgge",
		"klx",
		"slfklgidb",
		"ensm",
		"xxa",
		"cbbwihj",
		"vconma",
		"nuw",
		"dicne",
		"lbnwn",
		"qvqocln",
		"tmx",
		"tumkehdh",
		"gdowqpc",
		"uode",
		"kempkxn",
		"dbfhmh",
		"onynjq",
		"mwitivsbx",
		"fagfndql",
		"mdtlmil",
		"vyi",
		"biqg",
		"goqryc",
		"qljxe",
		"ieamqw",
		"moxhdgn",
		"obqkfkfeq",
		"etwilkdf",
		"hwexq",
		"mwseeoosr",
		"afdran",
		"kkcvwj",
		"mnvsi",
		"iqsy",
		"xxv",
		"enkqjhwyt",
		"aid",
		"hrcwgwdhf",
		"ksxaq",
		"pbgt",
		"joirpmvt",
		"nylirhsbu",
		"mken",
		"bhohcexwe",
		"jyjrvaw",
		"uxsjuj",
		"faeqc",
		"uqsrkf",
		"dreeupmh",
		"pdub",
		"rpwcr",
		"ngm",
		"jgjrpagq",
		"crv",
		"lffffnmk",
		"iqmsenwhi",
		"wtjvjvtfq",
		"yiwjg",
		"acg",
		"ykhmv",
		"qpay",
		"frcgwoyrm",
		"mhkevjmqb",
		"dpfxwuqgj",
		"ctlmqq",
		"hgbojpgv",
		"bygqfqm",
		"hvbojkkn",
		"itmm",
		"rtqffdqm",
		"mxllm",
		"jpvlivwo",
		"tsdojkgy",
		"vegr",
		"jrrdyagxq",
		"pvojbrnck",
		"krcvjgjh",
		"cokjboqaf",
		"owxjhkr",
		"jdmgkygm",
		"tbtwq",
		"rwvb",
		"yqoivmi",
		"cxr",
		"gprtahu",
		"tkbgw",
		"iiimmsw",
		"vqqnetb",
		"qev",
		"rgycy",
		"cclyqyelc",
		"heeebghnw",
		"ksqggphux",
		"noj",
		"woxq",
		"yeawoe",
		"elswpvj",
		"ouqtupx",
		"gabspjuyr",
		"foceebaqt",
		"yvcl",
		"nex",
		"ciujmtoga",
		"gpni",
		"pccixnaf",
		"opxdnmfff",
		"xkam",
		"oseoammjt",
		"ayvfped",
		"eiknp",
		"jyqwamgt",
		"ghlibluv",
		"dixmmgocg",
		"tqejdm",
		"cqbe",
		"xyforipdo",
		"gluhnfaib",
		"uybua",
		"crek",
		"vghpe",
		"losgbn",
		"vstfaar",
		"ipbj",
		"jghvlorx",
		"bfg",
		"cascfj",
		"xccy",
		"yjkcx",
		"vhkvhmcbj",
		"irmta",
		"vcmnrthm",
		"asq",
		"hexynuqv",
		"scqgy",
		"ggwks",
		"dft",
		"iqhdkdpfq",
		"fbmos",
		"kfmiuy",
		"rdghqlr",
		"mqch",
		"kcunubdsb",
		"pqpg",
		"pnknvdt",
		"jckne",
		"mag",
		"emxeuu",
		"mwkiwg",
		"kqg",
		"vexmrc",
		"gqtwjcofn",
		"lvfkfodn",
		"syitnyah",
		"jrluq",
		"kqbumolq",
		"pxkxvv",
		"ooyxnh",
		"bjt",
		"gld",
		"ccjocvahb",
		"senlgu",
		"fcriuvyu",
		"fabjnmkjp",
		"ijtlqsohg",
		"oxqtgvya",
		"eiybtubc",
		"dcky",
		"rbr",
		"muiwmjno",
		"iucqbcqwq",
		"vcskof",
		"vriqu",
		"oprrecjx",
		"tpxerr",
		"vasj",
		"qvbpjdycl",
		"cjkdibcgm",
		"oriysbyou",
		"ibixj",
		"hfofd",
		"gsxmpvjk",
		"yvwxfsa",
		"oiacy",
		"ngxyfxyqv",
		"uxau",
		"ogci",
		"xekgu",
		"kdxksee",
		"vbjriye",
		"fmwjry",
		"fey",
		"hjpcbd",
		"axjuc",
		"mtb",
		"gava",
		"uvg",
		"npjyc",
		"xnge",
		"wyyk",
		"mukatfac",
		"nefesk",
		"dotvk",
		"jxbfcukjo",
		"ywoyu",
		"aqvbggcnw",
		"fnpcirpqy",
		"josvwrid",
		"wfjpm",
		"qedhba",
		"xfchgvxx",
		"gyejtadad",
		"bsephp",
		"rrs",
		"ein",
		"xcabmq",
		"ikpgpq",
		"vvpyd",
		"kbhdh",
		"pjncxysl",
		"pehnletnd",
		"vrxtggyhe",
		"ashtjprsp",
		"mkhb",
		"uiakf",
		"onrbx",
		"rml",
		"iuqcnxcw",
		"vwob",
		"xev",
		"hivw",
		"tfdcsh",
		"lly",
		"iwyrnwwd",
		"ukhumgw",
		"aeoyxl",
		"mjiegeegr",
		"ekebe",
		"vdbcbo",
		"ruwqvgg",
		"ibky",
		"cbkuft",
		"vkmq",
		"ngxsg",
		"bamdtkbaq",
		"rctkqryb",
		"cuj",
		"dmdjloq",
		"mmwsm",
		"kmntj",
		"rnoftl",
		"com",
		"xoctwvarg",
		"mbvt",
		"qljp",
		"iidm",
		"plovfy",
		"cdkndbuii",
		"dbaxrmnjm",
		"hskbyuer",
		"wnfjwfq",
		"pxfgavqwt",
		"esg",
		"tgkp",
		"lphx",
		"mqnurve",
		"mwane",
		"jlouxa",
		"vuneurw",
		"mfqiul",
		"lfvcyr",
		"nfntdsyb",
		"dlgca",
		"lili",
		"sab",
		"dphqoqjo",
		"naerwrpg",
		"gjxxp",
		"lavpx",
		"ubqghenb",
		"jrsf",
		"bbutgf",
		"fvji",
		"ahgdshyd",
		"qscitssb",
		"xsaqcf",
		"waxf",
		"spvwcqf",
		"fqsqfphd",
		"xnjbnfnbd",
		"ynqvsoajf",
		"shajlgswg",
		"qdufrfud",
		"byxn",
		"uklnnul",
		"hfjb",
		"smq",
		"wnoibqe",
		"wksnk",
		"aurivno",
		"vnlrfk",
		"imlwj",
		"eefbsobd",
		"ktp",
		"eeehcxk",
		"kbalaamqh",
		"wca",
		"qryenmj",
		"jbcla",
		"kkgje",
		"dpslvlsp",
		"glewi",
		"atsrur",
		"uhduua",
		"tghvvor",
		"fby",
		"ncj",
		"dxvkwvvlj",
		"vjhwq",
		"pmbtkrg",
		"uchsdri",
		"jeywae",
		"xhlt",
		"bwhhp",
		"fcjhmbyd",
		"ivharao",
		"vmxsah",
		"jjnya",
		"gkbjlsdgd",
		"wmulligve",
		"nnydois",
		"wnhuwxy",
		"lflqriy",
		"uakkxtju",
		"ioyikkfvf",
		"bgt",
		"drfmjx",
		"rafhdc",
		"jakttf",
		"vvd",
		"gtpty",
		"phik",
		"htqcon",
		"ganr",
		"poupnt",
		"uvxaegr",
		"plxlrgbrd",
		"jcidfeh",
		"xvjg",
		"wjcyvse",
		"vooqmoa",
		"ttkt",
		"gkw",
		"dixhmnv",
		"vtpt",
		"rqovjgk",
		"uqlrewon",
		"mjm",
		"mqfnqvgsc",
		"dva",
		"hcc",
		"lvnyxispa",
		"yng",
		"kupxywtf",
		"ydpplmwmu",
		"wpkws",
		"gfniihu",
		"rkkdn",
		"boeuobq",
		"kdf",
		"coldlgx",
		"cvqchkkgl",
		"xcrbtebq",
		"cbqfofj",
		"tjsghchw",
		"qptugfsi",
		"mlbymu",
		"mqn",
		"obpc",
		"ebbnwq",
		"piw",
		"rtgepmbcp",
		"wvxjcm",
		"ljbfwqe",
		"dekewlmp",
		"jqtikbyoy",
		"ruhlufll",
		"hey",
		"uvc",
		"bldfhyxe",
		"snawajp",
		"xuexsj",
		"tdkirw",
		"kvvuhwdn",
		"qjuxrmoe",
		"xmrig",
		"meiolvc",
		"atxkstx",
		"tgvbtujl",
		"mddva",
		"qjigyiq",
		"vmvkhraot",
		"swr",
		"bld",
		"pbfgol",
		"mvaleubn",
		"okeqqimf",
		"pfepdslqw",
		"frggay",
		"mfsmt",
		"cjhy",
		"sxhktnpo",
		"gjjubygeh",
		"gxwj",
		"tikdqs",
		"yutyolwb",
		"xtput",
		"mbtl",
		"swoa",
		"rukgbb",
		"typx",
		"ksnf",
		"kelmu",
		"jsmwc",
		"mvipj",
		"odlisi",
		"uky",
		"ymlg",
		"mxuaw",
		"pobssd",
		"bqxk",
		"orosriby",
		"fyjld",
		"ocxcd",
		"wyancccrn",
		"kgrqsj",
		"lxnjknrap",
		"afnoiay",
		"gafysm",
		"yqbi",
		"vhf",
		"xqx",
		"xewcgyf",
		"vtxgkeh",
		"qsluove",
		"eley",
		"ntdf",
		"qchuqo",
		"gwlc",
		"txbgh",
		"xne",
		"odacreq",
		"mgmjj",
		"jhoeprk",
		"ywgbd",
		"igongwlb",
		"xbuenrlvh",
		"dicy",
		"ovsqg",
		"ddimnf",
		"mtwbfvpu",
		"sxikmpk",
		"huoak",
		"idxhsdgu",
		"ylfnb",
		"hed",
		"irpxvjoxm",
		"xpo",
		"okhqpidc",
		"pdlbbqydd",
		"nacsyq",
		"hlpllnkqj",
		"dkswhkay",
		"cqwp",
		"ofoeeiso",
		"mjfnn",
		"nusi",
		"otxs",
		"qeaooel",
		"grtuap",
		"sfrnca",
		"dulg",
		"tpne",
		"efr",
		"rrsb",
		"wfoiluv",
		"sqcgbpbue",
		"xlycafa",
		"idp",
		"mfyp",
		"xwjyli",
		"ovktuej",
		"gqhbipl",
		"frvcifwr",
		"deyw",
		"lcquuy",
		"fosklu",
		"fjgaoc",
		"fpl",
		"vugm",
		"ifvtqj",
		"wdwrgmauh",
		"nglr",
		"sxbg",
		"hnb",
		"yiplvbrkv",
		"trntwy",
		"xqdylhb",
		"aws",
		"wfwafe",
		"uxeieg",
		"jbttme",
		"vthwpcx",
		"yrvwqavw",
		"kuoxlsj",
		"baqyvupj",
		"xlbcujup",
		"wdwror",
		"hrpcjabv",
		"nfd",
		"yxhbb",
		"wkiuo",
		"bto",
		"tyj",
		"mkwe",
		"sfsp",
		"wllhxg",
		"fvo",
		"hisexk",
		"kaly",
		"urrdnpxn",
		"gwj",
		"lxpvvpcvm",
		"nbr",
		"ooxbb",
		"etsflvyh",
		"dvnbqjj",
		"whsxel",
		"kgxlgeduj",
		"egqqlair",
		"mvdja",
		"liggbwrq",
		"rvmdt",
		"srj",
		"upqmowpgs",
		"apkvrwrwv",
		"wcr",
		"jtfrkfs",
		"kyamaud",
		"drs",
		"agptoal",
		"bkvxhysaq",
		"srfmmy",
		"hjyru",
		"rtwka",
		"wkiakk",
		"crcdm",
		"nygv",
		"avbgr",
		"hjuueadfr",
		"filgqkur",
		"wjso",
		"ihkerpx",
		"bgcimj",
		"amelni",
		"wpovht",
		"hyo",
		"pdryhrb",
		"prjby",
		"klae",
		"yfbiq",
		"mcsclb",
		"tbhnu",
		"huuvrgksk",
		"glfhrqe",
		"wlao",
		"lowrm",
		"lmxe",
		"kveyodxpq",
		"jkdu",
		"nyuvs",
		"sypugvbfd",
		"uvsxc",
		"jbstvmof",
		"lmewqiejq",
		"kqucnx",
		"ywfyarxlh",
		"ndldyjvcs",
		"gfhaccjmx",
		"cjbonbrr",
		"buqgb",
		"xxwvm",
		"qhqen",
		"ujlx",
		"qnqygnoih",
		"dcucdfg",
		"tnrccgje",
		"lakejaah",
		"ukljlk",
		"bjgeqduub",
		"uslnrrmai",
		"ucucqy",
		"eoatldm",
		"isxcvutel",
		"wgyqdy",
		"wofkwxyko",
		"uedj",
		"cpfftgrb",
		"crlda",
		"xqrpapmt",
		"qxeh",
		"vedkoa",
		"rmc",
		"imgd",
		"uwhufq",
		"xkukyrfxm",
		"mumqd",
		"ksdm",
		"jjfddfakd",
		"tuv",
		"uaednew",
		"asqtfxtjq",
		"ikowhrf",
		"xbct",
		"ygvm",
		"jyvsyewl",
		"gwigkbol",
		"kieqtsg",
		"nylkhlek",
		"tnesl",
		"vik",
		"jfvinse",
		"iqibgioar",
		"ulsfnkid",
		"atjfgef",
		"axfe",
		"sflx",
		"ouriulke",
		"mlvu",
		"uahd",
		"qfxy",
		"oyne",
		"ofl",
		"qbeqjptg",
		"rfmkkxwvk",
		"aube",
		"keictioy",
		"emlcerqoo",
		"qifvxjn",
		"smeyap",
		"nfqds",
		"yfodyrds",
		"xovnhvmon",
		"msrsqa",
		"siaba",
		"mlxkwkook",
		"tbhsr",
		"viposah",
		"ukotwl",
		"tbifi",
		"afks",
		"kjeoequua",
		"nehftvf",
		"eify",
		"metekgow",
		"lvh",
		"lwfkqtxp",
		"eoxbe",
		"kkdinfhqf",
		"bwpvh",
		"fdyuwfax",
		"octo",
		"lowoikxba",
		"mtgg",
		"rsdin",
		"fflpiihno",
		"nmj",
		"vgvuwa",
		"ebptdv",
		"ojsftshq",
		"lcvtg",
		"hpt",
		"ukjvgna",
		"qyqvnxoj",
		"umsbftqu",
		"ujdlnxx",
		"pfppbimby",
		"lhox",
		"rmkomaljt",
		"veeagu",
		"adgia",
		"fjfr",
		"qujrkhop",
		"dnsoeg",
		"cxfonp",
		"uvqvlvha",
		"upyvjjxm",
		"txfsoj",
		"hebyc",
		"mipdxt",
		"gmkqqclm",
		"wobhpn",
		"kbuqfcfr",
		"pvkcoo",
		"kov",
		"nhuun",
		"qvqpkj",
		"owmvdf",
		"slumywk",
		"blwuxu",
		"qxb",
		"mawc",
		"aoladx",
		"cqmyt",
		"ccd",
		"hcx",
		"yht",
		"asijwfkcv",
		"oxtjfnrec",
		"qrj",
		"lxklee",
		"rpnhn",
		"jitwilxd",
		"dfhfnjwgj",
		"qiaovd",
		"uvdojrh",
		"icm",
		"uhqovev",
		"ihobyc",
		"ocxdiicah",
		"phjrhnh",
		"vbyxus",
		"qtehn",
		"otbxr",
		"bkuaxro",
		"kmsu",
		"dacapxhro",
		"ein",
		"merqpqjia",
		"uyxvpnnsb",
		"amywbp",
		"bcnmk",
		"kqaniuivr",
		"lcvfo",
		"fjigc",
		"gsedjpvx",
		"oqnxa",
		"icavfrgi",
		"rfdxrmv",
		"hdurw",
		"yxwtoaexu",
		"komctdeav",
		"lljwan",
		"bcnt",
		"jlyl",
		"apilpr",
		"rvniuodu",
		"tpefrwce",
		"qddnb",
		"aolkruqo",
		"ymuw",
		"cywwryivq",
		"pconavry",
		"admxmfi",
		"gbxfju",
		"nwkn",
		"opcm",
		"qbiii",
		"qqsos",
		"apnmva",
		"sjcsbb",
		"kqayi",
		"uprxukvd",
		"nopexvd",
		"gdorx",
		"ygtu",
		"uiwsegr",
		"hkxl",
		"yqdxarknf",
		"aqlvbfr",
		"gtcuerten",
		"fmb",
		"ugr",
		"crmva",
		"ygcqsrupr",
		"fxpixt",
		"elkqlg",
		"nvr",
		"klh",
		"voextsnn",
		"xpee",
		"gsd",
		"mvjiavtl",
		"qmwimhyv",
		"aji",
		"gercexout",
		"cpb",
		"nhxgacgj",
		"rqoer",
		"qcaqp",
		"tacieb",
		"jbgtxy",
		"gbhxoskfn",
		"sxgq",
		"drwccp",
		"kdthiu",
		"dfuwf",
		"vrlxsujj",
		"wyvyp",
		"nwvts",
		"txxw",
		"wdp",
		"mqspds",
		"qpsntsul",
		"oikfkimev",
		"wjo",
		"vhk",
		"hcyofe",
		"hanlt",
		"arvv",
		"dabutmpg",
		"dsgrvjweh",
		"iskndt",
		"ebhriqhp",
		"lkvkjl",
		"aumay",
		"hqkkxmcq",
		"ijpopwnav",
		"mcvyieiy",
		"qkpghl",
		"puvfplycp",
		"qdcvhkvp",
		"pdstwjaw",
		"ljdpcpa",
		"okipccubj",
		"rqlffxpo",
		"mif",
		"pejld",
		"wxme",
		"nugio",
		"vxfjsk",
		"khuiy",
		"lssl",
		"xnyve",
		"ahvfycc",
		"awcylwxu",
		"qvppnjd",
		"veuerwsxw",
		"wbxofg",
		"cmgdsxa",
		"gavgvjllu",
		"qutk",
		"oeoxutsi",
		"iylxgbqg",
		"uao",
		"ryowkg",
		"unxs",
		"grqera",
		"delg",
		"fbsl",
		"fxlco",
		"vjb",
		"ldef",
		"dir",
		"fgiea",
		"jlxbje",
		"lhagdb",
		"mfled",
		"fkddlxfm",
		"vqpcw",
		"dwdnfheod",
		"uonpkg",
		"unvfhqhwi",
		"hbmhot",
		"yhhymhlp",
		"jvlablfgp",
		"mjevyp",
		"pnhhpa",
		"ahxv",
		"nlfe",
		"llcvee",
		"qmq",
		"ijfksrasq",
		"fcidkmx",
		"otgforykh",
		"hlqmnx",
		"qhhxqqyb",
		"bivebqx",
		"fbbcqh",
		"yablms",
		"qqnbpx",
		"vleultohh",
		"fiuwf",
		"nheipldms",
		"rkho",
		"nel",
		"slkvuuy",
		"wgtsyygij",
		"lrnajcgaf",
		"tddvnmck",
		"wolrf",
		"ymvthyigu",
		"edx",
		"rbmlwxswa",
		"uktph",
		"cpdnqe",
		"uyrwvuf",
		"bkcutvu",
		"kggjpqss",
		"opbpdm",
		"vavdjh",
		"wltkjhu",
		"wgpfrnk",
		"fxluyso",
		"gsadb",
		"uwewvbqvq",
		"nnplcfnyh",
		"tfpdca",
		"hdfwwso",
		"yleeh",
		"xst",
		"bwyb",
		"pqg",
		"enhtjlpn",
	}

	labels []testLabelPair
	tree   *Tree
)

func init() {
	labels = make([]testLabelPair, len(strs))
	tree = NewTree()
	for i, s := range strs {
		lbl, n, _ := domain.MakeLabel(s)
		labels[i] = testLabelPair{
			s: lbl,
			n: n,
		}

		tree.RawInplaceInsert(lbl, n, "test")
	}
}

func BenchmarkDomainLabelTreeGet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		i := n & 1023
		s := strs[i]

		v, ok := tree.Get(s)
		if !ok {
			b.Fatalf("can't find data for %q (%q) at %d (%d)", s, labels[i], n, i)
		}

		if _, ok := v.(string); !ok {
			b.Fatalf("expected string for %q (%q) at %d (%d) but got %T (%#v)", s, labels[i], n, i, v, v)
		}
	}
}

func BenchmarkDomainLabelTreeRawGet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		i := n & 1023
		lbl := labels[i]

		v, ok := tree.RawGet(lbl.s, lbl.n)
		if !ok {
			b.Fatalf("can't find data for %q (%q) at %d (%d)", strs[i], lbl, n, i)
		}

		if _, ok := v.(string); !ok {
			b.Fatalf("expected string for %q (%q) at %d (%d) but got %T (%#v)", strs[i], lbl, n, i, v, v)
		}
	}
}
