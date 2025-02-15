<template>
    <div class="coming-soon-wrapper">
      <div class="coming-soon">
        <h1>Coming Soon</h1>
        <p>
          We are working hard to bring you a new experience. Stay tuned for updates and be the first to know when we launch!
        </p>
        <!-- Ad Image Section -->
        <div class="ad-container">
          <img src="../assets/refferalposter.png" alt="Advertisement" />
        </div>
        <!-- Email Subscription Form -->
        <form class="email-form" @submit.prevent="handleSubmit">
          <input
            type="email"
            v-model="email"
            placeholder="Enter your email address"
            required
          />
          <button type="submit">Notify Me</button>
        </form>
        <div class="footer">
          &copy; 2025 Your Company. All rights reserved.
        </div>
      </div>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from "vue";
  export default defineComponent( {
    name: 'ComingSoon',
    data() {
      return {
        email: '',
      }
    },
    methods: {
      handleSubmit() {
        // Handle email subscription logic here (e.g., API call)
        console.log('Email submitted:', this.email);
        // Clear the input field after submission
        this.email = '';
      },
    },
  });
  </script>
  
  <style scoped>
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }
  
  /* Wrapper that replaces the :host styles */
  .coming-soon-wrapper {
    height: 100vh;
    background: linear-gradient(135deg, #6a11cb, #2575fc);
    display: flex;
    align-items: center;
    justify-content: center;
    text-align: center;
    padding: 20px;
    font-family: 'Open Sans', sans-serif;
    color: #fff;
  }
  
  .coming-soon {
    max-width: 500px;
    width: 100%;
  }
  
  .coming-soon h1 {
    font-size: 3rem;
    font-weight: 600;
    margin-bottom: 20px;
  }
  
  .coming-soon p {
    font-size: 1.125rem;
    margin-bottom: 30px;
    line-height: 1.5;
  }
  
  /* Ad Image Styling */
  .ad-container {
    margin: 30px 0;
    display: flex;
    justify-content: center;
    align-items: center;
    overflow: hidden;
  }
  
  .ad-container img {
    max-width: 100%;
    max-height: 300px; /* Adjust this value as needed */
    width: auto;
    height: auto;
    object-fit: contain;
    border-radius: 8px;
  }
  
  .email-form {
    display: flex;
    justify-content: center;
  }
  
  .email-form input {
    padding: 12px 15px;
    border: none;
    border-radius: 4px 0 0 4px;
    font-size: 1rem;
    outline: none;
  }
  
  .email-form button {
    padding: 12px 20px;
    border: none;
    background-color: #fff;
    color: #2575fc;
    font-size: 1rem;
    font-weight: 600;
    border-radius: 0 4px 4px 0;
    cursor: pointer;
    transition: background-color 0.3s ease;
  }
  
  .email-form button:hover {
    background-color: #f1f1f1;
  }
  
  .footer {
    margin-top: 40px;
    font-size: 0.875rem;
    opacity: 0.8;
  }
  </style>
  
  
  