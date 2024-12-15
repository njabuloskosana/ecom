<template>
  <div id="app">
    <div class="app-heading">Cyber Spaza</div>
    <ProductList :products="products" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import ProductList from './components/ProductList.vue';
import axios from 'axios';

interface Product {
  id: number;
  name: string;
  description: string;
  image: string;
  quantity: number;
  created_time: string;
  price: number;
}

const products = ref<Product[]>([]);
const url = 'http://localhost:8080/api/v1/products';

const fetchProducts = async () => {
  try {
    const response = await axios.get<Product[]>(url);
    products.value = response.data;
    console.log('Fetched products:', products.value);
  } catch (error) {
    console.error('Failed to fetch products', error);
  }
};

onMounted(fetchProducts);
</script>

<!-- Global styles (not scoped) -->
<style>
html, body {
  margin: 0;
  padding: 0;
  width: 100%;
  background-color: #fafafa;
  font-family: "Helvetica Neue", Avenir, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  background-image: url('../src/assets/cyberstorebackground.jpg');
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-size: cover;
  background-position: center center;
}

body {
  margin-left: 0;
  margin-bottom: 0;
}
</style>

<!-- Component-scoped styles -->
<style scoped>
#app {
  display: flex;
  flex-direction: column;
  width: 100vw !important;
  background-color: transparent !important;
  position: relative;
  padding: 20px !important;
}

.app-heading {
  font-family: 'Montserrat', sans-serif;
  font-size: 3rem;
  color: #8afc7e; /* neon accent color from your existing design */
  text-transform: uppercase;
  letter-spacing: 2px;
  text-align: center;
  margin: 40px 0;
  position: relative;
  overflow: hidden;
  /* Neon glow */
  text-shadow: 0 0 5px #8afc7e, 0 0 10px #8afc7e, 0 0 20px #8afc7e, 0 0 40px #8afc7e;
}

/* Add a subtle "scanline" effect via a pseudo-element overlay */
.app-heading::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: repeating-linear-gradient(
    to bottom,
    transparent 0%,
    transparent 2%,
    rgba(255, 255, 255, 0.02) 2.5%,
    transparent 3%
  );
  pointer-events: none;
  animation: scan 5s linear infinite;
}

/* Subtle glitch effect using pseudo-elements */
.app-heading::after {
  content: "Cyber Spaza";
  position: absolute;
  top: 0;
  left: 0;
  color: #8afc7e;
  opacity: 0.7;
  text-shadow: 0 0 5px #8afc7e, 0 0 10px #8afc7e;
  clip: rect(0, 0, 0, 0);
  animation: glitch 2s infinite;
}

/* Keyframes for the scanline overlay */
@keyframes scan {
  0% { background-position: 0 0; }
  100% { background-position: 0 100%; }
}

/* Keyframes for a gentle glitch effect */
@keyframes glitch {
  0% {
    clip: rect(0, 0, 0, 0);
  }
  10% {
    clip: rect(0, 1000px, 0, 0);
    transform: translate(-5px, -2px);
  }
  20% {
    clip: rect(0, 0, 1000px, 0);
    transform: translate(5px, 2px);
  }
  30% {
    clip: rect(0, 1000px, 1000px, 0);
    transform: translate(-2px, -1px);
  }
  40% {
    clip: rect(0, 0, 0, 0);
    transform: translate(2px, 1px);
  }
  100% {
    clip: rect(0, 0, 0, 0);
    transform: translate(0,0);
  }
}
</style>
