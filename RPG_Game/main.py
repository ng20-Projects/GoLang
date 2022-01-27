isKnightSleep = isArcherSleep = isPrisonerSleep = isAllSleep = hasDog = True

def char_stat():
    isKnightSleep = bool(input("Knight State: "))
    isArcherSleep = bool(input("Archer State: "))
    isPrisonerSleep = bool(input("Prisoner State: "))
    hasDog = bool(input("Dog State: "))
    isAllSleep = bool(isKnightSleep and isArcherSleep and isPrisonerSleep)

def CanFastAttack():
    return isKnightSleep

def CanSpy():
    return isAllSleep

def CanSignalPrisoner():
    return isArcherSleep and (not isKnightSleep)

def CanFreePrisoner():
    return hasDog and isKnightSleep and isArcherSleep and (not isPrisonerSleep)

char_stat()
print("CanFastAttack: ", CanFastAttack())
print("CanSpy: ", CanSpy())
print("CanSignalPrisoner", CanSignalPrisoner())
print("CanFreePrisoner", CanFreePrisoner())

