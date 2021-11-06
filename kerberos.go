package winrm

import (
    "net"
    "net/http"
    "net/url"

    "github.com/dpotapov/go-spnego"
    "github.com/masterzen/winrm/soap"
)

// ClientKerberos provides a transport via Kerberos
type ClientKerberos struct {
    clientRequest
}

// ClientKerberosWithNoCanonicalize provides a transport via Kerberos, which do not
//  do forward & reverse DNS queries. Useful when you cant resolve reverse
// Needs PR#5 from github.com/dpotapov/go-spnego (or temporarily use github.com/yo000/go-spnego)
type ClientKerberosWithNoCanonicalize struct {
    clientRequest
	NoCanonicalize bool
}

// Transport creates the wrapped Kerberos transport
func (c *ClientKerberos) Transport(endpoint *Endpoint) error {
    c.clientRequest.Transport(endpoint)
    c.clientRequest.transport = &spnego.Transport{}
    return nil
}

// Transport creates the wrapped KerberosWithNoCanonicalize transport
func (c *ClientKerberosWithNoCanonicalize) Transport(endpoint *Endpoint) error {
    c.clientRequest.Transport(endpoint)
    c.clientRequest.transport = &spnego.Transport{NoCanonicalize: true}
    return nil
}

// Post make post to the winrm soap service (forwarded to clientRequest implementation)
func (c ClientKerberos) Post(client *Client, request *soap.SoapMessage) (string, error) {
    return c.clientRequest.Post(client, request)
}

//NewClientKerberosWithDial NewClientKerberosWithDial
func NewClientKerberosWithDial(dial func(network, addr string) (net.Conn, error)) *ClientKerberos {
    return &ClientKerberos{
        clientRequest{
            dial: dial,
        },
    }
}

//NewClientKerberosWithProxyFunc NewClientKerberosWithProxyFunc
func NewClientKerberosWithProxyFunc(proxyfunc func(req *http.Request) (*url.URL, error)) *ClientKerberos {
    return &ClientKerberos{
        clientRequest{
            proxyfunc: proxyfunc,
        },
    }
}

