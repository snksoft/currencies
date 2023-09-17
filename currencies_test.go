package currencies

import (
	"testing"
)

func TestCurrencies(t *testing.T) {
	doTest := func(code string, number int, decimalPlaces int, scale int, description string) {
		currency := GetCurrencyByCode(code)
		if currency == nil {
			t.Fatalf("Currency not found for code %s", code)
		} else if currency.GetCode() != code || currency.GetNumber() != number || currency.GetDecimalPlaces() != decimalPlaces || currency.GetScale() != scale || currency.GetDescription() != description {
			t.Fatalf("Incorrect currency retrieved by code (code: '%s', number: %d, decimalPlaces: %d, scale: %d, description:'%s') when requested (code: '%s', number: %d, decimalPlaces: %d, scale: %d, description:'%s')",
				currency.GetCode(), currency.GetNumber(), currency.GetDecimalPlaces(), currency.GetScale(), currency.GetDescription(), code, number, decimalPlaces, scale, description)
		}

		currency = GetCurrencyByNumber(number)
		if currency == nil {
			t.Fatalf("Currency not found for number %d", number)
		} else if currency.GetCode() != code || currency.GetNumber() != number || currency.GetDecimalPlaces() != decimalPlaces || currency.GetScale() != scale || currency.GetDescription() != description {
			t.Fatalf("Incorrect currency retrieved by number (code: '%s', number: %d, decimalPlaces: %d, scale: %d, description:'%s') when requested (code: '%s', number: %d, decimalPlaces: %d, scale: %d, description:'%s')",
				currency.GetCode(), currency.GetNumber(), currency.GetDecimalPlaces(), currency.GetScale(), currency.GetDescription(), code, number, decimalPlaces, scale, description)
		}
	}

	doTest("AED", 784, 2, 100, "United Arab Emirates dirham")
	doTest("ALL", 8, 2, 100, "Albanian lek")
	doTest("AUD", 36, 2, 100, "Australian dollar")
	doTest("BHD", 48, 3, 1000, "Bahraini dinar")
	doTest("BIF", 108, 0, 1, "Burundian franc")
	doTest("BYR", 974, 0, 1, "Belarusian ruble")
	doTest("CAD", 124, 2, 100, "Canadian dollar")
	doTest("CLF", 990, 4, 10000, "Unidad de Fomento (funds code)")
	doTest("CNY", 156, 2, 100, "Chinese yuan")
	doTest("EUR", 978, 2, 100, "Euro")
	doTest("FJD", 242, 2, 100, "Fiji dollar")
	doTest("GNF", 324, 0, 1, "Guinean franc")
	doTest("HKD", 344, 2, 100, "Hong Kong dollar")
	doTest("ISK", 352, 0, 1, "Icelandic króna")
	doTest("KWD", 414, 3, 1000, "Kuwaiti dinar")
	doTest("MRO", 478, 1, 10, "Mauritanian ouguiya")
	doTest("NZD", 554, 2, 100, "New Zealand dollar")
	doTest("USD", 840, 2, 100, "United States dollar")

}

func TestNonExistingCurrencies(t *testing.T) {
	doTestCode := func(code string) {
		currency := GetCurrencyByCode(code)
		if currency != nil {
			t.Fatalf("OOPS! A non existing currency actually found on the list for code %s", code)
		}
	}
	doTestNumber := func(number int) {
		currency := GetCurrencyByNumber(number)
		if currency != nil {
			t.Fatalf("OOPS! A non existing currency actually found on the list for number %d", number)
		}
	}

	doTestCode("XYZ")
	doTestCode("")
	doTestCode("ABCDS")

	doTestNumber(0)
	doTestNumber(1000)
	doTestNumber(999) // "no currency"
	doTestNumber(964) // actually, Palladium, but it should not be on the list anyway
}
