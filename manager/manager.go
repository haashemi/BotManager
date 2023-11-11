package manager

import (
	"sort"
	"sync"

	"github.com/haashemi/BotManagerBot/config"
	"github.com/samber/lo"
)

const BotsDataPath = "bots.json"

// Manager contains all whitelisted bots
type Manager struct {
	bots   map[string]*Bot
	botMut sync.RWMutex
}

// NewManager returns a new initialized Manager
//
// it loads the stored bots from previous runs, adds the manager's bot
// if not exists, and returns the initialized manager instance
func NewManager(config *config.Config) (*Manager, error) {
	// load all stored bots at BotsDataPath from previous runs
	bots, err := loadBots(BotsDataPath)
	if err != nil {
		return nil, err
	}

	manager := &Manager{
		// It just converts the bots slice to a map, don't even try to read it.
		bots: lo.Reduce(bots, func(agg map[string]*Bot, item *Bot, index int) map[string]*Bot { agg[item.Token] = item; return agg }, map[string]*Bot{}),
	}

	// add bot manager bot's token to the whitelist if not exists for some reason
	if _, ok := manager.bots[config.Bot.Token]; !ok {
		bot, err := NewBot(config.Bot.Token, config.TelegramBotAPI.Host)
		if err != nil {
			return nil, err
		}

		if err = manager.AddBot(bot); err != nil {
			return nil, err
		}
	}

	return manager, nil
}

// Bots returns the current bots in the manager ordered by their id.
func (m *Manager) Bots() []*Bot {
	// as maps are not ordered, we'll make it a slice.
	// we also take care of mutex locking.
	m.botMut.RLock()
	bots := lo.Values(m.bots)
	m.botMut.RUnlock()

	// Sort them so end-user (via the manager-bot) could see them in the same order every time.
	sort.Slice(bots, func(i, j int) bool { return bots[i].ID < bots[j].ID })

	return bots
}

// AddBot adds/replaces a new bot to the manager's bots.
func (m *Manager) AddBot(bot *Bot) error {
	m.botMut.Lock()
	m.bots[bot.Token] = bot
	m.botMut.Unlock()

	return save(m.Bots(), BotsDataPath)
}

// RemoveBot removes the bot from the manager's bots.
func (m *Manager) RemoveBot(token string) (*Bot, error) {
	m.botMut.Lock()
	bot, ok := m.bots[token]
	if ok {
		delete(m.bots, token)
	}
	m.botMut.Unlock()

	return bot, save(m.Bots(), BotsDataPath)
}

// IsBotExists checks the manager's bots to check the existence of the token.
func (m *Manager) IsBotExists(token string) bool {
	m.botMut.RLock()
	defer m.botMut.RUnlock()

	_, exists := m.bots[token]
	return exists
}
