package interfaces

import "ticktack/internal/status"

type HasStatus interface {
	SetStatus(status.Status)
	GetStatus() status.Status
}
