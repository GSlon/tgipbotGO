package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	dbs "github.com/GSlon/tgipbotGO/internal/dbservice"	// custom errors import
	"github.com/GSlon/tgipbotGO/internal/utils"
	_ "strconv"
	"errors"
	"fmt"
)

func (b *Bot) changeUserState(id int, state string) error {
	if err := b.service.SetUserState(id, state); err != nil {
		return err 
	}
	return nil
}

func (b *Bot) getUserState(id int) (string, error) {
	state, err := b.service.GetUserState(id);
	if err != nil {
		return "", err
	}
	return state, nil
}

func (b *Bot) getUser(id int) (string, error) {
	info, err := b.service.GetUserInfo(id)
	if err != nil {
		return "", err
	}
	return info, nil
}

func (b *Bot) getUserHistory(id int) (map[string]string, error) {
	logs, err := b.service.GetUserUniqueIPsExt(id)
	if err != nil {
		return logs, err
	}
	return logs, nil
}

// command handlers
func (b *Bot) handleUserStartCommand(message *tgbotapi.Message) error {
	// create user (if not exists)
	err := b.service.CreateUser(message.From.ID, message.Chat.ID, "default_user_state")
	if err != nil && !errors.Is(err, &dbs.UserAlreadyExistsError{}) {
		return errors.New("internal error")
	}

	b.SendMessage(message.Chat.ID, "user mode")
	return b.handleUserHelpCommand(message)
}

func (b *Bot) handleUserHelpCommand(message *tgbotapi.Message) error {
	b.SendMessage(message.Chat.ID, Userhelp)
	return nil
}

func (b *Bot) handleCheckIPCommand(message *tgbotapi.Message) error {
	if err := b.changeUserState(message.From.ID, "request_ip"); err != nil {
		return err
	}

	b.SendMessage(message.Chat.ID, "input ip")
	return nil
}

func (b *Bot) handleGetHistoryCommand(message *tgbotapi.Message) error {
	info, err := b.getUserHistory(message.From.ID) 
	if err != nil {
		return err
	}

	for ip, log := range info {
		b.SendMessage(message.Chat.ID, fmt.Sprintf("%s : %s", ip, log))
	}
	return nil
}

func (b *Bot) handleUserCancelCommand(message *tgbotapi.Message) error {
	if err := b.changeUserState(message.From.ID, "default_user_state"); err != nil {
		return err
	}

	b.SendMessage(message.Chat.ID, "operation cancelled")
	return nil	
}

func (b *Bot) handleBecomeAdminCommand(message *tgbotapi.Message) error {
	if err := b.createAdmin(message.From.ID); err != nil {
		return err
	}

	b.SendMessage(message.Chat.ID, "now, you are admin")
	return nil 
}

// message handlers
func (b *Bot) handleRequestIPMessage(message *tgbotapi.Message) error {
	// validate and get info about ip
	ip := message.Text

	if !utils.ValidateIPv4(ip) {
		return errors.New("invalid ip")
	}

	info, err := utils.GetIpInfo(ip)
	if err != nil {
		b.service.LogError(err.Error())
		return errors.New("internal error")
	}

	// если успешно, то создаем лог, возвращаем ответ и меняем state 
	// create userLog
	if err := b.service.CreateUserLog(message.From.ID, ip, info); err != nil {
		return err
	}

	b.SendMessage(message.Chat.ID, info)

	if err := b.changeUserState(message.From.ID, "default_user_state"); err != nil {
		return err
	}
	
	return nil
}

