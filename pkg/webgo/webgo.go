package webgo

func (w *WebGo) Start() {
	w.Server.Start(w.Logger)
}
