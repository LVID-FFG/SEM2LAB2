#include <iostream>
#include  <sstream>
using namespace std;

class SNode{
public:
    pair<int, string> data;
    SNode* next;
    SNode(pair<int, string> dt, SNode* nx) : data(dt)
                                , next(nx){}
};

class Stack{
public:
    SNode* head;
    Stack(SNode* node) : head(node){}
    Stack(pair<int, string> data) : head(new SNode{data, nullptr}){}
    Stack() : head(nullptr){}
};

void SPUSH(Stack& list, pair<int, string> data){
    SNode* new_SNode = new SNode{data, list.head};
    list.head = new_SNode;
}

void SPOP(Stack& list){
    SNode* Address = list.head -> next;
    delete list.head;
    list.head = Address;
}


ostream& operator<<(ostream& ss, pair<int, string> data){
    ss << data.first << " " << data.second << " ";
    return ss;
}

void SPRINT(Stack& list){
    SNode* Address = list.head;
    while (Address != nullptr){
        cout << Address->data << endl;
        Address = Address->next;
    }
}

void ReverseStack(Stack& list) {
    Stack reversed;
    SNode* current = list.head;
    while (current != nullptr) {
        SPUSH(reversed, current->data);
        current = current->next;
    }
    list.head = reversed.head;
}

void AsteroidCycle(Stack& aster){
    bool changed;
    do {
        changed = false;
        Stack temp;
        
        SNode* current = aster.head;
        while (current != nullptr) {
            if (temp.head != nullptr && 
                temp.head->data.second == "right" && 
                current->data.second == "left") {
                
                // Обработка столкновения
                if (temp.head->data.first == current->data.first) {
                    SPOP(temp);
                } else if (temp.head->data.first > current->data.first) {
                    temp.head->data.first -= current->data.first;
                } else {
                    current->data.first -= temp.head->data.first;
                    SPOP(temp);
                    SPUSH(temp, current->data);
                }
                changed = true;
            } else {
                SPUSH(temp, current->data);
            }
            current = current->next;
        }
        
        // Восстанавливаем исходный порядок
        aster.head = nullptr;
        current = temp.head;
        while (current != nullptr) {
            SPUSH(aster, current->data);
            current = current->next;
        }
        
    } while (changed);
}

void asteroid(){
    cout << "Введите начальные данные в формате:\n<масса> <направление (l/r)> stop" << endl;
    
    Stack stack;
    string input;
    
    cin.ignore();
    getline(cin, input);
    
    stringstream ss(input);
    string token;
    int mass = 0;
    string direction;
    bool expectMass = true;
    
    while(ss >> token){
        if (token == "stop"){
            break;
        }
        
        if (expectMass){
            mass = stoi(token);
            expectMass = false;
        } else {
            if (token == "l" || token == "r") {
                direction = (token == "l" ? "left" : "right");
                pair<int, string> ast = {mass, direction};
                SPUSH(stack, ast);
                expectMass = true;
            } else {
                cout << "Ошибка: направление должно быть 'l' или 'r', получено: " << token << endl;
                return;
            }
        }
    }
    
    // Обрабатываем астероиды
    ReverseStack(stack);
    AsteroidCycle(stack);
    cout << "Результат:" << endl;
    SPRINT(stack);
}