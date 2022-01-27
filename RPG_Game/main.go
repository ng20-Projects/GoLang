package main

var isKnightSleep, isPrisonerSleep, isArcherSleep, isAllSleep, hasDog bool

func CanFastAttack() bool {
	return isKnightSleep
}

func CanSpy() bool {
	return isAllSleep
}

func CanSignalPrisoner() bool {
	return !isPrisonerSleep && isArcherSleep
}

func CanFreePrisoner() bool {
	return ((hasDog && isArcherSleep) || (!isPrisonerSleep && isArcherSleep && isKnightSleep))
}

func main() {
	isArcherSleep, isKnightSleep, isPrisonerSleep = true, true, true
	isAllSleep = isPrisonerSleep && isArcherSleep && isKnightSleep
}
