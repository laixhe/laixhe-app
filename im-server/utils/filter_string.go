package utils

type filterString struct {
	filter map[string]string
}

func NewFilterString() *filterString{
	return &filterString{
		filter: make(map[string]string),
	}
}

func (f *filterString)Add(s string){
	f.filter[s] = s
}

func (f *filterString)Get() []string {
	if len(f.filter) == 0 {
		return nil
	}
	ss := make([]string, 0, len(f.filter))
	for v := range f.filter {
		ss = append(ss, v)
	}
	return ss
}
