package speech

import "testing"

func TestWitHandler_SendWitMessage(t *testing.T) {
	key := "FK733Z7UEV75KBHPBMEDPVSFBULTAWO2"

	type fields struct {
		witKey string
	}
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Running successful message test",
			fields:  fields{witKey: key},
			args:    args{"hello"},
			wantErr: false,
			want:    `{"_text"`,
		},
		{
			name:    "Running wrong api key test",
			fields:  fields{witKey: "wrong"},
			args:    args{"hello"},
			want:    `{"error"`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wh := &WitHandler{
				witKey: tt.fields.witKey,
			}
			got, err := wh.SendWitMessage(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("WitHandler.SendWitMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got[:8] != tt.want {
				t.Errorf("WitHandler.SendWitMessage() = %v", got[:8])
			}
		})
	}
}
