package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"fmt"
	"strconv"
	"errors"
)

// commands
var (
	commandStart = "start"
	commandHelp = "help"
	commandGetID = "get_id"

	// admin commands
	commandMailing = "mailing"
	commandAdminAdd = "add_admin"
	commandAdminDelete = "delete_admin"
	commandGetUserHistory = "get_history_by_tg"
	commandGetAllUsers = "get_users"
	commandGetUser = "get_user"
	commandAdminCancel = "cancel"

	// user commands
	commandCheckIP = "check_ip"
	commandGetHistory = "history"	// внутри как get_history_by_tg
	commandUserCancel = "cancel"
	commandBecomeAdmin = "become_admin" // debug
)

var Adminhelp = fmt.Sprintf(
	`/%s -> show all commands 
	/%s -> get your id
	/%s -> mailing all users 
	/%s -> add admin 
	/%s -> delete admin 
	/%s -> request history (need user id) 
	/%s -> info about all users 
	/%s -> info about user
	/%s -> cancel operation`, 
	commandHelp, commandGetID, commandMailing, commandAdminAdd, commandAdminDelete,
	commandGetUserHistory, commandGetAllUsers, commandGetUser, commandAdminCancel,
)

var Userhelp = fmt.Sprintf(
	`/%s -> show all commands
	/%s -> get your id
	/%s -> check ip
	/%s -> all previous ip
	/%s -> become an admin (debug)
	/%s -> cancel operation`,
	commandHelp, commandGetID, commandCheckIP, commandGetHistory, 
	commandBecomeAdmin, commandUserCancel,
)

// обработка команд
func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	isadmin, err := b.service.CheckAdminExists(message.From.ID)
	if err != nil {
		return err
	}

	if isadmin {
		switch message.Command() {
		case commandStart:
			return b.handleAdminStartCommand(message)
			
		case commandHelp:
			return b.handleAdminHelpCommand(message)
 
		case commandGetID:
			return b.handleGetIDCommand(message)

		case commandMailing:
			return b.handleMailingCommand(message)

		case commandAdminAdd:
			return b.handleAdminAddCommand(message)

		case commandAdminDelete:
			return b.handleAdminDeleteCommand(message)
			
		case commandGetUserHistory:
			return b.handleGetUserHistoryCommand(message)
		
		case commandGetAllUsers:
			return b.handleGetAllUsersCommand(message)

		case commandGetUser:
			return b.handleGetUserCommand(message)

		case commandAdminCancel:
			return b.handleAdminCancelCommand(message)

		default:
			return b.handleUnknownCommand(message)
		}

	} else {	// non-admin
		switch message.Command() {
		case commandStart:
			return b.handleUserStartCommand(message)
		
		case commandHelp:
			return b.handleUserHelpCommand(message)

		case commandGetID:
			return b.handleGetIDCommand(message)

		case commandCheckIP:
			return b.handleCheckIPCommand(message)

		case commandGetHistory:
			return b.handleGetHistoryCommand(message)

		case commandUserCancel:
			return b.handleUserCancelCommand(message)

		case commandBecomeAdmin:
			return b.handleBecomeAdminCommand(message)	// debug

		default:
			return b.handleUnknownCommand(message)
		}
	}
	
	return nil
}

// обработка сообщений
func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	isadmin, err := b.service.CheckAdminExists(message.From.ID)
	if err != nil {
		return err
	}

	if isadmin {
		state, err := b.getAdminState(message.From.ID)
		if err != nil {		
			return err
		}

		switch state {
		case "default_admin_state":
			return b.handleDefaultMessage(message)

		case "request_mailing":
			return b.handleMailingMessage(message)

		case "request_adminID_create":
			return b.handleAdminIdCreateMessage(message)

		case "request_adminID_delete":
			return b.handleAdminIdDeleteMessage(message)

		case "request_userID_history":
			return b.handleUserIdHistoryMessage(message)

		case "request_userID_info":
			return b.handleUserIdInfoMessage(message)

		default:
			return errors.New("internal error")
		}
	} else {
		state, err := b.getUserState(message.From.ID)
		if err != nil {		
			return err
		}

		switch state {
		case "default_user_state":
			return b.handleDefaultMessage(message)
			
		case "request_ip":
			return b.handleRequestIPMessage(message)

		default:
			return errors.New("internal error")
		}
	}
	
	return nil
}

func (b *Bot) handleGetIDCommand(message *tgbotapi.Message) error {
	b.SendMessage(message.Chat.ID, strconv.Itoa(message.From.ID))
	return nil
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	b.SendMessage(message.Chat.ID, "unknown command")
	return nil
}

func (b *Bot) handleDefaultMessage(message *tgbotapi.Message) error {
	b.SendMessage(message.Chat.ID, "no command was entered")
	return nil
}