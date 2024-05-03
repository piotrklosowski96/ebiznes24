import { Link, Outlet } from 'react-router-dom';
import './LayoutWrapper.css'

export default function LayoutWrapper() {
	return (
		<>
			<div className='navbar'>
				<Link to="/">Home</Link>
				<Link to="/products">Products</Link>
				<Link to="/payments">Payments</Link>
				<Link to="/cart">Cart</Link>
			</div>

			<Outlet/>
		</>
	)
}
