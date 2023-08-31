// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RequestPeriodicPayment struct {
	// La fecha se refiere a la zona horaria del ASPSP
	DayOfExecution *string `json:"dayOfExecution,omitempty"`
	// El último día aplicable de ejecución. Si no viene se trata de una orden permanente sin fin.
	EndDate *string `json:"endDate,omitempty"`
	// Define el comportamiento cuando las fechas del pago recurrente caen en fin de semana o festivo. Entonces el pago se ejecuta el día laboral anterior o posterior
	ExecutionRule *string `json:"executionRule,omitempty"`
	// La frecuencia del pago recurrente resultante de esta orden permanente
	Frequency *string `json:"frequency,omitempty"`
	// El primer día aplicable de ejecución desde esta fecha es el primer pago
	StartDate string `json:"startDate"`
}

func (o *RequestPeriodicPayment) GetDayOfExecution() *string {
	if o == nil {
		return nil
	}
	return o.DayOfExecution
}

func (o *RequestPeriodicPayment) GetEndDate() *string {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *RequestPeriodicPayment) GetExecutionRule() *string {
	if o == nil {
		return nil
	}
	return o.ExecutionRule
}

func (o *RequestPeriodicPayment) GetFrequency() *string {
	if o == nil {
		return nil
	}
	return o.Frequency
}

func (o *RequestPeriodicPayment) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}
