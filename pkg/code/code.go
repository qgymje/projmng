package code

type Code struct {
	ErrorCode
	detail error
}

var Success = New(OK, nil)

func New(c ErrorCode, err error) *Code {
	return &Code{
		ErrorCode: c,
		detail:    err,
	}
}

func (c *Code) String() string {
	if c.detail != nil {
		return c.detail.Error()
	}
	return ""
}

func (c *Code) Error() string {
	return c.ErrorCode.String()
}

//go:generate stringer -type=ErrorCode
type ErrorCode int

const (
	OK                 ErrorCode = 0
	CreateProjectError ErrorCode = 1000
	UpdateProjectError ErrorCode = 1001
	DeleteProjectError ErrorCode = 1002
)
