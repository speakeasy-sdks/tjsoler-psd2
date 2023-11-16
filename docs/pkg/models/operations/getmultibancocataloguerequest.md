# GetMultibancoCatalogueRequest


## Fields

| Field                                                                                                                                                                                                                       | Type                                                                                                                                                                                                                        | Required                                                                                                                                                                                                                    | Description                                                                                                                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Digest`                                                                                                                                                                                                                    | *string*                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                          | Es contenido si viaja el campo Signature. Ej: Digest: SHA-256=NzdmZjA4YjY5M2M2NDYyMmVjOWFmMGNmYTZiNTU3MjVmNDI4NTRlMzJkYzE3ZmNmMDE3ZGFmMjhhNTc5OTU3OQ==                                                                      |
| `Signature`                                                                                                                                                                                                                 | *string*                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                          | Firma de la petición por el TPP.                                                                                                                                                                                            |
| `TPPSignatureCertificate`                                                                                                                                                                                                   | *string*                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                          | Certificado del TPP usado para firmar la petición, en base64, sin cabecera, pie ni saltos de linea. Ej: TPP-Signature-Certificate: MIIHgzCCBmugAwIBAgIIZzZvBQlt0UcwDQYJ………….KoZIhvcNAQELBQAwSTELMAkGA1UEBhMCVVMxEzARBgNVBA  |
| `XRequestID`                                                                                                                                                                                                                | *string*                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                          | Identificador único de la transacción asignado por el TPP. Ej: X-Request-ID: 1b3ab8e8-0fd5-43d2-946e-d75958b172e7                                                                                                           |
| `AspName`                                                                                                                                                                                                                   | *string*                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                          | Nombre del ASPSP al que desea realizar la petición.                                                                                                                                                                         |
| `InstructedAmount`                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                          | Importe de pago                                                                                                                                                                                                             |
| `MultibancoPaymentType`                                                                                                                                                                                                     | *string*                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                          | Tipo de pago multibanco                                                                                                                                                                                                     |
| `PaymentReference`                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                          | Referencia de la operación                                                                                                                                                                                                  |
| `RequestedExecutionDate`                                                                                                                                                                                                    | **string*                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                          | Fecha de la ejecución del pago                                                                                                                                                                                              |