#include <iostream>
#include <string>
#include <sstream>

using namespace std;

struct Node {
    string key;
    Node* prev;
    Node* next;
    
    Node(string k) : key(k), prev(nullptr), next(nullptr) {}
};

class LRUCache {
private:
    int capacity;
    Node* head;
    Node* tail;

    // Поиск узла по ключу
    Node* findNode(string key) {
        Node* current = head->next;
        while (current != tail) {
            if (current->key == key) {
                return current;
            }
            current = current->next;
        }
        return nullptr;
    }

    // Удаление узла из списка
    void removeNode(Node* node) {
        node->prev->next = node->next;
        node->next->prev = node->prev;
    }

    // Добавление узла в начало списка (после head)
    void addToFront(Node* node) {
        node->next = head->next;
        node->prev = head;
        head->next->prev = node;
        head->next = node;
    }

    // Перемещение узла в начало
    void moveToFront(Node* node) {
        removeNode(node);
        addToFront(node);
    }

    // Удаление LRU элемента
    void removeLRU() {
        Node* lru = tail->prev;
        removeNode(lru);
        delete lru;
    }

public:
    LRUCache(int cap) : capacity(cap) {
        // Создаем фиктивные голову и хвост для упрощения операций
        head = new Node("");
        tail = new Node("");
        head->next = tail;
        tail->prev = head;
    }

    ~LRUCache() {
        Node* current = head;
        while (current) {
            Node* next = current->next;
            delete current;
            current = next;
        }
    }

    // GET операция
    bool get(string key) {
        Node* node = findNode(key);
        if (node != nullptr) {
            moveToFront(node);
            return true;
        }
        return false;
    }

    // SET операция
    void set(string key) {
        Node* node = findNode(key);
        if (node != nullptr) {
            // Ключ уже существует - просто перемещаем в начало
            moveToFront(node);
        } else {
            // Новый ключ
            if (getSize() >= capacity) {
                removeLRU();
            }
            
            // Создаем новый узел и добавляем в начало
            Node* newNode = new Node(key);
            addToFront(newNode);
        }
    }

    // Вспомогательная функция для получения размера
    int getSize() {
        int size = 0;
        Node* current = head->next;
        while (current != tail) {
            size++;
            current = current->next;
        }
        return size;
    }

    void printCache() {
        cout << "Кэш: ";
        Node* current = head->next;
        while (current != tail) {
            cout << "[\"" << current->key << "\"] ";
            current = current->next;
        }
        cout << endl;
    }
};

void LRU() {
    cout << "Режим LRU-кеша (LRUCASH)" << endl;
    cout << "Введите размер кеша:" << endl;
    
    int cacheSize;
    cin >> cacheSize;
    
    if (cacheSize <= 0) {
        cout << "Ошибка: размер кеша должен быть положительным числом" << endl;
        return;
    }
    
    LRUCache cache(cacheSize);
    cout << "LRU-кеш создан с размером " << cacheSize << endl;
    
    cin.ignore(); // Очищаем буфер после ввода числа
    
    while (true) {
        cout << endl << "Доступные команды:" << endl;
        cout << "SET <ключ> - добавить элемент в кеш" << endl;
        cout << "GET <ключ> - проверить наличие элемента в кеше" << endl;
        cout << "PRINT - вывести текущее состояние кеша" << endl;
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
        else if (command == "SET") {
            string key;
            if (ss >> key) {
                cache.set(key);
                cout << "Элемент добавлен в кеш: ключ='" << key << "'" << endl;
                cache.printCache();
            } else {
                cout << "Ошибка: неверный формат команды. Используйте: SET <ключ>" << endl;
            }
        }
        else if (command == "GET") {
            string key;
            if (ss >> key) {
                bool found = cache.get(key);
                if (found) {
                    cout << "Элемент найден в кеше: ключ='" << key << "'" << endl;
                } else {
                    cout << "Элемент отсутствует в кеше: ключ='" << key << "'" << endl;
                }
                cache.printCache();
            } else {
                cout << "Ошибка: неверный формат команды. Используйте: GET <ключ>" << endl;
            }
        }
        else if (command == "PRINT") {
            cache.printCache();
        }
        else {
            cout << "Ошибка: неизвестная команда '" << command << "'" << endl;
        }
    }
}