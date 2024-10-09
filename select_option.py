def select_option(user_selected_input):
    match user_selected_input:
        case 1:
            return {"add_item": True}
        case 2: 
            return {"subtract_value": True}
        case 3:
            return {"add_value": True}
        case 4:
            return {"remove_item": True}
        case 5:
            return {"break": True}