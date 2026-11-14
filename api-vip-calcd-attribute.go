package jx3api

import (
	"context"
	"encoding/json"
)

type CalcdAttributeResponse struct {
	BodyId         int               `json:"bodyId"`
	BodyName       string            `json:"bodyName"`
	CacheTime      int64             `json:"cacheTime"`
	CampId         int               `json:"campId"`
	CampName       string            `json:"campName"`
	EnchantEffects []json.RawMessage `json:"enchantEffects,omitempty"`
	EquipList      []struct {
		AColorDiamond []int   `json:"aColorDiamond,omitempty"`
		AMountDiamond [][]int `json:"aMountDiamond,omitempty"`
		BelongSchool  *string `json:"belongSchool,omitempty"`
		ColorDiamond  *struct {
			IconId   string `json:"iconId"`
			ItemAttr struct {
				AtPhysicsAttackPowerBase int `json:"atPhysicsAttackPowerBase"`
				AtPhysicsCriticalStrike  int `json:"atPhysicsCriticalStrike"`
				AtPhysicsOvercomeBase    int `json:"atPhysicsOvercomeBase"`
			} `json:"itemAttr"`
			ItemId      int    `json:"itemId"`
			ItemName    string `json:"itemName"`
			ItemQuality string `json:"itemQuality"`
			ItemScore   int    `json:"itemScore"`
		} `json:"colorDiamond,omitempty"`
		DwPermanentEnchantID *int `json:"dwPermanentEnchantID,omitempty"`
		DwTabIndex           *int `json:"dwTabIndex,omitempty"`
		DwTabType            *int `json:"dwTabType,omitempty"`
		DwTemporaryEnchantID *int `json:"dwTemporaryEnchantID,omitempty"`
		IconId               *int `json:"iconId,omitempty"`
		ItemAttrib           []struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"itemAttrib,omitempty"`
		ItemDesc  []string `json:"itemDesc,omitempty"`
		ItemIndex *int     `json:"itemIndex,omitempty"`
		ItemKind  *string  `json:"itemKind,omitempty"`
		ItemLevel *int     `json:"itemLevel,omitempty"`
		ItemMagic []struct {
			ItemDesc  json.RawMessage `json:"itemDesc"`
			ItemName  string          `json:"itemName"`
			ItemValue struct {
				Param1Max int `json:"Param1Max"`
				Param1Min int `json:"Param1Min"`
			} `json:"itemValue"`
			TypeName string `json:"typeName"`
		} `json:"itemMagic,omitempty"`
		ItemName         *string                    `json:"itemName,omitempty"`
		ItemPlus         json.RawMessage            `json:"itemPlus,omitempty"`
		ItemQuality      *int                       `json:"itemQuality,omitempty"`
		ItemSchool       *string                    `json:"itemSchool,omitempty"`
		ItemSource       *string                    `json:"itemSource,omitempty"`
		ItemType         *string                    `json:"itemType,omitempty"`
		MagicType        *string                    `json:"magicType,omitempty"`
		MaxEnchant       map[string]json.RawMessage `json:"maxEnchant,omitempty"`
		MaxStrengthLevel *int                       `json:"maxStrengthLevel,omitempty"`
		MinEnchant       *struct {
			Attribute1ID     *string `json:"attribute1ID,omitempty"`
			Attribute1Name   *string `json:"attribute1Name,omitempty"`
			Attribute1Value1 *string `json:"attribute1Value1,omitempty"`
			Attribute1Value2 *string `json:"attribute1Value2,omitempty"`
			DestItemSubType  *string `json:"destItemSubType,omitempty"`
			ItemId           *string `json:"itemId,omitempty"`
			ItemName         *string `json:"itemName,omitempty"`
			ItemScore        *string `json:"itemScore,omitempty"`
			ItemType         *string `json:"itemType,omitempty"`
		} `json:"minEnchant,omitempty"`
		MountDiamond []struct {
			BindType    int    `json:"bindType"`
			IconId      int    `json:"iconId"`
			ItemDesc    string `json:"itemDesc"`
			ItemId      int    `json:"itemId"`
			ItemName    string `json:"itemName"`
			ItemQuality int    `json:"itemQuality"`
		} `json:"mountDiamond,omitempty"`
		NItemIndex     *int    `json:"nItemIndex,omitempty"`
		NStrengthLevel *int    `json:"nStrengthLevel,omitempty"`
		NeedLevel      *int    `json:"needLevel,omitempty"`
		SchoolName     *string `json:"schoolName,omitempty"`
		SetId          *string `json:"setId,omitempty"`
		StrengthLevel  *int    `json:"strengthLevel,omitempty"`
		TabId          *int    `json:"tabId,omitempty"`
		TabIndex       *int    `json:"tabIndex,omitempty"`
	} `json:"equipList"`
	EquipScore *int   `json:"equipScore,omitempty"`
	EquipType  string `json:"equipType"`
	ForceId    int    `json:"forceId"`
	ForceName  string `json:"forceName"`
	GlobalId   string `json:"globalId"`
	KungfuId   string `json:"kungfuId"`
	KungfuName string `json:"kungfuName"`
	KungfuType string `json:"kungfuType"`
	PanelList  []struct {
		Name    string   `json:"name"`
		Percent *float64 `json:"percent,omitempty"`
		Value   int      `json:"value"`
	} `json:"panelList,omitempty"`
	RoleId     string `json:"roleId"`
	RoleName   string `json:"roleName"`
	SeasonName string `json:"seasonName"`
	ServerName string `json:"serverName"`
	ShowInfo   struct {
		ShowAvatar string `json:"showAvatar"`
		ShowHash   string `json:"showHash"`
		ShowIndex  int    `json:"showIndex"`
	} `json:"showInfo"`
	Suits []struct {
		Attributes []struct {
			Name   string `json:"name"`
			Pieces int    `json:"pieces"`
			Value  int    `json:"value"`
		} `json:"attributes"`
		Count   int    `json:"count"`
		Name    string `json:"name"`
		Recipes int    `json:"recipes"`
	} `json:"suits,omitempty"`
	TongName string `json:"tongName"`
	ZoneName string `json:"zoneName"`
}

// CalcdAttribute 查询角色已收录的装备与属性面板。
func (c *Client) CalcdAttribute(ctx context.Context, server string, name string, format string, token string) (*CalcdAttributeResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Format string `json:"format,omitempty"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Format: format,
		Token:  token,
	}
	return requestData[CalcdAttributeResponse](ctx, c, "/calcd/attribute", params)
}
