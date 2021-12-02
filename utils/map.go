package utils

func GetMap(randomNum int) string {
	maps := []string{"Woods", "Factory", "Shoreline", "Customs", "Reserve", "Interchange"}
	return maps[randomNum-1]
}
