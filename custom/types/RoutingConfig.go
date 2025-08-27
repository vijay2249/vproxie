package types

import (
	"errors"
	"log"
)

type RoutingConfig struct {
	Routing []Routing `yaml:"routingConfig"`
}

type Routing struct {
	Domain string `yaml:"domain"`
	DefaultServer string `yaml:"default"`
	SubdomainRouting map[string]string `yaml:"subdomainRouting"`
	EndpointRouting map[string]string `yaml:"endpointRouting"`
}

func EmptyRoutingStruct() Routing { return Routing {"", "", make(map[string]string, 0), make(map[string]string, 0)} }

/*
   -------------------------    TODO 
	 1. Add interface to give user to write custom function to get service url from subdomain and endpoint configuration
*/

func (config *RoutingConfig) GetDomainConfigDetails(domain string) (Routing, error) {
	for _, routeConfig := range config.Routing {
		if routeConfig.Domain == domain {
			return routeConfig, nil
		}
	}
	log.Printf("no such config found for domain: %s\n", domain)
	return EmptyRoutingStruct(), errors.New("no such config found for domain: " + domain + "\n")
}

func (config *RoutingConfig) GetAllDomains() (domains []string) {
	if config == nil {
		log.Fatalf("no domains mentioned in config to forward requests to, please check configuration file again.")
	}
	for _, routing := range config.Routing {
		domains = append(domains, routing.Domain)
	}
	return domains
}

func (config *Routing) GetServiceURLBySubdomain(subdomain string) (string, error) {
	if routeConfig, ok := config.SubdomainRouting[subdomain]; ok {
		return routeConfig, nil
	} else {
		log.Printf("no such subdomain config found for subdomain: %s\n", subdomain)
    return "", errors.New("no such subdomain config found for subdomain: " + subdomain + "\n")
	}
}

func (config *Routing) GetServiceURLByEndPoint(endpoint string) (string, error) {
	if routeConfig, ok := config.EndpointRouting[endpoint]; ok {
		return routeConfig, nil
	} else {
		log.Printf("no such endpoint config found for endpoint: %s\n", endpoint)
		return "", errors.New("no such endpoint config found for endpoint: " + endpoint + "\n")
	}
}

func (config *Routing) GetDefaultServiceURL() string { return config.DefaultServer }

func (config *RoutingConfig) GetRoutingByDomain(domain, subdomain string) (string, error) {
	routing, err := config.GetDomainConfigDetails(domain)
	if err != nil {
		log.Printf("error getting routing for domain: %s\n", domain)
		return "", err
	}
	serviceURL, err := routing.GetServiceURLBySubdomain(subdomain)
	if err != nil {
		log.Printf("error getting service url for subdomain: %s\n", subdomain)
    return "", err
	}
	return serviceURL, nil
}

func (config *RoutingConfig) GetRoutingByEndpoint(domain, endpoint string) (string, error) {
	routing, err := config.GetDomainConfigDetails(domain)
	if err != nil {
		log.Printf("error getting routing for domain: %s\n", domain)
		return "", err
	}
	serviceURL, err := routing.GetServiceURLByEndPoint(endpoint)
	if err != nil {
		log.Printf("error getting service url for endpoint: %s\n", endpoint)
    return "", err
	}
	return serviceURL, nil
}


type Port struct {
	Port string `yaml:"port"`
}

func (config *Port) GetPortNumber() string { return config.Port }