<template>
    <div class="form-container">
      <h2>Bike Owner Registration</h2>
      <p class="subtitle">Register as a bike owner to list your bikes and find qualified drivers</p>
      
      <form @submit.prevent="submitForm">
        <div class="form-row">
          <div class="form-group">
            <label for="fullName">Full Name</label>
            <input 
              type="text" 
              id="fullName" 
              v-model="bikeOwner.fullName" 
              placeholder="Enter your full name"
              required
            />
          </div>
          
          <div class="form-group">
            <label for="businessName">Business Name (Optional)</label>
            <input 
              type="text" 
              id="businessName" 
              v-model="bikeOwner.businessName" 
              placeholder="Enter your business name"
            />
          </div>
        </div>
        
        <div class="form-group">
          <label for="email">Email</label>
          <input 
            type="email" 
            id="email" 
            v-model="bikeOwner.email" 
            placeholder="Enter your email address"
            required
          />
        </div>
        
        <div class="form-group">
          <label for="phoneNumber">Phone Number</label>
          <input 
            type="tel" 
            id="phoneNumber" 
            v-model="bikeOwner.phoneNumber" 
            placeholder="Enter your phone number"
            required
          />
        </div>
        
        <div class="form-row">
          <div class="form-group">
            <label for="location">Primary Location</label>
            <input 
              type="text" 
              id="location" 
              v-model="bikeOwner.location" 
              placeholder="City or area"
              required
            />
          </div>
          
          <div class="form-group">
            <label for="numberOfBikes">Number of Bikes</label>
            <select 
              id="numberOfBikes" 
              v-model="bikeOwner.numberOfBikes"
              required
            >
              <option value="" disabled selected>How many bikes do you own?</option>
              <option v-for="n in 10" :key="n" :value="n">{{ n }}</option>
              <option value="11">More than 10</option>
            </select>
          </div>
        </div>
        
        <div class="form-group">
          <label for="password">Create Password</label>
          <input 
            type="password" 
            id="password" 
            v-model="bikeOwner.password" 
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
        
        <button type="submit" class="submit-button">Create Owner Account</button>
      </form>
      
      <p class="login-link">
        Already have an account? <a href="#">Log in here</a>
      </p>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent, ref } from 'vue';
  import { BikeOwner } from '../types/BikeOwner'; // Adjust the import path as necessary
  
  export default defineComponent({
    name: 'BikeOwnerForm',
    setup() {
      const bikeOwner = ref<BikeOwner>({
        fullName: '',
        businessName: '',
        email: '',
        phoneNumber: '',
        location: '',
        numberOfBikes: 0,
        password: ''
      });
      
      const confirmPassword = ref('');
      const passwordError = ref('');
      
      const submitForm = () => {
        if (bikeOwner.value.password !== confirmPassword.value) {
          passwordError.value = 'Passwords do not match';
          return;
        }
        
        passwordError.value = '';
        console.log('Bike Owner Form Submitted:', bikeOwner.value);
        // Here you would typically make an API call to register the bike owner
        alert('Bike owner registration successful!');
      };
      
      return {
        bikeOwner,
        confirmPassword,
        passwordError,
        submitForm
      };
    }
  });
  </script>