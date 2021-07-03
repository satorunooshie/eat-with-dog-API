package errors

import (
	"net/http"
	"strconv"
	"testing"
)

func Test_appError_IsCritical(t *testing.T) {
	tests := []struct {
		name string
		err  AppError
		want bool
	}{
		{
			name: strconv.Itoa(http.StatusNotFound),
			err:  NotFound.New(),
			want: false,
		},
		{
			name: strconv.Itoa(http.StatusInternalServerError),
			err:  SystemDefault.New(),
			want: true,
		},
		{
			name: strconv.Itoa(http.StatusBadRequest) + " critical",
			err:  InvalidParameter.New().Critical(),
			want: true,
		},
		{
			name: strconv.Itoa(http.StatusBadRequest) + " wrap critical",
			err:  InvalidParameter.Wrap(InvalidParameter.New().Critical()),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.IsCritical(); got != tt.want {
				t.Errorf("IsCritical() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appError_IsInfo(t *testing.T) {
	tests := []struct {
		name string
		err  AppError
		want bool
	}{
		{
			name: strconv.Itoa(http.StatusNotFound),
			err:  NotFound.New(),
			want: true,
		},
		{
			name: strconv.Itoa(http.StatusBadRequest),
			err:  InvalidParameter.New(),
			want: true,
		},
		{
			name: strconv.Itoa(http.StatusInternalServerError),
			err:  SystemDefault.New(),
			want: false,
		},
		{
			name: strconv.Itoa(http.StatusBadRequest) + " info",
			err:  InvalidParameter.New().Info(),
			want: true,
		},
		{
			name: strconv.Itoa(http.StatusBadRequest) + " wrap info",
			err:  InvalidParameter.Wrap(InvalidParameter.New().Info()),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.IsInfo(); got != tt.want {
				t.Errorf("IsInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appError_IsPanic(t *testing.T) {
	tests := []struct {
		name string
		err  AppError
		want bool
	}{
		{
			name: strconv.Itoa(http.StatusNotFound),
			err:  NotFound.New(),
			want: false,
		},
		{
			name: strconv.Itoa(http.StatusInternalServerError),
			err:  SystemDefault.New(),
			want: false,
		},
		{
			name: strconv.Itoa(http.StatusInternalServerError) + " panic",
			err:  SystemPanic.New().Panic(),
			want: true,
		},
		{
			name: strconv.Itoa(http.StatusInternalServerError) + " wrap panic",
			err:  SystemPanic.Wrap(InvalidParameter.New().Panic()),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.IsPanic(); got != tt.want {
				t.Errorf("IsPanic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appError_IsWarn(t *testing.T) {
	tests := []struct {
		name string
		err  AppError
		want bool
	}{
		{
			name: strconv.Itoa(http.StatusNotFound),
			err:  NotFound.New(),
			want: false,
		},
		{
			name: strconv.Itoa(http.StatusBadRequest),
			err:  InvalidParameter.New(),
			want: false,
		},
		{
			name: strconv.Itoa(http.StatusInternalServerError),
			err:  SystemDefault.New(),
			want: false,
		},
		{
			name: strconv.Itoa(http.StatusBadRequest) + " warn",
			err:  SystemUnknown.New().Warn(),
			want: true,
		},
		{
			name: strconv.Itoa(http.StatusBadRequest) + " wrap warn",
			err:  SystemUnknown.Wrap(SystemUnknown.New().Warn()),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.IsWarn(); got != tt.want {
				t.Errorf("IsWarn() = %v, want %v", got, tt.want)
			}
		})
	}
}
