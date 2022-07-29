package worker

import (
	"fmt"
	"time"

	"github.com/adnvilla/boletia/repository"
)

func ApiModelToRepoModel(apimodel CurrencyLatest) ([]repository.Currencies, time.Time) {
	var lastUpdatedAt = apimodel.Meta.LastUpdatedAt
	var currencies = []repository.Currencies{}

	currencies = addValidCurrency(currencies, apimodel.Data.Ada.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Aed.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Afn.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.All.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Amd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Ang.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Aoa.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Ars.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Aud.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Avax.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Awg.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Azn.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Bam.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Bbd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Bdt.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Bgn.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Bhd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Bif.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Bmd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Bnb.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Bnd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Bob.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Brl.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Bsd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Btc.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Btn.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Bwp.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Byn.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Byr.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Bzd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Cad.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Cdf.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Chf.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Clf.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Clp.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Cny.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Cop.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Crc.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Cuc.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Cup.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Cve.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Czk.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Djf.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Dkk.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Dop.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Dot.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Dzd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Egp.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Ern.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Etb.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Eth.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Eur.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Fjd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Fkp.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Gbp.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Gel.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Ggp.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Ghs.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Gip.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Gmd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Gnf.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Gtq.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Gyd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Hkd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Hnl.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Hrk.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Htg.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Huf.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Idr.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Ils.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Imp.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Inr.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Iqd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Irr.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Isk.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Jep.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Jmd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Jod.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Jpy.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Kes.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Kgs.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Khr.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Kmf.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Kpw.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Krw.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Kwd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Kyd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Kzt.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Lak.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Lbp.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Lkr.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Lrd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Lsl.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Ltc.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Ltl.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Lvl.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Lyd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Mad.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Matic.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Mdl.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Mga.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Mkd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Mmk.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Mnt.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Mop.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Mro.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Mur.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Mvr.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Mwk.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Mxn.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Myr.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Mzn.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Nad.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Ngn.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Nio.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Nok.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Npr.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Nzd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Omr.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Pab.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Pen.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Pgk.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Php.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Pkr.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Pln.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Pyg.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Qar.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Ron.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Rsd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Rub.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Rwf.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Sar.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Sbd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Scr.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Sdg.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Sek.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Sgd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Shp.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Sll.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Sol.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Sos.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Srd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Std.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Svc.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Syp.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Szl.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Thb.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Tjs.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Tmt.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Tnd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Top.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Try.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Ttd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Twd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Tzs.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Uah.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Ugx.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Usd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Uyu.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Uzs.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Vef.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Vnd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Vuv.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Wst.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Xaf.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Xag.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Xau.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Xcd.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Xdr.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Xof.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Xpf.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Xrp.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Yer.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Zar.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Zmk.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Zmw.Currency, lastUpdatedAt)
	currencies = addValidCurrency(currencies, apimodel.Data.Zwl.Currency, lastUpdatedAt)

	return currencies, lastUpdatedAt
}

func addValidCurrency(currencies []repository.Currencies, currency Currency, lastUpdatedAt time.Time) []repository.Currencies {
	c := repository.Currencies{
		Currency:      currency.Code,
		Value:         fmt.Sprintf("%f", currency.Value),
		LastUpdatedAt: lastUpdatedAt,
	}
	if c.Currency != "" && c.Value != "" {
		c.LastUpdatedAt = lastUpdatedAt
		currencies = append(currencies, c)
	}
	return currencies
}

func RepoModelTApimodel() {

}