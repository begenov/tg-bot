package handler

/*
func employerHandler(update tgbotapi.Update) {
    message := update.Message
    chatID := message.Chat.ID

    // Check if user already exists in database
    user, err := db.GetUserByChatID(chatID)
    if err != nil {
        log.Printf("Error getting user: %s", err)
        sendMessage(chatID, "An error occurred. Please try again later.")
        return
    }

    // If user does not exist, prompt them to provide their information
    if user == nil {
        sendMessage(chatID, "Please provide your name, age, and gender:")
        // Set user role to employer
        user = &db.User{ChatID: chatID, Role: "employer"}
        err = db.CreateUser(user)
        if err != nil {
            log.Printf("Error creating user: %s", err)
            sendMessage(chatID, "An error occurred. Please try again later.")
            return
        }
        return
    }

    // If user already exists, handle message based on its content
    switch message.Text {
    case "/profile":
        // Retrieve user profile from database and send it to the user
        profile, err := db.GetEmployerProfile(user.ID)
        if err != nil {
            log.Printf("Error getting employer profile: %s", err)
            sendMessage(chatID, "An error occurred. Please try again later.")
            return
        }
        sendMessage(chatID, formatProfile(profile))
    case "/post_job":
        // Prompt the user to provide job details
        sendMessage(chatID, "Please provide job details:")
        // Set user state to "posting_job"
        err = db.SetUserState(user.ID, "posting_job")
        if err != nil {
            log.Printf("Error setting user state: %s", err)
            sendMessage(chatID, "An error occurred. Please try again later.")
            return
        }
    default:
        // If the message does not match any command, prompt the user with a list of available commands
        sendMessage(chatID, "Invalid command. Available commands: /profile, /post_job")
    }
}

*/
