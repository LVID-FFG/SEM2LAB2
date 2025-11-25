#include <iostream>
#include <sstream>

using namespace std;

struct MChain{
    string key;
    string data;
    MChain* next;
    MChain(string key, string data) : key(key), data(data), next(nullptr){}
};

size_t genHashMorf(int size, string key){
    size_t result = 1245;
    for (size_t i = 0; i < key.size(); i++){
        result += (i * key[i]) % size;
    }
    return result % size;
}

struct HashTableMChain{
    MChain** table;
    size_t size;
    
    HashTableMChain(size_t sz) : size(sz){
        table = new MChain*[sz];
        for (size_t i = 0; i < size; i++) {
            table[i] = nullptr;
        }
    }
    
    ~HashTableMChain(){
        for (size_t i = 0; i < size; i++) {
            MChain* current = table[i];
            while (current != nullptr) {
                MChain* next = current->next;
                delete current;
                current = next;
            }
        }
        delete[] table;
    }

    void insert(string key, string value){
        size_t Hash = genHashMorf(size, key);

        if (table[Hash] == nullptr){
            table[Hash] = new MChain(key, value);
        }else{
            MChain* address = table[Hash];
            while (address->next != nullptr && address->key != key) 
                address = address->next;
            if (address->key == key) {
                address->data = value;
            } else {
                address->next = new MChain(key, value);
            }
        }
    }

    string find(string key){
        size_t Hash = genHashMorf(size, key);
        MChain* address = table[Hash];
        if (address == nullptr){
            return "";
        } 
        while (address != nullptr){
            if (address->key == key){
                return address->data;
            }
            address = address->next;
        }
        return "";
    }
};

struct morfTable{
    HashTableMChain tableIn;
    HashTableMChain tableOut;
    morfTable() : tableIn(10), tableOut(10){}
    
    bool searchSimbolMorf(string& str1, string& str2, string symbol) {
        string existingMapping = tableIn.find(symbol);
        
        if (!existingMapping.empty()) {
            // Проверяем корректность существующего отображения
            for (int i = 0; i < str1.size(); i++) {
                if (str1[i] == symbol[0] && str2[i] != existingMapping[0]) {
                    return false;
                }
            }
            return true;
        }
        
        // Ищем кандидата для отображения
        char candidate = 0;
        bool candidateFound = false;
        
        for (int i = 0; i < str1.size(); i++) {
            if (str1[i] == symbol[0]) {
                if (!candidateFound) {
                    candidate = str2[i];
                    candidateFound = true;
                } else if (str2[i] != candidate) {
                    return false;
                }
            }
        }
        
        // Проверяем что разные символы не отображаются в 1
        string candidateStr(1, candidate);
        string reverseMapping = tableOut.find(candidateStr);
        if (!reverseMapping.empty() && reverseMapping != symbol) {
            return false;
        }
        
        // Сохраняем отображения
        tableIn.insert(symbol, candidateStr);
        tableOut.insert(candidateStr, symbol);
        return true;
    }
};

void isMorf(string str1, string str2){
    if (str1.size() != str2.size()) {
        cout << "FALSE" << endl;
        return;
    }
    
    morfTable table;
    
    // Проверяем отображение из str1 в str2
    for (int i = 0; i < str1.size(); i++){
        string symbol(1, str1[i]);
        if (!table.searchSimbolMorf(str1, str2, symbol)){
            cout << "FALSE" << endl;
            return;
        }
    }
    
    cout << "TRUE" << endl;
}

void morf(){
    string str1;
    string str2;
    cout << "Введите строки" << endl;
    cin >> str1;
    cin >> str2;
    isMorf(str1, str2);
}