# Import necessary modules
import os
import json
from pathlib import Path
import pprint

# Import functions
from import_json import import_json
from create_json import create_json
from export_json import export_json
from add_data import add_data
from user_confirm import user_confirm
from subtract_value import subtract_value
from add_value import add_value
from valid_dictionary import valid_dictionary
from menu import menu
from print_dict_values import print_dict_values

# The main function for the soda-counter
def main():
    # All variables defined for the main function specifically
    runningDirectory = os.getcwd()
    scriptDirectory = Path(__file__).resolve().parent
    databaseDirectory = Path(f"{scriptDirectory}\\sodas.json")
    user_exit = bool(False)

    ## Clear terminal for cleaner output
    os.system('cls')
    
    ## Menu-options for the function menu:
    options =  {
        "add_item" : "Add item to database",
        "remove_item" : "Remove item from database",
        "subtract_value" : "Subtract count from item",
        "add_value" : "Add count to item",
        "break" : "Exit",
    }

    # Import content from json
    try:
        if databaseDirectory.exists():
            database_data = import_json(databaseDirectory)
        elif not databaseDirectory.exists():
            create_json(databaseDirectory)
            database_data = {}  # Initialize database_data after creating the file
        else:
            raise Exception("An unknown error occurred")

    except FileNotFoundError as fileNotFound_Error:
        print (f"{fileNotFound_Error}")

    except json.JSONDecodeError as jsonDecode_Error:
        print (f"JSON Decode error: {jsonDecode_Error}")

    except Exception as unknown_Error:
        print (f"Import JSON: {unknown_Error}")

    # Verify contents
    try:
        # Check if there is data in the variable
        if not database_data:

            print (f"You have no sodas in your database.")

            user_input = user_confirm("Want to add some?")

            while (user_input == True):
                user_data_input = input("Add sodas to the database (separate with comma ','): ")
                addDatabaseElements = user_data_input.split(",")
                database_data = add_data(database_data, addDatabaseElements, 0)
                user_input = user_confirm("Want to add more?")

        while (user_exit == False):

            if not (valid_dictionary(database_data)):
                raise TypeError("Data in database_data doesn't seem to be a dictionary")

            else:
                # Print the current stock stored in the database
                print (f"Current stock:\n")
                print_dict_values(database_data,'Soda',25,'Count',5)
                print (f"\n\n")

                user_selected_option = dict(menu(options)) #select_option(int(input("\nInput number for option: ")))
                
                if user_selected_option == {"add_item": True}:
                    # Clear terminal for cleaner output
                    os.system('cls')

                    user_input = bool(True)

                    while (user_input == True):
                        user_data_input = input("Add sodas to the database (separate with comma ','): ")
                        addDatabaseElements = user_data_input.split(",")
                        database_data = add_data(database_data, addDatabaseElements, 0)
                        user_input = user_confirm("Want to add more?")

                elif user_selected_option == {"subtract_value": True}:
                #ser_subtract_input = user_confirm("\n\nWanna subtract from your inventory?")

                    # Clear terminal for cleaner output
                    os.system('cls')

                    print_dict_values(database_data,'Soda',25,'Count',5,True,10)
                    print (f"\n\n")
                    
                    user_input_item_to_subtract = int(input("Select index for soda: ")) - 1
                    user_input_amount_to_subtract = int(input("How many to subtract: "))
                    database_data = subtract_value(database_data,user_input_item_to_subtract,user_input_amount_to_subtract)

                    # Clear terminal for cleaner output
                    os.system('cls')

                    print (f"Updated stock:\n")
                    print(f"{'Soda':<25} {'Count':<5}")
                    print("-" * 30)
                    for key, value in database_data.items():
                        print(f"{key:<25} {value:<5}")
                
                elif user_selected_option == {"add_value": True}:
                #ser_subtract_input = user_confirm("\n\nWanna subtract from your inventory?")

                    # Clear terminal for cleaner output
                    os.system('cls')

                    print_dict_values(database_data,'Soda',25,'Count',5,True,10)
                    print (f"\n\n")
                    
                    user_input_item_to_add = int(input("Select index for soda: ")) - 1
                    user_input_amount_to_add = int(input("How many to add: "))
                    database_data = add_value(database_data,user_input_item_to_add,user_input_amount_to_add)

                    # Clear terminal for cleaner output
                    os.system('cls')

                    print (f"Updated stock:\n")
                    print(f"{'Soda':<25} {'Count':<5}")
                    print("-" * 30)
                    for key, value in database_data.items():
                        print(f"{key:<25} {value:<5}")

                elif user_selected_option == {"remove_item": True}:
                    print (f"To be added")
                
                elif user_selected_option == {"break": True}:
                    break
            
        #user_exit = user_confirm("Want to close the program?")
            # Export updated data to database_file when done
        export_json(databaseDirectory, database_data)
        print (f"Data successfully exported to {databaseDirectory}")

    # If the variable doesnt exist this exception raises and we continue from there
    except NameError as name_error:
        return (f"Seems to be an error with a variable: {name_error}")

    # In case of mis-input
    except ValueError as value_Error:
        print (f"Bad value: {value_Error}")

    # I don't know what dafuq happened
    except Exception as unknown_Error:
        print (f"An unknown error occurred (main): {unknown_Error}")

# If file is named "main", then run the function main()
if __name__ == "__main__":
    main()

# Clear terminal for cleaner output
os.system('cls')