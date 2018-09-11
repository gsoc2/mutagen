package encoding

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/havoc-io/mutagen/pkg/url"
)

// TestProtocolBuffersCycle tests a Protocol Buffers marshal/save/load/unmarshal
// cycle.
func TestProtocolBuffersCycle(t *testing.T) {
	// Create an empty temporary file and defer its cleanup.
	file, err := ioutil.TempFile("", "mutagen_encoding")
	if err != nil {
		t.Fatal("unable to create temporary file:", err)
	} else if err = file.Close(); err != nil {
		t.Fatal("unable to close temporary file:", err)
	}
	defer os.Remove(file.Name())

	// Create a Protocol Buffers message that we can test with.
	message := &url.URL{
		Protocol: url.Protocol_SSH,
		Username: "George",
		Hostname: "washington",
		Port:     1776,
		Path:     "/by/land/or/by/sea",
	}
	if err := MarshalAndSaveProtobuf(file.Name(), message); err != nil {
		t.Fatal("unable to marshal and save Protocol Buffers message:", err)
	}

	// Reload the message.
	decoded := &url.URL{}
	if err := LoadAndUnmarshalProtobuf(file.Name(), decoded); err != nil {
		t.Fatal("unable to load and unmarshal Protocol Buffers message:", err)
	}

	// Verify that contents were preserved. We have to explicitly compare
	// members because the Protocol Buffers generated struct contains fields
	// that can't be compared by value.
	match := decoded.Protocol == message.Protocol &&
		decoded.Username == message.Username &&
		decoded.Hostname == message.Hostname &&
		decoded.Port == message.Port &&
		decoded.Path == message.Path
	if !match {
		t.Error("decoded Protocol Buffers message did not match original:", decoded, "!=", message)
	}
}

const (
	// testProtobufEncodingNMessages is the number of messages to send/receive
	// in TestProtobufEncoding.
	testProtobufEncodingNMessages = 100
)

func TestProtobufEncoding(t *testing.T) {
	// Create a buffer to use as our stream.
	stream := &bytes.Buffer{}

	// Create an encoder/decoder pair.
	encoder := NewProtobufEncoder(stream)
	decoder := NewProtobufDecoder(stream)

	// Set test message parameters.
	protocol := url.Protocol_SSH
	username := "George"
	hostname := "washington"
	path := "/by/land/or/by/sea"

	// Write a sequence of SSH URL messages with increasing port values.
	message := &url.URL{
		Protocol: protocol,
		Username: username,
		Hostname: hostname,
		Path:     path,
	}
	for i := 0; i < testProtobufEncodingNMessages; i++ {
		message.Port = uint32(i)
		if err := encoder.Encode(message); err != nil {
			t.Fatal("unable to encode message:", err)
		}
	}

	// Read a sequence of URL messages and verify their port values.
	for i := 0; i < testProtobufEncodingNMessages; i++ {
		*message = url.URL{}
		if err := decoder.Decode(message); err != nil {
			t.Fatal("unable to decode message:", err)
		} else if message.Protocol != protocol {
			t.Error("protocol mismatch in received message")
		} else if message.Username != username {
			t.Error("username mismatch in received message")
		} else if message.Hostname != hostname {
			t.Error("hostname mismatch in received message")
		} else if message.Port != uint32(i) {
			t.Error("hostname mismatch in received message")
		} else if message.Path != path {
			t.Error("path mismatch in received message")
		}
	}
}
