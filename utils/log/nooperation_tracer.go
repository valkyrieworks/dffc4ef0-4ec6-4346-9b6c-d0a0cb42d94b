package log

type nooperationTracer struct{}

//
var _ Tracer = (*nooperationTracer)(nil)

//
func FreshNooperationTracer() Tracer { return &nooperationTracer{} }

func (nooperationTracer) Details(string, ...any)  {}
func (nooperationTracer) Diagnose(string, ...any) {}
func (nooperationTracer) Failure(string, ...any) {}

func (l *nooperationTracer) Using(...any) Tracer {
	return l
}
