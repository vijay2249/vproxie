package types

import (
	"errors"
	"log"

	// "github.com/vijay2249/vproxie/constant"
)

type RoutingConfig struct {
	Routing []Routing `yaml:"routingConfig"`
}

type Routing struct {
	Domain string `yaml:"domain"`
	SubdomainRouting map[string]string `yaml:"subdomainRouting"`
	EndpointRouting map[string]string `yaml:"endpointRouting"`
}

func EmptyRoute() Routing {
	return Routing{"", make(map[string]string, 0), make(map[string]string, 0)}
}

func init(){
	necessaryFlags := log.Ldate|log.LUTC|log.Lshortfile
	log.SetFlags(necessaryFlags)
}

func (config *RoutingConfig) GetRoutingDetailsByDomain(domain string) (Routing, error) {
	for _, routing := range config.Routing {
		if routing.Domain == domain {
			return routing, nil
		}
	}
	log.Printf("no config present to route traffic for domain: %s\n", domain)
	return EmptyRoute(), errors.ErrUnsupported
}

func (config *Routing) GetURLBySubdomain(subdomain string) (string, error) {
	if url, ok := config.SubdomainRouting[subdomain]; ok {
		return url, nil
	}
	log.Printf("no config to present to route %s subdomain request\n", subdomain)
	return "", errors.ErrUnsupported
}

func (config *Routing) GetURLByEndpoint(endpoint string) (string, error) {
	if url, ok := config.EndpointRouting[endpoint]; ok {
		return url, nil
	}
	log.Printf("no config to present to route %s endpoint request\n", endpoint)
	return "", errors.ErrUnsupported
}

func (config *RoutingConfig) GetBackendServiceURLBySubdomain(domain, subdomain string) (string, error) {
	routing, err := config.GetRoutingDetailsByDomain(domain)
	if err != nil {
		log.Printf("please revalidate config values")
		return "", errors.ErrUnsupported
	}

	url, err := routing.GetURLBySubdomain(subdomain)
	if err != nil {
		log.Printf("please revalidate config values")
		return "", errors.ErrUnsupported
	}
	return url, nil
}

func (config *RoutingConfig) GetBackendServiceURLByEndpoint(domain, endpoint string) (string, error) {
	routing, err := config.GetRoutingDetailsByDomain(domain)
	if err != nil {
		log.Printf("please revalidate config values")
		return "", errors.ErrUnsupported
	}

	url, err := routing.GetURLByEndpoint(endpoint)
	if err != nil {
		log.Printf("please revalidate config values")
		return "", errors.ErrUnsupported
	}
	return url, nil
}