#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <map>
#include <random>

using namespace std;

struct Array {
    vector<int> el;
};

void search_command(vector<string> *param, map<string, Array> *arrays);
void loadArray(string name_array, string file_name, map<string, Array> *arrays);
void saveArray(string name_array, string file_name, map<string, Array> *arrays);
void randArray(string name_array, string count, string lb, string rb, map<string, Array> *arrays);
void concatArray(string array_to_save, string name_array2, map<string, Array> *arrays);
void freeArray(string name_array, map<string, Array> *arrays);
void removeElements(string name_array, string start, string count, map<string, Array> *arrays);
void copyElements(string name_array, string start, string end, string array_save, map<string, Array> *arrays);
void sortArray(string param, map<string, Array> *arrays);
void shuffleArray(string name_array, map<string, Array> *arrays);
void statsArray(string name_array, map<string, Array> *arrays);
void printArray(string name_array, string option, map<string, Array> *arrays);
void printArrayRange(string name_array, string start, string end, map<string, Array> *arrays);

int main() {
    map<string, Array> arrays;

    //Открываем файл для чтения
    ifstream file_input;
    file_input.open("G:\\CLion 2023.3.3\\projects\\laba_array\\commands.txt");
    if (!file_input.is_open()) {
        cerr << "Unable to open file!" << endl;
        return 1;
    }

    string command;
    while (getline(file_input, command)) {
        string temp;
        vector<string> param;

        size_t size_command = command.length();
        for (int i = 0; i < size_command; i++) {
            if (command[i] == ';') {
                if (!temp.empty()){
                    param.push_back(temp);
                    temp.clear();
                }
            }
            else if (command[i] == ' ' || command[i] == ',' || command[i] == '(' || command[i] == ')') {
                if (!temp.empty()){
                    param.push_back(temp);
                    temp.clear();
                }
            }
            else {
                temp.push_back((char)tolower(command[i]));
            }
        }

        search_command(&param, &arrays);
    }

    file_input.close();
}


void search_command(vector<string> *param, map<string, Array> *arrays) {
    string name_command = (*param)[0];

    if (name_command == "load") {
        loadArray((*param)[1], (*param)[2], arrays);
    }
    else if (name_command == "save") {
        saveArray((*param)[1], (*param)[2], arrays);
    }
    else if (name_command == "rand") {
        randArray((*param)[1], (*param)[2], (*param)[3], (*param)[4], arrays);
    }
    else if (name_command == "concat") {
        concatArray((*param)[1], (*param)[2], arrays);
    }
    else if (name_command == "free") {
        freeArray((*param)[1], arrays);
    }
    else if (name_command == "remove") {
        removeElements((*param)[1], (*param)[2], (*param)[3], arrays);
    }
    else if (name_command == "copy") {
        copyElements((*param)[1], (*param)[2], (*param)[3], (*param)[4], arrays);
    }
    else if (name_command == "sort") {
        sortArray((*param)[1], arrays);
    }
    else if (name_command == "shuffle") {
        shuffleArray((*param)[1], arrays);
    }
    else if (name_command == "stats") {
        statsArray((*param)[1], arrays);
    }
    else if (name_command == "print") {
        if (((*param).size() - 1) == 2) {
            printArray((*param)[1], (*param)[2], arrays);
        }
        else if (((*param).size() - 1) == 3) {
            printArrayRange((*param)[1], (*param)[2], (*param)[3], arrays);
        }
    }
    else {
        cout << "Invalid command!" << endl;
    }

}

void loadArray(string name_array, string file_name, map<string, Array> *arrays) {
    ifstream file("G:\\CLion 2023.3.3\\projects\\laba_array\\" + file_name);
    if (!file.is_open()) {
        std::cerr << "Error: not open file" << file_name << std::endl;
        return;
    }

    (*arrays)[name_array].el.clear(); // Очищаем массив перед загрузкой

    int num;
    while (file >> num) {
        (*arrays)[name_array].el.push_back(num);
    }

    std::cout << "Array " << name_array << " load from file " << file_name << std::endl;
}

void saveArray(string name_array, string file_name, map<string, Array> *arrays) {
    ofstream output_file("G:\\CLion 2023.3.3\\projects\\" + file_name);


    if (!output_file.is_open()) {
        std::cerr << "Error: not open file " << file_name << std::endl;
        return;
    }

    vector<int> array = (*arrays)[name_array].el;
    for (auto element : array) {
        output_file << element << endl;
    }

    std::cout << "Array " << name_array << " save on file " << file_name << std::endl;
}

void randArray(string name_array, string count, string lb, string rb, map<string, Array> *arrays) {
    int count_int = stoi(count), lb_int = stoi(lb), rb_int = stoi(rb);

    random_device rd;
    mt19937 gen(rd());
    uniform_int_distribution<int> dist(lb_int, rb_int);

    (*arrays)[name_array].el.clear();
    for (int i = 0; i < count_int; ++i) {
        (*arrays)[name_array].el.push_back(dist(gen));
    }

    std::cout << "Array " << name_array << " adding random integer " << std::endl;
}

void concatArray(string array_to_save, string name_array2, map<string, Array> *arrays) {
    vector<int> *array_save = &(*arrays)[array_to_save].el;
    vector<int> array2 = (*arrays)[name_array2].el;

    for (auto element : array2) {
        (*array_save).push_back(element);
    }

    std::cout << "Array " << array_to_save << " and " << name_array2 << " union on " << array_to_save << std::endl;
}

void freeArray(string name_array, map<string, Array> *arrays) {
    (*arrays)[name_array].el.clear();

    cout << "Array " << name_array << " clear " << endl;
}

void removeElements(string name_array, string start, string count, map<string, Array> *arrays) {
    int start_int = stoi(start), count_int = stoi(count);
    vector<int> *array = &(*arrays)[name_array].el;

    if (start_int < 0 || start_int >= (*array).size()) {
        cerr << "Error: index" << endl;
        return;
    }

    for (int j = 0; j < count_int; j++) {
        (*array).erase((*array).begin() + start_int);
    }

    cout << "RemoveElements - true" << endl;
}

void copyElements(string name_array, string start, string end, string array_save, map<string, Array> *arrays) {
    int start_int = stoi(start), end_int = stoi(end);
    vector<int> array1 = (*arrays)[name_array].el;
    vector<int> *array_ToSave = &(*arrays)[array_save].el;

    if (start_int < 0 || start_int >= array1.size() || end_int < 0 || end_int >= array1.size()) {
        cerr << "Error index" << endl;
        return;
    }

    for (int i = start_int; i <= end_int; i++) {
        (*array_ToSave).push_back(array1[i]);
    }

    cout << "CopyElements - true" << endl;
}

void quickSort(vector <int> &arr, int left, int right, char option) {
    int i = left, j = right;
    int tmp;
    int pivot = arr[(left + right) / 2];

    while (i <= j) {
        if (option == '+'){
            while (arr[i] < pivot)
                i++;
            while (arr[j] > pivot)
                j--;
        }
        else if (option == '-') {
            while (arr[i] > pivot)
                i++;
            while (arr[j] < pivot)
                j--;
        }

        if (i <= j) {
            tmp = arr[i];
            arr[i] = arr[j];
            arr[j] = tmp;
            i++;
            j--;
        }
    }


    if (left < j)
        quickSort(arr, left, j, option);
    if (i < right)
        quickSort(arr, i, right, option);

}

void sortArray(string param, map<string, Array> *arrays) {
    string name_array;
    char option = param[1];
    name_array.push_back(param[0]);

    quickSort((*arrays)[name_array].el, 0, (*arrays)[name_array].el.size(), option);

    cout << "SortArray - true" << endl;
}

void quickSortShuffle(std::vector<int>& arr, int left, int right) {
    if (left < right) {
        int pivotIndex = left + (right - left) / 2;
        int pivot = arr[pivotIndex];

        // Разделение на две части
        int i = left, j = right;
        while (i <= j) {
            while (arr[i] < pivot) i++;
            while (arr[j] > pivot) j--;
            if (i <= j) {
                std::swap(arr[i], arr[j]);
                i++;
                j--;
            }
        }

        // Рекурсивная сортировка левой и правой части
        quickSortShuffle(arr, left, j);
        quickSortShuffle(arr, i, right);
    }
}

void shuffleArray(string name_array, map<string, Array> *arrays) {
    quickSortShuffle((*arrays)[name_array].el, 0, (*arrays)[name_array].el.size());

    cout << "ShuffleArray - true" << endl;
}

void statsArray(string name_array, map<string, Array> *arrays) {
    /*Взападло писать
     *
     * вывести в стандартный поток вывода статистическую информацию о массиве
    A: размер массива, максимальный и минимальный элемент (и их индексы), наиболее
    часто встречающийся элемент (если таковых несколько, вывести максимальный из них
    по значению), среднее значение элементов, максимальное из значений отклонений
    элементов от среднего значения;*/

    cout << "TODO : statsArray" << endl;
    // TODO
}

void printArray(string name_array, string option, map<string, Array> *arrays) {
    if (option == "all") {
        vector<int> array = (*arrays)[name_array].el;

        for (auto element : array) {
            cout << element << " ";
        }
        cout << endl;
    }
    else {
        int el_index = stoi(option);

        cout << "Element - " << (*arrays)[name_array].el[el_index] << endl;
    }
}

void printArrayRange(string name_array, string start, string end, map<string, Array> *arrays) {
    int start_index = stoi(start), end_index = stoi(end);
    vector<int> array = (*arrays)[name_array].el;

    if (start_index < 0 || start_index >= array.size() || end_index < 0 || end_index >= array.size()) {
        cerr << "Error index" << endl;
        return;
    }
    cout << "PrintRange ";
    for (int i = start_index; i < end_index; i++) {
        cout << array[i] << " ";
    }
    cout << endl;

}