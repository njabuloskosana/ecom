<template>
    <div class="welcome-container">
      <div class="welcome-card">
        <h1>Welcome, {{ userInformation.userName }}!</h1>
        <p>Your account is active.</p>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { computed } from "vue";
  import { useRoute } from "vue-router";
  
  // Simulate fetching user information from local storage
  const userInformation = computed(() => {
    const storedUser = localStorage.getItem("userInformation");
    return storedUser ? JSON.parse(storedUser) : { userName: "Guest" };
  });
  </script>
  
  <style scoped>
  .welcome-container {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    background-color: #f5f5f5;
  }
  
  .welcome-card {
    background: #ffffff;
    padding: 30px;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    width: 100%;
    max-width: 400px;
    text-align: center;
  }
  
  .welcome-card h1 {
    font-size: 24px;
    font-weight: bold;
    color: #333;
  }
  
  .welcome-card p {
    font-size: 16px;
    color: #666;
  }
  </style>