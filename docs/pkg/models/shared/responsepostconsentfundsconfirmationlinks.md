# ResponsePostConsentFundsConfirmationLinks

Lista de hipervínculos para ser reconocidos por el HUB. Tipos soportados en esta respuesta: scaRedirect: en caso de SCA por redirección. Link donde el navegador del PSU debe ser redireccionado por el Hub. startAuthorisation: en caso de que un inicio explícito de la autorización de la transacción sea necesario (no hay selección del método SCA) startAuthorisationWithAuthenticationMethodSelection: link al end-point de autorización donde el sub-recurso de autorización tiene que ser generado mientras se selecciona el método SCA. Este enlace es contenido bajo las mismas condiciones que el campo "scaMethods". self: link al recurso de inicio de pago creado por esta petición. status: link para recuperar el estado de la transacción. scaStatus: link para consultar el estado SCA correspondiente al sub-recurso de autorización. Este link es solo contenido si un sub-recurso de autorización ha sido creado. (Si no hay autorización explícita este enlace debe venir obligatoriamente)


## Fields

| Field       | Type        | Required    | Description |
| ----------- | ----------- | ----------- | ----------- |