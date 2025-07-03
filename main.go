package main

import (
	"fmt"
	"log"
	"time"

	tele "gopkg.in/telebot.v4"
)

var (
	totalMsg = `
До вычета налога: %v
После вычета налога: %v
Аванс: %v
Зарплата: %v
`
	configMsg = `
Ставка за этот месяц: %v
Ставка за предыдущий месяц: %v
Кол-во смен за этот месяц: %v
Кол-во смен за предыдущий месяц: %v
Премия: %v%%
Премия наставника: %v
`
)

type config struct {
	currentPerHour      int
	previousPerHour     int
	currentCountShifts  int
	previousCountShifts int
	bonusPercent        int
	coachBonus          int
}

func main() {
	pref := tele.Settings{
		Token:  "8095045960:AAHihLPiP0-ynJKVus-EMTh5eOQVdGQlwmg",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	cfg := &config{}
	b.Handle("/start", func(c tele.Context) error {
		c.Send(calculateSalary(cfg))
		return c.Send(createConfigMsg(cfg))
	})

	b.Start()

}

func calculateSalary(cfg *config) string {
	currentMonth := (cfg.currentPerHour * 11) * cfg.currentCountShifts
	previousMonth := (cfg.previousPerHour * 11) * cfg.previousCountShifts
	bonus := (previousMonth * cfg.bonusPercent) / 100
	tax := ((currentMonth + bonus) * 13) / 100
	total := (currentMonth + bonus) + cfg.coachBonus
	totalWithTax := total - tax
	salary := (totalWithTax * 60) / 100
	avance := (totalWithTax * 40) / 100
	msg := fmt.Sprintf(totalMsg, total, totalWithTax, avance, salary)
	return msg
}

func createConfigMsg(cfg *config) string {
	return fmt.Sprintf(configMsg,
		cfg.currentPerHour,
		cfg.previousPerHour,
		cfg.currentCountShifts,
		cfg.previousCountShifts,
		cfg.bonusPercent,
		cfg.coachBonus)
}
