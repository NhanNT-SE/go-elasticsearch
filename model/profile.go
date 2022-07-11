package model

type Profile struct {
	Username      string `bson:"username" json:"username"`
	Email         string `bson:"email" json:"email"`
	Bio           string `bson:"bio" json:"bio"`
	PersonalSite  string `bson:"personal_site" json:"personalSite"`
	WalletAddress string `bson:"wallet_address" json:"walletAddress"`
	AvatarURL     string `bson:"avatar_url" json:"avatarUrl"`
	CoverURL      string `bson:"cover_url" json:"coverUrl"`
}

func (Profile) CollectionName() string {
	return ColProfile
}
