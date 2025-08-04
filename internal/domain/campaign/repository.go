package campaign

// Repository abstracts the persistence of Campaign entities.
type Repository interface {
	Save(campaign *Campaign) error
}
