// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ResponseAccountTransactionsLinks - Lista de hipervínculos para ser reconocidos por el TPP. Tipos soportados en esta respuesta: download: Enlace de descarga de los datos de la consulta realizada, cuando los datos devueltos tienen un tamaño grande. Solo para camt-data.
type ResponseAccountTransactionsLinks struct {
}

// ResponseAccountTransactionsAccounts - Identificador de la cuenta que se está consultando. Nota: recomendado usarlo ya que podría pasar a parámetro obligatorio en futuras versiones.
type ResponseAccountTransactionsAccounts struct {
}

// ResponseAccountTransactionsBalances - Una lista de balances con respecto a una cuenta.
type ResponseAccountTransactionsBalances struct {
}

// ResponseAccountTransactionsTppMessage - Mensaje para el TPP enviado a través del HUB.
type ResponseAccountTransactionsTppMessage struct {
}

// ResponseAccountTransactionsTransactions - Devolución de los datos en formato JSON, cuando los datos devueltos tienen un tamaño pequeño. Este reporte contiene las transacciones resultantes de los parámetros de consulta.
type ResponseAccountTransactionsTransactions struct {
}

// ResponseAccountTransactions - HTTP/1.1 200 Ok
type ResponseAccountTransactions struct {
	// Lista de hipervínculos para ser reconocidos por el TPP. Tipos soportados en esta respuesta: download: Enlace de descarga de los datos de la consulta realizada, cuando los datos devueltos tienen un tamaño grande. Solo para camt-data.
	Links *ResponseAccountTransactionsLinks
	// Identificador de la cuenta que se está consultando. Nota: recomendado usarlo ya que podría pasar a parámetro obligatorio en futuras versiones.
	Accounts *ResponseAccountTransactionsAccounts
	// Una lista de balances con respecto a una cuenta.
	Balances *ResponseAccountTransactionsBalances
	// Texto enviado al TPP a través del HUB para ser mostrado al PSU.
	PsuMessage *string
	// Mensaje para el TPP enviado a través del HUB.
	TppMessage *ResponseAccountTransactionsTppMessage
	// Devolución de los datos en formato JSON, cuando los datos devueltos tienen un tamaño pequeño. Este reporte contiene las transacciones resultantes de los parámetros de consulta.
	Transactions *ResponseAccountTransactionsTransactions
}

func (o *ResponseAccountTransactions) GetLinks() *ResponseAccountTransactionsLinks {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *ResponseAccountTransactions) GetAccounts() *ResponseAccountTransactionsAccounts {
	if o == nil {
		return nil
	}
	return o.Accounts
}

func (o *ResponseAccountTransactions) GetBalances() *ResponseAccountTransactionsBalances {
	if o == nil {
		return nil
	}
	return o.Balances
}

func (o *ResponseAccountTransactions) GetPsuMessage() *string {
	if o == nil {
		return nil
	}
	return o.PsuMessage
}

func (o *ResponseAccountTransactions) GetTppMessage() *ResponseAccountTransactionsTppMessage {
	if o == nil {
		return nil
	}
	return o.TppMessage
}

func (o *ResponseAccountTransactions) GetTransactions() *ResponseAccountTransactionsTransactions {
	if o == nil {
		return nil
	}
	return o.Transactions
}
