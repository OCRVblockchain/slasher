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

func ShowRemoved(identityResponse *msp.IdentityResponse) {
	fmt.Printf("Removed identity:\nID: %s\nAffiliation: %s\nAttributes: %v\nType: %s\nMaxEnrollments: %d\nSecret: %s\nCAName: %s\n",
		identityResponse.ID,
		identityResponse.Affiliation,
		identityResponse.Attributes,
		identityResponse.Type,
		identityResponse.MaxEnrollments,
		identityResponse.Secret,
		identityResponse.CAName)
}
