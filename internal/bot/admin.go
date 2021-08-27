package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"strconv"
	"errors"
	"fmt"
)

func (b *Bot) changeAdminState(id int, state string) error {
	if err := b.service.SetAdminState(id, state); err != nil {
		return err
	}
	return nil
}

func (b *Bot) getAdminState(id int) (string, error) {
	state, err := b.service.GetAdminState(id);
	if err != nil {
		return "", err
	}
	return state, nil
}

func (b *Bot) createAdmin(id int) (error) {
	if err := b.service.CreateAdmin(id, "default_admin_state"); err != nil {
		return err
	} 
	return nil
} 

func (b *Bot) validateID(input string) (int, error) {
	value, err := strconv.ParseInt(input, 0, 64)
	if err != nil {
		return -1, err
	}

	return int(value), nil
}

// command handlers
func (b *Bot) handleAdminStartCommand(message *tgbotapi.Message) error {
	if err := b.changeAdminState(message.From.ID, "default_admin_state"); err != nil {
		return err
	}

	b.SendMessage(message.Chat.ID, "admin mode")
	return b.handleAdminHelpCommand(message)
}

func (b *Bot) handleAdminHelpCommand(message *tgbotapi.Message) error {
	b.SendMessage(message.Chat.ID, Adminhelp)
	return nil
}

func (b *Bot) handleMailingCommand(message *tgbotapi.Message) error {
	if err := b.changeAdminState(message.From.ID, "request_mailing"); err != nil {
		return err
	}

	b.SendMessage(message.Chat.ID, "input message for mailing")
	return nil
}

func (b *Bot) handleAdminAddCommand(message *tgbotapi.Message) error {
	if err := b.changeAdminState(message.From.ID, "request_adminID_create"); err != nil {
		return err
	}

	b.SendMessage(message.Chat.ID, "input id for adding new admin")
	return nil
}

func (b *Bot) handleAdminDeleteCommand(message *tgbotapi.Message) error {
	if err := b.changeAdminState(message.From.ID, "request_adminID_delete"); err != nil {
		return err
	}

	b.SendMessage(message.Chat.ID, "input id for deleting admin")
	return nil
}

func (b *Bot) handleGetUserHistoryCommand(message *tgbotapi.Message) error {
	if err := b.changeAdminState(message.From.ID, "request_userID_history"); err != nil {
		return err
	}

	b.SendMessage(message.Chat.ID, "input user id")
	return nil
}

func (b *Bot) handleGetAllUsersCommand(message *tgbotapi.Message) error {
	info, err := b.service.GetAllUsersInfo()
	if err != nil {
		return err
	}

	for _, user := range info {
		b.SendMessage(message.Chat.ID, user)
	}
	return nil
}

func (b *Bot) handleGetUserCommand(message *tgbotapi.Message) error {
	if err := b.changeAdminState(message.From.ID, "request_userID_info"); err != nil {
		return err
	}

	b.SendMessage(message.Chat.ID, "input user id")
	return nil
}

func (b *Bot) handleAdminCancelCommand(message *tgbotapi.Message) error {
	if err := b.changeAdminState(message.From.ID, "default_admin_state"); err != nil {
		return err
	}

	b.SendMessage(message.Chat.ID, "operation cancelled")
	return nil	
}

// message handlers
func (b *Bot) handleMailingMessage(message *tgbotapi.Message) error {
	chats, err := b.service.GetAllUsersChatID()
	if err != nil {
		return err
	}

	for _, chatID := range chats {
		b.SendMessage(chatID, message.Text)
	}

	if err := b.changeAdminState(message.From.ID, "default_admin_state"); err != nil {
		return err
	}	

	return nil
}

func (b *Bot) handleAdminIdCreateMessage(message *tgbotapi.Message) error {
	id, err := b.validateID(message.Text)
	if err != nil {
		return errors.New("invalid id")
	}

	if err := b.createAdmin(id); err != nil {
		return err
	}

	b.SendMessage(message.Chat.ID, "new admin created")
	
	if err := b.changeAdminState(message.From.ID, "default_admin_state"); err != nil {
		return err
	}	

	return nil
}

func (b *Bot) handleAdminIdDeleteMessage(message *tgbotapi.Message) error {
	id, err := b.validateID(message.Text)
	if err != nil {
		return errors.New("invalid id")
	}

	if err := b.service.DeleteAdmin(id); err != nil {
		return err
	} 

	b.SendMessage(message.Chat.ID, "admin deleted")
	
	if id != message.From.ID { // удалили себя -> некому менять state 
		if err := b.changeAdminState(message.From.ID, "default_admin_state"); err != nil {
			return err
		}	
	} else {
		return b.handleUserStartCommand(message)
	}

	return nil
}

func (b *Bot) handleUserIdHistoryMessage(message *tgbotapi.Message) error {
	id, err := b.validateID(message.Text)
	if err != nil {
		return errors.New("invalid id")
	}

	info, err := b.service.GetUserUniqueIPsExt(id)
	if err != nil {
		return err
	}

	for ip, log := range info {
		b.SendMessage(message.Chat.ID, fmt.Sprintf("%s : %s", ip, log))
	}

	if err := b.changeAdminState(message.From.ID, "default_admin_state"); err != nil {
		return err
	}	

	return nil
}

func (b *Bot) handleUserIdInfoMessage(message *tgbotapi.Message) error {
	id, err := b.validateID(message.Text)
	if err != nil {
		return errors.New("invalid id")
	}
	
	info, err := b.service.GetUserInfo(id) 
	if err != nil {
		return err
	}

	b.SendMessage(message.Chat.ID, info)

	if err := b.changeAdminState(message.From.ID, "default_admin_state"); err != nil {
		return err
	}	

	return nil
}

