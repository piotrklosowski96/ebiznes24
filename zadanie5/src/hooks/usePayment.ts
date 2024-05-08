import { useState } from "react";
import { PaymentResponse, PaymentsService } from "../client";

type UsePaymentReturn = [Payment[], () => void]

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

	return [ payments, createNewPayment ]
}