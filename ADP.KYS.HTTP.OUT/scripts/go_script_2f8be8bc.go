package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

// ---------- Целевые структуры (соответствуют out.txt) ----------
// mapstructure-теги читают из inner (ключи с суффиксами _innerNNN),
// json-теги пишут чистые ключи для out.

type PassportOut struct {
	Series     string `mapstructure:"series_inner381"     json:"series"`
	Number     string `mapstructure:"number_inner95"      json:"number"`
	IssueDate  string `mapstructure:"issueDate_inner447"  json:"issueDate"`
	IssuerCode string `mapstructure:"issuerCode_inner862" json:"issuerCode"`
}

type ClientOut struct {
	ClientID   string      `mapstructure:"clientId_inner42"   json:"clientId"`
	FirstName  string      `mapstructure:"firstName_inner891" json:"firstName"`
	LastName   string      `mapstructure:"lastName_inner256"  json:"lastName"`
	MiddleName string      `mapstructure:"middleName_inner77" json:"middleName"`
	BirthDate  string      `mapstructure:"birthDate_inner503" json:"birthDate"`
	Inn        string      `mapstructure:"inn_inner618"       json:"inn"`
	Snils      string      `mapstructure:"snils_inner9"       json:"snils"`
	Passport   PassportOut `mapstructure:"passport_inner724"  json:"passport"`
}

type BiometryOut struct {
	SelfieURL        string `mapstructure:"selfieUrl_inner777"        json:"selfieUrl"`
	LivenessVideoURL string `mapstructure:"livenessVideoUrl_inner312" json:"livenessVideoUrl"`
}

type ConsentsOut struct {
	PdnProcessing bool   `mapstructure:"pdnProcessing_inner126" json:"pdnProcessing"`
	GrantedAt     string `mapstructure:"grantedAt_inner961"     json:"grantedAt"`
}

type OutMessage struct {
	Client           ClientOut   `mapstructure:"client_inner137"         json:"client"`
	VerificationType string      `mapstructure:"verificationType_inner530" json:"verificationType"`
	Checks           []string    `mapstructure:"checks_inner214"          json:"checks"`
	Biometry         BiometryOut `mapstructure:"biometry_inner68"         json:"biometry"`
	Consents         ConsentsOut `mapstructure:"consents_inner845"        json:"consents"`
}

// ---------- Основная функция обработки ----------

func Process_2f8be8bc(msg Message, output MessageSender) Message {
	log.Printf("🐹 Go Script: обработка сообщения")
	log.Printf("   • ID: %s", msg.ID)
	log.Printf("   • Body: %s", string(msg.Body))

	// --- 1. Перенос трёх HTTP-заголовков в msg.Headers (если их там нет) ---
	if msg.Headers == nil {
		msg.Headers = make(map[string]string)
	}
	for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
		// если заголовок уже есть в Headers — не трогаем
		if _, ok := msg.Headers[key]; ok {
			continue
		}
		// иначе пробуем взять из Data (например, если платформа положила туда)
		if msg.Data != nil {
			if v, ok := msg.Data[key].(string); ok && v != "" {
				msg.Headers[key] = v
				delete(msg.Data, key) // убираем из тела, чтобы не дублировать
			}
		}
	}

	// --- 2. Маппинг тела inner -> out (как раньше, через mapstructure) ---
	var raw map[string]interface{}
	if err := json.Unmarshal(msg.Body, &raw); err != nil {
		log.Printf("❌ Ошибка парсинга JSON: %v", err)
		sendOrWarn(msg, output)
		return msg
	}

	var out OutMessage
	if err := mapstructure.Decode(raw, &out); err != nil {
		log.Printf("❌ Ошибка mapstructure.Decode: %v", err)
		sendOrWarn(msg, output)
		return msg
	}

	newBody, err := json.Marshal(out)
	if err != nil {
		log.Printf("❌ Ошибка сериализации JSON: %v", err)
		sendOrWarn(msg, output)
		return msg
	}

	msg.Body = newBody
	msg.IsJSON = true
	msg.Data = out

	log.Printf("   • Body после маппинга: %s", string(msg.Body))
	log.Printf("   • Headers: %v", msg.Headers)
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