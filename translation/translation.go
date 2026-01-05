package translation

import (
	"fmt"
	"os"
	"strings"

	"github.com/OkyHp/plg_utils/s2sdk"

	"gopkg.in/yaml.v3"
)

// ChatColors - цвета для чата (символы управления цветом)
var ChatColors = map[string]string{
	"Default":     "\x01",
	"White":       "\x01",
	"DarkRed":     "\x02",
	"Green":       "\x04",
	"LightYellow": "\x09",
	"LightBlue":   "\x0B",
	"Olive":       "\x05",
	"Lime":        "\x06",
	"Red":         "\x07",
	"LightPurple": "\x03",
	"Purple":      "\x0E",
	"Grey":        "\x08",
	"Yellow":      "\x09",
	"Gold":        "\x10",
	"Silver":      "\x0A",
	"Blue":        "\x0B",
	"DarkBlue":    "\x0C",
	"BlueGrey":    "\x0A",
	"Magenta":     "\x0E",
	"LightRed":    "\x0F",
	"Orange":      "\x10",
	"Darkred":     "\x02", // Obsolete, но оставляем для совместимости
	"NewLine":     "\u2029",
}

// translations - глобальное хранилище переводов
// map[ключ_перевода]map[язык]текст
var translations = make(map[string]map[string]string)

// LoadTranslation - загружает переводы из YAML файла
func LoadTranslation(dir, fileName string) error {
	pathToFile := fmt.Sprintf("%s/%s.yml", dir, fileName)

	// Читаем файл
	data, err := os.ReadFile(pathToFile)
	if err != nil {
		return fmt.Errorf("failed to read translation file: %w", err)
	}

	// Парсим YAML
	var yamlData map[string]map[string]string
	if err := yaml.Unmarshal(data, &yamlData); err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}

	// Сохраняем переводы глобально
	for key, translationsByLang := range yamlData {
		translations[key] = translationsByLang
	}

	return nil
}

// processPlaceholders - заменяет плейсхолдеры ChatColors на соответствующие символы
func processPlaceholders(input string) string {
	result := input
	for name, value := range ChatColors {
		placeholder := fmt.Sprintf("{%s}", name)
		result = strings.ReplaceAll(result, placeholder, value)
	}
	return result
}

// getTranslationText - получает текст перевода по ключу и языку
func getTranslationText(key, lang string) (string, bool) {
	if trans, ok := translations[key]; ok {
		if text, ok := trans[lang]; ok {
			return text, true
		}

		// Если запрошенный язык не найден, пытаемся вернуть английский
		if text, ok := trans["en"]; ok {
			return text, true
		}
	}

	return "", false
}

// processArgs - заменяет плейсхолдеры {0}, {1}, {2}... на аргументы
func processArgs(text string, args ...interface{}) string {
	if len(args) == 0 {
		return text
	}

	// Заменяем {0}, {1}, {2}... на %s для fmt.Sprintf
	result := text
	for i, arg := range args {
		placeholder := fmt.Sprintf("{%d}", i)
		result = strings.ReplaceAll(result, placeholder, fmt.Sprintf("%v", arg))
	}
	return result
}

// GetTranslation - получает перевод по ключу и языку
// Если запрошенный язык не найден, возвращает английский (en) по умолчанию
// Автоматически заменяет плейсхолдеры {ChatColors.*} на соответствующие символы
// func GetTranslation(key, lang string) string {
// 	text, found := getTranslationText(key, lang)
// 	if !found {
// 		return ""
// 	}
// 	return processPlaceholders(text)
// }

// GetTranslationF - получает перевод по ключу и языку с поддержкой аргументов
// Поддерживает плейсхолдеры {0}, {1}, {2}... для замены на аргументы
// Если запрошенный язык не найден, возвращает английский (en) по умолчанию
// Автоматически заменяет плейсхолдеры {ChatColors.*} на соответствующие символы
func GetTranslation(playerSlot int32, key string, args ...interface{}) string {
	lang := "en"
	if len(s2sdk.GetClientLanguage(playerSlot)) >= 2 {
		lang = s2sdk.GetClientLanguage(playerSlot)[:2]
	}

	text, found := getTranslationText(key, lang)
	if !found {
		return ""
	}

	// Сначала заменяем аргументы, потом цвета
	text = processArgs(text, args...)
	return processPlaceholders(text)
}
