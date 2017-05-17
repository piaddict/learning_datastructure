#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int lotto[6];

int DuCheck(int i) {
	int j;

	lotto[i] = rand()%45+1;
	for(j=0; j<i; j++) {
		if(lotto[i]==lotto[j]) DuCheck(i);
	}
	return lotto[i];
}

int main(void) {
	int i;
	srand(time(NULL));
	for(i=0 ; i<6 ; i++) {
		printf(" %d \n", DuCheck(i));
	}
	return 0;
}