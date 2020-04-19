#include "hamming.h"

int hamming::compute(string a, string b){
    if (a.length() != b.length()){
		throw std::domain_error("Unequal size");
	}

	int dist = 0;
	for(int i=0; i<a.length(); i++){
		if (a[i] != b[i]){
			dist = dist + 1;
		}
	}
	return dist;
}
