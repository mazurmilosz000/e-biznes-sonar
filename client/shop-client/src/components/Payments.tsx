import { useCart, CartItem } from "../context/CartContext";
import { api } from "../api/api";

const Payments = () => {
  const { cart, clearCart } = useCart();

  const handleCheckout = async (cartItems: CartItem[]) => {
    try {
      await api.post("http://localhost:8080/checkout", {
        items: cartItems.map(item => ({
          product_id: item.id,
          name: item.name,
          price: item.price,
          quantity: item.quantity,
        }))
      });
      alert("Order placed successfully!");
      clearCart();
    } catch (error) {
      console.error("Error during checkout:", error);
    }
  };

  const handleClick = () => {
    handleCheckout(cart);
  };

  return (
    <div>
      <h2>Checkout</h2>
      <button onClick={handleClick} disabled={cart.length === 0}>
        Pay Now
      </button>
    </div>
  );
};

export default Payments;