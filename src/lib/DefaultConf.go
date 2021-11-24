package lib

func DefaultConf() {
	conf = Conf{
		Browser:  "chrome",
		Env:      "dev",
		Headless: true,
		Port:     4444,
		Width:    1920,
		Height:   1028,
	}

	SetCaps(conf)
}
