package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

type InnerFsFeatures_c9d0e1f9 struct {
    Age                int     `mapstructure:"age"                json:"age_inner701"`
    IncomeMonthly      int     `mapstructure:"incomeMonthly"      json:"incomeMonthly_inner702"`
    Dti                float64 `mapstructure:"dti"                json:"dti_inner703"`
    BkiScore           int     `mapstructure:"bkiScore"           json:"bkiScore_inner704"`
    ActiveLoans        int     `mapstructure:"activeLoans"        json:"activeLoans_inner705"`
    OverdueMaxDays     int     `mapstructure:"overdueMaxDays"     json:"overdueMaxDays_inner706"`
    ClientTenureMonths int     `mapstructure:"clientTenureMonths" json:"clientTenureMonths_inner707"`
    SalaryClient       bool    `mapstructure:"salaryClient"       json:"salaryClient_inner708"`
    AntifraudRiskScore int     `mapstructure:"antifraudRiskScore" json:"antifraudRiskScore_inner709"`
}
type InnerMessage_c9d0e1f9 struct {
    RequestID       string               `mapstructure:"requestId"       json:"requestId_inner711"`
    FeatureVectorID string               `mapstructure:"featureVectorId" json:"featureVectorId_inner712"`
    Features        InnerFsFeatures_c9d0e1f9 `mapstructure:"features"        json:"features_inner713"`
    ComputedAt      string               `mapstructure:"computedAt"      json:"computedAt_inner714"`
}

func Process_0f95d31a(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script (out→inner, ID c9d0e1f9)")
    if msg.Headers == nil { msg.Headers = make(map[string]string) }
    for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
        if _, ok := msg.Headers[key]; ok { continue }
        if msg.Data != nil { if v, ok := msg.Data[key].(string); ok && v != "" { msg.Headers[key] = v } }
    }
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil { log.Printf("❌ parse: %v", err); sendOrWarn(msg, output); return msg }
    var inner InnerMessage_c9d0e1f9
    if err := mapstructure.Decode(raw, &inner); err != nil { log.Printf("❌ decode: %v", err); sendOrWarn(msg, output); return msg }
    newBody, err := json.Marshal(inner)
    if err != nil { log.Printf("❌ marshal: %v", err); sendOrWarn(msg, output); return msg }
    msg.Body = newBody; msg.IsJSON = true; msg.Data = inner
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