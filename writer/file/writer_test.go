package file

import (
	"testing"
)

type mockFile struct {
	a string
}

func (s *mockFile) Name() string {
	return ""
}

func (s *mockFile) Write(p []byte) (n int, err error) {
	s.a = string(p)

	return 0, nil
}

func TestNew(t *testing.T) {
	type args struct {
		file *mockFile
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "stdout",
			args: args{
				file: &mockFile{},
			},
			want: "test",
		},
		{
			name: "stderr",
			args: args{
				file: &mockFile{},
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := New(tt.args.file)

			w.Write("Test1")

			if tt.args.file.a != "Test1" {
				t.Error("Пошел нахуй")
			}
		})
	}
}
