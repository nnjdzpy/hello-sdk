package hello

func NewClient() *Client {
	return &Client{Hello: "hello-sdk"}
}

type Client struct {
	Hello string
}
