def subtract_value(database,item_index,value_to_subtract):
    try:
        # Convert index to key
        keys = list(database.keys())
        key = keys[item_index]

        item_current_value = database[key]
        database[key] = item_current_value - value_to_subtract
        return database
    except Exception as unknown_error:
        return Exception(f"Unknown error (subtract_value): ", unknown_error)