package ecs

type Entity struct {
	Components []Component
}

func TryGetComponent[T any](e *Entity) *T {
	for _, component := range e.Components {
		switch component.(type) {
		case T:
			c := component.(T)
			return &c
		}
	}
	return nil
}
