# AuthenticationObject


## Fields

| Field                                                                                                                                                                                                                | Type                                                                                                                                                                                                                 | Required                                                                                                                                                                                                             | Description                                                                                                                                                                                                          | Example                                                                                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `AuthenticationMethodID`                                                                                                                                                                                             | *string*                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                   | Id del método de autenticación proporcionado por el ASPSP.                                                                                                                                                           | SMS_OTP                                                                                                                                                                                                              |
| `AuthenticationType`                                                                                                                                                                                                 | *string*                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                   | Tipo del método de autenticación. Ver anexo 9.6 Tipos de autenticación para más información.                                                                                                                         | SMS_OTP                                                                                                                                                                                                              |
| `AuthenticationVersion`                                                                                                                                                                                              | **string*                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                   | Versión de la herramienta asociada al authenticationType.                                                                                                                                                            | 1.0                                                                                                                                                                                                                  |
| `Explanation`                                                                                                                                                                                                        | **string*                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                   | Información detallada acerca del método SCA para el PSU.                                                                                                                                                             |                                                                                                                                                                                                                      |
| `Name`                                                                                                                                                                                                               | *string*                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                   | Nombre del método de autenticación definido por el PSU en la banca online del ASPSP. Alternativamente podría ser una descripción proporcionada por el ASPSP. Si el TPP lo tiene disponible, debe presentarlo al PSU. | SMS OTP al teléfono 666777888                                                                                                                                                                                        |