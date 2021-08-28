package handlers


func ExecFnIgnoringError(fn func() error) {
	_ = fn()
}