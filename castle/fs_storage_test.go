package castle

import (
	"os"
	"path"
	"reflect"
	"testing"
)

func TestFSBackend_Read(t *testing.T) {
	_ = os.Mkdir(path.Join(os.TempDir(), "flying_castle"), os.ModeDir)
	file, err := os.Create(path.Join(path.Join(os.TempDir(), "flying_castle", "test")))
	if err != nil {
		panic(err)
	}
	_, _ = file.Write([]byte("hello"))
	_ = file.Close()
	type fields struct {
		dataPath string
	}
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "NoError", fields: struct{ dataPath string }{dataPath: path.Join(os.TempDir(), "flying_castle")}, args: struct{ fileName string }{fileName: path.Join(path.Join(os.TempDir(), "flying_castle", "test"))}, want: []byte("hello"), wantErr: false},
		{name: "FileNotFound", fields: struct{ dataPath string }{dataPath: path.Join(os.TempDir(), "flying_castle")}, args: struct{ fileName string }{fileName: "test"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			F := FSBackend{
				dataPath: tt.fields.dataPath,
			}
			got, err := F.Read(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
	_ = os.RemoveAll(path.Join(os.TempDir(), "flying_castle"))
}

func TestFSBackend_Write(t *testing.T) {
	_ = os.Mkdir(path.Join(os.TempDir(), "flying_castle"), os.ModeDir)
	type fields struct {
		dataPath string
	}
	type args struct {
		fileName  string
		chunkData []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{name: "NoError", fields: fields{dataPath: path.Join(os.TempDir(), "flying_castle")}, args: args{
			fileName:  "test",
			chunkData: []byte("hello"),
		}, want: path.Join(path.Join(os.TempDir(), "flying_castle"), "test"), wantErr: false},
		{name: "IncorrectDatapath", fields: fields{dataPath: "/flying_castle_nope"}, args: args{
			fileName:  "test",
			chunkData: []byte("hello"),
		}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			F := FSBackend{
				dataPath: tt.fields.dataPath,
			}
			got, err := F.Write(tt.args.fileName, tt.args.chunkData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Write() got = %v, want %v", got, tt.want)
			}
		})
	}
}
