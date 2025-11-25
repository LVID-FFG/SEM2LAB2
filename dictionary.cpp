#include <iostream> //cAnnot cannOt fOund pAge stop thE pAge cAnnot be found - ввод проверки
#include <algorithm>
#include <sstream>

using namespace std;

struct Dictionary{
    size_t size;
    size_t cap;
    string* data;
    Dictionary() : size(0), cap(10)
                , data(new string[10]){}
    void addWord(string str){
        data[size] = str;
        size++;
        if (size == cap){
            string* newdata = new string[2*cap];
            cap *= 2;
            for(int i = 0; i < size; i++) newdata[i] = data[i];
            delete data;
            data = newdata;
        }
    }

    int errors(string str){
        int result = 0;
        stringstream ss(str);
        while (ss >> str){
            string str_down = str;
            transform(str.begin(), str.end(), str_down.begin(), ::tolower); //нет ударений
            if (str_down == str){
                result++;
                continue;
            }

            bool manyUdr = false;
            for(size_t i = 0, udr = 0 ; i < str.size(); i++){
                    if (str[i] != str_down[i]) udr++;
                    if (udr > 1) {
                        manyUdr = true;
                        break;
                    }
                }
            if (manyUdr){
                result++;
                continue;
            }
            bool haveWord = false; //есть ли слово в принципе
            for(int i = 0; i < size; i++){
                string data_down = data[i];
                transform(data[i].begin(), data[i].end(), data_down.begin(), ::tolower);
                if (str_down == data_down){
                    haveWord = true;
                };
            }

            bool noError = false;
            for(int i = 0; i < size; i++){
                if (str == data[i]){
                    noError = true;
                };
            }

            if (haveWord && !(noError)) result++;
            
        }
        return result;
    }
    

    ~Dictionary(){
        delete[] data;
    }
};

ostream& operator<<(ostream& ss, Dictionary& dic){
    for(int i = 0; i < dic.size; i++) ss << dic.data[i] << " ";
    ss << endl;
    return ss;
}

void dictionary(){
    cout << "Введите словарь (слова через пробел), затем слово 'stop', затем строку для проверки:" << endl;
    
    Dictionary dict;
    string input;
    
    cin.ignore();
    getline(cin, input);
    
    stringstream ss(input);
    string usCin;
    
    while(ss >> usCin){
        if (usCin == "stop"){
            break;
        }
        dict.addWord(usCin);
    }
    
    string text_to_check;
    if (getline(ss, text_to_check)) {
        // Убираем лишние пробелы в начале
        text_to_check.erase(0, text_to_check.find_first_not_of(" "));
    }
    
    // Если в строке не осталось текста для проверки, читаем новую строку
    if (text_to_check.empty()) {
        cout << "Введите строку для проверки:" << endl;
        getline(cin, text_to_check);
    }
    
    // Проверяем ошибки
    int error_count = dict.errors(text_to_check);
    
    
    cout << "Словарь: " << dict;
    cout << "Проверяемая строка: " << text_to_check << endl;
    cout << "Количество ошибок: " << error_count << endl;
}