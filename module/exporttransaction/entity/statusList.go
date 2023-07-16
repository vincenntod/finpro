package entity

const StatusSize = 4

var StatusArray = [StatusSize]string{"SUCCESS", "WAITING_FOR_DEBITTED", "FAILD", ""}

func IsValidStatus(status string) bool {
	for _, sts := range StatusArray {
		if sts == status {
			return true
		}
	}
	return false
}
