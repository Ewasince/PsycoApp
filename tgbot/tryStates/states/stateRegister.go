package states

import (
	msg "PsychoBot/messages"
	. "PsychoBot/teleBotStateLib"
	"PsychoBot/tryStates/context"
	. "StorageModule/models"
	"StorageModule/repo"
	"errors"
	"gorm.io/gorm"
)

var RegisterState = newBotStateWrapper(
	"Register state",
	&BotMessages{msg.Greeting, msg.Register},
	BotMessageHandler(exitMessageHandlerRegisterState),
	nil,
	messageHandlerRegisterState,
)

func exitMessageHandlerRegisterState(c BotContext) ([]string, error) {
	ctx := *c.(*context.MyBotContext)
	if ctx.IsPatientRegistered() {
		_ = ctx.CreateAndSendMessage(msg.CantCreatePatient)
		return []string{}, errors.New("patient was complete register, but wasn't registered ")
	}
	return []string{msg.RegisterComplete}, nil
}

func messageHandlerRegisterState(c BotContext) (HandlerResponse, error) {
	ctx := *c.(*context.MyBotContext)

	var user *User
	user, err := repo.GetUserByUsername(ctx.MessageText)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		_ = ctx.CreateAndSendMessage(msg.UserNotFound)
		return HandlerResponse{}, nil
	}
	if err != nil {
		return HandlerResponse{}, err
	}

	patient := &Patient{
		BaseModel: BaseModel{
			Model: gorm.Model{},
		},
		Name:     ctx.Message.From.FirstName,
		LastName: ctx.Message.From.LastName,
		Email:    "",
		Username: ctx.Message.From.UserName,
		Password: "",
		UserId:   user.ID,
		TgId:     ctx.PatientTgId,
		TgChatId: &ctx.Message.Chat.ID,
	}
	err = repo.CreatePatient(patient)
	if err != nil {
		_ = ctx.CreateAndSendMessage(msg.CantCreatePatient)
		return HandlerResponse{}, nil
	}
	ctx.Patient = patient
	return HandlerResponse{
		NextState:      &DefaultState,
		TransitionType: GoState,
	}, nil
}
