package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

type InnerKycClient_<ID> struct {
    FirstName  string             `mapstructure:"firstName_inner201"  json:"firstName"`
    LastName   string             `mapstructure:"lastName_inner202"   json:"lastName"`
    MiddleName string             `mapstructure:"middleName_inner203" json:"middleName"`
    BirthDate  string             `mapstructure:"birthDate_inner204"  json:"birthDate"`
    Inn        string             `mapstructure:"inn_inner205"        json:"inn"`
    Snils      string             `mapstructure:"snils_inner206"      json:"snils"`
    Passport   InnerPassport_<ID> `mapstructure:"passport_inner207"   json:"passport"`
}

type InnerBiometryRequest_<ID> struct {
    SelfieURL        string `mapstructure:"selfieUrl_inner211"        json:"selfieUrl"`
    LivenessVideoURL string `mapstructure:"livenessVideoUrl_inner212" json:"livenessVideoUrl"`
}

type InnerMessage_<ID> struct {
    RequestID     string                    `mapstructure:"requestId_inner221"     json:"requestId"`
    ApplicationID string                    `mapstructure:"applicationId_inner222" json:"applicationId"`
    Client        InnerKycClient_<ID>       `mapstructure:"client_inner223"        json:"client"`
    Checks        []string                  `mapstructure:"checks_inner224"        json:"checks"`
    Biometry      InnerBiometryRequest_<ID> `mapstructure:"biometry_inner225"      json:"biometry"`
}

// Process_ce39f6a8 - основная функция обработки сообщения.
// msg    - входящее сообщение
// output - MessageSender, куда отправлять результат (может быть nil).
//
// Вся отправка сообщений дальше делается через output.Send(...).
// Если функция ничего не отправляет, сообщение дальше не уйдёт.
// Возвращаемое значение игнорируется — можно вернуть msg или любое другое.
func Process_ce39f6a8(msg Message, output MessageSender) Message {
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