#include <iostream>
#include <algorithm>
#include <sstream>

using namespace std;

size_t genHashFree(int size, string key){
    size_t result = 1245;
    for (size_t i = 0; i < key.size(); i++){
        result += i*key[i] % size;
    }
    return result % size;
}

template<typename V>
struct HashTableFree{
    pair<string, V>* table;
    size_t size;
    

    HashTableFree(size_t sz) : size(sz){
        table = new pair<string, V>[sz];
        for(int i = 0; i < size; i++) table[i].first = "_empty_";
    }
    ~HashTableFree(){
        delete[] table;
    }

    // Преобразуем любой ключ в строку
    template<typename K>
    string keyToString(K key) {
        if constexpr (is_same_v<K, string>) { //constexpr
            return key;
        } else {
            return to_string(key);
        }
    }

    template<typename K>
    void insert(K key, V value){
        string strKey = keyToString(key);
        size_t Hash = genHashFree(size, strKey);
        if (table[Hash].first == "_empty_"){
            table[Hash] = {strKey, value};
            cout << "Элемент добавлен: ключ='" << key << "', значение='" << value << "'" << endl;
        }else{
            int i = 1;
            while(1){
                if (table[(Hash + i)%size].first == "_empty_"){
                    table[(Hash + i)%size] = {strKey, value};
                    cout << "Элемент добавлен: ключ='" << key << "', значение='" << value << "'" << endl;
                    return;
                } 
                i++;
                if (i > size*2){
                    cout << "Свободное место отсутствует в таблице" << endl;
                    return;
                }
            }
        }
    }

    template<typename K>
    void remove(K key){
        string strKey = keyToString(key);
        size_t Hash = genHashFree(size, strKey);
        if (table[Hash].first == strKey){
            table[Hash].first = "_empty_";
            cout << "Элемент успешно удалён" << endl;
        }else{
            int i = 1;
            while(1){
                if (table[(Hash + i)%size].first == strKey){
                    table[(Hash + i)%size].first = "_empty_";
                    cout << "Элемент успешно удалён" << endl;
                    return;
                } 
                i++;
                if (i > size*2){
                    cout << "Элемент отсутствует в таблице" << endl;
                    return;
                }
            }
        }
    }

    template<typename K>
    void find(K key){
        string strKey = keyToString(key);
        size_t Hash = genHashFree(size, strKey);
        if (table[Hash].first == strKey){
            cout << "Data = " << table[Hash].second << endl;
        }else{
            int i = 1;
            while(1){
                if (table[(Hash + i)%size].first == strKey){
                    cout << "Data = " << table[(Hash + i)%size].second << endl;
                    return;
                } 
                i++;
                if (i > size*2){
                    cout << "Элемент отсутствует в таблице" << endl;
                    return;
                }
            }
        }
    }
};

void hashTableFree(){
    cout << "Введите размер хеш-таблицы:" << endl;
    
    size_t tableSize;
    cin >> tableSize;
    
    if (tableSize <= 0) {
        cout << "Ошибка: размер таблицы должен быть положительным числом" << endl;
        return;
    }
    
    // ИЗМЕНЕНИЕ: используем string вместо int
    HashTableFree<string> hashTable(tableSize);
    cout << "Хеш-таблица создана с размером " << tableSize << endl;
    
    cin.ignore(); // Очищаем буфер после ввода числа
    
    while (true) {
        cout << endl << "Доступные команды:" << endl;
        cout << "INSERT <ключ> <значение> - добавить элемент" << endl;
        cout << "REMOVE <ключ> - удалить элемент" << endl;
        cout << "FIND <ключ> - найти элемент" << endl;
        cout << "EXIT - выход из программы" << endl;
        cout << "Введите команду:" << endl;
        
        string input;
        getline(cin, input);
        
        stringstream ss(input);
        string command;
        ss >> command;
        
        if (command == "EXIT") {
            cout << "Выход из программы" << endl;
            break;
        }
        else if (command == "INSERT") {
            string key;
            string value; // ИЗМЕНЕНИЕ: string вместо int
            if (ss >> key >> value) {
                hashTable.insert(key, value);
            } else {
                cout << "Ошибка: неверный формат команды. Используйте: INSERT <ключ> <значение>" << endl;
            }
        }
        else if (command == "REMOVE") {
            string key;
            if (ss >> key) {
                cout << "Попытка удаления элемента с ключом='" << key << "'..." << endl;
                hashTable.remove(key);
            } else {
                cout << "Ошибка: неверный формат команды. Используйте: REMOVE <ключ>" << endl;
            }
        }
        else if (command == "FIND") {
            string key;
            if (ss >> key) {
                hashTable.find(key);
            } else {
                cout << "Ошибка: неверный формат команды. Используйте: FIND <ключ>" << endl;
            }
        }
        else {
            cout << "Ошибка: неизвестная команда '" << command << "'" << endl;
        }
    }
}