package lib

type Cookies struct {
	kv map[string]string
}

func (cookies *Cookies) Add(key string, value string) *Cookies {
	if cookies.kv == nil {
		cookies.kv = make(map[string]string)
	}

	cookies.kv[key] = value

	return cookies
}

func (cookies *Cookies) AddBulk(kv map[string]string) *Cookies {
	if cookies.kv == nil {
		cookies.kv = make(map[string]string)
	}

	for k, v := range kv {
		cookies.kv[k] = v
	}

	return cookies
}
