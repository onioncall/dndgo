package create

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/onioncall/dndgo/character-management/handlers"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/logger"
)

const (
	orange    = lipgloss.Color("#FFA500")
	lightBlue = lipgloss.Color("#5DC9E2")
	cream     = lipgloss.Color("#F9F6F0")
	darkGray  = lipgloss.Color("#767676")
)

var (
	primaryStyle   = lipgloss.NewStyle().Foreground(lightBlue)
	secondaryStyle = lipgloss.NewStyle().Foreground(orange)
	tertiaryStyle  = lipgloss.NewStyle().Foreground(cream)
)

const (
	basicInfoPage = iota
	abilitiesPage
	skillsPage
	spellsPage
	weaponsPage
	wornEquipmentPage
	backpackPage
	dumpPage
)

type Model struct {
	inputs            []textinput.Model
	focused           int
	err               error
	width             int
	height            int
	existingNames     []string
	addButtonFocused  bool
	nextButtonFocused bool
	backButtonFocused bool
	viewportOffset    int
	currentPage       int
	character         *models.Character
}

func NewModel() Model {
	inputs := basicInfoInputs()
	names, err := handlers.GetCharacterNames()
	if err != nil {
		logger.Errorf("Create Character TUI Page failed to get character names:\n%v", err)
	}

	return Model{
		inputs:            inputs,
		focused:           0,
		err:               nil,
		nextButtonFocused: false,
		existingNames:     names,
		viewportOffset:    0,
		character:         &models.Character{},
	}
}

func (m *Model) View() string {
	switch m.currentPage {
	case basicInfoPage:
		return m.BasicInfoPageView()
	case abilitiesPage:
		return m.AbilitiesPageView()
	case skillsPage:
		return m.SkillsPageView()
	case spellsPage:
		return m.SpellsPageView()
	case weaponsPage:
		return m.WeaponsPageView()
	case wornEquipmentPage:
		return m.EquipmentPageView()
	case backpackPage:
		return m.BackpackPageView()
	case dumpPage:
		return m.DumpView()
	default:
		return "Unknown page"
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch m.currentPage {
	case basicInfoPage:
		return m.UpdateBasicInfoPage(msg)
	case abilitiesPage:
		return m.UpdateAbilitiesPage(msg)
	case skillsPage:
		return m.UpdateSkillsPage(msg)
	case spellsPage:
		return m.UpdateSpellsPage(msg)
	case weaponsPage:
		return m.UpdateWeaponsPage(msg)
	case wornEquipmentPage:
		return m.UpdateEquipmentPage(msg)
	case backpackPage:
		return m.UpdateBackpackPage(msg)
	case dumpPage:
		return m.DumpUpdate(msg)
	default:
		return m, nil
	}
}

func (m *Model) nextInput(numButtons int) {
	m.focused = (m.focused + 1) % (len(m.inputs) + numButtons)
}

func (m *Model) prevInput(numButtons int) {
	m.focused--
	if m.focused < 0 {
		m.focused = len(m.inputs) + numButtons - 1
	}
}

func (m *Model) updateViewportGeneric(linesPerField, headerLines, footerLines int) {
	if m.height == 0 {
		return
	}

	availableLines := m.height - headerLines - footerLines
	visibleFields := max(availableLines/linesPerField, 1)

	maxOffset := max(0, len(m.inputs)-visibleFields)

	if m.nextButtonFocused {
		m.viewportOffset = maxOffset
		return
	}

	if m.focused >= m.viewportOffset+visibleFields {
		m.viewportOffset = m.focused - visibleFields + 1
	}

	if m.focused < m.viewportOffset {
		m.viewportOffset = m.focused
	}

	m.viewportOffset = min(m.viewportOffset, maxOffset)
}

func (m Model) renderError() string {
	if m.err == nil {
		return ""
	}

	errorStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FF0000")).
		Padding(0, 1).
		Foreground(lipgloss.Color("#FF0000"))

	return "\n\n" + errorStyle.Render(m.err.Error())
}

func getScrollIndicators(startIndex int, endIndex int, totalItems int, visibleItems int) string {
	scrollIndicators := ""

	// If everything fits on the screen, we don't need to worry about scroll indicatiors
	if visibleItems >= totalItems {
		return scrollIndicators
	}

	if startIndex > 0 {
		scrollIndicators += "▲ "
	}
	if endIndex < totalItems {
		scrollIndicators += "▼"
	}
	scrollIndicators = strings.TrimSpace(scrollIndicators) + "\n"
	return scrollIndicators
}
