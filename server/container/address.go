package container

import (
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddressesContainer 获得当前所在环境的各个局域网IP
// 1 获取该环境所在的所有地址，包括tcp/udp
// 2 找出ip和非回环地址
// 3 append到result
func AddressesContainer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []string
		adders, err := net.InterfaceAddrs()
		if err != nil {
			log.Fatal(err)
		}
		for _, addr := range adders {
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					result = append(result, ipNet.IP.String())
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"addresses": result,
		})
	}
}
