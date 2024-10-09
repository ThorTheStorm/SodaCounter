def add_data(database, input_data, amount = None):
    try:
        for item in input_data:
            if item in database:
                raise ValueError("Value already exists")
            else:
                database[item] = amount if amount is not None else 0
        return database
    except ValueError as value_error:
        return (f"Value error (add_data): {value_error}")
    except Exception as unknown_error:
        return (f"An unknown error occurred (add_data): {unknown_error}")