package worker

import "time"

type Currency struct {
	Code  string  `json:"code,omitempty"`
	Value float64 `json:"value,omitempty"`
}

// this is autogenerated by https://mholt.github.io/json-to-go/
type CurrencyLatest struct {
	Meta struct {
		LastUpdatedAt time.Time `json:"last_updated_at,omitempty"`
	} `json:"meta,omitempty"`
	Data struct {
		Ada struct {
			Currency
		} `json:"ADA,omitempty"`
		Aed struct {
			Currency
		} `json:"AED,omitempty"`
		Afn struct {
			Currency
		} `json:"AFN,omitempty"`
		All struct {
			Currency
		} `json:"ALL,omitempty"`
		Amd struct {
			Currency
		} `json:"AMD,omitempty"`
		Ang struct {
			Currency
		} `json:"ANG,omitempty"`
		Aoa struct {
			Currency
		} `json:"AOA,omitempty"`
		Ars struct {
			Currency
		} `json:"ARS,omitempty"`
		Aud struct {
			Currency
		} `json:"AUD,omitempty"`
		Avax struct {
			Currency
		} `json:"AVAX,omitempty"`
		Awg struct {
			Currency
		} `json:"AWG,omitempty"`
		Azn struct {
			Currency
		} `json:"AZN,omitempty"`
		Bam struct {
			Currency
		} `json:"BAM,omitempty"`
		Bbd struct {
			Currency
		} `json:"BBD,omitempty"`
		Bdt struct {
			Currency
		} `json:"BDT,omitempty"`
		Bgn struct {
			Currency
		} `json:"BGN,omitempty"`
		Bhd struct {
			Currency
		} `json:"BHD,omitempty"`
		Bif struct {
			Currency
		} `json:"BIF,omitempty"`
		Bmd struct {
			Currency
		} `json:"BMD,omitempty"`
		Bnb struct {
			Currency
		} `json:"BNB,omitempty"`
		Bnd struct {
			Currency
		} `json:"BND,omitempty"`
		Bob struct {
			Currency
		} `json:"BOB,omitempty"`
		Brl struct {
			Currency
		} `json:"BRL,omitempty"`
		Bsd struct {
			Currency
		} `json:"BSD,omitempty"`
		Btc struct {
			Currency
		} `json:"BTC,omitempty"`
		Btn struct {
			Currency
		} `json:"BTN,omitempty"`
		Bwp struct {
			Currency
		} `json:"BWP,omitempty"`
		Byn struct {
			Currency
		} `json:"BYN,omitempty"`
		Byr struct {
			Currency
		} `json:"BYR,omitempty"`
		Bzd struct {
			Currency
		} `json:"BZD,omitempty"`
		Cad struct {
			Currency
		} `json:"CAD,omitempty"`
		Cdf struct {
			Currency
		} `json:"CDF,omitempty"`
		Chf struct {
			Currency
		} `json:"CHF,omitempty"`
		Clf struct {
			Currency
		} `json:"CLF,omitempty"`
		Clp struct {
			Currency
		} `json:"CLP,omitempty"`
		Cny struct {
			Currency
		} `json:"CNY,omitempty"`
		Cop struct {
			Currency
		} `json:"COP,omitempty"`
		Crc struct {
			Currency
		} `json:"CRC,omitempty"`
		Cuc struct {
			Currency
		} `json:"CUC,omitempty"`
		Cup struct {
			Currency
		} `json:"CUP,omitempty"`
		Cve struct {
			Currency
		} `json:"CVE,omitempty"`
		Czk struct {
			Currency
		} `json:"CZK,omitempty"`
		Djf struct {
			Currency
		} `json:"DJF,omitempty"`
		Dkk struct {
			Currency
		} `json:"DKK,omitempty"`
		Dop struct {
			Currency
		} `json:"DOP,omitempty"`
		Dot struct {
			Currency
		} `json:"DOT,omitempty"`
		Dzd struct {
			Currency
		} `json:"DZD,omitempty"`
		Egp struct {
			Currency
		} `json:"EGP,omitempty"`
		Ern struct {
			Currency
		} `json:"ERN,omitempty"`
		Etb struct {
			Currency
		} `json:"ETB,omitempty"`
		Eth struct {
			Currency
		} `json:"ETH,omitempty"`
		Eur struct {
			Currency
		} `json:"EUR,omitempty"`
		Fjd struct {
			Currency
		} `json:"FJD,omitempty"`
		Fkp struct {
			Currency
		} `json:"FKP,omitempty"`
		Gbp struct {
			Currency
		} `json:"GBP,omitempty"`
		Gel struct {
			Currency
		} `json:"GEL,omitempty"`
		Ggp struct {
			Currency
		} `json:"GGP,omitempty"`
		Ghs struct {
			Currency
		} `json:"GHS,omitempty"`
		Gip struct {
			Currency
		} `json:"GIP,omitempty"`
		Gmd struct {
			Currency
		} `json:"GMD,omitempty"`
		Gnf struct {
			Currency
		} `json:"GNF,omitempty"`
		Gtq struct {
			Currency
		} `json:"GTQ,omitempty"`
		Gyd struct {
			Currency
		} `json:"GYD,omitempty"`
		Hkd struct {
			Currency
		} `json:"HKD,omitempty"`
		Hnl struct {
			Currency
		} `json:"HNL,omitempty"`
		Hrk struct {
			Currency
		} `json:"HRK,omitempty"`
		Htg struct {
			Currency
		} `json:"HTG,omitempty"`
		Huf struct {
			Currency
		} `json:"HUF,omitempty"`
		Idr struct {
			Currency
		} `json:"IDR,omitempty"`
		Ils struct {
			Currency
		} `json:"ILS,omitempty"`
		Imp struct {
			Currency
		} `json:"IMP,omitempty"`
		Inr struct {
			Currency
		} `json:"INR,omitempty"`
		Iqd struct {
			Currency
		} `json:"IQD,omitempty"`
		Irr struct {
			Currency
		} `json:"IRR,omitempty"`
		Isk struct {
			Currency
		} `json:"ISK,omitempty"`
		Jep struct {
			Currency
		} `json:"JEP,omitempty"`
		Jmd struct {
			Currency
		} `json:"JMD,omitempty"`
		Jod struct {
			Currency
		} `json:"JOD,omitempty"`
		Jpy struct {
			Currency
		} `json:"JPY,omitempty"`
		Kes struct {
			Currency
		} `json:"KES,omitempty"`
		Kgs struct {
			Currency
		} `json:"KGS,omitempty"`
		Khr struct {
			Currency
		} `json:"KHR,omitempty"`
		Kmf struct {
			Currency
		} `json:"KMF,omitempty"`
		Kpw struct {
			Currency
		} `json:"KPW,omitempty"`
		Krw struct {
			Currency
		} `json:"KRW,omitempty"`
		Kwd struct {
			Currency
		} `json:"KWD,omitempty"`
		Kyd struct {
			Currency
		} `json:"KYD,omitempty"`
		Kzt struct {
			Currency
		} `json:"KZT,omitempty"`
		Lak struct {
			Currency
		} `json:"LAK,omitempty"`
		Lbp struct {
			Currency
		} `json:"LBP,omitempty"`
		Lkr struct {
			Currency
		} `json:"LKR,omitempty"`
		Lrd struct {
			Currency
		} `json:"LRD,omitempty"`
		Lsl struct {
			Currency
		} `json:"LSL,omitempty"`
		Ltc struct {
			Currency
		} `json:"LTC,omitempty"`
		Ltl struct {
			Currency
		} `json:"LTL,omitempty"`
		Lvl struct {
			Currency
		} `json:"LVL,omitempty"`
		Lyd struct {
			Currency
		} `json:"LYD,omitempty"`
		Mad struct {
			Currency
		} `json:"MAD,omitempty"`
		Matic struct {
			Currency
		} `json:"MATIC,omitempty"`
		Mdl struct {
			Currency
		} `json:"MDL,omitempty"`
		Mga struct {
			Currency
		} `json:"MGA,omitempty"`
		Mkd struct {
			Currency
		} `json:"MKD,omitempty"`
		Mmk struct {
			Currency
		} `json:"MMK,omitempty"`
		Mnt struct {
			Currency
		} `json:"MNT,omitempty"`
		Mop struct {
			Currency
		} `json:"MOP,omitempty"`
		Mro struct {
			Currency
		} `json:"MRO,omitempty"`
		Mur struct {
			Currency
		} `json:"MUR,omitempty"`
		Mvr struct {
			Currency
		} `json:"MVR,omitempty"`
		Mwk struct {
			Currency
		} `json:"MWK,omitempty"`
		Mxn struct {
			Currency
		} `json:"MXN,omitempty"`
		Myr struct {
			Currency
		} `json:"MYR,omitempty"`
		Mzn struct {
			Currency
		} `json:"MZN,omitempty"`
		Nad struct {
			Currency
		} `json:"NAD,omitempty"`
		Ngn struct {
			Currency
		} `json:"NGN,omitempty"`
		Nio struct {
			Currency
		} `json:"NIO,omitempty"`
		Nok struct {
			Currency
		} `json:"NOK,omitempty"`
		Npr struct {
			Currency
		} `json:"NPR,omitempty"`
		Nzd struct {
			Currency
		} `json:"NZD,omitempty"`
		Omr struct {
			Currency
		} `json:"OMR,omitempty"`
		Pab struct {
			Currency
		} `json:"PAB,omitempty"`
		Pen struct {
			Currency
		} `json:"PEN,omitempty"`
		Pgk struct {
			Currency
		} `json:"PGK,omitempty"`
		Php struct {
			Currency
		} `json:"PHP,omitempty"`
		Pkr struct {
			Currency
		} `json:"PKR,omitempty"`
		Pln struct {
			Currency
		} `json:"PLN,omitempty"`
		Pyg struct {
			Currency
		} `json:"PYG,omitempty"`
		Qar struct {
			Currency
		} `json:"QAR,omitempty"`
		Ron struct {
			Currency
		} `json:"RON,omitempty"`
		Rsd struct {
			Currency
		} `json:"RSD,omitempty"`
		Rub struct {
			Currency
		} `json:"RUB,omitempty"`
		Rwf struct {
			Currency
		} `json:"RWF,omitempty"`
		Sar struct {
			Currency
		} `json:"SAR,omitempty"`
		Sbd struct {
			Currency
		} `json:"SBD,omitempty"`
		Scr struct {
			Currency
		} `json:"SCR,omitempty"`
		Sdg struct {
			Currency
		} `json:"SDG,omitempty"`
		Sek struct {
			Currency
		} `json:"SEK,omitempty"`
		Sgd struct {
			Currency
		} `json:"SGD,omitempty"`
		Shp struct {
			Currency
		} `json:"SHP,omitempty"`
		Sll struct {
			Currency
		} `json:"SLL,omitempty"`
		Sol struct {
			Currency
		} `json:"SOL,omitempty"`
		Sos struct {
			Currency
		} `json:"SOS,omitempty"`
		Srd struct {
			Currency
		} `json:"SRD,omitempty"`
		Std struct {
			Currency
		} `json:"STD,omitempty"`
		Svc struct {
			Currency
		} `json:"SVC,omitempty"`
		Syp struct {
			Currency
		} `json:"SYP,omitempty"`
		Szl struct {
			Currency
		} `json:"SZL,omitempty"`
		Thb struct {
			Currency
		} `json:"THB,omitempty"`
		Tjs struct {
			Currency
		} `json:"TJS,omitempty"`
		Tmt struct {
			Currency
		} `json:"TMT,omitempty"`
		Tnd struct {
			Currency
		} `json:"TND,omitempty"`
		Top struct {
			Currency
		} `json:"TOP,omitempty"`
		Try struct {
			Currency
		} `json:"TRY,omitempty"`
		Ttd struct {
			Currency
		} `json:"TTD,omitempty"`
		Twd struct {
			Currency
		} `json:"TWD,omitempty"`
		Tzs struct {
			Currency
		} `json:"TZS,omitempty"`
		Uah struct {
			Currency
		} `json:"UAH,omitempty"`
		Ugx struct {
			Currency
		} `json:"UGX,omitempty"`
		Usd struct {
			Currency
		} `json:"USD,omitempty"`
		Uyu struct {
			Currency
		} `json:"UYU,omitempty"`
		Uzs struct {
			Currency
		} `json:"UZS,omitempty"`
		Vef struct {
			Currency
		} `json:"VEF,omitempty"`
		Vnd struct {
			Currency
		} `json:"VND,omitempty"`
		Vuv struct {
			Currency
		} `json:"VUV,omitempty"`
		Wst struct {
			Currency
		} `json:"WST,omitempty"`
		Xaf struct {
			Currency
		} `json:"XAF,omitempty"`
		Xag struct {
			Currency
		} `json:"XAG,omitempty"`
		Xau struct {
			Currency
		} `json:"XAU,omitempty"`
		Xcd struct {
			Currency
		} `json:"XCD,omitempty"`
		Xdr struct {
			Currency
		} `json:"XDR,omitempty"`
		Xof struct {
			Currency
		} `json:"XOF,omitempty"`
		Xpf struct {
			Currency
		} `json:"XPF,omitempty"`
		Xrp struct {
			Currency
		} `json:"XRP,omitempty"`
		Yer struct {
			Currency
		} `json:"YER,omitempty"`
		Zar struct {
			Currency
		} `json:"ZAR,omitempty"`
		Zmk struct {
			Currency
		} `json:"ZMK,omitempty"`
		Zmw struct {
			Currency
		} `json:"ZMW,omitempty"`
		Zwl struct {
			Currency
		} `json:"ZWL,omitempty"`
	} `json:"data,omitempty"`
}
