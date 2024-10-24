package typed_strings

type TaskSettingsStatus string

const (
	TaskSettingsStatusActive   TaskSettingsStatus = "ACTIVE"
	TaskSettingsStatusInActive TaskSettingsStatus = "INACTIVE"
)

func (tss TaskSettingsStatus) String() string {
	return string(tss)
}
