package domain

type FileHome struct {
	Image       string `form:"image" validate:"required"`
	Foto        string `form:"foto" validate:"required"`
	PembicaraID uint   `json:"pid"`
}

type FileSeminar struct {
	Rundown string `form:"rundown" validate:"required"`
	Materi  string `form:"materi" validate:"required"`
	CV      string `form:"cv" validate:"required"`
}

type FileRakerda struct {
	Rundown     string `form:"rundown" validate:"required"`
	Materi      string `form:"materi" validate:"required"`
	MPengabdian string `form:"m_abdi" validate:"require"`
}

type PID struct {
	Rundown  string    `form:"rundown" validate:"required"`
	Poster   string    `form:"poster" validate:"required"`
	Artikels []Artikel `json:"artikel"`
}
