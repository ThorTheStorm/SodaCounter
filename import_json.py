import json

def import_json(source):
    try:
        with open(source, 'r') as database: # Open data-file. 'r' is for read-only
            database_data = json.load(database)

        return database_data
        
        if not database_data:
            raise Exception(f"No data in database. You should consider adding something")

    except FileNotFoundError as fileNotFound_Error:
        return FileNotFoundError(f"A 'FileNotFound' error: {fileNotFound_Error}")

    except json.JSONDecodeError as jsonDecode_Error:
        return json.JSONDecodeError(f"An error occurred decoding the json: {jsonDecode_Error}")

    except Exception as unknown_Error:
        return Exception(f"An unknown error occurred: {unknown_Error}")