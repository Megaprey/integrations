package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

type ClientOut_b2c3d4e2 struct {
    FirstName  string               `mapstructure:"firstName_inner201"  json:"firstName"`
    LastName   string               `mapstructure:"lastName_inner202"   json:"lastName"`
    MiddleName string               `mapstructure:"middleName_inner203" json:"middleName"`
    BirthDate  string               `mapstructure:"birthDate_inner204"  json:"birthDate"`
    Inn        string               `mapstructure:"inn_inner205"        json:"inn"`
    Snils      string               `mapstructure:"snils_inner206"      json:"snils"`
    Passport   PassportOut_b2c3d4e2 `mapstructure:"passport_inner207"   json:"passport"`
}
type PassportOut_b2c3d4e2 struct {
    Series string `mapstructure:"series_inner141" json:"series"`
    Number string `mapstructure:"number_inner142" json:"number"`
}
type BiometryOut_b2c3d4e2 struct {
    SelfieURL        string `mapstructure:"selfieUrl_inner211"        json:"selfieUrl"`
    LivenessVideoURL string `mapstructure:"livenessVideoUrl_inner212" json:"livenessVideoUrl"`
}
type OutMessage_b2c3d4e2 struct {
    RequestID     string             `mapstructure:"requestId_inner221"     json:"requestId"`
    ApplicationID string             `mapstructure:"applicationId_inner222" json:"applicationId"`
    Client        ClientOut_b2c3d4e2 `mapstructure:"client_inner223"        json:"client"`
    Checks        []string           `mapstructure:"checks_inner224"        json:"checks"`
    Biometry      BiometryOut_b2c3d4e2 `mapstructure:"biometry_inner225"    json:"biometry"`
}

func Process_2f8be8bc(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script (inner→out, ID b2c3d4e2)")
    if msg.Headers == nil { msg.Headers = make(map[string]string) }
    for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
        if _, ok := msg.Headers[key]; ok { continue }
        if msg.Data != nil { if v, ok := msg.Data[key].(string); ok && v != "" { msg.Headers[key] = v; delete(msg.Data, key) } }
    }
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil { log.Printf("❌ parse: %v", err); sendOrWarn(msg, output); return msg }
    var out OutMessage_b2c3d4e2
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