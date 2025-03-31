import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Products from "./components/Products";
import Cart from "./components/Cart";
import Payments from "./components/Payments";
import { CartProvider } from "./context/CartContext";

const App = () => {
  return (
    <CartProvider>
      <Router>
        <nav>
          <Link to="/">Products</Link> | <Link to="/cart">Cart</Link> | <Link to="/checkout">Checkout</Link>
        </nav>
        <Routes>
          <Route path="/" element={<Products />} />
          <Route path="/cart" element={<Cart />} />
          <Route path="/checkout" element={<Payments />} />
        </Routes>
      </Router>
    </CartProvider>
  );
};

export default App;