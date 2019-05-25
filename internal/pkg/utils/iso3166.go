package utils

import (
	"github.com/pariz/gountries"
	"strconv"
	"strings"
)

/*
 * id to iso3166 mappings
 * http://www.dmr-marc.net/FAQ/GSM-ID.html
 */
var idToCountryCode = map[int]string{
	1023: "ca",
	110:  "us",
	111:  "us",
	112:  "us",
	113:  "us",
	114:  "us",
	115:  "us",
	202:  "gr",
	204:  "nl",
	206:  "be",
	208:  "fr",
	212:  "mc",
	213:  "ad",
	214:  "es",
	216:  "hu",
	218:  "ba",
	219:  "hr",
	220:  "cs",
	222:  "it",
	226:  "ro",
	228:  "ch",
	230:  "cz",
	231:  "sk",
	232:  "at",
	234:  "gb",
	235:  "gb",
	238:  "dk",
	240:  "se",
	242:  "no",
	244:  "fi",
	246:  "lt",
	247:  "lv",
	248:  "ee",
	250:  "ru",
	255:  "ua",
	257:  "by",
	259:  "md",
	260:  "pl",
	262:  "de",
	263:  "de",
	266:  "gi",
	268:  "pt",
	270:  "lu",
	272:  "ie",
	274:  "is",
	278:  "mt",
	280:  "cy",
	284:  "bg",
	286:  "tr",
	288:  "fo",
	290:  "gl",
	292:  "sm",
	293:  "si",
	294:  "mk",
	295:  "li",
	297:  "cs",
	302:  "ca",
	310:  "us",
	311:  "us",
	312:  "us",
	313:  "us",
	314:  "us",
	315:  "us",
	316:  "us",
	330:  "pr",
	332:  "vg",
	334:  "mx",
	340:  "mq",
	342:  "bb",
	352:  "gd",
	358:  "lc",
	360:  "vc",
	362:  "an",
	364:  "bs",
	366:  "dm",
	370:  "do",
	372:  "ht",
	374:  "tt",
	376:  "vi",
	401:  "kz",
	404:  "in",
	415:  "lb",
	418:  "iq",
	419:  "kw",
	420:  "sa",
	421:  "ye",
	422:  "om",
	425:  "il",
	427:  "qa",
	431:  "ae",
	432:  "ir",
	440:  "jp",
	441:  "jp",
	450:  "kr",
	452:  "vn",
	454:  "hk",
	455:  "mo",
	457:  "la",
	460:  "cn",
	466:  "tw",
	470:  "bd",
	502:  "my",
	505:  "au",
	510:  "id",
	515:  "ph",
	520:  "th",
	525:  "sg",
	530:  "nz",
	5351: "gu",
	5371: "pg",
	546:  "nc",
	603:  "dz",
	604:  "ma",
	617:  "mu",
	620:  "gh",
	621:  "ng",
	638:  "dj",
	639:  "ke",
	643:  "mz",
	645:  "zm",
	647:  "re",
	648:  "zw",
	649:  "na",
	653:  "sz",
	655:  "za",
	702:  "bz",
	704:  "gt",
	708:  "hn",
	710:  "ni",
	712:  "cr",
	714:  "pa",
	716:  "pe",
	722:  "ar",
	724:  "br",
	730:  "cl",
	732:  "co",
	734:  "ve",
	738:  "gy",
	740:  "ec",
	744:  "py",
	748:  "uy",
	910:  "de",
	950:  "ru",
}

func GetCountryCodeForID(id int) string {
	s := strconv.Itoa(id)

	var longestPrefixId int = 0
	var longestLength int = 0

	for k, _ := range idToCountryCode {
		prefix := strconv.Itoa(k)
		if strings.HasPrefix(s, prefix) && len(prefix) > longestLength {
			longestPrefixId = k
			longestLength = len(prefix)
		}
	}

	return idToCountryCode[longestPrefixId]
}

func GetCountryNameForID(id int) string {
	iso := GetCountryCodeForID(id)
	if iso == "" {
		return "Unknown"
	}

	query := gountries.New()
	country, err := query.FindCountryByAlpha(iso)
	if err != nil {
		return "Unknown"
	}

	return country.Name.Common
}
