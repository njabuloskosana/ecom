<template>
    <div class="form-container">
      <h2>Driver Registration</h2>
      <p class="subtitle">Create your driver profile to connect with bike owners and find opportunities</p>
      
      <form @submit.prevent="submitForm">
        <div class="form-row">
          <div class="form-group">
            <label for="firstName">First Name</label>
            <input 
              type="text" 
              id="firstName" 
              v-model="driver.firstName" 
              placeholder="Enter your first name"
              required
            />
          </div>
          
          <div class="form-group">
            <label for="lastName">Last Name</label>
            <input 
              type="text" 
              id="lastName" 
              v-model="driver.lastName" 
              placeholder="Enter your last name"
              required
            />
          </div>
        </div>
        
        <div class="form-group">
          <label for="email">Email</label>
          <input 
            type="email" 
            id="email" 
            v-model="driver.email" 
            placeholder="Enter your email address"
            required
          />
        </div>
        
        <div class="form-group">
          <label for="phoneNumber">Phone Number</label>
          <input 
            type="tel" 
            id="phoneNumber" 
            v-model="driver.phoneNumber" 
            placeholder="Enter your phone number"
            required
          />
        </div>
        
        <div class="form-group">
          <label for="password">Create Password</label>
          <input 
            type="password" 
            id="password" 
            v-model="driver.password" 
            placeholder="Create a secure password"
            required
          />
        </div>
        
        <div class="form-group">
          <label for="confirmPassword">Confirm Password</label>
          <input 
            type="password" 
            id="confirmPassword" 
            v-model="confirmPassword" 
            placeholder="Confirm your password"
            required
          />
          <span class="error-message" v-if="passwordError">{{ passwordError }}</span>
        </div>
        
        <button type="submit" class="submit-button">Create Driver Account</button>
      </form>
      
      <p class="login-link">
        Already have an account? <a href="#">Log in here</a>
      </p>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent, ref } from 'vue';
  import { Driver } from '../types/Driver'; // Adjust the import path as necessary
  
  export default defineComponent({
    name: 'DriverForm',
    setup() {
      const driver = ref<Driver>({
        firstName: '',
        lastName: '',
        email: '',
        phoneNumber: '',
        password: ''
      });
      
      const confirmPassword = ref('');
      const passwordError = ref('');
      
      const submitForm = () => {
        if (driver.value.password !== confirmPassword.value) {
          passwordError.value = 'Passwords do not match';
          return;
        }
        
        passwordError.value = '';
        console.log('Driver Form Submitted:', driver.value);
        // Here you would typically make an API call to register the driver
        alert('Driver registration successful!');
      };
      
      return {
        driver,
        confirmPassword,
        passwordError,
        submitForm
      };
    }
  });
  </script>