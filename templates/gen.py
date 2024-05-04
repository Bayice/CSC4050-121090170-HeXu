import os

def generate_template_files(base_name, custom_string, output_dir):
    # 定义要生成的文件名列表
    file_names = [
        f"{base_name}.gtpl",
    ]
    os.makedirs(output_dir, exist_ok=True)

    # 遍历文件名列表，创建空文件
    for file_name in file_names:
        # 拼接文件完整路径
        full_path = os.path.join(output_dir, file_name)
        # 创建空文件
        with open(full_path, "w") as f:
            f.write("")  # 写入空内容

        print(f"Created file: {full_path}")


	# BpfMapTypeArray         MapType = "BPF_MAP_TYPE_ARRAY"
	# BpfMapTypeHashOfMaps    MapType = "BPF_MAP_TYPE_HASH_OF_MAPS"
	# BpfMapTypeArrayOfMaps   MapType = "BPF_MAP_TYPE_ARRAY_OF_MAPS"
	# BpfMapTypeQueue         MapType = "BPF_MAP_TYPE_QUEUE"
	# BpfMapTypeStack         MapType = "BPF_MAP_TYPE_STACK"
	# BpfMapTypePercpuHash    MapType = "BPF_MAP_TYPE_PERCPU_HASH"
	# BpfMapTypeLruHash       MapType = "BPF_MAP_TYPE_LRU_HASH"
	# BpfMapTypeLruPercpuHash MapType = "BPF_MAP_TYPE_LRU_PERCPU_HASH"



# 示例调用
custom_string = ["BPF_MAP_TYPE_ARRAY","BPF_MAP_TYPE_HASH_OF_MAPS","BPF_MAP_TYPE_ARRAY_OF_MAPS","BPF_MAP_TYPE_QUEUE", "BPF_MAP_TYPE_STACK",   "BPF_MAP_TYPE_PERCPU_HASH",   "BPF_MAP_TYPE_LRU_HASH" ,"BPF_MAP_TYPE_LRU_PERCPU_HASH"]
output_directory = custom_string

# 调用生成函数
for i in custom_string:
    generate_template_files(i, i, i)
