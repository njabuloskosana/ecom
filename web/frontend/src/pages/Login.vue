<template>
  <div class="login-container">
    <Loader :isLoading="isLoading" />
    <NotificationModal
      :isVisible="isErrorModalVisible"
      title="Login Failed"
      :message="errorMessage"
      cancelText="Cancel"
      confirmText="Retry"
      @cancel="hideErrorModal"
      @confirm="retryLogin"
    />
    <div  v-if="!isLoggedIn && !isLoading" class="login-card">
      <h2 class="login-title">Sign In</h2>
      <form @submit.prevent="handleLogin" class="login-form">
        <!-- Email Input -->
        <div class="input-group">
          <label for="email" class="input-label">Email Address</label>
          <input
            type="email"
            id="email"
            v-model="email"
            placeholder="Enter your email"
            required
            class="input-field"
          />
        </div>

        <!-- Password Input -->
        <div class="input-group">
          <label for="password" class="input-label">Password</label>
          <input
            type="password"
            id="password"
            v-model="password"
            placeholder="Enter your password"
            required
            class="input-field"
          />
        </div>

        <!-- Forgot Password Link -->
        <div class="forgot-password">
          <router-link to="/" class="forgot-password-link"
            >Discover More Rewards</router-link
          >
        </div>

        <!-- Login Button -->
        <button type="submit" class="login-button">Login</button>

        <!-- Signup Link -->
        <div class="signup-link">
          Don't have an account? 
          <router-link to="/signup" class="signup-link-text">Sign Up</router-link>
        </div>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import axios from "axios";
import { useRouter } from "vue-router";
import Loader from "../components/Loader.vue";
import NotificationModal from "../components/Notification.vue";// Ensure the path is correct
import Config from "@/config";
import type { LoginResponse } from "../interfaces/index";



export default defineComponent({
  components: {
    NotificationModal, // Register the NotificationModal component
    Loader, // Register the Loader component
  },
  setup() {
    const email = ref<string>("");
    const password = ref<string>("");
    const isLoggedIn = ref<boolean>(false);
    const isLoading = ref<boolean>(false);
    const isErrorModalVisible = ref<boolean>(false);
    const errorMessage = ref<string>("Invalid credentials or server error.");
    const baseUrl = Config.BASE_API_URL_PRODUCTION;
    const router = useRouter();

    // Query the local storage for the JWT token and user info then navigate the user to the correct UI
    const checkLoginStatus = () => {
      const jwtToken = localStorage.getItem("jwtToken");
      const userInfo = localStorage.getItem("userInfo");
      const userProfile = localStorage.getItem("userProfile");

      if (jwtToken && userInfo && userProfile) {
      isLoggedIn.value = true;
      // Navigate to the correct UI based on roles and permissions
      router.push("/home-elarduspark?tenant=elarduspark");
      }
    };

    // Call the function to check login status on component mount
    checkLoginStatus();

    // Method to handle login
    const handleLogin = async () => {
      if (!email.value || !password.value) {
        alert("Please fill in all fields.");
        return;
      }

      try {
        isLoading.value = true; // Show the loader

        const url = baseUrl + "/User/Login";
        const response: any = await axios.post(url, {
          email: email.value,
          username: email.value,
          password: password.value,
        });

       const user : LoginResponse = response.data;

        // Handle successful login
        console.log("Login Successful:", response);

        // Store JWT token in local storage
        localStorage.setItem("jwtToken", user.jwtToken);
        localStorage.setItem("userInfo", JSON.stringify(user.userInformation));
        localStorage.setItem("userProfile", JSON.stringify(user.userProfile));

        // Set login status and navigate to welcome page
        isLoggedIn.value = true;
        router.push("/home-elarduspark?tenant=elarduspark");
       console.log("Navigate to the correct UI based on your roles and permissions:");
      } catch (error) {
        errorMessage.value = "Invalid credentials or server error.";
        isErrorModalVisible.value = true; // Show the error modal
      } finally {
        isLoading.value = false; // Hide the loader
      }
    };

    const hideErrorModal = () => {
  isErrorModalVisible.value = false;
};

// Retry login
const retryLogin = () => {
  hideErrorModal(); // Hide the modal
  handleLogin(); // Retry the login process
};

    return {
      email,
      password,
      isLoggedIn,
      isLoading,
      isErrorModalVisible,
      errorMessage,
      handleLogin,
      retryLogin,
      hideErrorModal
    };
  },
});


</script>

<style scoped>
/* General Styles */
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  min-width: 100%;
  background-color: #0056b3;
  background-image: url("../assets/login-bkgnd.jpg");
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-size: cover;
  background-position: center center;
}

.login-card {
  background: #ffffff;
  padding: 30px 40px;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  width: 90%;
  max-width: 500px;
}

.login-title {
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

.forgot-password {
  text-align: right;
  margin-bottom: 15px;
}

.forgot-password-link {
  color: #007bff;
  text-decoration: none;
  font-size: 14px;
}

.forgot-password-link:hover {
  text-decoration: underline;
}

.login-button {
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

.login-button:hover {
  background-color: #0056b3;
}

.signup-link {
  text-align: center;
  margin-top: 15px;
  font-size: 14px;
}

.signup-link-text {
  color: #007bff;
  text-decoration: none;
}

.signup-link-text:hover {
  text-decoration: underline;
}

/* Responsive Design */
@media (max-width: 600px) {
  .login-card {
    padding: 20px;
  }

  .login-title {
    font-size: 20px;
  }

  .input-field {
    font-size: 14px;
  }

  .login-button {
    font-size: 14px;
  }
}
</style>