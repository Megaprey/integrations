package main

import (
    "encoding/json"
    "log"
    "github.com/go-viper/mapstructure/v2"
)

type StatusApprovedTermsOut_f8a9b0c8 struct {
    Amount         int     `mapstructure:"amount_inner1201"         json:"amount"`
    Currency       string  `mapstructure:"currency_inner1202"       json:"currency"`
    TermMonths     int     `mapstructure:"termMonths_inner1203"     json:"termMonths"`
    Rate           float64 `mapstructure:"rate_inner1204"           json:"rate"`
    MonthlyPayment int     `mapstructure:"monthlyPayment_inner1205" json:"monthlyPayment"`
    TotalCost      int     `mapstructure:"totalCost_inner1206"      json:"totalCost"`
}
type StatusSummaryOut_f8a9b0c8 struct {
    Kyc       string                 `mapstructure:"kyc_inner1211"       json:"kyc"`
    Bki       string                 `mapstructure:"bki_inner1212"       json:"bki"`
    Antifraud string                 `mapstructure:"antifraud_inner1213" json:"antifraud"`
    Scoring   map[string]interface{} `mapstructure:"scoring_inner1214"   json:"scoring"`
}
type StatusNextActionOut_f8a9b0c8 struct {
    Action   string `mapstructure:"action_inner1221"   json:"action"`
    Required bool   `mapstructure:"required_inner1222" json:"required"`
}
type OutMessage_f8a9b0c8 struct {
    NotificationID string                       `mapstructure:"notificationId_inner1231" json:"notificationId"`
    ApplicationID  string                       `mapstructure:"applicationId_inner1232"  json:"applicationId"`
    CrmRequestID   string                       `mapstructure:"crmRequestId_inner1233"   json:"crmRequestId"`
    Status         string                       `mapstructure:"status_inner1234"         json:"status"`
    Stage          string                       `mapstructure:"stage_inner1235"          json:"stage"`
    Decision       string                       `mapstructure:"decision_inner1236"       json:"decision"`
    DecidedAt      string                       `mapstructure:"decidedAt_inner1237"      json:"decidedAt"`
    ApprovedTerms  StatusApprovedTermsOut_f8a9b0c8 `mapstructure:"approvedTerms_inner1238" json:"approvedTerms"`
    Summary        StatusSummaryOut_f8a9b0c8      `mapstructure:"summary_inner1239"       json:"summary"`
    ReasonCodes    []string                     `mapstructure:"reasonCodes_inner1240"    json:"reasonCodes"`
    NextActions    []StatusNextActionOut_f8a9b0c8 `mapstructure:"nextActions_inner1241"   json:"nextActions"`
    Timestamp      string                       `mapstructure:"timestamp_inner1242"      json:"timestamp"`
}

func Process_f9208db8(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script (inner→out, ID f8a9b0c8)")
    if msg.Headers == nil { msg.Headers = make(map[string]string) }
    for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
        if _, ok := msg.Headers[key]; ok { continue }
        if msg.Data != nil { if v, ok := msg.Data[key].(string); ok && v != "" { msg.Headers[key] = v; delete(msg.Data, key) } }
    }
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil { log.Printf("❌ parse: %v", err); sendOrWarn(msg, output); return msg }
    var out OutMessage_f8a9b0c8
    if err := mapstructure.Decode(raw, &out); err != nil { log.Printf("❌ decode: %v", err); sendOrWarn(msg, output); return msg }
    newBody, err := json.Marshal(out)
    if err != nil { log.Printf("❌ marshal: %v", err); sendOrWarn(msg, output); return msg }
    msg.Body = newBody; msg.IsJSON = true; msg.Data = out
    log.Printf("   • Body: %s", string(msg.Body))
    sendOrWarn(msg, output); return msg
}

func sendOrWarn(msg Message, output MessageSender) {
    if output != nil { if err := output.Send(msg); err != nil { log.Printf("❌ send: %v", err) } } else { log.Printf("⚠️ output not set") }
}