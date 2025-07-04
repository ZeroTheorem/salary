package main

import (
	"fmt"
	"github.com/leekchan/accounting"
	tele "gopkg.in/telebot.v4"
	"log"
	"strconv"
	"time"
)

var (
	totalMsg = `
Доход за текущий месяц составил:

<b>%v</b>

----------------------------
<b>Яндекс</b>
Премия: <b>%v</b>
Аванс: <b>%v</b>
Зарплата: <b>%v</b>

Всего: <b>%v</b>

<b>Вклады</b>
Всего: <b>%v</b>

----------------------------
`
	configMsg = `
Ставка за этот месяц: %v
Ставка за предыдущий месяц: %v
Кол-во смен за этот месяц: %v
Кол-во смен за предыдущий месяц: %v
Премия: %v%%
Премия наставника: %v
Компенсация за интернет: %v
Проценты по вкладу: %v
`
	option string
)

type config struct {
	currentPerHour      float64
	previousPerHour     float64
	currentCountShifts  float64
	previousCountShifts float64
	bonusPercent        float64
	coachBonus          float64
	internetBonus       float64
	depositPercent      float64
}

func main() {
	//-----------Initial bot--------------------
	pref := tele.Settings{
		Token:     "8095045960:AAHihLPiP0-ynJKVus-EMTh5eOQVdGQlwmg",
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tele.ModeHTML,
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	//-----------Create main menu---------------
	mainMenu := &tele.ReplyMarkup{}
	settingsBtn := mainMenu.Data("Настройки", "settings")
	calculateBtn := mainMenu.Data("Расчитать", "calculate")
	mainMenu.Inline(
		mainMenu.Row(calculateBtn),
		mainMenu.Row(settingsBtn),
	)

	//-----------Create settings menu------------
	settingsMenu := &tele.ReplyMarkup{}
	currentPerHourBtn := settingsMenu.Data("Ставка за этот месяц", "currentPerHour")
	previousPerHourBtn := settingsMenu.Data("Cтавка за прошлый месяц", "previousPerHour")
	currentCountShiftsBtn := settingsMenu.Data("Кол-во смен за этот месяц", "currentCountShifts")
	previousCountShiftsBtn := settingsMenu.Data("Кол-во смен за прошлый месяц", "previousCountShifts")
	bonusPercentBtn := settingsMenu.Data("Премия", "bonusPercent")
	coachBonusBtn := settingsMenu.Data("Премия наставника", "coachBonus")
	internetBonusBtn := settingsMenu.Data("Компенсация за интернет", "internetBonus")
	depositPercentBtn := settingsMenu.Data("Проценты по вкладу", "depositPercent")
	previousBtn := settingsMenu.Data("Назад", "previous")
	settingsMenu.Inline(
		settingsMenu.Row(currentPerHourBtn),
		settingsMenu.Row(previousPerHourBtn),
		settingsMenu.Row(currentCountShiftsBtn),
		settingsMenu.Row(previousCountShiftsBtn),
		settingsMenu.Row(bonusPercentBtn),
		settingsMenu.Row(coachBonusBtn),
		settingsMenu.Row(internetBonusBtn),
		settingsMenu.Row(depositPercentBtn),
		settingsMenu.Row(previousBtn),
	)

	//-----------Initial config------------------
	cfg := &config{}

	//-----------Initial formater----------------
	ac := &accounting.Accounting{
		Symbol:    "₽",
		Precision: 2,
		Thousand:  " ",
		Decimal:   ".",
	}

	//-----------Create handlres-----------------
	b.Handle("/start", func(c tele.Context) error {
		return c.Send(createConfigMsg(cfg), mainMenu)
	})
	b.Handle(&calculateBtn, func(c tele.Context) error {
		return c.Send(calculateSalary(cfg, ac))
	})
	b.Handle(&previousBtn, func(c tele.Context) error {
		return c.Send(createConfigMsg(cfg), mainMenu)
	})
	b.Handle(&settingsBtn, func(c tele.Context) error {
		return c.Send("Настройки", settingsMenu)
	})
	b.Handle(&currentPerHourBtn, func(c tele.Context) error {
		option = c.Callback().Unique
		return c.Send("Введите значение")
	})
	b.Handle(&previousPerHourBtn, func(c tele.Context) error {
		option = c.Callback().Unique
		return c.Send("Введите значение")
	})
	b.Handle(&currentCountShiftsBtn, func(c tele.Context) error {
		option = c.Callback().Unique
		return c.Send("Введите значение")
	})
	b.Handle(&previousCountShiftsBtn, func(c tele.Context) error {
		option = c.Callback().Unique
		return c.Send("Введите значение")
	})
	b.Handle(&bonusPercentBtn, func(c tele.Context) error {
		option = c.Callback().Unique
		return c.Send("Введите значение")
	})
	b.Handle(&coachBonusBtn, func(c tele.Context) error {
		option = c.Callback().Unique
		return c.Send("Введите значение")
	})
	b.Handle(&internetBonusBtn, func(c tele.Context) error {
		option = c.Callback().Unique
		return c.Send("Введите значение")
	})
	b.Handle(&depositPercentBtn, func(c tele.Context) error {
		option = c.Callback().Unique
		return c.Send("Введите значение")
	})
	b.Handle(tele.OnText, func(c tele.Context) error {
		val := c.Message().Text
		ival, err := strconv.ParseFloat(val, 64)

		if err != nil {
			return c.Send("Введите число")
		}

		switch option {
		case "currentPerHour":
			cfg.currentPerHour = ival
			option = ""
			return c.Send(createConfigMsg(cfg), settingsMenu)
		case "previousPerHour":
			cfg.previousPerHour = ival
			option = ""
			return c.Send(createConfigMsg(cfg), settingsMenu)
		case "currentCountShifts":
			cfg.currentCountShifts = ival
			option = ""
			return c.Send(createConfigMsg(cfg), settingsMenu)
		case "previousCountShifts":
			cfg.previousCountShifts = ival
			option = ""
			return c.Send(createConfigMsg(cfg), settingsMenu)
		case "bonusPercent":
			cfg.bonusPercent = ival
			option = ""
			return c.Send(createConfigMsg(cfg), settingsMenu)
		case "coachBonus":
			cfg.coachBonus = ival
			option = ""
			return c.Send(createConfigMsg(cfg), settingsMenu)
		case "internetBonus":
			cfg.internetBonus = ival
			option = ""
			return c.Send(createConfigMsg(cfg), settingsMenu)
		case "depositPercent":
			cfg.depositPercent = ival
			option = ""
			return c.Send(createConfigMsg(cfg), settingsMenu)

		default:
			return c.Send("Выберите нужную опицию", settingsMenu)
		}

	})

	//-----------Start bot-----------------------
	b.Start()

}

func calculateSalary(cfg *config, ac *accounting.Accounting) string {
	currentMonth := cfg.currentPerHour * 11 * cfg.currentCountShifts
	previousMonth := cfg.previousPerHour * 11 * cfg.previousCountShifts
	currentMonthWithTax := currentMonth - (currentMonth*13)/100
	previousMonthWithTax := previousMonth - (previousMonth*13)/100
	bonus := previousMonthWithTax*cfg.bonusPercent/100 + cfg.coachBonus
	salary := previousMonthWithTax*60/100 + cfg.internetBonus
	avance := currentMonthWithTax * 40 / 100
	totalsalary := salary + avance + bonus
	total := totalsalary + cfg.depositPercent
	msg := fmt.Sprintf(
		totalMsg,
		ac.FormatMoney(total),
		ac.FormatMoney(bonus),
		ac.FormatMoney(avance),
		ac.FormatMoney(salary),
		ac.FormatMoney(totalsalary),
		ac.FormatMoney(cfg.depositPercent),
	)
	return msg
}

func createConfigMsg(cfg *config) string {
	return fmt.Sprintf(configMsg,
		cfg.currentPerHour,
		cfg.previousPerHour,
		cfg.currentCountShifts,
		cfg.previousCountShifts,
		cfg.bonusPercent,
		cfg.coachBonus,
		cfg.internetBonus,
		cfg.depositPercent,
	)
}
