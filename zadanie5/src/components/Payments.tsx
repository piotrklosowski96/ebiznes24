import usePayment from "../hooks/usePayment.ts";

export default function Payments() {
	const [ payments, createNewPayment ] = usePayment();

	return (
		<div>
			<h1>Payments</h1>
			<div className="payments">
				{
					payments?.map((payment: Payment) => (
						<div style={{border: "1px solid", padding: "10px", display: "block"}}>
							<h2>{payment.id}</h2>
							<h3>{payment.status}</h3>
						</div>
					))
				}
				<button onClick={createNewPayment}>Create new payment</button>
			</div>
		</div>
	)
}