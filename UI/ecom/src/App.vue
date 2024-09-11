<template>
  <ProductList :products="products" />
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
    console.log('Fetched products:', products.value); // Check the actual value here
  } catch (error) {
    console.error('Failed to fetch products', error);
  }
};

onMounted(fetchProducts);
</script>

<style scoped>
/* Add your styles here */
</style>