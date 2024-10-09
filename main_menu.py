def main_menu(options):
    # Let user select what they want to do
    for index, (key, value) in enumerate(options.items()):
        # Index of item +1
        item_index = index + 1

        print (f"{item_index}: {value}")

    user_input = int(input("\nChoose option: "))

    selected_option = list(options.keys())[(user_input - 1)]

    return({selected_option:True})