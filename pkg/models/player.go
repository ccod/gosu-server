package models

// Player is going to be basically my user struct
type Player struct {
	AccountID      int          `gorm:"primary_key" json:"accountId"`
	ProfileID      int          `json:"profileId"`
	ProfileURL     string       `json:"profileUrl"`
	AvatarURL      string       `json:"avatarUrl"`
	ClanTag        string       `json:"clanTag"`
	ClanName       string       `json:"clanName"`
	DisplayName    string       `json:"displayName"`
	RealmID        int          `json:"realmId"`
	RegionID       int          `json:"regionId"`
	Admin          bool         `gorm:"default:false" json:"admin"`
	LadderRecord   LadderRecord `json:"ladderRecord"`
	LadderRecordID int          `json:"-"`
}
