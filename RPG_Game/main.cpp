#include<stdio.h>
#include<iostream>

using namespace std;

bool isKnightSleep, isPrisonerSleep, isArcherSleep, hasDog, isAllSleep;

bool CanFastAttack(){
    return isKnightSleep;
}

bool CanSpy(){
    return isAllSleep;
}

bool CanSignalPrisoner(){
    return isArcherSleep && !isKnightSleep;
}

bool CanFreePrisoner(){
    return isArcherSleep && isKnightSleep && !isPrisonerSleep && hasDog;
}

int main(){
    // isKnightSleep = isArcherSleep = isPrisonerSleep = hasDog = true;
    cout << "Enter the state of characters - Knight | Archer | Prisoner | Dog \n";
    cin >> isKnightSleep;
    cin >> isArcherSleep;
    cin >> isPrisonerSleep;
    cin >> isAllSleep;
    isAllSleep = isKnightSleep && isPrisonerSleep && isArcherSleep;
    cout<<isKnightSleep<<isArcherSleep<<isPrisonerSleep<<isAllSleep<<hasDog<<endl;
    cout<<CanFastAttack()<<endl;
    cout<<CanSpy()<<endl;
    cout<<CanSignalPrisoner()<<endl;
    cout<<CanFreePrisoner()<<endl;

    return 0;
}