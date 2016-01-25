package lang

import (
	"os"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// ByName implements sort.Interface for []Language based on
// the Description field.
type ByName []Language

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Description < a[j].Description }

// Language defines a recognized language for
type Language struct {
	ID            string   // ID
	Alias         []string // Available alias (as input for Subify)
	Description   string   // Description of the language
	OpenSubtitles string   // Open Subtitles id for this language, empty if not available
	SubDB         string   // SubDB id for this language, empty if not available
}

// Langs is the type for a list of languages
type Langs []Language

// GetLanguage get the language from an id
func (l Langs) GetLanguage(id string) (lang *Language) {
	for _, v := range Languages {
		if v.ID == id {
			return &v
		}
		for _, a := range v.Alias {
			if a == id {
				return &v
			}
		}
	}

	return nil
}

// Print prints the languages
func (l Langs) Print(all bool) {

	table := tablewriter.NewWriter(os.Stdout)

	if all {
		table.SetHeader([]string{"Language", "Id(s)", "Available ?"})
	} else {
		table.SetHeader([]string{"Language", "Available Id(s)"})
	}

	langs := Languages
	sort.Sort(ByName(langs))

	for _, l := range langs {
		if all {
			available := "No"
			if l.OpenSubtitles != "" || l.SubDB != "" {
				available = "Yes"
			}
			values := []string{
				l.Description,                             // Language
				strings.Join(append(l.Alias, l.ID), ", "), // Id(s)
				available, // Available ?
			}
			table.Append(values)
		} else if l.OpenSubtitles != "" || l.SubDB != "" {
			values := []string{
				l.Description,                             // Language
				strings.Join(append(l.Alias, l.ID), ", "), // Available Id(s)
			}
			table.Append(values)
		}
	}
	table.SetAutoWrapText(false)
	table.SetColWidth(50)
	table.Render() // Send output
}

//Languages is the list of all langauges
var Languages = Langs{
	{"aar", []string{"aa"}, "Afar, afar", "", ""},
	{"abk", []string{"ab"}, "Abkhazian", "", ""},
	{"ace", []string{}, "Achinese", "", ""},
	{"ach", []string{}, "Acoli", "", ""},
	{"ada", []string{}, "Adangme", "", ""},
	{"ady", []string{}, "Adyghé", "", ""},
	{"afa", []string{}, "Afro-Asiatic (Other)", "", ""},
	{"afh", []string{}, "Afrihili", "", ""},
	{"afr", []string{"af"}, "Afrikaans", "afr", ""},
	{"ain", []string{}, "Ainu", "", ""},
	{"aka", []string{"ak"}, "Akan", "", ""},
	{"akk", []string{}, "Akkadian", "", ""},
	{"alb", []string{"sq"}, "Albanian", "alb", ""},
	{"ale", []string{}, "Aleut", "", ""},
	{"alg", []string{}, "Algonquian languages", "", ""},
	{"alt", []string{}, "Southern Altai", "", ""},
	{"amh", []string{"am"}, "Amharic", "", ""},
	{"ang", []string{}, "English, Old", "", ""},
	{"apa", []string{}, "Apache languages", "", ""},
	{"ara", []string{"ar"}, "Arabic", "ara", ""},
	{"arc", []string{}, "Aramaic", "", ""},
	{"arg", []string{"an"}, "Aragonese", "", ""},
	{"arm", []string{"hy"}, "Armenian", "arm", ""},
	{"arn", []string{}, "Araucanian", "", ""},
	{"arp", []string{}, "Arapaho", "", ""},
	{"art", []string{}, "Artificial (Other)", "", ""},
	{"arw", []string{}, "Arawak", "", ""},
	{"asm", []string{"as"}, "Assamese", "", ""},
	{"ast", []string{}, "Asturian, Bable", "", ""},
	{"ath", []string{}, "Athapascan languages", "", ""},
	{"aus", []string{}, "Australian languages", "", ""},
	{"ava", []string{"av"}, "Avaric", "", ""},
	{"ave", []string{"ae"}, "Avestan", "", ""},
	{"awa", []string{}, "Awadhi", "", ""},
	{"aym", []string{"ay"}, "Aymara", "", ""},
	{"aze", []string{"az"}, "Azerbaijani", "", ""},
	{"bad", []string{}, "Banda", "", ""},
	{"bai", []string{}, "Bamileke languages", "", ""},
	{"bak", []string{"ba"}, "Bashkir", "", ""},
	{"bal", []string{}, "Baluchi", "", ""},
	{"bam", []string{"bm"}, "Bambara", "", ""},
	{"ban", []string{}, "Balinese", "", ""},
	{"baq", []string{"eu"}, "Basque", "", ""},
	{"bas", []string{}, "Basa", "", ""},
	{"bat", []string{}, "Baltic (Other)", "", ""},
	{"bej", []string{}, "Beja", "", ""},
	{"bel", []string{"be"}, "Belarusian", "bel", ""},
	{"bem", []string{}, "Bemba", "", ""},
	{"ben", []string{"bn"}, "Bengali", "ben", ""},
	{"ber", []string{}, "Berber (Other)", "", ""},
	{"bho", []string{}, "Bhojpuri", "", ""},
	{"bih", []string{"bh"}, "Bihari", "", ""},
	{"bik", []string{}, "Bikol", "", ""},
	{"bin", []string{}, "Bini", "", ""},
	{"bis", []string{"bi"}, "Bislama", "", ""},
	{"bla", []string{}, "Siksika", "", ""},
	{"bnt", []string{}, "Bantu (Other)", "", ""},
	{"bos", []string{"bs"}, "Bosnian", "bos", ""},
	{"bra", []string{}, "Braj", "", ""},
	{"bre", []string{"br"}, "Breton", "", ""},
	{"btk", []string{}, "Batak (Indonesia)", "", ""},
	{"bua", []string{}, "Buriat", "", ""},
	{"bug", []string{}, "Buginese", "", ""},
	{"bul", []string{"bg"}, "Bulgarian", "bul", ""},
	{"bur", []string{"my"}, "Burmese", "bur", ""},
	{"byn", []string{}, "Blin", "", ""},
	{"cad", []string{}, "Caddo", "", ""},
	{"cai", []string{}, "Central American Indian", "", ""},
	{"car", []string{}, "Carib", "", ""},
	{"cat", []string{"ca"}, "Catalan", "cat", ""},
	{"cau", []string{}, "Caucasian (Other)", "", ""},
	{"ceb", []string{}, "Cebuano", "", ""},
	{"cel", []string{}, "Celtic (Other)", "", ""},
	{"cha", []string{"ch"}, "Chamorro", "", ""},
	{"chb", []string{}, "Chibcha", "", ""},
	{"che", []string{"ce"}, "Chechen", "", ""},
	{"chg", []string{}, "Chagatai", "", ""},
	{"chi", []string{"zh"}, "Chinese (simplified)", "chi", ""},
	{"chk", []string{}, "Chuukese", "", ""},
	{"chm", []string{}, "Mari", "", ""},
	{"chn", []string{}, "Chinook jargon", "", ""},
	{"cho", []string{}, "Choctaw", "", ""},
	{"chp", []string{}, "Chipewyan", "", ""},
	{"chr", []string{}, "Cherokee", "", ""},
	{"chu", []string{"cu"}, "Church Slavic", "", ""},
	{"chv", []string{"cv"}, "Chuvash", "", ""},
	{"chy", []string{}, "Cheyenne", "", ""},
	{"cmc", []string{}, "Chamic languages", "", ""},
	{"cop", []string{}, "Coptic", "", ""},
	{"cor", []string{"kw"}, "Cornish", "", ""},
	{"cos", []string{"co"}, "Corsican", "", ""},
	{"cpe", []string{}, "Creoles and pidgins, English", "", ""},
	{"cpf", []string{}, "Creoles and pidgins, French", "", ""},
	{"cpp", []string{}, "Creoles and pidgins, Portug", "", ""},
	{"cre", []string{"cr"}, "Cree", "", ""},
	{"crh", []string{}, "Crimean Tatar", "", ""},
	{"crp", []string{}, "Creoles and pidgins (Other)", "", ""},
	{"csb", []string{}, "Kashubian", "", ""},
	{"cus", []string{}, "Cushitic (Other)", "", ""},
	{"cze", []string{"cs"}, "Czech", "cze", ""},
	{"dak", []string{}, "Dakota", "", ""},
	{"dan", []string{"da"}, "Danish", "dan", ""},
	{"dar", []string{}, "Dargwa", "", ""},
	{"day", []string{}, "Dayak", "", ""},
	{"del", []string{}, "Delaware", "", ""},
	{"den", []string{}, "Slave (Athapascan)", "", ""},
	{"dgr", []string{}, "Dogrib", "", ""},
	{"din", []string{}, "Dinka", "", ""},
	{"div", []string{"dv"}, "Divehi", "", ""},
	{"doi", []string{}, "Dogri", "", ""},
	{"dra", []string{}, "Dravidian (Other)", "", ""},
	{"dua", []string{}, "Duala", "", ""},
	{"dum", []string{}, "Dutch, Middle", "", ""},
	{"dut", []string{"nl"}, "Dutch", "dut", "nl"},
	{"dyu", []string{}, "Dyula", "", ""},
	{"dzo", []string{"dz"}, "Dzongkha", "", ""},
	{"efi", []string{}, "Efik", "", ""},
	{"egy", []string{}, "Egyptian (Ancient)", "", ""},
	{"eka", []string{}, "Ekajuk", "", ""},
	{"elx", []string{}, "Elamite", "", ""},
	{"eng", []string{"en"}, "English", "eng", "en"},
	{"enm", []string{}, "English, Middle", "", ""},
	{"epo", []string{"eo"}, "Esperanto", "epo", ""},
	{"est", []string{"et"}, "Estonian", "est", ""},
	{"ewe", []string{"ee"}, "Ewe", "", ""},
	{"ewo", []string{}, "Ewondo", "", ""},
	{"fan", []string{}, "Fang", "", ""},
	{"fao", []string{"fo"}, "Faroese", "", ""},
	{"fat", []string{}, "Fanti", "", ""},
	{"fij", []string{"fj"}, "Fijian", "", ""},
	{"fil", []string{}, "Filipino", "", ""},
	{"fin", []string{"fi"}, "Finnish", "fin", ""},
	{"fiu", []string{}, "Finno-Ugrian (Other)", "", ""},
	{"fon", []string{}, "Fon", "", ""},
	{"fre", []string{"fr"}, "French", "fre", "fr"},
	{"frm", []string{}, "French, Middle", "", ""},
	{"fro", []string{}, "French, Old", "", ""},
	{"fry", []string{"fy"}, "Frisian", "", ""},
	{"ful", []string{"ff"}, "Fulah", "", ""},
	{"fur", []string{}, "Friulian", "", ""},
	{"gaa", []string{}, "Ga", "", ""},
	{"gay", []string{}, "Gayo", "", ""},
	{"gba", []string{}, "Gbaya", "", ""},
	{"gem", []string{}, "Germanic (Other)", "", ""},
	{"geo", []string{"ka"}, "Georgian", "geo", ""},
	{"ger", []string{"de"}, "German", "ger", ""},
	{"gez", []string{}, "Geez", "", ""},
	{"gil", []string{}, "Gilbertese", "", ""},
	{"gla", []string{"gd"}, "Gaelic", "", ""},
	{"gle", []string{"ga"}, "Irish", "", ""},
	{"glg", []string{"gl"}, "Galician", "glg", ""},
	{"glv", []string{"gv"}, "Manx", "", ""},
	{"gmh", []string{}, "German, Middle High)", "", ""},
	{"goh", []string{}, "German, Old High", "", ""},
	{"gon", []string{}, "Gondi", "", ""},
	{"gor", []string{}, "Gorontalo", "", ""},
	{"got", []string{}, "Gothic", "", ""},
	{"grb", []string{}, "Grebo", "", ""},
	{"grc", []string{}, "Greek, Ancient (to 1453)", "", ""},
	{"ell", []string{"el"}, "Greek", "ell", ""},
	{"grn", []string{"gn"}, "Guarani", "", ""},
	{"guj", []string{"gu"}, "Gujarati", "", ""},
	{"gwi", []string{}, "Gwich´in", "", ""},
	{"hai", []string{}, "Haida", "", ""},
	{"hat", []string{"ht"}, "Haitian", "", ""},
	{"hau", []string{"ha"}, "Hausa", "", ""},
	{"haw", []string{}, "Hawaiian", "", ""},
	{"heb", []string{"he"}, "Hebrew", "heb", ""},
	{"her", []string{"hz"}, "Herero", "", ""},
	{"hil", []string{}, "Hiligaynon", "", ""},
	{"him", []string{}, "Himachali", "", ""},
	{"hin", []string{"hi"}, "Hindi", "hin", ""},
	{"hit", []string{}, "Hittite", "", ""},
	{"hmn", []string{}, "Hmong", "", ""},
	{"hmo", []string{"ho"}, "Hiri Motu", "", ""},
	{"hrv", []string{"hr"}, "Croatian", "hrv", ""},
	{"hun", []string{"hu"}, "Hungarian", "hun", ""},
	{"hup", []string{}, "Hupa", "", ""},
	{"iba", []string{}, "Iban", "", ""},
	{"ibo", []string{"ig"}, "Igbo", "", ""},
	{"ice", []string{"is"}, "Icelandic", "ice", ""},
	{"ido", []string{"io"}, "Ido", "", ""},
	{"iii", []string{"ii"}, "Sichuan Yi", "", ""},
	{"ijo", []string{}, "Ijo", "", ""},
	{"iku", []string{"iu"}, "Inuktitut", "", ""},
	{"ile", []string{"ie"}, "Interlingue", "", ""},
	{"ilo", []string{}, "Iloko", "", ""},
	{"ina", []string{"ia"}, "Interlingua", "", ""},
	{"inc", []string{}, "Indic (Other)", "", ""},
	{"ind", []string{"id"}, "Indonesian", "ind", ""},
	{"ine", []string{}, "Indo-European (Other)", "", ""},
	{"inh", []string{}, "Ingush", "", ""},
	{"ipk", []string{"ik"}, "Inupiaq", "", ""},
	{"ira", []string{}, "Iranian (Other)", "", ""},
	{"iro", []string{}, "Iroquoian languages", "", ""},
	{"ita", []string{"it"}, "Italian", "ita", "it"},
	{"jav", []string{"jv"}, "Javanese", "", ""},
	{"jpn", []string{"ja"}, "Japanese", "jpn", ""},
	{"jpr", []string{}, "Judeo-Persian", "", ""},
	{"jrb", []string{}, "Judeo-Arabic", "", ""},
	{"kaa", []string{}, "Kara-Kalpak", "", ""},
	{"kab", []string{}, "Kabyle", "", ""},
	{"kac", []string{}, "Kachin", "", ""},
	{"kal", []string{"kl"}, "Kalaallisut", "", ""},
	{"kam", []string{}, "Kamba", "", ""},
	{"kan", []string{"kn"}, "Kannada", "", ""},
	{"kar", []string{}, "Karen", "", ""},
	{"kas", []string{"ks"}, "Kashmiri", "", ""},
	{"kau", []string{"kr"}, "Kanuri", "", ""},
	{"kaw", []string{}, "Kawi", "", ""},
	{"kaz", []string{"kk"}, "Kazakh", "kaz", ""},
	{"kbd", []string{}, "Kabardian", "", ""},
	{"kha", []string{}, "Khasi", "", ""},
	{"khi", []string{}, "Khoisan (Other)", "", ""},
	{"khm", []string{"km"}, "Khmer", "", ""},
	{"kho", []string{}, "Khotanese", "", ""},
	{"kik", []string{"ki"}, "Kikuyu", "", ""},
	{"kin", []string{"rw"}, "Kinyarwanda", "", ""},
	{"kir", []string{"ky"}, "Kirghiz", "", ""},
	{"kmb", []string{}, "Kimbundu", "", ""},
	{"kok", []string{}, "Konkani", "", ""},
	{"kom", []string{"kv"}, "Komi", "", ""},
	{"kon", []string{"kg"}, "Kongo", "", ""},
	{"kor", []string{"ko"}, "Korean", "kor", ""},
	{"kos", []string{}, "Kosraean", "", ""},
	{"kpe", []string{}, "Kpelle", "", ""},
	{"krc", []string{}, "Karachay-Balkar", "", ""},
	{"kro", []string{}, "Kru", "", ""},
	{"kru", []string{}, "Kurukh", "", ""},
	{"kua", []string{"kj"}, "Kuanyama", "", ""},
	{"kum", []string{}, "Kumyk", "", ""},
	{"kur", []string{"ku"}, "Kurdish", "", ""},
	{"kut", []string{}, "Kutenai", "", ""},
	{"lad", []string{}, "Ladino", "", ""},
	{"lah", []string{}, "Lahnda", "", ""},
	{"lam", []string{}, "Lamba", "", ""},
	{"lao", []string{"lo"}, "Lao", "", ""},
	{"lat", []string{"la"}, "Latin", "", ""},
	{"lav", []string{"lv"}, "Latvian", "lav", ""},
	{"lez", []string{}, "Lezghian", "", ""},
	{"lim", []string{"li"}, "Limburgan", "", ""},
	{"lin", []string{"ln"}, "Lingala", "", ""},
	{"lit", []string{"lt"}, "Lithuanian", "lit", ""},
	{"lol", []string{}, "Mongo", "", ""},
	{"loz", []string{}, "Lozi", "", ""},
	{"ltz", []string{"lb"}, "Luxembourgish", "ltz", ""},
	{"lua", []string{}, "Luba-Lulua", "", ""},
	{"lub", []string{"lu"}, "Luba-Katanga", "", ""},
	{"lug", []string{"lg"}, "Ganda", "", ""},
	{"lui", []string{}, "Luiseno", "", ""},
	{"lun", []string{}, "Lunda", "", ""},
	{"luo", []string{}, "Luo (Kenya and Tanzania)", "", ""},
	{"lus", []string{}, "Lushai", "", ""},
	{"mac", []string{"mk"}, "Macedonian", "mac", ""},
	{"mad", []string{}, "Madurese", "", ""},
	{"mag", []string{}, "Magahi", "", ""},
	{"mah", []string{"mh"}, "Marshallese", "", ""},
	{"mai", []string{}, "Maithili", "", ""},
	{"mak", []string{}, "Makasar", "", ""},
	{"mal", []string{"ml"}, "Malayalam", "", ""},
	{"man", []string{}, "Mandingo", "", ""},
	{"mao", []string{"mi"}, "Maori", "", ""},
	{"map", []string{}, "Austronesian (Other)", "", ""},
	{"mar", []string{"mr"}, "Marathi", "", ""},
	{"mas", []string{}, "Masai", "", ""},
	{"may", []string{"ms"}, "Malay", "", ""},
	{"mdf", []string{}, "Moksha", "", ""},
	{"mdr", []string{}, "Mandar", "", ""},
	{"men", []string{}, "Mende", "", ""},
	{"mga", []string{}, "Irish, Middle", "", ""},
	{"mic", []string{}, "Mi'kmaq", "", ""},
	{"min", []string{}, "Minangkabau", "", ""},
	{"mis", []string{}, "Miscellaneous languages", "", ""},
	{"mkh", []string{}, "Mon-Khmer (Other)", "", ""},
	{"mlg", []string{"mg"}, "Malagasy", "", ""},
	{"mlt", []string{"mt"}, "Maltese", "", ""},
	{"mnc", []string{}, "Manchu", "", ""},
	{"mni", []string{"ma"}, "Manipuri", "mni", ""},
	{"mno", []string{}, "Manobo languages", "", ""},
	{"moh", []string{}, "Mohawk", "", ""},
	{"mol", []string{"mo"}, "Moldavian", "", ""},
	{"mon", []string{"mn"}, "Mongolian", "mon", ""},
	{"mos", []string{}, "Mossi", "", ""},
	{"mwl", []string{}, "Mirandese", "", ""},
	{"mul", []string{}, "Multiple languages", "", ""},
	{"mun", []string{}, "Munda languages", "", ""},
	{"mus", []string{}, "Creek", "", ""},
	{"mwr", []string{}, "Marwari", "", ""},
	{"myn", []string{}, "Mayan languages", "", ""},
	{"myv", []string{}, "Erzya", "", ""},
	{"nah", []string{}, "Nahuatl", "", ""},
	{"nai", []string{}, "North American Indian", "", ""},
	{"nap", []string{}, "Neapolitan", "", ""},
	{"nau", []string{"na"}, "Nauru", "", ""},
	{"nav", []string{"nv"}, "Navajo", "", ""},
	{"nbl", []string{"nr"}, "Ndebele, South", "", ""},
	{"nde", []string{"nd"}, "Ndebele, North", "", ""},
	{"ndo", []string{"ng"}, "Ndonga", "", ""},
	{"nds", []string{}, "Low German", "", ""},
	{"nep", []string{"ne"}, "Nepali", "", ""},
	{"new", []string{}, "Nepal Bhasa", "", ""},
	{"nia", []string{}, "Nias", "", ""},
	{"nic", []string{}, "Niger-Kordofanian (Other)", "", ""},
	{"niu", []string{}, "Niuean", "", ""},
	{"nno", []string{"nn"}, "Norwegian Nynorsk", "", ""},
	{"nob", []string{"nb"}, "Norwegian Bokmal", "", ""},
	{"nog", []string{}, "Nogai", "", ""},
	{"non", []string{}, "Norse, Old", "", ""},
	{"nor", []string{"no"}, "Norwegian", "nor", ""},
	{"nso", []string{}, "Northern Sotho", "", ""},
	{"nub", []string{}, "Nubian languages", "", ""},
	{"nwc", []string{}, "Classical Newari", "", ""},
	{"nya", []string{"ny"}, "Chichewa", "", ""},
	{"nym", []string{}, "Nyamwezi", "", ""},
	{"nyn", []string{}, "Nyankole", "", ""},
	{"nyo", []string{}, "Nyoro", "", ""},
	{"nzi", []string{}, "Nzima", "", ""},
	{"oci", []string{"oc"}, "Occitan", "", ""},
	{"oji", []string{"oj"}, "Ojibwa", "", ""},
	{"ori", []string{"or"}, "Oriya", "", ""},
	{"orm", []string{"om"}, "Oromo", "", ""},
	{"osa", []string{}, "Osage", "", ""},
	{"oss", []string{"os"}, "Ossetian", "", ""},
	{"ota", []string{}, "Turkish, Ottoman", "", ""},
	{"oto", []string{}, "Otomian languages", "", ""},
	{"paa", []string{}, "Papuan (Other)", "", ""},
	{"pag", []string{}, "Pangasinan", "", ""},
	{"pal", []string{}, "Pahlavi", "", ""},
	{"pam", []string{}, "Pampanga", "", ""},
	{"pan", []string{"pa"}, "Panjabi", "", ""},
	{"pap", []string{}, "Papiamento", "", ""},
	{"pau", []string{}, "Palauan", "", ""},
	{"peo", []string{}, "Persian, Old", "", ""},
	{"per", []string{"fa"}, "Persian", "per", ""},
	{"phi", []string{}, "Philippine (Other)", "", ""},
	{"phn", []string{}, "Phoenician", "", ""},
	{"pli", []string{"pi"}, "Pali", "", ""},
	{"pol", []string{"pl"}, "Polish", "pol", "pl"},
	{"pon", []string{}, "Pohnpeian", "", ""},
	{"por", []string{"pt"}, "Portuguese", "por", ""},
	{"pra", []string{}, "Prakrit languages", "", ""},
	{"pro", []string{}, "Provençal, Old", "", ""},
	{"pus", []string{"ps"}, "Pushto", "", ""},
	{"que", []string{"qu"}, "Quechua", "", ""},
	{"raj", []string{}, "Rajasthani", "", ""},
	{"rap", []string{}, "Rapanui", "", ""},
	{"rar", []string{}, "Rarotongan", "", ""},
	{"roa", []string{}, "Romance (Other)", "", ""},
	{"roh", []string{"rm"}, "Raeto-Romance", "", ""},
	{"rom", []string{}, "Romany", "", ""},
	{"run", []string{"rn"}, "Rundi", "", ""},
	{"rup", []string{}, "Aromanian", "", ""},
	{"rus", []string{"ru"}, "Russian", "rus", ""},
	{"sad", []string{}, "Sandawe", "", ""},
	{"sag", []string{"sg"}, "Sango", "", ""},
	{"sah", []string{}, "Yakut", "", ""},
	{"sai", []string{}, "South American Indian (Other)", "", ""},
	{"sal", []string{}, "Salishan languages", "", ""},
	{"sam", []string{}, "Samaritan Aramaic", "", ""},
	{"san", []string{"sa"}, "Sanskrit", "", ""},
	{"sas", []string{}, "Sasak", "", ""},
	{"sat", []string{}, "Santali", "", ""},
	{"scc", []string{"sr"}, "Serbian", "scc", ""},
	{"scn", []string{}, "Sicilian", "", ""},
	{"sco", []string{}, "Scots", "", ""},
	{"sel", []string{}, "Selkup", "", ""},
	{"sem", []string{}, "Semitic (Other)", "", ""},
	{"sga", []string{}, "Irish, Old", "", ""},
	{"sgn", []string{}, "Sign Languages", "", ""},
	{"shn", []string{}, "Shan", "", ""},
	{"sid", []string{}, "Sidamo", "", ""},
	{"sin", []string{"si"}, "Sinhalese", "sin", ""},
	{"sio", []string{}, "Siouan languages", "", ""},
	{"sit", []string{}, "Sino-Tibetan (Other)", "", ""},
	{"sla", []string{}, "Slavic (Other)", "", ""},
	{"slo", []string{"sk"}, "Slovak", "slo", ""},
	{"slv", []string{"sl"}, "Slovenian", "slv", ""},
	{"sma", []string{}, "Southern Sami", "", ""},
	{"sme", []string{"se"}, "Northern Sami", "", ""},
	{"smi", []string{}, "Sami languages (Other)", "", ""},
	{"smj", []string{}, "Lule Sami", "", ""},
	{"smn", []string{}, "Inari Sami", "", ""},
	{"smo", []string{"sm"}, "Samoan", "", ""},
	{"sms", []string{}, "Skolt Sami", "", ""},
	{"sna", []string{"sn"}, "Shona", "", ""},
	{"snd", []string{"sd"}, "Sindhi", "", ""},
	{"snk", []string{}, "Soninke", "", ""},
	{"sog", []string{}, "Sogdian", "", ""},
	{"som", []string{"so"}, "Somali", "", ""},
	{"son", []string{}, "Songhai", "", ""},
	{"sot", []string{"st"}, "Sotho, Southern", "", ""},
	{"spa", []string{"es"}, "Spanish", "spa", "es"},
	{"srd", []string{"sc"}, "Sardinian", "", ""},
	{"srr", []string{}, "Serer", "", ""},
	{"ssa", []string{}, "Nilo-Saharan (Other)", "", ""},
	{"ssw", []string{"ss"}, "Swati", "", ""},
	{"suk", []string{}, "Sukuma", "", ""},
	{"sun", []string{"su"}, "Sundanese", "", ""},
	{"sus", []string{}, "Susu", "", ""},
	{"sux", []string{}, "Sumerian", "", ""},
	{"swa", []string{"sw"}, "Swahili", "swa", ""},
	{"swe", []string{"sv"}, "Swedish", "swe", "sv"},
	{"syr", []string{"sy"}, "Syriac", "syr", ""},
	{"tah", []string{"ty"}, "Tahitian", "", ""},
	{"tai", []string{}, "Tai (Other)", "", ""},
	{"tam", []string{"ta"}, "Tamil", "tam", ""},
	{"tat", []string{"tt"}, "Tatar", "", ""},
	{"tel", []string{"te"}, "Telugu", "", ""},
	{"tem", []string{}, "Timne", "", ""},
	{"ter", []string{}, "Tereno", "", ""},
	{"tet", []string{}, "Tetum", "", ""},
	{"tgk", []string{"tg"}, "Tajik", "", ""},
	{"tgl", []string{"tl"}, "Tagalog", "", ""},
	{"tha", []string{"th"}, "Thai", "tha", ""},
	{"tib", []string{"bo"}, "Tibetan", "", ""},
	{"tig", []string{}, "Tigre", "", ""},
	{"tir", []string{"ti"}, "Tigrinya", "", ""},
	{"tiv", []string{}, "Tiv", "", ""},
	{"tkl", []string{}, "Tokelau", "", ""},
	{"tlh", []string{}, "Klingon", "", ""},
	{"tli", []string{}, "Tlingit", "", ""},
	{"tmh", []string{}, "Tamashek", "", ""},
	{"tog", []string{}, "Tonga (Nyasa)", "", ""},
	{"ton", []string{"to"}, "Tonga (Tonga Islands)", "", ""},
	{"tpi", []string{}, "Tok Pisin", "", ""},
	{"tsi", []string{}, "Tsimshian", "", ""},
	{"tsn", []string{"tn"}, "Tswana", "", ""},
	{"tso", []string{"ts"}, "Tsonga", "", ""},
	{"tuk", []string{"tk"}, "Turkmen", "", ""},
	{"tum", []string{}, "Tumbuka", "", ""},
	{"tup", []string{}, "Tupi languages", "", ""},
	{"tur", []string{"tr"}, "Turkish", "tur", "tr"},
	{"tut", []string{}, "Altaic (Other)", "", ""},
	{"tvl", []string{}, "Tuvalu", "", ""},
	{"twi", []string{"tw"}, "Twi", "", ""},
	{"tyv", []string{}, "Tuvinian", "", ""},
	{"udm", []string{}, "Udmurt", "", ""},
	{"uga", []string{}, "Ugaritic", "", ""},
	{"uig", []string{"ug"}, "Uighur", "", ""},
	{"ukr", []string{"uk"}, "Ukrainian", "ukr", ""},
	{"umb", []string{}, "Umbundu", "", ""},
	{"und", []string{}, "Undetermined", "", ""},
	{"urd", []string{"ur"}, "Urdu", "", ""},
	{"uzb", []string{"uz"}, "Uzbek", "", ""},
	{"vai", []string{}, "Vai", "", ""},
	{"ven", []string{"ve"}, "Venda", "", ""},
	{"vie", []string{"vi"}, "Vietnamese", "vie", ""},
	{"vol", []string{"vo"}, "Volapük", "", ""},
	{"vot", []string{}, "Votic", "", ""},
	{"wak", []string{}, "Wakashan languages", "", ""},
	{"wal", []string{}, "Walamo", "", ""},
	{"war", []string{}, "Waray", "", ""},
	{"was", []string{}, "Washo", "", ""},
	{"wel", []string{"cy"}, "Welsh", "", ""},
	{"wen", []string{}, "Sorbian languages", "", ""},
	{"wln", []string{"wa"}, "Walloon", "", ""},
	{"wol", []string{"wo"}, "Wolof", "", ""},
	{"xal", []string{}, "Kalmyk", "", ""},
	{"xho", []string{"xh"}, "Xhosa", "", ""},
	{"yao", []string{}, "Yao", "", ""},
	{"yap", []string{}, "Yapese", "", ""},
	{"yid", []string{"yi"}, "Yiddish", "", ""},
	{"yor", []string{"yo"}, "Yoruba", "", ""},
	{"ypk", []string{}, "Yupik languages", "", ""},
	{"zap", []string{}, "Zapotec", "", ""},
	{"zen", []string{}, "Zenaga", "", ""},
	{"zha", []string{"za"}, "Zhuang", "", ""},
	{"znd", []string{}, "Zande", "", ""},
	{"zul", []string{"zu"}, "Zulu", "", ""},
	{"zun", []string{}, "Zuni", "", ""},
	{"rum", []string{"ro"}, "Romanian", "rum", "ro"},
	{"pob", []string{"pb"}, "Portuguese (BR)", "pob", "pt"},
	{"mne", []string{"me"}, "Montenegrin", "mne", ""},
	{"zht", []string{"zt"}, "Chinese (traditional)", "zht", ""},
	{"zhe", []string{"ze"}, "Chinese bilingual", "zhe", ""},
}
