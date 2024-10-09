import json
from pathlib import Path

def create_json(destination):

    # Import necessary components
    from user_confirm import user_confirm

    # Create the JSON if it doesn't exist
    createDatabaseFile = user_confirm("Create database?")

    if createDatabaseFile == True:
        Path(f"{destination}").touch()
        with open(destination, 'w') as database: # Open data-file. 'r' is for read-only
            database.write("{}")
        return (f"Created database: {destination}")
    elif createDatabaseFile == False:
        print (f"Chosen not to create database. Exiting...")
        exit
    else:
        raise Exception("Non-boolean value in variable createDatabaseFile")