package dnsimple

import (
	"fmt"
	"strconv"
)

// CertificatesService handles communication with the certificate related
// methods of the DNSimple API.
//
// See https://developer.dnsimple.com/v2/domains/certificates
type CertificatesService struct {
	client *Client
}

// Certificate represents a Certificate in DNSimple.
type Certificate struct {
	ID                  int    `json:"id,omitempty"`
	DomainID            int    `json:"domain_id,omitempty"`
	CommonName          string `json:"common_name,omitempty"`
	Years               int    `json:"years,omitempty"`
	State               string `json:"state,omitempty"`
	AuthorityIdentifier string `json:"authority_identifier,omitempty"`
	CreatedAt           string `json:"created_at,omitempty"`
	UpdatedAt           string `json:"updated_at,omitempty"`
	ExpiresOn           string `json:"expires_on,omitempty"`

	CertificateRequest string `json:"csr,omitempty"`
}

func certificatePath(accountID, domainIdentifier, certificateID string) string {
	path := fmt.Sprintf("%v/certificates", domainPath(accountID, domainIdentifier))

	if certificateID != "" {
		path += fmt.Sprintf("/%v", certificateID)
	}
	return path
}

// CertificateResponse represents a response from an API method that returns a Certificate struct.
type CertificateResponse struct {
	Response
	Data *Certificate `json:"data"`
}

// CertificatesResponse represents a response from an API method that returns a collection of Certificate struct.
type CertificatesResponse struct {
	Response
	Data []Certificate `json:"data"`
}

// ListCertificates list the certificates for a domain.
//
// See https://developer.dnsimple.com/v2/domains/certificates#list
func (s *CertificatesService) ListCertificates(accountID, domainIdentifier string, options *ListOptions) (*CertificatesResponse, error) {
	path := versioned(certificatePath(accountID, domainIdentifier, ""))
	certificatesResponse := &CertificatesResponse{}

	path, err := addURLQueryOptions(path, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.get(path, certificatesResponse)
	if err != nil {
		return certificatesResponse, err
	}

	certificatesResponse.HttpResponse = resp
	return certificatesResponse, nil
}

// GetCertificate fetches the certificate.
//
// See https://developer.dnsimple.com/v2/domains/certificates#get
func (s *CertificatesService) GetCertificate(accountID, domainIdentifier string, certificateID int) (*CertificateResponse, error) {
	path := versioned(certificatePath(accountID, domainIdentifier, strconv.Itoa(certificateID)))
	certificateResponse := &CertificateResponse{}

	resp, err := s.client.get(path, certificateResponse)
	if err != nil {
		return nil, err
	}

	certificateResponse.HttpResponse = resp
	return certificateResponse, nil
}
