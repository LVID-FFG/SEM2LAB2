#include <iostream>
#include <algorithm>
#include <sstream>

using namespace std;

struct HNode{
    char symbol;
    double probability;
    string code;
    HNode* parent;
    HNode* left;
    HNode* right;
    HNode(pair<double, char> sym) : symbol(sym.second)
                                , probability(sym.first)
                                , code("")
                                , parent(nullptr)
                                , left(nullptr)
                                , right(nullptr){}
    HNode(HNode* l, HNode* r) : symbol('\0')
                                , probability(l -> probability + r->probability)
                                , code("")
                                , parent(nullptr)
                                , left(l)
                                , right(r){}
};

ostream& operator<<(ostream& ss, pair<char, string> dc){
    ss << dc.first << " " << dc.second;
    return ss;
}

ostream& operator<<(ostream& ss, pair<double, char> dc){
    ss << dc.first << " " << dc.second;
    return ss;
}

// Функция для создания кодов
void createCode(HNode* node, string currentCode = "" ){
    if (!node) return;
    
    node->code = currentCode;
    
    if (node->left != nullptr) createCode(node->left, currentCode + "0");
    if (node->right != nullptr) createCode(node->right, currentCode + "1");
}

void buildCodeTable(pair<char, string>* codeTable, HNode* curentNode, int& index){
    if (curentNode == nullptr) return;
    if (curentNode -> symbol == '\0'){
        buildCodeTable(codeTable, curentNode -> left, index);
        buildCodeTable(codeTable, curentNode -> right, index);
        return;
    }
    codeTable[index] = {curentNode -> symbol, curentNode -> code};
    index++;
}

struct HafmanCode{
    pair<char, string>* codeTable;
    size_t size;
    HafmanCode(string str){
        string unique;
        for (char s : str) if (unique.find(s) == string::npos) unique += s;

        pair<double, char> Table[unique.size()];
        int* j = new int(0);
        for (char i : unique){
            int cnt = count(str.begin(), str.end(), i);
            Table[(*j)++] = {static_cast<double>(cnt)/static_cast<double>(str.size()), i};
        }
        delete j;

        sort(Table, Table + unique.size(), [](pair<double, char> a, pair<double, char> b){
            return  a.first < b.first;
        });

        HNode* workTable[unique.size()];
        for(int i = 0; i < unique.size(); i++) {
            HNode* node = new HNode(Table[i]);
            workTable[i] = node;
        }

        size_t sizeWorkTable = unique.size();
        while(sizeWorkTable != 1){
            HNode* node = new HNode(workTable[0], workTable[1]);
            workTable[0] -> parent = node;
            workTable[1] -> parent = node;
            sizeWorkTable--;
            for (int i = 0; i < sizeWorkTable; i++) workTable[i] = workTable[i+1];
            workTable[0] = node;
            sort(workTable, workTable + sizeWorkTable, [](HNode* a, HNode* b){
                return  a -> probability < b -> probability;
            });
        }

        HNode* head;
        head = workTable[0];
        createCode(head);

        codeTable = new pair<char, string>[unique.size()];
        size = unique.size();
        int index = 0;
        buildCodeTable(codeTable, head, index);
        
        string result = "";
        for(char i : str){
            for(int j = 0; j < unique.size(); j++) if(codeTable[j].first == i) result += codeTable[j].second;
        }
        cout << "Код: " << result << endl;
        cout << "Таблица кодирования:" << endl;
        for(int i = 0; i < unique.size(); i++) cout << codeTable[i] << endl;
    }
    
    // Добавляем деструктор для очистки памяти
    ~HafmanCode() {
        delete[] codeTable;
    }
    
    void decode(string code){
        string result = "";
        string decodeSymbol = "";
        for (int i = 0; i < code.size(); i++){
            decodeSymbol += code[i];
            for(int j = 0; j < size; j++){
                if(codeTable[j].second == decodeSymbol) {
                    result += codeTable[j].first;
                    decodeSymbol = "";
                }
            } 
        }
        cout << "Декодированная строка: " << result << endl;
    }
};

void hafman(){
    cout << "Введите строку для кодирования:" << endl;
    
    string input;
    cin.ignore();
    getline(cin, input);
    HafmanCode huffman(input);

    cout << endl << "Введите код для декодирования (или 'stop' для выхода):" << endl;
    string code_to_decode;
    getline(cin, code_to_decode);
    
    while (code_to_decode != "stop") {
        huffman.decode(code_to_decode);
        
        cout << endl << "Введите следующий код для декодирования (или 'stop' для выхода):" << endl;
        getline(cin, code_to_decode);
    }
}