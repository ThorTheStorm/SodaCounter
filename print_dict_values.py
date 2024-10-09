def print_dict_values(data,key_name,key_width,value_name,value_width,index = False,index_width = 0):
    if index:
        index_name = str("Index")
        print(f"{index_name:<{index_width}} {key_name:<{key_width}} {value_name:<{value_width}}")
        print("-" * (index_width + key_width + value_width))
    else:
        print(f"{key_name:<{key_width}} {value_name:<{value_width}}")
        print("-" * (key_width + value_width))

    for index_value, (key, value) in enumerate(data.items(), 1):
        if index:
            print(f"{index_value:<{index_width}} {key:<{key_width}} {value:<{value_width}}")
        else:
            print(f"{key:<{key_width}} {value:<{value_width}}")