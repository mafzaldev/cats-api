package api

type CatParams struct {
	Name string
}

type CatResponse struct {
	Code  int
	Name  string
	Image string
	Tags  []string
}

type Error struct {
	Code    int
	Message string
}
