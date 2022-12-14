package tls

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"
)

var validityNotAfter = time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)

// Issue returns CA certificate, TLS certificate and TLS private key
func Issue(hosts []string) (caCert []byte, tlsCert []byte, tlsKey []byte, err error) {
	rsaBits := 2048
	priv, err := rsa.GenerateKey(rand.Reader, rsaBits)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("generate rsa key: %w", err)
	}
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("generate serial number for root: %w", err)
	}
	subject := pkix.Name{
		Organization: []string{"Root CA"},
	}
	caTemplate := x509.Certificate{
		SerialNumber:          serialNumber,
		Subject:               subject,
		NotBefore:             time.Now(),
		NotAfter:              validityNotAfter,
		KeyUsage:              x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("generate CA certificate: %w", err)
	}
	certOut := &bytes.Buffer{}
	err = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if err != nil {
		return nil, nil, nil, fmt.Errorf("encode CA certificate: %w", err)
	}
	cert := certOut.Bytes()

	serialNumber, err = rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("generate serial number for client: %w", err)
	}
	subject = pkix.Name{
		Organization: []string{"PGO"},
	}
	issuer := pkix.Name{
		Organization: []string{"Root CA"},
	}
	tlsTemplate := x509.Certificate{
		SerialNumber:          serialNumber,
		Subject:               subject,
		Issuer:                issuer,
		NotBefore:             time.Now(),
		NotAfter:              validityNotAfter,
		DNSNames:              hosts,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA:                  false,
	}
	clientKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("generate client key: %w", err)
	}
	tlsDerBytes, err := x509.CreateCertificate(rand.Reader, &tlsTemplate, &caTemplate, &clientKey.PublicKey, priv)
	if err != nil {
		return nil, nil, nil, err
	}
	tlsCertOut := &bytes.Buffer{}
	err = pem.Encode(tlsCertOut, &pem.Block{Type: "CERTIFICATE", Bytes: tlsDerBytes})
	if err != nil {
		return nil, nil, nil, fmt.Errorf("encode TLS  certificate: %w", err)
	}
	tlsCert = tlsCertOut.Bytes()

	keyOut := &bytes.Buffer{}
	block := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(clientKey)}
	err = pem.Encode(keyOut, block)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("encode RSA private key: %w", err)
	}
	privKey := keyOut.Bytes()

	return cert, tlsCert, privKey, nil
}
