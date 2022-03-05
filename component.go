package contract

type Components struct {
	ActionsRows []*ActionsRow `json:"actionsRows"`
	Buttons     []*Button     `json:"buttons"`
	SelectMenus []*SelectMenu `json:"selectMenus"`
	TextInputs  []*TextInput  `json:"textInputs"`
}

type ActionsRow struct {
	Buttons     []*Button     `json:"buttons"`
	SelectMenus []*SelectMenu `json:"selectMenus"`
	TextInputs  []*TextInput  `json:"textInputs"`
}

type ComponentEmoji struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Animated bool   `json:"animated,omitempty"`
}

type Button struct {
	Label    string         `json:"label"`
	Style    uint           `json:"style"`
	Disabled bool           `json:"disabled"`
	Emoji    ComponentEmoji `json:"emoji"`
	Url      string         `json:"url,omitempty"`
	CustomId string         `json:"customId,omitempty"`
}

type SelectMenuOption struct {
	Label       string         `json:"label,omitempty"`
	Value       string         `json:"value"`
	Description string         `json:"description"`
	Emoji       ComponentEmoji `json:"emoji"`
	Default     bool           `json:"default"`
}

type SelectMenu struct {
	CustomId    string             `json:"customId,omitempty"`
	Placeholder string             `json:"placeholder"`
	MinValues   *int               `json:"minValues,omitempty"`
	MaxValues   int                `json:"maxValues,omitempty"`
	Options     []SelectMenuOption `json:"options"`
	Disabled    bool               `json:"disabled"`
}

type TextInput struct {
	CustomId    string `json:"customId"`
	Label       string `json:"label"`
	Style       uint   `json:"style"`
	Placeholder string `json:"placeholder,omitempty"`
	Value       string `json:"value,omitempty"`
	Required    bool   `json:"required,omitempty"`
	MinLength   int    `json:"minLength,omitempty"`
	MaxLength   int    `json:"maxLength,omitempty"`
}
