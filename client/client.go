package client


import  ("log"
	"github.com/sgjp/go-coap"
)

func OutTuple(tuple string) string{
	rv := sendCoapMsg("/out",tuple,coap.PUT,coap.Confirmable)
	return string(rv.Payload)
}

func RdTuple(tuple string) string{
	rv := sendCoapMsg("/rd",tuple,coap.GET,coap.Confirmable)
	return string(rv.Payload)
}

func InTuple(tuple string) string{
	rv := sendCoapMsg("/in",tuple,coap.GET,coap.Confirmable)
	return string(rv.Payload)
}


func sendCoapMsg(path string, payload string, coapCode coap.COAPCode, coapType coap.COAPType) *coap.Message{
	req := coap.Message{
		Type:      coapType,
		Code:      coapCode,
		MessageID: 12345,
		Payload:   []byte(payload),
	}


	req.SetOption(coap.MaxAge, 3)
	req.SetPathString(path)

	c, err := coap.Dial("udp", "localhost:5683")
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	rv, err := c.Send(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	if rv != nil {
		//log.Printf("Response payload: %s", rv.Payload)
	}
	return rv
}
