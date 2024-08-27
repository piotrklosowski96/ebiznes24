import { useState } from "react";
import { PaymentResponse, PaymentsService } from "../client";

type UsePaymentReturn = [Payment[], () => void, () => void]

export default function usePayment(): UsePaymentReturn {
	const [payments, setPayments] = useState<Payment[]>([]);

	const createNewPayment = () => {
		PaymentsService.createPayment({body: {status: "INITIATING"}}).then((response: PaymentResponse) => {
			const newPayment: Payment = {
				id: response.id,
				status: response.status,
			}
			setPayments([
				...payments,
				newPayment
			])
		})
	}

	const getPayments = () => {
		PaymentsService.getPayments({
			limit: 1000,
			offset: 0,
		}).then(paymentsResponse => {
			const payments = paymentsResponse.payments?.map(p => {
				return {
					id: p.id,
					status: p.status,
				} as Payment
			}) || []

			setPayments(payments)
		})
	}

	return [ payments, getPayments, createNewPayment ]
}