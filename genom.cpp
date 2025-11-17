#include <iostream>
#include "setInclude.h"

using namespace std;


void genome(){
    string str1;
    string str2;
    StringSet genom1;
    StringSet genom2;
    cout << "Введите геномы" << endl;
    cin >> str1;
    cin >> str2;
    for (size_t i = 0; i < str1.size() - 1; i++){
        string addElem = "";
        addElem += str1[i];
        addElem += str1[i+1];
        genom1.add(addElem);
    }
    for (size_t i = 0; i < str2.size() - 1; i++){
        string addElem = "";
        addElem += str2[i];
        addElem += str2[i+1];
        genom2.add(addElem);
    }
    StringSet intersectionGenom = genom1.intersectionWith(genom2);
    cout << "Similarity = " << intersectionGenom.size() << endl;
}