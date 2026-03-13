package botEngine

func (r *RequestContext) Set(key string, item interface{}) {
	r.customContext[key] = item
}

func (r *RequestContext) Get(key string) (interface{}, bool) {
	item, ok := r.customContext[key]
	return item, ok
}
