# ResponseGetStatusMultibankPayment


## Fields

| Field                                                                                                                                                                            | Type                                                                                                                                                                             | Required                                                                                                                                                                         | Description                                                                                                                                                                      | Example                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `FundsAvailable`                                                                                                                                                                 | **bool*                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                               | Este dato es contenido si es soportado por el ASPSP, si una confirmación de fondos ha sido realizada y si el "transactionStatus" es alguno de los siguientes •ACTC•  ACWC•  ACCP | true                                                                                                                                                                             |
| `PsuMessage`                                                                                                                                                                     | **string*                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                               | Texto enviado al TPP a través del HUB para ser mostrado al PSU.                                                                                                                  | Mensaje de ejemplo                                                                                                                                                               |
| `TppMessages`                                                                                                                                                                    | [*ResponseGetStatusMultibankPaymentTppMessages](../../models/shared/responsegetstatusmultibankpaymenttppmessages.md)                                                             | :heavy_minus_sign:                                                                                                                                                               | Mensaje para el TPP enviado a través del HUB.                                                                                                                                    |                                                                                                                                                                                  |
| `TransactionStatus`                                                                                                                                                              | *string*                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                               | Estado de la transacción.                                                                                                                                                        | ACCP                                                                                                                                                                             |