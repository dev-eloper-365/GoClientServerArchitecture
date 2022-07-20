package audio

import (
	"bufio"
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func Player(connection net.Conn) (err error) {

	reader := bufio.NewReader(connection)
	audio_raw, err := reader.ReadString('\n')
	audio := strings.TrimSpace(audio_raw)

	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	println(audio)

	if audio == "stop" {
		return
	} else {
		if _, err := os.Stat(mydir + "/" + audio); err == nil {
			//if file exists
			nbytes, err := connection.Write([]byte("correct\n"))
			if err != nil {
				panic(err)
			}
			if error := run(audio, connection); error != nil {
				log.Fatal(error)
			}
			nbytes, erro := connection.Write([]byte("[+] Played audio\n"))
			if erro != nil {
				panic(erro)
			}
			fmt.Println("[+]", nbytes, "bytes written\n")
			return nil
		} else {
			println("[-]File doesn't exist")
			// path/to/whatever does *not* exist
			nbytes, err := connection.Write([]byte("incorrect\n"))
			if err != nil {
				panic(err)
			}
			fmt.Println("[+]", nbytes, "bytes written\n")

			return nil
		}
	}
}

func run(audio string, connection net.Conn) error {
	f, err := os.Open(audio)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, ready, err := oto.NewContext(d.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-ready

	p := c.NewPlayer(d)
	defer p.Close()
	p.Play()

	fmt.Printf("Length: %d[bytes]\n", d.Length())
	for {
		for range time.Tick(time.Second * 1) {
			reader := bufio.NewReader(connection)
			raw_response, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			response := strings.TrimSpace(raw_response)
			println(response + "#")
			if response == "" {
				if p.IsPlaying() {
					return nil
				}
			}
		}
	}
	return nil
}
