package log

type noopTracer struct{}

//
var _ Tracer = (*noopTracer)(nil)

//
func NewNoopTracer() Tracer { return &noopTracer{} }

func (noopTracer) Details(string, ...any)  {}
func (noopTracer) Diagnose(string, ...any) {}
func (noopTracer) Fault(string, ...any) {}

func (l *noopTracer) With(...any) Tracer {
	return l
}
