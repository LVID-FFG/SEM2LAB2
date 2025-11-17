#include <iostream>
#include <algorithm>
#include <sstream>

using namespace std;

struct morfTable{
    pair<char, char>* table = new pair<char, char>[10];
    size_t size = 0;
    size_t cap = 10;

    bool searchSimbolMorf(string& str1, string& str2, char symbol){
        // Сначала проверяем, есть ли символ уже в таблице
        for(int i = 0; i < size; i++){
            if (table[i].first == symbol) {
                // Если символ есть в таблице, проверяем что отображение корректно
                char expectedChar = table[i].second;
                for(size_t j = 0; j < str1.size(); j++) {
                    if(str1[j] == symbol && str2[j] != expectedChar) {
                        return false;
                    }
                }
                return true;
            }
        }
        
        // Если символа нет в таблице, находим соответствующий символ из str2
        char mappedChar = '\0';
        for(size_t i = 0; i < str1.size(); i++){
            if(str1[i] == symbol) {
                if(mappedChar == '\0') {
                    mappedChar = str2[i];
                } else if(mappedChar != str2[i]) {
                    return false;  // Один символ отображается в разные - не изоморфно
                }
            }
        }
        
        // Проверяем, что mappedChar не используется для другого символа
        for(int i = 0; i < size; i++){
            if(table[i].second == mappedChar) {
                return false;  // Два разных символа отображаются в один - не изоморфно
            }
        }
        
        // Добавляем новое отображение в таблицу
        table[size] = {symbol, mappedChar};
        size++;
        if (size == cap){
            cap *= 2;
            pair<char, char>* newTable = new pair<char, char>[cap];
            for(int i = 0; i < size; i++) newTable[i] = table[i];
            delete[] table;
            table = newTable;
        }
        return true;
    }

    ~morfTable() {
        delete[] table;
    }
};

void isMorf(string str1, string str2){
    if (str1.size() != str2.size()) {
        cout << "FALSE" << endl;
        return;
    }
    morfTable table;
    for (int i = 0; i < str1.size(); i++){
        if (!(table.searchSimbolMorf(str1, str2, str1[i]))){
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