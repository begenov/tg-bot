package handler

/*
func jobSeekerHandler(update tgbotapi.Update) {
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
		// Set user role to job seeker
		user = &db.User{ChatID: chatID, Role: "job seeker"}
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
		profile, err := db.GetJobSeekerProfile(user.ID)
		if err != nil {
			log.Printf("Error getting job seeker profile: %s", err)
			sendMessage(chatID, "An error occurred. Please try again later.")
			return
		}
		sendMessage(chatID, formatProfile(profile))
	case "/search":
		// Perform a job search and send the results to the user
		results, err := searchJobs(user)
		if err != nil {
			log.Printf("Error searching for jobs: %s", err)
			sendMessage(chatID, "An error occurred. Please try again later.")
			return
		}
		sendMessage(chatID, formatSearchResults(results))
	default:
		// If the message does not match any command, prompt the user with a list of available commands
		sendMessage(chatID, "Invalid command. Available commands: /profile, /search")
	}
}
*/
