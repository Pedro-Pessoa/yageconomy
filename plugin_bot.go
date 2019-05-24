package yageconomy

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/jonas747/dcmd"
	"github.com/jonas747/discordgo"
	"github.com/jonas747/dstate"
	"github.com/jonas747/yageconomy/models"
	"github.com/jonas747/yagpdb/bot"
	"github.com/jonas747/yagpdb/bot/eventsystem"
	"github.com/jonas747/yagpdb/commands"
	"github.com/jonas747/yagpdb/common"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
	"math/rand"
)

var _ bot.BotInitHandler = (*Plugin)(nil)
var _ commands.CommandProvider = (*Plugin)(nil)

var gofont, _ = truetype.Parse(goregular.TTF)

var CategoryEconomy = &dcmd.Category{
	Name:        "Economy",
	Description: "Ecnonomy commands",
	HelpEmoji:   "$",
	EmbedColor:  0x35afed,
}
var CategoryWaifu = &dcmd.Category{
	Name:        "Waifu",
	Description: "Waifu commands",
	HelpEmoji:   "$",
	EmbedColor:  0xed369e,
}

const (
	ColorBlue = 0x4595e0
	ColorRed  = 0xe04545
)

func (p *Plugin) AddCommands() {

	// commands.AddRootCommands(cmds...)
	commands.AddRootCommandsWithMiddlewares([]dcmd.MiddleWareFunc{economyCmdMiddleware}, CoreCommands...)
	commands.AddRootCommandsWithMiddlewares([]dcmd.MiddleWareFunc{economyCmdMiddleware}, GameCommands...)

	waifuContainer := commands.CommandSystem.Root.Sub("waifu", "wf")
	waifuContainer.NotFound = commands.CommonContainerNotFoundHandler(waifuContainer, "")

	waifuContainer.AddMidlewares(economyCmdMiddleware)

	waifuContainer.AddCommand(WaifuCmdTop, WaifuCmdTop.GetTrigger())
	waifuContainer.AddCommand(WaifuCmdInfo, WaifuCmdInfo.GetTrigger())
	waifuContainer.AddCommand(WaifuCmdClaim, WaifuCmdClaim.GetTrigger())
	waifuContainer.AddCommand(WaifuCmdReset, WaifuCmdReset.GetTrigger())
	waifuContainer.AddCommand(WaifuCmdTransfer, WaifuCmdTransfer.GetTrigger())
	waifuContainer.AddCommand(WaifuCmdDivorce, WaifuCmdDivorce.GetTrigger())
	waifuContainer.AddCommand(WaifuCmdAffinity, WaifuCmdAffinity.GetTrigger())
	waifuContainer.AddCommand(WaifuCmdGift, WaifuCmdGift.GetTrigger())
	waifuContainer.AddCommand(WaifuShopAdd, WaifuShopAdd.GetTrigger().SetMiddlewares(economyAdminMiddleware))
	waifuContainer.AddCommand(WaifuShopEdit, WaifuShopEdit.GetTrigger().SetMiddlewares(economyAdminMiddleware))
	waifuContainer.AddCommand(WaifuCmdDel, WaifuCmdDel.GetTrigger().SetMiddlewares(economyAdminMiddleware))

	commands.AddRootCommandsWithMiddlewares([]dcmd.MiddleWareFunc{economyCmdMiddleware}, ShopCommands...)
	commands.AddRootCommandsWithMiddlewares([]dcmd.MiddleWareFunc{economyCmdMiddleware, economyAdminMiddleware}, ShopAdminCommands...)
}

func (p *Plugin) BotInit() {
	eventsystem.AddHandlerAsyncLast(handleMessageCreate, eventsystem.EventMessageCreate)

}

func economyCmdMiddleware(inner dcmd.RunFunc) dcmd.RunFunc {
	return func(data *dcmd.Data) (interface{}, error) {
		config, err := models.FindEconomyConfigG(data.Context(), data.GS.ID)
		if err != nil {
			if errors.Cause(err) == sql.ErrNoRows {
				config = DefaultConfig(data.GS.ID)
			} else {
				return "Failed retrieving economy config", err
			}
		}

		if !config.Enabled {
			return "Economy is disabled on this server, you can enable it in the control panel.", nil
		}

		ctx := context.WithValue(data.Context(), CtxKeyConfig, config)
		account, _, err := GetCreateAccount(ctx, data.Msg.Author.ID, data.GS.ID, config.StartBalance)
		if err != nil {
			return "Failed creating or retrieving your economy account", err
		}

		ctx = context.WithValue(ctx, CtxKeyUser, account)
		return inner(data.WithContext(ctx))
	}
}

func economyAdminMiddleware(inner dcmd.RunFunc) dcmd.RunFunc {
	return func(data *dcmd.Data) (interface{}, error) {
		ms := commands.ContextMS(data.Context())
		conf := CtxConfig(data.Context())

		if !common.ContainsInt64SliceOneOf(ms.Roles, conf.Admins) {
			return ErrorEmbed(ms, "This command requires you to be an economy admin"), nil
		}

		return inner(data)
	}
}

func GetCreateAccount(ctx context.Context, userID int64, guildID int64, startBalance int64) (account *models.EconomyUser, created bool, err error) {
	return GetCreateAccountExec(common.PQ, ctx, userID, guildID, startBalance)
}

func GetCreateAccountExec(exec boil.ContextExecutor, ctx context.Context, userID int64, guildID int64, startBalance int64) (account *models.EconomyUser, created bool, err error) {
	account, err = models.FindEconomyUser(ctx, exec, guildID, userID)
	if err == nil {
		return account, false, nil
	}

	if errors.Cause(err) != sql.ErrNoRows {
		return nil, false, err
	}

	account = &models.EconomyUser{
		GuildID:   guildID,
		UserID:    userID,
		MoneyBank: startBalance,
	}

	err = account.Insert(ctx, exec, boil.Infer())
	if err != nil {
		return nil, false, err
	}

	return account, true, nil
}

type CtxKey int

const (
	CtxKeyConfig CtxKey = iota
	CtxKeyUser
)

func CtxConfig(c context.Context) *models.EconomyConfig {
	return c.Value(CtxKeyConfig).(*models.EconomyConfig)
}

func CtxUser(c context.Context) *models.EconomyUser {
	return c.Value(CtxKeyUser).(*models.EconomyUser)
}

func UserEmebdAuthor(ms *dstate.MemberState) *discordgo.MessageEmbedAuthor {

	user := ms.DGoUser()

	return &discordgo.MessageEmbedAuthor{
		Name:    user.Username + "#" + user.Discriminator,
		IconURL: user.AvatarURL("128"),
	}
}

func SimpleEmbedResponse(ms *dstate.MemberState, msgF string, args ...interface{}) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Author:      UserEmebdAuthor(ms),
		Color:       ColorBlue,
		Description: fmt.Sprintf(msgF, args...),
	}
}

func ErrorEmbed(ms *dstate.MemberState, msgF string, args ...interface{}) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Author:      UserEmebdAuthor(ms),
		Color:       ColorRed,
		Description: fmt.Sprintf(msgF, args...),
	}
}

func handleMessageCreate(evt *eventsystem.EventData) {
	msg := evt.MessageCreate()
	if msg.Author == nil || msg.Author.Bot {
		return
	}

	conf, err := models.FindEconomyConfigG(evt.Context(), msg.GuildID)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			conf = DefaultConfig(msg.GuildID)
		} else {
			logger.WithError(err).WithField("guild", msg.GuildID).Error("failed retrieving economy config")
			return
		}
	}

	if !conf.Enabled {
		return
	}

	if conf.ChatmoneyAmountMax > 0 {
		// gen chat money maybe?
		amount := rand.Int63n(conf.ChatmoneyAmountMax-conf.ChatmoneyAmountMin) + conf.ChatmoneyAmountMin

		result, err := common.PQ.Exec(`UPDATE economy_users SET last_chatmoney_claim = now(), money_wallet = money_wallet + $4
			WHERE guild_id = $1 AND user_id = $2 AND EXTRACT(EPOCH FROM (now() - last_chatmoney_claim))  > $3`,
			msg.GuildID, msg.Author.ID, conf.ChatmoneyFrequency, amount)
		if err != nil {
			logger.WithField("guild", msg.GuildID).WithError(err).Error("failed claiming chatmoney")
			return
		}

		rows, err := result.RowsAffected()
		if err != nil {
			logger.WithField("guild", msg.GuildID).WithError(err).Error("failed claiming chatmoney, rows")
			return
		}

		if rows > 0 {
			logger.Infof("Gave %s (%d) chat money of %d", msg.Author.Username, msg.Author.ID, amount)
		}
	}

	// maybe plant?
	if common.ContainsInt64Slice(conf.AutoPlantChannels, msg.ChannelID) {
		if rand.Float64() < 0.98 {
			return
		}

		amount := rand.Int63n(conf.AutoPlantMax-conf.AutoPlantMin) + conf.AutoPlantMin

		availableChars := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

		pw := ""
		for i := 0; i < 6; i++ {
			char := availableChars[rand.Intn(len(availableChars))]
			pw += char
		}

		// Plant!
		err = PlantMoney(context.Background(), conf, msg.ChannelID, 0, int(amount), pw)
		if err != nil {
			logger.WithError(err).WithField("guild", msg.GuildID).WithField("channel", msg.ChannelID).Error("failed planting money")
		}
	}
}

var errAlreadyPlantInChannel = errors.New("Already money planted in this channel")

func PlantMoney(ctx context.Context, conf *models.EconomyConfig, channelID, author int64, amount int, password string) error {
	_, err := models.FindEconomyPlantG(ctx, channelID)
	if err == nil {
		return errAlreadyPlantInChannel
	}

	if errors.Cause(err) != sql.ErrNoRows {
		return err
	}

	r, err := models.FindEconomyPickImageG(ctx, conf.GuildID)
	if err != nil && errors.Cause(err) != sql.ErrNoRows {
		return err
	}

	m := &models.EconomyPlant{
		ChannelID: channelID,
		GuildID:   conf.GuildID,
		AuthorID:  author,
		Amount:    int64(amount),
		Password:  password,
	}

	err = m.InsertG(ctx, boil.Infer())
	if err != nil {
		return err
	}

	img, _, err := image.Decode(bytes.NewReader(r.Image))
	if err != nil {
		return err
	}

	var msg *discordgo.Message
	if r != nil {
		c := gg.NewContextForImage(img)
		c.SetFontFace(truetype.NewFace(gofont, &truetype.Options{
			Size: 36,
		}))

		c.SetColor(color.RGBA{
			1, 1, 1, 156,
		})
		width, height := c.MeasureString(password)
		c.DrawRectangle(0, 0, width+20, height+20)
		c.Fill()

		c.SetColor(color.RGBA{
			255, 255, 255, 255,
		})

		c.DrawString(password, 5, height+5)

		buf := bytes.NewBuffer(nil)
		err = c.EncodePNG(buf)
		if err != nil {
			return err
		}

		msgContent := fmt.Sprintf("%d random %s appeared! Pick them up with `pick <password>`", amount, conf.CurrencySymbol)
		msg, err = common.BotSession.ChannelFileSendWithMessage(channelID, msgContent, "plant.png", buf)
	} else {
		// fallback if no image is set
		embed := &discordgo.MessageEmbed{
			Description: fmt.Sprintf("%d random %s appeared! Pick them up with `pick %s`", amount, conf.CurrencySymbol, password),
			Color:       ColorBlue,
		}
		msg, err = common.BotSession.ChannelMessageSendEmbed(channelID, embed)
	}

	if err != nil {
		return err
	}

	m.MessageID = msg.ID
	_, err = m.UpdateG(ctx, boil.Whitelist("message_id"))

	return err
}

var (
	ErrInsufficientFunds = errors.New("Insufficient funds")
)

// TransferMoneyWallet transfers money from one users wallet to another
// both from and to is optional
// out and in amount can be different in certain cases (such as gambling)
func TransferMoneyWallet(ctx context.Context, tx *sql.Tx, conf *models.EconomyConfig, checkFunds bool, from, to int64, outAmount, inAmount int64) (err error) {
	createdTX := false
	if tx == nil {
		createdTX = true
		tx, err = common.PQ.Begin()
		if err != nil {
			return err
		}
	}

	// make sure the origin account is created
	if from != 0 {
		account, _, err := GetCreateAccountExec(tx, ctx, from, conf.GuildID, conf.StartBalance)
		if err != nil {
			if createdTX {
				tx.Rollback()
			}
			return err
		}
		if checkFunds && account.MoneyWallet < outAmount {
			if createdTX {
				tx.Rollback()
			}
			return ErrInsufficientFunds
		}
	}

	// make sure the destination account is created
	if to != 0 {
		_, _, err = GetCreateAccountExec(tx, ctx, to, conf.GuildID, conf.StartBalance)
		if err != nil {
			if createdTX {
				tx.Rollback()
			}
			return err
		}
	}

	// update origin account
	if from != 0 {
		_, err := tx.Exec("UPDATE economy_users SET money_wallet = money_wallet - $3 WHERE user_id = $2 AND guild_id = $1", conf.GuildID, from, outAmount)
		if err != nil {
			if createdTX {
				tx.Rollback()
			}
			return err
		}
	}

	// update the destination account
	if to != 0 {
		_, err = tx.Exec("UPDATE economy_users SET money_wallet = money_wallet + $3 WHERE user_id = $2 AND guild_id = $1", conf.GuildID, to, inAmount)
		if err != nil {
			if createdTX {
				tx.Rollback()
			}
			return err
		}
	}

	if createdTX {
		return tx.Commit()
	}

	return err
}