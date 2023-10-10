package manifest

type Manifest struct {
	Name            string `json:"name"`
	ShortName       string `json:"short_name"`
	Icons           []Icon `json:"icons"`
	ThemeColor      string `json:"theme_color"`
	BackgroundColor string `json:"background_color"`
	Display         string `json:"display"`
}

type Icon struct {
	Source string `json:"src"`
	Sizes  string `json:"sizes"`
	Type   string `json:"type"`
}

func NewManifest(env string) Manifest {
	switch env {
	case "production":
		return Manifest{
			Name:      "HomeOS Shrtnr",
			ShortName: "Shrtnr",
			Icons: []Icon{
				{
					"/dist/vite.svg",
					"48x48 72x72 96x96 144x144 192x192",
					"image/svg",
				},
			},
			ThemeColor:      "#367fa9",
			BackgroundColor: "#367fa9",
			Display:         "standalone",
		}
	default:
		return Manifest{
			Name:      "HomeOS Shrtnr",
			ShortName: "Shrtnr",
			Icons: []Icon{
				{
					"http://localhost:5173/vite.svg",
					"48x48 72x72 96x96 144x144 192x192",
					"image/svg",
				},
			},
			ThemeColor:      "#367fa9",
			BackgroundColor: "#367fa9",
			Display:         "standalone",
		}
	}

}
