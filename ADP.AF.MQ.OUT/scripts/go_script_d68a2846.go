package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

// ---------- Целевые структуры (соответствуют inner.txt) ----------
// mapstructure-теги читают из out (без суффиксов),
// json-теги пишут ключи с суффиксами _innerNNN.
// Имена типов уникальны в пакете main благодаря суффиксу _d68a2846.

type InnerPassportValidity_d68a2846 struct {
	Result string `mapstructure:"result" json:"result_inner831"`
}

type InnerInnSnilsMatch_d68a2846 struct {
	Result string `mapstructure:"result" json:"result_inner608"`
}

type InnerFnsDetails_d68a2846 struct {
	InnActive bool `mapstructure:"innActive" json:"innActive_inner667"`
}

type InnerFnsCheck_d68a2846 struct {
	Result  string                     `mapstructure:"result"  json:"result_inner904"`
	Details InnerFnsDetails_d68a2846   `mapstructure:"details" json:"details_inner22"`
}

type InnerMvdCheck_d68a2846 struct {
	Result string `mapstructure:"result" json:"result_inner513"`
}

type InnerFsspCheck_d68a2846 struct {
	Result      string        `mapstructure:"result"      json:"result_inner391"`
	Proceedings []interface{} `mapstructure:"proceedings" json:"proceedings_inner176"`
}

type InnerBankruptcyCheck_d68a2846 struct {
	Result string `mapstructure:"result" json:"result_inner725"`
}

type InnerChecks_d68a2846 struct {
	PassportValidity InnerPassportValidity_d68a2846 `mapstructure:"passportValidity" json:"passportValidity_inner56"`
	InnSnilsMatch    InnerInnSnilsMatch_d68a2846    `mapstructure:"innSnilsMatch"    json:"innSnilsMatch_inner294"`
	FnsCheck         InnerFnsCheck_d68a2846         `mapstructure:"fnsCheck"         json:"fnsCheck_inner375"`
	MvdCheck         InnerMvdCheck_d68a2846         `mapstructure:"mvdCheck"         json:"mvdCheck_inner148"`
	FsspCheck        InnerFsspCheck_d68a2846        `mapstructure:"fsspCheck"        json:"fsspCheck_inner820"`
	BankruptcyCheck  InnerBankruptcyCheck_d68a2846  `mapstructure:"bankruptcyCheck"  json:"bankruptcyCheck_inner459"`
}

type InnerBiometryResult_d68a2846 struct {
	Liveness  string  `mapstructure:"liveness"  json:"liveness_inner888"`
	FaceMatch float64 `mapstructure:"faceMatch" json:"faceMatch_inner241"`
}

type InnerMessage_d68a2846 struct {
	Status         string                        `mapstructure:"status"         json:"status_inner482"`
	VerifiedAt     string                        `mapstructure:"verifiedAt"     json:"verifiedAt_inner117"`
	Checks         InnerChecks_d68a2846          `mapstructure:"checks"         json:"checks_inner739"`
	BiometryResult InnerBiometryResult_d68a2846  `mapstructure:"biometryResult" json:"biometryResult_inner63"`
	RiskFlags      []interface{}                 `mapstructure:"riskFlags"      json:"riskFlags_inner507"`
}

// ---------- Основная функция обработки ----------

func Process_d68a2846(msg Message, output MessageSender) Message {
	log.Printf("🐹 Go Script: обработка сообщения")
	log.Printf("   • ID: %s", msg.ID)
	log.Printf("   • Body: %s", string(msg.Body))

	// --- 1. Гарантируем, что три заголовка есть в msg.Headers ---
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

	// --- 2. Маппинг тела out -> inner ---
	var raw map[string]interface{}
	if err := json.Unmarshal(msg.Body, &raw); err != nil {
		log.Printf("❌ Ошибка парсинга JSON: %v", err)
		sendOrWarn_d68a2846(msg, output)
		return msg
	}

	var inner InnerMessage_d68a2846
	if err := mapstructure.Decode(raw, &inner); err != nil {
		log.Printf("❌ Ошибка mapstructure.Decode: %v", err)
		sendOrWarn_d68a2846(msg, output)
		return msg
	}

	newBody, err := json.Marshal(inner)
	if err != nil {
		log.Printf("❌ Ошибка сериализации JSON: %v", err)
		sendOrWarn_d68a2846(msg, output)
		return msg
	}

	msg.Body = newBody
	msg.IsJSON = true

	// Data в Message объявлен как map[string]interface{},
	// поэтому структуру туда класть нельзя — распаковываем newBody обратно.
	var newData map[string]interface{}
	if err := json.Unmarshal(newBody, &newData); err == nil {
		msg.Data = newData
	} else {
		log.Printf("⚠️ Не удалось распаковать newBody в Data: %v", err)
	}

	log.Printf("   • Body после маппинга: %s", string(msg.Body))
	log.Printf("   • Headers (уйдут в HTTP): %v", msg.Headers)
	sendOrWarn_d68a2846(msg, output)
	return msg
}

func sendOrWarn_d68a2846(msg Message, output MessageSender) {
	if output != nil {
		if err := output.Send(msg); err != nil {
			log.Printf("❌ Ошибка отправки: %v", err)
		}
	} else {
		log.Printf("⚠️ output sender не установлен")
	}
}