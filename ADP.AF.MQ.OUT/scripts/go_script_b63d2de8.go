package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

type AfClientOut_f6a7b8c6 struct {
    ClientID  string               `mapstructure:"clientId_inner501"  json:"clientId"`
    FirstName string               `mapstructure:"firstName_inner502" json:"firstName"`
    LastName  string               `mapstructure:"lastName_inner503"  json:"lastName"`
    BirthDate string               `mapstructure:"birthDate_inner504" json:"birthDate"`
    Inn       string               `mapstructure:"inn_inner505"       json:"inn"`
    Phone     string               `mapstructure:"phone_inner506"     json:"phone"`
    Passport  PassportOut_f6a7b8c6 `mapstructure:"passport_inner507"  json:"passport"`
}
type PassportOut_f6a7b8c6 struct {
    Series string `mapstructure:"series_inner141" json:"series"`
    Number string `mapstructure:"number_inner142" json:"number"`
}
type AfEmployeeOut_f6a7b8c6 struct {
    EmployeeID string `mapstructure:"employeeId_inner511" json:"employeeId"`
    BranchCode string `mapstructure:"branchCode_inner512" json:"branchCode"`
}
type OutMessage_f6a7b8c6 struct {
    RequestID     string               `mapstructure:"requestId_inner521"     json:"requestId"`
    ApplicationID string               `mapstructure:"applicationId_inner522" json:"applicationId"`
    Client        AfClientOut_f6a7b8c6  `mapstructure:"client_inner523"        json:"client"`
    Channel       string               `mapstructure:"channel_inner524"       json:"channel"`
    Employee      AfEmployeeOut_f6a7b8c6 `mapstructure:"employee_inner525"     json:"employee"`
    Checks        []string              `mapstructure:"checks_inner526"        json:"checks"`
}

func Process_b63d2de8(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script (inner→out, ID f6a7b8c6)")
    if msg.Headers == nil { msg.Headers = make(map[string]string) }
    for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
        if _, ok := msg.Headers[key]; ok { continue }
        if msg.Data != nil { if v, ok := msg.Data[key].(string); ok && v != "" { msg.Headers[key] = v; delete(msg.Data, key) } }
    }
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil { log.Printf("❌ parse: %v", err); sendOrWarn(msg, output); return msg }
    var out OutMessage_f6a7b8c6
    if err := mapstructure.Decode(raw, &out); err != nil { log.Printf("❌ decode: %v", err); sendOrWarn(msg, output); return msg }
    newBody, err := json.Marshal(out)
    if err != nil { log.Printf("❌ marshal: %v", err); sendOrWarn(msg, output); return msg }
    msg.Body = newBody; msg.IsJSON = true; msg.Data = out
    log.Printf("   • Body: %s", string(msg.Body))
    sendOrWarn(msg, output); return msg
}

func sendOrWarn_b63d2de8(msg Message, output MessageSender) {
	if output != nil {
		if err := output.Send(msg); err != nil {
			log.Printf("❌ Ошибка отправки: %v", err)
		}
	} else {
		log.Printf("⚠️ output sender не установлен")
	}
}