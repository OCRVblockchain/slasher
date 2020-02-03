package helpers

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

func ShowRevoked(RevocationResponse *msp.RevocationResponse) {
	fmt.Println("Revoked certs:")
	for index, revokedCert := range RevocationResponse.RevokedCerts {
		fmt.Printf("%d. Serial: %s, AKI: %s\n", index, revokedCert.Serial, revokedCert.AKI)
	}
}
