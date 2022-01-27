#include<iostream>
#include<time.h>

int main(){
    clock_t t;
    t = clock();
    for(int i=0; i<10000; i++){
        continue;
    }
    t = clock() - t;
    double time_taken = ((double)t)/CLOCKS_PER_SEC; // in seconds
  
    printf("Time took %f seconds to execute \n", time_taken);

    return 0;
}