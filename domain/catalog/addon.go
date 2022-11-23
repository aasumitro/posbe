package catalog

type (
	Addon struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	IAddonRepository interface {
		SelectAll()
		SelectBy()
		Insert()
		Update()
		Delete()
	}

	IAddonService interface {
		Fetch()
		Find()
		Create()
		Update()
		Delete()
	}

	IAddonHandler interface {
		List()
		Show()
		Store()
		Update()
		Destroy()
	}
)
