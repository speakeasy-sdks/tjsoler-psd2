# ResponseAccountTransactions


## Fields

| Field                                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                              | Example                                                                                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Links`                                                                                                                                                                                                                                  | [*shared.ResponseAccountTransactionsLinks](../../../pkg/models/shared/responseaccounttransactionslinks.md)                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                       | Lista de hipervínculos para ser reconocidos por el TPP. Tipos soportados en esta respuesta: download: Enlace de descarga de los datos de la consulta realizada, cuando los datos devueltos tienen un tamaño grande. Solo para camt-data. |                                                                                                                                                                                                                                          |
| `Accounts`                                                                                                                                                                                                                               | [*shared.Accounts](../../../pkg/models/shared/accounts.md)                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                       | Identificador de la cuenta que se está consultando. Nota: recomendado usarlo ya que podría pasar a parámetro obligatorio en futuras versiones.                                                                                           |                                                                                                                                                                                                                                          |
| `Balances`                                                                                                                                                                                                                               | [*shared.ResponseAccountTransactionsBalances](../../../pkg/models/shared/responseaccounttransactionsbalances.md)                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                       | Una lista de balances con respecto a una cuenta.                                                                                                                                                                                         |                                                                                                                                                                                                                                          |
| `PsuMessage`                                                                                                                                                                                                                             | **string*                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                       | Texto enviado al TPP a través del HUB para ser mostrado al PSU.                                                                                                                                                                          | Informacion para PSU                                                                                                                                                                                                                     |
| `TppMessage`                                                                                                                                                                                                                             | [*shared.ResponseAccountTransactionsTppMessage](../../../pkg/models/shared/responseaccounttransactionstppmessage.md)                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                       | Mensaje para el TPP enviado a través del HUB.                                                                                                                                                                                            |                                                                                                                                                                                                                                          |
| `Transactions`                                                                                                                                                                                                                           | [*shared.Transactions](../../../pkg/models/shared/transactions.md)                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                       | Devolución de los datos en formato JSON, cuando los datos devueltos tienen un tamaño pequeño. Este reporte contiene las transacciones resultantes de los parámetros de consulta.                                                         |                                                                                                                                                                                                                                          |