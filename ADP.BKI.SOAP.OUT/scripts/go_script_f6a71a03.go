package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

// === Скрипт: XML → JSON для БКИ ===
func Process_f6a71a03(msg Message, sender MessageSender) error {
	log.Printf("🔄 XML → JSON (BKI): обработка %s", msg.ID)

	result, err := parseSoapBody(msg.Body, "BkiReportResponse")
	if err != nil {
		log.Printf("❌ Ошибка парсинга XML: %v", err)
		return err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		log.Printf("❌ Ошибка сериализации JSON: %v", err)
		return err
	}

	outMsg := Message{
		ID:            msg.ID,
		Body:          jsonBytes,
		Data:          result,
		IsJSON:        true,
		Headers:       msg.Headers,
		AggregateName: msg.AggregateName,
		AggregateID:   msg.AggregateID,
		CorrelationID: msg.CorrelationID,
		ReplyTo:       msg.ReplyTo,
	}

	log.Printf("✅ XML → JSON готов (BKI): %s", string(jsonBytes))

	if sender != nil {
		if err := sender.Send(outMsg); err != nil {
			log.Printf("❌ Ошибка отправки: %v", err)
			return err
		}
	}
	return nil
}

// parseSoapBody — вытаскивает указанный элемент из <soap:Body> и превращает в map
func parseSoapBody(data []byte, targetElement string) (map[string]interface{}, error) {
	decoder := xml.NewDecoder(strings.NewReader(string(data)))
	inBody := false

	for {
		tok, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		switch t := tok.(type) {
		case xml.StartElement:
			if t.Name.Local == "Body" {
				inBody = true
			} else if inBody && t.Name.Local == targetElement {
				val, err := decodeElement(decoder, &t)
				if err != nil {
					return nil, err
				}
				if m, ok := val.(map[string]interface{}); ok {
					return m, nil
				}
				return map[string]interface{}{targetElement: val}, nil
			}
		}
	}
	return nil, fmt.Errorf("элемент <%s> не найден", targetElement)
}

// decodeElement — рекурсивно возвращает значение (string/float64/bool/map/[]interface{})
func decodeElement(decoder *xml.Decoder, start *xml.StartElement) (interface{}, error) {
	var text strings.Builder
	children := make(map[string]interface{})
	hasChildren := false

	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		switch t := tok.(type) {
		case xml.StartElement:
			hasChildren = true
			childVal, err := decodeElement(decoder, &t)
			if err != nil {
				return nil, err
			}
			addToMap(children, t.Name.Local, childVal)
		case xml.CharData:
			text.Write(t)
		case xml.EndElement:
			if t.Name.Local == start.Name.Local {
				if hasChildren {
					// «Поднимаем» единственного ребёнка-массив для известных полей
					if shouldFlatten(start.Name.Local, children) {
						for _, v := range children {
							if arr, ok := v.([]interface{}); ok {
								return arr, nil
							}
						}
					}
					return children, nil
				}
				s := strings.TrimSpace(text.String())
				if s == "" {
					return nil, nil
				}
				// 1. Сначала число (приоритетнее, чтобы "0" → 0, а не false)
				if n, err := strconv.ParseFloat(s, 64); err == nil {
					return n, nil
				}
				// 2. Потом bool ("true"/"false")
				if b, err := strconv.ParseBool(s); err == nil {
					return b, nil
				}
				// 3. Иначе — строка
				return s, nil
			}
		}
	}
}

// shouldFlatten — «поднимать» ли единственного ребёнка-массив
func shouldFlatten(parentName string, children map[string]interface{}) bool {
	flattenParents := map[string]bool{
		"sources":     true,
		"documents":   true,
		"consents":    true,
		"items":       true,
		"checks":      true,
		"flags":       true,
		"reasonCodes": true,
		"proceedings": true,
	}
	if !flattenParents[parentName] {
		return false
	}
	if len(children) != 1 {
		return false
	}
	for _, v := range children {
		if _, ok := v.([]interface{}); ok {
			return true
		}
	}
	return false
}

// addToMap — добавляет значение, повторяющиеся ключи → массив
func addToMap(m map[string]interface{}, key string, value interface{}) {
	if existing, ok := m[key]; ok {
		if arr, ok := existing.([]interface{}); ok {
			m[key] = append(arr, value)
		} else {
			m[key] = []interface{}{existing, value}
		}
	} else {
		m[key] = value
	}
}
