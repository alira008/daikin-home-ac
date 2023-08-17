package server

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Mode uint8

const (
	Auto   Mode = 0
	Dry    Mode = 2
	Cooler Mode = 3
	Heat   Mode = 4
	Fan    Mode = 6
)

type FanSpeed uint8

const (
	One       FanSpeed = 3
	Two       FanSpeed = 4
	Three     FanSpeed = 5
	Four      FanSpeed = 6
	Five      FanSpeed = 7
	Automatic FanSpeed = 10
	Silent    FanSpeed = 11
)

type DaikinState struct {
	Temperature int      `json:"temperature"`
	Mode        Mode     `json:"mode"`
	Timer       bool     `json:"timer"`
	TimerDelay  int      `json:"timerDelay"`
	Power       bool     `json:"power"`
	FanSpeed    FanSpeed `json:"fanSpeed"`
	Swing       bool     `json:"swing"`
	Powerful    bool     `json:"powerful"`
	Econo       bool     `json:"econo"`
	Comfort     bool     `json:"comfort"`
}

func NewDaikinState() DaikinState {
	return DaikinState{
		Mode:     Cooler,
		FanSpeed: One,
	}
}

type Frame []byte

func encodeFrames(frames []Frame) (string, error) {
	var stringBuilder strings.Builder
	var err error
	zeroHigh := "\t\t450\t419\n"
	oneHigh := "\t\t450\t1286\n"
	gap := "\t\t450\t10000\n"
	_, err = stringBuilder.WriteString("begin remote\n\tname\tdaikin\n\tflags\tRAW_CODES\n\teps\t30\n\taeps\t100\n\tgap\t10000\n\tbegin\traw_codes\n\t\tname\tsignal\n")
	if err != nil {
		return "", err
	}

	// Loop through frames
	for _, f := range frames {
		// Loop through bytes of frame
		for _, b := range f {
			// Loop through bits of frame
			for j := 0; j < 8; j++ {
				// Check the j-th bit using a mask
				bit := (b >> j) & 1
				var code string
				if bit == 1 {
					code = oneHigh
				} else {
					code = zeroHigh
				}

				_, err = stringBuilder.WriteString(code)
				if err != nil {
					return "", err
				}
			}
		}

		// add gap between frames
		_, err = stringBuilder.WriteString(gap)
		if err != nil {
			return "", err
		}
	}

	_, err = stringBuilder.WriteString("\tend raw_codes\nend remote")
	if err != nil {
		return "", err
	}

	return stringBuilder.String(), err
}

func createMessage(ds *DaikinState) []Frame {
	var frames []Frame

	frame1 := initiateFrame(8)
	frame1[4] = 0xc5
	frame1[5] = 0x00

	if ds.Comfort {
		frame1[6] = 0x10
		frame1[7] = 0xe7
	} else {
		frame1[6] = 0x00
		frame1[7] = 0xd7
	}

	frame2 := initiateFrame(8)
	frame2[4] = 0x42
	frame2[5] = 0x00
	frame2[6] = 0x00
	frame2[7] = 0x54

	frame3 := createFrame3(ds)

	frames = append(frames, frame1)
	frames = append(frames, frame2)
	frames = append(frames, frame3)
	return frames
}

func initiateFrame(size int) Frame {
	frame := make(Frame, size)
	frame[0] = 0x11
	frame[1] = 0xda
	frame[2] = 0x27
	frame[3] = 0x00

	return frame
}

func createFrame3(ds *DaikinState) Frame {
	frame := initiateFrame(19)
	frame[4] = 0x00

	// Set bit for mode
	frame[5] |= byte(ds.Mode) << 4

	// Set bits for timer
	if ds.Timer {
		frame[5] |= 0x2
	} else {
		frame[5] |= 0x4
	}

	// Set bits for power
	if ds.Power {
		frame[5] |= 1
	} else {
		frame[5] &= ^byte(1)
	}

	// Set bits for temperature
	// take celsius and multiply by 2
	frame[6] |= byte(ds.Temperature * 2)

	// Set bits for fan speed
	frame[8] |= byte(ds.FanSpeed) << 4

	// Set bits for swing
	if ds.Swing {
		frame[8] |= 0xF
	}

	// Set bits for timer
	frame[10] |= byte(ds.TimerDelay)
	frame[11] |= byte(ds.TimerDelay >> 4)
	frame[12] |= byte(ds.TimerDelay >> 8)

	// Set bits for powerful
	if ds.Powerful {
		frame[13] |= 0x1
	}

	// Always 0xC1
	frame[15] = 0xc1

	// Set bits for Econo
	if ds.Econo {
		frame[16] = 0x84
	} else {
		frame[16] = 0x80
	}

	// Calculate checksum
	// add all previous values of the frame
	checksum := 0
	for i := 0; i < 18; i++ {
		checksum += int(frame[i])
	}
	frame[18] = byte(checksum) & 0xFF

	return frame
}

func transmitLircMessage(msg string) {
	// Write message into temp file
	f, err := os.Create("/tmp/daikin.lircd.conf")
	if err != nil {
		fmt.Println("error creating temp file for message: ", err.Error())
		return
	}
	_, err = f.WriteString(msg)
	if err != nil {
		fmt.Println("error writing to temp file for message: ", err.Error())
		return
	}

	// copy to /etc/lirc/lircd.conf.d/
	cmd := exec.Command("sudo", "cp", "/tmp/daikin.lircd.conf", "/etc/lirc/lircd.conf.d/")
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error copying temp lirc file: ", err.Error())
		return
	}

	// Restart lirc
	cmd = exec.Command("sudo", "systemctl", "restart", "lircd")
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error restarting lirc: ", err.Error())
		return
	}

	// Transmit message
	cmd = exec.Command("irsend", "SEND_ONCE", "daikin", "signal")
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error transmitting message: ", err.Error())
		return
	}
}

func (ds *DaikinState) Send() {
	frames := createMessage(ds)
	msg, err := encodeFrames(frames)
	if err != nil {
		fmt.Println("Coudl not encode frames")
		return
	}

	fmt.Println(msg)
	// transmitLircMessage(msg)
}
