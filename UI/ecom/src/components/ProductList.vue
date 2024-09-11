<template>
  <div class="container">
    <!-- Products Section on the left (60%) -->
    <div class="products-section">
      <h1>Products</h1>
      <div v-if="products.length === 0" class="no-products">No products available.</div>
      <div v-else class="product-grid">
        <div v-for="product in products" :key="product.id" class="product-card">
          <div class="product-image">
            <img :src="`/assets/images/${product.image}`" alt="Product Image" />
          </div>
          <div class="product-details">
            <h3>{{ product.name }}</h3>
            <p class="description">{{ product.description }}</p>
            <p class="price">Price: ${{ product.price.toFixed(2) }}</p>
            <input v-model.number="quantities[product.id]" type="number" min="1" class="quantity-input" />
            <button @click="addToCart(product)" class="add-to-cart-btn">Add to Cart</button>
          </div>
        </div>
      </div>
    </div>

    <!-- Cart Section on the right (30%) -->
    <div class="cart-section">
      <h2>Your Cart</h2>
      <div v-if="cart.length === 0" class="empty-cart">Your cart is empty.</div>
      <ul v-else class="cart-list">
        <li v-for="item in cart" :key="item.id" class="cart-item">
          <div class="cart-item-info">
            <h3>{{ item.name }}</h3>
            <p>Quantity: {{ item.quantity }}</p>
            <p>Total Price: ${{ (item.price * item.quantity).toFixed(2) }}</p>
          </div>
          <button @click="removeFromCart(item.id)" class="remove-item-btn">Remove</button>
        </li>
      </ul>

      <!-- Cart Total -->
      <div v-if="cart.length > 0" class="cart-total">
        <h3>Total: ${{ cartTotal.toFixed(2) }}</h3>
      </div>

      <div v-if="cart.length > 0" class="cart-actions">
        <button @click="clearCart" class="clear-cart-btn">Clear Cart</button>
        <button @click="makePayment" class="make-payment-btn">Make Payment</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed } from 'vue';

interface Product {
  id: number;
  name: string;
  description: string;
  image: string;
  price: number;
}

interface CartItem {
  id: number;
  name: string;
  price: number;
  quantity: number;
}

export default defineComponent({
  name: 'ProductList',
  props: {
    products: {
      type: Array as () => Product[],
      required: true
    }
  },
  setup() {
    const cart = ref<CartItem[]>([]);
    const quantities = ref<{ [key: number]: number }>({}); // Quantity for each product

    const addToCart = (product: Product) => {
      const quantity = Math.max(1, quantities.value[product.id] || 1); // Ensure quantity is at least 1
      const existingItem = cart.value.find(item => item.id === product.id);
      if (existingItem) {
        existingItem.quantity += quantity;
      } else {
        cart.value.push({
          id: product.id,
          name: product.name,
          price: product.price,
          quantity: quantity
        });
      }
      quantities.value[product.id] = 1; // Reset quantity input for this product
    };

    const removeFromCart = (productId: number) => {
      cart.value = cart.value.filter(item => item.id !== productId);
    };

    const clearCart = () => {
      cart.value = [];
    };

    const makePayment = () => {
      console.log('Payment made for the following items:', cart.value);
      clearCart();
    };

    // Calculate the cart total
    const cartTotal = computed(() => {
      return cart.value.reduce((total, item) => total + item.price * item.quantity, 0);
    });

    return {
      cart,
      quantities,
      addToCart,
      removeFromCart,
      clearCart,
      makePayment,
      cartTotal
    };
  }
});
</script>

<style scoped>
.container {
  display: flex;
  justify-content: space-between;
}

.products-section {
  width: 60%;
}

.cart-section {
  width: 30%;
  border-left: 1px solid #ccc;
  padding-left: 20px;
}

.product-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
}

.product-card {
  border: 1px solid #ccc;
  padding: 16px;
  text-align: center;
}

.product-image img {
  width: 100%;
  height: auto;
}

.quantity-input {
  width: 60px;
  margin-bottom: 8px;
}

.add-to-cart-btn {
  background-color: #4CAF50;
  color: white;
  border: none;
  padding: 8px 16px;
  cursor: pointer;
}

.add-to-cart-btn:hover {
  background-color: #45a049;
}

.cart-list {
  list-style: none;
  padding: 0;
}

.cart-item {
  margin-bottom: 16px;
}

.remove-item-btn {
  background-color: #f44336;
  color: white;
  border: none;
  padding: 8px 16px;
  cursor: pointer;
}

.remove-item-btn:hover {
  background-color: #e53935;
}

.cart-total {
  margin-top: 16px;
  font-size: 18px;
  font-weight: bold;
}

.cart-actions {
  display: flex;
  justify-content: space-between;
  margin-top: 16px;
}

.clear-cart-btn,
.make-payment-btn {
  background-color: #ff9800;
  color: white;
  border: none;
  padding: 8px 16px;
  cursor: pointer;
}

.make-payment-btn {
  background-color: #4CAF50;
}

.clear-cart-btn:hover,
.make-payment-btn:hover {
  background-color: #45a049;
}
</style>
