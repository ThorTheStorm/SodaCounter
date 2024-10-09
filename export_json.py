import json

def export_json(database_dir, data_input):
    # Convert dictionary to json
    json_output = json.dumps(data_input)

    try:
        # Export the json to file
        with open(database_dir, "w") as file:
            file.write(json_output)

        return bool(True)
    except FileNotFoundError as fileNotFound_Error:
        return (f"A 'FileNotFound' error: {fileNotFound_Error}")