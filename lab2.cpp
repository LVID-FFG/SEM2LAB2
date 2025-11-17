//g++ lab2.cpp asteroid.cpp dictionary.cpp genom.cpp hafman.cpp hashTableChain.cpp hashTableFree.cpp morf.cpp LRU.cpp setInclude.cpp; ./a.out; rm a.out

#include <iostream>
#include <fstream>
#include <string>
#include <vector>

#include "asteroid.h"
#include "genom.h"
#include "dictionary.h"
#include "hafman.h"
#include "hashTableChain.h"
#include "hashTableFree.h"
#include "morf.h"
#include "LRU.h"
using namespace std;



int main() {
   enum class command{
        EXIT = 0,
        ASTEROID,
        GENOME,
        DICTIONARY,
        HAFMAN,
        OPENHASH,
        CHAINHASH,
        MORF,
        LRUCASH,
        ERROR
    };
    while (1){
        cout << "Доступные программы: " << endl;
        cout << "Выход - EXIT" << endl;
        cout << "ASTEROID" << endl;
        cout << "GENOME" << endl;
        cout << "DICTIONARY" << endl;
        cout << "HAFMAN" << endl;
        cout << "OPENHASH" << endl;
        cout << "CHAINHASH" << endl;
        cout << "MORF" << endl;
        cout << "LRUCASH" << endl;
        string usCin;
        cin >> usCin;
        enum command command;
        if (usCin == "EXIT"){
            command = command::EXIT;
        }
        if (usCin == "ASTEROID"){
            command = command::ASTEROID;
        }
        if (usCin == "GENOME"){
            command = command::GENOME;
        }
        if (usCin == "DICTIONARY"){
            command = command::DICTIONARY;
        }
        if (usCin == "HAFMAN"){
            command = command::HAFMAN;
        }
        if (usCin == "OPENHASH"){
            command = command::OPENHASH;
        }
        if (usCin == "CHAINHASH"){
            command = command::CHAINHASH;
        }
        if (usCin == "MORF"){
            command = command::MORF;
        }
        if (usCin == "LRUCASH"){
            command = command::LRUCASH;
        }

        switch (command){
            case command::EXIT:
                cout << "Выход из программы" << endl;
                exit(0);
            case command::ASTEROID:
                cout << endl << "Запуск задания №1" << endl;
                asteroid();
                break;
            case command::GENOME:
                cout << endl << "Запуск задания №3" << endl;
                genome();
                break;
            case command::DICTIONARY:
                cout << endl << "Запуск задания №4" << endl;
                dictionary();
                break;
            case command::HAFMAN:
                cout << endl << "Запуск задания №5" << endl;
                hafman();
                break;
            case command::OPENHASH:
                cout << endl << "Запуск задания №6" << endl;
                hashTableFree();
                break;
            case command::CHAINHASH:
                cout << endl << "Запуск задания №7" << endl;
                hashTableChain();
                break;
            case command::MORF:
                cout << endl << "Запуск задания №8" << endl;
                morf();
                break;
            case command::LRUCASH:
                cout << endl << "Запуск задания №9" << endl;
                LRU();
                break;
            default:
                cout << "Ошибка ввода!" << endl;
                exit(1);
        }
        command = command::ERROR;
        cout << endl;
    }
}