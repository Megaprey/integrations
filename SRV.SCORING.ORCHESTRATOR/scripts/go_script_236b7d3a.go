package main

import (
    "encoding/json"
    "log"

    "github.com/go-viper/mapstructure/v2"
)

// =============================================================================
// ВХОДНЫЕ СТРУКТУРЫ (inner, читаем из canonical inner)
// =============================================================================

type ProductIn struct {
    Type       string `mapstructure:"type_inner101"       json:"type_inner101"`
    Amount     int    `mapstructure:"amount_inner102"     json:"amount_inner102"`
    Currency   string `mapstructure:"currency_inner103"   json:"currency_inner103"`
    TermMonths int    `mapstructure:"termMonths_inner104" json:"termMonths_inner104"`
}

type EmploymentIn struct {
    Status        string `mapstructure:"status_inner111"        json:"status_inner111"`
    EmployerName  string `mapstructure:"employerName_inner112"  json:"employerName_inner112"`
    MonthlyIncome int    `mapstructure:"monthlyIncome_inner113" json:"monthlyIncome_inner113"`
}

type PassportIn struct {
    Series string `mapstructure:"series_inner141" json:"series_inner141"`
    Number string `mapstructure:"number_inner142" json:"number_inner142"`
}

type ClientIn struct {
    ClientID   string       `mapstructure:"clientId_inner121"   json:"clientId_inner121"`
    FirstName  string       `mapstructure:"firstName_inner122"  json:"firstName_inner122"`
    LastName   string       `mapstructure:"lastName_inner123"   json:"lastName_inner123"`
    MiddleName string       `mapstructure:"middleName_inner124" json:"middleName_inner124"`
    BirthDate  string       `mapstructure:"birthDate_inner125"  json:"birthDate_inner125"`
    Inn        string       `mapstructure:"inn_inner126"        json:"inn_inner126"`
    Snils      string       `mapstructure:"snils_inner127"      json:"snils_inner127"`
    Passport   PassportIn   `mapstructure:"passport_inner128"   json:"passport_inner128"`
    Phone      string       `mapstructure:"phone_inner129"      json:"phone_inner129"`
    Employment EmploymentIn `mapstructure:"employment_inner130" json:"employment_inner130"`
}

type ConsentIn struct {
    Type        string `mapstructure:"type_inner151"        json:"type_inner151"`
    Granted     bool   `mapstructure:"granted_inner152"     json:"granted_inner152"`
    DocumentRef string `mapstructure:"documentRef_inner153" json:"documentRef_inner153"`
}

type DocumentIn struct {
    Type string `mapstructure:"type_inner161" json:"type_inner161"`
    URL  string `mapstructure:"url_inner162"  json:"url_inner162"`
}

type ApplicationPayload struct {
    CrmRequestID string        `mapstructure:"crmRequestId_inner171" json:"crmRequestId_inner171"`
    Channel      string        `mapstructure:"channel_inner173"      json:"channel_inner173"`
    BranchCode   string        `mapstructure:"branchCode_inner174"   json:"branchCode_inner174"`
    EmployeeID   string        `mapstructure:"employeeId_inner175"   json:"employeeId_inner175"`
    Product      ProductIn     `mapstructure:"product_inner176"      json:"product_inner176"`
    Client       ClientIn      `mapstructure:"client_inner177"       json:"client_inner177"`
    Consents     []ConsentIn   `mapstructure:"consents_inner178"     json:"consents_inner178"`
    Documents    []DocumentIn  `mapstructure:"documents_inner179"    json:"documents_inner179"`
}

type EnvelopeIn struct {
    MessageID     string `mapstructure:"messageId_inner2001"     json:"messageId_inner2001"`
    CorrelationID string `mapstructure:"correlationId_inner2002" json:"correlationId_inner2002"`
    ApplicationID string `mapstructure:"applicationId_inner2003" json:"applicationId_inner2003"`
    MessageType   string `mapstructure:"messageType_inner2004"   json:"messageType_inner2004"`
    SourceSystem  string `mapstructure:"sourceSystem_inner2005"  json:"sourceSystem_inner2005"`
    Timestamp     string `mapstructure:"timestamp_inner2006"     json:"timestamp_inner2006"`
}

type IncomingMessage struct {
    Envelope EnvelopeIn         `mapstructure:"envelope"          json:"envelope"`
    Payload  ApplicationPayload `mapstructure:"payload_inner2007" json:"payload_inner2007"`
}

// =============================================================================
// ИСХОДЯЩИЕ СТРУКТУРЫ (inner, пишем в canonical inner для OUT-адаптеров)
// =============================================================================

// ---------- 2.1 KYC ----------
type PassportKycOut struct {
    Series string `mapstructure:"series_inner141" json:"series_inner141"`
    Number string `mapstructure:"number_inner142" json:"number_inner142"`
}
type ClientKycOut struct {
    FirstName  string          `mapstructure:"firstName_inner201"  json:"firstName_inner201"`
    LastName   string          `mapstructure:"lastName_inner202"   json:"lastName_inner202"`
    MiddleName string          `mapstructure:"middleName_inner203" json:"middleName_inner203"`
    BirthDate  string          `mapstructure:"birthDate_inner204"  json:"birthDate_inner204"`
    Inn        string          `mapstructure:"inn_inner205"        json:"inn_inner205"`
    Snils      string          `mapstructure:"snils_inner206"      json:"snils_inner206"`
    Passport   PassportKycOut  `mapstructure:"passport_inner207"   json:"passport_inner207"`
}
type BiometryKycOut struct {
    SelfieURL        string `mapstructure:"selfieUrl_inner211"        json:"selfieUrl_inner211"`
    LivenessVideoURL string `mapstructure:"livenessVideoUrl_inner212" json:"livenessVideoUrl_inner212"`
}
type KycVerifyRequestOut struct {
    RequestID     string           `mapstructure:"requestId_inner221"     json:"requestId_inner221"`
    ApplicationID string           `mapstructure:"applicationId_inner222" json:"applicationId_inner222"`
    Client        ClientKycOut     `mapstructure:"client_inner223"        json:"client_inner223"`
    Checks        []string         `mapstructure:"checks_inner224"        json:"checks_inner224"`
    Biometry      BiometryKycOut   `mapstructure:"biometry_inner225"      json:"biometry_inner225"`
}

// ---------- 2.2 БКИ ----------
type BkiClientOut struct {
    FirstName      string `mapstructure:"firstName_inner401"      json:"firstName_inner401"`
    LastName       string `mapstructure:"lastName_inner402"       json:"lastName_inner402"`
    MiddleName     string `mapstructure:"middleName_inner403"     json:"middleName_inner403"`
    BirthDate      string `mapstructure:"birthDate_inner404"      json:"birthDate_inner404"`
    PassportSeries string `mapstructure:"passportSeries_inner405" json:"passportSeries_inner405"`
    PassportNumber string `mapstructure:"passportNumber_inner406" json:"passportNumber_inner406"`
    Inn            string `mapstructure:"inn_inner407"            json:"inn_inner407"`
    Snils          string `mapstructure:"snils_inner408"          json:"snils_inner408"`
}
type BkiConsentOut struct {
    Granted     bool   `mapstructure:"granted_inner421"     json:"granted_inner421"`
    DocumentRef string `mapstructure:"documentRef_inner422" json:"documentRef_inner422"`
}
type BkiReportRequestOut struct {
    RequestID      string         `mapstructure:"requestId_inner431"      json:"requestId_inner431"`
    ApplicationID  string         `mapstructure:"applicationId_inner432"  json:"applicationId_inner432"`
    Client         BkiClientOut   `mapstructure:"client_inner433"         json:"client_inner433"`
    RequestedFrom  []string       `mapstructure:"requestedFrom_inner434"  json:"requestedFrom_inner434"`
    RequestPurpose string         `mapstructure:"requestPurpose_inner435" json:"requestPurpose_inner435"`
    Consent        BkiConsentOut  `mapstructure:"consent_inner436"        json:"consent_inner436"`
}

// ---------- 2.3 Anti-Fraud ----------
type PassportAfOut struct {
    Series string `mapstructure:"series_inner141" json:"series_inner141"`
    Number string `mapstructure:"number_inner142" json:"number_inner142"`
}
type AfClientOut struct {
    ClientID  string         `mapstructure:"clientId_inner501"  json:"clientId_inner501"`
    FirstName string         `mapstructure:"firstName_inner502" json:"firstName_inner502"`
    LastName  string         `mapstructure:"lastName_inner503"  json:"lastName_inner503"`
    BirthDate string         `mapstructure:"birthDate_inner504" json:"birthDate_inner504"`
    Inn       string         `mapstructure:"inn_inner505"       json:"inn_inner505"`
    Phone     string         `mapstructure:"phone_inner506"     json:"phone_inner506"`
    Passport  PassportAfOut  `mapstructure:"passport_inner507"  json:"passport_inner507"`
}
type AfEmployeeOut struct {
    EmployeeID string `mapstructure:"employeeId_inner511" json:"employeeId_inner511"`
    BranchCode string `mapstructure:"branchCode_inner512" json:"branchCode_inner512"`
}
type FraudCheckRequestOut struct {
    RequestID     string          `mapstructure:"requestId_inner521"     json:"requestId_inner521"`
    ApplicationID string          `mapstructure:"applicationId_inner522" json:"applicationId_inner522"`
    Client        AfClientOut     `mapstructure:"client_inner523"        json:"client_inner523"`
    Channel       string          `mapstructure:"channel_inner524"       json:"channel_inner524"`
    Employee      AfEmployeeOut   `mapstructure:"employee_inner525"      json:"employee_inner525"`
    Checks        []string        `mapstructure:"checks_inner526"        json:"checks_inner526"`
}

// ---------- 2.4 Feature Store ----------
type FsBkiInputOut struct {
    Score          int `mapstructure:"score_inner601"          json:"score_inner601"`
    ActiveLoans    int `mapstructure:"activeLoans_inner602"    json:"activeLoans_inner602"`
    TotalDebt      int `mapstructure:"totalDebt_inner603"      json:"totalDebt_inner603"`
    OverdueMaxDays int `mapstructure:"overdueMaxDays_inner604" json:"overdueMaxDays_inner604"`
}
type FsAntifraudInputOut struct {
    RiskScore int    `mapstructure:"riskScore_inner611" json:"riskScore_inner611"`
    RiskLevel string `mapstructure:"riskLevel_inner612" json:"riskLevel_inner612"`
}
type FsInternalInputOut struct {
    ClientSince       string `mapstructure:"clientSince_inner621"       json:"clientSince_inner621"`
    AvgMonthlyBalance int    `mapstructure:"avgMonthlyBalance_inner622" json:"avgMonthlyBalance_inner622"`
    SalaryClient      bool   `mapstructure:"salaryClient_inner623"      json:"salaryClient_inner623"`
}
type FsProductInputOut struct {
    Amount     int    `mapstructure:"amount_inner631"     json:"amount_inner631"`
    TermMonths int    `mapstructure:"termMonths_inner632" json:"termMonths_inner632"`
    Type       string `mapstructure:"type_inner633"       json:"type_inner633"`
}
type FsInputsOut struct {
    Kyc       string                 `mapstructure:"kyc_inner641"       json:"kyc_inner641"`
    Bki       FsBkiInputOut          `mapstructure:"bki_inner642"       json:"bki_inner642"`
    Antifraud FsAntifraudInputOut    `mapstructure:"antifraud_inner643" json:"antifraud_inner643"`
    Internal  FsInternalInputOut     `mapstructure:"internal_inner644"  json:"internal_inner644"`
    Product   FsProductInputOut      `mapstructure:"product_inner645"   json:"product_inner645"`
}
type FeatureComputeRequestOut struct {
    RequestID     string         `mapstructure:"requestId_inner651"     json:"requestId_inner651"`
    ApplicationID string         `mapstructure:"applicationId_inner652" json:"applicationId_inner652"`
    ClientID      string         `mapstructure:"clientId_inner653"      json:"clientId_inner653"`
    FeatureGroups []string       `mapstructure:"featureGroups_inner654" json:"featureGroups_inner654"`
    Inputs        FsInputsOut    `mapstructure:"inputs_inner655"        json:"inputs_inner655"`
}

// ---------- 2.5 Scoring ----------
type ScoringFeaturesOut struct {
    Age                int     `mapstructure:"age_inner801"                json:"age_inner801"`
    IncomeMonthly      int     `mapstructure:"incomeMonthly_inner802"      json:"incomeMonthly_inner802"`
    Dti                float64 `mapstructure:"dti_inner803"                json:"dti_inner803"`
    BkiScore           int     `mapstructure:"bkiScore_inner804"           json:"bkiScore_inner804"`
    ActiveLoans        int     `mapstructure:"activeLoans_inner805"        json:"activeLoans_inner805"`
    OverdueMaxDays     int     `mapstructure:"overdueMaxDays_inner806"     json:"overdueMaxDays_inner806"`
    AntifraudRiskScore int     `mapstructure:"antifraudRiskScore_inner807" json:"antifraudRiskScore_inner807"`
}
type ScoringRequestOut struct {
    RequestID       string             `mapstructure:"requestId_inner811"       json:"requestId_inner811"`
    ApplicationID   string             `mapstructure:"applicationId_inner812"   json:"applicationId_inner812"`
    ModelHint       string             `mapstructure:"modelHint_inner813"       json:"modelHint_inner813"`
    FeatureVectorID string             `mapstructure:"featureVectorId_inner814" json:"featureVectorId_inner814"`
    Features        ScoringFeaturesOut `mapstructure:"features_inner815"       json:"features_inner815"`
    Context         map[string]string  `mapstructure:"context_inner816"         json:"context_inner816"`
}

// ---------- 2.6 Rules ----------
type RulesScoringOut struct {
    Score     int     `mapstructure:"score_inner901"     json:"score_inner901"`
    Pd        float64 `mapstructure:"pd_inner902"        json:"pd_inner902"`
    RiskGrade string  `mapstructure:"riskGrade_inner903" json:"riskGrade_inner903"`
}
type RulesContextOut struct {
    Scoring   RulesScoringOut        `mapstructure:"scoring_inner911"   json:"scoring_inner911"`
    Bki       map[string]interface{} `mapstructure:"bki_inner912"       json:"bki_inner912"`
    Client    map[string]interface{} `mapstructure:"client_inner913"    json:"client_inner913"`
    Antifraud map[string]interface{} `mapstructure:"antifraud_inner914" json:"antifraud_inner914"`
    Product   map[string]interface{} `mapstructure:"product_inner915"   json:"product_inner915"`
}
type RulesEvaluateRequestOut struct {
    RequestID     string            `mapstructure:"requestId_inner921"     json:"requestId_inner921"`
    ApplicationID string            `mapstructure:"applicationId_inner922" json:"applicationId_inner922"`
    RuleSet       string            `mapstructure:"ruleSet_inner923"       json:"ruleSet_inner923"`
    Context       RulesContextOut   `mapstructure:"context_inner924"       json:"context_inner924"`
}

// ---------- 2.7 Decision Engine ----------
type DecProductOut struct {
    Type       string `mapstructure:"type_inner1001"       json:"type_inner1001"`
    Amount     int    `mapstructure:"amount_inner1002"     json:"amount_inner1002"`
    Currency   string `mapstructure:"currency_inner1003"   json:"currency_inner1003"`
    TermMonths int    `mapstructure:"termMonths_inner1004" json:"termMonths_inner1004"`
}
type DecAggregatedOut struct {
    Kyc       map[string]interface{} `mapstructure:"kyc_inner1011"       json:"kyc_inner1011"`
    Bki       map[string]interface{} `mapstructure:"bki_inner1012"       json:"bki_inner1012"`
    Antifraud map[string]interface{} `mapstructure:"antifraud_inner1013" json:"antifraud_inner1013"`
    Scoring   map[string]interface{} `mapstructure:"scoring_inner1014"   json:"scoring_inner1014"`
    Rules     map[string]interface{} `mapstructure:"rules_inner1015"     json:"rules_inner1015"`
}
type DecisionRequestOut struct {
    RequestID         string              `mapstructure:"requestId_inner1021"         json:"requestId_inner1021"`
    ApplicationID     string              `mapstructure:"applicationId_inner1022"     json:"applicationId_inner1022"`
    Product           DecProductOut       `mapstructure:"product_inner1023"           json:"product_inner1023"`
    AggregatedResults DecAggregatedOut    `mapstructure:"aggregatedResults_inner1024" json:"aggregatedResults_inner1024"`
}

// ---------- 2.8 Pricing ----------
type PricingClientOut struct {
    ClientID     string `mapstructure:"clientId_inner1101"     json:"clientId_inner1101"`
    RiskGrade    string `mapstructure:"riskGrade_inner1102"    json:"riskGrade_inner1102"`
    SalaryClient bool   `mapstructure:"salaryClient_inner1103" json:"salaryClient_inner1103"`
}
type PricingTermsOut struct {
    Amount     int    `mapstructure:"amount_inner1111"     json:"amount_inner1111"`
    Currency   string `mapstructure:"currency_inner1112"   json:"currency_inner1112"`
    TermMonths int    `mapstructure:"termMonths_inner1113" json:"termMonths_inner1113"`
}
type PricingScoringOut struct {
    Pd    float64 `mapstructure:"pd_inner1121"    json:"pd_inner1121"`
    Score int     `mapstructure:"score_inner1122" json:"score_inner1122"`
}
type PricingRequestOut struct {
    RequestID     string                `mapstructure:"requestId_inner1131"     json:"requestId_inner1131"`
    ApplicationID string                `mapstructure:"applicationId_inner1132" json:"applicationId_inner1132"`
    Product       string                `mapstructure:"product_inner1133"       json:"product_inner1133"`
    Client        PricingClientOut      `mapstructure:"client_inner1134"        json:"client_inner1134"`
    Terms         PricingTermsOut       `mapstructure:"terms_inner1135"         json:"terms_inner1135"`
    Scoring       PricingScoringOut     `mapstructure:"scoring_inner1136"       json:"scoring_inner1136"`
}

// =============================================================================
// ЕДИНЫЙ МЕТОД ОРКЕСТРАТОРА (inner → inner)
// =============================================================================

func Process_236b7d3a(msg Message, output MessageSender) Message {
    log.Printf("🐹 Orchestrator: входящее ID=%s", msg.ID)

    // --- 1. Перенос заголовков ---
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

    // --- 2. Парсинг входящего сообщения через mapstructure (inner → inner) ---
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil {
        log.Printf("❌ Ошибка парсинга JSON: %v", err)
        return msg
    }

    var incoming IncomingMessage
    if err := mapstructure.Decode(raw, &incoming); err != nil {
        log.Printf("❌ Ошибка mapstructure.Decode: %v", err)
        return msg
    }

    // --- 3. Проверка типа сообщения ---
    if incoming.Envelope.MessageType != "ApplicationSubmitted" {
        log.Printf("⚠️ Не ApplicationSubmitted, пропуск: %q", incoming.Envelope.MessageType)
        return msg
    }

    app := incoming.Payload
    env := incoming.Envelope

    // --- 4. Формируем и отправляем все 8 исходящих ---

    // 2.1 KYC
    kycReq := KycVerifyRequestOut{
        RequestID:     "req-kyc-" + env.ApplicationID,
        ApplicationID: env.ApplicationID,
        Client: ClientKycOut{
            FirstName:  app.Client.FirstName,
            LastName:   app.Client.LastName,
            MiddleName: app.Client.MiddleName,
            BirthDate:  app.Client.BirthDate,
            Inn:        app.Client.Inn,
            Snils:      app.Client.Snils,
            Passport: PassportKycOut{
                Series: app.Client.Passport.Series,
                Number: app.Client.Passport.Number,
            },
        },
        Checks: []string{"PASSPORT_VALIDITY", "INN_SNILS_MATCH", "FNS_CHECK", "MVD_CHECK", "FSSP_CHECK", "BANKRUPTCY_CHECK"},
        Biometry: BiometryKycOut{
            SelfieURL:        "S3://docs/selfie.jpg",
            LivenessVideoURL: "S3://docs/liveness.mp4",
        },
    }
    sendAsync("CRM", "KYC", "ADP.KYC.HTTP.OUT", env, kycReq, output)

    // 2.2 БКИ
    bkiReq := BkiReportRequestOut{
        RequestID:     "req-bki-" + env.ApplicationID,
        ApplicationID: env.ApplicationID,
        Client: BkiClientOut{
            FirstName:      app.Client.FirstName,
            LastName:       app.Client.LastName,
            MiddleName:     app.Client.MiddleName,
            BirthDate:      app.Client.BirthDate,
            PassportSeries: app.Client.Passport.Series,
            PassportNumber: app.Client.Passport.Number,
            Inn:            app.Client.Inn,
            Snils:          app.Client.Snils,
        },
        RequestedFrom:  []string{"NBKI", "OKB"},
        RequestPurpose: "CREDIT_APPLICATION",
        Consent: BkiConsentOut{
            Granted:     true,
            DocumentRef: "S3://consents/bki.pdf",
        },
    }
    sendAsync("CRM", "BKI", "ADP.BKI.SOAP.OUT", env, bkiReq, output)

    // 2.3 Anti-Fraud
    afReq := FraudCheckRequestOut{
        RequestID:     "req-af-" + env.ApplicationID,
        ApplicationID: env.ApplicationID,
        Client: AfClientOut{
            ClientID:  app.Client.ClientID,
            FirstName: app.Client.FirstName,
            LastName:  app.Client.LastName,
            BirthDate: app.Client.BirthDate,
            Inn:       app.Client.Inn,
            Phone:     app.Client.Phone,
            Passport: PassportAfOut{
                Series: app.Client.Passport.Series,
                Number: app.Client.Passport.Number,
            },
        },
        Channel: app.Channel,
        Employee: AfEmployeeOut{
            EmployeeID: app.EmployeeID,
            BranchCode: app.BranchCode,
        },
        Checks: []string{"BLACKLIST_PHONE", "BLACKLIST_PASSPORT", "DUPLICATE_APPLICATION", "EMPLOYEE_FRAUD_PATTERN"},
    }
    sendAsync("CRM", "ANTIFRAUD", "ADP.ANTIFRAUD.MQ.OUT", env, afReq, output)

    // 2.4 Feature Store
    fsReq := FeatureComputeRequestOut{
        RequestID:     "req-fs-" + env.ApplicationID,
        ApplicationID: env.ApplicationID,
        ClientID:      app.Client.ClientID,
        FeatureGroups: []string{"CLIENT_DEMOGRAPHICS", "BKI_AGGREGATES", "INTERNAL_BEHAVIOR", "PRODUCT_CONTEXT", "ANTIFRAUD_SIGNALS"},
        Inputs: FsInputsOut{
            Kyc: "VERIFIED",
            Bki: FsBkiInputOut{
                Score:          720,
                ActiveLoans:    2,
                TotalDebt:      180000,
                OverdueMaxDays: 0,
            },
            Antifraud: FsAntifraudInputOut{
                RiskScore: 12,
                RiskLevel: "LOW",
            },
            Internal: FsInternalInputOut{
                ClientSince:       "2018-06-01",
                AvgMonthlyBalance: 45000,
                SalaryClient:      true,
            },
            Product: FsProductInputOut{
                Amount:     app.Product.Amount,
                TermMonths: app.Product.TermMonths,
                Type:       app.Product.Type,
            },
        },
    }
    sendAsync("CRM", "FEATURE_STORE", "ADP.FEATURESTORE.HTTP.OUT", env, fsReq, output)

    // 2.5 Scoring
    scReq := ScoringRequestOut{
        RequestID:       "req-sc-" + env.ApplicationID,
        ApplicationID:   env.ApplicationID,
        ModelHint:       "CONSUMER_LOAN_PD_V3",
        FeatureVectorID: "fv-" + env.ApplicationID,
        Features: ScoringFeaturesOut{
            Age:                41,
            IncomeMonthly:      app.Client.Employment.MonthlyIncome,
            Dti:                0.35,
            BkiScore:           720,
            ActiveLoans:        2,
            OverdueMaxDays:     0,
            AntifraudRiskScore: 12,
        },
        Context: map[string]string{
            "product": app.Product.Type,
            "channel": app.Channel,
        },
    }
    sendAsync("CRM", "SCORING", "ADP.SCORING.MQ.OUT", env, scReq, output)

    // 2.6 Rules
    rulesReq := RulesEvaluateRequestOut{
        RequestID:     "req-rules-" + env.ApplicationID,
        ApplicationID: env.ApplicationID,
        RuleSet:       "CONSUMER_LOAN_POLICY_V5",
        Context: RulesContextOut{
            Scoring: RulesScoringOut{
                Score:     685,
                Pd:        0.042,
                RiskGrade: "B",
            },
            Bki: map[string]interface{}{
                "overdueMaxDays": 0,
                "activeLoans":    2,
            },
            Client: map[string]interface{}{
                "age":           41,
                "incomeMonthly": app.Client.Employment.MonthlyIncome,
                "dti":           0.35,
                "salaryClient":  true,
            },
            Antifraud: map[string]interface{}{
                "riskLevel": "LOW",
            },
            Product: map[string]interface{}{
                "amount":     app.Product.Amount,
                "termMonths": app.Product.TermMonths,
            },
        },
    }
    sendAsync("CRM", "RULES", "ADP.RULES.HTTP.OUT", env, rulesReq, output)

    // 2.7 Decision Engine
    decReq := DecisionRequestOut{
        RequestID:     "req-dec-" + env.ApplicationID,
        ApplicationID: env.ApplicationID,
        Product: DecProductOut{
            Type:       app.Product.Type,
            Amount:     app.Product.Amount,
            Currency:   app.Product.Currency,
            TermMonths: app.Product.TermMonths,
        },
        AggregatedResults: DecAggregatedOut{
            Kyc: map[string]interface{}{
                "status": "VERIFIED",
            },
            Bki: map[string]interface{}{
                "worstStatus":           "GOOD",
                "hasCurrentDelinquency": false,
                "score":                 720,
            },
            Antifraud: map[string]interface{}{
                "riskLevel": "LOW",
                "riskScore": 12,
            },
            Scoring: map[string]interface{}{
                "score":        685,
                "pd":           0.042,
                "riskGrade":    "B",
                "modelVersion": "3.2.1",
            },
            Rules: map[string]interface{}{
                "policyDecision": "PASS",
                "limits": map[string]interface{}{
                    "maxAmount":     750000,
                    "maxTermMonths": 60,
                },
            },
        },
    }
    sendAsync("CRM", "DECISION_ENGINE", "ADP.DECISION.MQ.OUT", env, decReq, output)

    // 2.8 Pricing
    priceReq := PricingRequestOut{
        RequestID:     "req-price-" + env.ApplicationID,
        ApplicationID: env.ApplicationID,
        Product:       app.Product.Type,
        Client: PricingClientOut{
            ClientID:     app.Client.ClientID,
            RiskGrade:    "B",
            SalaryClient: true,
        },
        Terms: PricingTermsOut{
            Amount:     app.Product.Amount,
            Currency:   app.Product.Currency,
            TermMonths: app.Product.TermMonths,
        },
        Scoring: PricingScoringOut{
            Pd:    0.042,
            Score: 685,
        },
    }
    sendAsync("CRM", "PRICING", "ADP.PRICING.HTTP.OUT", env, priceReq, output)

    log.Printf("✅ Все 8 исходящих сообщений отправлены асинхронно")
    return msg
}

// =============================================================================
// ВСПОМОГАТЕЛЬНАЯ ФУНКЦИЯ
// =============================================================================

// sendAsync маршалит inner-структуру в JSON с ключами _innerNNN
// и отправляет в OUT-адаптер асинхронно.
func sendAsync(sender, receiver, serviceName string, env EnvelopeIn, payload interface{}, output MessageSender) {
    body, err := json.Marshal(payload)
    if err != nil {
        log.Printf("❌ Ошибка маршалинга для %s: %v", receiver, err)
        return
    }

    var dataMap map[string]interface{}
    if err := json.Unmarshal(body, &dataMap); err != nil {
        log.Printf("❌ Ошибка приведения к map для %s: %v", receiver, err)
        dataMap = nil
    }

    outMsg := Message{
        ID:     env.MessageID + "-" + receiver,
        Body:   body,
        IsJSON: true,
        Data:   dataMap,
        Headers: map[string]string{
            "Sender":      sender,
            "Receiver":    receiver,
            "ServiceName": serviceName,
        },
    }

    if output != nil {
        if err := output.Send(outMsg); err != nil {
            log.Printf("❌ Ошибка отправки в %s: %v", receiver, err)
        } else {
            log.Printf("📤 Отправлено в %s (ServiceName=%s)", receiver, serviceName)
        }
    } else {
        log.Printf("⚠️ output sender не установлен, %s не отправлен", receiver)
    }
}