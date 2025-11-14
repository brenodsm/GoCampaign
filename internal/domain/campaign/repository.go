package campaign

// Repository abstracts the persistence of Campaign entities.
type Repository interface {
	Save(campaign *Campaign) error
	GetAll() ([]Campaign, error)
	GetByID(id string) (*Campaign, error)
	UpdateStatus(id, status string) error
	Delete(id string) error
}
