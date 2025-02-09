<template>
    <div class="signup-container">
      <!-- Loader -->
      <Loader :isLoading="isLoading" />
  
      <!-- Error Modal -->
      <NotificationModal
        :isVisible="isErrorModalVisible"
        :title="errorMessageTitle"
        :message="errorMessage"
        cancelText="Cancel"
        confirmText="Retry"
        @cancel="hideErrorModal"
        @confirm="retrySignup"
      />
  
      <!-- Signup Form -->
      <div v-if="!isSignedUp && !isLoading" class="signup-card">
        <h2 class="signup-title">Sign Up</h2>
        <form @submit.prevent="handleSignup" class="signup-form" autocomplete="on">
          <!-- Name Input -->
          <div class="input-group">
            <label for="firstName" class="input-label">Name *</label>
            <input
              type="text"
              id="firstName"
              name="firstName"
              v-model="user.profile.firstName"
              placeholder="Enter your first name"
              required
              class="input-field"
              autocomplete="given-name"
            />
          </div>
  
          <!-- Surname Input -->
          <div class="input-group">
            <label for="lastName" class="input-label">Surname *</label>
            <input
              type="text"
              id="lastName"
              name="lastName"
              v-model="user.profile.lastName"
              placeholder="Enter your surname"
              required
              class="input-field"
              autocomplete="family-name"
            />
          </div>
  
          <!-- Email Input -->
          <div class="input-group">
            <label for="email" class="input-label">Email *</label>
            <input
              type="email"
              id="email"
              name="email"
              v-model="user.userEmail"
              placeholder="Enter your email"
              required
              class="input-field"
              autocomplete="email"
            />
          </div>
  
          <!-- Phone Number Input -->
          <div class="input-group">
            <label for="phone" class="input-label">Phone Number *</label>
            <input
              type="tel"
              id="phone"
              name="phone"
              v-model="user.profile.phoneNumber"
              placeholder="Enter your phone number"
              required
              class="input-field"
              autocomplete="tel"
            />
          </div>
  
          <!-- Password Input -->
          <div class="input-group">
            <label for="password" class="input-label">Password *</label>
            <input
              type="password"
              id="password"
              name="password"
              v-model="user.password"
              placeholder="Enter your password"
              required
              class="input-field"
              autocomplete="new-password"
            />
          </div>
  
          <!-- Confirm Password Input (Optional) -->
          <div class="input-group">
            <label for="confirmPassword" class="input-label">Confirm Password *</label>
            <input
              type="password"
              id="confirmPassword"
              name="confirmPassword"
              placeholder="Re-enter your password"
              required
              class="input-field"
              autocomplete="new-password"
            />
          </div>
  
          <!-- Signup Button -->
          <button type="submit" class="signup-button">SIGNUP</button>
  
          <!-- Login Link -->
          <div class="login-link">
            Already have an account? 
            <router-link to="/login" class="login-link-text">Login</router-link>
          </div>
        </form>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref } from "vue";
  import axios from "axios";
  import { useRouter } from "vue-router";
  import Loader from "../components/Loader.vue";
  import NotificationModal from "../components/Notification.vue"; 
  import Config from "../config/index"// Ensure the path is correct
import type { User } from "@/interfaces";
  

  
  // Reactive state variables
  const user = ref<User>({
    userEmail: "",
    userName: "",
    password: "",
    notificationPreferences: [0,2],
    roleId: [2], // Default role ID (e.g., Customer)
    profile: {
      id: 0,
      userId: 0,
      firstName: "",
      lastName: "",
      phoneNumber: "",
      alternativeNumber: "",
    },
  });
  
  const isLoading = ref<boolean>(false);
  const isErrorModalVisible = ref<boolean>(false);
  const errorMessage = ref<string>("");
  const errorMessageTitle = ref<string>("Signup Failed");
  const isSignedUp = ref<boolean>(false);
  
  const router = useRouter();
  const baseUrl = Config.BASE_API_URL_PRODUCTION;
  
  // Handle signup
  const handleSignup = async () => {
    if (
      !user.value.userEmail ||
      !user.value.password ||
      !user.value.profile.firstName ||
      !user.value.profile.lastName ||
      !user.value.profile.phoneNumber
    ) {
      alert("Please fill in all required fields.");
      return;
    }
  
    try {
      isLoading.value = true; // Show the loader
      const url = baseUrl + "/User/Admin";
    user.value.profile.alternativeNumber = user.value.profile.phoneNumber; // Set alternativeNumber to phoneNumber
    user.value.userEmail = user.value.userEmail; // Convert email to lowercase
    user.value.userName = user.value.userEmail; // Set username to email
    const response = await axios.post(url, {
      ...user.value,
      userName: user.value.userEmail, // Use email as username
    });
  
      isSignedUp.value = true;
  
      // Redirect to login page after signup
      router.push("/login");
    } catch (error) {
      console.error("Failed to signup", error);
      errorMessage.value = "Invalid details or server error.";
      errorMessageTitle.value = "Signup Failed";
      isErrorModalVisible.value = true; // Show the error modal
    } finally {
      isLoading.value = false; // Hide the loader
    }
  };
  
  // Hide error modal
  const hideErrorModal = () => {
    isErrorModalVisible.value = false;
  };
  
  // Retry signup
  const retrySignup = () => {
    hideErrorModal(); // Hide the modal
    handleSignup(); // Retry the signup process
  };
  </script>
  
  <style scoped>
  /* General Styles */
  .signup-container {
    color: #333;
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    background-image: url("../assets/login-bkgnd.jpg");
    background-repeat: no-repeat;
  background-attachment: fixed;
  background-size: cover;
  background-position: center center;
  }
  
  .signup-card {
    background: #ffffff;
    padding: 30px 40px;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    width: 90%;
    max-width: 500px;
  }
  
  .signup-title {
    font-size: 24px;
    font-weight: bold;
    margin-bottom: 20px;
    color: #333;
    text-align: center;
  }
  
  .input-group {
    margin-bottom: 15px;
  }
  
  .input-label {
    display: block;
    font-size: 14px;
    color: #666;
    margin-bottom: 5px;
  }
  
  .input-field {
    width: 100%;
    padding: 10px;
    font-size: 16px;
    border: 1px solid #ddd;
    border-radius: 4px;
    transition: border-color 0.3s ease;
  }
  
  .input-field:focus {
    border-color: #007bff;
    outline: none;
  }
  
  .signup-button {
    width: 100%;
    padding: 12px;
    font-size: 16px;
    font-weight: bold;
    color: #fff;
    background-color: #007bff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s ease;
  }
  
  .signup-button:hover {
    background-color: #0056b3;
  }
  
  .login-link {
    text-align: center;
    margin-top: 15px;
    font-size: 14px;
  }
  
  .login-link-text {
    color: #007bff;
    text-decoration: none;
  }
  
  .login-link-text:hover {
    text-decoration: underline;
  }
  
  /* Responsive Design */
  @media (max-width: 600px) {
    .signup-card {
      padding: 20px;
    }
    .signup-title {
      font-size: 20px;
    }
    .input-field {
      font-size: 14px;
    }
    .signup-button {
      font-size: 14px;
    }
  }
  </style>