package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

// ---------- Целевые структуры (соответствуют inner.txt) ----------
// mapstructure-теги читают из out (без суффиксов),
// json-теги пишут ключи с суффиксами _innerNNN.

type InnerResult struct {
	Result string `mapstructure:"result" json:"result_inner831"`
}

type InnerFnsDetails struct {
	InnActive bool `mapstructure:"innActive" json:"innActive_inner667"`
}

type InnerFnsCheck struct {
	Result  string           `mapstructure:"result" json:"result_inner904"`
	Details InnerFnsDetails  `mapstructure:"details" json:"details_inner22"`
}

type InnerFsspCheck struct {
	Result      string        `mapstructure:"result" json:"result_inner391"`
	Proceedings []interface{} `mapstructure:"proceedings" json:"proceedings_inner176"`
}

type InnerChecks struct {
	PassportValidity InnerResult    `mapstructure:"passportValidity" json:"passportValidity_inner56"`
	InnSnilsMatch    InnerResult    `mapstructure:"innSnilsMatch" json:"innSnilsMatch_inner294"`
	FnsCheck         InnerFnsCheck  `mapstructure:"fnsCheck" json:"fnsCheck_inner375"`
	MvdCheck         InnerResult    `mapstructure:"mvdCheck" json:"mvdCheck_inner148"`
	FsspCheck        InnerFsspCheck `mapstructure:"fsspCheck" json:"fsspCheck_inner820"`
	BankruptcyCheck  InnerResult    `mapstructure:"bankruptcyCheck" json:"bankruptcyCheck_inner459"`
}

type InnerBiometryResult struct {
	Liveness  string  `mapstructure:"liveness" json:"liveness_inner888"`
	FaceMatch float64 `mapstructure:"faceMatch" json:"faceMatch_inner241"`
}

type InnerMessage struct {
	Status         string             `mapstructure:"status" json:"status_inner482"`
	VerifiedAt     string             `mapstructure:"verifiedAt" json:"verifiedAt_inner117"`
	Checks         InnerChecks        `mapstructure:"checks" json:"checks_inner739"`
	BiometryResult InnerBiometryResult `mapstructure:"biometryResult" json:"biometryResult_inner63"`
	RiskFlags      []interface{}      `mapstructure:"riskFlags" json:"riskFlags_inner507"`
}

// ---------- Основная функция обработки ----------

func Process_4b405a14(msg Message, output MessageSender) Message {
	log.Printf("🐹 Go Script: обработка сообщения")
	log.Printf("   • ID: %s", msg.ID)
	log.Printf("   • Body: %s", string(msg.Body))

	// --- 1. Гарантируем, что три заголовка есть в msg.Headers ---
	// MQInput уже положил их туда, но на случай отсутствия — подстрахуемся.
	if msg.Headers == nil {
		msg.Headers = make(map[string]string)
	}
	for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
		if _, ok := msg.Headers[key]; ok {
			continue
		}
		if msg.Data != nil {
			if v, ok := msg.Data[key].(string); ok && v != "" {
				msg.Headers[key] = v
			}
		}
	}

	// --- 2. Маппинг тела out -> inner (структуры InnerMessage из предыдущего ответа) ---
	var raw map[string]interface{}
	if err := json.Unmarshal(msg.Body, &raw); err != nil {
		log.Printf("❌ Ошибка парсинга JSON: %v", err)
		sendOrWarn(msg, output)
		return msg
	}

	var inner InnerMessage
	if err := mapstructure.Decode(raw, &inner); err != nil {
		log.Printf("❌ Ошибка mapstructure.Decode: %v", err)
		sendOrWarn(msg, output)
		return msg
	}

	newBody, err := json.Marshal(inner)
	if err != nil {
		log.Printf("❌ Ошибка сериализации JSON: %v", err)
		sendOrWarn(msg, output)
		return msg
	}

	msg.Body = newBody
	msg.IsJSON = true
	msg.Data = inner

	log.Printf("   • Body после маппинга: %s", string(msg.Body))
	log.Printf("   • Headers (уйдут в HTTP): %v", msg.Headers)
	sendOrWarn(msg, output)
	return msg
}

func sendOrWarn(msg Message, output MessageSender) {
	if output != nil {
		if err := output.Send(msg); err != nil {
			log.Printf("❌ Ошибка отправки: %v", err)
		}
	} else {
		log.Printf("⚠️ output sender не установлен")
	}
}