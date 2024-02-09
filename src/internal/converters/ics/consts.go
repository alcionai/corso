package ics

var (
	// Map from Graph API recurrence type to iCal recurrence type
	GraphToICalIndex = map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
		"fourth": 4,
		"last":   -1,
	}

	// Map from Graph API day of week representation to iCal day of week representation
	GraphToICalDOW = map[string]string{
		"sunday":    "SU",
		"monday":    "MO",
		"tuesday":   "TU",
		"wednesday": "WE",
		"thursday":  "TH",
		"friday":    "FR",
		"saturday":  "SA",
	}
)

// Map from Window time zone to TZ database time zone
// https://github.com/closeio/sync-engine/blob/1ce0e1ad0104a2ab2479da09b073c86f4feee5f9/inbox/events/timezones.py#L6
var GraphTimeZoneToTZ = map[string]string{
	"AUS Central Standard Time":       "Australia/Darwin",
	"AUS Eastern Standard Time":       "Australia/Sydney",
	"Afghanistan Standard Time":       "Asia/Kabul",
	"Alaskan Standard Time":           "America/Anchorage",
	"Aleutian Standard Time":          "America/Adak",
	"Altai Standard Time":             "Asia/Barnaul",
	"Arab Standard Time":              "Asia/Riyadh",
	"Arabian Standard Time":           "Asia/Dubai",
	"Arabic Standard Time":            "Asia/Baghdad",
	"Argentina Standard Time":         "America/Buenos_Aires",
	"Astrakhan Standard Time":         "Europe/Astrakhan",
	"Atlantic Standard Time":          "America/Halifax",
	"Aus Central W. Standard Time":    "Australia/Eucla",
	"Azerbaijan Standard Time":        "Asia/Baku",
	"Azores Standard Time":            "Atlantic/Azores",
	"Bahia Standard Time":             "America/Bahia",
	"Bangladesh Standard Time":        "Asia/Dhaka",
	"Belarus Standard Time":           "Europe/Minsk",
	"Bougainville Standard Time":      "Pacific/Bougainville",
	"Canada Central Standard Time":    "America/Regina",
	"Cape Verde Standard Time":        "Atlantic/Cape_Verde",
	"Caucasus Standard Time":          "Asia/Yerevan",
	"Cen. Australia Standard Time":    "Australia/Adelaide",
	"Central America Standard Time":   "America/Guatemala",
	"Central Asia Standard Time":      "Asia/Almaty",
	"Central Brazilian Standard Time": "America/Cuiaba",
	"Central Europe Standard Time":    "Europe/Budapest",
	"Central European Standard Time":  "Europe/Warsaw",
	"Central Pacific Standard Time":   "Pacific/Guadalcanal",
	"Central Standard Time":           "America/Chicago",
	"Central Standard Time (Mexico)":  "America/Mexico_City",
	"Chatham Islands Standard Time":   "Pacific/Chatham",
	"China Standard Time":             "Asia/Shanghai",
	"Cuba Standard Time":              "America/Havana",
	"Dateline Standard Time":          "Etc/GMT+12",
	"E. Africa Standard Time":         "Africa/Nairobi",
	"E. Australia Standard Time":      "Australia/Brisbane",
	"E. Europe Standard Time":         "Europe/Chisinau",
	"E. South America Standard Time":  "America/Sao_Paulo",
	"Easter Island Standard Time":     "Pacific/Easter",
	"Eastern Standard Time":           "America/New_York",
	"Eastern Standard Time (Mexico)":  "America/Cancun",
	"Egypt Standard Time":             "Africa/Cairo",
	"Ekaterinburg Standard Time":      "Asia/Yekaterinburg",
	"FLE Standard Time":               "Europe/Kiev",
	"Fiji Standard Time":              "Pacific/Fiji",
	"GMT Standard Time":               "Europe/London",
	"GTB Standard Time":               "Europe/Bucharest",
	"Georgian Standard Time":          "Asia/Tbilisi",
	"Greenland Standard Time":         "America/Godthab",
	"Greenwich Standard Time":         "Atlantic/Reykjavik",
	"Haiti Standard Time":             "America/Port-au-Prince",
	"Hawaiian Standard Time":          "Pacific/Honolulu",
	"India Standard Time":             "Asia/Calcutta",
	"Iran Standard Time":              "Asia/Tehran",
	"Israel Standard Time":            "Asia/Jerusalem",
	"Jordan Standard Time":            "Asia/Amman",
	"Kaliningrad Standard Time":       "Europe/Kaliningrad",
	"Korea Standard Time":             "Asia/Seoul",
	"Libya Standard Time":             "Africa/Tripoli",
	"Line Islands Standard Time":      "Pacific/Kiritimati",
	"Lord Howe Standard Time":         "Australia/Lord_Howe",
	"Magadan Standard Time":           "Asia/Magadan",
	"Magallanes Standard Time":        "America/Punta_Arenas",
	"Marquesas Standard Time":         "Pacific/Marquesas",
	"Mauritius Standard Time":         "Indian/Mauritius",
	"Middle East Standard Time":       "Asia/Beirut",
	"Montevideo Standard Time":        "America/Montevideo",
	"Morocco Standard Time":           "Africa/Casablanca",
	"Mountain Standard Time":          "America/Denver",
	"Mountain Standard Time (Mexico)": "America/Chihuahua",
	"Myanmar Standard Time":           "Asia/Rangoon",
	"N. Central Asia Standard Time":   "Asia/Novosibirsk",
	"Namibia Standard Time":           "Africa/Windhoek",
	"Nepal Standard Time":             "Asia/Katmandu",
	"New Zealand Standard Time":       "Pacific/Auckland",
	"Newfoundland Standard Time":      "America/St_Johns",
	"Norfolk Standard Time":           "Pacific/Norfolk",
	"North Asia East Standard Time":   "Asia/Irkutsk",
	"North Asia Standard Time":        "Asia/Krasnoyarsk",
	"North Korea Standard Time":       "Asia/Pyongyang",
	"Omsk Standard Time":              "Asia/Omsk",
	"Pacific SA Standard Time":        "America/Santiago",
	"Pacific Standard Time":           "America/Los_Angeles",
	"Pacific Standard Time (Mexico)":  "America/Tijuana",
	"Pakistan Standard Time":          "Asia/Karachi",
	"Paraguay Standard Time":          "America/Asuncion",
	"Qyzylorda Standard Time":         "Asia/Qyzylorda",
	"Romance Standard Time":           "Europe/Paris",
	"Russia Time Zone 10":             "Asia/Srednekolymsk",
	"Russia Time Zone 11":             "Asia/Kamchatka",
	"Russia Time Zone 3":              "Europe/Samara",
	"Russian Standard Time":           "Europe/Moscow",
	"SA Eastern Standard Time":        "America/Cayenne",
	"SA Pacific Standard Time":        "America/Bogota",
	"SA Western Standard Time":        "America/La_Paz",
	"SE Asia Standard Time":           "Asia/Bangkok",
	"Saint Pierre Standard Time":      "America/Miquelon",
	"Sakhalin Standard Time":          "Asia/Sakhalin",
	"Samoa Standard Time":             "Pacific/Apia",
	"Sao Tome Standard Time":          "Africa/Sao_Tome",
	"Saratov Standard Time":           "Europe/Saratov",
	"Singapore Standard Time":         "Asia/Singapore",
	"South Africa Standard Time":      "Africa/Johannesburg",
	"South Sudan Standard Time":       "Africa/Juba",
	"Sri Lanka Standard Time":         "Asia/Colombo",
	"Sudan Standard Time":             "Africa/Khartoum",
	"Syria Standard Time":             "Asia/Damascus",
	"Taipei Standard Time":            "Asia/Taipei",
	"Tasmania Standard Time":          "Australia/Hobart",
	"Tocantins Standard Time":         "America/Araguaina",
	"Tokyo Standard Time":             "Asia/Tokyo",
	"Tomsk Standard Time":             "Asia/Tomsk",
	"Tonga Standard Time":             "Pacific/Tongatapu",
	"Transbaikal Standard Time":       "Asia/Chita",
	"Turkey Standard Time":            "Europe/Istanbul",
	"Turks And Caicos Standard Time":  "America/Grand_Turk",
	"US Eastern Standard Time":        "America/Indianapolis",
	"US Mountain Standard Time":       "America/Phoenix",
	"UTC":                             "Etc/UTC",
	"UTC+12":                          "Etc/GMT-12",
	"UTC+13":                          "Etc/GMT-13",
	"UTC-02":                          "Etc/GMT+2",
	"UTC-08":                          "Etc/GMT+8",
	"UTC-09":                          "Etc/GMT+9",
	"UTC-11":                          "Etc/GMT+11",
	"Ulaanbaatar Standard Time":       "Asia/Ulaanbaatar",
	"Venezuela Standard Time":         "America/Caracas",
	"Vladivostok Standard Time":       "Asia/Vladivostok",
	"Volgograd Standard Time":         "Europe/Volgograd",
	"W. Australia Standard Time":      "Australia/Perth",
	"W. Central Africa Standard Time": "Africa/Lagos",
	"W. Europe Standard Time":         "Europe/Berlin",
	"W. Mongolia Standard Time":       "Asia/Hovd",
	"West Asia Standard Time":         "Asia/Tashkent",
	"West Bank Standard Time":         "Asia/Hebron",
	"West Pacific Standard Time":      "Pacific/Port_Moresby",
	"Yakutsk Standard Time":           "Asia/Yakutsk",
	"Yukon Standard Time":             "America/Whitehorse",
	"tzone://Microsoft/Utc":           "Etc/UTC",
}

// Map from alternatives to the canonical time zone name
// There mapping are currently generated by manually going on the
// values in the GraphTimeZoneToTZ which is not available in the tzdb
var CanonicalTimeZoneMap = map[string]string{
	"Africa/Asmara":        "Africa/Asmera",
	"Asia/Calcutta":        "Asia/Kolkata",
	"Asia/Rangoon":         "Asia/Yangon",
	"Asia/Saigon":          "Asia/Ho_Chi_Minh",
	"Europe/Kiev":          "Europe/Kyiv",
	"Europe/Warsaw":        "Europe/Warszawa",
	"America/Buenos_Aires": "America/Argentina/Buenos_Aires",
	"America/Godthab":      "America/Nuuk",
	// NOTE: "Atlantic/Raykjavik" missing in tzdb but is in MS list

	"Etc/UTC": "UTC", // simplifying the time zone name
}
