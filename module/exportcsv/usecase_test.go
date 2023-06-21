package exportcsv

import (
	"testing"
)

func TestTransactionStringConverter(t *testing.T) {
	var positifData []Transaction
	positifData = append(positifData, Transaction{2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", false, false, false, false, 0.000000, "2023-04-17", false})
	positifData = append(positifData, Transaction{2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", false, false, false, false, 0.000000, "2023-04-17", false})
	positifData = append(positifData, Transaction{241, "5655429080696649", "9018506264648524", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", false, false, false, false, 0.000000, "2023-04-17", false})
	positifData = append(positifData, Transaction{242, "1262435180570130", "1798502071284385", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", false, false, false, false, 0.000000, "2023-04-17", false})

}
