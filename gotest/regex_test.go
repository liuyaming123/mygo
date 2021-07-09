package gotest

import (
	"encoding/json"
	"fmt"
	"regexp"
	"testing"
)

// AudioLoudnorm 音频响度信息
type AudioLoudnorm struct {
	InputI            string `json:"input_i"`
	InputTP           string `json:"input_tp"`
	InputLRA          string `json:"input_lra"`
	InputThresh       string `json:"input_thresh"`
	OutputI           string `json:"output_i"`
	OutputTP          string `json:"output_tp"`
	OutputLRA         string `json:"output_lra"`
	OutputThresh      string `json:"output_thresh"`
	NormalizationType string `json:"normalization_type"`
	TargetOffset      string `json:"target_offset"`
}

func TestRegex(t *testing.T) {
	text := `
	the heartbeart url: http://127.0.0.1:8090/live/liveService/live/transProHeartBeat/770
	[tcp @ 0x80ac700] [2021-07-09 16:17:24]   Connection to tcp://127.0.0.1:8090 failed: Connection refused
	[2021-07-09 16:17:24]   heartbeat sends failed, last time: 290s, current time: 292s
	[2021-07-09 16:17:24]   size=N/A time=00:04:53.97 bitrate=N/A speed=3.56x processed=97.99%    
	size=N/A time=00:04:55.87 bitrate=N/A speed=3.56x processed=98.62%    
	the heartbeart url: http://127.0.0.1:8090/live/liveService/live/transProHeartBeat/770
	[tcp @ 0x7314680] [2021-07-09 16:17:26]   Connection to tcp://127.0.0.1:8090 failed: Connection refused
	[2021-07-09 16:17:26]   heartbeat sends failed, last time: 292s, current time: 295s
	[2021-07-09 16:17:26]   size=N/A time=00:04:59.97 bitrate=N/A speed=3.57x processed=99.99%    
	[2021-07-09 16:17:26]   the heartbeart url: http://127.0.0.1:8090/live/liveService/live/transProHeartBeat/770
	[tcp @ 0x7b412c0] [2021-07-09 16:17:26]   Connection to tcp://127.0.0.1:8090 failed: Connection refused
	[2021-07-09 16:17:26]   heartbeat sends failed, last time: 295s, current time: 299s
	[2021-07-09 16:17:26]   size=N/A time=00:05:00.00 bitrate=N/A speed=3.57x processed=100.00%    
	video:0kB audio:35798kB subtitle:0kB other streams:0kB global headers:0kB muxing overhead: unknown
	[Parsed_loudnorm_0 @ 0x72e66c0] [2021-07-09 16:17:26]   
	{
		"input_i" : "-28.51",
		"input_tp" : "-3.30",
		"input_lra" : "13.20",
		"input_thresh" : "-40.02",
		"output_i" : "-22.82",
		"output_tp" : "-2.00",
		"output_lra" : "8.50",
		"output_thresh" : "-34.03",
		"normalization_type" : "dynamic",
		"target_offset" : "-0.18"
	}
`
	compile, err := regexp.Compile("{(.|\n)*?input_i(.|\n)*?target_offset(.|\n)*?}")

	if err != nil {
		fmt.Println(err)
	}

	res := compile.FindString(text)

	fmt.Println(res)
	// fmt.Println([]byte(res))

	audioLoudnorm1 := &AudioLoudnorm{}

	err = json.Unmarshal([]byte(res), audioLoudnorm1)
	if err != nil {
		fmt.Println("json Unmarshal failed!")
		return
	}
	fmt.Printf("audioLoudnorm1------------------------> %#v\n", audioLoudnorm1)

	info, err := json.Marshal(audioLoudnorm1)
	if err != nil {
		fmt.Println("json Marshal failed!")
		return
	}
	fmt.Printf("%s", info)
}
