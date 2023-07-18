package test

type Api interface {
	get(name string) string
}

func Get(name string, api Api) string {
	return name + api.get(name) + "aaa"
}
