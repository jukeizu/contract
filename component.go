package contract

type Components struct {
	ActionsRows []*ActionsRow `json:"actionsRows,omitempty" yaml:"actionsRows,omitempty"`
	Buttons     []*Button     `json:"buttons,omitempty" yaml:"buttons,omitempty"`
	SelectMenus []*SelectMenu `json:"selectMenus,omitempty" yaml:"selectMenus,omitempty"`
	TextInputs  []*TextInput  `json:"textInputs,omitempty" yaml:"textInputs,omitempty"`
}

type ActionsRow struct {
	Buttons     []*Button     `json:"buttons,omitempty" yaml:"buttons,omitempty"`
	SelectMenus []*SelectMenu `json:"selectMenus,omitempty" yaml:"selectMenus,omitempty"`
	TextInputs  []*TextInput  `json:"textInputs,omitempty" yaml:"textInputs,omitempty"`
}

type ComponentEmoji struct {
	Id       string `json:"id,omitempty" yaml:"id,omitempty"`
	Name     string `json:"name,omitempty" yaml:"name,omitempty"`
	Animated bool   `json:"animated,omitempty" yaml:"animated,omitempty"`
}

type Button struct {
	Label    string         `json:"label" yaml:"label,omitempty"`
	Style    uint           `json:"style" yaml:"style,omitempty"`
	Disabled bool           `json:"disabled" yaml:"disabled,omitempty"`
	Emoji    ComponentEmoji `json:"emoji" yaml:"emoji,omitempty"`
	Url      string         `json:"url,omitempty" yaml:"url,omitempty"`
	CustomId string         `json:"customId,omitempty" yaml:"customId,omitempty"`
}

type SelectMenuOption struct {
	Label       string         `json:"label,omitempty" yaml:"label,omitempty"`
	Value       string         `json:"value" yaml:"value,omitempty"`
	Description string         `json:"description" yaml:"description,omitempty"`
	Emoji       ComponentEmoji `json:"emoji" yaml:"emoji,omitempty"`
	Default     bool           `json:"default" yaml:"default,omitempty"`
}

type SelectMenu struct {
	CustomId    string             `json:"customId,omitempty" yaml:"customId,omitempty"`
	Placeholder string             `json:"placeholder" yaml:"placeholder,omitempty"`
	MinValues   *int               `json:"minValues,omitempty" yaml:"minValues,omitempty"`
	MaxValues   int                `json:"maxValues,omitempty" yaml:"maxValues,omitempty"`
	Options     []SelectMenuOption `json:"options" yaml:"options,omitempty"`
	Disabled    bool               `json:"disabled" yaml:"disabled,omitempty"`
}

type TextInput struct {
	CustomId    string `json:"customId" yaml:"customId,omitempty"`
	Label       string `json:"label" yaml:"label,omitempty"`
	Style       uint   `json:"style" yaml:"style,omitempty"`
	Placeholder string `json:"placeholder,omitempty" yaml:"placeholder,omitempty"`
	Value       string `json:"value,omitempty" yaml:"value,omitempty"`
	Required    bool   `json:"required,omitempty" yaml:"required,omitempty"`
	MinLength   int    `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	MaxLength   int    `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
}
