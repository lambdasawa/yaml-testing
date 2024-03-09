#include <iostream>
#include <fstream>
#include "yaml-cpp/yaml.h"

int main(int argc, char* argv[]) {
    std::string path = argv[1];
    YAML::Node obj = YAML::LoadFile(path);
    std::cout << obj << "\n";
    return 0;
}
