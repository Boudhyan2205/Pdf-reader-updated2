//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package sigutil ;import (_g "bytes";_gb "crypto";_c "crypto/x509";_dgf "encoding/asn1";_ab "encoding/pem";_d "errors";_ga "fmt";_bc "github.com/unidoc/timestamp";_gg "github.com/unidoc/unipdf/v3/common";_aa "golang.org/x/crypto/ocsp";_b "io";_dg "io/ioutil";
_bb "net/http";_ae "time";);

// TimestampClient represents a RFC 3161 timestamp client.
// It is used to obtain signed tokens from timestamp authority servers.
type TimestampClient struct{

// HTTPClient is the HTTP client used to make timestamp requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_bb .Client ;};

// NewTimestampRequest returns a new timestamp request based
// on the specified options.
func NewTimestampRequest (body _b .Reader ,opts *_bc .RequestOptions )(*_bc .Request ,error ){if opts ==nil {opts =&_bc .RequestOptions {};};if opts .Hash ==0{opts .Hash =_gb .SHA256 ;};if !opts .Hash .Available (){return nil ,_c .ErrUnsupportedAlgorithm ;
};_bbe :=opts .Hash .New ();if _ ,_efc :=_b .Copy (_bbe ,body );_efc !=nil {return nil ,_efc ;};return &_bc .Request {HashAlgorithm :opts .Hash ,HashedMessage :_bbe .Sum (nil ),Certificates :opts .Certificates ,TSAPolicyOID :opts .TSAPolicyOID ,Nonce :opts .Nonce },nil ;
};

// CRLClient represents a CRL (Certificate revocation list) client.
// It is used to request revocation data from CRL servers.
type CRLClient struct{

// HTTPClient is the HTTP client used to make CRL requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_bb .Client ;};

// GetIssuer retrieves the issuer of the provided certificate.
func (_bg *CertClient )GetIssuer (cert *_c .Certificate )(*_c .Certificate ,error ){for _ ,_bba :=range cert .IssuingCertificateURL {_bbb ,_aaf :=_bg .Get (_bba );if _aaf !=nil {_gg .Log .Debug ("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074 \u0064\u006f\u0077\u006e\u006c\u006f\u0061\u0064\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0066\u006f\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066ic\u0061\u0074\u0065\u0020\u0025\u0076\u003a\u0020\u0025\u0076",cert .Subject .CommonName ,_aaf );
continue ;};return _bbb ,nil ;};return nil ,_ga .Errorf ("\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063a\u0074e\u0020\u006e\u006f\u0074\u0020\u0066\u006fu\u006e\u0064");};

// NewTimestampClient returns a new timestamp client.
func NewTimestampClient ()*TimestampClient {return &TimestampClient {HTTPClient :_cgc ()}};

// MakeRequest makes a CRL request to the specified server and returns the
// response. If a server URL is not provided, it is extracted from the certificate.
func (_eg *CRLClient )MakeRequest (serverURL string ,cert *_c .Certificate )([]byte ,error ){if _eg .HTTPClient ==nil {_eg .HTTPClient =_cgc ();};if serverURL ==""{if len (cert .CRLDistributionPoints )==0{return nil ,_d .New ("\u0063e\u0072\u0074i\u0066\u0069\u0063\u0061t\u0065\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070ec\u0069\u0066\u0079 \u0061\u006ey\u0020\u0043\u0052\u004c\u0020\u0073e\u0072\u0076e\u0072\u0073");
};serverURL =cert .CRLDistributionPoints [0];};_fe ,_dc :=_eg .HTTPClient .Get (serverURL );if _dc !=nil {return nil ,_dc ;};defer _fe .Body .Close ();_dfe ,_dc :=_dg .ReadAll (_fe .Body );if _dc !=nil {return nil ,_dc ;};if _fc ,_ :=_ab .Decode (_dfe );
_fc !=nil {_dfe =_fc .Bytes ;};return _dfe ,nil ;};

// NewOCSPClient returns a new OCSP client.
func NewOCSPClient ()*OCSPClient {return &OCSPClient {HTTPClient :_cgc (),Hash :_gb .SHA1 }};

// NewCertClient returns a new certificate client.
func NewCertClient ()*CertClient {return &CertClient {HTTPClient :_cgc ()}};

// GetEncodedToken executes the timestamp request and returns the DER encoded
// timestamp token bytes.
func (_ega *TimestampClient )GetEncodedToken (serverURL string ,req *_bc .Request )([]byte ,error ){if serverURL ==""{return nil ,_ga .Errorf ("\u006d\u0075\u0073\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061m\u0070\u0020\u0073\u0065\u0072\u0076\u0065r\u0020\u0055\u0052\u004c");
};if req ==nil {return nil ,_ga .Errorf ("\u0074\u0069\u006de\u0073\u0074\u0061\u006dp\u0020\u0072\u0065\u0071\u0075\u0065\u0073t\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c");};_ad ,_dbe :=req .Marshal ();if _dbe !=nil {return nil ,_dbe ;
};_ee :=_ega .HTTPClient ;if _ee ==nil {_ee =_cgc ();};_egaf ,_dbe :=_ee .Post (serverURL ,"a\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0074\u0069\u006d\u0065\u0073t\u0061\u006d\u0070-\u0071u\u0065\u0072\u0079",_g .NewBuffer (_ad ));
if _dbe !=nil {return nil ,_dbe ;};defer _egaf .Body .Close ();_dfb ,_dbe :=_dg .ReadAll (_egaf .Body );if _dbe !=nil {return nil ,_dbe ;};if _egaf .StatusCode !=_bb .StatusOK {return nil ,_ga .Errorf ("\u0075\u006e\u0065x\u0070\u0065\u0063\u0074e\u0064\u0020\u0048\u0054\u0054\u0050\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0064",_egaf .StatusCode );
};var _de struct{Version _dgf .RawValue ;Content _dgf .RawValue ;};if _ ,_dbe =_dgf .Unmarshal (_dfb ,&_de );_dbe !=nil {return nil ,_dbe ;};return _de .Content .FullBytes ,nil ;};

// Get retrieves the certificate at the specified URL.
func (_ce *CertClient )Get (url string )(*_c .Certificate ,error ){if _ce .HTTPClient ==nil {_ce .HTTPClient =_cgc ();};_e ,_df :=_ce .HTTPClient .Get (url );if _df !=nil {return nil ,_df ;};defer _e .Body .Close ();_dgff ,_df :=_dg .ReadAll (_e .Body );
if _df !=nil {return nil ,_df ;};if _gd ,_ :=_ab .Decode (_dgff );_gd !=nil {_dgff =_gd .Bytes ;};_f ,_df :=_c .ParseCertificate (_dgff );if _df !=nil {return nil ,_df ;};return _f ,nil ;};

// CertClient represents a X.509 certificate client. Its primary purpose
// is to download certificates.
type CertClient struct{

// HTTPClient is the HTTP client used to make certificate requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_bb .Client ;};

// OCSPClient represents a OCSP (Online Certificate Status Protocol) client.
// It is used to request revocation data from OCSP servers.
type OCSPClient struct{

// HTTPClient is the HTTP client used to make OCSP requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_bb .Client ;

// Hash is the hash function  used when constructing the OCSP
// requests. If zero, SHA-1 will be used.
Hash _gb .Hash ;};func _cgc ()*_bb .Client {return &_bb .Client {Timeout :5*_ae .Second }};

// MakeRequest makes a OCSP request to the specified server and returns
// the parsed and raw responses. If a server URL is not provided, it is
// extracted from the certificate.
func (_cf *OCSPClient )MakeRequest (serverURL string ,cert ,issuer *_c .Certificate )(*_aa .Response ,[]byte ,error ){if _cf .HTTPClient ==nil {_cf .HTTPClient =_cgc ();};if serverURL ==""{if len (cert .OCSPServer )==0{return nil ,nil ,_d .New ("\u0063e\u0072\u0074i\u0066\u0069\u0063a\u0074\u0065\u0020\u0064\u006f\u0065\u0073 \u006e\u006f\u0074\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0079\u0020\u004f\u0043S\u0050\u0020\u0073\u0065\u0072\u0076\u0065\u0072\u0073");
};serverURL =cert .OCSPServer [0];};_ag ,_ef :=_aa .CreateRequest (cert ,issuer ,&_aa .RequestOptions {Hash :_cf .Hash });if _ef !=nil {return nil ,nil ,_ef ;};_ea ,_ef :=_cf .HTTPClient .Post (serverURL ,"\u0061p\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u006fc\u0073\u0070\u002d\u0072\u0065\u0071\u0075\u0065\u0073\u0074",_g .NewReader (_ag ));
if _ef !=nil {return nil ,nil ,_ef ;};defer _ea .Body .Close ();_dbd ,_ef :=_dg .ReadAll (_ea .Body );if _ef !=nil {return nil ,nil ,_ef ;};if _dd ,_ :=_ab .Decode (_dbd );_dd !=nil {_dbd =_dd .Bytes ;};_fb ,_ef :=_aa .ParseResponseForCert (_dbd ,cert ,issuer );
if _ef !=nil {return nil ,nil ,_ef ;};return _fb ,_dbd ,nil ;};

// NewCRLClient returns a new CRL client.
func NewCRLClient ()*CRLClient {return &CRLClient {HTTPClient :_cgc ()}};

// IsCA returns true if the provided certificate appears to be a CA certificate.
func (_ca *CertClient )IsCA (cert *_c .Certificate )bool {return cert .IsCA &&_g .Equal (cert .RawIssuer ,cert .RawSubject );};