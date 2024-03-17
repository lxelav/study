#include <iostream>
#include <map>
#include <fstream>
#include <vector>

using namespace std;

struct Variables {
    map<string, int> var;
};

void showVar(vector<Variables> &area);

int main(int argc, char* argv[]) {
    for (int i = 1; i < argc; ++i) {
        vector<Variables> area;

        ifstream file_input("G:\\CLion 2023.3.3\\projects\\laba3\\labav3_1\\" + (string)argv[i]);
        if (!file_input.is_open()) {
            cerr << "Error: file not open(((" << endl;
            return 1;
        }

        string command;
        while (getline(file_input, command)) {
            //Теперь нужно обрабатывать каждую строку
            if (command == "{") {
                area.push_back(Variables());
            }
            else if (command == "}") {
                area.pop_back();
            }
            else if (command == "ShowVar;") {
                showVar(area);
            }
            else {
                string name_var, value_var;
                bool flag = false;

                for (auto symbol : command) {
                    if (symbol == '=') flag = true;
                    else if (!flag) name_var.push_back(symbol);
                    else value_var.push_back(symbol);
                }

                size_t len_area = area.size();
                area[len_area - 1].var.insert({name_var, stoi(value_var)});
            }
        }
    }

}


void showVar(vector<Variables> &area) {
    map<string, int> result_var;

    for (const auto &scope : area) {
        for (const auto &variable : scope.var) {
            result_var[variable.first] = variable.second;
        }
    }

    cout << "Visible variables:" << endl;
    for (const auto& variable : result_var) {
        cout << variable.first << " = " << variable.second << endl;
    }
}