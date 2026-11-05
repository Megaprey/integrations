package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

type BkiClientOut_d4e5f6a4 struct {
    FirstName      string `mapstructure:"firstName_inner401"      json:"firstName"`
    LastName       string `mapstructure:"lastName_inner402"       json:"lastName"`
    MiddleName     string `mapstructure:"middleName_inner403"     json:"middleName"`
    BirthDate      string `mapstructure:"birthDate_inner404"      json:"birthDate"`
    PassportSeries string `mapstructure:"passportSeries_inner405" json:"passportSeries"`
    PassportNumber string `mapstructure:"passportNumber_inner406" json:"passportNumber"`
    Inn            string `mapstructure:"inn_inner407"            json:"inn"`
    Snils          string `mapstructure:"snils_inner408"          json:"snils"`
}
type BkiConsentOut_d4e5f6a4 struct {
    Granted     bool   `mapstructure:"granted_inner421"     json:"granted"`
    DocumentRef string `mapstructure:"documentRef_inner422" json:"documentRef"`
}
type OutMessage_d4e5f6a4 struct {
    RequestID      string                `mapstructure:"requestId_inner431"      json:"requestId"`
    ApplicationID  string                `mapstructure:"applicationId_inner432"  json:"applicationId"`
    Client         BkiClientOut_d4e5f6a4 `mapstructure:"client_inner433"         json:"client"`
    RequestedFrom  []string              `mapstructure:"requestedFrom_inner434"  json:"requestedFrom"`
    RequestPurpose string                `mapstructure:"requestPurpose_inner435" json:"requestPurpose"`
    Consent        BkiConsentOut_d4e5f6a4 `mapstructure:"consent_inner436"       json:"consent"`
}

func Process_93d9fa01(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script (inner→out, ID d4e5f6a4)")
    if msg.Headers == nil { msg.Headers = make(map[string]string) }
    for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
        if _, ok := msg.Headers[key]; ok { continue }
        if msg.Data != nil { if v, ok := msg.Data[key].(string); ok && v != "" { msg.Headers[key] = v; delete(msg.Data, key) } }
    }
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil { log.Printf("❌ parse: %v", err); sendOrWarn(msg, output); return msg }
    var out OutMessage_d4e5f6a4
    if err := mapstructure.Decode(raw, &out); err != nil { log.Printf("❌ decode: %v", err); sendOrWarn(msg, output); return msg }
    newBody, err := json.Marshal(out)
    if err != nil { log.Printf("❌ marshal: %v", err); sendOrWarn(msg, output); return msg }
    msg.Body = newBody; msg.IsJSON = true; msg.Data = out
    log.Printf("   • Body: %s", string(msg.Body))
    sendOrWarn(msg, output); return msg
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