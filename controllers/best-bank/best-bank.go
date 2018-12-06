package bestBankController

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/oreuta/easytrip/models"
	"github.com/oreuta/easytrip/services/best-bank"
	"github.com/oreuta/easytrip/translate"
)

type bestBankController struct {
	beego.Controller
	BestService bestBankService.BestBankServiceInterface
}

func New(s bestBankService.BestBankServiceInterface) *bestBankController {
	return &bestBankController{BestService: s}
}

func (r *bestBankController) Get() {

	translate := translate.New()
	lang := r.GetString("lang")
	if lang != "" {
		translate.Lang = lang
		r.Ctx.SetCookie("lang", translate.Lang)
	} else {
		translate.Lang = r.Ctx.GetCookie("lang")
		if translate.Lang == "" {
			translate.Lang = "en-US"
		}
	}
	translate.Path = "conf/locale_" + translate.Lang + ".ini"
	r.Data["i18n"] = translate.Tr

	toolbox.StatisticsMap.AddStatistics("GET", "/best", "&controllers.bestBankController.bestBankController", time.Duration(15000))
	inpData := models.MainRequest{
		Currency: r.GetStrings("currency"),
		Option:   r.GetString("option"),
		Bank:     r.GetStrings("bank"),
	}

	{
		i := 0
		if inpData.Currency == nil {
			r.Data["warningCurrency"] = "*Select Currency"
			i++
		}
		if inpData.Bank == nil {
			r.Data["warningBank"] = "*Select Bank"
			i++
		}
		if i > 0 {
			r.TplName = "index.tpl"
			r.Data["isWarnMess"] = true
			return
		}
	}

	sale, buy, err := r.BestService.GetBestBanks(inpData)
	if err != nil {
		beego.Error("GetBestBanks func in BestService: %v", err)
		return
	}
	r.Layout = "bestBank_layout.tpl"
	r.TplName = "bestBank.tpl"
	r.Data["Buy"] = buy
	r.Data["TitleBuy"] = ""
	r.Data["Sale"] = sale
	r.Data["TitleSale"] = ""
	if buy != nil {
		r.Data["TitleBuy"] = "Best Buy"
	}
	if sale != nil {
		r.Data["TitleSale"] = "Best Sale"
	}
}
