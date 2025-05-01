package permission

const (
	Read    int64 = 1 << 0
	Create  int64 = 1 << 1
	Update  int64 = 1 << 2
	Delete  int64 = 1 << 3
	Active  int64 = 1 << 4
	Publish int64 = 1 << 5
	Archive int64 = 1 << 6
)

func HasPermission(permissionValue int64, permission int64) bool {
	return permissionValue&permission != 0
}

func AddPermission(permissionValue *int64, permission int64) {
	*permissionValue |= permission
}
