package utils

func IsTypeUIShop(typeUI int8) bool {
	return typeUI == 20 || typeUI == 21 || typeUI == 22 || typeUI == 23 || typeUI == 24 || typeUI == 25 || typeUI == 26 || typeUI == 27 || typeUI == 28 || typeUI == 29 || typeUI == 16 || typeUI == 17 || typeUI == 18 || typeUI == 19 || typeUI == 2 || typeUI == 6 || typeUI == 8 || typeUI == 34
}

func IsTypeMounts(typeUI int8) bool {
	return typeUI == 29 && typeUI <= 33
}

func IsTypeUIShopLock(typeUI int8) bool {
	return typeUI == 7 || typeUI == 9
}

func IsTypeUIStore(typeUI int8) bool {
	return typeUI == 14
}

func IsTypeUIBook(typeUI int8) bool {
	return typeUI == 15
}

func IsTypeUIFashion(typeUI int8) bool {
	return typeUI == 32
}

func IsTypeUIClanShop(typeUI int8) bool {
	return typeUI == 34
}

func IsTypeUIME(typeUI int8) bool {
	return typeUI == 5 || typeUI == 3 || typeUI == 4 || typeUI == 39
}
