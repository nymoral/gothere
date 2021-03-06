package utils

func HasShortName(name string) bool {
	_, flag := gamesMap[name]
	return flag
}

func GetShortNames(name1, name2 string) (string, string) {
	return gamesMap[name1], gamesMap[name2]
}

var gamesMap = map[string]string{"Afganistanas": "Af", "Airija": "Ie", "Albanija": "Al", "Alžyras": "Dz", "Argentina": "Ar", "Armėnija": "Am", "Australija": "Au", "Austrija": "At", "Azerbaidžanas": "Az", "Baltarusija": "By", "Belgija": "Be", "Bolivija": "Bo", "Bosnija ir Hercogovina": "Ba", "Brazilija": "Br", "Bulgarija": "Bg", "Čekijos Respublika": "Cz", "Čilė": "Cl", "Danija": "Dk", "Didžioji Britanija": "Gb", "Dominikos Sandrauga": "Dm", "Dominikos Respublika": "Do", "Dramblio Kaulo Krantas": "Ci", "Ekvadoras": "Ec", "Egiptas": "Eg", "Estija": "Ee", "Gajana": "Gy", "Gana": "Gh", "Grenlandija": "Gl", "Graikija": "Gr", "Gruzija": "Ge", "Hong Kongas": "Hk", "Hondūras": "Hn", "Indija": "In", "Irakas": "Iq", "Iranas": "Ir", "Islandija": "Is", "Ispanija": "Es", "Italija": "It", "Izraelis": "Il", "Japonija": "Jp", "Jugoslavija": "Y", "Jungtinė Karalystė": "Gb", "Jungtinės Amerikos Valstijos": "Us", "Kanada": "Ca", "Kamerūnas": "Cm", "Kazachstanas": "Kz", "Kinija": "Cn", "Kirgizstanas": "Kg", "Kolumbija": "Co", "Kroatija": "Hr", "Kuba": "Cu", "Kosta Rika": "Cr", "Latvija": "Lv", "Lenkija": "Pl", "Lietuva": "Lt", "Liuksemburgas": "Lu", "Makedonija": "Mk", "Marokas": "Ma", "Meksika": "Mx", "Moldavija": "Md", "Mongolija": "Mn", "Naujoji Zelandija": "Nz", "Nyderlandai": "Nl", "Nigerija": "Ng", "Norvegija": "No", "Paragvajus": "Py", "Per": "Pe", "Pietų Afrika ": "Za", "Pietų Korėja": "Kr", "Portugalija": "Pt", "Prancūzija": "Fr", "Rumunija": "Ro", "Rusija": "Ru", "Singapūras": "Sg", "Slovakija": "Sk", "Slovėnija": "Si", "Surinamas": "Sr", "Suomija": "Fi", "Šiaurės Korėja": "Kp", "Švedija": "Se", "Šveicarija": "Ch", "Tailandas": "Th", "Taivanis": "Tw", "Tadžikistanas": "Tj", "Turkija": "Tr", "Turkmėnija": "Tm", "Ukraina": "Ua", "Urugvajus": "Uy", "Uzbekija": "Uz", "Vatikanas": "Va", "Venesuela": "Ve", "Vengrija": "Hu", "Vietnamas": "Vn", "Vokietija": "De", "Zambija": "Zm"}
