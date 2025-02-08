<template>
    <div v-if="isLoading" class="loader-container">
      <div class="loader-content">
        <div class="spinner"></div>
        <p v-if="loadingMessage" class="loading-message">{{ loadingMessage }}</p>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { defineProps } from "vue";
  
  // Define props for the loader
  defineProps({
    isLoading: {
      type: Boolean,
      required: true,
      default: false,
    },
    loadingMessage: {
      type: String,
      default: "Loading...", // Default loading message
    },
  });
  </script>
  
  <style scoped>
  /* Loader container */
  .loader-container {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: rgba(255, 255, 255, 0.9); /* Semi-transparent white background */
    backdrop-filter: blur(8px); /* Soften the background */
    z-index: 9999; /* Ensure it's on top of everything */
  }
  
  /* Loader content wrapper */
  .loader-content {
    text-align: center;
  }
  
  /* Spinner styles */
  .spinner {
    border: 8px solid rgba(0, 123, 255, 0.2); /* Subtle blue border */
    border-top: 8px solid #007bff; /* Highlighted blue border */
    border-radius: 50%;
    width: 80px;
    height: 80px;
    animation: spin 1.2s linear infinite;
  
    /* Add glowing effect */
    box-shadow: 
      0 0 15px rgba(0, 123, 255, 0.9), /* Inner glow */
      0 0 30px rgba(0, 123, 255, 0.7), /* Outer glow */
      0 0 50px rgba(0, 123, 255, 0.5); /* Larger glow */
  }
  
  /* Keyframes for spinning animation */
  @keyframes spin {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }
  
  /* Loading message */
  .loading-message {
    margin-top: 16px;
    font-size: 16px;
    font-weight: bold;
    color: #333;
  }
  </style>