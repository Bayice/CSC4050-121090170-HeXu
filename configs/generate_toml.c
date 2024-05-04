#include <stdio.h>
#include <stdlib.h>
#include <time.h>

// 函数类型枚举
enum BpfFuncType {
    BpfFuncTypeKprobeTime,
    BpfFuncTypeKprobeCount,
    BpfFuncTypeKprobeSampleArgs,
    BpfFuncTypeKprobeSampleReturn,
    BpfFuncTypeUprobeTime,
    BpfFuncTypeUprobeCount,
    BpfFuncTypeUprobeSampleArgs,
    BpfFuncTypeUprobeSampleReturn
};

// 地图类型枚举
enum BpfMapType {
    BpfMapTypeArray,
    BpfMapTypeHashOfMaps,
    BpfMapTypeArrayOfMaps,
    BpfMapTypeQueue,
    BpfMapTypeStack,
    BpfMapTypePercpuHash,
    BpfMapTypeLruHash,
    BpfMapTypeLruPercpuHash
};

// 随机生成指定范围内的整数
int generate_random(int min, int max) {
    return rand() % (max - min + 1) + min;
}

// 随机选择函数类型
enum BpfFuncType random_function_type() {
    return generate_random(0, 7); // 0-7 是 BpfFuncType 枚举的范围
}

// 随机选择地图类型
enum BpfMapType random_map_type() {
    return generate_random(0, 7); // 0-7 是 BpfMapType 枚举的范围
}

// 随机生成函数和结构体名称
void generate_random_names(char *name, int length) {
    const char *characters = "abcdefghijklmnopqrstuvwxyz0123456789";
    for (int i = 0; i < length; ++i) {
        name[i] = characters[rand() % 36];
    }
    name[length] = '\0';
}

int main() {
    srand(time(NULL)); // 初始化随机数生成器

    FILE *file = fopen("df.toml", "w");
    if (file == NULL) {
        perror("Error opening file");
        return 1;
    }

    // 输出 TOML 配置文件头部
    fprintf(file, "[EBPFProgram]\n");
    fprintf(file, "ObjectBinary = \"path/to/compiled/ebpf/program.o\"\n\n");

    // 随机生成 1 到 4 个函数和结构体
    int num_functions = generate_random(1, 4);
    for (int i = 0; i < num_functions; ++i) {
        char function_name[10];
        generate_random_names(function_name, 8);
        enum BpfFuncType function_type = random_function_type();

        // 输出随机生成的函数配置
        fprintf(file, "[[EBPFProgram.Functions]]\n");
        fprintf(file, "Name = \"%s\"\n", function_name);
        fprintf(file, "Aim = \"%s\"\n\n", 
               (function_type == BpfFuncTypeKprobeTime) ? "BpfFuncTypeKprobeTime" :
               (function_type == BpfFuncTypeKprobeCount) ? "BpfFuncTypeKprobeCount" :
               (function_type == BpfFuncTypeKprobeSampleArgs) ? "BpfFuncTypeKprobeSampleArgs" :
               (function_type == BpfFuncTypeKprobeSampleReturn) ? "BpfFuncTypeKprobeSampleReturn" :
               (function_type == BpfFuncTypeUprobeTime) ? "BpfFuncTypeUprobeTime" :
               (function_type == BpfFuncTypeUprobeCount) ? "BpfFuncTypeUprobeCount" :
               (function_type == BpfFuncTypeUprobeSampleArgs) ? "BpfFuncTypeUprobeSampleArgs" :
               "BpfFuncTypeUprobeSampleReturn");

        // 输出随机生成的结构体配置
        fprintf(file, "[[EBPFProgram.Structs]]\n");
        fprintf(file, "Name = \"%s\"\n", function_name);
        fprintf(file, "ForTransmission = true\n\n");
    }

    // 随机生成 1 到 4 个地图配置
    int num_maps = generate_random(1, 4);
    for (int i = 0; i < num_maps; ++i) {
        char map_name[10];
        generate_random_names(map_name, 8);
        enum BpfMapType map_type = random_map_type();

        // 输出随机生成的地图配置
        fprintf(file, "[[EBPFProgram.Maps]]\n");
        fprintf(file, "Type = \"%s\"\n", 
               (map_type == BpfMapTypeArray) ? "BpfMapTypeArray" :
               (map_type == BpfMapTypeHashOfMaps) ? "BpfMapTypeHashOfMaps" :
               (map_type == BpfMapTypeArrayOfMaps) ? "BpfMapTypeArrayOfMaps" :
               (map_type == BpfMapTypeQueue) ? "BpfMapTypeQueue" :
               (map_type == BpfMapTypeStack) ? "BpfMapTypeStack" :
               (map_type == BpfMapTypePercpuHash) ? "BpfMapTypePercpuHash" :
               (map_type == BpfMapTypeLruHash) ? "BpfMapTypeLruHash" :
               "BpfMapTypeLruPercpuHash");
        fprintf(file, "Key = \"u32\"\n");
        fprintf(file, "Value = \"u64\"\n");
        fprintf(file, "MaxEntries = %d\n", generate_random(10, 1000));
        fprintf(file, "Name = \"%s\"\n\n", map_name);
    }

    fclose(file); // 关闭文件

    return 0;
}
