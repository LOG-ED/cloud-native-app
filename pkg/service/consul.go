package service

import (
	"fmt"
	"net"
	consul "github.com/hashicorp/consul/api"
)

func (s *Service) registerConsul() {
	c, err := consul.NewClient(&consul.Config{
		Address: "consul:8500",
	})
	if err != nil {
		panic("Failed to connect to Consul agent")
	}
	s.ConsulAgent = c.Agent()

	ifaces, err := net.Interfaces()
	// handle err
	var ip net.IP
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			fmt.Println(ip.String())
		}
	}

	serviceDef := &consul.AgentServiceRegistration{
		Name: s.Name,
		Address: ip.String(),
		Port: s.Port,
		Check: &consul.AgentServiceCheck{
			TTL: s.TTL.String(),
		},
	}

	if err := s.ConsulAgent.ServiceRegister(serviceDef); err != nil {
		panic("Failed to register Service")
	}
}