#include <iostream>
#include <algorithm>
#include <sstream>

using namespace std;

struct Chain{
    string key;
    string data;
    Chain* next;
    Chain(string key, string data) : key(key), data(data), next(nullptr){}
};

size_t genHash(int size, string key){
    size_t result = 1245;
    for (size_t i = 0; i < key.size(); i++){
        result += i*key[i] % size;
    }
    return result % size;
}

struct HashTableChain{
    Chain** table;
    size_t size;
    
    HashTableChain(size_t sz) : size(sz){
        table = new Chain*[sz];
        fill(table, table + size, nullptr);
    }
    ~HashTableChain(){
        for (size_t i = 0; i < size; i++) {
            Chain* current = table[i];
            while (current != nullptr) {
                Chain* next = current->next;
                delete current;
                current = next;
            }
        }
        delete[] table;
    }

    void insert(string key, string value){
        size_t Hash = genHash(size, key);
        if (table[Hash] == nullptr){
            table[Hash] = new Chain(key, value);
        }else{
            Chain* address = table[Hash];
            while (address->next != nullptr) address = address->next;
            address->next = new Chain(key, value);
        }
    }

    void remove(string key){
        size_t Hash = genHash(size, key);
        Chain* address = table[Hash];
        
        if (address == nullptr) {
            cout << "Элемент отсутствует в таблице" << endl;
            return;
        }
        
        if (address->key == key){
            table[Hash] = address->next;
            delete address;
            return;
        }
        
        while(address->next != nullptr && address->next->key != key){
            address = address->next;
        }
        
        if (address->next == nullptr){
            cout << "Элемент отсутствует в таблице" << endl;
            return;
        }
        
        Chain* deleteChain = address->next;
        address->next = address->next->next;
        delete deleteChain;
    }

    void find(string key){
        size_t Hash = genHash(size, key);
        Chain* address = table[Hash];
        if (address == nullptr){
            cout << "Элемент отсутствует в таблице" << endl;
            return;
        } 
        while (address != nullptr){
            if (address->key == key){
                cout << "Элемент найден, data = " << address->data << endl; 
                return;
            }
            address = address->next;
        }
        cout << "Элемент отсутствует в таблице" << endl;
    }
};

void hashTableChain(){
    cout << "Режим хеш-таблицы с цепочками (CHAINHASH)" << endl;
    cout << "Введите размер хеш-таблицы:" << endl;
    
    size_t tableSize;
    cin >> tableSize;
    
    if (tableSize <= 0) {
        cout << "Ошибка: размер таблицы должен быть положительным числом" << endl;
        return;
    }
    
    HashTableChain hashTable(tableSize);
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
            string value;
            if (ss >> key >> value) {
                hashTable.insert(key, value);
                cout << "Элемент добавлен: ключ='" << key << "', значение='" << value << "'" << endl;
            } else {
                cout << "Ошибка: неверный формат команды. Используйте: INSERT <ключ> <значение>" << endl;
            }
        }
        else if (command == "REMOVE") {
            string key;
            if (ss >> key) {
                hashTable.remove(key);
                cout << "Попытка удаления элемента с ключом='" << key << "'" << endl;
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