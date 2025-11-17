#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <algorithm>
#include <map>

using namespace std;

enum class Color{
    RED,
    BLACK
};

class RBNode{
    public:
        string data;
        Color color;
        RBNode* parent;
        RBNode* left;
        RBNode* right;
        
        RBNode(const string& value) : data(value), color(Color::RED), parent(nullptr), left(nullptr), right(nullptr) {}
};

// Предварительные объявления функций красно-черного дерева
class RBTree;
void addRBNode(RBTree* tree, const string& value);
void delRBNode(RBTree* tree, const string& value);
RBNode* getRBNode(const RBTree* tree, const string& value);
void print(const RBTree* tree);

class RBTree{
    public:
        RBNode* root;
        RBTree() : root(nullptr) {}
        
        // Итератор для обхода дерева
        class Iterator {
        private:
            RBNode* current;
            RBNode* findMin(RBNode* node) {
                while (node && node->left) node = node->left;
                return node;
            }
            
            RBNode* findSuccessor(RBNode* node) {
                if (node->right) return findMin(node->right);
                
                RBNode* parent = node->parent;
                while (parent && node == parent->right) {
                    node = parent;
                    parent = parent->parent;
                }
                return parent;
            }
            
        public:
            Iterator(RBNode* node) : current(node) {}
            
            string& operator*() { return current->data; }
            string* operator->() { return &current->data; }
            
            Iterator& operator++() {
                current = findSuccessor(current);
                return *this;
            }
            
            Iterator operator++(int) {
                Iterator temp = *this;
                ++(*this);
                return temp;
            }
            
            bool operator==(const Iterator& other) const { return current == other.current; }
            bool operator!=(const Iterator& other) const { return current != other.current; }
        };
        
        Iterator begin() const {
            return Iterator(findMin(root));
        }
        
        Iterator end() const {
            return Iterator(nullptr);
        }
        
    private:
        RBNode* findMin(RBNode* node) const {
            while (node && node->left) node = node->left;
            return node;
        }
};

class StringSet {
private:
    RBTree tree;

public:
    // Конструкторы
    StringSet() = default;
    
    StringSet(const vector<string>& elements) {
        for (const auto& elem : elements) {
            add(elem);
        }
    }
    
    // Базовые операции множества
    void add(const string& value) {
        if (contains(value)) return;
        ::addRBNode(&tree, value);
    }
    
    void remove(const string& value) {
        ::delRBNode(&tree, value);
    }
    
    bool contains(const string& value) const {
        return ::getRBNode(&tree, value) != nullptr;
    }
    
    size_t size() const {
        size_t count = 0;
        for (auto it = begin(); it != end(); ++it) count++;
        return count;
    }
    
    bool empty() const {
        return tree.root == nullptr;
    }
    
    void clear() {
        clearTree(tree.root);
        tree.root = nullptr;
    }
    
    // Операции множества
    StringSet unionWith(const StringSet& other) const {
        StringSet result = *this;
        for (const auto& elem : other) {
            result.add(elem);
        }
        return result;
    }
    
    StringSet intersectionWith(const StringSet& other) const {
        StringSet result;
        for (const auto& elem : *this) {
            if (other.contains(elem)) {
                result.add(elem);
            }
        }
        return result;
    }
    
    StringSet differenceWith(const StringSet& other) const {
        StringSet result;
        for (const auto& elem : *this) {
            if (!other.contains(elem)) {
                result.add(elem);
            }
        }
        return result;
    }
    
    // Проверка принадлежности
    bool operator==(const StringSet& other) const {
        if (size() != other.size()) return false;
        for (const auto& elem : *this) {
            if (!other.contains(elem)) return false;
        }
        return true;
    }
    
    bool operator!=(const StringSet& other) const {
        return !(*this == other);
    }
    
    // Итераторы
    RBTree::Iterator begin() const { return tree.begin(); }
    RBTree::Iterator end() const { return tree.end(); }
    
    // Вывод
    void print() const {
        ::print(&tree);
    }
    
    // Вектор элементов
    vector<string> toVector() const {
        vector<string> result;
        for (const auto& elem : *this) {
            result.push_back(elem);
        }
        return result;
    }

private:
    void clearTree(RBNode* node) {
        if (node == nullptr) return;
        clearTree(node->left);
        clearTree(node->right);
        delete node;
    }
};

// Функции красно-черного дерева
void leftRotate(RBTree* tree, RBNode* x){
    RBNode* y = x->right;
    x->right = y->left;
    if (y->left != nullptr){
        y->left->parent = x;
    }
    y->parent = x->parent;
    if (x->parent == nullptr){
        tree->root = y;
    }else{
        if(x == x->parent->left) x->parent->left = y;
        else x->parent->right = y;
    }
    y->left = x;
    x->parent = y;
}

void rightRotate(RBTree* tree, RBNode* x){
    RBNode* y = x->left;
    x->left = y->right;
    if (y->right != nullptr){
        y->right->parent = x;
    }
    y->parent = x->parent;
    if (x->parent == nullptr){
        tree->root = y;
    }else{
        if(x == x->parent->left) x->parent->left = y;
        else x->parent->right = y;
    }
    y->right = x;
    x->parent = y;
}

void fixAdd(RBTree* tree, RBNode* z){
    while (z->parent != nullptr && z->parent->color == Color::RED){
        
        if (z->parent->parent == nullptr) {
            break;
        }

        if(z->parent == z->parent->parent->left){
            RBNode* y = z->parent->parent->right;
            
            if (y != nullptr && y->color == Color::RED){
                z->parent->color = Color::BLACK;
                y->color = Color::BLACK;
                z->parent->parent->color = Color::RED;
                z = z->parent->parent;
            }
            else{
                if (z == z->parent->right){
                    z = z->parent;
                    leftRotate(tree, z);
                }
                z->parent->color = Color::BLACK;
                z->parent->parent->color = Color::RED;
                rightRotate(tree, z->parent->parent);
            }
        }
        else{
            RBNode* y = z->parent->parent->left;
            
            if (y != nullptr && y->color == Color::RED){
                z->parent->color = Color::BLACK;
                y->color = Color::BLACK;
                z->parent->parent->color = Color::RED;
                z = z->parent->parent;
            }
            else{
                if (z == z->parent->left){
                    z = z->parent;
                    rightRotate(tree, z);
                }
                z->parent->color = Color::BLACK;
                z->parent->parent->color = Color::RED;
                leftRotate(tree, z->parent->parent);
            }
        }
        
        if (z == tree->root) {
            break;
        }
    }
    tree->root->color = Color::BLACK;
}

void addRBNode(RBTree* tree, const string& value){
    RBNode* new_RBNode = new RBNode(value);
    if (tree->root == nullptr){
        new_RBNode->color = Color::BLACK;
        tree->root = new_RBNode;
        return;
    }
    RBNode* address = tree->root;
    while(true){
        if(value < address->data){
            if (address->left == nullptr){
                new_RBNode->parent = address;
                address->left = new_RBNode;
                fixAdd(tree, new_RBNode);
                return;
            }else{
                address = address->left;
            }
        }
        else{
            if (address->right == nullptr){
                new_RBNode->parent = address;
                address->right = new_RBNode;
                fixAdd(tree, new_RBNode);
                return;
            }else{
                address = address->right;
            }
        }
    }  
}

RBNode* getRBNode(const RBTree* tree, const string& value) {
    if (tree->root == nullptr) {
        return nullptr;
    }
    
    RBNode* address = tree->root;
    while (address != nullptr) {
        if (value == address->data) {
            return address;
        } else if (value < address->data) {
            address = address->left;
        } else {
            address = address->right;
        }
    }
    
    return nullptr;
}

RBNode* treeMinimum(RBTree* tree, RBNode* node) {
    if (node == nullptr) return nullptr;
    RBNode* address = node;
    while (address->left != nullptr) {
        address = address->left;
    }
    return address;
}

void deleteFix(RBTree* tree, RBNode* x) {
    // Если x nullptr, значит дерево пустое или это последний узел
    if (x == nullptr) return;
    
    while (x != tree->root && x->color == Color::BLACK) {
        if (x == x->parent->left) {
            RBNode* w = x->parent->right;
            if (w == nullptr) break;
            
            if (w->color == Color::RED) {
                w->color = Color::BLACK;
                x->parent->color = Color::RED;
                leftRotate(tree, x->parent);
                w = x->parent->right;
                if (w == nullptr) break;
            }
            
            bool leftBlack = (w->left == nullptr) || (w->left->color == Color::BLACK);
            bool rightBlack = (w->right == nullptr) || (w->right->color == Color::BLACK);
            
            if (leftBlack && rightBlack) {
                w->color = Color::RED;
                x = x->parent;
            } else {
                if (w->right == nullptr || w->right->color == Color::BLACK) {
                    if (w->left != nullptr) {
                        w->left->color = Color::BLACK;
                    }
                    w->color = Color::RED;
                    rightRotate(tree, w);
                    w = x->parent->right;
                    if (w == nullptr) break;
                }
                
                w->color = x->parent->color;
                x->parent->color = Color::BLACK;
                if (w->right != nullptr) {
                    w->right->color = Color::BLACK;
                }
                leftRotate(tree, x->parent);
                x = tree->root;
            }
        } else {
            RBNode* w = x->parent->left;
            if (w == nullptr) break;
            
            if (w->color == Color::RED) {
                w->color = Color::BLACK;
                x->parent->color = Color::RED;
                rightRotate(tree, x->parent);
                w = x->parent->left;
                if (w == nullptr) break;
            }
            
            bool leftBlack = (w->left == nullptr) || (w->left->color == Color::BLACK);
            bool rightBlack = (w->right == nullptr) || (w->right->color == Color::BLACK);
            
            if (leftBlack && rightBlack) {
                w->color = Color::RED;
                x = x->parent;
            } else {
                if (w->left == nullptr || w->left->color == Color::BLACK) {
                    if (w->right != nullptr) {
                        w->right->color = Color::BLACK;
                    }
                    w->color = Color::RED;
                    leftRotate(tree, w);
                    w = x->parent->left;
                    if (w == nullptr) break;
                }
                
                w->color = x->parent->color;
                x->parent->color = Color::BLACK;
                if (w->left != nullptr) {
                    w->left->color = Color::BLACK;
                }
                rightRotate(tree, x->parent);
                x = tree->root;
            }
        }
    }
    
    if (x != nullptr) {
        x->color = Color::BLACK;
    }
}

void transplant(RBTree* tree, RBNode* u, RBNode* v) {
    if (u->parent == nullptr) {
        tree->root = v;
    } else if (u == u->parent->left) {
        u->parent->left = v;
    } else {
        u->parent->right = v;
    }
    
    if (v != nullptr) {
        v->parent = u->parent;
    }
}

void delRBNode(RBTree* tree, const string& value) {
    RBNode* z = getRBNode(tree, value);
    if (z == nullptr) return;
    
    RBNode* y = z;
    RBNode* x = nullptr;
    Color y_original_color = y->color;
    
    if (z->left == nullptr) {
        x = z->right;
        transplant(tree, z, x);
    } else if (z->right == nullptr) {
        x = z->left;
        transplant(tree, z, x);
    } else {
        y = treeMinimum(tree, z->right);
        y_original_color = y->color;
        x = y->right;
        
        if (y->parent != z) {
            transplant(tree, y, x);
            y->right = z->right;
            if (y->right != nullptr) {
                y->right->parent = y;
            }
        } else {
            if (x != nullptr) {
                x->parent = y;
            }
        }
        
        transplant(tree, z, y);
        y->left = z->left;
        if (y->left != nullptr) {
            y->left->parent = y;
        }
        y->color = z->color;
    }
    
    delete z;
    
    if (y_original_color == Color::BLACK) {
        if (x != nullptr) {
            deleteFix(tree, x);
        }
        // Если x nullptr, значит удалили последний узел
    }
}

// База данных множеств
class SetDatabase {
private:
    map<string, StringSet> sets;

public:
    // Добавление множества
    void addSet(const string& name, const StringSet& set) {
        sets[name] = set;
    }
    
    // Удаление множества
    void removeSet(const string& name) {
        sets.erase(name);
    }
    
    // Получение множества
    StringSet* getSet(const string& name) {
        auto it = sets.find(name);
        if (it != sets.end()) {
            return &it->second;
        }
        return nullptr;
    }
    
    // Проверка существования множества
    bool containsSet(const string& name) const {
        return sets.find(name) != sets.end();
    }
    
    // Очистка множества (удаление всех элементов)
    void clearSet(const string& name) {
        auto it = sets.find(name);
        if (it != sets.end()) {
            it->second.clear();
        }
    }
    
    // Получение всех имен множеств
    vector<string> getSetNames() const {
        vector<string> names;
        for (const auto& pair : sets) {
            names.push_back(pair.first);
        }
        return names;
    }
    
    // Сохранение в файл
    void saveToFile(const string& filename) {
        ofstream file(filename);
        if (!file.is_open()) {
            cerr << "Ошибка: не удалось открыть файл для записи: " << filename << endl;
            return;
        }
        
        file << "SETS " << sets.size() << endl;
        for (const auto& pair : sets) {
            file << "SET " << pair.first << endl;
            
            // Сохраняем множество в порядке предварительного обхода
            vector<string> setData;
            for (const auto& elem : pair.second) {
                setData.push_back(elem);
            }
            file << setData.size() << endl;
            for (const auto& data : setData) {
                file << data << endl;
            }
        }
        
        file.close();
        cout << "База данных сохранена в " << filename << endl;
    }
    
    // Загрузка из файла
    void loadFromFile(const string& filename) {
        ifstream file(filename);
        if (!file.is_open()) {
            cerr << "Ошибка: не удалось открыть файл для чтения: " << filename << endl;
            return;
        }
        
        sets.clear();
        string section;
        
        while (file >> section) {
            if (section == "SETS") {
                int count;
                file >> count;
                for (int i = 0; i < count; i++) {
                    string type, name;
                    file >> type >> name;
                    
                    int elementCount;
                    file >> elementCount;
                    file.ignore(); // Пропускаем перевод строки
                    
                    vector<string> setData(elementCount);
                    for (int j = 0; j < elementCount; j++) {
                        getline(file, setData[j]);
                    }
                    
                    StringSet newSet;
                    for (const auto& elem : setData) {
                        newSet.add(elem);
                    }
                    sets[name] = newSet;
                }
            }
        }
        
        file.close();
        cout << "База данных загружена из " << filename << endl;
    }
};

// Главная функция
int main(int argc, char* argv[]) {
    string filename, query;

    // Разбор аргументов командной строки
    for (int i = 1; i < argc; ++i) {
        string arg = argv[i];
        if (arg == "--file" && i + 1 < argc) {
            filename = argv[++i];
        } else if (arg == "--query" && i + 1 < argc) {
            query = argv[++i];
        }
    }

    if (filename.empty() || query.empty()) {
        cerr << "Использование: " << argv[0] << " --file <filename> --query <query>" << endl;
        cerr << "Доступные операции:" << endl;
        cerr << "  SETCREATE <setname> - создать новое множество" << endl;
        cerr << "  SETADD <setname> <value> - добавить элемент в множество" << endl;
        cerr << "  SETDEL <setname> <value> - удалить элемент из множества" << endl;
        cerr << "  SET_AT <setname> <value> - проверить наличие элемента" << endl;
        cerr << "  SET_UNION <setname1> <setname2> <resultname> - объединение множеств" << endl;
        cerr << "  SET_INTERSECT <setname1> <setname2> <resultname> - пересечение множеств" << endl;
        cerr << "  SET_DIFF <setname1> <setname2> <resultname> - разность множеств" << endl;
        cerr << "  SET_PRINT <setname> - вывести все элементы множества" << endl;
        cerr << "  SET_LIST - вывести список всех множеств" << endl;
        cerr << "  SETREMOVE <setname> - удалить множество" << endl;
        cerr << "  SETCLEAR <setname> - очистить все элементы множества" << endl;
        return 1;
    }

    // Загрузка базы данных
    SetDatabase database;
    database.loadFromFile(filename);

    // Разбор запроса
    vector<string> tokens;
    size_t start = 0, end = 0;
    while ((end = query.find(' ', start)) != string::npos) {
        tokens.push_back(query.substr(start, end - start));
        start = end + 1;
    }
    tokens.push_back(query.substr(start));

    if (tokens.empty()) {
        cerr << "Неверный формат запроса" << endl;
        return 1;
    }

    string operation = tokens[0];
    bool needSave = false;

    // Обработка операций
    if (operation == "SETCREATE" && tokens.size() == 2) {
        string setName = tokens[1];
        if (!database.containsSet(setName)) {
            database.addSet(setName, StringSet());
            needSave = true;
            cout << "Множество '" << setName << "' создано" << endl;
        } else {
            cout << "Множество '" << setName << "' уже существует" << endl;
        }
    } else if (operation == "SETADD" && tokens.size() == 3) {
        string setName = tokens[1];
        string value = tokens[2];
        StringSet* set = database.getSet(setName);
        if (set) {
            set->add(value);
            needSave = true;
            cout << "Элемент '" << value << "' добавлен в множество '" << setName << "'" << endl;
        } else {
            cerr << "Множество '" << setName << "' не найдено" << endl;
        }
    } else if (operation == "SETDEL" && tokens.size() == 3) {
        string setName = tokens[1];
        string value = tokens[2];
        StringSet* set = database.getSet(setName);
        if (set) {
            set->remove(value);
            needSave = true;
            cout << "Элемент '" << value << "' удален из множества '" << setName << "'" << endl;
        } else {
            cerr << "Множество '" << setName << "' не найдено" << endl;
        }
    } else if (operation == "SET_AT" && tokens.size() == 3) {
        string setName = tokens[1];
        string value = tokens[2];
        StringSet* set = database.getSet(setName);
        if (set) {
            cout << (set->contains(value) ? "true" : "false") << endl;
        } else {
            cerr << "Множество '" << setName << "' не найдено" << endl;
        }
    } else if (operation == "SET_UNION" && tokens.size() == 4) {
        string setName1 = tokens[1];
        string setName2 = tokens[2];
        string resultName = tokens[3];
        StringSet* set1 = database.getSet(setName1);
        StringSet* set2 = database.getSet(setName2);
        if (set1 && set2) {
            StringSet result = set1->unionWith(*set2);
            database.addSet(resultName, result);
            needSave = true;
            cout << "Объединение множеств '" << setName1 << "' и '" << setName2 << "' сохранено как '" << resultName << "'" << endl;
        } else {
            cerr << "Одно из множеств не найдено" << endl;
        }
    } else if (operation == "SET_INTERSECT" && tokens.size() == 4) {
        string setName1 = tokens[1];
        string setName2 = tokens[2];
        string resultName = tokens[3];
        StringSet* set1 = database.getSet(setName1);
        StringSet* set2 = database.getSet(setName2);
        if (set1 && set2) {
            StringSet result = set1->intersectionWith(*set2);
            database.addSet(resultName, result);
            needSave = true;
            cout << "Пересечение множеств '" << setName1 << "' и '" << setName2 << "' сохранено как '" << resultName << "'" << endl;
        } else {
            cerr << "Одно из множеств не найдено" << endl;
        }
    } else if (operation == "SET_DIFF" && tokens.size() == 4) {
        string setName1 = tokens[1];
        string setName2 = tokens[2];
        string resultName = tokens[3];
        StringSet* set1 = database.getSet(setName1);
        StringSet* set2 = database.getSet(setName2);
        if (set1 && set2) {
            StringSet result = set1->differenceWith(*set2);
            database.addSet(resultName, result);
            needSave = true;
            cout << "Разность множеств '" << setName1 << "' и '" << setName2 << "' сохранена как '" << resultName << "'" << endl;
        } else {
            cerr << "Одно из множеств не найдено" << endl;
        }
    } else if (operation == "SET_PRINT" && tokens.size() == 2) {
        string setName = tokens[1];
        StringSet* set = database.getSet(setName);
        if (set) {
            cout << "Элементы множества '" << setName << "':" << endl;
            for (const auto& elem : *set) {
                cout << elem << endl;
            }
            cout << "Всего элементов: " << set->size() << endl;
        } else {
            cerr << "Множество '" << setName << "' не найдено" << endl;
        }
    } else if (operation == "SET_LIST") {
        vector<string> setNames = database.getSetNames();
        cout << "Доступные множества:" << endl;
        for (const auto& name : setNames) {
            cout << "- " << name << endl;
        }
    } else if (operation == "SETREMOVE" && tokens.size() == 2) {
        string setName = tokens[1];
        if (database.containsSet(setName)) {
            database.removeSet(setName);
            needSave = true;
            cout << "Множество '" << setName << "' удалено" << endl;
        } else {
            cerr << "Множество '" << setName << "' не найдено" << endl;
        }
    } else if (operation == "SETCLEAR" && tokens.size() == 2) {
        string setName = tokens[1];
        StringSet* set = database.getSet(setName);
        if (set) {
            set->clear();
            needSave = true;
            cout << "Все элементы множества '" << setName << "' удалены" << endl;
        } else {
            cerr << "Множество '" << setName << "' не найдено" << endl;
        }
    } else {
        cerr << "Неизвестная команда или неверный формат: " << operation << endl;
        return 1;
    }

    // Сохранение изменений в файл
    if (needSave) {
        database.saveToFile(filename);
    }

    return 0;
}