package clientCore

import (
	"testing"
	"anonymous-messaging/packet_format"
	"github.com/stretchr/testify/assert"
	"anonymous-messaging/publics"
	"os"
	"fmt"
	"reflect"
	"strconv"
)

var mixClient MixClient
var mixPubs []publics.MixPubs

func TestMain(m *testing.M) {
	mixClient = *NewMixClient("MixClient", 1, 0)

	m1 := publics.MixPubs{"Mix1", "localhost", "3330", 0}
	m2 := publics.MixPubs{"Mix2", "localhost", "3331", 0}
	mixPubs = []publics.MixPubs{m1, m2}

	code := m.Run()
	os.Exit(code)
}

func TestMixClientEncode(t *testing.T) {

	message := "Hello world"
	path := mixPubs
	delays := []float64{1.4, 2.5, 2.3}

	encoded := mixClient.EncodeMessage(message, path, delays)
	expected := packet_format.Encode(message, path, delays)
	assert.Equal(t, encoded, expected, "The packets should be the same")
}

func TestMixClientDecode(t *testing.T) {

	packet := packet_format.NewPacket("Message", []float64{0.1, 0.2, 0.3}, mixPubs, nil)

	decoded := mixClient.DecodeMessage(packet)
	expected := packet_format.Decode(packet)

	assert.Equal(t, decoded, expected, "The packets should be the same")
}

func TestGenerateDelaySequence(t *testing.T) {
	delays := mixClient.GenerateDelaySequence(100, 5)
	if len(delays) != 5 {
		t.Error("Wrong length")
	}
	if reflect.TypeOf(delays).Elem().Kind() != reflect.Float64 {
		t.Error("Incorrect type of generated delays")
	}
}

func TestGetRandomMixSequence(t *testing.T) {
	// test two cases: the one when len is smaller than all mixes and the one when length is larger / the same
	var mixes []publics.MixPubs
	for i:=0; i < 5; i++ {
		mixes = append(mixes, publics.NewMixPubs(fmt.Sprintf("Mix%d", i), "localhost", strconv.Itoa(3330+i), int64(i)))
	}

	var sequence []publics.MixPubs
	sequence = mixClient.GetRandomMixSequence(mixes, 6)
	assert.Equal(t, 5, len(sequence), "When the given length is larger than the number of active nodes, the path should be " +
		"the sequence of all active mixes")

	sequence = mixClient.GetRandomMixSequence(mixes, 3)
	assert.Equal(t, 3, len(sequence), "When the given length is larger than the number of active nodes, the path should be " +
		"the sequence of all active mixes")
}